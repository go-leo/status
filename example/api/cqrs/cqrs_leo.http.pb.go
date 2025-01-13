// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package cqrs

import (
	bytes "bytes"
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
)

func appendCQRSHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/pb.CQRS/CreateUser").Methods("POST").Path("/pb.CQRS/CreateUser")
	router.NewRoute().Name("/pb.CQRS/DeleteUser").Methods("POST").Path("/pb.CQRS/DeleteUser")
	router.NewRoute().Name("/pb.CQRS/UpdateUser").Methods("POST").Path("/pb.CQRS/UpdateUser")
	router.NewRoute().Name("/pb.CQRS/FindUser").Methods("POST").Path("/pb.CQRS/FindUser")
	return router
}
func AppendCQRSHttpServerRoutes(router *mux.Router, svc CQRSService, middlewares ...endpoint.Middleware) *mux.Router {
	transports := newCQRSHttpServerTransports(svc, middlewares...)
	router = appendCQRSHttpRoutes(router)
	router.Get("/pb.CQRS/CreateUser").Handler(transports.CreateUser())
	router.Get("/pb.CQRS/DeleteUser").Handler(transports.DeleteUser())
	router.Get("/pb.CQRS/UpdateUser").Handler(transports.UpdateUser())
	router.Get("/pb.CQRS/FindUser").Handler(transports.FindUser())
	return router
}

func NewCQRSHttpClient(target string, opts ...httpx.ClientOption) CQRSService {
	options := httpx.NewClientOptions(opts...)
	transports := newCQRSHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newCQRSClientEndpoints(target, transports, options.Builder(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newCQRSClientService(endpoints, httpx.HttpClient)
}

type CQRSHttpServerTransports interface {
	CreateUser() http.Handler
	DeleteUser() http.Handler
	UpdateUser() http.Handler
	FindUser() http.Handler
}

type CQRSHttpServerRequestDecoder interface {
	CreateUser() http1.DecodeRequestFunc
	DeleteUser() http1.DecodeRequestFunc
	UpdateUser() http1.DecodeRequestFunc
	FindUser() http1.DecodeRequestFunc
}

type CQRSHttpServerResponseEncoder interface {
	CreateUser() http1.EncodeResponseFunc
	DeleteUser() http1.EncodeResponseFunc
	UpdateUser() http1.EncodeResponseFunc
	FindUser() http1.EncodeResponseFunc
}

type CQRSHttpClientRequestEncoder interface {
	CreateUser(instance string) http1.CreateRequestFunc
	DeleteUser(instance string) http1.CreateRequestFunc
	UpdateUser(instance string) http1.CreateRequestFunc
	FindUser(instance string) http1.CreateRequestFunc
}

type CQRSHttpClientResponseDecoder interface {
	CreateUser() http1.DecodeResponseFunc
	DeleteUser() http1.DecodeResponseFunc
	UpdateUser() http1.DecodeResponseFunc
	FindUser() http1.DecodeResponseFunc
}

type cQRSHttpServerTransports struct {
	endpoints       CQRSServerEndpoints
	requestDecoder  CQRSHttpServerRequestDecoder
	responseEncoder CQRSHttpServerResponseEncoder
}

func (t *cQRSHttpServerTransports) CreateUser() http.Handler {
	return http1.NewServer(
		t.endpoints.CreateUser(context.TODO()),
		t.requestDecoder.CreateUser(),
		t.responseEncoder.CreateUser(),
		http1.ServerBefore(httpx.EndpointInjector("/pb.CQRS/CreateUser")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *cQRSHttpServerTransports) DeleteUser() http.Handler {
	return http1.NewServer(
		t.endpoints.DeleteUser(context.TODO()),
		t.requestDecoder.DeleteUser(),
		t.responseEncoder.DeleteUser(),
		http1.ServerBefore(httpx.EndpointInjector("/pb.CQRS/DeleteUser")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *cQRSHttpServerTransports) UpdateUser() http.Handler {
	return http1.NewServer(
		t.endpoints.UpdateUser(context.TODO()),
		t.requestDecoder.UpdateUser(),
		t.responseEncoder.UpdateUser(),
		http1.ServerBefore(httpx.EndpointInjector("/pb.CQRS/UpdateUser")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *cQRSHttpServerTransports) FindUser() http.Handler {
	return http1.NewServer(
		t.endpoints.FindUser(context.TODO()),
		t.requestDecoder.FindUser(),
		t.responseEncoder.FindUser(),
		http1.ServerBefore(httpx.EndpointInjector("/pb.CQRS/FindUser")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func newCQRSHttpServerTransports(svc CQRSService, middlewares ...endpoint.Middleware) CQRSHttpServerTransports {
	endpoints := newCQRSServerEndpoints(svc, middlewares...)
	return &cQRSHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  cQRSHttpServerRequestDecoder{},
		responseEncoder: cQRSHttpServerResponseEncoder{},
	}
}

type cQRSHttpServerRequestDecoder struct{}

func (cQRSHttpServerRequestDecoder) CreateUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &CreateUserRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}
func (cQRSHttpServerRequestDecoder) DeleteUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &DeleteUserRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}
func (cQRSHttpServerRequestDecoder) UpdateUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &UpdateUserRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}
func (cQRSHttpServerRequestDecoder) FindUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &FindUserRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}

type cQRSHttpServerResponseEncoder struct{}

func (cQRSHttpServerResponseEncoder) CreateUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*emptypb.Empty)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}
func (cQRSHttpServerResponseEncoder) DeleteUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*DeleteUserResponse)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}
func (cQRSHttpServerResponseEncoder) UpdateUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*UpdateUserResponse)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}
func (cQRSHttpServerResponseEncoder) FindUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*GetUserResponse)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}

