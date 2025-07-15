package status

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	statuspb "github.com/go-leo/status/proto/leo/status"
	"google.golang.org/genproto/googleapis/rpc/code"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// From converts various types of status representations into a unified Status interface.
// It supports conversion from:
// - *sampleStatus (internal implementation)
// - Status interface
// - *statuspb.Status (protobuf)
// - *rpcstatus.Status (gRPC status)
// - *grpcstatus.Status (gRPC wrapper)
// - Any type implementing GRPCStatus() method
// - *http.Response (HTTP response)
// - error (including context errors, URL errors, and gRPC errors)
// Returns the converted Status and a boolean indicating if conversion was successful.
func From(obj any) (Status, bool) {
	switch st := obj.(type) {
	case *sampleStatus:
		return st, true
	case Status:
		return st, true
	case *statuspb.Status:
		return &sampleStatus{st: st}, true
	case *rpcstatus.Status:
		return fromRpcStatus(st), true
	case *grpcstatus.Status:
		return fromRpcStatus(st.Proto()), true
	case interface{ GRPCStatus() *grpcstatus.Status }:
		return fromRpcStatus(st.GRPCStatus().Proto()), true
	case *http.Response:
		return fromHttpResponse(st)
	case error:
		return fromError(st)
	default:
		return New(codes.Unknown, Message("%+v", obj)), false
	}
}

// fromRpcStatus converts a gRPC status proto (*rpcstatus.Status) to Status interface.
// It preserves:
// - Status code (converted to code.Code)
// - Message (if present)
// - All original details (via statuspb.FromDetails)
func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := statuspb.FromGrpcDetails(grpcProto.Details)
	st.RpcStatus = code.Code(grpcProto.Code)
	if len(grpcProto.GetMessage()) > 0 {
		st.Message = grpcProto.GetMessage()
	}
	return &sampleStatus{st: st}
}

// fromHttpResponse converts an HTTP response to Status interface.
// It:
// 1. Attempts to parse response body as statuspb.Status (JSON)
// 2. Sets HTTP status code from response
// 3. Extracts headers marked with special key (kKey) and stores them in Details
// Returns the converted Status and true (always succeeds)
func fromHttpResponse(resp *http.Response) (Status, bool) {
	st := &statuspb.Status{}
	if data, err := io.ReadAll(resp.Body); err == nil {
		_ = protojson.Unmarshal(data, st)
	}
	st.HttpStatus = int32(resp.StatusCode)
	if keys := strings.Split(resp.Header.Get(kKey), kSeparator); len(keys) > 0 {
		st.Details = &statuspb.Details{
			Header: &statuspb.Header{},
		}
		for _, key := range keys {
			for _, value := range resp.Header[key] {
				st.Details.Header.Values = append(st.Details.Header.Values, &rpchttp.HttpHeader{Key: key, Value: value})
			}
		}
	}
	return &sampleStatus{st: st}, true
}

// fromError converts various error types to Status interface.
// Special cases handled:
// - context.DeadlineExceeded → codes.DeadlineExceeded
// - context.Canceled → codes.Canceled
// - url.Error → codes.Unavailable
// - sampleStatus errors (preserved as-is)
// - gRPC errors (via grpcstatus.FromError)
// Other errors become codes.Unknown status.
// Returns the Status and a boolean indicating if it was a recognized error type.
func fromError(err error) (Status, bool) {
	if errors.Is(err, context.DeadlineExceeded) {
		return New(codes.DeadlineExceeded, Message(err.Error())), true
	}
	if errors.Is(err, context.Canceled) {
		return New(codes.Canceled, Message(err.Error())), true
	}
	if urlErr := new(url.Error); errors.As(err, &urlErr) {
		return New(codes.Unavailable, Message(urlErr.Error())), true
	}
	if statusErr := new(sampleStatus); errors.As(err, &statusErr) {
		return statusErr, true
	}
	grpcStatus, ok := grpcstatus.FromError(err)
	if ok {
		return fromRpcStatus(grpcStatus.Proto()), true
	}
	return New(codes.Unknown, Message(err.Error())), false
}
