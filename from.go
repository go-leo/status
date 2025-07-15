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

func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := statuspb.FromDetails(grpcProto.Details)
	st.RpcStatus = code.Code(grpcProto.Code)
	if len(grpcProto.GetMessage()) > 0 {
		st.Message = &statuspb.Message{Value: grpcProto.GetMessage()}
	}
	return &sampleStatus{st: st}
}

func fromHttpResponse(resp *http.Response) (Status, bool) {
	st := &statuspb.Status{}
	if data, err := io.ReadAll(resp.Body); err == nil {
		_ = protojson.Unmarshal(data, st)
	}
	st.HttpStatus = &statuspb.HttpStatus{Value: int32(resp.StatusCode)}
	if keys := strings.Split(resp.Header.Get(kKey), kSeparator); len(keys) > 0 {
		st.Details = &statuspb.Details{
			Header: &statuspb.Header{},
		}
		for _, key := range keys {
			values := resp.Header.Values(key)
			for _, value := range values {
				st.Details.Header.Values = append(st.Details.Header.Values, &rpchttp.HttpHeader{Key: key, Value: value})
			}
		}
	}
	return &sampleStatus{st: st}, true
}

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
