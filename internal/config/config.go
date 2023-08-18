package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	// telegram client variables
	Phone   string
	AppId   int
	ApiHash string
	// api variables
	Port int

	MongoConnectString string
}

func Read() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	phone := os.Getenv("TGA_PHONE")
	if len(phone) == 0 {
		log.Fatal("No 'TGA_PHONE' variable defined")
	}

	appId, err := strconv.Atoi(os.Getenv("TGA_APP_ID"))
	if err != nil {
		log.Fatal("invalid value of 'TGA_APP_ID'")
	}

	apiHash := os.Getenv("TGA_API_HASH")
	if len(apiHash) == 0 {
		log.Fatal("No 'TGA_API_HASH' variable defined")
	}

	mongoConnectString := os.Getenv("TGA_MONGO_CONNECTION")
	if len(mongoConnectString) == 0 {
		log.Fatal("No 'TGA_MONGO_CONNECTION' variable defined")
	}

	port, err := strconv.Atoi(os.Getenv("TGA_PORT"))
	if err != nil {
		port = 9001
		log.Printf("Unable to find application port, fallback to %v", port)
	}

	return &AppConfig{
		Phone:              phone,
		AppId:              appId,
		ApiHash:            apiHash,
		Port:               port,
		MongoConnectString: mongoConnectString,
	}
}
