package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "user/pkg/endpoint"
	pb "user/pkg/grpc/pb"
)

// makeGetUserByIdHandler creates the handler logic
func makeGetUserByIdHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserByIdEndpoint, decodeGetUserByIdRequest, encodeGetUserByIdResponse, options...)
}

// decodeGetUserByIdResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserById request.
// TODO implement the decoder
func decodeGetUserByIdRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Decoder is not impelemented")
}

// encodeGetUserByIdResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserByIdResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'User' Encoder is not impelemented")
}
func (g *grpcServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdReply, error) {
	_, rep, err := g.getUserById.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserByIdReply), nil
}
