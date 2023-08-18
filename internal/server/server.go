package server

import (
	"fmt"
	"net/http"
	"tclient/internal/config"
	"tclient/internal/db"

	"github.com/anonyindian/gotgproto"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.AppConfig
	Client *gotgproto.Client
	Db     *db.DBClient
}

func (app Application) Run() error {
	router := gin.Default()

	router.GET("/create", app.createChat)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", app.Config.Port),
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	fmt.Printf("Server is running on port %v", app.Config.Port)
	return nil
}
