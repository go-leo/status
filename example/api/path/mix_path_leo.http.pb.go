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

func appendMixPathHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/leo.example.path.v1.MixPath/MixPath").Methods("GET").Path("/v1/{string}/{opt_string}/{wrap_string}/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	return router
}

// =========================== http server ===========================

func AppendMixPathHttpRoutes(router *mux.Router, svc MixPathService, middlewares ...endpoint.Middleware) *mux.Router {
	endpoints := newMixPathServerEndpoints(svc, middlewares...)
	router = appendMixPathHttpRoutes(router)
	router.Get("/leo.example.path.v1.MixPath/MixPath").Handler(_MixPath_MixPath_HttpServer_Transport(endpoints))
	return router
}

// =========================== http client ===========================

type mixPathHttpClientTransports struct {
	scheme        string
	router        *mux.Router
	clientOptions []http.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *mixPathHttpClientTransports) MixPath(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http.ClientOption{
		http.ClientBefore(httpx.OutgoingMetadataInjector),
		http.ClientBefore(httpx.OutgoingTimeLimiter),
		http.ClientBefore(httpx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := http.NewExplicitClient(
		_MixPath_MixPath_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_MixPath_MixPath_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newMixPathHttpClientTransports(scheme string, clientOptions []http.ClientOption, middlewares []endpoint.Middleware) MixPathClientTransportsV2 {
	return &mixPathHttpClientTransports{
		scheme:        scheme,
		router:        appendMixPathHttpRoutes(mux.NewRouter()),
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
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

// =========================== http transport ===========================

func _MixPath_MixPath_HttpServer_Transport(endpoints MixPathEndpoints) *http.Server {
	return http.NewServer(
		endpoints.MixPath(context.TODO()),
		_MixPath_MixPath_HttpServer_RequestDecoder,
		_MixPath_MixPath_HttpServer_ResponseEncoder,
		http.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.MixPath/MixPath")),
		http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
		http.ServerBefore(httpx.IncomingMetadataInjector),
		http.ServerBefore(httpx.IncomingTimeLimiter),
		http.ServerFinalizer(httpx.CancelInvoker),
		http.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func _MixPath_MixPath_HttpClient_Transport(target string, router *mux.Router, options ...httpx.ClientTransportOption) func() (transportx.ClientTransport, error) {
	return func() (transportx.ClientTransport, error) {
		return httpx.NewClientTransport(
			target,
			httpx.ClientFactory(
				_MixPath_MixPath_HttpClient_RequestEncoder(router),
				_MixPath_MixPath_HttpClient_ResponseDecoder,
				http.ClientBefore(httpx.OutgoingMetadataInjector),
				http.ClientBefore(httpx.OutgoingTimeLimiter),
				http.ClientBefore(httpx.OutgoingStain),
			),
			options...,
		)
	}
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
