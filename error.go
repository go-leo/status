package status

import (
	"errors"
	"fmt"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (err *Error) Message() string {
	if err == nil {
		return ""
	}
	if causeAny := err.GetCause(); causeAny != nil {
		if causeProto := causeAny.GetError(); causeProto != nil {
			causeErr, err := causeProto.UnmarshalNew()
			if err != nil {
				panic(err)
			}
			data, err := protojson.Marshal(causeErr)
			if err != nil {
				panic(err)
			}
			return string(data)
		}
		if causeMsg := causeAny.GetMessage(); causeMsg != nil {
			return causeMsg.GetValue()
		}
	}
	if errorInfo := err.GetDetail().GetErrorInfo(); errorInfo != nil {
		data, err := protojson.Marshal(errorInfo)
		if err != nil {
			panic(err)
		}
		return string(data)
	}
	return err.GetGrpcStatus().GetMessage()
}

func (err *Error) Details() []proto.Message {
	details := err.GetGrpcStatus().GetDetails()
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

func (err *Error) AppendHeader(headers []*httpstatus.HttpHeader) []*httpstatus.HttpHeader {
	// add cause info to header
	if err.GetCause() != nil {
		info, err := anypb.New(err.GetCause())
		if err != nil {
			panic(err)
		}
		data, err := protojson.Marshal(info)
		if err != nil {
			panic(err)
		}
		item := &httpstatus.HttpHeader{
			Key:   kStatusCauseKey,
			Value: string(data),
		}
		headers = append(headers, item)
	}

	// add detail info to header
	if err.GetDetail() != nil {
		info, err := anypb.New(err.GetDetail())
		if err != nil {
			panic(err)
		}
		data, err := protojson.Marshal(info)
		if err != nil {
			panic(err)
		}
		item := &httpstatus.HttpHeader{
			Key:   kStatusDetailKey,
			Value: string(data),
		}
		headers = append(headers, item)
	}

	// add grpc status info to header
	if err.GetGrpcStatus() != nil {
		info, err := anypb.New(err.GetGrpcStatus())
		if err != nil {
			panic(err)
		}
		data, err := protojson.Marshal(info)
		if err != nil {
			panic(err)
		}
		item := &httpstatus.HttpHeader{
			Key:   kStatusGrpcKey,
			Value: string(data),
		}
		headers = append(headers, item)
	}
	return headers
}

func (err *Error) AppendDetails(details []*anypb.Any) []*anypb.Any {
	// add cause info to details
	if err.GetCause() != nil {
		cause, err := anypb.New(err.GetCause())
		if err != nil {
			panic(err)
		}
		details = append(details, cause)
	}

	// add detail info to details
	if err.GetDetail() != nil {
		detail, err := anypb.New(err.GetDetail())
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}

	// add http status info to details
	if err.GetHttpStatus() != nil {
		httpStatus, err := anypb.New(err.GetHttpStatus())
		if err != nil {
			panic(err)
		}
		details = append(details, httpStatus)
	}
	return details
}

func (*Cause) Wrap(err error) *Cause {
	if err == nil {
		return nil
	}
	causeProto, ok := err.(proto.Message)
	if !ok {
		// if err is not proto.Message, return message
		return &Cause{
			Cause: &Cause_Message{
				Message: wrapperspb.String(fmt.Sprintf("%+v", err)),
			},
		}
	}
	causeAny, err := anypb.New(causeProto)
	if err != nil {
		panic(err)
	}
	return &Cause{
		Cause: &Cause_Error{
			Error: causeAny,
		},
	}
}

func (cause *Cause) Unwrap() error {
	// if no cause, return nil
	if cause == nil {
		return nil
	}
	causeAny := cause.GetError()
	if causeAny == nil {
		// if cause is not proto.Message, return message
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
