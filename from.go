package status

import (
	"context"
	"errors"
	internalstatus "github.com/go-leo/status/internal/status"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"net/url"
)

func From(obj any) Status {
	switch st := obj.(type) {
	case *sampleStatus:
		return st
	case *internalstatus.Error:
		return &sampleStatus{
			err: st,
		}
	case Status:
		return st
	case *rpcstatus.Status:
		return fromRpcStatus(st)
	case *grpcstatus.Status:
		return fromRpcStatus(st.Proto())
	case interface{ GRPCStatus() *grpcstatus.Status }:
		return fromRpcStatus(st.GRPCStatus().Proto())
	case error:
		return fromError(st)
	default:
		return Unknown(Message("%s", obj))
	}
}

func fromError(err error) Status {
	if errors.Is(err, context.DeadlineExceeded) {
		return DeadlineExceeded()
	}
	if errors.Is(err, context.Canceled) {
		return Canceled()
	}
	if urlErr := new(url.Error); errors.As(err, &urlErr) {
		return Unavailable()
	}
	if statusErr := new(sampleStatus); errors.As(err, &statusErr) {
		return statusErr
	}
	if grpcStatus, ok := grpcstatus.FromError(err); ok {
		return fromRpcStatus(grpcStatus.Proto())
	}
	return Unknown(Message(err.Error()))
}

func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := newStatus(codes.Code(grpcProto.Code))
	st.err.GrpcStatus.Message = grpcProto.GetMessage()
	st.err.DetailInfo, st.err.HttpStatus = st.err.FromGrpcDetails(grpcProto.GetDetails())
	return st
}
