package main

import (
	"log"

	"github.com/goququ/tclient/internal/client"
	"github.com/goququ/tclient/internal/config"
	"github.com/goququ/tclient/internal/db"
	"github.com/goququ/tclient/internal/server"
)

func Main() error {
	log.Print("Start reading config")
	appConfig, err := config.Read()
	if err != nil {
		return err
	}
	log.Print("Config ready")

	log.Print("Start creating the telegram client")
	tgClient, err := client.Create(appConfig)
	if err != nil {
		return err
	}
	log.Print("Telegram client created")

	log.Print("Creating db client")
	dbClient, err := db.New(appConfig)
	if err != nil {
		return err
	}
	defer dbClient.Disconnect()
	log.Print("DB client created")

	log.Print("Creating server")
	app := server.Application{
		Config: appConfig,
		Client: tgClient,
		Db:     dbClient,
	}

	err = app.Run()

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Main()
	if err != nil {
		log.Fatal(err)
	}
}
