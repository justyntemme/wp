package service

import (
	"context"

	"github.com/justyntemme/wp/dal"
	"go-micro.dev/v4/util/log"
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

	return "GetVoteByID function has been called", nil
}

func GetVotesByClubId(ctx context.Context, id string) (result string, err error) {
	result, err = dal.GetVotesByClubId(ctx, id)
	if err != nil {
		log.Errorf(err.Error())
	}

	return result, nil
}

func GetVotesByUserId(ctx context.Context, id string) (result string, err error) {

	return "response", err
}
