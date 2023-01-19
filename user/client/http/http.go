package http

import (
	"bytes"
	"context"
	"encoding/json"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpoint1 "github.com/justyntemme/wp/user/pkg/endpoint"
	http2 "github.com/justyntemme/wp/user/pkg/http"
	service "github.com/justyntemme/wp/user/pkg/service"
	"io/ioutil"
	http1 "net/http"
	"net/url"
	"strings"
)

// New returns an AddService backed by an HTTP server living at the remote
// instance. We expect instance to come from a service discovery system, so
// likely of the form "host:port".
func New(instance string, options map[string][]http.ClientOption) (service.UserService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var getUserByIdEndpoint endpoint.Endpoint
	{
		getUserByIdEndpoint = http.NewClient("POST", copyURL(u, "/get-user-by-id"), encodeHTTPGenericRequest, decodeGetUserByIdResponse, options["GetUserById"]...).Endpoint()
	}

	var updateUserByIdEndpoint endpoint.Endpoint
	{
		updateUserByIdEndpoint = http.NewClient("POST", copyURL(u, "/update-user-by-id"), encodeHTTPGenericRequest, decodeUpdateUserByIdResponse, options["UpdateUserById"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		GetUserByIdEndpoint:    getUserByIdEndpoint,
		UpdateUserByIdEndpoint: updateUserByIdEndpoint,
	}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// SON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http1.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeGetUserByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeGetUserByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.GetUserByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeUpdateUserByIdResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeUpdateUserByIdResponse(_ context.Context, r *http1.Response) (interface{}, error) {
	if r.StatusCode != http1.StatusOK {
		return nil, http2.ErrorDecoder(r)
	}
	var resp endpoint1.UpdateUserByIdResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func copyURL(base *url.URL, path string) (next *url.URL) {
	n := *base
	n.Path = path
	next = &n
	return
}
