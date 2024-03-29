// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: storezhang/transfer/file.proto

package transfer

import (
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

// 文件
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 文件名
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 基础路径
	Base string `protobuf:"bytes,4,opt,name=base,proto3" json:"base,omitempty"`
	// 类型
	Type Type `protobuf:"varint,5,opt,name=type,proto3,enum=storezhang.transfer.Type" json:"type,omitempty"`
	// 存储
	Storage *anypb.Any `protobuf:"bytes,10,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storezhang_transfer_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_storezhang_transfer_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_storezhang_transfer_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *File) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNKNOWN
}

func (x *File) GetStorage() *anypb.Any {
	if x != nil {
		return x.Storage
	}
	return nil
}

var File_storezhang_transfer_file_proto protoreflect.FileDescriptor

var file_storezhang_transfer_file_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x42, 0x45, 0x0a, 0x18, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7a, 0x68,
	0x61, 0x6e, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x7a, 0x68, 0x61, 0x6e, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x3b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storezhang_transfer_file_proto_rawDescOnce sync.Once
	file_storezhang_transfer_file_proto_rawDescData = file_storezhang_transfer_file_proto_rawDesc
)

func file_storezhang_transfer_file_proto_rawDescGZIP() []byte {
	file_storezhang_transfer_file_proto_rawDescOnce.Do(func() {
		file_storezhang_transfer_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_storezhang_transfer_file_proto_rawDescData)
	})
	return file_storezhang_transfer_file_proto_rawDescData
}

var file_storezhang_transfer_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storezhang_transfer_file_proto_goTypes = []interface{}{
	(*File)(nil),      // 0: storezhang.transfer.File
	(Type)(0),         // 1: storezhang.transfer.Type
	(*anypb.Any)(nil), // 2: google.protobuf.Any
}
var file_storezhang_transfer_file_proto_depIdxs = []int32{
	1, // 0: storezhang.transfer.File.type:type_name -> storezhang.transfer.Type
	2, // 1: storezhang.transfer.File.storage:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_storezhang_transfer_file_proto_init() }
func file_storezhang_transfer_file_proto_init() {
	if File_storezhang_transfer_file_proto != nil {
		return
	}
	file_storezhang_transfer_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storezhang_transfer_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			RawDescriptor: file_storezhang_transfer_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storezhang_transfer_file_proto_goTypes,
		DependencyIndexes: file_storezhang_transfer_file_proto_depIdxs,
		MessageInfos:      file_storezhang_transfer_file_proto_msgTypes,
	}.Build()
	File_storezhang_transfer_file_proto = out.File
	file_storezhang_transfer_file_proto_rawDesc = nil
	file_storezhang_transfer_file_proto_goTypes = nil
	file_storezhang_transfer_file_proto_depIdxs = nil
}
