// Code generated by protoc-gen-leo-cqrs. DO NOT EDIT.

package demo

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	command "github.com/go-leo/leo/v3/example/internal/demo/command"
	query "github.com/go-leo/leo/v3/example/internal/demo/query"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// DemoAssembler responsible for completing the transformation between domain model objects and DTOs
type DemoAssembler interface {
	// FromCreateUserRequest convert request to query arguments
	FromCreateUserRequest(ctx context.Context, request *CreateUserRequest) (*query.CreateUserArgs, context.Context, error)
	// ToCreateUserResponse convert query result to response
	ToCreateUserResponse(ctx context.Context, request *CreateUserRequest, res *query.CreateUserRes) (*CreateUserResponse, error)
	// FromDeleteUserRequest convert request to command arguments
	FromDeleteUserRequest(ctx context.Context, request *DeleteUsersRequest) (*command.DeleteUserArgs, context.Context, error)
	// FromUpdateUserRequest convert request to command arguments
	FromUpdateUserRequest(ctx context.Context, request *UpdateUserRequest) (*command.UpdateUserArgs, context.Context, error)
	// FromGetUserRequest convert request to query arguments
	FromGetUserRequest(ctx context.Context, request *GetUserRequest) (*query.GetUserArgs, context.Context, error)
	// ToGetUserResponse convert query result to response
	ToGetUserResponse(ctx context.Context, request *GetUserRequest, res *query.GetUserRes) (*GetUserResponse, error)
	// FromGetUsersRequest convert request to query arguments
	FromGetUsersRequest(ctx context.Context, request *GetUsersRequest) (*query.GetUsersArgs, context.Context, error)
	// ToGetUsersResponse convert query result to response
	ToGetUsersResponse(ctx context.Context, request *GetUsersRequest, res *query.GetUsersRes) (*GetUsersResponse, error)
	// FromUploadUserAvatarRequest convert request to command arguments
	FromUploadUserAvatarRequest(ctx context.Context, request *UploadUserAvatarRequest) (*command.UploadUserAvatarArgs, context.Context, error)
	// FromGetUserAvatarRequest convert request to query arguments
	FromGetUserAvatarRequest(ctx context.Context, request *GetUserAvatarRequest) (*query.GetUserAvatarArgs, context.Context, error)
	// ToGetUserAvatarResponse convert query result to response
	ToGetUserAvatarResponse(ctx context.Context, request *GetUserAvatarRequest, res *query.GetUserAvatarRes) (*httpbody.HttpBody, error)
}

// demoCqrsService implement the DemoService with CQRS pattern
type demoCqrsService struct {
	bus       cqrs.Bus
	assembler DemoAssembler
}

func (svc *demoCqrsService) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	args, ctx, err := svc.assembler.FromCreateUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	res, err := svc.bus.Query(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToCreateUserResponse(ctx, request, res.(*query.CreateUserRes))
}

func (svc *demoCqrsService) DeleteUser(ctx context.Context, request *DeleteUsersRequest) (*emptypb.Empty, error) {
	command, ctx, err := svc.assembler.FromDeleteUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if err := svc.bus.Exec(ctx, command); err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (svc *demoCqrsService) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*emptypb.Empty, error) {
	command, ctx, err := svc.assembler.FromUpdateUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if err := svc.bus.Exec(ctx, command); err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (svc *demoCqrsService) GetUser(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	args, ctx, err := svc.assembler.FromGetUserRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	res, err := svc.bus.Query(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToGetUserResponse(ctx, request, res.(*query.GetUserRes))
}

func (svc *demoCqrsService) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	args, ctx, err := svc.assembler.FromGetUsersRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	res, err := svc.bus.Query(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToGetUsersResponse(ctx, request, res.(*query.GetUsersRes))
}

func (svc *demoCqrsService) UploadUserAvatar(ctx context.Context, request *UploadUserAvatarRequest) (*emptypb.Empty, error) {
	command, ctx, err := svc.assembler.FromUploadUserAvatarRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	if err := svc.bus.Exec(ctx, command); err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (svc *demoCqrsService) GetUserAvatar(ctx context.Context, request *GetUserAvatarRequest) (*httpbody.HttpBody, error) {
	args, ctx, err := svc.assembler.FromGetUserAvatarRequest(ctx, request)
	if err != nil {
		return nil, err
	}
	res, err := svc.bus.Query(ctx, args)
	if err != nil {
		return nil, err
	}
	return svc.assembler.ToGetUserAvatarResponse(ctx, request, res.(*query.GetUserAvatarRes))
}

func NewDemoCqrsService(
	createUser query.CreateUser,
	deleteUser command.DeleteUser,
	updateUser command.UpdateUser,
	getUser query.GetUser,
	getUsers query.GetUsers,
	uploadUserAvatar command.UploadUserAvatar,
	getUserAvatar query.GetUserAvatar,
	assembler DemoAssembler,
) (DemoService, error) {
	var bus cqrs.SampleBus
	if err := bus.RegisterQuery(createUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterCommand(deleteUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterCommand(updateUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterQuery(getUser); err != nil {
		return nil, err
	}
	if err := bus.RegisterQuery(getUsers); err != nil {
		return nil, err
	}
	if err := bus.RegisterCommand(uploadUserAvatar); err != nil {
		return nil, err
	}
	if err := bus.RegisterQuery(getUserAvatar); err != nil {
		return nil, err
	}
	return &demoCqrsService{
		bus:       &bus,
		assembler: assembler,
	}, nil
}
