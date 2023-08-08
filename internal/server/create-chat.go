package server

import (
	"log"
	"net/http"

	"github.com/anonyindian/gotgproto/storage"
	"github.com/gin-gonic/gin"
	"github.com/gotd/td/tg"
)

func (h AppHandlers) createChat(c *gin.Context) {
	creationContext := h.client.CreateContext()
	chat, err := creationContext.CreateChat("test", []int64{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}

	storage.AddPeer(chat.GetID(), storage.DefaultAccessHash, storage.TypeChat, storage.DefaultUsername)

	fullChat, err := creationContext.GetChat(chat.GetID())
	if err != nil {
		log.Println("Unable to get full chat")
		c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}

	invite, ok := fullChat.GetExportedInvite()
	if !ok {
		log.Println("Unable to get invite link")
		c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}

	var link string
	if exportedInvite, ok := invite.(*tg.ChatInviteExported); ok {
		link = exportedInvite.Link
	} else {
		log.Println("Unable to get access to invite link field")
		c.AbortWithError(http.StatusInternalServerError, err).SetType(gin.ErrorTypePrivate)
		return
	}

	res := gin.H{
		"chatID":   chat.GetID(),
		"title":    chat.GetTitle(),
		"joinLink": link,
	}

	c.JSON(http.StatusOK, res)
}
