// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: pb/error.proto

package pb

import (
	_ "github.com/go-leo/status"
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
	Errors_OK                         Errors = 0
	Errors_DownloadCancelled          Errors = 1
	Errors_RequestGoogleFailed        Errors = 2
	Errors_InvalidPassword            Errors = 3
	Errors_ResourceExpired            Errors = 4
	Errors_UserNotFound               Errors = 5
	Errors_UserAlreadyExists          Errors = 6
	Errors_DeleteUserPermissionDenied Errors = 7
	Errors_TokenExpired               Errors = 16
	Errors_HDD_FULL                   Errors = 8
	Errors_UserDeleted                Errors = 9
	Errors_UserCreateFailed           Errors = 10
	Errors_VideoOutOfRange            Errors = 11
	Errors_MethodNotImplemented       Errors = 12
	Errors_RedisCrash                 Errors = 13
	Errors_RequestRejected            Errors = 14
	Errors_MysqlDataLoss              Errors = 15
	Errors_FileDownloadFailed         Errors = 17
	Errors_FileUploadFailed           Errors = 18
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0:  "OK",
		1:  "DownloadCancelled",
		2:  "RequestGoogleFailed",
		3:  "InvalidPassword",
		4:  "ResourceExpired",
		5:  "UserNotFound",
		6:  "UserAlreadyExists",
		7:  "DeleteUserPermissionDenied",
		16: "TokenExpired",
		8:  "HDD_FULL",
		9:  "UserDeleted",
		10: "UserCreateFailed",
		11: "VideoOutOfRange",
		12: "MethodNotImplemented",
		13: "RedisCrash",
		14: "RequestRejected",
		15: "MysqlDataLoss",
		17: "FileDownloadFailed",
		18: "FileUploadFailed",
	}
	Errors_value = map[string]int32{
		"OK":                         0,
		"DownloadCancelled":          1,
		"RequestGoogleFailed":        2,
		"InvalidPassword":            3,
		"ResourceExpired":            4,
		"UserNotFound":               5,
		"UserAlreadyExists":          6,
		"DeleteUserPermissionDenied": 7,
		"TokenExpired":               16,
		"HDD_FULL":                   8,
		"UserDeleted":                9,
		"UserCreateFailed":           10,
		"VideoOutOfRange":            11,
		"MethodNotImplemented":       12,
		"RedisCrash":                 13,
		"RequestRejected":            14,
		"MysqlDataLoss":              15,
		"FileDownloadFailed":         17,
		"FileUploadFailed":           18,
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
	return file_pb_error_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_pb_error_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_pb_error_proto_rawDescGZIP(), []int{0}
}

var File_pb_error_proto protoreflect.FileDescriptor

