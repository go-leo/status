// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package helloworld

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	urlx "github.com/go-leo/gox/netx/urlx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	io "io"
	http "net/http"
	url "net/url"
)

func appendGreeterHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/helloworld.Greeter/SayHello").Methods("GET").Path("/helloworld/{name}")
	return router
}
func AppendGreeterHttpServerRoutes(router *mux.Router, svc GreeterService, middlewares ...endpoint.Middleware) *mux.Router {
	transports := newGreeterHttpServerTransports(svc, middlewares...)
	router = appendGreeterHttpRoutes(router)
	router.Get("/helloworld.Greeter/SayHello").Handler(transports.SayHello())
	return router
}

func NewGreeterHttpClient(target string, opts ...httpx.ClientOption) GreeterService {
	options := httpx.NewClientOptions(opts...)
	transports := newGreeterHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newGreeterClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newGreeterClientService(endpoints, httpx.HttpClient)
}

type GreeterHttpServerTransports interface {
	SayHello() http.Handler
}

type GreeterHttpServerRequestDecoder interface {
	SayHello() http1.DecodeRequestFunc
}

type GreeterHttpServerResponseEncoder interface {
	SayHello() http1.EncodeResponseFunc
}

type GreeterHttpClientRequestEncoder interface {
	SayHello(instance string) http1.CreateRequestFunc
}

type GreeterHttpClientResponseDecoder interface {
	SayHello() http1.DecodeResponseFunc
}

type greeterHttpServerTransports struct {
	endpoints       GreeterServerEndpoints
	requestDecoder  GreeterHttpServerRequestDecoder
	responseEncoder GreeterHttpServerResponseEncoder
}

func (t *greeterHttpServerTransports) SayHello() http.Handler {
	return http1.NewServer(
		t.endpoints.SayHello(context.TODO()),
		t.requestDecoder.SayHello(),
		t.responseEncoder.SayHello(),
		http1.ServerBefore(httpx.EndpointInjector("/helloworld.Greeter/SayHello")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func newGreeterHttpServerTransports(svc GreeterService, middlewares ...endpoint.Middleware) GreeterHttpServerTransports {
	endpoints := newGreeterServerEndpoints(svc, middlewares...)
	return &greeterHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  greeterHttpServerRequestDecoder{},
		responseEncoder: greeterHttpServerResponseEncoder{},
	}
}

type greeterHttpServerRequestDecoder struct{}

func (greeterHttpServerRequestDecoder) SayHello() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &HelloRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.Name = vars.Get("name")
		if varErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
		}
		return req, nil
	}
}

type greeterHttpServerResponseEncoder struct{}

func (greeterHttpServerResponseEncoder) SayHello() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*HelloReply)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}

type greeterHttpClientTransports struct {
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  GreeterHttpClientRequestEncoder
	responseDecoder GreeterHttpClientResponseDecoder
}

func (t *greeterHttpClientTransports) SayHello(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.SayHello(instance),
		t.responseDecoder.SayHello(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newGreeterHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) GreeterClientTransports {
	return &greeterHttpClientTransports{
		clientOptions: clientOptions,
		middlewares:   middlewares,
		requestEncoder: greeterHttpClientRequestEncoder{
			scheme: scheme,
			router: appendGreeterHttpRoutes(mux.NewRouter()),
		},
		responseDecoder: greeterHttpClientResponseDecoder{},
	}
}

type greeterHttpClientRequestEncoder struct {
	router *mux.Router
	scheme string
}

func (e greeterHttpClientRequestEncoder) SayHello(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*HelloRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var pairs []string
		pairs = append(pairs, "name", req.GetName())
		path, err := e.router.Get("/helloworld.Greeter/SayHello").URLPath(pairs...)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		queries := url.Values{}
		target := &url.URL{
			Scheme:   e.scheme,
			Host:     instance,
			Path:     path.Path,
			RawQuery: queries.Encode(),
		}
		r, err := http.NewRequestWithContext(ctx, "GET", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return r, nil
	}
}

type greeterHttpClientResponseDecoder struct{}

func (greeterHttpClientResponseDecoder) SayHello() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &HelloReply{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
