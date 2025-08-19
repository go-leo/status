package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-leo/status"
	helloworldpb "github.com/go-leo/status/example/helloworld"
	statuspb "github.com/go-leo/status/example/status"
	"golang.org/x/sync/errgroup"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	http_port = flag.Int("http_port", 60062, "The http server port")
	grpc_port = flag.Int("grpc_port", 50052, "The grpc server port")
)

func main() {
	flag.Parse()
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		// http bff
		conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", *grpc_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		client := helloworldpb.NewGreeterClient(conn)
		mux := http.NewServeMux()
		mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			content, _ := io.ReadAll(r.Body)
			name := string(content)
			log.Printf("Received: %v", name)
			resp, err := client.SayHello(r.Context(), &helloworldpb.HelloRequest{
				Name: name,
			})
			if err == nil {
				_, _ = w.Write([]byte(resp.GetMessage()))
				return
			}
			st, ok := status.FromError(err)
			if !ok {
				_, _ = w.Write([]byte(err.Error()))
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
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *http_port))
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
		return nil
	})
	eg.Go(func() error {
		// grpc micro service
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpc_port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		helloworldpb.RegisterGreeterServer(s, &server{})
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	switch in.GetName() {
	case "Default":
		return nil, statuspb.ErrDefault(status.ErrorInfo("reason", "domain", map[string]string{"key": "value"}))
	case "JustRpcStatus":
		return nil, statuspb.ErrJustRpcStatus(status.RetryInfo(time.Second))
	case "JustHttpStatus":
		return nil, statuspb.ErrJustHttpStatus(status.DebugInfo([]string{"stack entry"}, "stack entry"))
	case "JustMessage":
		return nil, statuspb.ErrJustMessage(status.QuotaFailure([]*errdetails.QuotaFailure_Violation{{Subject: "subject", Description: "description"}}))
	case "AllHave":
		return nil, statuspb.ErrAllHave(status.PreconditionFailure([]*errdetails.PreconditionFailure_Violation{{Subject: "subject", Description: "description"}}))
	case "Custom":
		return nil, status.New(
			codes.Unknown,
			status.Message("custom message"),
			status.BadRequest([]*errdetails.BadRequest_FieldViolation{{Field: "field", Description: "description"}}),
			status.RequestInfo("request_id", "serving_data"),
			status.ResourceInfo("resource_type", "resource_name", "owner", "description"),
			status.Help([]*errdetails.Help_Link{{Url: "url", Description: "description"}}),
			status.LocalizedMessage("locale", "message"),
		)
	}
	return &helloworldpb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
