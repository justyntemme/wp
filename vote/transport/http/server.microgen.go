// Code generated by microgen 1.0.4. DO NOT EDIT.

package transporthttp

import (
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	transport "github.com/justyntemme/wp/vote/transport"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/get-vote-byid").Handler(
		http.NewServer(
			endpoints.GetVoteByIdEndpoint,
			_Decode_GetVoteById_Request,
			_Encode_GetVoteById_Response,
			opts...))
	mux.Methods("POST").Path("/get-votes-byuser-id").Handler(
		http.NewServer(
			endpoints.GetVotesByUserIdEndpoint,
			_Decode_GetVotesByUserId_Request,
			_Encode_GetVotesByUserId_Response,
			opts...))
	mux.Methods("POST").Path("/get-votes-byclub-id").Handler(
		http.NewServer(
			endpoints.GetVotesByClubIdEndpoint,
			_Decode_GetVotesByClubId_Request,
			_Encode_GetVotesByClubId_Response,
			opts...))
	return mux
}
