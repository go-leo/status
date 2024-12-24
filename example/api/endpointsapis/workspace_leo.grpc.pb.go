// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package endpointsapis

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// =========================== grpc server ===========================

type WorkspacesGrpcServerTransports interface {
	ListWorkspaces() *grpc.Server
	GetWorkspace() *grpc.Server
	CreateWorkspace() *grpc.Server
	UpdateWorkspace() *grpc.Server
	DeleteWorkspace() *grpc.Server
}

type workspacesGrpcServerTransports struct {
	listWorkspaces  *grpc.Server
	getWorkspace    *grpc.Server
	createWorkspace *grpc.Server
	updateWorkspace *grpc.Server
	deleteWorkspace *grpc.Server
}

func (t *workspacesGrpcServerTransports) ListWorkspaces() *grpc.Server {
	return t.listWorkspaces
}

func (t *workspacesGrpcServerTransports) GetWorkspace() *grpc.Server {
	return t.getWorkspace
}

func (t *workspacesGrpcServerTransports) CreateWorkspace() *grpc.Server {
	return t.createWorkspace
}

func (t *workspacesGrpcServerTransports) UpdateWorkspace() *grpc.Server {
	return t.updateWorkspace
}

func (t *workspacesGrpcServerTransports) DeleteWorkspace() *grpc.Server {
	return t.deleteWorkspace
}

func newWorkspacesGrpcServerTransports(endpoints WorkspacesEndpoints) WorkspacesGrpcServerTransports {
	return &workspacesGrpcServerTransports{
		listWorkspaces:  _Workspaces_ListWorkspaces_GrpcServer_Transport(endpoints),
		getWorkspace:    _Workspaces_GetWorkspace_GrpcServer_Transport(endpoints),
		createWorkspace: _Workspaces_CreateWorkspace_GrpcServer_Transport(endpoints),
		updateWorkspace: _Workspaces_UpdateWorkspace_GrpcServer_Transport(endpoints),
		deleteWorkspace: _Workspaces_DeleteWorkspace_GrpcServer_Transport(endpoints),
	}
}

type workspacesGrpcServer struct {
	listWorkspaces  *grpc.Server
	getWorkspace    *grpc.Server
	createWorkspace *grpc.Server
	updateWorkspace *grpc.Server
	deleteWorkspace *grpc.Server
}

func (s *workspacesGrpcServer) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	ctx, rep, err := s.listWorkspaces.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*ListWorkspacesResponse), nil
}

func (s *workspacesGrpcServer) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.getWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *workspacesGrpcServer) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.createWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *workspacesGrpcServer) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.updateWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *workspacesGrpcServer) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.deleteWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesGrpcServer(svc WorkspacesService, middlewares ...endpoint.Middleware) WorkspacesService {
	endpoints := newWorkspacesServerEndpoints(svc, middlewares...)
	transports := newWorkspacesGrpcServerTransports(endpoints)
	return &workspacesGrpcServer{
		listWorkspaces:  transports.ListWorkspaces(),
		getWorkspace:    transports.GetWorkspace(),
		createWorkspace: transports.CreateWorkspace(),
		updateWorkspace: transports.UpdateWorkspace(),
		deleteWorkspace: transports.DeleteWorkspace(),
	}
}

// =========================== grpc client ===========================

type workspacesGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *workspacesGrpcClientTransports) ListWorkspaces(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.endpointsapis.v1.Workspaces",
		"ListWorkspaces",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		ListWorkspacesResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *workspacesGrpcClientTransports) GetWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.endpointsapis.v1.Workspaces",
		"GetWorkspace",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Workspace{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *workspacesGrpcClientTransports) CreateWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.endpointsapis.v1.Workspaces",
		"CreateWorkspace",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Workspace{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *workspacesGrpcClientTransports) UpdateWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.endpointsapis.v1.Workspaces",
		"UpdateWorkspace",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		Workspace{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *workspacesGrpcClientTransports) DeleteWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"google.example.endpointsapis.v1.Workspaces",
		"DeleteWorkspace",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func newWorkspacesGrpcClientTransports(
	dialOptions []grpc1.DialOption,
	clientOptions []grpc.ClientOption,
	middlewares []endpoint.Middleware,
) WorkspacesClientTransports {
	return &workspacesGrpcClientTransports{
		dialOptions:   dialOptions,
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

func NewWorkspacesGrpcClient(target string, opts ...grpcx.ClientOption) WorkspacesService {
	options := grpcx.NewClientOptions(opts...)
	transports := newWorkspacesGrpcClientTransports(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newWorkspacesClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newWorkspacesClientService(endpoints, grpcx.GrpcClient)
}

// =========================== grpc transport ===========================

func _Workspaces_ListWorkspaces_GrpcServer_Transport(endpoints WorkspacesEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.ListWorkspaces(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Workspaces_GetWorkspace_GrpcServer_Transport(endpoints WorkspacesEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.GetWorkspace(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/GetWorkspace")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Workspaces_CreateWorkspace_GrpcServer_Transport(endpoints WorkspacesEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.CreateWorkspace(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Workspaces_UpdateWorkspace_GrpcServer_Transport(endpoints WorkspacesEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.UpdateWorkspace(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Workspaces_DeleteWorkspace_GrpcServer_Transport(endpoints WorkspacesEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.DeleteWorkspace(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}
