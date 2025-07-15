package status

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	statuspb "github.com/go-leo/status/proto/leo/status"
	"golang.org/x/exp/maps"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

// Status defines a unified interface for handling both gRPC and HTTP statuses
// It combines error handling with rich status information including:
// - Protocol-specific status codes (gRPC and HTTP)
// - Metadata and details
// - Header information
// - Various error details types
type Status interface {
	error

	// Identifier returns a unique string identifier for this status
	// Used to distinguish between different status instances with same codes
	Identifier() string

	// Code returns the gRPC status code
	Code() codes.Code

	// Message returns the human-readable status message
	Message() string

	// GRPCStatus converts the status to gRPC's *status.Status format
	// Implements the interface expected by gRPC error handling
	GRPCStatus() *grpcstatus.Status

	// Is checks if this status matches another error/status
	// Implements the future errors.Is functionality
	Is(target error) bool

	// StatusCode returns the HTTP status code
	StatusCode() int

	// Headers returns the HTTP headers associated with this status
	Headers() http.Header

	// Marshaler provides JSON serialization capability
	json.Marshaler

	// ErrorInfo returns extended error information if present
	ErrorInfo() *errdetails.ErrorInfo

	// RetryInfo returns retry timing information if present
	RetryInfo() *errdetails.RetryInfo

	// DebugInfo returns debugging information if present
	DebugInfo() *errdetails.DebugInfo

	// QuotaFailure returns quota violation details if present
	QuotaFailure() *errdetails.QuotaFailure

	// PreconditionFailure returns precondition violation details if present
	PreconditionFailure() *errdetails.PreconditionFailure

	// BadRequest returns field validation errors if present
	BadRequest() *errdetails.BadRequest

	// RequestInfo returns request metadata if present
	RequestInfo() *errdetails.RequestInfo

	// ResourceInfo returns resource details if present
	ResourceInfo() *errdetails.ResourceInfo

	// Help returns help/documentation links if present
	Help() *errdetails.Help

	// LocalizedMessage returns translated messages if present
	LocalizedMessage() *errdetails.LocalizedMessage

	// Extra returns custom protocol buffer extensions if present
	Extra() *statuspb.Extra
}

// sampleStatus is the concrete implementation of the Status interface
var _ Status = (*sampleStatus)(nil)

type sampleStatus struct {
	st *statuspb.Status
}

// Error returns a formatted string representation of the status
// Format: "status: rpc-status = CODE, http-status = CODE, desc = MESSAGE"
func (st *sampleStatus) Error() string {
	return fmt.Sprintf("status: rpc-status = %s, http-status = %d, desc = %s", st.Code(), st.StatusCode(), st.Message())
}

// Identifier returns the unique identifier string
func (st *sampleStatus) Identifier() string {
	return st.st.GetIdentifier()
}

// Code returns the gRPC status code
func (st *sampleStatus) Code() codes.Code {
	return codes.Code(st.st.GetRpcStatus())
}

// Message returns the status message
func (st *sampleStatus) Message() string {
	return st.st.GetMessage()
}

// GRPCStatus converts the status to gRPC's Status type
func (st *sampleStatus) GRPCStatus() *grpcstatus.Status {
	grpcStatus := &spb.Status{
		Code:    int32(st.Code()),
		Message: st.Message(),
		Details: st.st.GrpcDetails(),
	}
	return grpcstatus.FromProto(grpcStatus)
}

// Is implements status equality comparison
// Returns true if both statuses have matching:
// - gRPC codes
// - HTTP status codes
// - Identifiers
func (st *sampleStatus) Is(target error) bool {
	targetStatus, ok := From(target)
	if !ok {
		return false
	}
	return targetStatus.Code() == st.Code() &&
		targetStatus.StatusCode() == st.StatusCode() &&
		targetStatus.Identifier() == st.Identifier()
}

// StatusCode returns the HTTP status code
func (st *sampleStatus) StatusCode() int {
	return int(st.st.GetHttpStatus())
}

// Headers returns the HTTP headers including special status headers
// Automatically adds the kKey header containing all header keys
func (st *sampleStatus) Headers() http.Header {
	values := st.st.GetDetails().GetHeader().GetValues()
	header := make(http.Header, len(values))
	keys := make(map[string]struct{}, len(values))
	for _, item := range values {
		header.Add(item.GetKey(), item.GetValue())
		keys[item.GetKey()] = struct{}{}
	}
	header.Add(kKey, strings.Join(maps.Keys(keys), kSeparator))
	return header
}

// MarshalJSON provides JSON serialization of the status
// Omits http status and header details from serialization to avoid redundancy
func (st *sampleStatus) MarshalJSON() ([]byte, error) {
	httpBody := &statuspb.Status{
		Identifier: st.st.GetIdentifier(),
		RpcStatus:  st.st.GetRpcStatus(),
		Details:    st.st.GetDetails(),
		Message:    st.st.GetMessage(),
	}
	if st.st.GetDetails() != nil && st.st.GetDetails().GetHeader() != nil {
		httpBody.Details.Header = nil
	}
	return protojson.Marshal(httpBody)
}

// ErrorInfo returns the embedded ErrorInfo details if present
func (st *sampleStatus) ErrorInfo() *errdetails.ErrorInfo {
	return st.st.GetDetails().GetErrorInfo()
}

// RetryInfo returns the embedded RetryInfo details if present
func (st *sampleStatus) RetryInfo() *errdetails.RetryInfo {
	return st.st.GetDetails().GetRetryInfo()
}

// DebugInfo returns the embedded DebugInfo details if present
func (st *sampleStatus) DebugInfo() *errdetails.DebugInfo {
	return st.st.GetDetails().GetDebugInfo()
}

// QuotaFailure returns the embedded QuotaFailure details if present
func (st *sampleStatus) QuotaFailure() *errdetails.QuotaFailure {
	return st.st.GetDetails().GetQuotaFailure()
}

// PreconditionFailure returns the embedded PreconditionFailure details if present
func (st *sampleStatus) PreconditionFailure() *errdetails.PreconditionFailure {
	return st.st.GetDetails().GetPreconditionFailure()
}

// BadRequest returns the embedded BadRequest details if present
func (st *sampleStatus) BadRequest() *errdetails.BadRequest {
	return st.st.GetDetails().GetBadRequest()
}

// RequestInfo returns the embedded RequestInfo details if present
func (st *sampleStatus) RequestInfo() *errdetails.RequestInfo {
	return st.st.GetDetails().GetRequestInfo()
}

// ResourceInfo returns the embedded ResourceInfo details if present
func (st *sampleStatus) ResourceInfo() *errdetails.ResourceInfo {
	return st.st.GetDetails().GetResourceInfo()
}

// Help returns the embedded Help details if present
func (st *sampleStatus) Help() *errdetails.Help {
	return st.st.GetDetails().GetHelp()
}

// LocalizedMessage returns the embedded LocalizedMessage details if present
func (st *sampleStatus) LocalizedMessage() *errdetails.LocalizedMessage {
	return st.st.GetDetails().GetLocalizedMessage()
}

// Extra returns the embedded protocol buffer extensions if present
func (st *sampleStatus) Extra() *statuspb.Extra {
	return &statuspb.Extra{Values: st.st.GetDetails().GetExtra()}
}
