package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/go-leo/status"
	helloworldpb "github.com/go-leo/status/example/helloworld"
	statuspb "github.com/go-leo/status/example/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworldpb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworldpb.HelloRequest{Name: *name})
	if err != nil {
		var st status.Status
		var ok bool
		if st, ok = statuspb.IsDefault(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("default error: %v, json: %s", st, jsonData)
		} else if st, ok = statuspb.IsJustRpcStatus(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("just rpc status error: %v, json: %s", st, jsonData)
		} else if st, ok = statuspb.IsJustHttpStatus(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("just http status error: %v, json: %s", st, jsonData)
		} else if st, ok = statuspb.IsJustMessage(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("just message error: %v, json: %s", st, jsonData)
		} else if st, ok = statuspb.IsAllHave(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("all have error: %v, json: %s", st, jsonData)
		} else {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("custom error: %v, json: %s", st, jsonData)
		}
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
