package status

import (
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
)

type Option func(st *sampleStatus)

// Wrap wraps the cause error into the Status.
func Wrap(err error) Option {
	return func(st *sampleStatus) {
		if err == nil {
			return
		}
		causeProto, ok := err.(proto.Message)
		if !ok {
			st.err.Cause = &Cause{Cause: &Cause_Message{Message: wrapperspb.String(fmt.Sprintf("%+v", err))}}
			return
		}
		causeAny, err := anypb.New(causeProto)
		if err != nil {
			panic(err)
		}
		st.err.Cause = &Cause{Cause: &Cause_Error{Error: causeAny}}
	}
}

func Message(format string, a ...any) Option {
	return func(st *sampleStatus) {
		if len(a) <= 0 {
			st.err.GrpcStatus.Message = format
			return
		}
		st.err.GrpcStatus.Message = fmt.Sprintf(format, a...)
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

// HttpHeader sets the http header info.
func HttpHeader(infos ...*httpstatus.HttpHeader) Option {
	return func(st *sampleStatus) {
		st.err.HttpStatus.Headers = append(st.err.HttpStatus.Headers, infos...)
	}
}

// HttpBody sets the http body.
func HttpBody(info *wrapperspb.BytesValue) Option {
	return func(st *sampleStatus) {
		st.err.HttpStatus.Body = info.GetValue()
	}
}

// ErrorInfo sets the error info.
func ErrorInfo(info *errdetails.ErrorInfo) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.ErrorInfo = info
	}
}

// RetryInfo sets the retry info.
func RetryInfo(info *errdetails.RetryInfo) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.RetryInfo = info
	}
}

// DebugInfo sets the debug info.
func DebugInfo(info *errdetails.DebugInfo) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.DebugInfo = info
	}
}

// QuotaFailure sets the quota failure info.
func QuotaFailure(info *errdetails.QuotaFailure) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.QuotaFailure = info
	}
}

// PreconditionFailure sets the precondition failure info.
func PreconditionFailure(info *errdetails.PreconditionFailure) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.PreconditionFailure = info
	}
}

// BadRequest sets the bad request info.
func BadRequest(info *errdetails.BadRequest) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.BadRequest = info
	}
}

// RequestInfo sets the request info.
func RequestInfo(info *errdetails.RequestInfo) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.RequestInfo = info
	}
}

// ResourceInfo sets the resource info.
func ResourceInfo(info *errdetails.ResourceInfo) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.ResourceInfo = info
	}
}

// Help sets the help info.
func Help(info *errdetails.Help) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.Help = info
	}
}

// LocalizedMessage sets the localized message info.
func LocalizedMessage(info *errdetails.LocalizedMessage) Option {
	return func(st *sampleStatus) {
		if st.err.Detail == nil {
			st.err.Detail = &Detail{}
		}
		st.err.Detail.LocalizedMessage = info
	}
}

// Details adds additional details to the Status as protocol buffer messages.
func Details(details ...proto.Message) Option {
	return func(st *sampleStatus) {
		for _, detail := range details {
			switch item := detail.(type) {
			case *Cause:
				st.err.Cause = item
			case *wrapperspb.StringValue:
				Message(item.GetValue())(st)
			case *httpstatus.HttpHeader:
				HttpHeader(item)(st)
			case *wrapperspb.BytesValue:
				HttpBody(item)(st)
			case *errdetails.ErrorInfo:
				ErrorInfo(item)(st)
			case *errdetails.RetryInfo:
				RetryInfo(item)(st)
			case *errdetails.DebugInfo:
				DebugInfo(item)(st)
			case *errdetails.QuotaFailure:
				QuotaFailure(item)(st)
			case *errdetails.PreconditionFailure:
				PreconditionFailure(item)(st)
			case *errdetails.BadRequest:
				BadRequest(item)(st)
			case *errdetails.RequestInfo:
				RequestInfo(item)(st)
			case *errdetails.ResourceInfo:
				ResourceInfo(item)(st)
			case *errdetails.Help:
				Help(item)(st)
			case *errdetails.LocalizedMessage:
				LocalizedMessage(item)(st)
			default:
				value, err := anypb.New(item)
				if err != nil {
					panic(err)
				}
				st.err.GrpcStatus.Details = append(st.err.GrpcStatus.Details, value)
			}
		}
	}
}
