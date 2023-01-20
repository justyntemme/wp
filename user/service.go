package service

import "context"

// @microgen middleware, logging, grpc, http, recovering, main
// @protobuf github.com/justyntemme/whats-poppin/user/proto
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetUserById(ctx context.Context, s string) (rs string, err error)
}
