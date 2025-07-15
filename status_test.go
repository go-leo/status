package status

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	statuspb "github.com/go-leo/status/proto/leo/status"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/grpc/codes"
)

func TestSampleStatus_Error(t *testing.T) {
	st := &sampleStatus{
		st: &statuspb.Status{
			RpcStatus:  code.Code(codes.NotFound),
			HttpStatus: 404,
			Message:    "Not Found",
		},
	}
	want := "status: rpc-status = NotFound, http-status = 404, desc = Not Found"
	if got := st.Error(); got != want {
		t.Errorf("Error() = %v, want %v", got, want)
	}
}

func TestSampleStatus_Identifier(t *testing.T) {
	st := &sampleStatus{
		st: &statuspb.Status{
			Identifier: "test-id",
		},
	}
	want := "test-id"
	if got := st.Identifier(); got != want {
		t.Errorf("Identifier() = %v, want %v", got, want)
	}
}

func TestSampleStatus_Code(t *testing.T) {
	st := &sampleStatus{
		st: &statuspb.Status{
			RpcStatus: code.Code(codes.Internal),
		},
	}
	want := codes.Internal
	if got := st.Code(); got != want {
		t.Errorf("Code() = %v, want %v", got, want)
	}
}

func TestSampleStatus_Message(t *testing.T) {
	st := &sampleStatus{
		st: &statuspb.Status{
			Message: "test message",
		},
	}
	want := "test message"
	if got := st.Message(); got != want {
		t.Errorf("Message() = %v, want %v", got, want)
	}
}

func TestSampleStatus_StatusCode(t *testing.T) {
	st := &sampleStatus{
		st: &statuspb.Status{
			HttpStatus: 500,
		},
	}
	want := 500
	if got := st.StatusCode(); got != want {
		t.Errorf("StatusCode() = %v, want %v", got, want)
	}
}

func TestSampleStatus_Headers(t *testing.T) {
	st := &sampleStatus{
		st: &statuspb.Status{
			Details: &statuspb.Details{
				Header: &statuspb.Header{
					Values: []*httpstatus.HttpHeader{
						{Key: "Content-Type", Value: "application/json"},
						{Key: "X-Request-Id", Value: "12345"},
					},
				},
			},
		},
	}

	headers := st.Headers()
	if got := headers.Get("Content-Type"); got != "application/json" {
		t.Errorf("Headers() Content-Type = %v, want %v", got, "application/json")
	}
	if got := headers.Get("X-Request-Id"); got != "12345" {
		t.Errorf("Headers() X-Request-Id = %v, want %v", got, "12345")
	}

	keysHeader := headers.Get(kKey)
	if keysHeader == "" {
		t.Error("Headers() missing keys header")
	}
	if !strings.Contains(keysHeader, "Content-Type") || !strings.Contains(keysHeader, "X-Request-Id") {
		t.Errorf("Headers() keys header = %v, should contain both Content-Type and X-Request-Id", keysHeader)
	}
}

func TestSampleStatus_MarshalJSON(t *testing.T) {
	st := &sampleStatus{
		st: &statuspb.Status{
			Identifier: "test-id",
			RpcStatus:  code.Code(codes.InvalidArgument),
			HttpStatus: 400,
			Message:    "参数错误",
		},
	}

	data, err := st.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal error = %v", err)
	}

	if id := result["identifier"]; id != "test-id" {
		t.Errorf("MarshalJSON() identifier = %v, want %v", id, "test-id")
	}
	if rpcStatus := result["rpcStatus"]; rpcStatus != "INVALID_ARGUMENT" { // codes.OK is 0
		t.Errorf("MarshalJSON() rpcStatus = %v, want %v", rpcStatus, codes.InvalidArgument.String() )
	}
	if msg := result["message"]; msg != "参数错误" {
		t.Errorf("MarshalJSON() message = %v, want %v", msg, "OK")
	}
}

func TestSampleStatus_ErrorInfo(t *testing.T) {
	expected := &errdetails.ErrorInfo{
		Reason:   "test-reason",
		Domain:   "test-domain",
		Metadata: map[string]string{"key": "value"},
	}
	st := &sampleStatus{
		st: &statuspb.Status{
			Details: &statuspb.Details{
				ErrorInfo: expected,
			},
		},
	}
	if got := st.ErrorInfo(); !reflect.DeepEqual(got, expected) {
		t.Errorf("ErrorInfo() = %v, want %v", got, expected)
	}
}

