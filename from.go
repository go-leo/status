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
		return FromRpcStatus(st), true
	case *grpcstatus.Status:
		return FromRpcStatus(st.Proto()), true
	case interface{ GRPCStatus() *grpcstatus.Status }:
		return FromRpcStatus(st.GRPCStatus().Proto()), true
	case *http.Response:
		return FromHttpResponse(st)
	case error:
		return FromError(st)
	default:
		return New(codes.Unknown, Message("%+v", obj)), false
	}
}

// FromGrpcStatus converts a gRPC status object to a local Status type.
// This is a convenience wrapper around FromRpcStatus that handles the proto conversion.
//
// Parameters:
//
//	grpcStatus: *grpcstatus.Status - The gRPC status object to convert
//
// Returns:
//
//	Status - The converted local status representation
func FromGrpcStatus(grpcStatus *grpcstatus.Status) Status {
	// Simply delegates to FromRpcStatus after extracting the proto message
	return FromRpcStatus(grpcStatus.Proto())
}

// FromRpcStatus converts a gRPC status proto (*rpcstatus.Status) to Status interface.
// It preserves:
// - Status code (converted to code.Code)
// - Message (if present)
// - All original details (via statuspb.FromDetails)
func FromRpcStatus(rpcStatus *rpcstatus.Status) Status {
	st := statuspb.FromGrpcDetails(rpcStatus.Details)
	st.RpcStatus = code.Code(rpcStatus.Code)
	if len(rpcStatus.GetMessage()) > 0 {
		st.Message = rpcStatus.GetMessage()
	}
	return &sampleStatus{st: st}
}

// FromHttpResponse converts an HTTP response to Status interface.
// It:
// 1. Attempts to parse response body as statuspb.Status (JSON)
// 2. Sets HTTP status code from response
// 3. Extracts headers marked with special key (kKey) and stores them in Details
// Returns the converted Status and true (always succeeds)
func FromHttpResponse(resp *http.Response) (Status, bool) {
	statusKeys, ok := resp.Header[kKey]
	if !ok {
		return nil, false
	}
	st := &statuspb.Status{}
	if data, err := io.ReadAll(resp.Body); err == nil {
		_ = protojson.Unmarshal(data, st)
	}
	st.HttpStatus = int32(resp.StatusCode)
	if keys := strings.Split(statusKeys[0], kSeparator); len(keys) > 0 {
		if st.Details == nil {
			st.Details = &statuspb.Details{}
		}
		if st.Details.Header == nil {
			st.Details.Header = &statuspb.Header{}
		}
		for _, key := range keys {
			for _, value := range resp.Header.Values(key) {
				st.Details.Header.Values = append(st.Details.Header.Values, &rpchttp.HttpHeader{Key: key, Value: value})
			}
		}
	}
	return &sampleStatus{st: st}, true
}

// FromError converts various error types to Status interface.
// Special cases handled:
// - context.DeadlineExceeded → codes.DeadlineExceeded
// - context.Canceled → codes.Canceled
// - url.Error → codes.Unavailable
// - sampleStatus errors (preserved as-is)
// - gRPC errors (via grpcstatus.FromError)
// Other errors become codes.Unknown status.
// Returns the Status and a boolean indicating if it was a recognized error type.
func FromError(err error) (Status, bool) {
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
		return FromRpcStatus(grpcStatus.Proto()), true
	}
	return New(codes.Unknown, Message(err.Error())), false
}
