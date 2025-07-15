package status

import (
	"encoding/json"

	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

var (
	_Identifier          = &Identifier{}
	_HttpStatus          = &HttpStatus{}
	_ErrorInfo           = &errdetails.ErrorInfo{}
	_RetryInfo           = &errdetails.RetryInfo{}
	_DebugInfo           = &errdetails.DebugInfo{}
	_QuotaFailure        = &errdetails.QuotaFailure{}
	_PreconditionFailure = &errdetails.PreconditionFailure{}
	_BadRequest          = &errdetails.BadRequest{}
	_RequestInfo         = &errdetails.RequestInfo{}
	_ResourceInfo        = &errdetails.ResourceInfo{}
	_Help                = &errdetails.Help{}
	_LocalizedMessage    = &errdetails.LocalizedMessage{}
	_Header              = &Header{}
	_Extra               = &Extra{}
)

func (x *Status) GrpcDetails() []*anypb.Any {
	if x == nil || x.Details == nil {
		return nil
	}
	info := x.GetDetails().GetHeader()
	if info == nil {
		x.HttpDetails()
	}
	infoAny, err := anypb.New(info)
	if err != nil {
		panic(err)
	}
	return append(x.HttpDetails(), infoAny)
}

func (x *Status) HttpDetails() []*anypb.Any {
	if x == nil || x.Details == nil {
		return nil
	}
	var details []*anypb.Any
	if info := x.GetIdentifier(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	if info := x.GetHttpStatus(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
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
	if info := x.GetDetails().GetExtra(); info != nil {
		infoAny, err := anypb.New(info)
		if err != nil {
			panic(err)
		}
		details = append(details, infoAny)
	}
	return details
}

func FromDetails(details []*anypb.Any) *Status {
	st := &Status{}
	for _, value := range details {
		switch {
		case value.MessageIs(_Identifier):
			st.Identifier = &Identifier{}
			err := value.UnmarshalTo(st.Identifier)
			if err != nil {
				panic(err)
			}
		case value.MessageIs(_HttpStatus):
			st.HttpStatus = &HttpStatus{}
			err := value.UnmarshalTo(st.HttpStatus)
			if err != nil {
				panic(err)
			}
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
			if st.Details.Extra == nil {
				st.Details.Extra = &Extra{}
			}
			err := value.UnmarshalTo(st.Details.Header)
			if err != nil {
				panic(err)
			}
		default:
			if st.Details == nil {
				st.Details = &Details{}
			}
			if st.Details.Extra == nil {
				st.Details.Extra = &Extra{}
			}
			st.Details.Extra.Values = append(st.Details.Extra.Values, value)
		}
	}
	return st
}

func (x *Identifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.GetValue())
}

func (x *Identifier) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &x.Value)
}

func (x *HttpStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.GetValue())
}

func (x *HttpStatus) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &x.Value)
}

func (x *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.GetValue())
}

func (x *Message) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &x.Value)
}

func (x *Header) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.GetValues())
}

func (x *Header) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &x.Values)
}

func (x *Extra) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.GetValues())
}

func (x *Extra) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &x.Values)
}
