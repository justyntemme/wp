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

// Writes
func AddVoteForClub(ctx context.Context, cuid string, userId string) error {
	vote := bson.M{
		"cuid": cuid,
		"uuid": userId,
	}
	// Get a handle to the orders collection
	votesCollection := client.Database("wp").Collection("votes")
	_, err := votesCollection.InsertOne(ctx, vote)
	if err != nil {
		return fmt.Errorf("Error adding vote: %v", err)
	}
	return nil
}

func query(ctx context.Context, pipeline []bson.D) (result string, err error) {
	var results []bson.M

	// Get a handle to the orders collection
	collection := client.Database("wp").Collection("votes")

	// Execute the aggregation pipeline
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Fatal(err)
	}
	cursor.All(ctx, &results)
	output, err := json.Marshal(results)

	return string(output), err
}

// Reads
func GetVotesByClubId(ctx context.Context, id string) (result string, err error) {
	var pipeline []bson.D

	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "cuid", Value: id},
		},
		},
	}
	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "uuid", Value: 1},
		}},
	}

	pipeline = append(pipeline, matchStage, projectStage)

	result, err = query(ctx, pipeline)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return result, err

}

func GetVotesByUserId(ctx context.Context, id string) (result string, err error) {
	var pipeline []bson.D

	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "uuid", Value: id},
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

	return result, err

}

func GetClubsWithMostVotes(ctx context.Context) (result string, err error) {
	var pipeline []bson.D

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$cuid"},
			{Key: "count", Value: bson.D{
				{Key: "$sum", Value: 1},
			}},
		}},
	}

	sortStage := bson.D{
		{Key: "$sort", Value: bson.D{
			{Key: "count", Value: -1},
		}},
	}

	pipeline = append(pipeline, groupStage, sortStage)

	result, err = query(ctx, pipeline)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	return result, err
}
