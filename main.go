package main

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
}

