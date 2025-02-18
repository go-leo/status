package status

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-leo/gox/protox"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"net/http"
	"slices"
)

type Status interface {
	error

	// With wraps the current Status with the given options and return new Status.
	With(opts ...Option) Status

	// Code returns the status code.
	Code() codes.Code

	// Message returns the message.
	Message() string

	// Unwrap unwraps the cause error from the current Status.
	Unwrap() error

	// GRPCStatus returns the gRPC Status.
	// see: https://github.com/grpc/grpc-go/blame/8528f4387f276518050f2b71a9dee1e3fb19d924/status/status.go#L100
	// type grpcstatus interface{ GRPCStatus() *Status }
	GRPCStatus() *grpcstatus.Status

	// HTTPStatus returns the HTTP Status.
	HTTPStatus() *httpstatus.HttpResponse

	// Is implements future errors.Is functionality.
	Is(target error) bool

	// Equals checks if the current status is equal to the target status by
	// comparing gRPC status code and http status code.
	// It does not compare the details.
	Equals(target error) bool

	// StatusCode returns the http status code.
	StatusCode() int

	// Headers returns the http header info.
	Headers() http.Header

	// Marshaler implements json.Marshaler.
	json.Marshaler

	// ErrorInfo returns the error info.
	ErrorInfo() *errdetails.ErrorInfo

	// RetryInfo returns the retry info.
	RetryInfo() *errdetails.RetryInfo

	// DebugInfo returns the debug info.
	DebugInfo() *errdetails.DebugInfo

	// QuotaFailure returns the quota failure info.
	QuotaFailure() *errdetails.QuotaFailure

	// PreconditionFailure returns the precondition failure info.
	PreconditionFailure() *errdetails.PreconditionFailure

	// BadRequest returns the bad request info.
	BadRequest() *errdetails.BadRequest

	// RequestInfo returns the request info.
	RequestInfo() *errdetails.RequestInfo

	// ResourceInfo returns the resource info.
	ResourceInfo() *errdetails.ResourceInfo

	// Help returns the help info.
	Help() *errdetails.Help

	// LocalizedMessage returns the localized message info.
	LocalizedMessage() *errdetails.LocalizedMessage

	// Details returns additional details from the Status
	Details() []proto.Message
}

type sampleStatus struct {
	err *Error
}

func newStatus(code codes.Code) *sampleStatus {
	var statusCode int
	switch code {
	case codes.OK:
		statusCode = http.StatusOK
	case codes.Canceled:
		statusCode = 499
	case codes.Unknown:
		statusCode = http.StatusInternalServerError
	case codes.InvalidArgument:
		statusCode = http.StatusBadRequest
	case codes.DeadlineExceeded:
		statusCode = http.StatusGatewayTimeout
	case codes.NotFound:
		statusCode = http.StatusNotFound
	case codes.AlreadyExists:
		statusCode = http.StatusConflict
	case codes.PermissionDenied:
		statusCode = http.StatusForbidden
	case codes.ResourceExhausted:
		statusCode = http.StatusTooManyRequests
	case codes.FailedPrecondition:
		statusCode = http.StatusBadRequest
	case codes.Aborted:
		statusCode = http.StatusConflict
	case codes.OutOfRange:
		statusCode = http.StatusBadRequest
	case codes.Unimplemented:
		statusCode = http.StatusNotImplemented
	case codes.Internal:
		statusCode = http.StatusInternalServerError
	case codes.Unavailable:
		statusCode = http.StatusServiceUnavailable
	case codes.DataLoss:
		statusCode = http.StatusInternalServerError
	case codes.Unauthenticated:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}
	st := &sampleStatus{
		err: &Error{
			Cause:  nil,
			Detail: &Detail{},
			HttpStatus: &httpstatus.HttpResponse{
				Status: int32(statusCode),
			},
			GrpcStatus: &rpcstatus.Status{
				Code: int32(codes.ResourceExhausted),
			},
		},
	}
	return st
}

// Status wraps a pointer of a Status proto.
func (st *sampleStatus) Error() string {
	return fmt.Sprintf("status: code = %s, status-code = %d, desc = %s", st.Code(), st.StatusCode(), st.Message())
}

