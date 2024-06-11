// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package path

import (
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	urlx "github.com/go-leo/gox/netx/urlx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	grpc1 "google.golang.org/grpc"
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
	MixPath() endpoint.Endpoint
}

type MixPathFactories interface {
	MixPath() sd.Factory
}

type mixPathEndpoints struct {
	svc         MixPathService
	middlewares []endpoint.Middleware
}

func (e *mixPathEndpoints) MixPath() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.MixPath(ctx, request.(*MixPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewMixPathEndpoints(svc MixPathService, middlewares ...endpoint.Middleware) MixPathEndpoints {
	return &mixPathEndpoints{svc: svc, middlewares: middlewares}
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

func NewMixPathGrpcServerTransports(endpoints MixPathEndpoints) MixPathGrpcServerTransports {
	return &mixPathGrpcServerTransports{
		mixPath: grpc.NewServer(
			endpoints.MixPath(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.MixPath/MixPath")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
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

func NewMixPathGrpcServer(transports MixPathGrpcServerTransports) MixPathService {
	return &mixPathGrpcServer{
		mixPath: transports.MixPath(),
	}
}

// =========================== grpc client ===========================

type MixPathGrpcClientTransports interface {
	MixPath() *grpc.Client
}

type mixPathGrpcClientTransports struct {
	mixPath *grpc.Client
}

func (t *mixPathGrpcClientTransports) MixPath() *grpc.Client {
	return t.mixPath
}

func NewMixPathGrpcClientTransports(conn *grpc1.ClientConn) MixPathGrpcClientTransports {
	return &mixPathGrpcClientTransports{
		mixPath: grpc.NewClient(
			conn,
			"leo.example.path.v1.MixPath",
			"MixPath",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
	}
}

type mixPathGrpcClientEndpoints struct {
	transports  MixPathGrpcClientTransports
	middlewares []endpoint.Middleware
}

func (e *mixPathGrpcClientEndpoints) MixPath() endpoint.Endpoint {
	return endpointx.Chain(e.transports.MixPath().Endpoint(), e.middlewares...)
}

func NewMixPathGrpcClientEndpoints(transports MixPathGrpcClientTransports, middlewares ...endpoint.Middleware) MixPathEndpoints {
	return &mixPathGrpcClientEndpoints{transports: transports, middlewares: middlewares}
}

type mixPathGrpcClient struct {
	mixPath endpoint.Endpoint
}

func (c *mixPathGrpcClient) MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.MixPath/MixPath")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.mixPath(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewMixPathGrpcClient(endpoints MixPathEndpoints) MixPathService {
	return &mixPathGrpcClient{
		mixPath: endpoints.MixPath(),
	}
}

type mixPathGrpcClientFactories struct {
	endpoints func(transports MixPathGrpcClientTransports) MixPathEndpoints
	opts      []grpc1.DialOption
}

func (f *mixPathGrpcClientFactories) MixPath() sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc1.NewClient(instance, f.opts...)
		if err != nil {
			return nil, nil, err
		}
		endpoints := f.endpoints(NewMixPathGrpcClientTransports(conn))
		return endpoints.MixPath(), conn, nil
	}
}

func NewMixPathGrpcClientFactories(endpoints func(transports MixPathGrpcClientTransports) MixPathEndpoints, opts ...grpc1.DialOption) MixPathFactories {
	return &mixPathGrpcClientFactories{endpoints: endpoints, opts: opts}
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

func NewMixPathHttpServerTransports(endpoints MixPathEndpoints) MixPathHttpServerTransports {
	return &mixPathHttpServerTransports{
		mixPath: http.NewServer(
			endpoints.MixPath(),
			func(ctx context.Context, r *http1.Request) (any, error) {
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
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.MixPath/MixPath")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
	}
}

func NewMixPathHttpServerHandler(endpoints MixPathHttpServerTransports) http1.Handler {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.path.v1.MixPath/MixPath").Methods("GET").Path("/v1/{string}/{opt_string}/{wrap_string}/classes/{class}/shelves/{shelf}/books/{book}/families/{family}").Handler(endpoints.MixPath())
	return router
}

// =========================== http client ===========================

type MixPathHttpClientTransports interface {
	MixPath() *http.Client
}

type mixPathHttpClientTransports struct {
	mixPath *http.Client
}

func (t *mixPathHttpClientTransports) MixPath() *http.Client {
	return t.mixPath
}

func NewMixPathHttpClientTransports(scheme string, instance string) MixPathHttpClientTransports {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.path.v1.MixPath/MixPath").Methods("GET").Path("/v1/{string}/{opt_string}/{wrap_string}/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	return &mixPathHttpClientTransports{
		mixPath: http.NewExplicitClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*MixPathRequest)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				namedPathParameter := req.GetEmbed().GetWrapString().GetValue()
				namedPathValues := strings.Split(namedPathParameter, "/")
				if len(namedPathValues) != 8 {
					return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
				}
				pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
				pairs = append(pairs, "string", req.GetString_(), "opt_string", req.GetOptString(), "wrap_string", req.GetWrapString().GetValue())
				path, err := router.Get("/leo.example.path.v1.MixPath/MixPath").URLPath(pairs...)
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
	}
}

type mixPathHttpClient struct {
	mixPath endpoint.Endpoint
}

func (c *mixPathHttpClient) MixPath(ctx context.Context, request *MixPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.MixPath/MixPath")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.mixPath(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewMixPathHttpClient(transports MixPathHttpClientTransports, middlewares ...endpoint.Middleware) MixPathService {
	return &mixPathHttpClient{
		mixPath: endpointx.Chain(transports.MixPath().Endpoint(), middlewares...),
	}
}
