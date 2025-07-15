package status

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	statuspb "github.com/go-leo/status/proto/leo/status"
	"google.golang.org/genproto/googleapis/rpc/code"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

func TestFrom(t *testing.T) {
	// Test case 1: *sampleStatus input
	t.Run("SampleStatus", func(t *testing.T) {
		input := &sampleStatus{st: &statuspb.Status{RpcStatus: code.Code_OK}}
		result, ok := From(input)
		if !ok {
			t.Error("Expected ok=true for *sampleStatus input")
		}
		if result != input {
			t.Error("Expected same *sampleStatus instance to be returned")
		}
	})

	// Test case 2: Status interface input
	t.Run("StatusInterface", func(t *testing.T) {
		input := &sampleStatus{st: &statuspb.Status{RpcStatus: code.Code_OK}}
		result, ok := From(Status(input))
		if !ok {
			t.Error("Expected ok=true for Status interface input")
		}
		if result != input {
			t.Error("Expected same Status instance to be returned")
		}
	})

	// Test case 3: *statuspb.Status input
	t.Run("ProtoStatus", func(t *testing.T) {
		input := &statuspb.Status{RpcStatus: code.Code_OK}
		result, ok := From(input)
		if !ok {
			t.Error("Expected ok=true for *statuspb.Status input")
		}
		if result.(*sampleStatus).st != input {
			t.Error("Expected new sampleStatus wrapping the input proto")
		}
	})

	// Test case 4: *rpcstatus.Status input
	t.Run("RpcStatus", func(t *testing.T) {
		input := &rpcstatus.Status{Code: int32(codes.OK)}
		result, ok := From(input)
		if !ok {
			t.Error("Expected ok=true for *rpcstatus.Status input")
		}
		if result.Code() != codes.OK {
			t.Errorf("Expected code=OK, got %v", result.Code())
		}
	})

	// Test case 5: *grpcstatus.Status input
	t.Run("GrpcStatus", func(t *testing.T) {
		input := grpcstatus.New(codes.OK, "ok")
		result, ok := From(input)
		if !ok {
			t.Error("Expected ok=true for *grpcstatus.Status input")
		}
		if result.Code() != codes.OK {
			t.Errorf("Expected code=OK, got %v", result.Code())
		}
	})

	// Test case 6: GRPCStatus() interface input
	t.Run("GRPCStatusInterface", func(t *testing.T) {
		type gs interface{ GRPCStatus() *grpcstatus.Status }
		err := grpcstatus.New(codes.Internal, "").Err()
		input := err.(gs)
		result, ok := From(input)
		if !ok {
			t.Error("Expected ok=true for GRPCStatus() interface input")
		}
		if result.Code() != codes.Internal {
			t.Errorf("Expected code=OK, got %v", result.Code())
		}
	})

	// Test case 7: *http.Response input
	t.Run("HttpResponse", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"rpcStatus":"OK"}`))
		}))
		defer ts.Close()

		resp, err := http.Get(ts.URL)
		if err != nil {
			t.Fatal(err)
		}

		result, ok := From(resp)
		if !ok {
			t.Error("Expected ok=true for *http.Response input")
		}
		if result.Code() != codes.OK {
			t.Errorf("Expected code=OK, got %v", result.Code())
		}
		if result.StatusCode() != http.StatusOK {
			t.Errorf("Expected status code=200, got %v", result.StatusCode())
		}
	})

	// Test case 8: error input (context deadline)
	t.Run("ContextDeadlineError", func(t *testing.T) {
		err := context.DeadlineExceeded
		result, ok := From(err)
		if !ok {
			t.Error("Expected ok=true for context.DeadlineExceeded")
		}
		if result.Code() != codes.DeadlineExceeded {
			t.Errorf("Expected code=DeadlineExceeded, got %v", result.Code())
		}
	})

	// Test case 9: error input (url error)
	t.Run("URLError", func(t *testing.T) {
		err := &url.Error{Op: "GET", URL: "http://test", Err: errors.New("test")}
		result, ok := From(err)
		if !ok {
			t.Error("Expected ok=true for url.Error")
		}
		if result.Code() != codes.Unavailable {
			t.Errorf("Expected code=Unavailable, got %v", result.Code())
		}
	})

	// Test case 10: unknown type input
	t.Run("UnknownType", func(t *testing.T) {
		result, ok := From(123)
		if ok {
			t.Error("Expected ok=false for unknown type")
		}
		if result.Code() != codes.Unknown {
			t.Errorf("Expected code=Unknown, got %v", result.Code())
		}
	})
}

func TestFromRpcStatus(t *testing.T) {
	// Test with message
	input := &rpcstatus.Status{
		Code:    int32(codes.OK),
		Message: "test message",
	}
	result := fromRpcStatus(input)
	if result.Code() != codes.OK {
		t.Errorf("Expected code=OK, got %v", result.Code())
	}
	if result.Message() != "test message" {
		t.Errorf("Expected message='test message', got '%v'", result.Message())
	}

	// Test without message
	input = &rpcstatus.Status{Code: int32(codes.NotFound)}
	result = fromRpcStatus(input)
	if result.Code() != codes.NotFound {
		t.Errorf("Expected code=NotFound, got %v", result.Code())
	}
	if result.Message() != "" {
		t.Errorf("Expected empty message, got '%v'", result.Message())
	}
}

func TestFromHttpResponse(t *testing.T) {
	// Test with valid JSON body
	t.Run("WithBody", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(strings.NewReader(`{"rpcStatus":"NOT_FOUND"}`)),
		}
		result, ok := fromHttpResponse(resp)
		if !ok {
			t.Error("Expected ok=true")
		}
		if result.Code() != codes.NotFound {
			t.Errorf("Expected code=NotFound, got %v", result.Code())
		}
		if result.StatusCode() != http.StatusNotFound {
			t.Errorf("Expected status code=404, got %v", result.StatusCode())
		}
	})

	// Test with headers
	t.Run("WithHeaders", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("")),
			Header: http.Header{
				kKey:      []string{strings.Join([]string{"header1", "header2"}, kSeparator)},
				"header1": []string{"value1"},
				"header2": []string{"value2"},
			},
		}
		result, ok := fromHttpResponse(resp)
		if !ok {
			t.Error("Expected ok=true")
		}
		headers := result.Headers()
		if len(headers) == 0 {
			t.Error("Expected non-empty headers")
		}
	})

	// Test with invalid body
	t.Run("InvalidBody", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("invalid json")),
		}
		result, ok := fromHttpResponse(resp)
		if !ok {
			t.Error("Expected ok=true")
		}
		if result.Code() != codes.OK {
			t.Errorf("Expected code=OK, got %v", result.Code())
		}
	})
}

func TestFromError(t *testing.T) {
	// Test context errors
	t.Run("ContextErrors", func(t *testing.T) {
		// DeadlineExceeded
		result, ok := fromError(context.DeadlineExceeded)
		if !ok {
			t.Error("Expected ok=true for context.DeadlineExceeded")
		}
		if result.Code() != codes.DeadlineExceeded {
			t.Errorf("Expected code=DeadlineExceeded, got %v", result.Code())
		}

		// Canceled
		result, ok = fromError(context.Canceled)
		if !ok {
			t.Error("Expected ok=true for context.Canceled")
		}
		if result.Code() != codes.Canceled {
			t.Errorf("Expected code=Canceled, got %v", result.Code())
		}
	})

	// Test URL error
	t.Run("URLError", func(t *testing.T) {
		err := &url.Error{Op: "GET", URL: "http://test", Err: errors.New("test")}
		result, ok := fromError(err)
		if !ok {
			t.Error("Expected ok=true for url.Error")
		}
		if result.Code() != codes.Unavailable {
			t.Errorf("Expected code=Unavailable, got %v", result.Code())
		}
	})

	// Test sampleStatus error
	t.Run("SampleStatusError", func(t *testing.T) {
		err := &sampleStatus{st: &statuspb.Status{RpcStatus: code.Code_INTERNAL}}
		result, ok := fromError(err)
		if !ok {
			t.Error("Expected ok=true for sampleStatus error")
		}
		if result.Code() != codes.Internal {
			t.Errorf("Expected code=Internal, got %v", result.Code())
		}
	})

	// Test gRPC error
	t.Run("GRPCError", func(t *testing.T) {
		err := grpcstatus.New(codes.NotFound, "not found").Err()
		result, ok := fromError(err)
		if !ok {
			t.Error("Expected ok=true for gRPC error")
		}
		if result.Code() != codes.NotFound {
			t.Errorf("Expected code=NotFound, got %v", result.Code())
		}
	})

	// Test unknown error
	t.Run("UnknownError", func(t *testing.T) {
		err := errors.New("some error")
		result, ok := fromError(err)
		if ok {
			t.Error("Expected ok=false for unknown error")
		}
		if result.Code() != codes.Unknown {
			t.Errorf("Expected code=Unknown, got %v", result.Code())
		}
	})
}
