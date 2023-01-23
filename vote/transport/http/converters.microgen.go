// Code generated by microgen 1.0.4. DO NOT EDIT.

// Please, do not change functions names!
package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	transport "github.com/justyntemme/wp/vote/transport"
	"io/ioutil"
	"net/http"
	"path"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func CommonHTTPResponseEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func _Decode_GetVoteById_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetVoteByIdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_GetVotesByUserId_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetVotesByUserIdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_GetVotesByClubId_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetVotesByClubIdRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_GetVoteById_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetVoteByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_GetVotesByUserId_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetVotesByUserIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_GetVotesByClubId_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetVotesByClubIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Encode_GetVoteById_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "get-vote-byid")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetVotesByUserId_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "get-votes-byuser-id")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetVotesByClubId_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "get-votes-byclub-id")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetVoteById_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetVotesByUserId_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetVotesByClubId_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
