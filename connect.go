package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func connect() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb+srv://user-service:RCIcEjYEtMJTuhHS@platform-test-5fvzo.gcp.mongodb.net/test?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Organisations").Collection("organisations")
	return collection
}
