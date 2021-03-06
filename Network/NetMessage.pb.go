// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: NetMessage.proto

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

type NetMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head    *ProtocolHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	Content []byte        `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *NetMessage) Reset() {
	*x = NetMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NetMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessage) ProtoMessage() {}

func (x *NetMessage) ProtoReflect() protoreflect.Message {
	mi := &file_NetMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMessage.ProtoReflect.Descriptor instead.
func (*NetMessage) Descriptor() ([]byte, []int) {
	return file_NetMessage_proto_rawDescGZIP(), []int{0}
}

func (x *NetMessage) GetHead() *ProtocolHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *NetMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type ProtocolHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId      uint32 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"`
	Cmd      uint32 `protobuf:"varint,2,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Index    uint32 `protobuf:"varint,3,opt,name=Index,proto3" json:"Index,omitempty"`
	DataSize uint32 `protobuf:"varint,4,opt,name=DataSize,proto3" json:"DataSize,omitempty"`
	CheckSum uint32 `protobuf:"varint,5,opt,name=CheckSum,proto3" json:"CheckSum,omitempty"`
}

func (x *ProtocolHead) Reset() {
	*x = ProtocolHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NetMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtocolHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtocolHead) ProtoMessage() {}

func (x *ProtocolHead) ProtoReflect() protoreflect.Message {
	mi := &file_NetMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtocolHead.ProtoReflect.Descriptor instead.
func (*ProtocolHead) Descriptor() ([]byte, []int) {
	return file_NetMessage_proto_rawDescGZIP(), []int{1}
}

func (x *ProtocolHead) GetUId() uint32 {
	if x != nil {
		return x.UId
	}
	return 0
}

func (x *ProtocolHead) GetCmd() uint32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *ProtocolHead) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ProtocolHead) GetDataSize() uint32 {
	if x != nil {
		return x.DataSize
	}
	return 0
}

func (x *ProtocolHead) GetCheckSum() uint32 {
	if x != nil {
		return x.CheckSum
	}
	return 0
}

var File_NetMessage_proto protoreflect.FileDescriptor

var file_NetMessage_proto_rawDesc = []byte{
	0x0a, 0x10, 0x4e, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x51, 0x0a, 0x0a, 0x4e,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x48, 0x65, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04,
	0x48, 0x65, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x80,
	0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x55, 0x49,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x43, 0x6d, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75,
	0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NetMessage_proto_rawDescOnce sync.Once
	file_NetMessage_proto_rawDescData = file_NetMessage_proto_rawDesc
)

func file_NetMessage_proto_rawDescGZIP() []byte {
	file_NetMessage_proto_rawDescOnce.Do(func() {
		file_NetMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_NetMessage_proto_rawDescData)
	})
	return file_NetMessage_proto_rawDescData
}

var file_NetMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_NetMessage_proto_goTypes = []interface{}{
	(*NetMessage)(nil),   // 0: Network.NetMessage
	(*ProtocolHead)(nil), // 1: Network.ProtocolHead
}
var file_NetMessage_proto_depIdxs = []int32{
	1, // 0: Network.NetMessage.Head:type_name -> Network.ProtocolHead
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_NetMessage_proto_init() }
func file_NetMessage_proto_init() {
	if File_NetMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NetMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMessage); i {
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
		file_NetMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtocolHead); i {
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
			RawDescriptor: file_NetMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NetMessage_proto_goTypes,
		DependencyIndexes: file_NetMessage_proto_depIdxs,
		MessageInfos:      file_NetMessage_proto_msgTypes,
	}.Build()
	File_NetMessage_proto = out.File
	file_NetMessage_proto_rawDesc = nil
	file_NetMessage_proto_goTypes = nil
	file_NetMessage_proto_depIdxs = nil
}
