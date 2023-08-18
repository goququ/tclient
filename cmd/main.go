package main

import (
	"log"
	"tclient/internal/client"
	"tclient/internal/config"
	"tclient/internal/db"
	"tclient/internal/server"
)

func main() {
	appConfig := config.Read()
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
