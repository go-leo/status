// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package path

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

type PathGrpcServerTransports interface {
	BoolPath() *grpc.Server
	Int32Path() *grpc.Server
	Int64Path() *grpc.Server
	Uint32Path() *grpc.Server
	Uint64Path() *grpc.Server
	FloatPath() *grpc.Server
	DoublePath() *grpc.Server
	StringPath() *grpc.Server
	EnumPath() *grpc.Server
}

type pathGrpcServerTransports struct {
	boolPath   *grpc.Server
	int32Path  *grpc.Server
	int64Path  *grpc.Server
	uint32Path *grpc.Server
	uint64Path *grpc.Server
	floatPath  *grpc.Server
	doublePath *grpc.Server
	stringPath *grpc.Server
	enumPath   *grpc.Server
}

func (t *pathGrpcServerTransports) BoolPath() *grpc.Server {
	return t.boolPath
}

func (t *pathGrpcServerTransports) Int32Path() *grpc.Server {
	return t.int32Path
}

func (t *pathGrpcServerTransports) Int64Path() *grpc.Server {
	return t.int64Path
}

func (t *pathGrpcServerTransports) Uint32Path() *grpc.Server {
	return t.uint32Path
}

func (t *pathGrpcServerTransports) Uint64Path() *grpc.Server {
	return t.uint64Path
}

func (t *pathGrpcServerTransports) FloatPath() *grpc.Server {
	return t.floatPath
}

func (t *pathGrpcServerTransports) DoublePath() *grpc.Server {
	return t.doublePath
}

func (t *pathGrpcServerTransports) StringPath() *grpc.Server {
	return t.stringPath
}

func (t *pathGrpcServerTransports) EnumPath() *grpc.Server {
	return t.enumPath
}

func newPathGrpcServerTransports(endpoints PathEndpoints) PathGrpcServerTransports {
	return &pathGrpcServerTransports{
		boolPath:   _Path_BoolPath_GrpcServer_Transport(endpoints),
		int32Path:  _Path_Int32Path_GrpcServer_Transport(endpoints),
		int64Path:  _Path_Int64Path_GrpcServer_Transport(endpoints),
		uint32Path: _Path_Uint32Path_GrpcServer_Transport(endpoints),
		uint64Path: _Path_Uint64Path_GrpcServer_Transport(endpoints),
		floatPath:  _Path_FloatPath_GrpcServer_Transport(endpoints),
		doublePath: _Path_DoublePath_GrpcServer_Transport(endpoints),
		stringPath: _Path_StringPath_GrpcServer_Transport(endpoints),
		enumPath:   _Path_EnumPath_GrpcServer_Transport(endpoints),
	}
}

type pathGrpcServer struct {
	boolPath   *grpc.Server
	int32Path  *grpc.Server
	int64Path  *grpc.Server
	uint32Path *grpc.Server
	uint64Path *grpc.Server
	floatPath  *grpc.Server
	doublePath *grpc.Server
	stringPath *grpc.Server
	enumPath   *grpc.Server
}

func (s *pathGrpcServer) BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.boolPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.int32Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.int64Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.uint32Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.uint64Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.floatPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.doublePath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.stringPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGrpcServer) EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.enumPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewPathGrpcServer(svc PathService, middlewares ...endpoint.Middleware) PathService {
	endpoints := newPathServerEndpoints(svc, middlewares...)
	transports := newPathGrpcServerTransports(endpoints)
	return &pathGrpcServer{
		boolPath:   transports.BoolPath(),
		int32Path:  transports.Int32Path(),
		int64Path:  transports.Int64Path(),
		uint32Path: transports.Uint32Path(),
		uint64Path: transports.Uint64Path(),
		floatPath:  transports.FloatPath(),
		doublePath: transports.DoublePath(),
		stringPath: transports.StringPath(),
		enumPath:   transports.EnumPath(),
	}
}

// =========================== grpc client ===========================

type pathGrpcClientTransports struct {
	boolPath   transportx.ClientTransport
	int32Path  transportx.ClientTransport
	int64Path  transportx.ClientTransport
	uint32Path transportx.ClientTransport
	uint64Path transportx.ClientTransport
	floatPath  transportx.ClientTransport
	doublePath transportx.ClientTransport
	stringPath transportx.ClientTransport
	enumPath   transportx.ClientTransport
}

func (t *pathGrpcClientTransports) BoolPath() transportx.ClientTransport {
	return t.boolPath
}

