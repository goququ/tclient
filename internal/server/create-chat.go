package server

import (
	"fmt"
	"net/http"

	"github.com/anonyindian/gotgproto"
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

func createTelegramChat(c *gotgproto.Client, p *CreateChatParams) (*chatsBundle, error) {
	creationContext := c.CreateContext()
	chat, err := creationContext.CreateChat(p.Title, []int64{})
	if err != nil {
		return nil, fmt.Errorf("unable to create chat")
	}

	storage.AddPeer(chat.GetID(), storage.DefaultAccessHash, storage.TypeChat, storage.DefaultUsername)

	fullChat, err := creationContext.GetChat(chat.GetID())
	if err != nil {
		return nil, fmt.Errorf("unable to create chat")
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

func (h AppHandlers) createChat(c *gin.Context) {
	params, err := getCreateChatParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, createApiError(err))
	}

	chatBundle, err := createTelegramChat(h.client, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createApiError(err))
		return
	}

	link, err := getChatLink(chatBundle.fullChat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, createApiError(err))
		return
	}

	res := gin.H{
		"id":    chatBundle.fullChat.GetID(),
		"title": chatBundle.chat.GetTitle(),
		"link":  link,
	}

	c.JSON(http.StatusOK, res)
}
