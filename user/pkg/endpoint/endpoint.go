package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "user/pkg/service"
)

// GetUserByIdRequest collects the request parameters for the GetUserById method.
type GetUserByIdRequest struct {
	S string `json:"s"`
}

// GetUserByIdResponse collects the response parameters for the GetUserById method.
type GetUserByIdResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetUserByIdEndpoint returns an endpoint that invokes GetUserById on the service.
func MakeGetUserByIdEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIdRequest)
		rs, err := s.GetUserById(ctx, req.S)
		return GetUserByIdResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserByIdResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetUserById implements Service. Primarily useful in a client.
func (e Endpoints) GetUserById(ctx context.Context, s string) (rs string, err error) {
	request := GetUserByIdRequest{S: s}
	response, err := e.GetUserByIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserByIdResponse).Rs, response.(GetUserByIdResponse).Err
}
