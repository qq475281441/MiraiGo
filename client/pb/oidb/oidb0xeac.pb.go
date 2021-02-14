// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: oidb0xeac.proto

package oidb

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EACReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupCode *uint64 `protobuf:"varint,1,opt,name=groupCode" json:"groupCode,omitempty"`
	Seq       *uint32 `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
	Random    *uint32 `protobuf:"varint,3,opt,name=random" json:"random,omitempty"`
}

func (x *EACReqBody) Reset() {
	*x = EACReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xeac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EACReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EACReqBody) ProtoMessage() {}

func (x *EACReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xeac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EACReqBody.ProtoReflect.Descriptor instead.
func (*EACReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xeac_proto_rawDescGZIP(), []int{0}
}

func (x *EACReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *EACReqBody) GetSeq() uint32 {
	if x != nil && x.Seq != nil {
		return *x.Seq
	}
	return 0
}

func (x *EACReqBody) GetRandom() uint32 {
	if x != nil && x.Random != nil {
		return *x.Random
	}
	return 0
}

type EACRspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wording    *string `protobuf:"bytes,1,opt,name=wording" json:"wording,omitempty"`
	DigestUin  *uint64 `protobuf:"varint,2,opt,name=digestUin" json:"digestUin,omitempty"`
	DigestTime *uint32 `protobuf:"varint,3,opt,name=digestTime" json:"digestTime,omitempty"`
	//optional DigestMsg msg = 4;
	ErrorCode *uint32 `protobuf:"varint,10,opt,name=errorCode" json:"errorCode,omitempty"`
}

func (x *EACRspBody) Reset() {
	*x = EACRspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xeac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EACRspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EACRspBody) ProtoMessage() {}

func (x *EACRspBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xeac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EACRspBody.ProtoReflect.Descriptor instead.
func (*EACRspBody) Descriptor() ([]byte, []int) {
	return file_oidb0xeac_proto_rawDescGZIP(), []int{1}
}

func (x *EACRspBody) GetWording() string {
	if x != nil && x.Wording != nil {
		return *x.Wording
	}
	return ""
}

func (x *EACRspBody) GetDigestUin() uint64 {
	if x != nil && x.DigestUin != nil {
		return *x.DigestUin
	}
	return 0
}

func (x *EACRspBody) GetDigestTime() uint32 {
	if x != nil && x.DigestTime != nil {
		return *x.DigestTime
	}
	return 0
}

func (x *EACRspBody) GetErrorCode() uint32 {
	if x != nil && x.ErrorCode != nil {
		return *x.ErrorCode
	}
	return 0
}

var File_oidb0xeac_proto protoreflect.FileDescriptor

var file_oidb0xeac_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x65, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x0a, 0x45, 0x41, 0x43, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x45, 0x41, 0x43, 0x52,
	0x73, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x3b, 0x6f, 0x69, 0x64, 0x62,
}

var (
	file_oidb0xeac_proto_rawDescOnce sync.Once
	file_oidb0xeac_proto_rawDescData = file_oidb0xeac_proto_rawDesc
)

func file_oidb0xeac_proto_rawDescGZIP() []byte {
	file_oidb0xeac_proto_rawDescOnce.Do(func() {
		file_oidb0xeac_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0xeac_proto_rawDescData)
	})
	return file_oidb0xeac_proto_rawDescData
}

var file_oidb0xeac_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oidb0xeac_proto_goTypes = []interface{}{
	(*EACReqBody)(nil), // 0: EACReqBody
	(*EACRspBody)(nil), // 1: EACRspBody
}
var file_oidb0xeac_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oidb0xeac_proto_init() }
func file_oidb0xeac_proto_init() {
	if File_oidb0xeac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0xeac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EACReqBody); i {
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
		file_oidb0xeac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EACRspBody); i {
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
			RawDescriptor: file_oidb0xeac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0xeac_proto_goTypes,
		DependencyIndexes: file_oidb0xeac_proto_depIdxs,
		MessageInfos:      file_oidb0xeac_proto_msgTypes,
	}.Build()
	File_oidb0xeac_proto = out.File
	file_oidb0xeac_proto_rawDesc = nil
	file_oidb0xeac_proto_goTypes = nil
	file_oidb0xeac_proto_depIdxs = nil
}
