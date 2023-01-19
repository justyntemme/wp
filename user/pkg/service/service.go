package service

import "context"

// UserService describes the service.
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)

	GetUserById(ctx context.Context, s string) (rs string, err error)
	UpdateUserById(ctx context.Context, s string) (rs string, err error)
}

type basicUserService struct{}

func (b *basicUserService) GetUserById(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of GetUserById
	return rs, err
}
func (b *basicUserService) UpdateUserById(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of UpdateUserById
	return rs, err
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	return &basicUserService{}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
