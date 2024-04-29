package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/demo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	endpoints := demo.NewdemoServiceEndpoints(newDemoService())
	service := demo.NewDemoServiceGRPCServer(endpoints, []endpoint.Middleware{})
	demo.RegisterDemoServiceServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type demoService struct {
}

func (d demoService) CreateUser(ctx context.Context, request *demo.CreateUserRequest) (*emptypb.Empty, error) {
	fmt.Println(request)
	return new(emptypb.Empty), nil
}

func (d demoService) UpdateUser(ctx context.Context, request *demo.UpdateUserRequest) (*emptypb.Empty, error) {
	fmt.Println(request)
	return new(emptypb.Empty), nil
}

func (d demoService) GetUser(ctx context.Context, request *demo.GetUserRequest) (*demo.GetUserResponse, error) {
	fmt.Println(request)
	return &demo.GetUserResponse{
		Name:   request.GetName(),
		Age:    request.GetAge(),
		Salary: request.GetSalary(),
		Token:  request.GetToken(),
	}, nil
}

func (d demoService) GetUsers(ctx context.Context, request *demo.GetUsersRequest) (*demo.GetUsersResponse, error) {
	fmt.Println(request)
	var users []*demo.GetUsersResponse_User
	for i := 0; i < int(request.PageSize); i++ {
		users = append(users, &demo.GetUsersResponse_User{
			Name:   randx.HexString(10),
			Age:    randx.Int31n(100),
			Salary: randx.Float64(),
			Token:  randx.NumericString(10),
		})
	}
	return &demo.GetUsersResponse{
		Users: users,
	}, nil
}

func (d demoService) DeleteUser(ctx context.Context, request *demo.DeleteUsersRequest) (*emptypb.Empty, error) {
	fmt.Println(request)
	return new(emptypb.Empty), nil
}

func newDemoService() demo.DemoServiceServer {
	return &demoService{}
}
