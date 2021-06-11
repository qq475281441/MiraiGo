// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: oidb0xb77.proto

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

type DB77ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId       uint64           `protobuf:"varint,1,opt,name=appId,proto3" json:"appId,omitempty"`
	AppType     uint32           `protobuf:"varint,2,opt,name=appType,proto3" json:"appType,omitempty"`
	MsgStyle    uint32           `protobuf:"varint,3,opt,name=msgStyle,proto3" json:"msgStyle,omitempty"`
	SenderUin   uint64           `protobuf:"varint,4,opt,name=senderUin,proto3" json:"senderUin,omitempty"`
	ClientInfo  *DB77ClientInfo  `protobuf:"bytes,5,opt,name=clientInfo,proto3" json:"clientInfo,omitempty"`
	TextMsg     string           `protobuf:"bytes,6,opt,name=textMsg,proto3" json:"textMsg,omitempty"`
	ExtInfo     *DB77ExtInfo     `protobuf:"bytes,7,opt,name=extInfo,proto3" json:"extInfo,omitempty"`
	SendType    uint32           `protobuf:"varint,10,opt,name=sendType,proto3" json:"sendType,omitempty"`
	RecvUin     uint64           `protobuf:"varint,11,opt,name=recvUin,proto3" json:"recvUin,omitempty"`
	RichMsgBody *DB77RichMsgBody `protobuf:"bytes,12,opt,name=richMsgBody,proto3" json:"richMsgBody,omitempty"`
}

func (x *DB77ReqBody) Reset() {
	*x = DB77ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xb77_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DB77ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DB77ReqBody) ProtoMessage() {}

func (x *DB77ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xb77_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DB77ReqBody.ProtoReflect.Descriptor instead.
func (*DB77ReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xb77_proto_rawDescGZIP(), []int{0}
}

func (x *DB77ReqBody) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DB77ReqBody) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *DB77ReqBody) GetMsgStyle() uint32 {
	if x != nil {
		return x.MsgStyle
	}
	return 0
}

func (x *DB77ReqBody) GetSenderUin() uint64 {
	if x != nil {
		return x.SenderUin
	}
	return 0
}

func (x *DB77ReqBody) GetClientInfo() *DB77ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *DB77ReqBody) GetTextMsg() string {
	if x != nil {
		return x.TextMsg
	}
	return ""
}

func (x *DB77ReqBody) GetExtInfo() *DB77ExtInfo {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *DB77ReqBody) GetSendType() uint32 {
	if x != nil {
		return x.SendType
	}
	return 0
}

func (x *DB77ReqBody) GetRecvUin() uint64 {
	if x != nil {
		return x.RecvUin
	}
	return 0
}

func (x *DB77ReqBody) GetRichMsgBody() *DB77RichMsgBody {
	if x != nil {
		return x.RichMsgBody
	}
	return nil
}

type DB77ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform           uint32 `protobuf:"varint,1,opt,name=platform,proto3" json:"platform,omitempty"`
	SdkVersion         string `protobuf:"bytes,2,opt,name=sdkVersion,proto3" json:"sdkVersion,omitempty"`
	AndroidPackageName string `protobuf:"bytes,3,opt,name=androidPackageName,proto3" json:"androidPackageName,omitempty"`
	AndroidSignature   string `protobuf:"bytes,4,opt,name=androidSignature,proto3" json:"androidSignature,omitempty"`
	IosBundleId        string `protobuf:"bytes,5,opt,name=iosBundleId,proto3" json:"iosBundleId,omitempty"`
	PcSign             string `protobuf:"bytes,6,opt,name=pcSign,proto3" json:"pcSign,omitempty"`
}

func (x *DB77ClientInfo) Reset() {
	*x = DB77ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xb77_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DB77ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DB77ClientInfo) ProtoMessage() {}

func (x *DB77ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xb77_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DB77ClientInfo.ProtoReflect.Descriptor instead.
func (*DB77ClientInfo) Descriptor() ([]byte, []int) {
	return file_oidb0xb77_proto_rawDescGZIP(), []int{1}
}

func (x *DB77ClientInfo) GetPlatform() uint32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *DB77ClientInfo) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *DB77ClientInfo) GetAndroidPackageName() string {
	if x != nil {
		return x.AndroidPackageName
	}
	return ""
}

