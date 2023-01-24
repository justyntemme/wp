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
