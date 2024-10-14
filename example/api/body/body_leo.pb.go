// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package body

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
	url "net/url"
)

// =========================== endpoints ===========================

type BodyService interface {
	StarBody(ctx context.Context, request *User) (*emptypb.Empty, error)
	NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error)
	NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error)
	HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error)
	HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error)
}

type BodyEndpoints interface {
	StarBody(ctx context.Context) endpoint.Endpoint
	NamedBody(ctx context.Context) endpoint.Endpoint
	NonBody(ctx context.Context) endpoint.Endpoint
	HttpBodyStarBody(ctx context.Context) endpoint.Endpoint
	HttpBodyNamedBody(ctx context.Context) endpoint.Endpoint
}

type BodyClientTransports interface {
	StarBody() transportx.ClientTransport
	NamedBody() transportx.ClientTransport
	NonBody() transportx.ClientTransport
	HttpBodyStarBody() transportx.ClientTransport
	HttpBodyNamedBody() transportx.ClientTransport
}

type BodyFactories interface {
	StarBody(middlewares ...endpoint.Middleware) sd.Factory
	NamedBody(middlewares ...endpoint.Middleware) sd.Factory
	NonBody(middlewares ...endpoint.Middleware) sd.Factory
	HttpBodyStarBody(middlewares ...endpoint.Middleware) sd.Factory
	HttpBodyNamedBody(middlewares ...endpoint.Middleware) sd.Factory
}

type BodyEndpointers interface {
	StarBody() sd.Endpointer
	NamedBody() sd.Endpointer
	NonBody() sd.Endpointer
	HttpBodyStarBody() sd.Endpointer
	HttpBodyNamedBody() sd.Endpointer
}

type bodyServerEndpoints struct {
	svc         BodyService
	middlewares []endpoint.Middleware
}

func (e *bodyServerEndpoints) StarBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.StarBody(ctx, request.(*User))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyServerEndpoints) NamedBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedBody(ctx, request.(*UserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyServerEndpoints) NonBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NonBody(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyServerEndpoints) HttpBodyStarBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyStarBody(ctx, request.(*httpbody.HttpBody))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyServerEndpoints) HttpBodyNamedBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyNamedBody(ctx, request.(*HttpBody))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func newBodyServerEndpoints(svc BodyService, middlewares ...endpoint.Middleware) BodyEndpoints {
	return &bodyServerEndpoints{svc: svc, middlewares: middlewares}
}

type bodyClientEndpoints struct {
	transports  BodyClientTransports
	middlewares []endpoint.Middleware
}

func (e *bodyClientEndpoints) StarBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.StarBody().Endpoint(ctx), e.middlewares...)
}

func (e *bodyClientEndpoints) NamedBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.NamedBody().Endpoint(ctx), e.middlewares...)
}

func (e *bodyClientEndpoints) NonBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.NonBody().Endpoint(ctx), e.middlewares...)
}

func (e *bodyClientEndpoints) HttpBodyStarBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyStarBody().Endpoint(ctx), e.middlewares...)
}

func (e *bodyClientEndpoints) HttpBodyNamedBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyNamedBody().Endpoint(ctx), e.middlewares...)
}

