// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: error.proto

package status

import (
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	http "google.golang.org/genproto/googleapis/rpc/http"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Error is the error.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// detail is the details.
	DetailInfo *DetailInfo `protobuf:"bytes,1,opt,name=detail_info,json=detailInfo,proto3" json:"detail_info,omitempty"`
	// http_status is the http status.
	HttpStatus *http.HttpResponse `protobuf:"bytes,2,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	// grpc_status is the grpc status.
	GrpcStatus *status.Status `protobuf:"bytes,3,opt,name=grpc_status,json=grpcStatus,proto3" json:"grpc_status,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetDetailInfo() *DetailInfo {
	if x != nil {
		return x.DetailInfo
	}
	return nil
}

func (x *Error) GetHttpStatus() *http.HttpResponse {
	if x != nil {
		return x.HttpStatus
	}
	return nil
}

func (x *Error) GetGrpcStatus() *status.Status {
	if x != nil {
		return x.GrpcStatus
	}
	return nil
}

type DetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This distinguish between two Status objects as being the same when
	// both code and status are identical.
	Identifier *Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// error_info is the error info.
	ErrorInfo *errdetails.ErrorInfo `protobuf:"bytes,2,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	// retry_info is the retry info.
	RetryInfo *errdetails.RetryInfo `protobuf:"bytes,3,opt,name=retry_info,json=retryInfo,proto3" json:"retry_info,omitempty"`
	// debug_info is the debug info.
	DebugInfo *errdetails.DebugInfo `protobuf:"bytes,4,opt,name=debug_info,json=debugInfo,proto3" json:"debug_info,omitempty"`
	// quota_failure is the quota failure.
	QuotaFailure *errdetails.QuotaFailure `protobuf:"bytes,5,opt,name=quota_failure,json=quotaFailure,proto3" json:"quota_failure,omitempty"`
	// precondition_failure is the precondition failure.
	PreconditionFailure *errdetails.PreconditionFailure `protobuf:"bytes,6,opt,name=precondition_failure,json=preconditionFailure,proto3" json:"precondition_failure,omitempty"`
	// bad_request is the bad request.
	BadRequest *errdetails.BadRequest `protobuf:"bytes,7,opt,name=bad_request,json=badRequest,proto3" json:"bad_request,omitempty"`
	// request_info is the request info.
	RequestInfo *errdetails.RequestInfo `protobuf:"bytes,8,opt,name=request_info,json=requestInfo,proto3" json:"request_info,omitempty"`
	// resource_info is the resource info.
	ResourceInfo *errdetails.ResourceInfo `protobuf:"bytes,9,opt,name=resource_info,json=resourceInfo,proto3" json:"resource_info,omitempty"`
	// help is the help.
	Help *errdetails.Help `protobuf:"bytes,10,opt,name=help,proto3" json:"help,omitempty"`
	// localized_message is the localized message.
	LocalizedMessage *errdetails.LocalizedMessage `protobuf:"bytes,11,opt,name=localized_message,json=localizedMessage,proto3" json:"localized_message,omitempty"`
	// details are the other detail.
	Details []*anypb.Any `protobuf:"bytes,12,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *DetailInfo) Reset() {
	*x = DetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailInfo) ProtoMessage() {}

func (x *DetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailInfo.ProtoReflect.Descriptor instead.
func (*DetailInfo) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{1}
}

func (x *DetailInfo) GetIdentifier() *Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *DetailInfo) GetErrorInfo() *errdetails.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *DetailInfo) GetRetryInfo() *errdetails.RetryInfo {
	if x != nil {
		return x.RetryInfo
	}
	return nil
}

func (x *DetailInfo) GetDebugInfo() *errdetails.DebugInfo {
	if x != nil {
		return x.DebugInfo
	}
	return nil
}

func (x *DetailInfo) GetQuotaFailure() *errdetails.QuotaFailure {
	if x != nil {
		return x.QuotaFailure
	}
	return nil
}

func (x *DetailInfo) GetPreconditionFailure() *errdetails.PreconditionFailure {
	if x != nil {
		return x.PreconditionFailure
	}
	return nil
}

func (x *DetailInfo) GetBadRequest() *errdetails.BadRequest {
	if x != nil {
		return x.BadRequest
	}
	return nil
}

