// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package demo

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/kitx/endpointx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type gRPCDemoServiceServer struct {
	createUser grpc.Handler

	updateUser grpc.Handler

	getUser grpc.Handler

	getUsers grpc.Handler

	deleteUser grpc.Handler
}

func (s *gRPCDemoServiceServer) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.createUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *gRPCDemoServiceServer) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.updateUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *gRPCDemoServiceServer) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	ctx, rep, err := s.getUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*GetUserResponse), nil
}

func (s *gRPCDemoServiceServer) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	ctx, rep, err := s.getUsers.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*GetUsersResponse), nil
}

func (s *gRPCDemoServiceServer) DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.deleteUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewDemoServiceGRPCServer(
	endpoints interface {
		CreateUser() endpoint.Endpoint
		UpdateUser() endpoint.Endpoint
		GetUser() endpoint.Endpoint
		GetUsers() endpoint.Endpoint
		DeleteUser() endpoint.Endpoint
	},
	mdw []endpoint.Middleware,
	opts ...grpc.ServerOption,
) interface {
	CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
	UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
	GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
	GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
	DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
} {
	return &gRPCDemoServiceServer{
		createUser: grpc.NewServer(
			endpointx.Chain(endpoints.CreateUser(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		updateUser: grpc.NewServer(
			endpointx.Chain(endpoints.UpdateUser(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		getUser: grpc.NewServer(
			endpointx.Chain(endpoints.GetUser(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		getUsers: grpc.NewServer(
			endpointx.Chain(endpoints.GetUsers(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		deleteUser: grpc.NewServer(
			endpointx.Chain(endpoints.DeleteUser(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
	}
}

type gRPCDemoServiceClient struct {
	createUser endpoint.Endpoint
	updateUser endpoint.Endpoint
	getUser    endpoint.Endpoint
	getUsers   endpoint.Endpoint
	deleteUser endpoint.Endpoint
}

func (c *gRPCDemoServiceClient) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	rep, err := c.createUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *gRPCDemoServiceClient) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error) {
	rep, err := c.updateUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *gRPCDemoServiceClient) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	rep, err := c.getUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*GetUserResponse), nil
}

func (c *gRPCDemoServiceClient) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	rep, err := c.getUsers(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*GetUsersResponse), nil
}

func (c *gRPCDemoServiceClient) DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewDemoServiceGRPCClient(
	conn *grpc1.ClientConn,
	mdw []endpoint.Middleware,
	opts ...grpc.ClientOption,
) interface {
	CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
	UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error)
	GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error)
	GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error)
	DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error)
} {
	return &gRPCDemoServiceClient{
		createUser: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.demo.v1.DemoService",
				"CreateUser",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		updateUser: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.demo.v1.DemoService",
				"UpdateUser",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		getUser: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.demo.v1.DemoService",
				"GetUser",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				GetUserResponse{},
				opts...,
			).Endpoint(),
			mdw...),
		getUsers: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.demo.v1.DemoService",
				"GetUsers",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				GetUsersResponse{},
				opts...,
			).Endpoint(),
			mdw...),
		deleteUser: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.demo.v1.DemoService",
				"DeleteUser",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
	}
}
