// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: RPCMessage.proto

package Network

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RPCArgType int32

const (
	RPCArgType_Unknown  RPCArgType = 0
	RPCArgType_INT      RPCArgType = 1
	RPCArgType_UInt     RPCArgType = 2
	RPCArgType_Long     RPCArgType = 3
	RPCArgType_ULong    RPCArgType = 4
	RPCArgType_Short    RPCArgType = 5
	RPCArgType_UShort   RPCArgType = 6
	RPCArgType_Double   RPCArgType = 7
	RPCArgType_Float    RPCArgType = 8
	RPCArgType_String   RPCArgType = 9
	RPCArgType_Byte     RPCArgType = 10
	RPCArgType_Bool     RPCArgType = 11
	RPCArgType_Bytes    RPCArgType = 31
	RPCArgType_PBObject RPCArgType = 32
)

// Enum value maps for RPCArgType.
var (
	RPCArgType_name = map[int32]string{
		0:  "Unknown",
		1:  "INT",
		2:  "UInt",
		3:  "Long",
		4:  "ULong",
		5:  "Short",
		6:  "UShort",
		7:  "Double",
		8:  "Float",
		9:  "String",
		10: "Byte",
		11: "Bool",
		31: "Bytes",
		32: "PBObject",
	}
	RPCArgType_value = map[string]int32{
		"Unknown":  0,
		"INT":      1,
		"UInt":     2,
		"Long":     3,
		"ULong":    4,
		"Short":    5,
		"UShort":   6,
		"Double":   7,
		"Float":    8,
		"String":   9,
		"Byte":     10,
		"Bool":     11,
		"Bytes":    31,
		"PBObject": 32,
	}
)

func (x RPCArgType) Enum() *RPCArgType {
	p := new(RPCArgType)
	*p = x
	return p
}

func (x RPCArgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RPCArgType) Descriptor() protoreflect.EnumDescriptor {
	return file_RPCMessage_proto_enumTypes[0].Descriptor()
}

func (RPCArgType) Type() protoreflect.EnumType {
	return &file_RPCMessage_proto_enumTypes[0]
}

func (x RPCArgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RPCArgType.Descriptor instead.
func (RPCArgType) EnumDescriptor() ([]byte, []int) {
	return file_RPCMessage_proto_rawDescGZIP(), []int{0}
}

type RPCMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string       `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	RPCRawArgs []*RPCRawArg `protobuf:"bytes,2,rep,name=RPCRawArgs,proto3" json:"RPCRawArgs,omitempty"`
}

func (x *RPCMessage) Reset() {
	*x = RPCMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPCMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCMessage) ProtoMessage() {}

func (x *RPCMessage) ProtoReflect() protoreflect.Message {
	mi := &file_RPCMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCMessage.ProtoReflect.Descriptor instead.
func (*RPCMessage) Descriptor() ([]byte, []int) {
	return file_RPCMessage_proto_rawDescGZIP(), []int{0}
}

func (x *RPCMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RPCMessage) GetRPCRawArgs() []*RPCRawArg {
	if x != nil {
		return x.RPCRawArgs
	}
	return nil
}

type RPCRawArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawValueType RPCArgType `protobuf:"varint,1,opt,name=RawValueType,proto3,enum=Network.RPCArgType" json:"RawValueType,omitempty"`
	RawValue     []byte     `protobuf:"bytes,2,opt,name=RawValue,proto3" json:"RawValue,omitempty"`
}

func (x *RPCRawArg) Reset() {
	*x = RPCRawArg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RPCMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCRawArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCRawArg) ProtoMessage() {}

func (x *RPCRawArg) ProtoReflect() protoreflect.Message {
	mi := &file_RPCMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCRawArg.ProtoReflect.Descriptor instead.
func (*RPCRawArg) Descriptor() ([]byte, []int) {
	return file_RPCMessage_proto_rawDescGZIP(), []int{1}
}

func (x *RPCRawArg) GetRawValueType() RPCArgType {
	if x != nil {
		return x.RawValueType
	}
	return RPCArgType_Unknown
}

func (x *RPCRawArg) GetRawValue() []byte {
	if x != nil {
		return x.RawValue
	}
	return nil
}

var File_RPCMessage_proto protoreflect.FileDescriptor

var file_RPCMessage_proto_rawDesc = []byte{
	0x0a, 0x10, 0x52, 0x50, 0x43, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x54, 0x0a, 0x0a, 0x52,
	0x50, 0x43, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a,
	0x0a, 0x52, 0x50, 0x43, 0x52, 0x61, 0x77, 0x41, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x52, 0x50, 0x43, 0x52,
	0x61, 0x77, 0x41, 0x72, 0x67, 0x52, 0x0a, 0x52, 0x50, 0x43, 0x52, 0x61, 0x77, 0x41, 0x72, 0x67,
	0x73, 0x22, 0x60, 0x0a, 0x09, 0x52, 0x50, 0x43, 0x52, 0x61, 0x77, 0x41, 0x72, 0x67, 0x12, 0x37,
	0x0a, 0x0c, 0x52, 0x61, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x52,
	0x50, 0x43, 0x41, 0x72, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x52, 0x61, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x52, 0x61, 0x77, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0xa8, 0x01, 0x0a, 0x0a, 0x52, 0x50, 0x43, 0x41, 0x72, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x49, 0x6e, 0x74,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x6f, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x55, 0x4c, 0x6f, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x10,
	0x09, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x79, 0x74, 0x65, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x6f, 0x6f, 0x6c, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x10, 0x1f,
	0x12, 0x0c, 0x0a, 0x08, 0x50, 0x42, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10, 0x20, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RPCMessage_proto_rawDescOnce sync.Once
	file_RPCMessage_proto_rawDescData = file_RPCMessage_proto_rawDesc
)

func file_RPCMessage_proto_rawDescGZIP() []byte {
	file_RPCMessage_proto_rawDescOnce.Do(func() {
		file_RPCMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_RPCMessage_proto_rawDescData)
	})
	return file_RPCMessage_proto_rawDescData
}

var file_RPCMessage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RPCMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_RPCMessage_proto_goTypes = []interface{}{
	(RPCArgType)(0),    // 0: Network.RPCArgType
	(*RPCMessage)(nil), // 1: Network.RPCMessage
	(*RPCRawArg)(nil),  // 2: Network.RPCRawArg
}
var file_RPCMessage_proto_depIdxs = []int32{
	2, // 0: Network.RPCMessage.RPCRawArgs:type_name -> Network.RPCRawArg
	0, // 1: Network.RPCRawArg.RawValueType:type_name -> Network.RPCArgType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RPCMessage_proto_init() }
func file_RPCMessage_proto_init() {
	if File_RPCMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RPCMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCMessage); i {
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
		file_RPCMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCRawArg); i {
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
			RawDescriptor: file_RPCMessage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RPCMessage_proto_goTypes,
		DependencyIndexes: file_RPCMessage_proto_depIdxs,
		EnumInfos:         file_RPCMessage_proto_enumTypes,
		MessageInfos:      file_RPCMessage_proto_msgTypes,
	}.Build()
	File_RPCMessage_proto = out.File
	file_RPCMessage_proto_rawDesc = nil
	file_RPCMessage_proto_goTypes = nil
	file_RPCMessage_proto_depIdxs = nil
}
