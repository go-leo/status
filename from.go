package status

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	interstatusx "github.com/go-leo/status/internal/status"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"net/url"
	"strconv"
)

func From(obj any) Status {
	switch err := obj.(type) {
	case *sampleStatus:
		return err
	case Status:
		return err
	case *interstatusx.Error:
		return &sampleStatus{
			err: err,
		}
	case *rpcstatus.Status:
		return fromRpcStatus(err)
	case *grpcstatus.Status:
		return fromRpcStatus(err.Proto())
	case interface{ GRPCStatus() *grpcstatus.Status }:
		return From(err.GRPCStatus())
	case codes.Code:
		return fromGrpcCode(err)
	case int:
		return fromHttpCode(err)
	case error:
		return fromError(err)
	default:
		return ErrUnknown
	}
}

func FromGrpcError(err error) Status {
	if err == nil {
		return nil
	}
	if grpcStatus, ok := grpcstatus.FromError(err); ok {
		return fromRpcStatus(grpcStatus.Proto())
	}
	return ErrUnknown.With(Wrap(err))
}

func fromError(err error) Status {
	if errors.Is(err, context.DeadlineExceeded) {
		return ErrDeadlineExceeded.With(Wrap(err))
	}
	if errors.Is(err, context.Canceled) {
		return ErrCanceled.With(Wrap(err))
	}
	if urlErr := new(url.Error); errors.As(err, &urlErr) {
		return ErrUnavailable.With(Message(strconv.Quote(fmt.Sprintf("%s %s: %s", urlErr.Op, urlErr.URL, urlErr.Err))), Wrap(urlErr))
	}
	if statusErr := new(sampleStatus); errors.As(err, &statusErr) {
		return statusErr
	}
	if grpcStatus, ok := grpcstatus.FromError(err); ok {
		return From(grpcStatus)
	}
	return fromDefaultErrorEncoder(err)
}

func fromDefaultErrorEncoder(err error) Status {
	statusCode := http.StatusInternalServerError
	if sc, ok := err.(httptransport.StatusCoder); ok {
		statusCode = sc.StatusCode()
	}
	statusErr := fromHttpCode(statusCode)

	// body
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = "application/json; charset=utf-8", jsonBody
		}
	}

	// header
	header := make([]*httpstatus.HttpHeader, 0)
	header = append(header, &httpstatus.HttpHeader{
		Key:   "Content-Type",
		Value: contentType,
	})
	if headerer, ok := err.(httptransport.Headerer); ok {
		for key, values := range headerer.Headers() {
			for _, value := range values {
				header = append(header, &httpstatus.HttpHeader{
					Key:   key,
					Value: value,
				})
			}
		}
	}
	return statusErr.With(HttpHeader(header...), HttpBody(wrapperspb.Bytes(body)))
}

func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := &sampleStatus{
		err: &interstatusx.Error{
			GrpcStatus: &rpcstatus.Status{
				Code:    grpcProto.GetCode(),
				Message: grpcProto.GetMessage(),
			},
		},
	}
	for _, value := range grpcProto.GetDetails() {
		switch {
		case value.MessageIs(&interstatusx.Cause{}):
			st.err.Cause = new(interstatusx.Cause)
			_ = value.UnmarshalTo(st.err.Cause)
		case value.MessageIs(&interstatusx.Detail{}):
			st.err.Detail = new(interstatusx.Detail)
			_ = value.UnmarshalTo(st.err.Detail)
		case value.MessageIs(&httpstatus.HttpResponse{}):
			st.err.HttpStatus = new(httpstatus.HttpResponse)
			_ = value.UnmarshalTo(st.err.HttpStatus)
		default:
			st.err.GrpcStatus.Details = append(st.err.GrpcStatus.Details, value)
		}
	}
	return st
}

// fromGrpcCode converts a gRPC status code to Status.
func fromGrpcCode(code codes.Code) Status {
	statusErr, ok := kGrpcToHttpCode[code]
	if ok {
		return statusErr
	}
	return ErrUnknown
}

var kHttpToGrpcCode = map[int]Status{
	http.StatusBadRequest:         ErrInternal,
	http.StatusUnauthorized:       ErrUnauthenticated,
	http.StatusForbidden:          ErrPermissionDenied,
	http.StatusNotFound:           ErrUnimplemented,
	http.StatusTooManyRequests:    ErrUnavailable,
	http.StatusBadGateway:         ErrUnavailable,
	http.StatusServiceUnavailable: ErrUnavailable,
	http.StatusGatewayTimeout:     ErrUnavailable,
}

// fromHttpCode converts an HTTP status code to Status.
// See: [HTTP to gRPC Status Code Mapping]: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
func fromHttpCode(code int) Status {
	statusErr, ok := kHttpToGrpcCode[code]
	if ok {
		return statusErr
	}
	return ErrUnknown
}
