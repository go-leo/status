package status

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

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
