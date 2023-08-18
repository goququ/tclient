package db

import (
	"context"
	"log"
	"tclient/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	client *mongo.Client
}

func New(c *config.AppConfig) (*DBClient, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(c.MongoConnectString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return &DBClient{
		client,
	}, nil
}
