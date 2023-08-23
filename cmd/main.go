package main

import (
	"log"

	"github.com/goququ/tclient/internal/client"
	"github.com/goququ/tclient/internal/config"
	"github.com/goququ/tclient/internal/db"
	"github.com/goququ/tclient/internal/server"
)

func main() {
	appConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	tgClient, err := client.Create(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	dbClient, err := db.New(appConfig)
	if err != nil {
		log.Fatal(err)
	}

	app := server.Application{
		Config: appConfig,
		Client: tgClient,
		Db:     dbClient,
	}

	app.Run()
}