var file_pb_error_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x70, 0x62, 0x1a, 0x18, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd9, 0x06,
	0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x1a, 0x0e, 0xb0, 0xb7, 0x22, 0x00, 0xda, 0x94, 0x28, 0x06, 0xe6, 0x88, 0x90, 0xe5, 0x8a, 0x9f,
	0x12, 0x2e, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x1a, 0x17, 0xb0, 0xb7, 0x22, 0x01, 0xda, 0x94, 0x28,
	0x0f, 0xe4, 0xb8, 0x8b, 0xe8, 0xbd, 0xbd, 0xe8, 0xa2, 0xab, 0xe5, 0x8f, 0x96, 0xe6, 0xb6, 0x88,
	0x12, 0x33, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x1a, 0x1a, 0xb0, 0xb7, 0x22, 0x02, 0xda,
	0x94, 0x28, 0x12, 0xe8, 0xaf, 0xb7, 0xe6, 0xb1, 0x82, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0xe5,
	0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x29, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x03, 0x1a, 0x14, 0xb0, 0xb7, 0x22, 0x03,
	0xda, 0x94, 0x28, 0x0c, 0xe5, 0xaf, 0x86, 0xe7, 0xa0, 0x81, 0xe6, 0x97, 0xa0, 0xe6, 0x95, 0x88,
	0x12, 0x29, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x10, 0x04, 0x1a, 0x14, 0xb0, 0xb7, 0x22, 0x04, 0xda, 0x94, 0x28, 0x0c, 0xe8,
	0xb5, 0x84, 0xe6, 0xba, 0x90, 0xe8, 0xbf, 0x87, 0xe6, 0x9c, 0x9f, 0x12, 0x29, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x05, 0x1a, 0x17, 0xb0,
	0xb7, 0x22, 0x05, 0xda, 0x94, 0x28, 0x0f, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe6, 0x9c, 0xaa,
	0xe6, 0x89, 0xbe, 0xe5, 0x88, 0xb0, 0x12, 0x2e, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x10, 0x06, 0x1a, 0x17, 0xb0,
	0xb7, 0x22, 0x06, 0xda, 0x94, 0x28, 0x0f, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe5, 0xb7, 0xb2,
	0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x12, 0x3a, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x6e, 0x69, 0x65, 0x64, 0x10, 0x07, 0x1a, 0x1a, 0xb0, 0xb7, 0x22, 0x07, 0xda, 0x94, 0x28, 0x12,
	0xe6, 0x97, 0xa0, 0xe6, 0x9d, 0x83, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe7, 0x94, 0xa8, 0xe6,
	0x88, 0xb7, 0x12, 0x28, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x10, 0x10, 0x1a, 0x16, 0xb0, 0xb7, 0x22, 0x10, 0xda, 0x94, 0x28, 0x0e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0xe5, 0xb7, 0xb2, 0xe8, 0xbf, 0x87, 0xe6, 0x9c, 0x9f, 0x12, 0x28, 0x0a, 0x08,
	0x48, 0x44, 0x44, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x08, 0x1a, 0x1a, 0xb0, 0xb7, 0x22, 0x08,
	0xda, 0x94, 0x28, 0x12, 0xe7, 0xa1, 0xac, 0xe7, 0x9b, 0x98, 0xe7, 0xa9, 0xba, 0xe9, 0x97, 0xb4,
	0xe4, 0xb8, 0x8d, 0xe8, 0xb6, 0xb3, 0x12, 0x2e, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x09, 0x1a, 0x1d, 0xb0, 0xb7, 0x22, 0x09, 0xda, 0x94, 0x28,
	0x15, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe5, 0xb7, 0xb2, 0xe7, 0xbb, 0x8f, 0xe8, 0xa2, 0xab,
	0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0x12, 0x30, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x0a, 0x1a, 0x1a, 0xb0, 0xb7,
	0x22, 0x0a, 0xda, 0x94, 0x28, 0x12, 0xe5, 0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe7, 0x94, 0xa8, 0xe6,
	0x88, 0xb7, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2f, 0x0a, 0x0f, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x0b, 0x1a, 0x1a, 0xb0,
	0xb7, 0x22, 0x0b, 0xda, 0x94, 0x28, 0x12, 0xe8, 0xa7, 0x86, 0xe9, 0xa2, 0x91, 0xe8, 0xb6, 0x85,
	0xe5, 0x87, 0xba, 0xe8, 0x8c, 0x83, 0xe5, 0x9b, 0xb4, 0x12, 0x31, 0x0a, 0x14, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4e, 0x6f, 0x74, 0x49, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65,
	0x64, 0x10, 0x0c, 0x1a, 0x17, 0xb0, 0xb7, 0x22, 0x0c, 0xda, 0x94, 0x28, 0x0f, 0xe6, 0x96, 0xb9,
	0xe6, 0xb3, 0x95, 0xe6, 0x9c, 0xaa, 0xe5, 0xae, 0x9e, 0xe7, 0x8e, 0xb0, 0x12, 0x23, 0x0a, 0x0a,
	0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x72, 0x61, 0x73, 0x68, 0x10, 0x0d, 0x1a, 0x13, 0xb0, 0xb7,
	0x22, 0x0d, 0xda, 0x94, 0x28, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x73, 0xe5, 0xb4, 0xa9, 0xe6, 0xba,
	0x83, 0x12, 0x2c, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x10, 0x0e, 0x1a, 0x17, 0xb0, 0xb7, 0x22, 0x0e, 0xda, 0x94, 0x28, 0x0f,
	0xe8, 0xaf, 0xb7, 0xe6, 0xb1, 0x82, 0xe8, 0xa2, 0xab, 0xe6, 0x8b, 0x92, 0xe7, 0xbb, 0x9d, 0x12,
	0x2c, 0x0a, 0x0d, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x73, 0x73,
	0x10, 0x0f, 0x1a, 0x19, 0xb0, 0xb7, 0x22, 0x0f, 0xda, 0x94, 0x28, 0x11, 0x4d, 0x79, 0x73, 0x71,
	0x6c, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe4, 0xb8, 0xa2, 0xe5, 0xa4, 0xb1, 0x12, 0x2e, 0x0a,
	0x12, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x10, 0x11, 0x1a, 0x16, 0xda, 0x94, 0x28, 0x12, 0xe6, 0x96, 0x87, 0xe4, 0xbb,
	0xb6, 0xe4, 0xb8, 0x8b, 0xe8, 0xbd, 0xbd, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x14, 0x0a,
	0x10, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x10, 0x12, 0x1a, 0x04, 0xa0, 0xe5, 0x1f, 0x0d, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x65, 0x6f, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_error_proto_rawDescOnce sync.Once
	file_pb_error_proto_rawDescData = file_pb_error_proto_rawDesc
)

func file_pb_error_proto_rawDescGZIP() []byte {
	file_pb_error_proto_rawDescOnce.Do(func() {
		file_pb_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_error_proto_rawDescData)
	})
	return file_pb_error_proto_rawDescData
}

var file_pb_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_error_proto_goTypes = []interface{}{
	(Errors)(0), // 0: status.example.pb.Errors
}
var file_pb_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_error_proto_init() }
func file_pb_error_proto_init() {
	if File_pb_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_error_proto_goTypes,
		DependencyIndexes: file_pb_error_proto_depIdxs,
		EnumInfos:         file_pb_error_proto_enumTypes,
	}.Build()
	File_pb_error_proto = out.File
	file_pb_error_proto_rawDesc = nil
	file_pb_error_proto_goTypes = nil
	file_pb_error_proto_depIdxs = nil
}
