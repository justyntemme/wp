// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service.proto

package protobuf

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for VoteService service

func NewVoteServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VoteService service

type VoteService interface {
	GetVoteById(ctx context.Context, in *GetVoteByIdRequest, opts ...client.CallOption) (*GetVoteByIdResponse, error)
	GetVoteByClubId(ctx context.Context, in *GetVoteByClubIdRequest, opts ...client.CallOption) (*GetVoteByClubIdResponse, error)
	GetVoteByUserId(ctx context.Context, in *GetVoteByUserIdRequest, opts ...client.CallOption) (*GetVoteByUserIdResponse, error)
}

type voteService struct {
	c    client.Client
	name string
}

func NewVoteService(name string, c client.Client) VoteService {
	return &voteService{
		c:    c,
		name: name,
	}
}

func (c *voteService) GetVoteById(ctx context.Context, in *GetVoteByIdRequest, opts ...client.CallOption) (*GetVoteByIdResponse, error) {
	req := c.c.NewRequest(c.name, "VoteService.GetVoteById", in)
	out := new(GetVoteByIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteService) GetVoteByClubId(ctx context.Context, in *GetVoteByClubIdRequest, opts ...client.CallOption) (*GetVoteByClubIdResponse, error) {
	req := c.c.NewRequest(c.name, "VoteService.GetVoteByClubId", in)
	out := new(GetVoteByClubIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteService) GetVoteByUserId(ctx context.Context, in *GetVoteByUserIdRequest, opts ...client.CallOption) (*GetVoteByUserIdResponse, error) {
	req := c.c.NewRequest(c.name, "VoteService.GetVoteByUserId", in)
	out := new(GetVoteByUserIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VoteService service

type VoteServiceHandler interface {
	GetVoteById(context.Context, *GetVoteByIdRequest, *GetVoteByIdResponse) error
	GetVoteByClubId(context.Context, *GetVoteByClubIdRequest, *GetVoteByClubIdResponse) error
	GetVoteByUserId(context.Context, *GetVoteByUserIdRequest, *GetVoteByUserIdResponse) error
}

func RegisterVoteServiceHandler(s server.Server, hdlr VoteServiceHandler, opts ...server.HandlerOption) error {
	type voteService interface {
		GetVoteById(ctx context.Context, in *GetVoteByIdRequest, out *GetVoteByIdResponse) error
		GetVoteByClubId(ctx context.Context, in *GetVoteByClubIdRequest, out *GetVoteByClubIdResponse) error
		GetVoteByUserId(ctx context.Context, in *GetVoteByUserIdRequest, out *GetVoteByUserIdResponse) error
	}
	type VoteService struct {
		voteService
	}
	h := &voteServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VoteService{h}, opts...))
}

type voteServiceHandler struct {
	VoteServiceHandler
}

func (h *voteServiceHandler) GetVoteById(ctx context.Context, in *GetVoteByIdRequest, out *GetVoteByIdResponse) error {
	return h.VoteServiceHandler.GetVoteById(ctx, in, out)
}

func (h *voteServiceHandler) GetVoteByClubId(ctx context.Context, in *GetVoteByClubIdRequest, out *GetVoteByClubIdResponse) error {
	return h.VoteServiceHandler.GetVoteByClubId(ctx, in, out)
}

func (h *voteServiceHandler) GetVoteByUserId(ctx context.Context, in *GetVoteByUserIdRequest, out *GetVoteByUserIdResponse) error {
	return h.VoteServiceHandler.GetVoteByUserId(ctx, in, out)
}
