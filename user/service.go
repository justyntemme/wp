package service

import "context"

// @microgen middleware, logging, grpc, http, recovering
// @protobuf github.com/justyntemme/wp/user/proto
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetUserById(ctx context.Context, id string) (result string, err error)
}