func (t *pathGrpcClientTransports) Int32Path() transportx.ClientTransport {
	return t.int32Path
}

func (t *pathGrpcClientTransports) Int64Path() transportx.ClientTransport {
	return t.int64Path
}

func (t *pathGrpcClientTransports) Uint32Path() transportx.ClientTransport {
	return t.uint32Path
}

func (t *pathGrpcClientTransports) Uint64Path() transportx.ClientTransport {
	return t.uint64Path
}

func (t *pathGrpcClientTransports) FloatPath() transportx.ClientTransport {
	return t.floatPath
}

func (t *pathGrpcClientTransports) DoublePath() transportx.ClientTransport {
	return t.doublePath
}

func (t *pathGrpcClientTransports) StringPath() transportx.ClientTransport {
	return t.stringPath
}

func (t *pathGrpcClientTransports) EnumPath() transportx.ClientTransport {
	return t.enumPath
}

func NewPathGrpcClientTransports(target string, options ...transportx.ClientTransportOption) (PathClientTransports, error) {
	t := &pathGrpcClientTransports{}
	var err error
	t.boolPath, err = errorx.Break[transportx.ClientTransport](err)(_Path_BoolPath_GrpcClient_Transport(target, options...))
	t.int32Path, err = errorx.Break[transportx.ClientTransport](err)(_Path_Int32Path_GrpcClient_Transport(target, options...))
	t.int64Path, err = errorx.Break[transportx.ClientTransport](err)(_Path_Int64Path_GrpcClient_Transport(target, options...))
	t.uint32Path, err = errorx.Break[transportx.ClientTransport](err)(_Path_Uint32Path_GrpcClient_Transport(target, options...))
	t.uint64Path, err = errorx.Break[transportx.ClientTransport](err)(_Path_Uint64Path_GrpcClient_Transport(target, options...))
	t.floatPath, err = errorx.Break[transportx.ClientTransport](err)(_Path_FloatPath_GrpcClient_Transport(target, options...))
	t.doublePath, err = errorx.Break[transportx.ClientTransport](err)(_Path_DoublePath_GrpcClient_Transport(target, options...))
	t.stringPath, err = errorx.Break[transportx.ClientTransport](err)(_Path_StringPath_GrpcClient_Transport(target, options...))
	t.enumPath, err = errorx.Break[transportx.ClientTransport](err)(_Path_EnumPath_GrpcClient_Transport(target, options...))
	return t, err
}

type pathGrpcClient struct {
	endpoints PathEndpoints
}

func (c *pathGrpcClient) BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/BoolPath")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.BoolPath(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Int32Path")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.Int32Path(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Int64Path")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.Int64Path(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Uint32Path")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.Uint32Path(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/Uint64Path")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.Uint64Path(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/FloatPath")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.FloatPath(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/DoublePath")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.DoublePath(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/StringPath")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.StringPath(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGrpcClient) EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.Path/EnumPath")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.EnumPath(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewPathGrpcClient(transports PathClientTransports, middlewares ...endpoint.Middleware) PathService {
	endpoints := newPathClientEndpoints(transports, middlewares...)
	return &pathGrpcClient{endpoints: endpoints}
}

// =========================== grpc transport ===========================

func _Path_BoolPath_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.BoolPath(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/BoolPath")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_BoolPath_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"BoolPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_Int32Path_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.Int32Path(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/Int32Path")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_Int32Path_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"Int32Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_Int64Path_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.Int64Path(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/Int64Path")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_Int64Path_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"Int64Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_Uint32Path_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.Uint32Path(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/Uint32Path")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_Uint32Path_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"Uint32Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_Uint64Path_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.Uint64Path(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/Uint64Path")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_Uint64Path_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"Uint64Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_FloatPath_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.FloatPath(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/FloatPath")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_FloatPath_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"FloatPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_DoublePath_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.DoublePath(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/DoublePath")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_DoublePath_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"DoublePath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_StringPath_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.StringPath(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/StringPath")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_StringPath_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"StringPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}

func _Path_EnumPath_GrpcServer_Transport(endpoints PathEndpoints) *grpc.Server {
	return grpc.NewServer(
		endpoints.EnumPath(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.Path/EnumPath")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
	)
}

func _Path_EnumPath_GrpcClient_Transport(target string, options ...transportx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.Path",
				"EnumPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	}
}
