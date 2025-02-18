package status

import (
	"errors"
	"fmt"
	"github.com/go-leo/gox/protox"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
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
	GRPCStatus() *grpcstatus.Status

	// HTTPStatus returns the HTTP Status.
	HTTPStatus() *httpstatus.HttpResponse

	// Proto return the gRPC and HTTP status protocol buffers.
	Proto() (*rpcstatus.Status, *httpstatus.HttpResponse)

	// Is implements future errors.Is functionality.
	Is(target error) bool

	// Equals checks if the current status is equal to the target status by
	// comparing gRPC status code and http status code.
	// It does not compare the details.
	Equals(target error) bool

	// Headers gets the http header info.
	Headers() http.Header

	// HttpHeader gets the http header info.
	HttpHeader() []*httpstatus.HttpHeader

	// HttpBody gets the http body.
	HttpBody() *wrapperspb.BytesValue

	// ErrorInfo gets the error info.
	ErrorInfo() *errdetails.ErrorInfo

	// RetryInfo gets the retry info.
	RetryInfo() *errdetails.RetryInfo

	// DebugInfo gets the debug info.
	DebugInfo() *errdetails.DebugInfo

	// QuotaFailure gets the quota failure info.
	QuotaFailure() *errdetails.QuotaFailure

	// PreconditionFailure gets the precondition failure info.
	PreconditionFailure() *errdetails.PreconditionFailure

	// BadRequest gets the bad request info.
	BadRequest() *errdetails.BadRequest

	// RequestInfo gets the request info.
	RequestInfo() *errdetails.RequestInfo

	// ResourceInfo gets the resource info.
	ResourceInfo() *errdetails.ResourceInfo

	// Help gets the help info.
	Help() *errdetails.Help

	// LocalizedMessage gets the localized message info.
	LocalizedMessage() *errdetails.LocalizedMessage

	// Details return additional details from the Status
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
	grpcStatus := st.err.GetGrpcStatus()
	code := codes.Code(grpcStatus.GetCode())
	var message string
	if causeAny := st.err.GetCause(); causeAny != nil {
		message = st.causeMessage(causeAny)
	} else if errorInfo := st.err.GetDetail().GetErrorInfo(); errorInfo != nil {
		message = errorInfo.GetReason()
	} else {
		message = grpcStatus.GetMessage()
	}
	return fmt.Sprintf("statusx: code = %s, desc = %s", code, message)
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
	return st.err.GetGrpcStatus().GetMessage()
}

func (st *sampleStatus) causeMessage(causeAny *Cause) string {
	if causeProto := causeAny.GetError(); causeProto != nil {
		causeErr, _ := causeProto.UnmarshalNew()
		return causeErr.(error).Error()
	}
	if causeMsg := causeAny.GetMessage(); causeMsg != nil {
		return causeMsg.GetValue()
	}
	return ""
}

func (st *sampleStatus) GRPCStatus() *grpcstatus.Status {
	grpcStatus := st.err.GetGrpcStatus()
	// return new grpc statu
	grpcProto := &rpcstatus.Status{
		Code:    grpcStatus.GetCode(),
		Message: grpcStatus.GetMessage(),
	}

	// copy grpc status details
	details := make([]*anypb.Any, 0, len(grpcStatus.GetDetails())+3)

	// add cause info
	if st.err.GetCause() != nil {
		cause, _ := anypb.New(st.err.GetCause())
		details = append(details, cause)
	}

	// add detail
	if st.err.GetDetail() != nil {
		detail, _ := anypb.New(st.err.GetDetail())
		details = append(details, detail)
	}

	// add http status info
	if st.err.GetHttpStatus() != nil {
		httpStatus, _ := anypb.New(st.err.GetHttpStatus())
		details = append(details, httpStatus)
	}

	// add grpc status details
	details = append(details, grpcStatus.GetDetails()...)

	grpcProto.Details = details
	return grpcstatus.FromProto(grpcProto)
}

