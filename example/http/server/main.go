package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-leo/status"
	statuspb "github.com/go-leo/status/example/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

var port = flag.Int("port", 60061, "The server port")

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		content, _ := io.ReadAll(r.Body)
		name := string(content)
		log.Printf("Received: %v", name)
		var st status.Status
		switch name {
		case "Default":
			st = statuspb.ErrDefault(status.ErrorInfo("reason", "domain", map[string]string{"key": "value"}), status.Headers(http.Header{"key": []string{"value"}}))
		case "JustRpcStatus":
			st = statuspb.ErrJustRpcStatus(status.RetryInfo(time.Second), status.Headers(http.Header{"key": []string{"value"}}))
		case "JustHttpStatus":
			st = statuspb.ErrJustHttpStatus(status.DebugInfo([]string{"stack entry"}, "stack entry"), status.Headers(http.Header{"key": []string{"value"}}))
		case "JustMessage":
			st = statuspb.ErrJustMessage(status.QuotaFailure([]*errdetails.QuotaFailure_Violation{{Subject: "subject", Description: "description"}}), status.Headers(http.Header{"key": []string{"value"}}))
		case "AllHave":
			st = statuspb.ErrAllHave(status.PreconditionFailure([]*errdetails.PreconditionFailure_Violation{{Subject: "subject", Description: "description"}}), status.Headers(http.Header{"key": []string{"value"}}))
		case "Custom":
			st = status.New(
				codes.Unknown,
				status.Message("custom message"),
				status.BadRequest([]*errdetails.BadRequest_FieldViolation{{Field: "field", Description: "description"}}),
				status.RequestInfo("request_id", "serving_data"),
				status.ResourceInfo("resource_type", "resource_name", "owner", "description"),
				status.Help([]*errdetails.Help_Link{{Url: "url", Description: "description"}}),
				status.LocalizedMessage("locale", "message"),
				status.Headers(http.Header{"key": []string{"value"}}),
			)
		}
		if st == nil {
			_, _ = w.Write([]byte("Hello " + name))
			return
		}

		var contentType string
		var body []byte
		if jsonBody, marshalErr := st.MarshalJSON(); marshalErr == nil {
			contentType, body = "application/json; charset=utf-8", jsonBody
		}
		w.Header().Set("Content-Type", contentType)
		for k, values := range st.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(st.StatusCode())
		_, _ = w.Write(body)
	})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := http.Server{
		Handler: mux,
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
