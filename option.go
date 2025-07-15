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

// Option defines a function type that modifies a sampleStatus instance
type Option func(st *sampleStatus)

// New creates a new Status instance with the given gRPC status code and options
// It automatically sets the corresponding HTTP status code using util.ToHttpStatusCode
// and applies all provided Option functions to configure the Status
func New(grpcStatus codes.Code, opts ...Option) Status {
	httpStatus := util.ToHttpStatusCode(grpcStatus)
	st := &sampleStatus{
		st: &statuspb.Status{
			Identifier: fmt.Sprintf("%d-%d", grpcStatus, httpStatus),
			RpcStatus:  code.Code(grpcStatus),
			HttpStatus: int32(httpStatus),
		},
	}
	for _, opt := range opts {
		opt(st)
	}
	return st
}

// Identifier sets a unique identifier for the Status
// This helps distinguish between Status objects with identical codes
func Identifier(id string) Option {
	return func(st *sampleStatus) {
		st.st.Identifier = id
	}
}

// Message sets the human-readable message for the Status
// Supports format strings with variadic arguments like fmt.Sprintf
func Message(format string, a ...any) Option {
	return func(st *sampleStatus) {
		st.st.Message = fmt.Sprintf(format, a...)
	}
}

// HttpStatus explicitly sets the HTTP status code
// Overrides the default mapping from gRPC status code
func HttpStatus(code int) Option {
	return func(st *sampleStatus) {
		st.st.HttpStatus = int32(code)
	}
}

// Headers configures HTTP headers to be included in the Status
// The headers will be stored in the Status details
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

// ErrorInfo attaches error metadata to the Status
// Includes reason, domain, and arbitrary key-value pairs
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

// RetryInfo specifies retry delay information
// Used to indicate when clients should retry failed operations
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

// DebugInfo attaches debugging information to the Status
// Can include stack traces and additional debug details
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

// QuotaFailure specifies quota violation details
// Used when requests exceed quota limits
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

// PreconditionFailure describes precondition violations
// Used when API preconditions aren't met
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

// BadRequest describes invalid request fields
// Contains field-level validation errors
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

// RequestInfo attaches request metadata
// Includes request ID and serving data for tracking
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

// ResourceInfo describes affected resources
// Provides details about resource type, name, owner, etc.
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

// Help provides assistance links
// Contains documentation and support resources
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

// LocalizedMessage provides translated error messages
// Includes locale information and localized text
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

// Extra attaches arbitrary protocol buffer messages
// Used for custom extension data, will panic if serialization fails
func Extra(extra proto.Message) Option {
	return func(st *sampleStatus) {
		if st.st.Details == nil {
			st.st.Details = &statuspb.Details{}
		}
		value, err := anypb.New(extra)
		if err != nil {
			panic(err)
		}
		st.st.Details.Extra = append(st.st.Details.Extra, value)
	}
}