func newBodyClientEndpoints(transports BodyClientTransports, middlewares ...endpoint.Middleware) BodyEndpoints {
	return &bodyClientEndpoints{transports: transports, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc server ===========================

type BodyGrpcServerTransports interface {
	StarBody() *grpc.Server
	NamedBody() *grpc.Server
	NonBody() *grpc.Server
	HttpBodyStarBody() *grpc.Server
	HttpBodyNamedBody() *grpc.Server
}

type bodyGrpcServerTransports struct {
	starBody          *grpc.Server
	namedBody         *grpc.Server
	nonBody           *grpc.Server
	httpBodyStarBody  *grpc.Server
	httpBodyNamedBody *grpc.Server
}

func (t *bodyGrpcServerTransports) StarBody() *grpc.Server {
	return t.starBody
}

func (t *bodyGrpcServerTransports) NamedBody() *grpc.Server {
	return t.namedBody
}

func (t *bodyGrpcServerTransports) NonBody() *grpc.Server {
	return t.nonBody
}

func (t *bodyGrpcServerTransports) HttpBodyStarBody() *grpc.Server {
	return t.httpBodyStarBody
}

func (t *bodyGrpcServerTransports) HttpBodyNamedBody() *grpc.Server {
	return t.httpBodyNamedBody
}

func newBodyGrpcServerTransports(endpoints BodyEndpoints) BodyGrpcServerTransports {
	return &bodyGrpcServerTransports{
		starBody: grpc.NewServer(
			endpoints.StarBody(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/StarBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		),
		namedBody: grpc.NewServer(
			endpoints.NamedBody(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/NamedBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		),
		nonBody: grpc.NewServer(
			endpoints.NonBody(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/NonBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		),
		httpBodyStarBody: grpc.NewServer(
			endpoints.HttpBodyStarBody(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/HttpBodyStarBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		),
		httpBodyNamedBody: grpc.NewServer(
			endpoints.HttpBodyNamedBody(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/HttpBodyNamedBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		),
	}
}

type bodyGrpcServer struct {
	starBody          *grpc.Server
	namedBody         *grpc.Server
	nonBody           *grpc.Server
	httpBodyStarBody  *grpc.Server
	httpBodyNamedBody *grpc.Server
}

func (s *bodyGrpcServer) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx, rep, err := s.starBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.namedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, rep, err := s.nonBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyStarBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyNamedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewBodyGrpcServer(svc BodyService, middlewares ...endpoint.Middleware) BodyService {
	endpoints := newBodyServerEndpoints(svc, middlewares...)
	transports := newBodyGrpcServerTransports(endpoints)
	return &bodyGrpcServer{
		starBody:          transports.StarBody(),
		namedBody:         transports.NamedBody(),
		nonBody:           transports.NonBody(),
		httpBodyStarBody:  transports.HttpBodyStarBody(),
		httpBodyNamedBody: transports.HttpBodyNamedBody(),
	}
}

// =========================== grpc client ===========================

type bodyGrpcClientTransports struct {
	starBody          transportx.ClientTransport
	namedBody         transportx.ClientTransport
	nonBody           transportx.ClientTransport
	httpBodyStarBody  transportx.ClientTransport
	httpBodyNamedBody transportx.ClientTransport
}

func (t *bodyGrpcClientTransports) StarBody() transportx.ClientTransport {
	return t.starBody
}

func (t *bodyGrpcClientTransports) NamedBody() transportx.ClientTransport {
	return t.namedBody
}

func (t *bodyGrpcClientTransports) NonBody() transportx.ClientTransport {
	return t.nonBody
}

func (t *bodyGrpcClientTransports) HttpBodyStarBody() transportx.ClientTransport {
	return t.httpBodyStarBody
}

func (t *bodyGrpcClientTransports) HttpBodyNamedBody() transportx.ClientTransport {
	return t.httpBodyNamedBody
}

func NewBodyGrpcClientTransports(target string, options ...transportx.ClientTransportOption) (BodyClientTransports, error) {
	t := &bodyGrpcClientTransports{}
	var err error
	t.starBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.body.v1.Body",
				"StarBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	t.namedBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.body.v1.Body",
				"NamedBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	t.nonBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.body.v1.Body",
				"NonBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	t.httpBodyStarBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.body.v1.Body",
				"HttpBodyStarBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	t.httpBodyNamedBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.body.v1.Body",
				"HttpBodyNamedBody",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	return t, err
}

type bodyGrpcClient struct {
	endpoints BodyEndpoints
}

func (c *bodyGrpcClient) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/StarBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.StarBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NamedBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.NamedBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NonBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.NonBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyStarBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.HttpBodyStarBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyNamedBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.HttpBodyNamedBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewBodyGrpcClient(transports BodyClientTransports, middlewares ...endpoint.Middleware) BodyService {
	endpoints := newBodyClientEndpoints(transports, middlewares...)
	return &bodyGrpcClient{endpoints: endpoints}
}

// =========================== http server ===========================

type BodyHttpServerTransports interface {
	StarBody() *http.Server
	NamedBody() *http.Server
	NonBody() *http.Server
	HttpBodyStarBody() *http.Server
	HttpBodyNamedBody() *http.Server
}

type bodyHttpServerTransports struct {
	starBody          *http.Server
	namedBody         *http.Server
	nonBody           *http.Server
	httpBodyStarBody  *http.Server
	httpBodyNamedBody *http.Server
}

func (t *bodyHttpServerTransports) StarBody() *http.Server {
	return t.starBody
}

func (t *bodyHttpServerTransports) NamedBody() *http.Server {
	return t.namedBody
}

func (t *bodyHttpServerTransports) NonBody() *http.Server {
	return t.nonBody
}

func (t *bodyHttpServerTransports) HttpBodyStarBody() *http.Server {
	return t.httpBodyStarBody
}

func (t *bodyHttpServerTransports) HttpBodyNamedBody() *http.Server {
	return t.httpBodyNamedBody
}

func newBodyHttpServerTransports(endpoints BodyEndpoints) BodyHttpServerTransports {
	return &bodyHttpServerTransports{
		starBody: http.NewServer(
			endpoints.StarBody(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &User{}
				if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
					return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.With(statusx.Wrap(err))
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/StarBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadataInjector),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
			http.ServerBefore(httpx.TimeoutController),
			http.ServerFinalizer(httpx.CancelInvoker),
		),
		namedBody: http.NewServer(
			endpoints.NamedBody(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &UserRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(&req.User); err != nil {
					return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.With(statusx.Wrap(err))
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/NamedBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadataInjector),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
			http.ServerBefore(httpx.TimeoutController),
			http.ServerFinalizer(httpx.CancelInvoker),
		),
		nonBody: http.NewServer(
			endpoints.NonBody(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.With(statusx.Wrap(err))
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/NonBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadataInjector),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
			http.ServerBefore(httpx.TimeoutController),
			http.ServerFinalizer(httpx.CancelInvoker),
		),
		httpBodyStarBody: http.NewServer(
			endpoints.HttpBodyStarBody(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &httpbody.HttpBody{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
				}
				req.Data = body
				req.ContentType = r.Header.Get("Content-Type")
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.With(statusx.Wrap(err))
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/HttpBodyStarBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadataInjector),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
			http.ServerBefore(httpx.TimeoutController),
			http.ServerFinalizer(httpx.CancelInvoker),
		),
		httpBodyNamedBody: http.NewServer(
			endpoints.HttpBodyNamedBody(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &HttpBody{}
				req.Body = &httpbody.HttpBody{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
				}
				req.Body.Data = body
				req.Body.ContentType = r.Header.Get("Content-Type")
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.With(statusx.Wrap(err))
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/HttpBodyNamedBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadataInjector),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
			http.ServerBefore(httpx.TimeoutController),
			http.ServerFinalizer(httpx.CancelInvoker),
		),
	}
}

func AppendBodyHttpRouter(router *mux.Router, svc BodyService, middlewares ...endpoint.Middleware) *mux.Router {
	endpoints := newBodyServerEndpoints(svc, middlewares...)
	transports := newBodyHttpServerTransports(endpoints)
	router.NewRoute().Name("/leo.example.body.v1.Body/StarBody").Methods("POST").Path("/v1/star/body").Handler(transports.StarBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/NamedBody").Methods("POST").Path("/v1/named/body").Handler(transports.NamedBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/NonBody").Methods("GET").Path("/v1/user_body").Handler(transports.NonBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyStarBody").Methods("PUT").Path("/v1/http/body/star/body").Handler(transports.HttpBodyStarBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyNamedBody").Methods("PUT").Path("/v1/http/body/named/body").Handler(transports.HttpBodyNamedBody())
	return router
}

// =========================== http client ===========================

type bodyHttpClientTransports struct {
	starBody          transportx.ClientTransport
	namedBody         transportx.ClientTransport
	nonBody           transportx.ClientTransport
	httpBodyStarBody  transportx.ClientTransport
	httpBodyNamedBody transportx.ClientTransport
}

func (t *bodyHttpClientTransports) StarBody() transportx.ClientTransport {
	return t.starBody
}

func (t *bodyHttpClientTransports) NamedBody() transportx.ClientTransport {
	return t.namedBody
}

func (t *bodyHttpClientTransports) NonBody() transportx.ClientTransport {
	return t.nonBody
}

func (t *bodyHttpClientTransports) HttpBodyStarBody() transportx.ClientTransport {
	return t.httpBodyStarBody
}

func (t *bodyHttpClientTransports) HttpBodyNamedBody() transportx.ClientTransport {
	return t.httpBodyNamedBody
}

func NewBodyHttpClientTransports(target string, options ...transportx.ClientTransportOption) (BodyClientTransports, error) {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.body.v1.Body/StarBody").Methods("POST").Path("/v1/star/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/NamedBody").Methods("POST").Path("/v1/named/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/NonBody").Methods("GET").Path("/v1/user_body")
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyStarBody").Methods("PUT").Path("/v1/http/body/star/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyNamedBody").Methods("PUT").Path("/v1/http/body/named/body")
	t := &bodyHttpClientTransports{}
	var err error
	t.starBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				func(scheme string, instance string) http.CreateRequestFunc {
					return func(ctx context.Context, obj any) (*http1.Request, error) {
						if obj == nil {
							return nil, errors.New("request object is nil")
						}
						req, ok := obj.(*User)
						if !ok {
							return nil, fmt.Errorf("invalid request object type, %T", obj)
						}
						_ = req
						var body io.Reader
						var bodyBuf bytes.Buffer
						if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
							return nil, err
						}
						body = &bodyBuf
						contentType := "application/json; charset=utf-8"
						var pairs []string
						path, err := router.Get("/leo.example.body.v1.Body/StarBody").URLPath(pairs...)
						if err != nil {
							return nil, err
						}
						queries := url.Values{}
						target := &url.URL{
							Scheme:   scheme,
							Host:     instance,
							Path:     path.Path,
							RawQuery: queries.Encode(),
						}
						r, err := http1.NewRequestWithContext(ctx, "POST", target.String(), body)
						if err != nil {
							return nil, err
						}
						r.Header.Set("Content-Type", contentType)
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &emptypb.Empty{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	t.namedBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				func(scheme string, instance string) http.CreateRequestFunc {
					return func(ctx context.Context, obj any) (*http1.Request, error) {
						if obj == nil {
							return nil, errors.New("request object is nil")
						}
						req, ok := obj.(*UserRequest)
						if !ok {
							return nil, fmt.Errorf("invalid request object type, %T", obj)
						}
						_ = req
						var body io.Reader
						var bodyBuf bytes.Buffer
						if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetUser()); err != nil {
							return nil, err
						}
						body = &bodyBuf
						contentType := "application/json; charset=utf-8"
						var pairs []string
						path, err := router.Get("/leo.example.body.v1.Body/NamedBody").URLPath(pairs...)
						if err != nil {
							return nil, err
						}
						queries := url.Values{}
						target := &url.URL{
							Scheme:   scheme,
							Host:     instance,
							Path:     path.Path,
							RawQuery: queries.Encode(),
						}
						r, err := http1.NewRequestWithContext(ctx, "POST", target.String(), body)
						if err != nil {
							return nil, err
						}
						r.Header.Set("Content-Type", contentType)
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &emptypb.Empty{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadataInjector),
				http.ClientBefore(httpx.TimeoutController),
			),
			options...,
		)
	})
	t.nonBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				func(scheme string, instance string) http.CreateRequestFunc {
					return func(ctx context.Context, obj any) (*http1.Request, error) {
						if obj == nil {
							return nil, errors.New("request object is nil")
						}
						req, ok := obj.(*emptypb.Empty)
						if !ok {
							return nil, fmt.Errorf("invalid request object type, %T", obj)
						}
						_ = req
						var body io.Reader
						var pairs []string
						path, err := router.Get("/leo.example.body.v1.Body/NonBody").URLPath(pairs...)
						if err != nil {
							return nil, err
						}
						queries := url.Values{}
						target := &url.URL{
							Scheme:   scheme,
							Host:     instance,
							Path:     path.Path,
							RawQuery: queries.Encode(),
						}
						r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
						if err != nil {
							return nil, err
						}
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &emptypb.Empty{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	t.httpBodyStarBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				func(scheme string, instance string) http.CreateRequestFunc {
					return func(ctx context.Context, obj any) (*http1.Request, error) {
						if obj == nil {
							return nil, errors.New("request object is nil")
						}
						req, ok := obj.(*httpbody.HttpBody)
						if !ok {
							return nil, fmt.Errorf("invalid request object type, %T", obj)
						}
						_ = req
						var body io.Reader
						body = bytes.NewReader(req.GetData())
						contentType := req.GetContentType()
						var pairs []string
						path, err := router.Get("/leo.example.body.v1.Body/HttpBodyStarBody").URLPath(pairs...)
						if err != nil {
							return nil, err
						}
						queries := url.Values{}
						target := &url.URL{
							Scheme:   scheme,
							Host:     instance,
							Path:     path.Path,
							RawQuery: queries.Encode(),
						}
						r, err := http1.NewRequestWithContext(ctx, "PUT", target.String(), body)
						if err != nil {
							return nil, err
						}
						r.Header.Set("Content-Type", contentType)
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &emptypb.Empty{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	t.httpBodyNamedBody, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				func(scheme string, instance string) http.CreateRequestFunc {
					return func(ctx context.Context, obj any) (*http1.Request, error) {
						if obj == nil {
							return nil, errors.New("request object is nil")
						}
						req, ok := obj.(*HttpBody)
						if !ok {
							return nil, fmt.Errorf("invalid request object type, %T", obj)
						}
						_ = req
						var body io.Reader
						body = bytes.NewReader(req.GetBody().GetData())
						contentType := req.GetBody().GetContentType()
						var pairs []string
						path, err := router.Get("/leo.example.body.v1.Body/HttpBodyNamedBody").URLPath(pairs...)
						if err != nil {
							return nil, err
						}
						queries := url.Values{}
						target := &url.URL{
							Scheme:   scheme,
							Host:     instance,
							Path:     path.Path,
							RawQuery: queries.Encode(),
						}
						r, err := http1.NewRequestWithContext(ctx, "PUT", target.String(), body)
						if err != nil {
							return nil, err
						}
						r.Header.Set("Content-Type", contentType)
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &emptypb.Empty{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadataInjector),
			),
			options...,
		)
	})
	return t, err
}

type bodyHttpClient struct {
	endpoints BodyEndpoints
}

func (c *bodyHttpClient) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/StarBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.StarBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NamedBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.NamedBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NonBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.NonBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyStarBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.HttpBodyStarBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyNamedBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.HttpBodyNamedBody(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewBodyHttpClient(transports BodyClientTransports, middlewares ...endpoint.Middleware) BodyService {
	endpoints := newBodyClientEndpoints(transports, middlewares...)
	return &bodyHttpClient{endpoints: endpoints}
}
