// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package path

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	urlx "github.com/go-leo/gox/netx/urlx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	io "io"
	http1 "net/http"
	url "net/url"
	strings "strings"
)

// =========================== endpoints ===========================

type MixPathService interface {
	MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error)
}

type MixPathEndpoints interface {
	MixPath(ctx context.Context) endpoint.Endpoint
}

type MixPathClientTransports interface {
	MixPath() transportx.ClientTransport
}

type MixPathFactories interface {
	MixPath(middlewares ...endpoint.Middleware) sd.Factory
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

// =========================== cqrs ===========================

// =========================== grpc server ===========================

type MixPathGrpcServerTransports interface {
	MixPath() *grpc.Server
}

type mixPathGrpcServerTransports struct {
	mixPath *grpc.Server
}

func (t *mixPathGrpcServerTransports) MixPath() *grpc.Server {
	return t.mixPath
}

func newMixPathGrpcServerTransports(endpoints MixPathEndpoints) MixPathGrpcServerTransports {
	return &mixPathGrpcServerTransports{
		mixPath: grpc.NewServer(
			endpoints.MixPath(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.MixPath/MixPath")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		),
	}
}

type mixPathGrpcServer struct {
	mixPath *grpc.Server
}

func (s *mixPathGrpcServer) MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.mixPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewMixPathGrpcServer(svc MixPathService, middlewares ...endpoint.Middleware) MixPathService {
	endpoints := newMixPathServerEndpoints(svc, middlewares...)
	transports := newMixPathGrpcServerTransports(endpoints)
	return &mixPathGrpcServer{
		mixPath: transports.MixPath(),
	}
}

// =========================== grpc client ===========================

type mixPathGrpcClientTransports struct {
	mixPath transportx.ClientTransport
}

func (t *mixPathGrpcClientTransports) MixPath() transportx.ClientTransport {
	return t.mixPath
}

func NewMixPathGrpcClientTransports(target string, options ...transportx.ClientTransportOption) (MixPathClientTransports, error) {
	t := &mixPathGrpcClientTransports{}
	var err error
	t.mixPath, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			grpcx.ClientFactory(
				"leo.example.path.v1.MixPath",
				"MixPath",
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

type mixPathGrpcClient struct {
	endpoints MixPathEndpoints
}

func (c *mixPathGrpcClient) MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.MixPath/MixPath")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.MixPath(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewMixPathGrpcClient(transports MixPathClientTransports, middlewares ...endpoint.Middleware) MixPathService {
	endpoints := newMixPathClientEndpoints(transports, middlewares...)
	return &mixPathGrpcClient{endpoints: endpoints}
}

// =========================== http server ===========================

type MixPathHttpServerTransports interface {
	MixPath() *http.Server
}

type mixPathHttpServerTransports struct {
	mixPath *http.Server
}

func (t *mixPathHttpServerTransports) MixPath() *http.Server {
	return t.mixPath
}

func newMixPathHttpServerTransports(endpoints MixPathEndpoints) MixPathHttpServerTransports {
	return &mixPathHttpServerTransports{
		mixPath: http.NewServer(
			endpoints.MixPath(context.TODO()),
			_MixPath_MixPath_HttpServer_RequestDecoder,
			_MixPath_MixPath_HttpServer_ResponseEncoder,
			http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.MixPath/MixPath")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadataInjector),
			http.ServerBefore(httpx.IncomingTimeLimiter),
			http.ServerFinalizer(httpx.CancelInvoker),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
	}
}

func AppendMixPathHttpRouter(router *mux.Router, svc MixPathService, middlewares ...endpoint.Middleware) *mux.Router {
	endpoints := newMixPathServerEndpoints(svc, middlewares...)
	transports := newMixPathHttpServerTransports(endpoints)
	router.NewRoute().Name("/leo.example.path.v1.MixPath/MixPath").Methods("GET").Path("/v1/{string}/{opt_string}/{wrap_string}/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").Handler(transports.MixPath())
	return router
}

// =========================== http client ===========================

type mixPathHttpClientTransports struct {
	mixPath transportx.ClientTransport
}

func (t *mixPathHttpClientTransports) MixPath() transportx.ClientTransport {
	return t.mixPath
}

func NewMixPathHttpClientTransports(target string, options ...transportx.ClientTransportOption) (MixPathClientTransports, error) {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.path.v1.MixPath/MixPath").Methods("GET").Path("/v1/{string}/{opt_string}/{wrap_string}/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	t := &mixPathHttpClientTransports{}
	var err error
	t.mixPath, err = errorx.Break[transportx.ClientTransport](err)(func() (transportx.ClientTransport, error) {
		return transportx.NewClientTransport(
			target,
			httpx.ClientFactory(
				_MixPath_MixPath_HttpClient_RequestEncoder(router),
				_MixPath_MixPath_HttpClient_ResponseDecoder,
				http.ClientBefore(httpx.OutgoingMetadataInjector),
				http.ClientBefore(httpx.OutgoingTimeLimiter),
			),
			options...,
		)
	})
	return t, err
}

type mixPathHttpClient struct {
	endpoints MixPathEndpoints
}

func (c *mixPathHttpClient) MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.MixPath/MixPath")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.MixPath(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewMixPathHttpClient(transports MixPathClientTransports, middlewares ...endpoint.Middleware) MixPathService {
	endpoints := newMixPathClientEndpoints(transports, middlewares...)
	return &mixPathHttpClient{endpoints: endpoints}
}

// =========================== http coder ===========================

func _MixPath_MixPath_HttpServer_RequestDecoder(ctx context.Context, r *http1.Request) (any, error) {
	req := &MixPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Embed == nil {
		req.Embed = &NamedPathRequest{}
	}
	req.Embed.WrapString = wrapperspb.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	req.String_ = vars.Get("string")
	req.OptString = proto.String(vars.Get("opt_string"))
	req.WrapString = wrapperspb.String(vars.Get("wrap_string"))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	return req, nil
}

func _MixPath_MixPath_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http.CreateRequestFunc {
	return func(scheme string, instance string) http.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http1.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*MixPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetEmbed().GetWrapString().GetValue()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			pairs = append(pairs, "string", req.GetString_(), "opt_string", req.GetOptString(), "wrap_string", req.GetWrapString().GetValue())
			path, err := router.Get("/leo.example.path.v1.MixPath/MixPath").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
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
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _MixPath_MixPath_HttpServer_ResponseEncoder(ctx context.Context, w http1.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http1.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _MixPath_MixPath_HttpClient_ResponseDecoder(ctx context.Context, r *http1.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
