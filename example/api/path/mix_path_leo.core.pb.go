// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package path

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	transportx "github.com/go-leo/leo/v3/transportx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

type MixPathService interface {
	MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error)
}

type MixPathEndpoints interface {
	MixPath(ctx context.Context) endpoint.Endpoint
}

type MixPathClientTransports interface {
	MixPath() transportx.ClientTransport
}

type MixPathClientTransportsV2 interface {
	MixPath(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

type MixPathFactories interface {
	MixPath(ctx context.Context) sd.Factory
}

type MixPathEndpointers interface {
	MixPath() sd.Endpointer
}

type mixPathServerEndpoints struct {
	svc         MixPathService
	middlewares []endpoint.Middleware
}

func (e *mixPathServerEndpoints) MixPath(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.MixPath(ctx, request.(*MixPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func newMixPathServerEndpoints(svc MixPathService, middlewares ...endpoint.Middleware) MixPathEndpoints {
	return &mixPathServerEndpoints{svc: svc, middlewares: middlewares}
}

type mixPathClientEndpoints struct {
	transports  MixPathClientTransports
	middlewares []endpoint.Middleware
}

func (e *mixPathClientEndpoints) MixPath(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.MixPath().Endpoint(ctx), e.middlewares...)
}

func newMixPathClientEndpoints(transports MixPathClientTransports, middlewares ...endpoint.Middleware) MixPathEndpoints {
	return &mixPathClientEndpoints{transports: transports, middlewares: middlewares}
}

type mixPathFactories struct {
	transports MixPathClientTransportsV2
}

func (f *mixPathFactories) MixPath(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.MixPath(ctx, instance)
	}
}

func newMixPathFactories(transports MixPathClientTransportsV2) MixPathFactories {
	return &mixPathFactories{transports: transports}
}
