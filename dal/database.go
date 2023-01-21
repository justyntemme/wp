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

	return client
}

func Query(ctx context.Context, pipeline []bson.M) (result string, err error) {
	// Define the aggregation pipeline
	// pipeline := []bson.M{
	// 	{"$match": bson.M{"Id": "A"}},
	// 	{"$group": bson.M{"_id": "$cust_id", "total": bson.M{"$sum": "$amount"}}},
	// 	{"$sort": bson.M{"total": -1}},
	// }

	// Get a handle to the orders collection
	orders := client.Database("mydb").Collection("orders")

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