func (st *sampleStatus) HTTPStatus() *httpstatus.HttpResponse {
	return protox.Clone(st.err.GetHttpStatus())
}

func (st *sampleStatus) Proto() (*rpcstatus.Status, *httpstatus.HttpResponse) {
	return protox.Clone(st.err.GetGrpcStatus()), protox.Clone(st.err.GetHttpStatus())
}

func (st *sampleStatus) Is(target error) bool {
	var targetErr *sampleStatus
	if !errors.As(target, &targetErr) {
		return false
	}
	return proto.Equal(st.err, targetErr.err)
}

func (st *sampleStatus) Equals(target error) bool {
	var targetErr *sampleStatus
	if !errors.As(target, &targetErr) {
		return false
	}
	if st.err.GetGrpcStatus().GetCode() != targetErr.err.GetGrpcStatus().GetCode() {
		return false
	}
	if st.err.GetHttpStatus().GetStatus() != targetErr.err.GetHttpStatus().GetStatus() {
		return false
	}
	return true
}

func (st *sampleStatus) Unwrap() error {
	cause := st.err.GetCause()

	// if no cause, return nil
	if cause == nil {
		return nil
	}

	causeAny := cause.GetError()
	// if no cause error, return message
	if causeAny == nil {
		return errors.New(cause.GetMessage().GetValue())
	}

	// unmarshal cause error
	causeProto, err := causeAny.UnmarshalNew()
	if err != nil {
		panic(err)
	}
	// must be error
	return causeProto.(error)
}

func (st *sampleStatus) Headers() http.Header {
	header := make(http.Header)
	headers := st.err.GetHttpStatus().GetHeaders()
	for _, item := range headers {
		header.Add(item.GetKey(), item.GetValue())
	}
	return header
}

func (st *sampleStatus) HttpHeader() []*httpstatus.HttpHeader {
	return st.err.GetHttpStatus().GetHeaders()
}

func (st *sampleStatus) HttpBody() *wrapperspb.BytesValue {
	return wrapperspb.Bytes(st.err.GetHttpStatus().GetBody())
}

func (st *sampleStatus) ErrorInfo() *errdetails.ErrorInfo {
	return st.err.GetDetail().GetErrorInfo()
}

func (st *sampleStatus) RetryInfo() *errdetails.RetryInfo {
	return st.err.GetDetail().GetRetryInfo()
}

func (st *sampleStatus) DebugInfo() *errdetails.DebugInfo {
	return st.err.GetDetail().GetDebugInfo()
}

func (st *sampleStatus) QuotaFailure() *errdetails.QuotaFailure {
	return st.err.GetDetail().GetQuotaFailure()
}

func (st *sampleStatus) PreconditionFailure() *errdetails.PreconditionFailure {
	return st.err.GetDetail().GetPreconditionFailure()
}

func (st *sampleStatus) BadRequest() *errdetails.BadRequest {
	return st.err.GetDetail().GetBadRequest()
}

func (st *sampleStatus) RequestInfo() *errdetails.RequestInfo {
	return st.err.GetDetail().GetRequestInfo()
}

func (st *sampleStatus) ResourceInfo() *errdetails.ResourceInfo {
	return st.err.GetDetail().GetResourceInfo()
}

func (st *sampleStatus) Help() *errdetails.Help {
	return st.err.GetDetail().GetHelp()
}

func (st *sampleStatus) LocalizedMessage() *errdetails.LocalizedMessage {
	return st.err.GetDetail().GetLocalizedMessage()
}

func (st *sampleStatus) Details() []proto.Message {
	details := st.err.GetGrpcStatus().GetDetails()
	messages := make([]proto.Message, 0, len(details))
	for _, anyDetail := range details {
		detail, err := anyDetail.UnmarshalNew()
		if err != nil {
			panic(err)
		}
		messages = append(messages, detail)
	}
	return messages
}
