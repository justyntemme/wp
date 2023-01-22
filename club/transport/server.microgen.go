// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	club "github.com/justyntemme/wp/club"
)

func Endpoints(svc club.ClubService) EndpointsSet {
	return EndpointsSet{GetClubByIdEndpoint: GetClubByIdEndpoint(svc)}
}

func GetClubByIdEndpoint(svc club.ClubService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetClubByIdRequest)
		res0, res1 := svc.GetClubById(arg0, req.Id)
		return &GetClubByIdResponse{Result: res0}, res1
	}
}
