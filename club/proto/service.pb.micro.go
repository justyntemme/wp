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

// Api Endpoints for ClubService service

func NewClubServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ClubService service

type ClubService interface {
	GetClubById(ctx context.Context, in *GetClubByIdRequest, opts ...client.CallOption) (*GetClubByIdResponse, error)
	GetTopClubs(ctx context.Context, in *GetTopClubsRequest, opts ...client.CallOption) (*GetTopClubsResponse, error)
	GetTopClubsNearMe(ctx context.Context, in *GetTopClubsNearMeRequest, opts ...client.CallOption) (*GetTopClubsNearMeResponse, error)
	GetAllClubsNearMe(ctx context.Context, in *GetAllClubsNearMeRequest, opts ...client.CallOption) (*GetAllClubsNearMeResponse, error)
}

type clubService struct {
	c    client.Client
	name string
}

func NewClubService(name string, c client.Client) ClubService {
	return &clubService{
		c:    c,
		name: name,
	}
}

func (c *clubService) GetClubById(ctx context.Context, in *GetClubByIdRequest, opts ...client.CallOption) (*GetClubByIdResponse, error) {
	req := c.c.NewRequest(c.name, "ClubService.GetClubById", in)
	out := new(GetClubByIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubService) GetTopClubs(ctx context.Context, in *GetTopClubsRequest, opts ...client.CallOption) (*GetTopClubsResponse, error) {
	req := c.c.NewRequest(c.name, "ClubService.GetTopClubs", in)
	out := new(GetTopClubsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubService) GetTopClubsNearMe(ctx context.Context, in *GetTopClubsNearMeRequest, opts ...client.CallOption) (*GetTopClubsNearMeResponse, error) {
	req := c.c.NewRequest(c.name, "ClubService.GetTopClubsNearMe", in)
	out := new(GetTopClubsNearMeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubService) GetAllClubsNearMe(ctx context.Context, in *GetAllClubsNearMeRequest, opts ...client.CallOption) (*GetAllClubsNearMeResponse, error) {
	req := c.c.NewRequest(c.name, "ClubService.GetAllClubsNearMe", in)
	out := new(GetAllClubsNearMeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClubService service

type ClubServiceHandler interface {
	GetClubById(context.Context, *GetClubByIdRequest, *GetClubByIdResponse) error
	GetTopClubs(context.Context, *GetTopClubsRequest, *GetTopClubsResponse) error
	GetTopClubsNearMe(context.Context, *GetTopClubsNearMeRequest, *GetTopClubsNearMeResponse) error
	GetAllClubsNearMe(context.Context, *GetAllClubsNearMeRequest, *GetAllClubsNearMeResponse) error
}

func RegisterClubServiceHandler(s server.Server, hdlr ClubServiceHandler, opts ...server.HandlerOption) error {
	type clubService interface {
		GetClubById(ctx context.Context, in *GetClubByIdRequest, out *GetClubByIdResponse) error
		GetTopClubs(ctx context.Context, in *GetTopClubsRequest, out *GetTopClubsResponse) error
		GetTopClubsNearMe(ctx context.Context, in *GetTopClubsNearMeRequest, out *GetTopClubsNearMeResponse) error
		GetAllClubsNearMe(ctx context.Context, in *GetAllClubsNearMeRequest, out *GetAllClubsNearMeResponse) error
	}
	type ClubService struct {
		clubService
	}
	h := &clubServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ClubService{h}, opts...))
}

type clubServiceHandler struct {
	ClubServiceHandler
}

func (h *clubServiceHandler) GetClubById(ctx context.Context, in *GetClubByIdRequest, out *GetClubByIdResponse) error {
	return h.ClubServiceHandler.GetClubById(ctx, in, out)
}

func (h *clubServiceHandler) GetTopClubs(ctx context.Context, in *GetTopClubsRequest, out *GetTopClubsResponse) error {
	return h.ClubServiceHandler.GetTopClubs(ctx, in, out)
}

func (h *clubServiceHandler) GetTopClubsNearMe(ctx context.Context, in *GetTopClubsNearMeRequest, out *GetTopClubsNearMeResponse) error {
	return h.ClubServiceHandler.GetTopClubsNearMe(ctx, in, out)
}

func (h *clubServiceHandler) GetAllClubsNearMe(ctx context.Context, in *GetAllClubsNearMeRequest, out *GetAllClubsNearMeResponse) error {
	return h.ClubServiceHandler.GetAllClubsNearMe(ctx, in, out)
}
