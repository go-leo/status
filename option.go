package status

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-leo/status/internal/util"
	statuspb "github.com/go-leo/status/proto/leo/status"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Option func(st *sampleStatus)

func New(grpcStatus codes.Code, opts ...Option) Status {
	httpStatus := util.ToHttpStatusCode(grpcStatus)
	st := &sampleStatus{
		st: &statuspb.Status{
			Identifier: &statuspb.Identifier{Value: fmt.Sprintf("%d-%d", grpcStatus, httpStatus)},
			RpcStatus:  code.Code(grpcStatus),
			HttpStatus: &statuspb.HttpStatus{Value: int32(httpStatus)},
		},
	}
	for _, opt := range opts {
		opt(st)
	}
	return st
}

// Identifier sets the identifier of the Status.
// This distinguish between two Status objects as being the same when
// both code and status are identical.
func Identifier(id string) Option {
	return func(st *sampleStatus) {
		st.st.Identifier = &statuspb.Identifier{
			Value: id,
		}
	}
}

// Message sets the message of the Status.
func Message(format string, a ...any) Option {
	return func(st *sampleStatus) {
		st.st.Message = &statuspb.Message{Value: fmt.Sprintf(format, a...)}
	}
}

func HttpStatus(code int) Option {
	return func(st *sampleStatus) {
		st.st.HttpStatus = &statuspb.HttpStatus{Value: int32(code)}
	}
}

// Headers sets the http header info.
func Headers(header http.Header) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		if st.st.Details.Header == nil {
			st.st.Details.Header = &statuspb.Header{}
		}
		for key, values := range header {
			for _, value := range values {
				item := &httpstatus.HttpHeader{Key: key, Value: value}
				st.st.Details.Header.Values = append(st.st.Details.Header.Values, item)
			}
		}
	}
}

// ErrorInfo sets the error info.
func ErrorInfo(reason string, domain string, metadata map[string]string) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.ErrorInfo = &errdetails.ErrorInfo{
			Reason:   reason,
			Domain:   domain,
			Metadata: metadata,
		}
	}
}

// RetryInfo sets the retry info.
func RetryInfo(retryDelay time.Duration) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.RetryInfo = &errdetails.RetryInfo{
			RetryDelay: durationpb.New(retryDelay),
		}
	}
}

// DebugInfo sets the debug info.
func DebugInfo(stackEntries []string, detail string) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.DebugInfo = &errdetails.DebugInfo{
			StackEntries: stackEntries,
			Detail:       detail,
		}
	}
}

// QuotaFailure sets the quota failure info.
func QuotaFailure(violations []*errdetails.QuotaFailure_Violation) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.QuotaFailure = &errdetails.QuotaFailure{
			Violations: violations,
		}
	}
}

// PreconditionFailure sets the precondition failure info.
func PreconditionFailure(violations []*errdetails.PreconditionFailure_Violation) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.PreconditionFailure = &errdetails.PreconditionFailure{
			Violations: violations,
		}
	}
}

// BadRequest sets the bad request info.
func BadRequest(violations []*errdetails.BadRequest_FieldViolation) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.BadRequest = &errdetails.BadRequest{
			FieldViolations: violations,
		}
	}
}

// RequestInfo sets the request info.
func RequestInfo(requestId string, servingData string) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.RequestInfo = &errdetails.RequestInfo{
			RequestId:   requestId,
			ServingData: servingData,
		}
	}
}

// ResourceInfo sets the resource info.
func ResourceInfo(resourceType string, resourceName string, owner string, description string) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.ResourceInfo = &errdetails.ResourceInfo{
			ResourceType: resourceType,
			ResourceName: resourceName,
			Owner:        owner,
			Description:  description,
		}
	}
}

// Help sets the help info.
func Help(links []*errdetails.Help_Link) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.Help = &errdetails.Help{
			Links: links,
		}
	}
}

// LocalizedMessage sets the localized message info.
func LocalizedMessage(locale string, message string) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		st.st.Details.LocalizedMessage = &errdetails.LocalizedMessage{
			Locale:  locale,
			Message: message,
		}
	}
}

// Extra sets the extra info.
func Extra(extra proto.Message) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		value, err := anypb.New(extra)
		if err != nil {
			panic(err)
		}
		st.st.Details.Extra.Values = append(st.st.Details.Extra.Values, value)
	}
}
