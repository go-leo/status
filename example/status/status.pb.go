// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.3
// source: status.proto

package status

import (
	_ "github.com/go-leo/status/proto/leo/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Errors int32

const (
	Errors_Default        Errors = 0
	Errors_JustRpcStatus  Errors = 1
	Errors_JustHttpStatus Errors = 2
	Errors_JustMessage    Errors = 3
	Errors_AllHave        Errors = 4
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0: "Default",
		1: "JustRpcStatus",
		2: "JustHttpStatus",
		3: "JustMessage",
		4: "AllHave",
	}
	Errors_value = map[string]int32{
		"Default":        0,
		"JustRpcStatus":  1,
		"JustHttpStatus": 2,
		"JustMessage":    3,
		"AllHave":        4,
	}
)

func (x Errors) Enum() *Errors {
	p := new(Errors)
	*p = x
	return p
}

func (x Errors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Errors) Descriptor() protoreflect.EnumDescriptor {
	return file_status_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_status_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{0}
}

var File_status_proto protoreflect.FileDescriptor

var file_status_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x6c, 0x65, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x6c, 0x65, 0x6f, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa6, 0x01, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x0d, 0x4a, 0x75, 0x73, 0x74, 0x52, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x10, 0x01, 0x1a, 0x05, 0x88, 0x94, 0x9b, 0x0b, 0x03, 0x12, 0x1b, 0x0a, 0x0e, 0x4a, 0x75, 0x73,
	0x74, 0x48, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x02, 0x1a, 0x07, 0xc0,
	0x91, 0xf4, 0xb4, 0x02, 0x90, 0x03, 0x12, 0x23, 0x0a, 0x0b, 0x4a, 0x75, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x03, 0x1a, 0x12, 0xc2, 0xa0, 0xd5, 0xab, 0x08, 0x0c, 0x6a,
	0x75, 0x73, 0x74, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x41,
	0x6c, 0x6c, 0x48, 0x61, 0x76, 0x65, 0x10, 0x04, 0x1a, 0x1a, 0x88, 0x94, 0x9b, 0x0b, 0x03, 0xc0,
	0x91, 0xf4, 0xb4, 0x02, 0x91, 0x03, 0xc2, 0xa0, 0xd5, 0xab, 0x08, 0x08, 0x61, 0x6c, 0x6c, 0x20,
	0x68, 0x61, 0x76, 0x65, 0x1a, 0x0a, 0x88, 0xac, 0x2f, 0x0d, 0xd0, 0xa7, 0xc8, 0x03, 0xf4, 0x03,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_status_proto_rawDescOnce sync.Once
	file_status_proto_rawDescData = file_status_proto_rawDesc
)

func file_status_proto_rawDescGZIP() []byte {
	file_status_proto_rawDescOnce.Do(func() {
		file_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_status_proto_rawDescData)
	})
	return file_status_proto_rawDescData
}

var file_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_status_proto_goTypes = []any{
	(Errors)(0), // 0: leo.example.status.errors.Errors
}
var file_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_status_proto_init() }
func file_status_proto_init() {
	if File_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_status_proto_goTypes,
		DependencyIndexes: file_status_proto_depIdxs,
		EnumInfos:         file_status_proto_enumTypes,
	}.Build()
	File_status_proto = out.File
	file_status_proto_rawDesc = nil
	file_status_proto_goTypes = nil
	file_status_proto_depIdxs = nil
}
