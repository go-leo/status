// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package endpointsapis

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	errorx "github.com/go-leo/gox/errorx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	listWorkspaces  transportx.ClientTransport
	getWorkspace    transportx.ClientTransport
	createWorkspace transportx.ClientTransport
	updateWorkspace transportx.ClientTransport
	deleteWorkspace transportx.ClientTransport
}

func (t *workspacesGrpcClientTransports) ListWorkspaces() transportx.ClientTransport {
	return t.listWorkspaces
}

func (t *workspacesGrpcClientTransports) GetWorkspace() transportx.ClientTransport {
	return t.getWorkspace
}

func (t *workspacesGrpcClientTransports) CreateWorkspace() transportx.ClientTransport {
	return t.createWorkspace
}

func (t *workspacesGrpcClientTransports) UpdateWorkspace() transportx.ClientTransport {
	return t.updateWorkspace
}

func (t *workspacesGrpcClientTransports) DeleteWorkspace() transportx.ClientTransport {
	return t.deleteWorkspace
}

func NewWorkspacesGrpcClientTransports(target string, options ...transportx.ClientTransportOption) (WorkspacesClientTransports, error) {
	t := &workspacesGrpcClientTransports{}
	var err error
	t.listWorkspaces, err = errorx.Break[transportx.ClientTransport](err)(_Workspaces_ListWorkspaces_GrpcClient_Transport(target, options...))
	t.getWorkspace, err = errorx.Break[transportx.ClientTransport](err)(_Workspaces_GetWorkspace_GrpcClient_Transport(target, options...))
	t.createWorkspace, err = errorx.Break[transportx.ClientTransport](err)(_Workspaces_CreateWorkspace_GrpcClient_Transport(target, options...))
	t.updateWorkspace, err = errorx.Break[transportx.ClientTransport](err)(_Workspaces_UpdateWorkspace_GrpcClient_Transport(target, options...))
	t.deleteWorkspace, err = errorx.Break[transportx.ClientTransport](err)(_Workspaces_DeleteWorkspace_GrpcClient_Transport(target, options...))
	return t, err
}

type workspacesGrpcClient struct {
	endpoints WorkspacesEndpoints
}

func (c *workspacesGrpcClient) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.ListWorkspaces(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*ListWorkspacesResponse), nil
}

func (c *workspacesGrpcClient) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.GetWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.CreateWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.UpdateWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.DeleteWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesGrpcClient(transports WorkspacesClientTransports, middlewares ...endpoint.Middleware) WorkspacesService {
	endpoints := newWorkspacesClientEndpoints(transports, middlewares...)
	return &workspacesGrpcClient{endpoints: endpoints}
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

func _Workspaces_ListWorkspaces_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"google.example.endpointsapis.v1.Workspaces",
				"ListWorkspaces",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				ListWorkspacesResponse{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
				grpc.ClientBefore(grpcx.OutgoingStain),
			),
			options...,
		)
	}
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

func _Workspaces_GetWorkspace_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"google.example.endpointsapis.v1.Workspaces",
				"GetWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				Workspace{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
				grpc.ClientBefore(grpcx.OutgoingStain),
			),
			options...,
		)
	}
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

func _Workspaces_CreateWorkspace_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"google.example.endpointsapis.v1.Workspaces",
				"CreateWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				Workspace{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
				grpc.ClientBefore(grpcx.OutgoingStain),
			),
			options...,
		)
	}
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

func _Workspaces_UpdateWorkspace_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"google.example.endpointsapis.v1.Workspaces",
				"UpdateWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				Workspace{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
				grpc.ClientBefore(grpcx.OutgoingStain),
			),
			options...,
		)
	}
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

func _Workspaces_DeleteWorkspace_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"google.example.endpointsapis.v1.Workspaces",
				"DeleteWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
				grpc.ClientBefore(grpcx.OutgoingStain),
			),
			options...,
		)
	}
}
