package query

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/cqrs"
	"github.com/go-leo/leo/v3/example/internal/demo/model"
)

type GetUsersArgs struct {
	PageNo   int32
	PageSize int32
}

type GetUsersRes struct {
	List []*model.User
}

type GetUsers cqrs.QueryHandler[*GetUsersArgs, *GetUsersRes]

func NewGetUsers() GetUsers {
	return &getUsers{}
}

type getUsers struct {
}

func (h *getUsers) Handle(ctx context.Context, args *GetUsersArgs) (*GetUsersRes, error) {
	fmt.Println("get users", args)
	users := make([]*model.User, 0)
	for i := 0; i < int(args.PageSize); i++ {
		users = append(users, &model.User{
			UserId: randx.Uint64(),
			Name:   randx.HexString(12),
			Age:    randx.Int31n(50),
			Salary: float64(randx.Int31n(30000)),
			Token:  randx.NumericString(16),
			Avatar: randx.WordString(16),
		})
	}
	return &GetUsersRes{List: users}, nil
}
