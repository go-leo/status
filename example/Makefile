build-grpc:
	docker build -t status-grpc:v1 ./grpc/server

run-grpc:
	docker run -d -p 50051:50051 status-grpc:v1

call-grpc-default:
	go run ./grpc/client -name Default

call-grpc-JustRpcStatus:
	go run ./grpc/client -name JustRpcStatus

call-grpc-JustHttpStatus:
	go run ./grpc/client -name JustHttpStatus

call-grpc-JustMessage:
	go run ./grpc/client -name JustMessage

call-grpc-AllHave:
	go run ./grpc/client -name AllHave

call-grpc-Custom:
	go run ./grpc/client -name Custom

build-http:
	docker build -t status-http:v1 ./http/server

run-http:
	docker run -d -p 60061:60061 status-http:v1

call-http-default:
	go run ./http/client -name Default

call-http-JustRpcStatus:
	go run ./http/client -name JustRpcStatus

call-http-JustHttpStatus:
	go run ./http/client -name JustHttpStatus

call-http-JustMessage:
	go run ./http/client -name JustMessage

call-http-AllHave:
	go run ./http/client -name AllHave

call-http-Custom:
	go run ./http/client -name Custom

build-mix:
	docker build -t status-mix:v1 ./mix/server

run-mix:
	docker run -d -p 50052:50052 -p 60062:60062 status-mix:v1

call-mix-default:
	go run ./mix/client -name Default

call-mix-JustRpcStatus:
	go run ./mix/client -name JustRpcStatus

call-mix-JustHttpStatus:
	go run ./mix/client -name JustHttpStatus

call-mix-JustMessage:
	go run ./mix/client -name JustMessage

call-mix-AllHave:
	go run ./mix/client -name AllHave

call-mix-Custom:
	go run ./mix/client -name Custom
