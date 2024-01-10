package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1")
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify that the client has connected successfully
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func VoicesCollection() *mongo.Collection {
	client := InitMongoDB()

	if client == nil {
		log.Fatal("Cannot Connect to the Database")
		return nil
	}

	database := client.Database("VoiceForge")
    collection := database.Collection("voices")

	keys := bson.D{
		{"id", 1},
	}

	indexModel := mongo.IndexModel{
		Keys: keys,
		Options: options.Index().SetUnique(false),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
    if err != nil {
        log.Fatalln("Error creating index:", err)
    }

	

	return collection

}
