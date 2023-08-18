package db

import (
	"context"
	"fmt"
	"tclient/internal/schemas"

	"go.mongodb.org/mongo-driver/mongo"
)

func (c *DBClient) getChatsCollection() *mongo.Collection {
	collection := c.client.Database("tclient").Collection("chats")

	return collection
}

func (c *DBClient) NewChat(r *schemas.ChatRecord) error {
	collection := c.getChatsCollection()

	insertResult, err := collection.InsertOne(context.TODO(), r)
	if err != nil {
		return err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}
