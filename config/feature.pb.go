// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: protobuf/daemon/config/feature.proto

package config

import (
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

type Feature int32

const (
	Feature_UNKNOWN_FEATURE Feature = 0
	Feature_MESHNET         Feature = 1
	Feature_FILESHARE       Feature = 2
	Feature_NAT_TRAVERSAL   Feature = 3
	Feature_TELIO_ANALYTICS Feature = 4
)

// Enum value maps for Feature.
var (
	Feature_name = map[int32]string{
		0: "UNKNOWN_FEATURE",
		1: "MESHNET",
		2: "FILESHARE",
		3: "NAT_TRAVERSAL",
		4: "TELIO_ANALYTICS",
	}
	Feature_value = map[string]int32{
		"UNKNOWN_FEATURE": 0,
		"MESHNET":         1,
		"FILESHARE":       2,
		"NAT_TRAVERSAL":   3,
		"TELIO_ANALYTICS": 4,
	}
)

func (x Feature) Enum() *Feature {
	p := new(Feature)
	*p = x
	return p
}

func (x Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_daemon_config_feature_proto_enumTypes[0].Descriptor()
}

func (Feature) Type() protoreflect.EnumType {
	return &file_protobuf_daemon_config_feature_proto_enumTypes[0]
}

func (x Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Feature.Descriptor instead.
func (Feature) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_daemon_config_feature_proto_rawDescGZIP(), []int{0}
}

var File_protobuf_daemon_config_feature_proto protoreflect.FileDescriptor

var file_protobuf_daemon_config_feature_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x62,
	0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x4d, 0x45, 0x53, 0x48, 0x4e, 0x45, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46,
	0x49, 0x4c, 0x45, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x41,
	0x54, 0x5f, 0x54, 0x52, 0x41, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x13, 0x0a,
	0x0f, 0x54, 0x45, 0x4c, 0x49, 0x4f, 0x5f, 0x41, 0x4e, 0x41, 0x4c, 0x59, 0x54, 0x49, 0x43, 0x53,
	0x10, 0x04, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x6e, 0x6f,
	0x72, 0x64, 0x76, 0x70, 0x6e, 0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_daemon_config_feature_proto_rawDescOnce sync.Once
	file_protobuf_daemon_config_feature_proto_rawDescData = file_protobuf_daemon_config_feature_proto_rawDesc
)

func file_protobuf_daemon_config_feature_proto_rawDescGZIP() []byte {
	file_protobuf_daemon_config_feature_proto_rawDescOnce.Do(func() {
		file_protobuf_daemon_config_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_daemon_config_feature_proto_rawDescData)
	})
	return file_protobuf_daemon_config_feature_proto_rawDescData
}

var file_protobuf_daemon_config_feature_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_daemon_config_feature_proto_goTypes = []interface{}{
	(Feature)(0), // 0: config.Feature
}
var file_protobuf_daemon_config_feature_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_daemon_config_feature_proto_init() }
func file_protobuf_daemon_config_feature_proto_init() {
	if File_protobuf_daemon_config_feature_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_daemon_config_feature_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_daemon_config_feature_proto_goTypes,
		DependencyIndexes: file_protobuf_daemon_config_feature_proto_depIdxs,
		EnumInfos:         file_protobuf_daemon_config_feature_proto_enumTypes,
	}.Build()
	File_protobuf_daemon_config_feature_proto = out.File
	file_protobuf_daemon_config_feature_proto_rawDesc = nil
	file_protobuf_daemon_config_feature_proto_goTypes = nil
	file_protobuf_daemon_config_feature_proto_depIdxs = nil
}
