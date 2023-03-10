// Code generated by microgen 1.0.4. DO NOT EDIT.

package transport

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	vote "github.com/justyntemme/wp/vote"
)

func Endpoints(svc vote.VoteService) EndpointsSet {
	return EndpointsSet{
		GetVoteByIdEndpoint:      GetVoteByIdEndpoint(svc),
		GetVotesByClubIdEndpoint: GetVotesByClubIdEndpoint(svc),
		GetVotesByUserIdEndpoint: GetVotesByUserIdEndpoint(svc),
	}
}

func GetVoteByIdEndpoint(svc vote.VoteService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetVoteByIdRequest)
		res0, res1 := svc.GetVoteById(arg0, req.VoteId)
		return &GetVoteByIdResponse{Result: res0}, res1
	}
}

func GetVotesByUserIdEndpoint(svc vote.VoteService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetVotesByUserIdRequest)
		res0, res1 := svc.GetVotesByUserId(arg0, req.UserId)
		return &GetVotesByUserIdResponse{Result: res0}, res1
	}
}

func GetVotesByClubIdEndpoint(svc vote.VoteService) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetVotesByClubIdRequest)
		res0, res1 := svc.GetVotesByClubId(arg0, req.ClubId)
		return &GetVotesByClubIdResponse{Result: res0}, res1
	}
}
