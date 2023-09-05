package db

import (
	"context"
	"log"
	"time"

	"github.com/goququ/tclient/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	Client *mongo.Client
}

func New(c *config.AppConfig) (*DBClient, error) {
	log.Printf("MONGO: start applying connection uri: %v", c.MongoConnectString)

	clientOptions := options.Client().ApplyURI(c.MongoConnectString)
	log.Printf("MONGO: applying complete")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("MONGO: creating mongo db client")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	log.Printf("MONGO: Successfully created mongo db client")

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Printf("MONGO: Checking connection to db")
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}
	log.Printf("MONGO: Connection checked. OK")

	log.Println("Connected to MongoDB!")

	return &DBClient{
		Client: client,
	}, nil
}

func (dc *DBClient) Disconnect() {
	err := dc.Client.Disconnect(context.Background())
	if err != nil {
		log.Print("Failed to close DB connection")
	}
}
