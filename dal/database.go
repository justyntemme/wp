package dal

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
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://jupiter:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

func query(ctx context.Context, pipeline []bson.D) (result string, err error) {

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

func GetVotesByClubId(ctx context.Context, id string) (result string, err error) {

	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "$cuid", Value: id},
		},
		},
	}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "$uuid", Value: 1},
		}},
	}
	// pipeline = append(pipeline, bson.D{{"$match", bson.D{{"uuid", bson.D{{"$in", []string{id}}}}, bson.A{}}}})
	// match := bson.D{
	// 	"$match",
	// 	bson.E{"uuid", id}}
	var pipeline []bson.D

	pipeline = append(pipeline, matchStage, projectStage)

	return query(ctx, pipeline)

}
