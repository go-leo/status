// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cqrs

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
	cqrs "github.com/go-leo/leo/v3/cqrs"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	command "github.com/go-leo/leo/v3/example/internal/cqrs/command"
	query "github.com/go-leo/leo/v3/example/internal/cqrs/query"
	metadatax "github.com/go-leo/leo/v3/metadatax"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
	url "net/url"
)

// =========================== endpoints ===========================

type CQRSService interface {
	CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error)
	FindUser(ctx context.Context, request *FindUserRequest) (*GetUserResponse, error)
}

type CQRSEndpoints interface {
	CreateUser(ctx context.Context) endpoint.Endpoint
	FindUser(ctx context.Context) endpoint.Endpoint
}

type CQRSClientTransports interface {
	CreateUser() transportx.ClientTransport
	FindUser() transportx.ClientTransport
}

type CQRSFactories interface {
	CreateUser(middlewares ...endpoint.Middleware) sd.Factory
	FindUser(middlewares ...endpoint.Middleware) sd.Factory
}

type CQRSEndpointers interface {
	CreateUser() sd.Endpointer
	FindUser() sd.Endpointer
}

type cQRSServerEndpoints struct {
	svc         CQRSService
	middlewares []endpoint.Middleware
}

func (e *cQRSServerEndpoints) CreateUser(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateUser(ctx, request.(*CreateUserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *cQRSServerEndpoints) FindUser(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.FindUser(ctx, request.(*FindUserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func newCQRSServerEndpoints(svc CQRSService, middlewares ...endpoint.Middleware) CQRSEndpoints {
	return &cQRSServerEndpoints{svc: svc, middlewares: middlewares}
}

type cQRSClientEndpoints struct {
	transports  CQRSClientTransports
	middlewares []endpoint.Middleware
}

func (e *cQRSClientEndpoints) CreateUser(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.CreateUser().Endpoint(ctx), e.middlewares...)
}

func (e *cQRSClientEndpoints) FindUser(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.FindUser().Endpoint(ctx), e.middlewares...)
}

func newCQRSClientEndpoints(transports CQRSClientTransports, middlewares ...endpoint.Middleware) CQRSEndpoints {
	return &cQRSClientEndpoints{transports: transports, middlewares: middlewares}
}

// =========================== cqrs ===========================

// CQRSAssembler responsible for completing the transformation between domain model objects and DTOs
type CQRSAssembler interface {

	// FromCreateUserRequest convert request to command arguments
	FromCreateUserRequest(ctx context.Context, request *CreateUserRequest) (*command.CreateUserArgs, context.Context, error)

	// ToCreateUserResponse convert query result to response
	ToCreateUserResponse(ctx context.Context, request *CreateUserRequest, metadata metadatax.Metadata) (*emptypb.Empty, error)

	// FromFindUserRequest convert request to query arguments
	FromFindUserRequest(ctx context.Context, request *FindUserRequest) (*query.FindUserArgs, context.Context, error)

	// ToFindUserResponse convert query result to response
	ToFindUserResponse(ctx context.Context, request *FindUserRequest, res *query.FindUserRes) (*GetUserResponse, error)
}

// cQRSCqrsService implement the CQRSService with CQRS pattern
type cQRSCqrsService struct {
	bus       cqrs.Bus
	assembler CQRSAssembler
}

func (svc *cQRSCqrsService) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	args, ctx, err := svc.assembler.FromCreateUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	metadata, err := svc.bus.Exec(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToCreateUserResponse(ctx, request, metadata)
}

func (svc *cQRSCqrsService) FindUser(ctx context.Context, request *FindUserRequest) (*GetUserResponse, error) {
	args, ctx, err := svc.assembler.FromFindUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	res, err := svc.bus.Query(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToFindUserResponse(ctx, request, res.(*query.FindUserRes))
}

func NewCQRSCqrsService(bus cqrs.Bus, assembler CQRSAssembler) CQRSService {
	return &cQRSCqrsService{bus: bus, assembler: assembler}
}

func NewCQRSBus(
	createUser command.CreateUser,
	findUser query.FindUser,
) (cqrs.Bus, error) {
	bus := cqrs.NewBus()
	if err := bus.RegisterCommand(createUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterQuery(findUser); err != nil {
		return nil, err
	}
	return bus, nil
}

// =========================== grpc server ===========================

type CQRSGrpcServerTransports interface {
	CreateUser() *grpc.Server
	FindUser() *grpc.Server
}

type cQRSGrpcServerTransports struct {
	createUser *grpc.Server
	findUser   *grpc.Server
}

func (t *cQRSGrpcServerTransports) CreateUser() *grpc.Server {
	return t.createUser
}

func (t *cQRSGrpcServerTransports) FindUser() *grpc.Server {
	return t.findUser
}

func newCQRSGrpcServerTransports(endpoints CQRSEndpoints) CQRSGrpcServerTransports {
	return &cQRSGrpcServerTransports{
		createUser: grpc.NewServer(
			endpoints.CreateUser(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/pb.CQRS/CreateUser")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		findUser: grpc.NewServer(
			endpoints.FindUser(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/pb.CQRS/FindUser")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
	}
}

type cQRSGrpcServer struct {
	createUser *grpc.Server
	findUser   *grpc.Server
}

func (s *cQRSGrpcServer) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.createUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *cQRSGrpcServer) FindUser(ctx context.Context, request *FindUserRequest) (*GetUserResponse, error) {
	ctx, rep, err := s.findUser.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*GetUserResponse), nil
}

func NewCQRSGrpcServer(svc CQRSService, middlewares ...endpoint.Middleware) CQRSService {
	endpoints := newCQRSServerEndpoints(svc, middlewares...)
	transports := newCQRSGrpcServerTransports(endpoints)
	return &cQRSGrpcServer{
		createUser: transports.CreateUser(),
		findUser:   transports.FindUser(),
	}
}

// =========================== grpc client ===========================

type cQRSGrpcClientTransports struct {
	createUser transportx.ClientTransport
	findUser   transportx.ClientTransport
}

func (t *cQRSGrpcClientTransports) CreateUser() transportx.ClientTransport {
	return t.createUser
}

func (t *cQRSGrpcClientTransports) FindUser() transportx.ClientTransport {
	return t.findUser
}

func NewCQRSGrpcClientTransports(target string, options ...transportx.ClientTransportOption) (CQRSClientTransports, error) {
	t := &cQRSGrpcClientTransports{}
	var err error
	t.createUser, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"pb.CQRS",
				"CreateUser",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				grpc.ClientBefore(grpcx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.findUser, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"pb.CQRS",
				"FindUser",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				GetUserResponse{},
				grpc.ClientBefore(grpcx.OutgoingMetadata),
			),
			options...,
		)
	})
	return t, err
}

type cQRSGrpcClient struct {
	endpoints CQRSEndpoints
}

func (c *cQRSGrpcClient) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/pb.CQRS/CreateUser")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.CreateUser(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *cQRSGrpcClient) FindUser(ctx context.Context, request *FindUserRequest) (*GetUserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/pb.CQRS/FindUser")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.FindUser(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*GetUserResponse), nil
}

func NewCQRSGrpcClient(transports CQRSClientTransports, middlewares ...endpoint.Middleware) CQRSService {
	endpoints := newCQRSClientEndpoints(transports, middlewares...)
	return &cQRSGrpcClient{endpoints: endpoints}
}

// =========================== http server ===========================

type CQRSHttpServerTransports interface {
	CreateUser() *http.Server
	FindUser() *http.Server
}

type cQRSHttpServerTransports struct {
	createUser *http.Server
	findUser   *http.Server
}

func (t *cQRSHttpServerTransports) CreateUser() *http.Server {
	return t.createUser
}

func (t *cQRSHttpServerTransports) FindUser() *http.Server {
	return t.findUser
}

func newCQRSHttpServerTransports(endpoints CQRSEndpoints) CQRSHttpServerTransports {
	return &cQRSHttpServerTransports{
		createUser: http.NewServer(
			endpoints.CreateUser(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &CreateUserRequest{}
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
			http.ServerBefore(httpx.EndpointInjector("/pb.CQRS/CreateUser")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		findUser: http.NewServer(
			endpoints.FindUser(context.TODO()),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &FindUserRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
					return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*GetUserResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return statusx.ErrInternal.With(statusx.Wrap(err))
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/pb.CQRS/FindUser")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
	}
}

func NewCQRSHttpServerHandler(svc CQRSService, middlewares ...endpoint.Middleware) http1.Handler {
	endpoints := newCQRSServerEndpoints(svc, middlewares...)
	transports := newCQRSHttpServerTransports(endpoints)
	router := mux.NewRouter()
	router.NewRoute().Name("/pb.CQRS/CreateUser").Methods("POST").Path("/pb.CQRS/CreateUser").Handler(transports.CreateUser())
	router.NewRoute().Name("/pb.CQRS/FindUser").Methods("POST").Path("/pb.CQRS/FindUser").Handler(transports.FindUser())
	return router
}

// =========================== http client ===========================

type cQRSHttpClientTransports struct {
	createUser transportx.ClientTransport
	findUser   transportx.ClientTransport
}

func (t *cQRSHttpClientTransports) CreateUser() transportx.ClientTransport {
	return t.createUser
}

func (t *cQRSHttpClientTransports) FindUser() transportx.ClientTransport {
	return t.findUser
}

func NewCQRSHttpClientTransports(target string, options ...transportx.ClientTransportOption) (CQRSClientTransports, error) {
	router := mux.NewRouter()
	router.NewRoute().Name("/pb.CQRS/CreateUser").Methods("POST").Path("/pb.CQRS/CreateUser")
	router.NewRoute().Name("/pb.CQRS/FindUser").Methods("POST").Path("/pb.CQRS/FindUser")
	t := &cQRSHttpClientTransports{}
	var err error
	t.createUser, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				func(scheme string, instance string) http.CreateRequestFunc {
					return func(ctx context.Context, obj any) (*http1.Request, error) {
						if obj == nil {
							return nil, errors.New("request object is nil")
						}
						req, ok := obj.(*CreateUserRequest)
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
						path, err := router.Get("/pb.CQRS/CreateUser").URLPath(pairs...)
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
				http.ClientBefore(httpx.OutgoingMetadata),
			),
			options...,
		)
	})
	t.findUser, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				func(scheme string, instance string) http.CreateRequestFunc {
					return func(ctx context.Context, obj any) (*http1.Request, error) {
						if obj == nil {
							return nil, errors.New("request object is nil")
						}
						req, ok := obj.(*FindUserRequest)
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
						path, err := router.Get("/pb.CQRS/FindUser").URLPath(pairs...)
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
					resp := &GetUserResponse{}
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
	return t, err
}

type cQRSHttpClient struct {
	endpoints CQRSEndpoints
}

func (c *cQRSHttpClient) CreateUser(ctx context.Context, request *CreateUserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/pb.CQRS/CreateUser")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.CreateUser(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *cQRSHttpClient) FindUser(ctx context.Context, request *FindUserRequest) (*GetUserResponse, error) {
	ctx = endpointx.InjectName(ctx, "/pb.CQRS/FindUser")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.FindUser(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*GetUserResponse), nil
}

func NewCQRSHttpClient(transports CQRSClientTransports, middlewares ...endpoint.Middleware) CQRSService {
	endpoints := newCQRSClientEndpoints(transports, middlewares...)
	return &cQRSHttpClient{endpoints: endpoints}
}
