package main

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	Collection *mongo.Collection
}

var db = Database{Collection: connect()}

func main() {
	println("Main Function")
}
