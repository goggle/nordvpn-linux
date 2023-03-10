// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: cities.proto

package pb

import (
	config "github.com/NordSecurity/nordvpn-linux/config"
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

type CitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol  config.Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=config.Protocol" json:"protocol,omitempty"`
	Obfuscate bool            `protobuf:"varint,2,opt,name=obfuscate,proto3" json:"obfuscate,omitempty"`
	Country   string          `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *CitiesRequest) Reset() {
	*x = CitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitiesRequest) ProtoMessage() {}

func (x *CitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitiesRequest.ProtoReflect.Descriptor instead.
func (*CitiesRequest) Descriptor() ([]byte, []int) {
	return file_cities_proto_rawDescGZIP(), []int{0}
}

func (x *CitiesRequest) GetProtocol() config.Protocol {
	if x != nil {
		return x.Protocol
	}
	return config.Protocol(0)
}

func (x *CitiesRequest) GetObfuscate() bool {
	if x != nil {
		return x.Obfuscate
	}
	return false
}

func (x *CitiesRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

var File_cities_proto protoreflect.FileDescriptor

var file_cities_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0d, 0x43, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62, 0x66, 0x75,
	0x73, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x62, 0x66,
	0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x6f, 0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x6e, 0x6f, 0x72, 0x64,
	0x76, 0x70, 0x6e, 0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cities_proto_rawDescOnce sync.Once
	file_cities_proto_rawDescData = file_cities_proto_rawDesc
)

func file_cities_proto_rawDescGZIP() []byte {
	file_cities_proto_rawDescOnce.Do(func() {
		file_cities_proto_rawDescData = protoimpl.X.CompressGZIP(file_cities_proto_rawDescData)
	})
	return file_cities_proto_rawDescData
}

var file_cities_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cities_proto_goTypes = []interface{}{
	(*CitiesRequest)(nil), // 0: pb.CitiesRequest
	(config.Protocol)(0),  // 1: config.Protocol
}
var file_cities_proto_depIdxs = []int32{
	1, // 0: pb.CitiesRequest.protocol:type_name -> config.Protocol
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cities_proto_init() }
func file_cities_proto_init() {
	if File_cities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CitiesRequest); i {
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
			RawDescriptor: file_cities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cities_proto_goTypes,
		DependencyIndexes: file_cities_proto_depIdxs,
		MessageInfos:      file_cities_proto_msgTypes,
	}.Build()
	File_cities_proto = out.File
	file_cities_proto_rawDesc = nil
	file_cities_proto_goTypes = nil
	file_cities_proto_depIdxs = nil
}