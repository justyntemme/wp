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

// ClubServiceClient is the client API for ClubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClubServiceClient interface {
	GetClubById(ctx context.Context, in *GetClubByIdRequest, opts ...grpc.CallOption) (*GetClubByIdResponse, error)
}

type clubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClubServiceClient(cc grpc.ClientConnInterface) ClubServiceClient {
	return &clubServiceClient{cc}
}

func (c *clubServiceClient) GetClubById(ctx context.Context, in *GetClubByIdRequest, opts ...grpc.CallOption) (*GetClubByIdResponse, error) {
	out := new(GetClubByIdResponse)
	err := c.cc.Invoke(ctx, "/club.ClubService/GetClubById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClubServiceServer is the server API for ClubService service.
// All implementations must embed UnimplementedClubServiceServer
// for forward compatibility
type ClubServiceServer interface {
	GetClubById(context.Context, *GetClubByIdRequest) (*GetClubByIdResponse, error)
	mustEmbedUnimplementedClubServiceServer()
}

// UnimplementedClubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClubServiceServer struct {
}

func (UnimplementedClubServiceServer) GetClubById(context.Context, *GetClubByIdRequest) (*GetClubByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClubById not implemented")
}
func (UnimplementedClubServiceServer) mustEmbedUnimplementedClubServiceServer() {}

// UnsafeClubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClubServiceServer will
// result in compilation errors.
type UnsafeClubServiceServer interface {
	mustEmbedUnimplementedClubServiceServer()
}

func RegisterClubServiceServer(s grpc.ServiceRegistrar, srv ClubServiceServer) {
	s.RegisterService(&ClubService_ServiceDesc, srv)
}

func _ClubService_GetClubById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClubByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServiceServer).GetClubById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/club.ClubService/GetClubById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServiceServer).GetClubById(ctx, req.(*GetClubByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClubService_ServiceDesc is the grpc.ServiceDesc for ClubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "club.ClubService",
	HandlerType: (*ClubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClubById",
			Handler:    _ClubService_GetClubById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}