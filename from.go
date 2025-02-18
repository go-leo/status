package status

import (
	"context"
	"errors"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"net/url"
)

func From(obj any) Status {
	switch st := obj.(type) {
	case *sampleStatus:
		return st
	case *Error:
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
	return Unknown(Wrap(err))
}

func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := newStatus(codes.Code(grpcProto.Code))
	for _, value := range grpcProto.GetDetails() {
		switch {
		case value.MessageIs(&Cause{}):
			st.err.Cause = new(Cause)
			err := value.UnmarshalTo(st.err.Cause)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&Detail{}):
			st.err.Detail = new(Detail)
			err := value.UnmarshalTo(st.err.Detail)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(&httpstatus.HttpResponse{}):
			err := value.UnmarshalTo(st.err.HttpStatus)
			if err != nil {
				panic(err)
			}
		default:
			st.err.GrpcStatus.Details = append(st.err.GrpcStatus.Details, value)
		}
	}
	return st
}
