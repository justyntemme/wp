// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import endpoint "github.com/go-kit/kit/endpoint"

// EndpointsSet implements UserService API and used for transport purposes.
type OneToManyStreamEndpoint func(req interface{}, stream interface{}) error

type ManyToManyStreamEndpoint func(stream interface{}) error

type ManyToOneStreamEndpoint func(stream interface{}) error

type EndpointsSet struct {
	GetUserByIdEndpoint endpoint.Endpoint
}
