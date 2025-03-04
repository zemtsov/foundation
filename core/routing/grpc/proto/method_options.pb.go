// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.3
// source: method_options.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define enum for method type.
type MethodType int32

const (
	MethodType_METHOD_TYPE_TRANSACTION MethodType = 0
	MethodType_METHOD_TYPE_INVOKE      MethodType = 1
	MethodType_METHOD_TYPE_QUERY       MethodType = 2
)

// Enum value maps for MethodType.
var (
	MethodType_name = map[int32]string{
		0: "METHOD_TYPE_TRANSACTION",
		1: "METHOD_TYPE_INVOKE",
		2: "METHOD_TYPE_QUERY",
	}
	MethodType_value = map[string]int32{
		"METHOD_TYPE_TRANSACTION": 0,
		"METHOD_TYPE_INVOKE":      1,
		"METHOD_TYPE_QUERY":       2,
	}
)

func (x MethodType) Enum() *MethodType {
	p := new(MethodType)
	*p = x
	return p
}

func (x MethodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodType) Descriptor() protoreflect.EnumDescriptor {
	return file_method_options_proto_enumTypes[0].Descriptor()
}

func (MethodType) Type() protoreflect.EnumType {
	return &file_method_options_proto_enumTypes[0]
}

func (x MethodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodType.Descriptor instead.
func (MethodType) EnumDescriptor() ([]byte, []int) {
	return file_method_options_proto_rawDescGZIP(), []int{0}
}

// Define enum for method authorization.
type MethodAuth int32

const (
	MethodAuth_METHOD_AUTH_DEFAULT  MethodAuth = 0
	MethodAuth_METHOD_AUTH_ENABLED  MethodAuth = 1
	MethodAuth_METHOD_AUTH_DISABLED MethodAuth = 2
)

// Enum value maps for MethodAuth.
var (
	MethodAuth_name = map[int32]string{
		0: "METHOD_AUTH_DEFAULT",
		1: "METHOD_AUTH_ENABLED",
		2: "METHOD_AUTH_DISABLED",
	}
	MethodAuth_value = map[string]int32{
		"METHOD_AUTH_DEFAULT":  0,
		"METHOD_AUTH_ENABLED":  1,
		"METHOD_AUTH_DISABLED": 2,
	}
)

func (x MethodAuth) Enum() *MethodAuth {
	p := new(MethodAuth)
	*p = x
	return p
}

func (x MethodAuth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MethodAuth) Descriptor() protoreflect.EnumDescriptor {
	return file_method_options_proto_enumTypes[1].Descriptor()
}

func (MethodAuth) Type() protoreflect.EnumType {
	return &file_method_options_proto_enumTypes[1]
}

func (x MethodAuth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MethodAuth.Descriptor instead.
func (MethodAuth) EnumDescriptor() ([]byte, []int) {
	return file_method_options_proto_rawDescGZIP(), []int{1}
}

var file_method_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodType)(nil),
		Field:         50001,
		Name:          "foundation.method_type",
		Tag:           "varint,50001,opt,name=method_type,enum=foundation.MethodType",
		Filename:      "method_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodAuth)(nil),
		Field:         50002,
		Name:          "foundation.method_auth",
		Tag:           "varint,50002,opt,name=method_auth,enum=foundation.MethodAuth",
		Filename:      "method_options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional foundation.MethodType method_type = 50001;
	E_MethodType = &file_method_options_proto_extTypes[0]
	// optional foundation.MethodAuth method_auth = 50002;
	E_MethodAuth = &file_method_options_proto_extTypes[1]
)

var File_method_options_proto protoreflect.FileDescriptor

var file_method_options_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x58, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12,
	0x16, 0x0a, 0x12, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x02, 0x2a, 0x58,
	0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x13,
	0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x44, 0x49,
	0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x3a, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x3a, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd2, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x75, 0x74, 0x68, 0x42, 0x41,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x6f,
	0x69, 0x64, 0x65, 0x61, 0x6f, 0x70, 0x65, 0x6e, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_method_options_proto_rawDescOnce sync.Once
	file_method_options_proto_rawDescData = file_method_options_proto_rawDesc
)

func file_method_options_proto_rawDescGZIP() []byte {
	file_method_options_proto_rawDescOnce.Do(func() {
		file_method_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_method_options_proto_rawDescData)
	})
	return file_method_options_proto_rawDescData
}

var file_method_options_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_method_options_proto_goTypes = []any{
	(MethodType)(0),                    // 0: foundation.MethodType
	(MethodAuth)(0),                    // 1: foundation.MethodAuth
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_method_options_proto_depIdxs = []int32{
	2, // 0: foundation.method_type:extendee -> google.protobuf.MethodOptions
	2, // 1: foundation.method_auth:extendee -> google.protobuf.MethodOptions
	0, // 2: foundation.method_type:type_name -> foundation.MethodType
	1, // 3: foundation.method_auth:type_name -> foundation.MethodAuth
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_method_options_proto_init() }
func file_method_options_proto_init() {
	if File_method_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_method_options_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_method_options_proto_goTypes,
		DependencyIndexes: file_method_options_proto_depIdxs,
		EnumInfos:         file_method_options_proto_enumTypes,
		ExtensionInfos:    file_method_options_proto_extTypes,
	}.Build()
	File_method_options_proto = out.File
	file_method_options_proto_rawDesc = nil
	file_method_options_proto_goTypes = nil
	file_method_options_proto_depIdxs = nil
}
