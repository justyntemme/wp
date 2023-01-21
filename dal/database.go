package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

func Query(ctx context.Context, pipeline []bson.M) (result string, err error) {

	// Get a handle to the orders collection
	orders := client.Database("wp").Collection("users")

	// Execute the aggregation pipeline
	cursor, err := orders.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the result set
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
	return result, err
}
