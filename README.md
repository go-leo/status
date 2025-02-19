# Status

status 是一个微服务常用的状态码管理工具，参考 google Status 规范设计。

# 错误代码
下面是一个表格，其中包含google.rpc.Code中定义的所有gRPC错误代码及其原因的简短说明。
<table>
<thead>
<tr><th>HTTP</th><th>RPC</th><th>描述</th></tr>
</thead>
<tbody>
<tr><td>200</td><td><strong>OK</strong></td><td>没有错误</td></tr>
<tr><td>400</td><td><strong>INVALID_ARGUMENT</strong></td><td>客户端指定了无效的参数。 检查错误消息和错误详细信息以获取更多信息。</td></tr>
<tr><td>400</td><td><strong>FAILED_PRECONDITION</strong></td><td>请求不能在当前系统状态下执行，例如删除非空目录。</td></tr>
<tr><td>400</td><td><strong>OUT_OF_RANGE</strong></td><td>客户端指定了无效的范围。</td></tr>
<tr><td>401</td><td><strong>UNAUTHENTICATED</strong></td><td>由于遗失，无效或过期的OAuth令牌而导致请求未通过身份验证。</td></tr>
<tr><td>403</td><td><strong>PERMISSION_DENIED</strong></td><td>客户端没有足够的权限。这可能是因为OAuth令牌没有正确的范围，客户端没有权限，或者客户端项目尚未启用API。</td></tr>
<tr><td>404</td><td><strong>NOT_FOUND</strong></td><td>找不到指定的资源，或者该请求被未公开的原因（例如白名单）拒绝。</td></tr>
<tr><td>409</td><td><strong>ABORTED</strong></td><td>并发冲突，例如读-修改-写冲突。</td></tr>
<tr><td>409</td><td><strong>ALREADY_EXISTS</strong></td><td>客户端尝试创建的资源已存在。</td></tr>
<tr><td>429</td><td><strong>RESOURCE_EXHAUSTED</strong></td><td>资源配额达到速率限制。 客户端应该查找google.rpc.QuotaFailure错误详细信息以获取更多信息。</td></tr>
<tr><td>499</td><td><strong>CANCELLED</strong></td><td>客户端取消请求</td></tr>
<tr><td>500</td><td><strong>DATA_LOSS</strong></td><td>不可恢复的数据丢失或数据损坏。 客户端应该向用户报告错误。</td></tr>
<tr><td>500</td><td><strong>UNKNOWN</strong></td><td>未知的服务器错误。 通常是服务器错误。</td></tr>
<tr><td>500</td><td><strong>INTERNAL</strong></td><td>内部服务错误。 通常是服务器错误。</td></tr>
<tr><td>501</td><td><strong>NOT_IMPLEMENTED</strong></td><td>服务器未实现该API方法。</td></tr>
<tr><td>503</td><td><strong>UNAVAILABLE</strong></td><td>暂停服务。通常是服务器已经关闭。</td></tr>
<tr><td>504</td><td><strong>DEADLINE_EXCEEDED</strong></td><td>已超过请求期限。如果重复发生，请考虑降低请求的复杂性。</td></tr>
</tbody>
</table>

要处理错误，您可以检查返回状态码的描述，并相应地修改您的请求。

# Install
```
go get github.com/go-leo/status@latest
```

# 定义错误
```protobuf
syntax = "proto3";
package status.example.pb;
option go_package = "github.com/go-leo/example/pb;pb";

import "status/annotations.proto";

enum Errors {
  option (status.default_rpc_code) = INTERNAL;

  InvalidName = 0 [(status.rpc_code) = INVALID_ARGUMENT, (status.message) = "名称为空"];

  FileDownloadFailed = 1 [(status.message) = "文件下载失败"];

  FileUploadFailed = 2;
}
```
注意:
* 如果枚举值指定了 rpc_code, 则使用指定的 rpc_code
* 如果枚举值未指定 rpc_code, 则使用 default_rpc_code
* 如果既没有 default_rpc_code 也没有 rpc_code，默认用 Unknown

