// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.7.1
// source: translit.proto

// protoc --go_out=plugins=grpc:. *.proto

package translit

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

type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word string `protobuf:"bytes,1,opt,name=Word,proto3" json:"Word,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_translit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_translit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_translit_proto_rawDescGZIP(), []int{0}
}

func (x *Word) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

var File_translit_proto protoreflect.FileDescriptor

var file_translit_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x22, 0x1a, 0x0a, 0x04, 0x57, 0x6f,
	0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x57, 0x6f, 0x72, 0x64, 0x32, 0x3f, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x04, 0x45, 0x6e, 0x52,
	0x75, 0x12, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x2e, 0x57, 0x6f, 0x72,
	0x64, 0x1a, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x2e, 0x57, 0x6f, 0x72,
	0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6c, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_translit_proto_rawDescOnce sync.Once
	file_translit_proto_rawDescData = file_translit_proto_rawDesc
)

func file_translit_proto_rawDescGZIP() []byte {
	file_translit_proto_rawDescOnce.Do(func() {
		file_translit_proto_rawDescData = protoimpl.X.CompressGZIP(file_translit_proto_rawDescData)
	})
	return file_translit_proto_rawDescData
}

var file_translit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_translit_proto_goTypes = []interface{}{
	(*Word)(nil), // 0: translit.Word
}
var file_translit_proto_depIdxs = []int32{
	0, // 0: translit.Transliteration.EnRu:input_type -> translit.Word
	0, // 1: translit.Transliteration.EnRu:output_type -> translit.Word
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_translit_proto_init() }
func file_translit_proto_init() {
	if File_translit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_translit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
			RawDescriptor: file_translit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_translit_proto_goTypes,
		DependencyIndexes: file_translit_proto_depIdxs,
		MessageInfos:      file_translit_proto_msgTypes,
	}.Build()
	File_translit_proto = out.File
	file_translit_proto_rawDesc = nil
	file_translit_proto_goTypes = nil
	file_translit_proto_depIdxs = nil
}
