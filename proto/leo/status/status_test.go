package status

import (
	"testing"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	anypb "google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestGrpcDetails(t *testing.T) {
	// TC01: 包含所有字段的状态对象
	st := &Status{
		Identifier: "test-id",
		HttpStatus: 200,
		Details: &Details{
			ErrorInfo: &errdetails.ErrorInfo{Reason: "reason", Domain: "domain"},
			RetryInfo: &errdetails.RetryInfo{RetryDelay: nil},
		},
	}

	anyMessages := st.GrpcDetails()
	if len(anyMessages) != 4 { // Identifier and HttpStatus
		t.Errorf("Expected 2 Any messages, got %d", len(anyMessages))
	}

	// TC02: 空状态对象
	emptySt := &Status{}
	if anyMessages := emptySt.GrpcDetails(); anyMessages != nil {
		t.Errorf("Expected nil for empty status, got %v", anyMessages)
	}

	// TC03: 仅包含Identifier字段
	idOnly := &Status{
		Identifier: "test-id",
	}
	if anyMessages := idOnly.GrpcDetails(); len(anyMessages) != 1 {
		t.Errorf("Expected 1 Any message for Identifier, got %d", len(anyMessages))
	}
}

func TestFromDetails(t *testing.T) {
	// TC04: 包含所有类型Any消息
	identifierAny, _ := anypb.New(&Identifier{Value: "test-id"})
	httpStatusAny, _ := anypb.New(&HttpStatus{Value: 200})
	errorInfoAny, _ := anypb.New(&errdetails.ErrorInfo{Reason: "reason", Domain: "domain"})
	details := []*anypb.Any{identifierAny, httpStatusAny, errorInfoAny}
	result := FromGrpcDetails(details)

	if result.Identifier == "" || result.Identifier != "test-id" {
		t.Errorf("Identifier not correctly parsed")
	}

	if result.HttpStatus == 0 || result.HttpStatus != 200 {
		t.Errorf("HttpStatus not correctly parsed")
	}

	if result.Details == nil || result.Details.ErrorInfo == nil || result.Details.ErrorInfo.Reason != "reason" {
		t.Errorf("ErrorInfo not correctly parsed")
	}

	// TC05: 空Any消息切片
	if result := FromGrpcDetails(nil); result == nil {
		t.Errorf("Expected empty Status, got nil")
	}

	// TC06: 包含未知消息类型
	unknownAny, _ := anypb.New(&wrapperspb.StringValue{Value: "StringValue"})
	detailsWithUnknown := []*anypb.Any{unknownAny}
	result = FromGrpcDetails(detailsWithUnknown)
	if result.Details == nil || result.Details.Extra == nil || len(result.Details.Extra) != 1 {
		t.Errorf("Unknown message not correctly handled in Extra field")
	}
}
