package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-leo/status"
	helloworldpb "github.com/go-leo/status/example/helloworld"
	statuspb "github.com/go-leo/status/example/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var port = flag.Int("port", 50051, "The server port")

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

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
