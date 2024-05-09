// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package query

import (
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	mux "github.com/gorilla/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
)

type httpQueryClient struct {
	query endpoint.Endpoint
}

func (c *httpQueryClient) Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error) {
	rep, err := c.query(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewQueryHTTPClient(
	instance string,
	mdw []endpoint.Middleware,
	opts ...http.ClientOption,
) interface {
	Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
} {
	router := mux.NewRouter()
	router.NewRoute().
		Name("/leo.example.query.v1.Query/Query").
		Methods("GET").
		Path("/v1/query")
	return &httpQueryClient{
		query: endpointx.Chain(
			http.NewExplicitClient(
				func(ctx context.Context, obj interface{}) (*http1.Request, error) {
					req, ok := obj.(*QueryRequest)
					if !ok {
						return nil, fmt.Errorf("invalid request object type, %T", obj)
					}
					if req == nil {
						return nil, errors.New("request object is nil")
					}
					var method = "GET"
					var url string
					var body io.Reader
					var pairs []string
					path, err := router.Get("/leo.example.query.v1.Query/Query").URLPath(pairs...)
					if err != nil {
						return nil, err
					}
					url = fmt.Sprintf("%s://%s%s", "http", instance, path)
					r, err := http1.NewRequestWithContext(ctx, method, url, body)
					if err != nil {
						return nil, err
					}
					return r, nil
				},
				func(ctx context.Context, r *http1.Response) (interface{}, error) {
					return nil, nil
				},
				opts...,
			).Endpoint(),
			mdw...),
	}
}
