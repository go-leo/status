// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: annotations.proto

package status

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*code.Code)(nil),
		Field:         65108,
		Name:          "status.default_rpc_code",
		Tag:           "varint,65108,opt,name=default_rpc_code,enum=google.rpc.Code",
		Filename:      "annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*code.Code)(nil),
		Field:         70518,
		Name:          "status.rpc_code",
		Tag:           "varint,70518,opt,name=rpc_code,enum=google.rpc.Code",
		Filename:      "annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         82251,
		Name:          "status.message",
		Tag:           "bytes,82251,opt,name=message",
		Filename:      "annotations.proto",
	},
}

// Extension fields to descriptorpb.EnumOptions.
var (
	// 如果枚举有默认值，则默认值对应的rpc code
	//
	// optional google.rpc.Code default_rpc_code = 65108;
	E_DefaultRpcCode = &file_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional google.rpc.Code rpc_code = 70518;
	E_RpcCode = &file_annotations_proto_extTypes[1]
	// optional string message = 82251;
	E_Message = &file_annotations_proto_extTypes[2]
)

var File_annotations_proto protoreflect.FileDescriptor

var file_annotations_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x5a, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd4, 0xfc, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x70, 0x63, 0x43, 0x6f, 0x64, 0x65,
	0x3a, 0x50, 0x0a, 0x08, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xf6, 0xa6, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x72, 0x70, 0x63, 0x43, 0x6f,
	0x64, 0x65, 0x3a, 0x3d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xcb, 0x82, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.EnumOptions)(nil),      // 0: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
	(code.Code)(0),                        // 2: google.rpc.Code
}
var file_annotations_proto_depIdxs = []int32{
	0, // 0: status.default_rpc_code:extendee -> google.protobuf.EnumOptions
	1, // 1: status.rpc_code:extendee -> google.protobuf.EnumValueOptions
	1, // 2: status.message:extendee -> google.protobuf.EnumValueOptions
	2, // 3: status.default_rpc_code:type_name -> google.rpc.Code
	2, // 4: status.rpc_code:type_name -> google.rpc.Code
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	3, // [3:5] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_annotations_proto_init() }
func file_annotations_proto_init() {
	if File_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_annotations_proto_goTypes,
		DependencyIndexes: file_annotations_proto_depIdxs,
		ExtensionInfos:    file_annotations_proto_extTypes,
	}.Build()
	File_annotations_proto = out.File
	file_annotations_proto_rawDesc = nil
	file_annotations_proto_goTypes = nil
	file_annotations_proto_depIdxs = nil
}
