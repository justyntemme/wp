// Code generated by microgen 1.0.4. DO NOT EDIT.

package transporthttp

import (
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	transport "github.com/justyntemme/wp/club/transport"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/get-club-byid").Handler(
		http.NewServer(
			endpoints.GetClubByIdEndpoint,
			_Decode_GetClubById_Request,
			_Encode_GetClubById_Response,
			opts...))
	mux.Methods("POST").Path("/get-top-clubs").Handler(
		http.NewServer(
			endpoints.GetTopClubsEndpoint,
			_Decode_GetTopClubs_Request,
			_Encode_GetTopClubs_Response,
			opts...))
	mux.Methods("POST").Path("/get-top-clubs-near-me").Handler(
		http.NewServer(
			endpoints.GetTopClubsNearMeEndpoint,
			_Decode_GetTopClubsNearMe_Request,
			_Encode_GetTopClubsNearMe_Response,
			opts...))
	mux.Methods("POST").Path("/get-all-clubs-near-me").Handler(
		http.NewServer(
			endpoints.GetAllClubsNearMeEndpoint,
			_Decode_GetAllClubsNearMe_Request,
			_Encode_GetAllClubsNearMe_Response,
			opts...))
	return mux
}
