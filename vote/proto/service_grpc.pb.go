// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VoteServiceClient is the client API for VoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoteServiceClient interface {
	GetVoteById(ctx context.Context, in *GetVoteByIdRequest, opts ...grpc.CallOption) (*GetVoteByIdResponse, error)
	GetVotesByClubId(ctx context.Context, in *GetVotesByClubIdRequest, opts ...grpc.CallOption) (*GetVotesByClubIdResponse, error)
	GetVotesByUserId(ctx context.Context, in *GetVotesByUserIdRequest, opts ...grpc.CallOption) (*GetVotesByUserIdResponse, error)
}

type voteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoteServiceClient(cc grpc.ClientConnInterface) VoteServiceClient {
	return &voteServiceClient{cc}
}

func (c *voteServiceClient) GetVoteById(ctx context.Context, in *GetVoteByIdRequest, opts ...grpc.CallOption) (*GetVoteByIdResponse, error) {
	out := new(GetVoteByIdResponse)
	err := c.cc.Invoke(ctx, "/vote.VoteService/GetVoteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) GetVotesByClubId(ctx context.Context, in *GetVotesByClubIdRequest, opts ...grpc.CallOption) (*GetVotesByClubIdResponse, error) {
	out := new(GetVotesByClubIdResponse)
	err := c.cc.Invoke(ctx, "/vote.VoteService/GetVotesByClubId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voteServiceClient) GetVotesByUserId(ctx context.Context, in *GetVotesByUserIdRequest, opts ...grpc.CallOption) (*GetVotesByUserIdResponse, error) {
	out := new(GetVotesByUserIdResponse)
	err := c.cc.Invoke(ctx, "/vote.VoteService/GetVotesByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoteServiceServer is the server API for VoteService service.
// All implementations must embed UnimplementedVoteServiceServer
// for forward compatibility
type VoteServiceServer interface {
	GetVoteById(context.Context, *GetVoteByIdRequest) (*GetVoteByIdResponse, error)
	GetVotesByClubId(context.Context, *GetVotesByClubIdRequest) (*GetVotesByClubIdResponse, error)
	GetVotesByUserId(context.Context, *GetVotesByUserIdRequest) (*GetVotesByUserIdResponse, error)
	mustEmbedUnimplementedVoteServiceServer()
}

// UnimplementedVoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVoteServiceServer struct {
}

func (UnimplementedVoteServiceServer) GetVoteById(context.Context, *GetVoteByIdRequest) (*GetVoteByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoteById not implemented")
}
func (UnimplementedVoteServiceServer) GetVotesByClubId(context.Context, *GetVotesByClubIdRequest) (*GetVotesByClubIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVotesByClubId not implemented")
}
func (UnimplementedVoteServiceServer) GetVotesByUserId(context.Context, *GetVotesByUserIdRequest) (*GetVotesByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVotesByUserId not implemented")
}
func (UnimplementedVoteServiceServer) mustEmbedUnimplementedVoteServiceServer() {}

// UnsafeVoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoteServiceServer will
// result in compilation errors.
type UnsafeVoteServiceServer interface {
	mustEmbedUnimplementedVoteServiceServer()
}

func RegisterVoteServiceServer(s grpc.ServiceRegistrar, srv VoteServiceServer) {
	s.RegisterService(&VoteService_ServiceDesc, srv)
}

func _VoteService_GetVoteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVoteByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).GetVoteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vote.VoteService/GetVoteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).GetVoteById(ctx, req.(*GetVoteByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_GetVotesByClubId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVotesByClubIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).GetVotesByClubId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vote.VoteService/GetVotesByClubId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).GetVotesByClubId(ctx, req.(*GetVotesByClubIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoteService_GetVotesByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVotesByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoteServiceServer).GetVotesByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vote.VoteService/GetVotesByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoteServiceServer).GetVotesByUserId(ctx, req.(*GetVotesByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VoteService_ServiceDesc is the grpc.ServiceDesc for VoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vote.VoteService",
	HandlerType: (*VoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVoteById",
			Handler:    _VoteService_GetVoteById_Handler,
		},
		{
			MethodName: "GetVotesByClubId",
			Handler:    _VoteService_GetVotesByClubId_Handler,
		},
		{
			MethodName: "GetVotesByUserId",
			Handler:    _VoteService_GetVotesByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}