// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package body

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type bodyEndpoints struct {
	svc interface {
		Bool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptBool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepBool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapBool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Int32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptInt32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepInt32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapInt32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Uint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptUint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepUint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapUint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Fixed32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFixed32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFixed32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Int64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptInt64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepInt64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapInt64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Sint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptSint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepSint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Sfixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptSfixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepSfixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Uint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptUint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepUint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapUint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Fixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Float32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFloat32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFloat32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapFloat32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Float64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFloat64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFloat64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapFloat64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		String(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptString(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepString(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapString(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Bytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptBytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepBytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapBytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Enum(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptEnum(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepEnum(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Dictionary(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		HttpBody(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		HttpRequest(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
	}
}

func (e *bodyEndpoints) Bool() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Bool(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptBool() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptBool(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepBool() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepBool(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapBool() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapBool(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Int32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Int32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptInt32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptInt32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepInt32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepInt32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapInt32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapInt32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Uint32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Uint32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptUint32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptUint32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepUint32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepUint32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapUint32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapUint32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Fixed32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Fixed32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptFixed32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptFixed32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepFixed32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepFixed32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Int64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Int64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptInt64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptInt64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepInt64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepInt64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapInt64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapInt64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Sint64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Sint64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptSint64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptSint64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepSint64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepSint64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Sfixed64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Sfixed64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptSfixed64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptSfixed64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepSfixed64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepSfixed64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Uint64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Uint64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptUint64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptUint64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepUint64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepUint64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapUint64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapUint64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Fixed64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Fixed64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptFixed64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptFixed64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepFixed64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepFixed64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Float32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Float32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptFloat32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptFloat32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepFloat32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepFloat32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapFloat32() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapFloat32(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Float64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Float64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptFloat64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptFloat64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepFloat64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepFloat64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapFloat64() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapFloat64(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) String() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.String(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptString(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepString(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapString() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapString(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Bytes() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Bytes(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptBytes() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptBytes(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepBytes() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepBytes(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) WrapBytes() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.WrapBytes(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Enum() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Enum(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) OptEnum() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.OptEnum(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) RepEnum() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.RepEnum(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) Dictionary() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.Dictionary(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) HttpBody() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBody(ctx, request.(*BodyRequest))
	}
}

func (e *bodyEndpoints) HttpRequest() endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpRequest(ctx, request.(*BodyRequest))
	}
}

func NewbodyEndpoints(
	svc interface {
		Bool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptBool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepBool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapBool(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Int32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptInt32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepInt32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapInt32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Uint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptUint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepUint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapUint32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Fixed32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFixed32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFixed32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Int64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptInt64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepInt64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapInt64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Sint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptSint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepSint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Sfixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptSfixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepSfixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Uint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptUint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepUint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapUint64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Fixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFixed64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Float32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFloat32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFloat32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapFloat32(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Float64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptFloat64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepFloat64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapFloat64(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		String(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptString(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepString(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapString(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Bytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptBytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepBytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		WrapBytes(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Enum(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		OptEnum(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		RepEnum(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		Dictionary(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		HttpBody(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
		HttpRequest(ctx context.Context, request *BodyRequest) (*emptypb.Empty, error)
	},
) interface {
	Bool() endpoint.Endpoint
	OptBool() endpoint.Endpoint
	RepBool() endpoint.Endpoint
	WrapBool() endpoint.Endpoint
	Int32() endpoint.Endpoint
	OptInt32() endpoint.Endpoint
	RepInt32() endpoint.Endpoint
	WrapInt32() endpoint.Endpoint
	Uint32() endpoint.Endpoint
	OptUint32() endpoint.Endpoint
	RepUint32() endpoint.Endpoint
	WrapUint32() endpoint.Endpoint
	Fixed32() endpoint.Endpoint
	OptFixed32() endpoint.Endpoint
	RepFixed32() endpoint.Endpoint
	Int64() endpoint.Endpoint
	OptInt64() endpoint.Endpoint
	RepInt64() endpoint.Endpoint
	WrapInt64() endpoint.Endpoint
	Sint64() endpoint.Endpoint
	OptSint64() endpoint.Endpoint
	RepSint64() endpoint.Endpoint
	Sfixed64() endpoint.Endpoint
	OptSfixed64() endpoint.Endpoint
	RepSfixed64() endpoint.Endpoint
	Uint64() endpoint.Endpoint
	OptUint64() endpoint.Endpoint
	RepUint64() endpoint.Endpoint
	WrapUint64() endpoint.Endpoint
	Fixed64() endpoint.Endpoint
	OptFixed64() endpoint.Endpoint
	RepFixed64() endpoint.Endpoint
	Float32() endpoint.Endpoint
	OptFloat32() endpoint.Endpoint
	RepFloat32() endpoint.Endpoint
	WrapFloat32() endpoint.Endpoint
	Float64() endpoint.Endpoint
	OptFloat64() endpoint.Endpoint
	RepFloat64() endpoint.Endpoint
	WrapFloat64() endpoint.Endpoint
	String() endpoint.Endpoint
	OptString() endpoint.Endpoint
	RepString() endpoint.Endpoint
	WrapString() endpoint.Endpoint
	Bytes() endpoint.Endpoint
	OptBytes() endpoint.Endpoint
	RepBytes() endpoint.Endpoint
	WrapBytes() endpoint.Endpoint
	Enum() endpoint.Endpoint
	OptEnum() endpoint.Endpoint
	RepEnum() endpoint.Endpoint
	Dictionary() endpoint.Endpoint
	HttpBody() endpoint.Endpoint
	HttpRequest() endpoint.Endpoint
} {
	return &bodyEndpoints{svc: svc}
}
