package server

import (
	"fmt"
	"net/http"

	"github.com/goququ/tclient/internal/config"
	"github.com/goququ/tclient/internal/db"

	"github.com/anonyindian/gotgproto"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.AppConfig
	Client *gotgproto.Client
	Db     *db.DBClient
}

func (app Application) Run() error {
	if app.Config.EnvMode == config.Production {
		gin.SetMode(gin.ReleaseMode)
	}

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
