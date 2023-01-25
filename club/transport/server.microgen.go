// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	club "github.com/justyntemme/wp/club"
)

func Endpoints(svc club.ClubService) EndpointsSet {
	return EndpointsSet{
		GetAllClubsNearMeEndpoint: GetAllClubsNearMeEndpoint(svc),
		GetClubByIdEndpoint:       GetClubByIdEndpoint(svc),
		GetTopClubsEndpoint:       GetTopClubsEndpoint(svc),
		GetTopClubsNearMeEndpoint: GetTopClubsNearMeEndpoint(svc),
	}
}

func GetClubByIdEndpoint(svc club.ClubService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetClubByIdRequest)
		res0, res1 := svc.GetClubById(arg0, req.Id)
		return &GetClubByIdResponse{Result: res0}, res1
	}
}

func GetTopClubsEndpoint(svc club.ClubService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetTopClubsRequest)
		res0, res1 := svc.GetTopClubs(arg0, req.Limit)
		return &GetTopClubsResponse{Result: res0}, res1
	}
}

func GetTopClubsNearMeEndpoint(svc club.ClubService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetTopClubsNearMeRequest)
		res0, res1 := svc.GetTopClubsNearMe(arg0, req.Limit)
		return &GetTopClubsNearMeResponse{Result: res0}, res1
	}
}

func GetAllClubsNearMeEndpoint(svc club.ClubService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetAllClubsNearMeRequest)
		res0, res1 := svc.GetAllClubsNearMe(arg0, req.Limit)
		return &GetAllClubsNearMeResponse{Result: res0}, res1
	}
}
