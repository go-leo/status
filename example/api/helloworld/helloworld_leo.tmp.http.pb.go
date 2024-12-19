// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package helloworld

import (
	context "context"
	"github.com/aws/smithy-go/transport/http"
	endpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	urlx "github.com/go-leo/gox/netx/urlx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	io "io"
	http1 "net/http"
	url "net/url"
)

// =========================== http client ===========================

type greeterHttpTransport struct {
	scheme        string
	router        *mux.Router
	clientOptions []httptransport.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *greeterHttpTransport) SayHello(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []httptransport.ClientOption{
		httptransport.ClientBefore(httpx.OutgoingMetadataInjector),
		httptransport.ClientBefore(httpx.OutgoingTimeLimiter),
	}
	opts = append(opts, t.clientOptions...)
	client := httptransport.NewExplicitClient(
		_Greeter_SayHello_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_Greeter_SayHello_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newGreeterHttpTransport(scheme string, clientOptions []httptransport.ClientOption, middlewares []endpoint.Middleware) GreeterClientTransportsV2 {
	return &greeterHttpTransport{
		scheme:        scheme,
		router:        appendGreeterHttpRoutes(mux.NewRouter()),
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

type greeterHttpClient struct {
	target            string
	scheme            string
	router            *mux.Router
	clientOptions     []httptransport.ClientOption
	middlewares       []endpoint.Middleware
	EndpointerOptions []sd.EndpointerOption
	GreeterFactory    GreeterFactory
}

func (c *greeterHttpClient) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	ctx = endpointx.InjectName(ctx, "/helloworld.Greeter/SayHello")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	ctx = httpx.InjectTarget(ctx, c.target)
	var factory sd.Factory = func(instance string) (endpoint.Endpoint, io.Closer, error) {
		opts := []httptransport.ClientOption{
			httptransport.ClientBefore(httpx.OutgoingMetadataInjector),
			httptransport.ClientBefore(httpx.OutgoingTimeLimiter),
		}
		opts = append(opts, c.clientOptions...)
		client := httptransport.NewExplicitClient(
			_Greeter_SayHello_HttpClient_RequestEncoder(c.router)(c.scheme, c.target),
			_Greeter_SayHello_HttpClient_ResponseDecoder,
			opts...,
		)
		return client.Endpoint(), nil, nil
	}

	instancer := passthroughx.Instancer{}
	//instancer := consul.NewInstancer()
	endpointer := sd.NewEndpointer(instancer, c.GreeterFactory.SayHello(ctx), logx.FromContext(ctx), c.EndpointerOptions...)
	balancer := lb.NewRoundRobin(endpointer)
	endpoint, err := balancer.Endpoint()
	if err != nil {
		return nil, err
	}

	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*HelloReply), nil
}

func NewGreeterHttpClient(target string, middlewares []endpoint.Middleware, options ...httpx.ClientTransportOption) (GreeterService, error) {
	return &greeterHttpClient{
		target:         target,
		router:         appendGreeterHttpRoutes(mux.NewRouter()),
		middlewares:    middlewares,
		GreeterFactory: newGreeterFactories(newGreeterHttpTransport()),
	}, nil
}
