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

type Status interface {
	error

	// Identifier returns the identifier.
	Identifier() string

	// Code returns the status code.
	Code() codes.Code

	// Message returns the message.
	Message() string

	// GRPCStatus returns the gRPC Status.
	// see: https://github.com/grpc/grpc-go/blame/8528f4387f276518050f2b71a9dee1e3fb19d924/status/status.go#L100
	// type grpcstatus interface{ GRPCStatus() *Status }
	GRPCStatus() *grpcstatus.Status

	// Is implements future errors.Is functionality.
	Is(target error) bool

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

	// Extra returns additional detail from the Status
	Extra() *statuspb.Extra
}

var _ Status = (*sampleStatus)(nil)

type sampleStatus struct {
	st *statuspb.Status
}

func (st *sampleStatus) Error() string {
	return fmt.Sprintf("status: rpc-status = %s, http-status = %d, desc = %s", st.Code(), st.StatusCode(), st.Message())
}

func (st *sampleStatus) Identifier() string {
	return st.st.GetIdentifier().GetValue()
}

func (st *sampleStatus) Code() codes.Code {
	return codes.Code(st.st.GetRpcStatus())
}

func (st *sampleStatus) Message() string {
	return st.st.GetMessage().GetValue()
}

func (st *sampleStatus) GRPCStatus() *grpcstatus.Status {
	grpcStatus := &spb.Status{
		Code:    int32(st.Code()),
		Message: st.Message(),
		Details: st.st.GrpcDetails(),
	}
	return grpcstatus.FromProto(grpcStatus)
}

func (st *sampleStatus) Is(target error) bool {
	targetStatus, ok := From(target)
	if !ok {
		return false
	}
	return targetStatus.Code() == st.Code() &&
		targetStatus.StatusCode() == st.StatusCode() &&
		targetStatus.Identifier() == st.Identifier()
}

func (st *sampleStatus) StatusCode() int {
	return int(st.st.GetHttpStatus().GetValue())
}

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

func (st *sampleStatus) MarshalJSON() ([]byte, error) {
	return protojson.Marshal(st.st)
}

func (st *sampleStatus) ErrorInfo() *errdetails.ErrorInfo {
	return st.st.GetDetails().GetErrorInfo()
}

func (st *sampleStatus) RetryInfo() *errdetails.RetryInfo {
	return st.st.GetDetails().GetRetryInfo()
}

func (st *sampleStatus) DebugInfo() *errdetails.DebugInfo {
	return st.st.GetDetails().GetDebugInfo()
}

func (st *sampleStatus) QuotaFailure() *errdetails.QuotaFailure {
	return st.st.GetDetails().GetQuotaFailure()
}

func (st *sampleStatus) PreconditionFailure() *errdetails.PreconditionFailure {
	return st.st.GetDetails().GetPreconditionFailure()
}

func (st *sampleStatus) BadRequest() *errdetails.BadRequest {
	return st.st.GetDetails().GetBadRequest()
}

func (st *sampleStatus) RequestInfo() *errdetails.RequestInfo {
	return st.st.GetDetails().GetRequestInfo()
}

func (st *sampleStatus) ResourceInfo() *errdetails.ResourceInfo {
	return st.st.GetDetails().GetResourceInfo()
}

func (st *sampleStatus) Help() *errdetails.Help {
	return st.st.GetDetails().GetHelp()
}

func (st *sampleStatus) LocalizedMessage() *errdetails.LocalizedMessage {
	return st.st.GetDetails().GetLocalizedMessage()
}

func (st *sampleStatus) Extra() *statuspb.Extra {
	return st.st.GetDetails().GetExtra()
}
