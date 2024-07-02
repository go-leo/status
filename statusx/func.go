package statusx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/gox/slicex"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"

	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewError(c codes.Code, msg string) *Error {
	return &Error{e: &interstatusx.Error{
		GrpcStatus: &rpcstatus.Status{Code: int32(GrpcCodeFromCode(c)), Message: msg},
		HttpStatus: &httpstatus.Status{Code: int32(HttpStatusFromCode(c))},
	}}
}

func NewErrorf(c codes.Code, format string, a ...any) *Error {
	return NewError(c, fmt.Sprintf(format, a...))
}

func FromError(err error) (*Error, bool) {
	var statusErr *Error
	if errors.As(err, &statusErr) {
		return statusErr, true
	}
	grpcStatus, ok := grpcstatus.FromError(err)
	if ok {
		return FromStatus(grpcStatus), true
	}
	return nil, false
}

func FromGrpcError(err error) *Error {
	grpcStatus, _ := grpcstatus.FromError(err)
	return FromStatus(grpcStatus)
}

// FromProto returns an error representing the given Status proto.
func FromProto(s *rpcstatus.Status) *Error {
	httpAny := errorx.Ignore(anypb.New(&httpstatus.Status{}))
	index := -1
	var httpProto *httpstatus.Status
	for i, detailAny := range s.Details {
		if detailAny.GetTypeUrl() != httpAny.GetTypeUrl() {
			continue
		}
		detail, err := detailAny.UnmarshalNew()
		if err != nil {
			continue
		}
		detailProto, ok := detail.(*httpstatus.Status)
		if !ok {
			continue
		}
		httpProto = detailProto
		index = i
		break
	}
	grpcProto := protox.Clone(s)
	if index != -1 {
		grpcProto.Details = slicex.Delete(grpcProto.Details, index)
	}
	return &Error{e: &interstatusx.Error{GrpcStatus: grpcProto, HttpStatus: httpProto}}
}

// FromStatus returns an error representing s.  If s.Code is OK, returns nil.
func FromStatus(s *grpcstatus.Status) *Error {
	return FromProto(s.Proto())
}

// FromContextError converts a context error to Error
func FromContextError(err error) *Error {
	if err == nil {
		return nil
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return NewError(codes.DeadlineExceeded, err.Error())
	}
	if errors.Is(err, context.Canceled) {
		return NewError(codes.Canceled, err.Error())
	}
	return NewError(codes.Unknown, err.Error())
}
