package status

import (
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

// Protobuf message type references for Any type handling
var (
	_Identifier          = &Identifier{}                     // Identifier message type reference
	_HttpStatus          = &HttpStatus{}                     // HttpStatus message type reference
	_ErrorInfo           = &errdetails.ErrorInfo{}           // ErrorInfo message type reference
	_RetryInfo           = &errdetails.RetryInfo{}           // RetryInfo message type reference
	_DebugInfo           = &errdetails.DebugInfo{}           // DebugInfo message type reference
	_QuotaFailure        = &errdetails.QuotaFailure{}        // QuotaFailure message type reference
	_PreconditionFailure = &errdetails.PreconditionFailure{} // PreconditionFailure message type reference
	_BadRequest          = &errdetails.BadRequest{}          // BadRequest message type reference
	_RequestInfo         = &errdetails.RequestInfo{}         // RequestInfo message type reference
	_ResourceInfo        = &errdetails.ResourceInfo{}        // ResourceInfo message type reference
	_Help                = &errdetails.Help{}                // Help message type reference
	_LocalizedMessage    = &errdetails.LocalizedMessage{}    // LocalizedMessage message type reference
	_Header              = &Header{}                         // Header message type reference
	_Extra               = &Extra{}                          // Extra message type reference
)

// GrpcDetails converts all status information to Any protobuf messages
// Includes:
//   - Identifier
//   - HttpStatus
//   - All error details (ErrorInfo, RetryInfo, DebugInfo, etc.)
//
// Returns:
//   - []*anypb.Any: Slice of Any messages containing all status information
//
// Panics:
//   - If Any message creation fails
func (x *Status) GrpcDetails() []*anypb.Any {
	var details []*anypb.Any
	// Convert each piece of status information to Any protobuf message
	if info := x.GetIdentifier(); info != "" {
		infoAny, err := anypb.New(&Identifier{Value: info})
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetHttpStatus(); info != 0 {
		infoAny, err := anypb.New(&HttpStatus{Value: info})
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	// Convert all error details to Any messages
	if info := x.GetDetails().GetErrorInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetRetryInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetDebugInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetQuotaFailure(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetPreconditionFailure(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetBadRequest(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetRequestInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetResourceInfo(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetHelp(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetLocalizedMessage(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetHeader(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetDetails().GetExtra(); len(info) > 0 {
		infoAny, err := anypb.New(&Extra{Values: info})
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	return details
}

// FromGrpcDetails converts a slice of Any messages back to a Status proto
// Handles all supported detail types including:
//   - Identifier
//   - HttpStatus
//   - All error details types (ErrorInfo, RetryInfo, etc.)
//   - Header information
//   - Extra custom messages
//
// Returns:
//   - *Status: Reconstructed Status proto
//
// Panics:
//   - If Any message unmarshaling fails
func FromGrpcDetails(details []*anypb.Any) *Status {
	st := &Status{}
	for _, value := range details {
		switch {
		// Handle each supported message type
		case value.MessageIs(_Identifier):
			identifier := &Identifier{}
			err := value.UnmarshalTo(identifier)
			if err != nil {
				panic(err)
			}
			st.Identifier = identifier.GetValue()
		case value.MessageIs(_HttpStatus):
			httpStatus := &HttpStatus{}
			err := value.UnmarshalTo(httpStatus)
			if err != nil {
				panic(err)
			}
			st.HttpStatus = httpStatus.GetValue()
		case value.MessageIs(_ErrorInfo):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.ErrorInfo = &errdetails.ErrorInfo{}
			err := value.UnmarshalTo(st.Details.ErrorInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_RetryInfo):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.RetryInfo = &errdetails.RetryInfo{}
			err := value.UnmarshalTo(st.Details.RetryInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_DebugInfo):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.DebugInfo = &errdetails.DebugInfo{}
			err := value.UnmarshalTo(st.Details.DebugInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_QuotaFailure):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.QuotaFailure = &errdetails.QuotaFailure{}
			err := value.UnmarshalTo(st.Details.QuotaFailure)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_PreconditionFailure):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.PreconditionFailure = &errdetails.PreconditionFailure{}
			err := value.UnmarshalTo(st.Details.PreconditionFailure)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_BadRequest):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.BadRequest = &errdetails.BadRequest{}
			err := value.UnmarshalTo(st.Details.BadRequest)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_RequestInfo):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.RequestInfo = &errdetails.RequestInfo{}
			err := value.UnmarshalTo(st.Details.RequestInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_ResourceInfo):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.ResourceInfo = &errdetails.ResourceInfo{}
			err := value.UnmarshalTo(st.Details.ResourceInfo)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_Help):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.Help = &errdetails.Help{}
			err := value.UnmarshalTo(st.Details.Help)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_LocalizedMessage):
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.LocalizedMessage = &errdetails.LocalizedMessage{}
			err := value.UnmarshalTo(st.Details.LocalizedMessage)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_Header):
			if st.Details == nil {
				st.Details = &Details{}
			}
			if st.Details.Header == nil {
				st.Details.Header = &Header{}
			}
			err := value.UnmarshalTo(st.Details.Header)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_Extra):
			if st.Details == nil {
				st.Details = &Details{}
			}
			extra := &Extra{}
			err := value.UnmarshalTo(extra)
			if err != nil {
				panic(err)
			}
			st.Details.Extra = append(st.Details.Extra, extra.GetValues()...)
		default:
			// Handle unknown message types by storing in Extra
			if st.Details == nil {
				st.Details = &Details{}
			}
			st.Details.Extra = append(st.Details.Extra, value)
		}
	}
	return st
}
