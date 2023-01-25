package service

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

func GetClubsById(ctx context.Context, id string) (result string, err error) {

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

	var pipeline []bson.D

	pipeline = append(pipeline, matchStage, projectStage)

	return query(ctx, pipeline)

}

func GetAllClubsNearMe(ctx context.Context, lat float64, lon float64, limit int32) (result string, err error) {
	matchStage := bson.D{
		{Key: "$geoNear", Value: bson.D{
			{Key: "near", Value: bson.D{
				{Key: "type", Value: "Point"},
				{Key: "coordinates", Value: []float64{lon, lat}},
			}},
			{Key: "maxDistance", Value: limit},
			{Key: "spherical", Value: true},
			{Key: "distanceField", Value: "distance"},
		}},
	}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "cuid", Value: 1},
			{Key: "distance", Value: 1},
		}},
	}
	var pipeline []bson.D

	pipeline = append(pipeline, matchStage, projectStage)

	return query(ctx, pipeline)
}
