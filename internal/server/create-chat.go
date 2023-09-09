package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/goququ/tclient/internal/schemas"

	"github.com/anonyindian/gotgproto"
	"github.com/anonyindian/gotgproto/ext"
	"github.com/anonyindian/gotgproto/storage"
	"github.com/gin-gonic/gin"
	"github.com/gotd/td/tg"
)

type CreateChatParams struct {
	Title string `form:"title" binding:"required"`
	Sap   string `form:"sap" binding:"required"`
	Admin string `form:"admin" binding:"required"`
}

func getCreateChatParams(c *gin.Context) (*CreateChatParams, error) {
	var params CreateChatParams

	if err := c.ShouldBindQuery(&params); err != nil {
		return nil, err
	}

	return &params, nil
}

type chatsBundle struct {
	chat     *tg.Chat
	fullChat tg.ChatFullClass
}

func createTelegramChat(c *gotgproto.Client, p *CreateChatParams, retries int, retryDelaySec int) (*chatsBundle, error) {
	creationContext := c.CreateContext()

	var adminId int64
	if effChat, err := creationContext.ResolveUsername(p.Admin); err != nil {
		return nil, fmt.Errorf("unable to resolve username '%v', error: %v", p.Admin, err.Error())
	} else {
		adminId = effChat.GetID()
	}

	log.Printf("Creating chat '%v'", p.Title)
	var chat *tg.Chat
	for attept := 0; chat == nil && attept < retries; attept++ {
		newChat, err := creationContext.CreateChat(p.Title, []int64{adminId})
		if err != nil {
			log.Printf("Attempt: %v, Error: %v", attept+1, err.Error())
			time.Sleep(time.Second * time.Duration(retryDelaySec))
			continue
		}
		chat = newChat
	}

	if chat == nil {
		return nil, fmt.Errorf("unable to create chat")
	}
	chatId := chat.GetID()

	log.Printf("Created chat with id: %v", chatId)

	storage.AddPeer(chatId, storage.DefaultAccessHash, storage.TypeChat, storage.DefaultUsername)

	fullChat, err := creationContext.GetChat(chatId)
	if err != nil {
		return nil, fmt.Errorf("unable to get full chat :%v", err.Error())
	}

	_, err = creationContext.PromoteChatMember(chatId, adminId, &ext.EditAdminOpts{})
	if err != nil {
		log.Printf("WARNING: Unable to promote user %v to admin. Error: %v", p.Admin, err.Error())
	}

	return &chatsBundle{
		chat:     chat,
		fullChat: fullChat,
	}, nil
}

func getChatLink(c tg.ChatFullClass) (string, error) {
	invite, ok := c.GetExportedInvite()
	if !ok {
		return "", fmt.Errorf("unable to extract link from chat")
	}

	var link string
	if exportedInvite, ok := invite.(*tg.ChatInviteExported); ok {
		link = exportedInvite.Link
	} else {
		return "", fmt.Errorf("unable to get access to invite link field")
	}

	return link, nil
}

func (h *Application) createChat(c *gin.Context) {
	params, err := getCreateChatParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, createApiError(err))
	}

	chatBundle, err := createTelegramChat(h.Client, params, h.Config.RetryCount, h.Config.RetryDelaySec)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createApiError(err))
		return
	}

	link, err := getChatLink(chatBundle.fullChat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createApiError(err))
		return
	}

	record := schemas.ChatRecord{
		Id:    chatBundle.chat.GetID(),
		Link:  link,
		Admin: params.Admin,
		Sap:   params.Sap,
		Title: params.Title,
	}

	if err := h.Db.NewChat(&record); err != nil {
		log.Println("Unable to save chat to mongo &#v", record)
	}

	res := gin.H{
		"link":   link,
		"status": "success",
	}

	c.JSON(http.StatusOK, res)
}
