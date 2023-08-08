package config

import (
	"errors"
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
}

func Read() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	phone := os.Getenv("TGA_PHONE")
	if len(phone) == 0 {
		panic(errors.New("No 'TGA_PHONE' variable defined"))
	}

	appId, err := strconv.Atoi(os.Getenv("TGA_APP_ID"))
	if err != nil {
		panic(errors.New("invalid value of 'TGA_APP_ID'"))
	}

	apiHash := os.Getenv("TGA_API_HASH")
	if len(apiHash) == 0 {
		panic(errors.New("No 'TGA_API_HASH' variable defined"))
	}

	port, err := strconv.Atoi(os.Getenv("TGA_PORT"))
	if err != nil {
		port = 9001
	}

	return &AppConfig{
		Phone:   phone,
		AppId:   appId,
		ApiHash: apiHash,
		Port:    port,
	}
}