func (x *DB77ClientInfo) GetAndroidSignature() string {
	if x != nil {
		return x.AndroidSignature
	}
	return ""
}

func (x *DB77ClientInfo) GetIosBundleId() string {
	if x != nil {
		return x.IosBundleId
	}
	return ""
}

func (x *DB77ClientInfo) GetPcSign() string {
	if x != nil {
		return x.PcSign
	}
	return ""
}

type DB77ExtInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomFeatureId []uint32 `protobuf:"varint,11,rep,packed,name=customFeatureId,proto3" json:"customFeatureId,omitempty"`
	ApnsWording     string   `protobuf:"bytes,12,opt,name=apnsWording,proto3" json:"apnsWording,omitempty"`
	GroupSaveDbFlag uint32   `protobuf:"varint,13,opt,name=groupSaveDbFlag,proto3" json:"groupSaveDbFlag,omitempty"`
	ReceiverAppId   uint32   `protobuf:"varint,14,opt,name=receiverAppId,proto3" json:"receiverAppId,omitempty"`
	MsgSeq          uint64   `protobuf:"varint,15,opt,name=msgSeq,proto3" json:"msgSeq,omitempty"`
}

func (x *DB77ExtInfo) Reset() {
	*x = DB77ExtInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xb77_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DB77ExtInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DB77ExtInfo) ProtoMessage() {}

func (x *DB77ExtInfo) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xb77_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DB77ExtInfo.ProtoReflect.Descriptor instead.
func (*DB77ExtInfo) Descriptor() ([]byte, []int) {
	return file_oidb0xb77_proto_rawDescGZIP(), []int{2}
}

func (x *DB77ExtInfo) GetCustomFeatureId() []uint32 {
	if x != nil {
		return x.CustomFeatureId
	}
	return nil
}

func (x *DB77ExtInfo) GetApnsWording() string {
	if x != nil {
		return x.ApnsWording
	}
	return ""
}

func (x *DB77ExtInfo) GetGroupSaveDbFlag() uint32 {
	if x != nil {
		return x.GroupSaveDbFlag
	}
	return 0
}

func (x *DB77ExtInfo) GetReceiverAppId() uint32 {
	if x != nil {
		return x.ReceiverAppId
	}
	return 0
}

func (x *DB77ExtInfo) GetMsgSeq() uint64 {
	if x != nil {
		return x.MsgSeq
	}
	return 0
}

type DB77RichMsgBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	Summary    string `protobuf:"bytes,11,opt,name=summary,proto3" json:"summary,omitempty"`
	Brief      string `protobuf:"bytes,12,opt,name=brief,proto3" json:"brief,omitempty"`
	Url        string `protobuf:"bytes,13,opt,name=url,proto3" json:"url,omitempty"`
	PictureUrl string `protobuf:"bytes,14,opt,name=pictureUrl,proto3" json:"pictureUrl,omitempty"`
	Action     string `protobuf:"bytes,15,opt,name=action,proto3" json:"action,omitempty"`
	MusicUrl   string `protobuf:"bytes,16,opt,name=musicUrl,proto3" json:"musicUrl,omitempty"` //ImageInfo imageInfo = 17;
}

func (x *DB77RichMsgBody) Reset() {
	*x = DB77RichMsgBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xb77_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DB77RichMsgBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DB77RichMsgBody) ProtoMessage() {}

func (x *DB77RichMsgBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xb77_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DB77RichMsgBody.ProtoReflect.Descriptor instead.
func (*DB77RichMsgBody) Descriptor() ([]byte, []int) {
	return file_oidb0xb77_proto_rawDescGZIP(), []int{3}
}

func (x *DB77RichMsgBody) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DB77RichMsgBody) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *DB77RichMsgBody) GetBrief() string {
	if x != nil {
		return x.Brief
	}
	return ""
}