func (x *DetailInfo) GetRequestInfo() *errdetails.RequestInfo {
	if x != nil {
		return x.RequestInfo
	}
	return nil
}

func (x *DetailInfo) GetResourceInfo() *errdetails.ResourceInfo {
	if x != nil {
		return x.ResourceInfo
	}
	return nil
}

func (x *DetailInfo) GetHelp() *errdetails.Help {
	if x != nil {
		return x.Help
	}
	return nil
}

func (x *DetailInfo) GetLocalizedMessage() *errdetails.LocalizedMessage {
	if x != nil {
		return x.LocalizedMessage
	}
	return nil
}

func (x *DetailInfo) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

type Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Identifier) Reset() {
	*x = Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identifier) ProtoMessage() {}

func (x *Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identifier.ProtoReflect.Descriptor instead.
func (*Identifier) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{2}
}

func (x *Identifier) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x0b, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x39, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a,
	0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xca, 0x05, 0x0a, 0x0a, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34,
	0x0a, 0x0a, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x52, 0x13, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x61, 0x64, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x62, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x04, 0x68,
	0x65, 0x6c, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x52, 0x04, 0x68, 0x65, 0x6c,
	0x70, 0x12, 0x49, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x10, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x22, 0x0a, 0x0a,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_error_proto_goTypes = []interface{}{
	(*Error)(nil),                          // 0: status.Error
	(*DetailInfo)(nil),                     // 1: status.DetailInfo
	(*Identifier)(nil),                     // 2: status.Identifier
	(*http.HttpResponse)(nil),              // 3: google.rpc.HttpResponse
	(*status.Status)(nil),                  // 4: google.rpc.Status
	(*errdetails.ErrorInfo)(nil),           // 5: google.rpc.ErrorInfo
	(*errdetails.RetryInfo)(nil),           // 6: google.rpc.RetryInfo
	(*errdetails.DebugInfo)(nil),           // 7: google.rpc.DebugInfo
	(*errdetails.QuotaFailure)(nil),        // 8: google.rpc.QuotaFailure
	(*errdetails.PreconditionFailure)(nil), // 9: google.rpc.PreconditionFailure
	(*errdetails.BadRequest)(nil),          // 10: google.rpc.BadRequest
	(*errdetails.RequestInfo)(nil),         // 11: google.rpc.RequestInfo
	(*errdetails.ResourceInfo)(nil),        // 12: google.rpc.ResourceInfo
	(*errdetails.Help)(nil),                // 13: google.rpc.Help
	(*errdetails.LocalizedMessage)(nil),    // 14: google.rpc.LocalizedMessage
	(*anypb.Any)(nil),                      // 15: google.protobuf.Any
}
var file_error_proto_depIdxs = []int32{
	1,  // 0: status.Error.detail_info:type_name -> status.DetailInfo
	3,  // 1: status.Error.http_status:type_name -> google.rpc.HttpResponse
	4,  // 2: status.Error.grpc_status:type_name -> google.rpc.Status
	2,  // 3: status.DetailInfo.identifier:type_name -> status.Identifier
	5,  // 4: status.DetailInfo.error_info:type_name -> google.rpc.ErrorInfo
	6,  // 5: status.DetailInfo.retry_info:type_name -> google.rpc.RetryInfo
	7,  // 6: status.DetailInfo.debug_info:type_name -> google.rpc.DebugInfo
	8,  // 7: status.DetailInfo.quota_failure:type_name -> google.rpc.QuotaFailure
	9,  // 8: status.DetailInfo.precondition_failure:type_name -> google.rpc.PreconditionFailure
	10, // 9: status.DetailInfo.bad_request:type_name -> google.rpc.BadRequest
	11, // 10: status.DetailInfo.request_info:type_name -> google.rpc.RequestInfo
	12, // 11: status.DetailInfo.resource_info:type_name -> google.rpc.ResourceInfo
	13, // 12: status.DetailInfo.help:type_name -> google.rpc.Help
	14, // 13: status.DetailInfo.localized_message:type_name -> google.rpc.LocalizedMessage
	15, // 14: status.DetailInfo.details:type_name -> google.protobuf.Any
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_error_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identifier); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		MessageInfos:      file_error_proto_msgTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
