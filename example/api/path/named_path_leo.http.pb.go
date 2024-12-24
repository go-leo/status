// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package path

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	urlx "github.com/go-leo/gox/netx/urlx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
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

// =========================== http router ===========================

func appendNamedPathHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/NamedPathString").Methods("GET").Path("/v1/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/NamedPathOptString").Methods("GET").Path("/v1/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/NamedPathWrapString").Methods("GET").Path("/v1/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/EmbedNamedPathString").Methods("GET").Path("/v1/embed/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").Methods("GET").Path("/v1/embed/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").Methods("GET").Path("/v1/embed/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	return router
}

// =========================== http server ===========================

func AppendNamedPathHttpRoutes(router *mux.Router, svc NamedPathService, middlewares ...endpoint.Middleware) *mux.Router {
	endpoints := newNamedPathServerEndpoints(svc, middlewares...)
	router = appendNamedPathHttpRoutes(router)
	router.Get("/leo.example.path.v1.NamedPath/NamedPathString").Handler(_NamedPath_NamedPathString_HttpServer_Transport(endpoints))
	router.Get("/leo.example.path.v1.NamedPath/NamedPathOptString").Handler(_NamedPath_NamedPathOptString_HttpServer_Transport(endpoints))
	router.Get("/leo.example.path.v1.NamedPath/NamedPathWrapString").Handler(_NamedPath_NamedPathWrapString_HttpServer_Transport(endpoints))
	router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathString").Handler(_NamedPath_EmbedNamedPathString_HttpServer_Transport(endpoints))
	router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").Handler(_NamedPath_EmbedNamedPathOptString_HttpServer_Transport(endpoints))
	router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").Handler(_NamedPath_EmbedNamedPathWrapString_HttpServer_Transport(endpoints))
	return router
}

// =========================== http client ===========================

