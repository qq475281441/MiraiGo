// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: oidb0xD79.proto

package oidb

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

type D79ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq          uint64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Uin          uint64 `protobuf:"varint,2,opt,name=uin,proto3" json:"uin,omitempty"`
	CompressFlag uint32 `protobuf:"varint,3,opt,name=compress_flag,json=compressFlag,proto3" json:"compress_flag,omitempty"`
	Content      []byte `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	SenderUin    uint64 `protobuf:"varint,5,opt,name=sender_uin,json=senderUin,proto3" json:"sender_uin,omitempty"`
	Qua          []byte `protobuf:"bytes,6,opt,name=qua,proto3" json:"qua,omitempty"`
	WordExt      []byte `protobuf:"bytes,7,opt,name=word_ext,json=wordExt,proto3" json:"word_ext,omitempty"`
}

func (x *D79ReqBody) Reset() {
	*x = D79ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xD79_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D79ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D79ReqBody) ProtoMessage() {}

func (x *D79ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xD79_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D79ReqBody.ProtoReflect.Descriptor instead.
func (*D79ReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xD79_proto_rawDescGZIP(), []int{0}
}

func (x *D79ReqBody) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *D79ReqBody) GetUin() uint64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *D79ReqBody) GetCompressFlag() uint32 {
	if x != nil {
		return x.CompressFlag
	}
	return 0
}

func (x *D79ReqBody) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *D79ReqBody) GetSenderUin() uint64 {
	if x != nil {
		return x.SenderUin
	}
	return 0
}

func (x *D79ReqBody) GetQua() []byte {
	if x != nil {
		return x.Qua
	}
	return nil
}

func (x *D79ReqBody) GetWordExt() []byte {
	if x != nil {
		return x.WordExt
	}
	return nil
}

type D79RspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret          uint32      `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Seq          uint64      `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Uin          uint64      `protobuf:"varint,3,opt,name=uin,proto3" json:"uin,omitempty"`
	CompressFlag uint32      `protobuf:"varint,4,opt,name=compress_flag,json=compressFlag,proto3" json:"compress_flag,omitempty"`
	Content      *D79Content `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *D79RspBody) Reset() {
	*x = D79RspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xD79_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D79RspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D79RspBody) ProtoMessage() {}

func (x *D79RspBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xD79_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D79RspBody.ProtoReflect.Descriptor instead.
func (*D79RspBody) Descriptor() ([]byte, []int) {
	return file_oidb0xD79_proto_rawDescGZIP(), []int{1}
}

func (x *D79RspBody) GetRet() uint32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *D79RspBody) GetSeq() uint64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *D79RspBody) GetUin() uint64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *D79RspBody) GetCompressFlag() uint32 {
	if x != nil {
		return x.CompressFlag
	}
	return 0
}

func (x *D79RspBody) GetContent() *D79Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type D79Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SliceContent [][]byte `protobuf:"bytes,1,rep,name=slice_content,json=sliceContent,proto3" json:"slice_content,omitempty"`
}

func (x *D79Content) Reset() {
	*x = D79Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xD79_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D79Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D79Content) ProtoMessage() {}

func (x *D79Content) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xD79_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D79Content.ProtoReflect.Descriptor instead.
func (*D79Content) Descriptor() ([]byte, []int) {
	return file_oidb0xD79_proto_rawDescGZIP(), []int{2}
}

func (x *D79Content) GetSliceContent() [][]byte {
	if x != nil {
		return x.SliceContent
	}
	return nil
}

var File_oidb0xD79_proto protoreflect.FileDescriptor

var file_oidb0xD79_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x44, 0x37, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x0a, 0x44, 0x37, 0x39, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x69,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55,
	0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x75, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x71, 0x75, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x65, 0x78, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x78, 0x74, 0x22,
	0x8e, 0x01, 0x0a, 0x0a, 0x44, 0x37, 0x39, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x72, 0x65, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x44, 0x37, 0x39,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x31, 0x0a, 0x0a, 0x44, 0x37, 0x39, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x6f, 0x69, 0x64, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidb0xD79_proto_rawDescOnce sync.Once
	file_oidb0xD79_proto_rawDescData = file_oidb0xD79_proto_rawDesc
)

func file_oidb0xD79_proto_rawDescGZIP() []byte {
	file_oidb0xD79_proto_rawDescOnce.Do(func() {
		file_oidb0xD79_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0xD79_proto_rawDescData)
	})
	return file_oidb0xD79_proto_rawDescData
}

var file_oidb0xD79_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oidb0xD79_proto_goTypes = []interface{}{
	(*D79ReqBody)(nil), // 0: D79ReqBody
	(*D79RspBody)(nil), // 1: D79RspBody
	(*D79Content)(nil), // 2: D79Content
}
var file_oidb0xD79_proto_depIdxs = []int32{
	2, // 0: D79RspBody.content:type_name -> D79Content
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_oidb0xD79_proto_init() }
func file_oidb0xD79_proto_init() {
	if File_oidb0xD79_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0xD79_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D79ReqBody); i {
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
		file_oidb0xD79_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D79RspBody); i {
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
		file_oidb0xD79_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D79Content); i {
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
			RawDescriptor: file_oidb0xD79_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0xD79_proto_goTypes,
		DependencyIndexes: file_oidb0xD79_proto_depIdxs,
		MessageInfos:      file_oidb0xD79_proto_msgTypes,
	}.Build()
	File_oidb0xD79_proto = out.File
	file_oidb0xD79_proto_rawDesc = nil
	file_oidb0xD79_proto_goTypes = nil
	file_oidb0xD79_proto_depIdxs = nil
}
