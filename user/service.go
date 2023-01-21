package service

import (
	"context"

	dal "github.com/justyntemme/wp/dal"
	"gopkg.in/mgo.v2/bson"
)

// @microgen middleware, logging, grpc, http, recovering
// @protobuf github.com/justyntemme/wp/user/proto
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetUserById(ctx context.Context, id string) (result string, err error)
}

func GetUserById(ctx context.Context, id string) (result string, err error) {

	// Define the aggregation pipeline
	pipeline := []bson.M{
		{"$match": bson.M{"Id": "A"}},
		{"$group": bson.M{"_id": "$cust_id", "total": bson.M{"$sum": "$amount"}}},
		{"$sort": bson.M{"total": -1}},
	}
	dal.Query(ctx, pipeline)

	return "GetUserByID function has been called", nil
}
