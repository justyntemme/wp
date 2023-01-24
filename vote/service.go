package service

import (
	"context"
)

// @microgen middleware, logging, grpc, http, recovering
// @protobuf github.com/justyntemme/wp/vote/proto
type VoteService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetVoteById(ctx context.Context, VoteId string) (result string, err error)
	GetVotesByUserId(ctx context.Context, UserId string) (result string, err error)
	GetVotesByClubId(ctx context.Context, ClubId string) (result string, err error)
}

func GetVoteById(ctx context.Context, ClubId string) (result string, err error) {

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

// func GetVotesByClubId(ctx context.Context, VoteId string) (result string, err error) {

// 	log.Errorf("ID value is : " + VoteId)
// 	fmt.Println("get votes by club id " + VoteId)
// 	result, err = dal.GetVotesByClubId(ctx, VoteId)

// 	if err != nil {
// 		log.Errorf(err.Error())
// 	}

// 	return result, err
// }

func GetVotesByUserId(ctx context.Context, id string) (result string, err error) {

	return "response", err
}
