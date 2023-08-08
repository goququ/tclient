package main

import (
	"tclient/internal/client"
	"tclient/internal/config"
	"tclient/internal/server"
)

func main() {
	appConfig := config.Read()
	tgClient, err := client.Create(appConfig)
	if err != nil {
		panic(err)
	}
	// client.Idle()

	server.Run(appConfig, tgClient)
}
