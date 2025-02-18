module github.com/go-leo/status

go 1.23.0

require (
	github.com/go-leo/gox v0.0.0-20250214090007-f44ede121342
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250212204824-5a70512c5d8b
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.5
)

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
)

replace github.com/armon/go-metrics v0.4.1 => github.com/hashicorp/go-metrics v0.4.1