# 代码生成
```shell
protoc \
		--proto_path=. \
		--proto_path=./third_party \
		--status_out=. \
		--status_opt=paths=source_relative \
		*/*.proto
```
注意事项:
* 需要将 项目中 [code.proto](third_party/google/rpc/code.proto)[third_party](third_party) 和 [annotations.proto](leo/status/annotations.proto) 放到 third_party，参考[example](example)

# 生成后的代码
```go
// Code generated by protoc-gen-status. DO NOT EDIT.

package pb

import (
	status "github.com/go-leo/status"
)

var clean_ErrInvalidName = ErrInvalidName()

func ErrInvalidName(opts ...status.Option) status.Status {
	return status.InvalidArgument(append([]status.Option{status.Identifier("Errors_InvalidName"), status.Message("名称为空")}, opts...)...)
}

func IsInvalidName(err error) bool {
	return clean_ErrInvalidName.Is(status.From(err))
}

var clean_ErrFileDownloadFailed = ErrFileDownloadFailed()

func ErrFileDownloadFailed(opts ...status.Option) status.Status {
	return status.Internal(append([]status.Option{status.Identifier("Errors_FileDownloadFailed"), status.Message("文件下载失败")}, opts...)...)
}

func IsFileDownloadFailed(err error) bool {
	return clean_ErrFileDownloadFailed.Is(status.From(err))
}

var clean_ErrFileUploadFailed = ErrFileUploadFailed()

func ErrFileUploadFailed(opts ...status.Option) status.Status {
	return status.Internal(append([]status.Option{status.Identifier("Errors_FileUploadFailed")}, opts...)...)
}

func IsFileUploadFailed(err error) bool {
	return clean_ErrFileUploadFailed.Is(status.From(err))
}

```
注意事项:
* 上述示例中的`ErrInvalidPassword`会创建并返回一个`status.Status`
* 上述示例中的`IsInvalidPassword`会判断传入的`error`是否是`ErrInvalidPassword`错误(在没有修改`Identifier`情况下)

# gRPC中使用
## Server
```go
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-leo/status"
	"github.com/go-leo/status/example/pb"
	"github.com/google/uuid"
	"log"
	"net"

	helloworldpb "github.com/go-leo/status/example/cmd/grpc/helloworld"
	"google.golang.org/grpc"
)
var (
	port = flag.Int("port", 50051, "The server port")
)
type server struct {
	helloworldpb.UnimplementedGreeterServer
}
func (s *server) SayHello(_ context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	if in.GetName() == "" {
		return nil, pb.ErrInvalidName(status.RequestInfo(uuid.NewString(), in.GetName()))
	}
	log.Printf("Received: %v", in.GetName())
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
```
## Client
```go
package main

import (
	"context"
	"flag"
	"github.com/go-leo/status/example/pb"
	"log"
	"time"

	helloworldpb "github.com/go-leo/status/example/cmd/grpc/helloworld"
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
		st, ok := pb.IsInvalidName(err)
		if ok {
			log.Fatalf("could not greet: %v, identifier: %v, request info: %v", err, st.Identifier(), st.RequestInfo())
		}
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
```
## 运行Server
```
 go run ./main.go 
```

## 运行Client
传名称：
```
go run ./main.go -name world                                                                                                   1 ↵ ──(Thu,Feb20)─┘
```
输出
```
Greeting: Hello world  
```
不传名称：
```
 go run ./main.go                                                                                                                                                                   ──(Thu,Feb20)─┘

```
输出
```
could not greet: rpc error: code = InvalidArgument desc = 名称为空, identifier: Errors_InvalidName, request info: request_id:"95d0a652-898d-4494-81d7-cd3e1cf7200b"
exit status 1
```
注意事项：
* gRPC服务端可以调用`ErrInvalidName`方法，并添加额外的信息，返回`status.Status`，框架会兼容错误信息的响应。
* gRPC客户端可以还原出`status.Status`，同时还原出添加额外的信息。

# HTTP(在gors)中使用


# Reference

* [https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto](https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto)
* [https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto](https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto)
* [https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto](https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto)
* [https://cloud.google.com/apis/design/errors](https://cloud.google.com/apis/design/errors)