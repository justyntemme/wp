package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type repsonse struct {
	Id      string `json:"_id"`
	User_id string `json:"user_id"`
	Club_id string `json:"club_id"`
}

func init() {
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://jupiter:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

func query(ctx context.Context, pipeline []bson.D) (result string, err error) {
	var results []bson.M

	// Get a handle to the orders collection
	orders := client.Database("wp").Collection("votes")

	// Execute the aggregation pipeline
	cursor, err := orders.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &results)

	output, err := json.MarshalIndent(results, "", " ")
	return string(output), err
}

func DalGetVotesByClubId(ctx context.Context, id string) (result string, err error) {
	var pipeline []bson.D

	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "cuid", Value: id},
		},
		},
	}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "cuid", Value: 1},
		}},
	}

	pipeline = append(pipeline, matchStage, projectStage)

	result, err = query(ctx, pipeline)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	mapBytes, err := json.Marshal(result)
	mapString := string(mapBytes)
	return mapString, err

}
