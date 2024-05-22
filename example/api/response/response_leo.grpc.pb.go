// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package response

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http "google.golang.org/genproto/googleapis/rpc/http"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type responseGRPCServer struct {
	omittedResponse grpc.Handler

	starResponse grpc.Handler

	namedResponse grpc.Handler

	httpBodyResponse grpc.Handler

	httpBodyNamedResponse grpc.Handler

	httpRequestStarBody grpc.Handler
}

func (s *responseGRPCServer) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.omittedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGRPCServer) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.starResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGRPCServer) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.namedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGRPCServer) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	ctx, rep, err := s.httpBodyResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*httpbody.HttpBody), nil
}

func (s *responseGRPCServer) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	ctx, rep, err := s.httpBodyNamedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*HttpBody), nil
}

func (s *responseGRPCServer) HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*http.HttpResponse, error) {
	ctx, rep, err := s.httpRequestStarBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*http.HttpResponse), nil
}

func NewResponseGRPCServer(
	endpoints interface {
		OmittedResponse() endpoint.Endpoint
		StarResponse() endpoint.Endpoint
		NamedResponse() endpoint.Endpoint
		HttpBodyResponse() endpoint.Endpoint
		HttpBodyNamedResponse() endpoint.Endpoint
		HttpRequestStarBody() endpoint.Endpoint
	},
	opts []grpc.ServerOption,
	mdw ...endpoint.Middleware,
) interface {
	OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error)
	HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error)
	HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*http.HttpResponse, error)
} {
	return &responseGRPCServer{
		omittedResponse: grpc.NewServer(
			endpointx.Chain(endpoints.OmittedResponse(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		starResponse: grpc.NewServer(
			endpointx.Chain(endpoints.StarResponse(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		namedResponse: grpc.NewServer(
			endpointx.Chain(endpoints.NamedResponse(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		httpBodyResponse: grpc.NewServer(
			endpointx.Chain(endpoints.HttpBodyResponse(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		httpBodyNamedResponse: grpc.NewServer(
			endpointx.Chain(endpoints.HttpBodyNamedResponse(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		httpRequestStarBody: grpc.NewServer(
			endpointx.Chain(endpoints.HttpRequestStarBody(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
	}
}

type responseGRPCClient struct {
	omittedResponse       endpoint.Endpoint
	starResponse          endpoint.Endpoint
	namedResponse         endpoint.Endpoint
	httpBodyResponse      endpoint.Endpoint
	httpBodyNamedResponse endpoint.Endpoint
	httpRequestStarBody   endpoint.Endpoint
}

func (c *responseGRPCClient) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.omittedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseGRPCClient) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.starResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseGRPCClient) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.namedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseGRPCClient) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	rep, err := c.httpBodyResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*httpbody.HttpBody), nil
}

func (c *responseGRPCClient) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	rep, err := c.httpBodyNamedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*HttpBody), nil
}

func (c *responseGRPCClient) HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*http.HttpResponse, error) {
	rep, err := c.httpRequestStarBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*http.HttpResponse), nil
}

func NewResponseGRPCClient(
	conn *grpc1.ClientConn,
	opts []grpc.ClientOption,
	mdw ...endpoint.Middleware,
) interface {
	OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error)
	HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error)
	HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*http.HttpResponse, error)
} {
	return &responseGRPCClient{
		omittedResponse: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.response.v1.Response",
				"OmittedResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				UserResponse{},
				opts...,
			).Endpoint(),
			mdw...),
		starResponse: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.response.v1.Response",
				"StarResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				UserResponse{},
				opts...,
			).Endpoint(),
			mdw...),
		namedResponse: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.response.v1.Response",
				"NamedResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				UserResponse{},
				opts...,
			).Endpoint(),
			mdw...),
		httpBodyResponse: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.response.v1.Response",
				"HttpBodyResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				httpbody.HttpBody{},
				opts...,
			).Endpoint(),
			mdw...),
		httpBodyNamedResponse: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.response.v1.Response",
				"HttpBodyNamedResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				HttpBody{},
				opts...,
			).Endpoint(),
			mdw...),
		httpRequestStarBody: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.response.v1.Response",
				"HttpRequestStarBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				http.HttpResponse{},
				opts...,
			).Endpoint(),
			mdw...),
	}
}
