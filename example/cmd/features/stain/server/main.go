package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	stdconsul "github.com/hashicorp/consul/api"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	go runApi(8000)
	go runHttp(8080, "red")
	go runHttp(8081, "blue")
	go runHttp(8082, "yellow")
	go runHttp(18080, "red")
	go runHttp(18081, "blue")
	go runHttp(18082, "yellow")
	go runHttp(28080, "red")
	go runHttp(28081, "blue")
	go runHttp(28082, "yellow")
	go runGrpc(9090, "red")
	go runGrpc(9091, "blue")
	go runGrpc(9092, "yellow")
	go runGrpc(19090, "red")
	go runGrpc(19091, "blue")
	go runGrpc(19092, "yellow")
	go runGrpc(29090, "red")
	go runGrpc(29091, "blue")
	go runGrpc(29092, "yellow")
	select {}
}

func runApi(port int) {
	address := ":" + strconv.FormatInt(int64(port), 10)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	httpClient := helloworld.NewGreeterHttpClient(
		"consul://localhost:8500/demo.http?dc=dc1",
		httpx.InstancerBuilder(consulx.Builder{}),
		httpx.BalancerFactory(lbx.RandomFactory{}),
	)

	router := helloworld.AppendGreeterHttpServerRoutes(
		mux.NewRouter(),
		NewGreeterApiService(httpClient, address),
	)
	server := http.Server{Handler: router}
	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := server.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v\n", err)
		}
	}()
	select {}
}

type GreeterApiService struct {
	client  helloworld.GreeterService
	address string
}

func (g GreeterApiService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	hello, err := g.client.SayHello(ctx, request)
	if err != nil {
		return nil, err
	}
	hello.Message = "[client," + request.GetName() + "]" + "[api," + g.address + "]" + hello.GetMessage()
	return hello, nil
}

func NewGreeterApiService(client helloworld.GreeterService, address string) helloworld.GreeterService {
	return &GreeterApiService{client: client, address: address}
}

func runHttp(port int, color string) {
	address := ":" + strconv.FormatInt(int64(port), 10)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcClient := helloworld.NewGreeterGrpcClient(
		"consul://localhost:8500/demo.grpc?dc=dc1",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
		grpcx.InstancerBuilder(consulx.Builder{}),
		grpcx.BalancerFactory(lbx.RandomFactory{}),
	)

	router := helloworld.AppendGreeterHttpServerRoutes(
		mux.NewRouter(),
		NewGreeterHttpService(grpcClient, address, color),
	)
	server := http.Server{Handler: router}
	client, err := stdconsul.NewClient(&stdconsul.Config{
		Address:    "localhost:8500",
		Datacenter: "dc1",
	})
	if err != nil {
		panic(err)
	}
	registrar := consul.NewRegistrar(consul.NewClient(client), &stdconsul.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    "demo.http",
		Tags:    []string{color},
		Port:    port,
		Address: "127.0.0.1",
	}, logx.L())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		registrar.Deregister()
		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Printf("failed to shutdown: %v\n", err)
		}
	}()

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		registrar.Register()
		if err := server.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v\n", err)
		}
	}()
	select {}
}

type GreeterHttpService struct {
	client  helloworld.GreeterService
	address string
	color   string
}

func (g GreeterHttpService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	hello, err := g.client.SayHello(ctx, request)
	if err != nil {
		return nil, err
	}
	hello.Message = "[http," + g.color + "," + g.address + "]" + hello.GetMessage()
	return hello, nil
}

func NewGreeterHttpService(client helloworld.GreeterService, address string, color string) helloworld.GreeterService {
	return &GreeterHttpService{client: client, address: address, color: color}
}

func runGrpc(port int, color string) {
	address := ":" + strconv.FormatInt(int64(port), 10)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc1.NewServer()
	service := helloworld.NewGreeterGrpcServer(
		NewGreeterGrpcService(address, color),
	)
	helloworld.RegisterGreeterServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	client, err := stdconsul.NewClient(&stdconsul.Config{
		Address:    "localhost:8500",
		Datacenter: "dc1",
	})
	if err != nil {
		panic(err)
	}
	registrar := consul.NewRegistrar(consul.NewClient(client), &stdconsul.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    "demo.grpc",
		Tags:    []string{color},
		Port:    port,
		Address: "127.0.0.1",
	}, logx.L())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		registrar.Deregister()
		s.GracefulStop()
	}()

	go func() {
		registrar.Register()
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()
	select {}
}

type GreeterGrpcService struct {
	color   string
	address string
}

func (g GreeterGrpcService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	time.Sleep(time.Second)
	return &helloworld.HelloReply{
		Message: "[http," + g.color + "," + g.address + "]",
	}, nil
}

func NewGreeterGrpcService(address, color string) helloworld.GreeterService {
	return &GreeterGrpcService{address: address, color: color}
}
