package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env string

const (
	Production  Env = "production"
	Development Env = "development"
)

type AppConfig struct {
	// telegram client variables
	Phone   string
	AppId   int
	ApiHash string
	// api variables
	Port          int
	RetryCount    int
	RetryDelaySec int

	MongoConnectString string

	EnvMode Env
}

var env Env

func Read() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to find .env file")
	}

	if env == "" {
		env = Development
	}

	log.Printf("Application runs in %v mode", env)

	phone := os.Getenv("TGA_PHONE")
	if len(phone) == 0 {
		return nil, fmt.Errorf("No 'TGA_PHONE' variable defined")
	}

	appId, err := strconv.Atoi(os.Getenv("TGA_APP_ID"))
	if err != nil {
		return nil, fmt.Errorf("invalid value of 'TGA_APP_ID'")
	}

	apiHash := os.Getenv("TGA_API_HASH")
	if len(apiHash) == 0 {
		return nil, fmt.Errorf("No 'TGA_API_HASH' variable defined")
	}

	mongoConnectString := os.Getenv("TGA_MONGO_CONNECTION")
	if len(mongoConnectString) == 0 {
		return nil, fmt.Errorf("No 'TGA_MONGO_CONNECTION' variable defined")
	}

	port, err := strconv.Atoi(os.Getenv("TGA_PORT"))
	if err != nil {
		port = 9001
		log.Printf("Unable to find application port, fallback to %v", port)
	}

	retryCount, err := strconv.Atoi(os.Getenv("TGA_RETRY_COUNT"))
	if err != nil {
		log.Print("Invalid value of 'TGA_RETRY_COUNT'")
	}

	retryDelaySeconds := 5
	if val, err := strconv.Atoi(os.Getenv("TGA_RETRY_DELEY_SECONDS")); err != nil {
		log.Print("Invalid value of 'TGA_RETRY_DELEY_SECONDS'")
	} else {
		retryDelaySeconds = val
	}
	log.Printf("value of 'TGA_RETRY_DELEY_SECONDS' %v", retryDelaySeconds)

	return &AppConfig{
		Phone:              phone,
		AppId:              appId,
		ApiHash:            apiHash,
		Port:               port,
		MongoConnectString: mongoConnectString,
		EnvMode:            env,
		RetryCount:         retryCount,
		RetryDelaySec:      retryDelaySeconds,
	}, nil
}
