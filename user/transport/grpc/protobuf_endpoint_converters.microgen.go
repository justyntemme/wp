// Code generated by microgen 1.0.4. DO NOT EDIT.

// Please, do not change functions names!
package transportgrpc

import (
	"context"
	"errors"
	pb "github.com/justyntemme/wp/user/proto"
	transport "github.com/justyntemme/wp/user/transport"
)

func _Encode_GetUserById_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetUserByIdRequest")
	}
	req := request.(*transport.GetUserByIdRequest)
	return &pb.GetUserByIdRequest{Id: req.Id}, nil
}

func _Encode_GetUserById_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetUserByIdResponse")
	}
	resp := response.(*transport.GetUserByIdResponse)
	return &pb.GetUserByIdResponse{Result: resp.Result}, nil
}

func _Decode_GetUserById_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetUserByIdRequest")
	}
	req := request.(*pb.GetUserByIdRequest)
	return &transport.GetUserByIdRequest{Id: string(req.Id)}, nil
}

func _Decode_GetUserById_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetUserByIdResponse")
	}
	resp := response.(*pb.GetUserByIdResponse)
	return &transport.GetUserByIdResponse{Result: string(resp.Result)}, nil
}
