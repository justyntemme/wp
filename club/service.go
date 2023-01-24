package service

import (
	"context"
)

// @microgen middleware, logging, grpc, http, recovering
// @protobuf github.com/justyntemme/wp/club/proto
type ClubService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetClubById(ctx context.Context, id string) (result string, err error)
}
