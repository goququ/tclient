package server

import (
	"fmt"
	"tclient/internal/config"

	"github.com/anonyindian/gotgproto"
	"github.com/gin-gonic/gin"
)

type AppHandlers struct {
	config *config.AppConfig
	client *gotgproto.Client
}

func Run(c *config.AppConfig, tc *gotgproto.Client) {
	r := gin.Default()

	handlers := AppHandlers{c, tc}

	r.GET("/create", handlers.createChat)

	err := r.Run(fmt.Sprintf(":%v", c.Port))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is running on port %v", c.Port)
}
