// Code generated by microgen 1.0.4. DO NOT EDIT.

// Please, do not change functions names!
package transportgrpc

import (
	"context"
	"errors"
	pb "github.com/justyntemme/wp/vote/proto"
	transport "github.com/justyntemme/wp/vote/transport"
)

func _Encode_GetVoteById_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetVoteByIdRequest")
	}
	req := request.(*transport.GetVoteByIdRequest)
	return &pb.GetVoteByIdRequest{Id: req.Id}, nil
}

func _Encode_GetVoteById_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetVoteByIdResponse")
	}
	resp := response.(*transport.GetVoteByIdResponse)
	return &pb.GetVoteByIdResponse{Result: resp.Result}, nil
}

func _Decode_GetVoteById_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetVoteByIdRequest")
	}
	req := request.(*pb.GetVoteByIdRequest)
	return &transport.GetVoteByIdRequest{Id: string(req.Id)}, nil
}

func _Decode_GetVoteById_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetVoteByIdResponse")
	}
	resp := response.(*pb.GetVoteByIdResponse)
	return &transport.GetVoteByIdResponse{Result: string(resp.Result)}, nil
}
