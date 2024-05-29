// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package response

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http "google.golang.org/genproto/googleapis/rpc/http"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type ResponseService interface {
	OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error)
	HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error)
	HttpRequestStarBody(ctx context.Context, request *http.HttpRequest) (*http.HttpResponse, error)
}

type ResponseEndpoints interface {
	OmittedResponse() endpoint.Endpoint
	StarResponse() endpoint.Endpoint
	NamedResponse() endpoint.Endpoint
	HttpBodyResponse() endpoint.Endpoint
	HttpBodyNamedResponse() endpoint.Endpoint
	HttpRequestStarBody() endpoint.Endpoint
}

type responseEndpoints struct {
	svc         ResponseService
	middlewares []endpoint.Middleware
}

func (e *responseEndpoints) OmittedResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.OmittedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) StarResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.StarResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) NamedResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) HttpBodyResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) HttpBodyNamedResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyNamedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) HttpRequestStarBody() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpRequestStarBody(ctx, request.(*http.HttpRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewResponseEndpoints(svc ResponseService, middlewares ...endpoint.Middleware) ResponseEndpoints {
	return &responseEndpoints{svc: svc, middlewares: middlewares}
}
