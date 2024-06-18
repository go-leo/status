// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package response

import (
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
	grpc1 "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	io "io"
	http1 "net/http"
	url "net/url"
)

// =========================== endpoints ===========================

type ResponseService interface {
	OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error)
	HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error)
}

type ResponseEndpoints interface {
	OmittedResponse(ctx context.Context) endpoint.Endpoint
	StarResponse(ctx context.Context) endpoint.Endpoint
	NamedResponse(ctx context.Context) endpoint.Endpoint
	HttpBodyResponse(ctx context.Context) endpoint.Endpoint
	HttpBodyNamedResponse(ctx context.Context) endpoint.Endpoint
}

type ResponseClientTransports interface {
	OmittedResponse() transportx.ClientTransport
	StarResponse() transportx.ClientTransport
	NamedResponse() transportx.ClientTransport
	HttpBodyResponse() transportx.ClientTransport
	HttpBodyNamedResponse() transportx.ClientTransport
}

type ResponseFactories interface {
	OmittedResponse(middlewares ...endpoint.Middleware) sd.Factory
	StarResponse(middlewares ...endpoint.Middleware) sd.Factory
	NamedResponse(middlewares ...endpoint.Middleware) sd.Factory
	HttpBodyResponse(middlewares ...endpoint.Middleware) sd.Factory
	HttpBodyNamedResponse(middlewares ...endpoint.Middleware) sd.Factory
}

type ResponseEndpointers interface {
	OmittedResponse() sd.Endpointer
	StarResponse() sd.Endpointer
	NamedResponse() sd.Endpointer
	HttpBodyResponse() sd.Endpointer
	HttpBodyNamedResponse() sd.Endpointer
}

type responseServerEndpoints struct {
	svc         ResponseService
	middlewares []endpoint.Middleware
}

func (e *responseServerEndpoints) OmittedResponse(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.OmittedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseServerEndpoints) StarResponse(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.StarResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseServerEndpoints) NamedResponse(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseServerEndpoints) HttpBodyResponse(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseServerEndpoints) HttpBodyNamedResponse(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyNamedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewResponseServerEndpoints(svc ResponseService, middlewares ...endpoint.Middleware) ResponseEndpoints {
	return &responseServerEndpoints{svc: svc, middlewares: middlewares}
}

type responseClientEndpoints struct {
	transports  ResponseClientTransports
	middlewares []endpoint.Middleware
}

func (e *responseClientEndpoints) OmittedResponse(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.OmittedResponse().Endpoint(ctx), e.middlewares...)
}

func (e *responseClientEndpoints) StarResponse(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.StarResponse().Endpoint(ctx), e.middlewares...)
}

func (e *responseClientEndpoints) NamedResponse(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.NamedResponse().Endpoint(ctx), e.middlewares...)
}

func (e *responseClientEndpoints) HttpBodyResponse(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyResponse().Endpoint(ctx), e.middlewares...)
}

func (e *responseClientEndpoints) HttpBodyNamedResponse(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyNamedResponse().Endpoint(ctx), e.middlewares...)
}

