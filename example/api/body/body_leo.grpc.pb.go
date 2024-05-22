// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package body

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

type bodyGRPCServer struct {
	starBody grpc.Handler

	namedBody grpc.Handler

	nonBody grpc.Handler

	httpBodyStarBody grpc.Handler

	httpBodyNamedBody grpc.Handler

	httpRequestStarBody grpc.Handler
}

func (s *bodyGRPCServer) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx, rep, err := s.starBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGRPCServer) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.namedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGRPCServer) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, rep, err := s.nonBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGRPCServer) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyStarBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGRPCServer) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyNamedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGRPCServer) HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpRequestStarBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewBodyGRPCServer(
	endpoints interface {
		StarBody() endpoint.Endpoint
		NamedBody() endpoint.Endpoint
		NonBody() endpoint.Endpoint
		HttpBodyStarBody() endpoint.Endpoint
		HttpBodyNamedBody() endpoint.Endpoint
		HttpRequestStarBody() endpoint.Endpoint
	},
	opts []grpc.ServerOption,
	mdw ...endpoint.Middleware,
) interface {
	StarBody(ctx context.Context, request *User) (*emptypb.Empty, error)
	NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error)
	NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error)
	HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error)
	HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error)
	HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*emptypb.Empty, error)
} {
	return &bodyGRPCServer{
		starBody: grpc.NewServer(
			endpointx.Chain(endpoints.StarBody(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		namedBody: grpc.NewServer(
			endpointx.Chain(endpoints.NamedBody(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		nonBody: grpc.NewServer(
			endpointx.Chain(endpoints.NonBody(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		httpBodyStarBody: grpc.NewServer(
			endpointx.Chain(endpoints.HttpBodyStarBody(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		httpBodyNamedBody: grpc.NewServer(
			endpointx.Chain(endpoints.HttpBodyNamedBody(), mdw...),
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

type bodyGRPCClient struct {
	starBody            endpoint.Endpoint
	namedBody           endpoint.Endpoint
	nonBody             endpoint.Endpoint
	httpBodyStarBody    endpoint.Endpoint
	httpBodyNamedBody   endpoint.Endpoint
	httpRequestStarBody endpoint.Endpoint
}

func (c *bodyGRPCClient) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	rep, err := c.starBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGRPCClient) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	rep, err := c.namedBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGRPCClient) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	rep, err := c.nonBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGRPCClient) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	rep, err := c.httpBodyStarBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGRPCClient) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	rep, err := c.httpBodyNamedBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGRPCClient) HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*emptypb.Empty, error) {
	rep, err := c.httpRequestStarBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewBodyGRPCClient(
	conn *grpc1.ClientConn,
	opts []grpc.ClientOption,
	mdw ...endpoint.Middleware,
) interface {
	StarBody(ctx context.Context, request *User) (*emptypb.Empty, error)
	NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error)
	NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error)
	HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error)
	HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error)
	HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*emptypb.Empty, error)
} {
	return &bodyGRPCClient{
		starBody: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.body.v1.Body",
				"StarBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		namedBody: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.body.v1.Body",
				"NamedBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		nonBody: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.body.v1.Body",
				"NonBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		httpBodyStarBody: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.body.v1.Body",
				"HttpBodyStarBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		httpBodyNamedBody: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.body.v1.Body",
				"HttpBodyNamedBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		httpRequestStarBody: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.body.v1.Body",
				"HttpRequestStarBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
	}
}
