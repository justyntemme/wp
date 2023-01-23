package service

import (
	"context"
	"log"

	dal "github.com/justyntemme/wp/dal"
	"gopkg.in/mgo.v2/bson"
)

// @microgen middleware, logging, grpc, http, recovering
// @protobuf github.com/justyntemme/wp/vote/proto
type VoteService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetVoteById(ctx context.Context, id string) (result string, err error)
	GetVotesByUserId(ctx context.Context, id string) (result string, err error)
	GetVotesByClubId(ctx context.Context, id string) (result string, err error)
}

func GetVoteById(ctx context.Context, id string) (result string, err error) {

	//var pipeline []bson.M

	// Define the aggregation pipeline
	// pipeline = []bson.M{
	// 	{"$match": bson.M{"Id": "A"}},
	// 	// {"$group": bson.M{"_id": "$cust_id", "total": bson.M{"$sum": "$amount"}}},
	// 	// {"$sort": bson.M{"total": -1}},
	// }
	//
	dal.Query(ctx, nil)

	return "GetVoteByID function has been called", nil
}

func GetVotesByClubId(ctx context.Context, id string) (result string, err error) {

	result, err = dal.GetVotesByClubId(id)

	return result, err
}

func GetVotesByUserId(ctx context.Context, id string) (result string, err error) {

	//TODO create helper functions for bson.M and never use bson.M directly
	pipeline := []bson.M{
		{"$match": bson.M{"cuid": bson.M{"$in": []string{id}}}},
		{"$project": bson.M{"_id": 0, "uuid": 1}},
	}

	response, err := dal.Query(ctx, pipeline)
	if err != nil {
		log.Println(err)
	}

	return response, err
}