func TestSampleStatus_Is(t *testing.T) {
	tests := []struct {
		name   string
		target error
		want   bool
	}{
		{
			name: "same status",
			target: New(codes.NotFound,
				Identifier("test-id"),
				HttpStatus(404),
				Message("Not Found"),
			),
			want: true,
		},
		{
			name: "different status",
			target: New(codes.Internal,
				Identifier("test-id"),
				HttpStatus(500),
				Message("Internal Error"),
			),
			want: false,
		},
		{
			name:   "other error",
			target: errors.New("other error"),
			want:   false,
		},
	}

	st1 := New(codes.NotFound,
		Identifier("test-id"),
		HttpStatus(404),
		Message("Not Found"),
	).(*sampleStatus)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := st1.Is(tt.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		grpcCode   codes.Code
		options    []Option
		wantCode   codes.Code
		wantStatus int
		wantMsg    string
	}{
		{
			name:       "default OK",
			grpcCode:   codes.OK,
			wantCode:   codes.OK,
			wantStatus: 200,
		},
		{
			name:       "with message",
			grpcCode:   codes.NotFound,
			options:    []Option{Message("Resource %s not found", "user123")},
			wantCode:   codes.NotFound,
			wantStatus: 404,
			wantMsg:    "Resource user123 not found",
		},
		{
			name:       "with custom http status",
			grpcCode:   codes.Internal,
			options:    []Option{HttpStatus(599)},
			wantCode:   codes.Internal,
			wantStatus: 599,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := New(tt.grpcCode, tt.options...).(*sampleStatus)
			if got := st.Code(); got != tt.wantCode {
				t.Errorf("Code() = %v, want %v", got, tt.wantCode)
			}
			if got := st.StatusCode(); got != tt.wantStatus {
				t.Errorf("StatusCode() = %v, want %v", got, tt.wantStatus)
			}
			if tt.wantMsg != "" {
				if got := st.Message(); got != tt.wantMsg {
					t.Errorf("Message() = %v, want %v", got, tt.wantMsg)
				}
			}
		})
	}
}

func TestStatusDetails(t *testing.T) {
	t.Run("RetryInfo", func(t *testing.T) {
		st := New(codes.Unavailable, RetryInfo(5*time.Second)).(*sampleStatus)
		retryInfo := st.RetryInfo()
		if retryInfo == nil {
			t.Fatal("RetryInfo() returned nil")
		}
		if got := retryInfo.RetryDelay.AsDuration(); got != 5*time.Second {
			t.Errorf("RetryDelay = %v, want %v", got, 5*time.Second)
		}
	})

	t.Run("DebugInfo", func(t *testing.T) {
		st := New(codes.Internal,
			DebugInfo([]string{"stack1", "stack2"}, "test detail"),
		).(*sampleStatus)
		debugInfo := st.DebugInfo()
		if debugInfo == nil {
			t.Fatal("DebugInfo() returned nil")
		}
		if !reflect.DeepEqual(debugInfo.StackEntries, []string{"stack1", "stack2"}) {
			t.Errorf("StackEntries = %v, want %v", debugInfo.StackEntries, []string{"stack1", "stack2"})
		}
		if debugInfo.Detail != "test detail" {
			t.Errorf("Detail = %v, want %v", debugInfo.Detail, "test detail")
		}
	})

	t.Run("BadRequest", func(t *testing.T) {
		violations := []*errdetails.BadRequest_FieldViolation{
			{Field: "email", Description: "invalid format"},
		}
		st := New(codes.InvalidArgument, BadRequest(violations)).(*sampleStatus)
		badReq := st.BadRequest()
		if badReq == nil {
			t.Fatal("BadRequest() returned nil")
		}
		if !reflect.DeepEqual(badReq.FieldViolations, violations) {
			t.Errorf("FieldViolations = %v, want %v", badReq.FieldViolations, violations)
		}
	})

	t.Run("Extra", func(t *testing.T) {
		extra := &statuspb.Extra{}
		st := New(codes.OK, Extra(extra)).(*sampleStatus)
		if got, _ := st.Extra().GetValues()[0].UnmarshalNew(); !reflect.DeepEqual(got.(*statuspb.Extra).GetValues(), extra.GetValues()) {
			t.Errorf("Extra() = %v, want %v", got, extra)
		}
	})
}

func TestGRPCStatus(t *testing.T) {
	st := New(codes.NotFound,
		Message("Not Found"),
		ErrorInfo("RESOURCE_MISSING", "example.com", map[string]string{"resource": "user123"}),
	).(*sampleStatus)

	grpcStatus := st.GRPCStatus()
	if grpcStatus.Code() != codes.NotFound {
		t.Errorf("GRPCStatus().Code() = %v, want %v", grpcStatus.Code(), codes.NotFound)
	}
	if grpcStatus.Message() != "Not Found" {
		t.Errorf("GRPCStatus().Message() = %v, want %v", grpcStatus.Message(), "Not Found")
	}
	if len(grpcStatus.Details()) == 0 {
		t.Error("GRPCStatus().Details() is empty")
	}
}