type namedPathHttpClientTransports struct {
	scheme        string
	router        *mux.Router
	clientOptions []http.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *namedPathHttpClientTransports) NamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http.ClientOption{
		http.ClientBefore(httpx.OutgoingMetadataInjector),
		http.ClientBefore(httpx.OutgoingTimeLimiter),
		http.ClientBefore(httpx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := http.NewExplicitClient(
		_NamedPath_NamedPathString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_NamedPathString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) NamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http.ClientOption{
		http.ClientBefore(httpx.OutgoingMetadataInjector),
		http.ClientBefore(httpx.OutgoingTimeLimiter),
		http.ClientBefore(httpx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := http.NewExplicitClient(
		_NamedPath_NamedPathOptString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_NamedPathOptString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) NamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http.ClientOption{
		http.ClientBefore(httpx.OutgoingMetadataInjector),
		http.ClientBefore(httpx.OutgoingTimeLimiter),
		http.ClientBefore(httpx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := http.NewExplicitClient(
		_NamedPath_NamedPathWrapString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_NamedPathWrapString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) EmbedNamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http.ClientOption{
		http.ClientBefore(httpx.OutgoingMetadataInjector),
		http.ClientBefore(httpx.OutgoingTimeLimiter),
		http.ClientBefore(httpx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := http.NewExplicitClient(
		_NamedPath_EmbedNamedPathString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_EmbedNamedPathString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) EmbedNamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http.ClientOption{
		http.ClientBefore(httpx.OutgoingMetadataInjector),
		http.ClientBefore(httpx.OutgoingTimeLimiter),
		http.ClientBefore(httpx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := http.NewExplicitClient(
		_NamedPath_EmbedNamedPathOptString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_EmbedNamedPathOptString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) EmbedNamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http.ClientOption{
		http.ClientBefore(httpx.OutgoingMetadataInjector),
		http.ClientBefore(httpx.OutgoingTimeLimiter),
		http.ClientBefore(httpx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := http.NewExplicitClient(
		_NamedPath_EmbedNamedPathWrapString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_EmbedNamedPathWrapString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newNamedPathHttpClientTransports(scheme string, clientOptions []http.ClientOption, middlewares []endpoint.Middleware) NamedPathClientTransports {
	return &namedPathHttpClientTransports{
		scheme:        scheme,
		router:        appendNamedPathHttpRoutes(mux.NewRouter()),
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

type namedPathHttpClient struct {
	balancers NamedPathBalancers
}

func (c *namedPathHttpClient) NamedPathString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.NamedPath/NamedPathString")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	balancer, err := c.balancers.NamedPathString(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHttpClient) NamedPathOptString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.NamedPath/NamedPathOptString")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	balancer, err := c.balancers.NamedPathOptString(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHttpClient) NamedPathWrapString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.NamedPath/NamedPathWrapString")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	balancer, err := c.balancers.NamedPathWrapString(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHttpClient) EmbedNamedPathString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.NamedPath/EmbedNamedPathString")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	balancer, err := c.balancers.EmbedNamedPathString(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHttpClient) EmbedNamedPathOptString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.NamedPath/EmbedNamedPathOptString")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	balancer, err := c.balancers.EmbedNamedPathOptString(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *namedPathHttpClient) EmbedNamedPathWrapString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	balancer, err := c.balancers.EmbedNamedPathWrapString(ctx)
	if err != nil {
		return nil, err
	}
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewNamedPathHttpClient(target string, opts ...httpx.ClientOption) NamedPathService {
	options := httpx.NewClientOptions(opts...)
	transports := newNamedPathHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	factories := newNamedPathFactories(transports)
	endpointers := newNamedPathEndpointers(target, options.InstancerFactory(), factories, options.Logger(), options.EndpointerOptions()...)
	balancers := newNamedPathBalancers(options.BalancerFactory(), endpointers)
	return &namedPathHttpClient{balancers: balancers}
}

// =========================== http transport ===========================

func _NamedPath_NamedPathString_HttpServer_Transport(endpoints NamedPathEndpoints) *http.Server {
	return http.NewServer(
		endpoints.NamedPathString(context.TODO()),
		_NamedPath_NamedPathString_HttpServer_RequestDecoder,
		_NamedPath_NamedPathString_HttpServer_ResponseEncoder,
		http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/NamedPathString")),
		http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
		http.ServerBefore(httpx.IncomingMetadataInjector),
		http.ServerBefore(httpx.IncomingTimeLimiter),
		http.ServerFinalizer(httpx.CancelInvoker),
		http.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func _NamedPath_NamedPathOptString_HttpServer_Transport(endpoints NamedPathEndpoints) *http.Server {
	return http.NewServer(
		endpoints.NamedPathOptString(context.TODO()),
		_NamedPath_NamedPathOptString_HttpServer_RequestDecoder,
		_NamedPath_NamedPathOptString_HttpServer_ResponseEncoder,
		http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/NamedPathOptString")),
		http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
		http.ServerBefore(httpx.IncomingMetadataInjector),
		http.ServerBefore(httpx.IncomingTimeLimiter),
		http.ServerFinalizer(httpx.CancelInvoker),
		http.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func _NamedPath_NamedPathWrapString_HttpServer_Transport(endpoints NamedPathEndpoints) *http.Server {
	return http.NewServer(
		endpoints.NamedPathWrapString(context.TODO()),
		_NamedPath_NamedPathWrapString_HttpServer_RequestDecoder,
		_NamedPath_NamedPathWrapString_HttpServer_ResponseEncoder,
		http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/NamedPathWrapString")),
		http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
		http.ServerBefore(httpx.IncomingMetadataInjector),
		http.ServerBefore(httpx.IncomingTimeLimiter),
		http.ServerFinalizer(httpx.CancelInvoker),
		http.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func _NamedPath_EmbedNamedPathString_HttpServer_Transport(endpoints NamedPathEndpoints) *http.Server {
	return http.NewServer(
		endpoints.EmbedNamedPathString(context.TODO()),
		_NamedPath_EmbedNamedPathString_HttpServer_RequestDecoder,
		_NamedPath_EmbedNamedPathString_HttpServer_ResponseEncoder,
		http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathString")),
		http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
		http.ServerBefore(httpx.IncomingMetadataInjector),
		http.ServerBefore(httpx.IncomingTimeLimiter),
		http.ServerFinalizer(httpx.CancelInvoker),
		http.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func _NamedPath_EmbedNamedPathOptString_HttpServer_Transport(endpoints NamedPathEndpoints) *http.Server {
	return http.NewServer(
		endpoints.EmbedNamedPathOptString(context.TODO()),
		_NamedPath_EmbedNamedPathOptString_HttpServer_RequestDecoder,
		_NamedPath_EmbedNamedPathOptString_HttpServer_ResponseEncoder,
		http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString")),
		http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
		http.ServerBefore(httpx.IncomingMetadataInjector),
		http.ServerBefore(httpx.IncomingTimeLimiter),
		http.ServerFinalizer(httpx.CancelInvoker),
		http.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func _NamedPath_EmbedNamedPathWrapString_HttpServer_Transport(endpoints NamedPathEndpoints) *http.Server {
	return http.NewServer(
		endpoints.EmbedNamedPathWrapString(context.TODO()),
		_NamedPath_EmbedNamedPathWrapString_HttpServer_RequestDecoder,
		_NamedPath_EmbedNamedPathWrapString_HttpServer_ResponseEncoder,
		http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString")),
		http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
		http.ServerBefore(httpx.IncomingMetadataInjector),
		http.ServerBefore(httpx.IncomingTimeLimiter),
		http.ServerFinalizer(httpx.CancelInvoker),
		http.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

// =========================== http coder ===========================

func _NamedPath_NamedPathString_HttpServer_RequestDecoder(ctx context.Context, r *http1.Request) (any, error) {
	req := &NamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.String_ = fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family"))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	queries := r.URL.Query()
	var queryErr error
	req.OptString = proto.String(queries.Get("opt_string"))
	req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
	if queryErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
	}
	return req, nil
}

func _NamedPath_NamedPathString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http.CreateRequestFunc {
	return func(scheme string, instance string) http.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http1.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*NamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetString_()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathString").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			queries["opt_string"] = append(queries["opt_string"], req.GetOptString())
			queries["wrap_string"] = append(queries["wrap_string"], req.GetWrapString().GetValue())
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

func _NamedPath_NamedPathString_HttpServer_ResponseEncoder(ctx context.Context, w http1.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http1.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_NamedPathString_HttpClient_ResponseDecoder(ctx context.Context, r *http1.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_NamedPathOptString_HttpServer_RequestDecoder(ctx context.Context, r *http1.Request) (any, error) {
	req := &NamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.OptString = proto.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	queries := r.URL.Query()
	var queryErr error
	req.String_ = queries.Get("string")
	req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
	if queryErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
	}
	return req, nil
}

func _NamedPath_NamedPathOptString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http.CreateRequestFunc {
	return func(scheme string, instance string) http.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http1.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*NamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetOptString()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathOptString").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			queries["string"] = append(queries["string"], req.GetString_())
			queries["wrap_string"] = append(queries["wrap_string"], req.GetWrapString().GetValue())
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

func _NamedPath_NamedPathOptString_HttpServer_ResponseEncoder(ctx context.Context, w http1.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http1.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_NamedPathOptString_HttpClient_ResponseDecoder(ctx context.Context, r *http1.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_NamedPathWrapString_HttpServer_RequestDecoder(ctx context.Context, r *http1.Request) (any, error) {
	req := &NamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.WrapString = wrapperspb.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	queries := r.URL.Query()
	var queryErr error
	req.String_ = queries.Get("string")
	req.OptString = proto.String(queries.Get("opt_string"))
	if queryErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
	}
	return req, nil
}

func _NamedPath_NamedPathWrapString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http.CreateRequestFunc {
	return func(scheme string, instance string) http.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http1.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*NamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetWrapString().GetValue()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathWrapString").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			queries["string"] = append(queries["string"], req.GetString_())
			queries["opt_string"] = append(queries["opt_string"], req.GetOptString())
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

func _NamedPath_NamedPathWrapString_HttpServer_ResponseEncoder(ctx context.Context, w http1.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http1.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_NamedPathWrapString_HttpClient_ResponseDecoder(ctx context.Context, r *http1.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_EmbedNamedPathString_HttpServer_RequestDecoder(ctx context.Context, r *http1.Request) (any, error) {
	req := &EmbedNamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Embed == nil {
		req.Embed = &NamedPathRequest{}
	}
	req.Embed.String_ = fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family"))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	return req, nil
}

func _NamedPath_EmbedNamedPathString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http.CreateRequestFunc {
	return func(scheme string, instance string) http.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http1.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*EmbedNamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetEmbed().GetString_()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathString").URLPath(pairs...)
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

func _NamedPath_EmbedNamedPathString_HttpServer_ResponseEncoder(ctx context.Context, w http1.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http1.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_EmbedNamedPathString_HttpClient_ResponseDecoder(ctx context.Context, r *http1.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_EmbedNamedPathOptString_HttpServer_RequestDecoder(ctx context.Context, r *http1.Request) (any, error) {
	req := &EmbedNamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Embed == nil {
		req.Embed = &NamedPathRequest{}
	}
	req.Embed.OptString = proto.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	return req, nil
}

func _NamedPath_EmbedNamedPathOptString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http.CreateRequestFunc {
	return func(scheme string, instance string) http.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http1.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*EmbedNamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetEmbed().GetOptString()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").URLPath(pairs...)
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

func _NamedPath_EmbedNamedPathOptString_HttpServer_ResponseEncoder(ctx context.Context, w http1.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http1.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_EmbedNamedPathOptString_HttpClient_ResponseDecoder(ctx context.Context, r *http1.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_EmbedNamedPathWrapString_HttpServer_RequestDecoder(ctx context.Context, r *http1.Request) (any, error) {
	req := &EmbedNamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Embed == nil {
		req.Embed = &NamedPathRequest{}
	}
	req.Embed.WrapString = wrapperspb.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	return req, nil
}

func _NamedPath_EmbedNamedPathWrapString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http.CreateRequestFunc {
	return func(scheme string, instance string) http.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http1.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*EmbedNamedPathRequest)
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
			path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").URLPath(pairs...)
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

func _NamedPath_EmbedNamedPathWrapString_HttpServer_ResponseEncoder(ctx context.Context, w http1.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http1.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_EmbedNamedPathWrapString_HttpClient_ResponseDecoder(ctx context.Context, r *http1.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
