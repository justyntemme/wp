package service

import (
	"context"

	dal "github.com/justyntemme/wp/dal"
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

	dal.Query(ctx, nil)

	return "GetVoteByID function has been called", nil
}

func GetVotesByUserId(ctx context.Context, id string) (result string, err error) {

	dal.Query(ctx, nil)

	return "GetVoteByID function has been called", nil
}
