package main

import (
	"fmt"
	"github.com/go-leo/status"
	"github.com/go-leo/status/example/pb"
)

func main() {
	st := pb.ErrUserNotFound()
	fmt.Println(st)
	is := pb.IsUserNotFound(st)
	fmt.Println(is)

	st = pb.ErrUserNotFound(status.Message("查找用户失败"))
	fmt.Println(st)
	is = pb.IsUserNotFound(st)
	fmt.Println(is)
}