func NewResponseClientEndpoints(transports ResponseClientTransports, middlewares ...endpoint.Middleware) ResponseEndpoints {
	return &responseClientEndpoints{transports: transports, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc server ===========================

type ResponseGrpcServerTransports interface {
	OmittedResponse() *grpc.Server
	StarResponse() *grpc.Server
	NamedResponse() *grpc.Server
	HttpBodyResponse() *grpc.Server
	HttpBodyNamedResponse() *grpc.Server
}

type responseGrpcServerTransports struct {
	omittedResponse       *grpc.Server
	starResponse          *grpc.Server
	namedResponse         *grpc.Server
	httpBodyResponse      *grpc.Server
	httpBodyNamedResponse *grpc.Server
}

func (t *responseGrpcServerTransports) OmittedResponse() *grpc.Server {
	return t.omittedResponse
}

func (t *responseGrpcServerTransports) StarResponse() *grpc.Server {
	return t.starResponse
}

func (t *responseGrpcServerTransports) NamedResponse() *grpc.Server {
	return t.namedResponse
}

func (t *responseGrpcServerTransports) HttpBodyResponse() *grpc.Server {
	return t.httpBodyResponse
}

func (t *responseGrpcServerTransports) HttpBodyNamedResponse() *grpc.Server {
	return t.httpBodyNamedResponse
}

func NewResponseGrpcServerTransports(endpoints ResponseEndpoints) ResponseGrpcServerTransports {
	return &responseGrpcServerTransports{
		omittedResponse: grpc.NewServer(
			endpoints.OmittedResponse(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		starResponse: grpc.NewServer(
			endpoints.StarResponse(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/StarResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		namedResponse: grpc.NewServer(
			endpoints.NamedResponse(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		httpBodyResponse: grpc.NewServer(
			endpoints.HttpBodyResponse(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		httpBodyNamedResponse: grpc.NewServer(
			endpoints.HttpBodyNamedResponse(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
	}
}

type responseGrpcServer struct {
	omittedResponse       *grpc.Server
	starResponse          *grpc.Server
	namedResponse         *grpc.Server
	httpBodyResponse      *grpc.Server
	httpBodyNamedResponse *grpc.Server
}

func (s *responseGrpcServer) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.omittedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.starResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.namedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	ctx, rep, err := s.httpBodyResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*httpbody.HttpBody), nil
}

func (s *responseGrpcServer) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	ctx, rep, err := s.httpBodyNamedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*HttpBody), nil
}

func NewResponseGrpcServer(svc ResponseService, middlewares ...endpoint.Middleware) ResponseService {
	endpoints := NewResponseServerEndpoints(svc, middlewares...)
	transports := NewResponseGrpcServerTransports(endpoints)
	return &responseGrpcServer{
		omittedResponse:       transports.OmittedResponse(),
		starResponse:          transports.StarResponse(),
		namedResponse:         transports.NamedResponse(),
		httpBodyResponse:      transports.HttpBodyResponse(),
		httpBodyNamedResponse: transports.HttpBodyNamedResponse(),
	}
}

// =========================== grpc client ===========================

type responseGrpcClientTransports struct {
	omittedResponse       transportx.ClientTransport
	starResponse          transportx.ClientTransport
	namedResponse         transportx.ClientTransport
	httpBodyResponse      transportx.ClientTransport
	httpBodyNamedResponse transportx.ClientTransport
}

func (t *responseGrpcClientTransports) OmittedResponse() transportx.ClientTransport {
	return t.omittedResponse
}

func (t *responseGrpcClientTransports) StarResponse() transportx.ClientTransport {
	return t.starResponse
}

func (t *responseGrpcClientTransports) NamedResponse() transportx.ClientTransport {
	return t.namedResponse
}

func (t *responseGrpcClientTransports) HttpBodyResponse() transportx.ClientTransport {
	return t.httpBodyResponse
}

func (t *responseGrpcClientTransports) HttpBodyNamedResponse() transportx.ClientTransport {
	return t.httpBodyNamedResponse
}

func NewResponseGrpcClientTransports(
	target string,
	dialOption []grpc1.DialOption,
	options ...transportx.ClientTransportOption,
) (ResponseClientTransports, error) {
	t := &responseGrpcClientTransports{}
	var err error
	t.omittedResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				dialOption,
				"leo.example.response.v1.Response",
				"OmittedResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				UserResponse{},
				grpc.ClientBefore(grpcx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.starResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				dialOption,
				"leo.example.response.v1.Response",
				"StarResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				UserResponse{},
				grpc.ClientBefore(grpcx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.namedResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				dialOption,
				"leo.example.response.v1.Response",
				"NamedResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				UserResponse{},
				grpc.ClientBefore(grpcx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.httpBodyResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				dialOption,
				"leo.example.response.v1.Response",
				"HttpBodyResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				httpbody.HttpBody{},
				grpc.ClientBefore(grpcx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.httpBodyNamedResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				dialOption,
				"leo.example.response.v1.Response",
				"HttpBodyNamedResponse",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				HttpBody{},
				grpc.ClientBefore(grpcx.OutgoingMetadata),
			),
			options...,
		)
	})
	return t, err
}

type responseGrpcClient struct {
	endpoints ResponseEndpoints
}

func (c *responseGrpcClient) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/OmittedResponse")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.OmittedResponse(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*UserResponse), nil
}

func (c *responseGrpcClient) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/StarResponse")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.StarResponse(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*UserResponse), nil
}

func (c *responseGrpcClient) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/NamedResponse")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.NamedResponse(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*UserResponse), nil
}

func (c *responseGrpcClient) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/HttpBodyResponse")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.HttpBodyResponse(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*httpbody.HttpBody), nil
}

func (c *responseGrpcClient) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/HttpBodyNamedResponse")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.HttpBodyNamedResponse(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*HttpBody), nil
}

func NewResponseGrpcClient(transports ResponseClientTransports, middlewares ...endpoint.Middleware) ResponseService {
	endpoints := NewResponseClientEndpoints(transports, middlewares...)
	return &responseGrpcClient{endpoints: endpoints}
}

// =========================== http server ===========================

type ResponseHttpServerTransports interface {
	OmittedResponse() *http.Server
	StarResponse() *http.Server
	NamedResponse() *http.Server
	HttpBodyResponse() *http.Server
	HttpBodyNamedResponse() *http.Server
}

type responseHttpServerTransports struct {
	omittedResponse       *http.Server
	starResponse          *http.Server
	namedResponse         *http.Server
	httpBodyResponse      *http.Server
	httpBodyNamedResponse *http.Server
}

func (t *responseHttpServerTransports) OmittedResponse() *http.Server {
	return t.omittedResponse
}

func (t *responseHttpServerTransports) StarResponse() *http.Server {
	return t.starResponse
}

func (t *responseHttpServerTransports) NamedResponse() *http.Server {
	return t.namedResponse
}

func (t *responseHttpServerTransports) HttpBodyResponse() *http.Server {
	return t.httpBodyResponse
}

func (t *responseHttpServerTransports) HttpBodyNamedResponse() *http.Server {
	return t.httpBodyNamedResponse
}

func NewResponseHttpServerTransports(endpoints ResponseEndpoints) ResponseHttpServerTransports {
	return &responseHttpServerTransports{
		omittedResponse: http.NewServer(
			endpoints.OmittedResponse(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*UserResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.Wrap(err)
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		starResponse: http.NewServer(
			endpoints.StarResponse(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*UserResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.Wrap(err)
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/StarResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		namedResponse: http.NewServer(
			endpoints.NamedResponse(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*UserResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp.GetUser()); err != nil {
					return statusx.ErrInternal.Wrap(err)
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		httpBodyResponse: http.NewServer(
			endpoints.HttpBodyResponse(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*httpbody.HttpBody)
				w.Header().Set("Content-Type", resp.GetContentType())
				for _, src := range resp.GetExtensions() {
					dst, err := anypb.UnmarshalNew(src, proto.UnmarshalOptions{})
					if err != nil {
						return statusx.ErrInternal.Wrap(err)
					}
					metadata, ok := dst.(*structpb.Struct)
					if !ok {
						continue
					}
					for key, value := range metadata.GetFields() {
						w.Header().Add(key, string(errorx.Ignore(jsonx.Marshal(value))))
					}
				}
				w.WriteHeader(http1.StatusOK)
				if _, err := w.Write(resp.GetData()); err != nil {
					return statusx.ErrInternal.Wrap(err)
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		httpBodyNamedResponse: http.NewServer(
			endpoints.HttpBodyNamedResponse(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*HttpBody)
				w.Header().Set("Content-Type", resp.GetBody().GetContentType())
				for _, src := range resp.GetBody().GetExtensions() {
					dst, err := anypb.UnmarshalNew(src, proto.UnmarshalOptions{})
					if err != nil {
						return statusx.ErrInternal.Wrap(err)
					}
					metadata, ok := dst.(*structpb.Struct)
					if !ok {
						continue
					}
					for key, value := range metadata.GetFields() {
						w.Header().Add(key, string(errorx.Ignore(jsonx.Marshal(value))))
					}
				}
				w.WriteHeader(http1.StatusOK)
				if _, err := w.Write(resp.GetBody().GetData()); err != nil {
					return statusx.ErrInternal.Wrap(err)
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
	}
}

func NewResponseHttpServerHandler(svc ResponseService, middlewares ...endpoint.Middleware) http1.Handler {
	endpoints := NewResponseServerEndpoints(svc, middlewares...)
	transports := NewResponseHttpServerTransports(endpoints)
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.response.v1.Response/OmittedResponse").Methods("POST").Path("/v1/omitted/response").Handler(transports.OmittedResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/StarResponse").Methods("POST").Path("/v1/star/response").Handler(transports.StarResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/NamedResponse").Methods("POST").Path("/v1/named/response").Handler(transports.NamedResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyResponse").Methods("PUT").Path("/v1/http/body/omitted/response").Handler(transports.HttpBodyResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyNamedResponse").Methods("PUT").Path("/v1/http/body/named/response").Handler(transports.HttpBodyNamedResponse())
	return router
}

// =========================== http client ===========================

type responseHttpClientTransports struct {
	omittedResponse       transportx.ClientTransport
	starResponse          transportx.ClientTransport
	namedResponse         transportx.ClientTransport
	httpBodyResponse      transportx.ClientTransport
	httpBodyNamedResponse transportx.ClientTransport
}

func (t *responseHttpClientTransports) OmittedResponse() transportx.ClientTransport {
	return t.omittedResponse
}

func (t *responseHttpClientTransports) StarResponse() transportx.ClientTransport {
	return t.starResponse
}

func (t *responseHttpClientTransports) NamedResponse() transportx.ClientTransport {
	return t.namedResponse
}

func (t *responseHttpClientTransports) HttpBodyResponse() transportx.ClientTransport {
	return t.httpBodyResponse
}

func (t *responseHttpClientTransports) HttpBodyNamedResponse() transportx.ClientTransport {
	return t.httpBodyNamedResponse
}

func NewResponseHttpClientTransports(
	target string,
	scheme string,
	options ...transportx.ClientTransportOption,
) (ResponseClientTransports, error) {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.response.v1.Response/OmittedResponse").Methods("POST").Path("/v1/omitted/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/StarResponse").Methods("POST").Path("/v1/star/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/NamedResponse").Methods("POST").Path("/v1/named/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyResponse").Methods("PUT").Path("/v1/http/body/omitted/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyNamedResponse").Methods("PUT").Path("/v1/http/body/named/response")
	t := &responseHttpClientTransports{}
	var err error
	t.omittedResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				scheme,
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
						path, err := router.Get("/leo.example.response.v1.Response/OmittedResponse").URLPath(pairs...)
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
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &UserResponse{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.starResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				scheme,
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
						path, err := router.Get("/leo.example.response.v1.Response/StarResponse").URLPath(pairs...)
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
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &UserResponse{}
					if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.namedResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				scheme,
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
						path, err := router.Get("/leo.example.response.v1.Response/NamedResponse").URLPath(pairs...)
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
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &UserResponse{}
					if err := jsonx.NewDecoder(r.Body).Decode(&resp.User); err != nil {
						return nil, err
					}
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.httpBodyResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				scheme,
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
						path, err := router.Get("/leo.example.response.v1.Response/HttpBodyResponse").URLPath(pairs...)
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
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &httpbody.HttpBody{}
					resp.ContentType = r.Header.Get("Content-Type")
					body, err := io.ReadAll(r.Body)
					if err != nil {
						return nil, err
					}
					resp.Data = body
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.httpBodyNamedResponse, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				scheme,
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
						path, err := router.Get("/leo.example.response.v1.Response/HttpBodyNamedResponse").URLPath(pairs...)
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
						return r, nil
					}
				},
				func(ctx context.Context, r *http1.Response) (any, error) {
					if httpx.IsErrorResponse(r) {
						return nil, httpx.ErrorDecoder(ctx, r)
					}
					resp := &HttpBody{}
					resp.Body = &httpbody.HttpBody{}
					resp.Body.ContentType = r.Header.Get("Content-Type")
					body, err := io.ReadAll(r.Body)
					if err != nil {
						return nil, err
					}
					resp.Body.Data = body
					return resp, nil
				},
				http.ClientBefore(httpx.OutgoingMetadata),
			),
			options...,
		)
	})
	return t, err
}

type responseHttpClient struct {
	endpoints ResponseEndpoints
}

func (c *responseHttpClient) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/OmittedResponse")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.OmittedResponse(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseHttpClient) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/StarResponse")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.StarResponse(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseHttpClient) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/NamedResponse")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.NamedResponse(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseHttpClient) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/HttpBodyResponse")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.HttpBodyResponse(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*httpbody.HttpBody), nil
}

func (c *responseHttpClient) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.response.v1.Response/HttpBodyNamedResponse")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.HttpBodyNamedResponse(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*HttpBody), nil
}

func NewResponseHttpClient(transports ResponseClientTransports, middlewares ...endpoint.Middleware) ResponseService {
	endpoints := NewResponseClientEndpoints(transports, middlewares...)
	return &responseGrpcClient{endpoints: endpoints}
}
