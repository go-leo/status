package status

import (
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"net/http"
	"time"
)

type Option func(st *sampleStatus)

// Wrap wraps the cause error into the Status.
func Wrap(err error) Option {
	return func(st *sampleStatus) {
		st.err.Cause = st.err.Cause.Wrap(err)
	}
}

func Message(format string, a ...any) Option {
	return func(st *sampleStatus) {
		st.err.GrpcStatus.Message = fmt.Sprintf(format, a...)
	}
}

// Details adds additional details to the Status as protocol buffer messages.
func Details(details ...proto.Message) Option {
	return func(st *sampleStatus) {
		for _, item := range details {
			value, _ := anypb.New(item)
			st.err.GrpcStatus.Details = append(st.err.GrpcStatus.Details, value)
		}
	}
}

// Header sets the http header info.
func Header(header http.Header) Option {
	return func(st *sampleStatus) {
		for key, values := range header {
			for _, value := range values {
				item := &httpstatus.HttpHeader{Key: key, Value: value}
				st.err.HttpStatus.Headers = append(st.err.HttpStatus.Headers, item)
			}
		}
	}
}

// ErrorInfo sets the error info.
func ErrorInfo(reason string, domain string, metadata map[string]string) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.ErrorInfo = &errdetails.ErrorInfo{
			Reason:   reason,
			Domain:   domain,
			Metadata: metadata,
		}
	}
}

// RetryInfo sets the retry info.
func RetryInfo(retryDelay time.Duration) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.RetryInfo = &errdetails.RetryInfo{
			RetryDelay: durationpb.New(retryDelay),
		}
	}
}

// DebugInfo sets the debug info.
func DebugInfo(stackEntries []string, detail string) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.DebugInfo = &errdetails.DebugInfo{
			StackEntries: stackEntries,
			Detail:       detail,
		}
	}
}

// QuotaFailure sets the quota failure info.
func QuotaFailure(violations []*errdetails.QuotaFailure_Violation) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.QuotaFailure = &errdetails.QuotaFailure{
			Violations: violations,
		}
	}
}

// PreconditionFailure sets the precondition failure info.
func PreconditionFailure(violations []*errdetails.PreconditionFailure_Violation) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.PreconditionFailure = &errdetails.PreconditionFailure{
			Violations: violations,
		}
	}
}

// BadRequest sets the bad request info.
func BadRequest(violations []*errdetails.BadRequest_FieldViolation) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.BadRequest = &errdetails.BadRequest{
			FieldViolations: violations,
		}
	}
}

// RequestInfo sets the request info.
func RequestInfo(requestId string, servingData string) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.RequestInfo = &errdetails.RequestInfo{
			RequestId:   requestId,
			ServingData: servingData,
		}
	}
}

// ResourceInfo sets the resource info.
func ResourceInfo(resourceType string, resourceName string, owner string, description string) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.ResourceInfo = &errdetails.ResourceInfo{
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
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.Help = &errdetails.Help{
			Links: links,
		}
	}
}

// LocalizedMessage sets the localized message info.
func LocalizedMessage(locale string, message string) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.LocalizedMessage = &errdetails.LocalizedMessage{
			Locale:  locale,
			Message: message,
		}
	}
}
