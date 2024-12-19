// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package query

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	log "github.com/go-kit/log"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	transportx "github.com/go-leo/leo/v3/transportx"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

type QueryService interface {
	Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
}

type QueryEndpoints interface {
	Query(ctx context.Context) endpoint.Endpoint
}

type QueryClientTransports interface {
	Query() transportx.ClientTransport
}
type QueryClientTransportsV2 interface {
	Query(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

type QueryFactories interface {
	Query(ctx context.Context) sd.Factory
}

type QueryEndpointers interface {
	Query(ctx context.Context) sd.Endpointer
}

type queryServerEndpoints struct {
	svc         QueryService
	middlewares []endpoint.Middleware
}

func (e *queryServerEndpoints) Query(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.Query(ctx, request.(*QueryRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func newQueryServerEndpoints(svc QueryService, middlewares ...endpoint.Middleware) QueryEndpoints {
	return &queryServerEndpoints{svc: svc, middlewares: middlewares}
}

type queryClientEndpoints struct {
	transports  QueryClientTransports
	middlewares []endpoint.Middleware
}

func (e *queryClientEndpoints) Query(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.Query().Endpoint(ctx), e.middlewares...)
}
func newQueryClientEndpoints(transports QueryClientTransports, middlewares ...endpoint.Middleware) QueryEndpoints {
	return &queryClientEndpoints{transports: transports, middlewares: middlewares}
}

type queryFactories struct {
	transports QueryClientTransportsV2
}

func (f *queryFactories) Query(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.Query(ctx, instance)
	}
}
func newQueryFactories(transports QueryClientTransportsV2) QueryFactories {
	return &queryFactories{transports: transports}
}

type queryEndpointers struct {
	instancer sd.Instancer
	factories QueryFactories
	logger    log.Logger
	options   []sd.EndpointerOption
}

func (e *queryEndpointers) Query(ctx context.Context) sd.Endpointer {
	return sd.NewEndpointer(e.instancer, e.factories.Query(ctx), e.logger, e.options...)
}
func newQueryEndpointers(instancer sd.Instancer, factories QueryFactories, logger log.Logger, options ...sd.EndpointerOption) QueryEndpointers {
	return &queryEndpointers{instancer: instancer, factories: factories, logger: logger, options: options}
}