func (st *sampleStatus) With(opts ...Option) Status {
	clonedSt := &sampleStatus{
		err: protox.Clone(st.err),
	}
	for _, opt := range opts {
		opt(clonedSt)
	}
	return clonedSt
}

func (st *sampleStatus) Code() codes.Code {
	if st == nil || st.err == nil {
		return codes.OK
	}
	return codes.Code(st.err.GetGrpcStatus().GetCode())
}

func (st *sampleStatus) Message() string {
	if st == nil || st.err == nil {
		return ""
	}
	return st.err.Message()
}

func (st *sampleStatus) causeMessage(causeAny *Cause) string {
	if causeProto := causeAny.GetError(); causeProto != nil {
		causeErr, err := causeProto.UnmarshalNew()
		if err != nil {
			panic(err)
		}
		return causeErr.(error).Error()
	}
	if causeMsg := causeAny.GetMessage(); causeMsg != nil {
		return causeMsg.GetValue()
	}
	return ""
}

func (st *sampleStatus) GRPCStatus() *grpcstatus.Status {
	grpcStatus := protox.Clone(st.err.GetGrpcStatus())
	grpcStatus.Details = st.err.AppendDetails(grpcStatus.Details)
	return grpcstatus.FromProto(grpcStatus)
}

func (st *sampleStatus) HTTPStatus() *httpstatus.HttpResponse {
	httpStatus := protox.Clone(st.err.GetHttpStatus())
	httpStatus.Headers = st.err.AppendHeader(httpStatus.Headers)
	return httpStatus
}

func (st *sampleStatus) Is(target error) bool {
	var targetErr *sampleStatus
	if !errors.As(target, &targetErr) {
		return false
	}
	return proto.Equal(st.err, targetErr.err)
}

func (st *sampleStatus) Equals(target error) bool {
	targetStatus, ok := target.(Status)
	if !ok {
		return false
	}
	return targetStatus.Code() == st.Code() && targetStatus.StatusCode() == st.StatusCode()
}

func (st *sampleStatus) Unwrap() error {
	return st.err.GetCause().Unwrap()
}

func (st *sampleStatus) StatusCode() int {
	return int(st.err.GetHttpStatus().GetStatus())
}

func (st *sampleStatus) Headers() http.Header {
	header := make(http.Header)
	headers := st.err.AppendHeader(slices.Clone(st.err.GetHttpStatus().GetHeaders()))
	for _, item := range headers {
		header.Add(item.GetKey(), item.GetValue())
	}
	return header
}

func (st *sampleStatus) MarshalJSON() ([]byte, error) {
	return st.err.GetHttpStatus().GetBody(), nil
}

func (st *sampleStatus) ErrorInfo() *errdetails.ErrorInfo {
	return protox.Clone(st.err.GetDetail().GetErrorInfo())
}

func (st *sampleStatus) RetryInfo() *errdetails.RetryInfo {
	return protox.Clone(st.err.GetDetail().GetRetryInfo())
}

func (st *sampleStatus) DebugInfo() *errdetails.DebugInfo {
	return protox.Clone(st.err.GetDetail().GetDebugInfo())
}

func (st *sampleStatus) QuotaFailure() *errdetails.QuotaFailure {
	return protox.Clone(st.err.GetDetail().GetQuotaFailure())
}

func (st *sampleStatus) PreconditionFailure() *errdetails.PreconditionFailure {
	return protox.Clone(st.err.GetDetail().GetPreconditionFailure())
}

func (st *sampleStatus) BadRequest() *errdetails.BadRequest {
	return protox.Clone(st.err.GetDetail().GetBadRequest())
}

func (st *sampleStatus) RequestInfo() *errdetails.RequestInfo {
	return protox.Clone(st.err.GetDetail().GetRequestInfo())
}

func (st *sampleStatus) ResourceInfo() *errdetails.ResourceInfo {
	return protox.Clone(st.err.GetDetail().GetResourceInfo())
}

func (st *sampleStatus) Help() *errdetails.Help {
	return protox.Clone(st.err.GetDetail().GetHelp())
}

func (st *sampleStatus) LocalizedMessage() *errdetails.LocalizedMessage {
	return protox.Clone(st.err.GetDetail().GetLocalizedMessage())
}

func (st *sampleStatus) Details() []proto.Message {
	return st.err.Details()
}