func (x *DB77RichMsgBody) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DB77RichMsgBody) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

func (x *DB77RichMsgBody) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *DB77RichMsgBody) GetMusicUrl() string {
	if x != nil {
		return x.MusicUrl
	}
	return ""
}

var File_oidb0xb77_proto protoreflect.FileDescriptor

var file_oidb0xb77_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x62, 0x37, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x0b, 0x44, 0x42, 0x37, 0x37, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x0a, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x44, 0x42, 0x37, 0x37, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x74, 0x65, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x65, 0x78, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x42, 0x37, 0x37, 0x45, 0x78,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x76, 0x55, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x76, 0x55, 0x69, 0x6e, 0x12, 0x32, 0x0a, 0x0b, 0x72, 0x69, 0x63, 0x68, 0x4d, 0x73, 0x67, 0x42,
	0x6f, 0x64, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x44, 0x42, 0x37, 0x37,
	0x52, 0x69, 0x63, 0x68, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x0b, 0x72, 0x69, 0x63,
	0x68, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x22, 0xe2, 0x01, 0x0a, 0x0e, 0x44, 0x42, 0x37,
	0x37, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6f, 0x73, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6f, 0x73, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x22, 0xc1, 0x01,
	0x0a, 0x0b, 0x44, 0x42, 0x37, 0x37, 0x45, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x6e, 0x73, 0x57,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70,
	0x6e, 0x73, 0x57, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x61, 0x76, 0x65, 0x44, 0x62, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x61, 0x76, 0x65, 0x44, 0x62, 0x46,
	0x6c, 0x61, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x73, 0x67,
	0x53, 0x65, 0x71, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x53, 0x65,
	0x71, 0x22, 0xbd, 0x01, 0x0a, 0x0f, 0x44, 0x42, 0x37, 0x37, 0x52, 0x69, 0x63, 0x68, 0x4d, 0x73,
	0x67, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x69, 0x65, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x55, 0x72,
	0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x55, 0x72,
	0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x6f, 0x69, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oidb0xb77_proto_rawDescOnce sync.Once
	file_oidb0xb77_proto_rawDescData = file_oidb0xb77_proto_rawDesc
)

func file_oidb0xb77_proto_rawDescGZIP() []byte {
	file_oidb0xb77_proto_rawDescOnce.Do(func() {
		file_oidb0xb77_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0xb77_proto_rawDescData)
	})
	return file_oidb0xb77_proto_rawDescData
}

var file_oidb0xb77_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_oidb0xb77_proto_goTypes = []interface{}{
	(*DB77ReqBody)(nil),     // 0: DB77ReqBody
	(*DB77ClientInfo)(nil),  // 1: DB77ClientInfo
	(*DB77ExtInfo)(nil),     // 2: DB77ExtInfo
	(*DB77RichMsgBody)(nil), // 3: DB77RichMsgBody
}
var file_oidb0xb77_proto_depIdxs = []int32{
	1, // 0: DB77ReqBody.clientInfo:type_name -> DB77ClientInfo
	2, // 1: DB77ReqBody.extInfo:type_name -> DB77ExtInfo
	3, // 2: DB77ReqBody.richMsgBody:type_name -> DB77RichMsgBody
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_oidb0xb77_proto_init() }
func file_oidb0xb77_proto_init() {
	if File_oidb0xb77_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0xb77_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DB77ReqBody); i {
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
		file_oidb0xb77_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DB77ClientInfo); i {
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
		file_oidb0xb77_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DB77ExtInfo); i {
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
		file_oidb0xb77_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DB77RichMsgBody); i {
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
			RawDescriptor: file_oidb0xb77_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0xb77_proto_goTypes,
		DependencyIndexes: file_oidb0xb77_proto_depIdxs,
		MessageInfos:      file_oidb0xb77_proto_msgTypes,
	}.Build()
	File_oidb0xb77_proto = out.File
	file_oidb0xb77_proto_rawDesc = nil
	file_oidb0xb77_proto_goTypes = nil
	file_oidb0xb77_proto_depIdxs = nil
}