type cQRSHttpClientTransports struct {
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  CQRSHttpClientRequestEncoder
	responseDecoder CQRSHttpClientResponseDecoder
}

func (t *cQRSHttpClientTransports) CreateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.CreateUser(instance),
		t.responseDecoder.CreateUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *cQRSHttpClientTransports) DeleteUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.DeleteUser(instance),
		t.responseDecoder.DeleteUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *cQRSHttpClientTransports) UpdateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.UpdateUser(instance),
		t.responseDecoder.UpdateUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *cQRSHttpClientTransports) FindUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.FindUser(instance),
		t.responseDecoder.FindUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newCQRSHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) CQRSClientTransports {
	return &cQRSHttpClientTransports{
		clientOptions: clientOptions,
		middlewares:   middlewares,
		requestEncoder: cQRSHttpClientRequestEncoder{
			scheme: scheme,
			router: appendCQRSHttpRoutes(mux.NewRouter()),
		},
		responseDecoder: cQRSHttpClientResponseDecoder{},
	}
}

type cQRSHttpClientRequestEncoder struct {
	router *mux.Router
	scheme string
}

func (e cQRSHttpClientRequestEncoder) CreateUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*CreateUserRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		path, err := e.router.Get("/pb.CQRS/CreateUser").URLPath(pairs...)
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
		r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}
func (e cQRSHttpClientRequestEncoder) DeleteUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*DeleteUserRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		path, err := e.router.Get("/pb.CQRS/DeleteUser").URLPath(pairs...)
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
		r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}
func (e cQRSHttpClientRequestEncoder) UpdateUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*UpdateUserRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		path, err := e.router.Get("/pb.CQRS/UpdateUser").URLPath(pairs...)
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
		r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}
func (e cQRSHttpClientRequestEncoder) FindUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*FindUserRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		path, err := e.router.Get("/pb.CQRS/FindUser").URLPath(pairs...)
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
		r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}

type cQRSHttpClientResponseDecoder struct{}

func (cQRSHttpClientResponseDecoder) CreateUser() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &emptypb.Empty{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (cQRSHttpClientResponseDecoder) DeleteUser() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &DeleteUserResponse{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (cQRSHttpClientResponseDecoder) UpdateUser() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &UpdateUserResponse{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (cQRSHttpClientResponseDecoder) FindUser() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &GetUserResponse{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
