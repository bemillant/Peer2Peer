// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/proto.proto

package proto_proto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test string `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

var File_proto_proto_proto protoreflect.FileDescriptor

var file_proto_proto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x50, 0x45, 0x45, 0x52, 0x32, 0x50, 0x45, 0x45, 0x52, 0x22, 0x1d,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x32, 0x38, 0x0a,
	0x05, 0x74, 0x65, 0x73, 0x74, 0x54, 0x12, 0x2f, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x41, 0x12,
	0x12, 0x2e, 0x50, 0x45, 0x45, 0x52, 0x32, 0x50, 0x45, 0x45, 0x52, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x50, 0x45, 0x45, 0x52, 0x32, 0x50, 0x45, 0x45, 0x52, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65,
	0x6d, 0x69, 0x6c, 0x6c, 0x61, 0x6e, 0x74, 0x2f, 0x50, 0x65, 0x65, 0x72, 0x32, 0x50, 0x65, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_proto_proto_rawDescOnce sync.Once
	file_proto_proto_proto_rawDescData = file_proto_proto_proto_rawDesc
)

func file_proto_proto_proto_rawDescGZIP() []byte {
	file_proto_proto_proto_rawDescOnce.Do(func() {
		file_proto_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_proto_proto_rawDescData)
	})
	return file_proto_proto_proto_rawDescData
}

var file_proto_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_proto_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: PEER2PEER.Message
}
var file_proto_proto_proto_depIdxs = []int32{
	0, // 0: PEER2PEER.testT.testA:input_type -> PEER2PEER.Message
	0, // 1: PEER2PEER.testT.testA:output_type -> PEER2PEER.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_proto_proto_init() }
func file_proto_proto_proto_init() {
	if File_proto_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_proto_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_proto_proto_goTypes,
		DependencyIndexes: file_proto_proto_proto_depIdxs,
		MessageInfos:      file_proto_proto_proto_msgTypes,
	}.Build()
	File_proto_proto_proto = out.File
	file_proto_proto_proto_rawDesc = nil
	file_proto_proto_proto_goTypes = nil
	file_proto_proto_proto_depIdxs = nil
}