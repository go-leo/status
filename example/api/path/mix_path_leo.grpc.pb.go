// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package path

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type mixPathGRPCServer struct {
	mixPath grpc.Handler
}

func (s *mixPathGRPCServer) MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.mixPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewMixPathGRPCServer(
	endpoints interface {
		MixPath() endpoint.Endpoint
	},
	opts []grpc.ServerOption,
	middlewares ...endpoint.Middleware,
) interface {
	MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error)
} {
	return &mixPathGRPCServer{
		mixPath: grpc.NewServer(
			endpointx.Chain(endpoints.MixPath(), middlewares...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
	}
}

type mixPathGRPCClient struct {
	mixPath endpoint.Endpoint
}

func (c *mixPathGRPCClient) MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error) {
	rep, err := c.mixPath(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewMixPathGRPCClient(
	conn *grpc1.ClientConn,
	opts []grpc.ClientOption,
	middlewares ...endpoint.Middleware,
) interface {
	MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error)
} {
	return &mixPathGRPCClient{
		mixPath: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.MixPath",
				"MixPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			middlewares...),
	}
}
