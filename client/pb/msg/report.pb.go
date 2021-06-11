// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: report.proto

package msg

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

type PbMsgReadedReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrpReadReport []*PbGroupReadedReportReq   `protobuf:"bytes,1,rep,name=grpReadReport" json:"grpReadReport,omitempty"`
	DisReadReport []*PbDiscussReadedReportReq `protobuf:"bytes,2,rep,name=disReadReport" json:"disReadReport,omitempty"`
	C2CReadReport *PbC2CReadedReportReq       `protobuf:"bytes,3,opt,name=c2CReadReport" json:"c2CReadReport,omitempty"` //optional PbBindUinMsgReadedConfirmReq bindUinReadReport = 4;
}

func (x *PbMsgReadedReportReq) Reset() {
	*x = PbMsgReadedReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbMsgReadedReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbMsgReadedReportReq) ProtoMessage() {}

func (x *PbMsgReadedReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbMsgReadedReportReq.ProtoReflect.Descriptor instead.
func (*PbMsgReadedReportReq) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{0}
}

func (x *PbMsgReadedReportReq) GetGrpReadReport() []*PbGroupReadedReportReq {
	if x != nil {
		return x.GrpReadReport
	}
	return nil
}

func (x *PbMsgReadedReportReq) GetDisReadReport() []*PbDiscussReadedReportReq {
	if x != nil {
		return x.DisReadReport
	}
	return nil
}

func (x *PbMsgReadedReportReq) GetC2CReadReport() *PbC2CReadedReportReq {
	if x != nil {
		return x.C2CReadReport
	}
	return nil
}

type PbMsgReadedReportResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrpReadReport []*PbGroupReadedReportResp   `protobuf:"bytes,1,rep,name=grpReadReport" json:"grpReadReport,omitempty"`
	DisReadReport []*PbDiscussReadedReportResp `protobuf:"bytes,2,rep,name=disReadReport" json:"disReadReport,omitempty"`
	C2CReadReport *PbC2CReadedReportResp       `protobuf:"bytes,3,opt,name=c2CReadReport" json:"c2CReadReport,omitempty"` //optional PbBindUinMsgReadedConfirmResp bindUinReadReport = 4;
}

func (x *PbMsgReadedReportResp) Reset() {
	*x = PbMsgReadedReportResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbMsgReadedReportResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbMsgReadedReportResp) ProtoMessage() {}

func (x *PbMsgReadedReportResp) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbMsgReadedReportResp.ProtoReflect.Descriptor instead.
func (*PbMsgReadedReportResp) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{1}
}

func (x *PbMsgReadedReportResp) GetGrpReadReport() []*PbGroupReadedReportResp {
	if x != nil {
		return x.GrpReadReport
	}
	return nil
}

func (x *PbMsgReadedReportResp) GetDisReadReport() []*PbDiscussReadedReportResp {
	if x != nil {
		return x.DisReadReport
	}
	return nil
}

func (x *PbMsgReadedReportResp) GetC2CReadReport() *PbC2CReadedReportResp {
	if x != nil {
		return x.C2CReadReport
	}
	return nil
}

type PbGroupReadedReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupCode   *uint64 `protobuf:"varint,1,opt,name=groupCode" json:"groupCode,omitempty"`
	LastReadSeq *uint64 `protobuf:"varint,2,opt,name=lastReadSeq" json:"lastReadSeq,omitempty"`
}

func (x *PbGroupReadedReportReq) Reset() {
	*x = PbGroupReadedReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbGroupReadedReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbGroupReadedReportReq) ProtoMessage() {}

func (x *PbGroupReadedReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbGroupReadedReportReq.ProtoReflect.Descriptor instead.
func (*PbGroupReadedReportReq) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{2}
}

func (x *PbGroupReadedReportReq) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *PbGroupReadedReportReq) GetLastReadSeq() uint64 {
	if x != nil && x.LastReadSeq != nil {
		return *x.LastReadSeq
	}
	return 0
}

type PbDiscussReadedReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfUin     *uint64 `protobuf:"varint,1,opt,name=confUin" json:"confUin,omitempty"`
	LastReadSeq *uint64 `protobuf:"varint,2,opt,name=lastReadSeq" json:"lastReadSeq,omitempty"`
}

func (x *PbDiscussReadedReportReq) Reset() {
	*x = PbDiscussReadedReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbDiscussReadedReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbDiscussReadedReportReq) ProtoMessage() {}

func (x *PbDiscussReadedReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbDiscussReadedReportReq.ProtoReflect.Descriptor instead.
func (*PbDiscussReadedReportReq) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{3}
}

func (x *PbDiscussReadedReportReq) GetConfUin() uint64 {
	if x != nil && x.ConfUin != nil {
		return *x.ConfUin
	}
	return 0
}

func (x *PbDiscussReadedReportReq) GetLastReadSeq() uint64 {
	if x != nil && x.LastReadSeq != nil {
		return *x.LastReadSeq
	}
	return 0
}

type PbC2CReadedReportReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncCookie []byte             `protobuf:"bytes,1,opt,name=syncCookie" json:"syncCookie,omitempty"`
	PairInfo   []*UinPairReadInfo `protobuf:"bytes,2,rep,name=pairInfo" json:"pairInfo,omitempty"`
}

func (x *PbC2CReadedReportReq) Reset() {
	*x = PbC2CReadedReportReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbC2CReadedReportReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbC2CReadedReportReq) ProtoMessage() {}

func (x *PbC2CReadedReportReq) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbC2CReadedReportReq.ProtoReflect.Descriptor instead.
func (*PbC2CReadedReportReq) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{4}
}

func (x *PbC2CReadedReportReq) GetSyncCookie() []byte {
	if x != nil {
		return x.SyncCookie
	}
	return nil
}

func (x *PbC2CReadedReportReq) GetPairInfo() []*UinPairReadInfo {
	if x != nil {
		return x.PairInfo
	}
	return nil
}

type UinPairReadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerUin      *uint64 `protobuf:"varint,1,opt,name=peerUin" json:"peerUin,omitempty"`
	LastReadTime *uint32 `protobuf:"varint,2,opt,name=lastReadTime" json:"lastReadTime,omitempty"`
	CrmSig       []byte  `protobuf:"bytes,3,opt,name=crmSig" json:"crmSig,omitempty"`
	PeerType     *uint32 `protobuf:"varint,4,opt,name=peerType" json:"peerType,omitempty"`
	ChatType     *uint32 `protobuf:"varint,5,opt,name=chatType" json:"chatType,omitempty"`
	Cpid         *uint64 `protobuf:"varint,6,opt,name=cpid" json:"cpid,omitempty"`
	AioType      *uint32 `protobuf:"varint,7,opt,name=aioType" json:"aioType,omitempty"`
	ToTinyId     *uint64 `protobuf:"varint,9,opt,name=toTinyId" json:"toTinyId,omitempty"`
}

func (x *UinPairReadInfo) Reset() {
	*x = UinPairReadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UinPairReadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UinPairReadInfo) ProtoMessage() {}

func (x *UinPairReadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UinPairReadInfo.ProtoReflect.Descriptor instead.
func (*UinPairReadInfo) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{5}
}

func (x *UinPairReadInfo) GetPeerUin() uint64 {
	if x != nil && x.PeerUin != nil {
		return *x.PeerUin
	}
	return 0
}

func (x *UinPairReadInfo) GetLastReadTime() uint32 {
	if x != nil && x.LastReadTime != nil {
		return *x.LastReadTime
	}
	return 0
}

func (x *UinPairReadInfo) GetCrmSig() []byte {
	if x != nil {
		return x.CrmSig
	}
	return nil
}

func (x *UinPairReadInfo) GetPeerType() uint32 {
	if x != nil && x.PeerType != nil {
		return *x.PeerType
	}
	return 0
}

func (x *UinPairReadInfo) GetChatType() uint32 {
	if x != nil && x.ChatType != nil {
		return *x.ChatType
	}
	return 0
}

func (x *UinPairReadInfo) GetCpid() uint64 {
	if x != nil && x.Cpid != nil {
		return *x.Cpid
	}
	return 0
}

func (x *UinPairReadInfo) GetAioType() uint32 {
	if x != nil && x.AioType != nil {
		return *x.AioType
	}
	return 0
}

func (x *UinPairReadInfo) GetToTinyId() uint64 {
	if x != nil && x.ToTinyId != nil {
		return *x.ToTinyId
	}
	return 0
}

type PbGroupReadedReportResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result      *uint32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Errmsg      *string `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	GroupCode   *uint64 `protobuf:"varint,3,opt,name=groupCode" json:"groupCode,omitempty"`
	MemberSeq   *uint64 `protobuf:"varint,4,opt,name=memberSeq" json:"memberSeq,omitempty"`
	GroupMsgSeq *uint64 `protobuf:"varint,5,opt,name=groupMsgSeq" json:"groupMsgSeq,omitempty"`
}

func (x *PbGroupReadedReportResp) Reset() {
	*x = PbGroupReadedReportResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbGroupReadedReportResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbGroupReadedReportResp) ProtoMessage() {}

func (x *PbGroupReadedReportResp) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbGroupReadedReportResp.ProtoReflect.Descriptor instead.
func (*PbGroupReadedReportResp) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{6}
}

func (x *PbGroupReadedReportResp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *PbGroupReadedReportResp) GetErrmsg() string {
	if x != nil && x.Errmsg != nil {
		return *x.Errmsg
	}
	return ""
}

func (x *PbGroupReadedReportResp) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *PbGroupReadedReportResp) GetMemberSeq() uint64 {
	if x != nil && x.MemberSeq != nil {
		return *x.MemberSeq
	}
	return 0
}

func (x *PbGroupReadedReportResp) GetGroupMsgSeq() uint64 {
	if x != nil && x.GroupMsgSeq != nil {
		return *x.GroupMsgSeq
	}
	return 0
}

type PbDiscussReadedReportResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    *uint32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Errmsg    *string `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	ConfUin   *uint64 `protobuf:"varint,3,opt,name=confUin" json:"confUin,omitempty"`
	MemberSeq *uint64 `protobuf:"varint,4,opt,name=memberSeq" json:"memberSeq,omitempty"`
	ConfSeq   *uint64 `protobuf:"varint,5,opt,name=confSeq" json:"confSeq,omitempty"`
}

func (x *PbDiscussReadedReportResp) Reset() {
	*x = PbDiscussReadedReportResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbDiscussReadedReportResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbDiscussReadedReportResp) ProtoMessage() {}

func (x *PbDiscussReadedReportResp) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbDiscussReadedReportResp.ProtoReflect.Descriptor instead.
func (*PbDiscussReadedReportResp) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{7}
}

func (x *PbDiscussReadedReportResp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *PbDiscussReadedReportResp) GetErrmsg() string {
	if x != nil && x.Errmsg != nil {
		return *x.Errmsg
	}
	return ""
}

func (x *PbDiscussReadedReportResp) GetConfUin() uint64 {
	if x != nil && x.ConfUin != nil {
		return *x.ConfUin
	}
	return 0
}

func (x *PbDiscussReadedReportResp) GetMemberSeq() uint64 {
	if x != nil && x.MemberSeq != nil {
		return *x.MemberSeq
	}
	return 0
}

func (x *PbDiscussReadedReportResp) GetConfSeq() uint64 {
	if x != nil && x.ConfSeq != nil {
		return *x.ConfSeq
	}
	return 0
}

type PbC2CReadedReportResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     *uint32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Errmsg     *string `protobuf:"bytes,2,opt,name=errmsg" json:"errmsg,omitempty"`
	SyncCookie []byte  `protobuf:"bytes,3,opt,name=syncCookie" json:"syncCookie,omitempty"`
}

func (x *PbC2CReadedReportResp) Reset() {
	*x = PbC2CReadedReportResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbC2CReadedReportResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbC2CReadedReportResp) ProtoMessage() {}

func (x *PbC2CReadedReportResp) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbC2CReadedReportResp.ProtoReflect.Descriptor instead.
func (*PbC2CReadedReportResp) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{8}
}

func (x *PbC2CReadedReportResp) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *PbC2CReadedReportResp) GetErrmsg() string {
	if x != nil && x.Errmsg != nil {
		return *x.Errmsg
	}
	return ""
}

func (x *PbC2CReadedReportResp) GetSyncCookie() []byte {
	if x != nil {
		return x.SyncCookie
	}
	return nil
}

var File_report_proto protoreflect.FileDescriptor

var file_report_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3,
	0x01, 0x0a, 0x14, 0x50, 0x62, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3d, 0x0a, 0x0d, 0x67, 0x72, 0x70, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x50, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x52, 0x0d, 0x67, 0x72, 0x70, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x50, 0x62, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x63, 0x32, 0x43, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x50, 0x62, 0x43, 0x32, 0x43, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x52, 0x0d, 0x63, 0x32, 0x43, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x15, 0x50, 0x62, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3e,
	0x0a, 0x0d, 0x67, 0x72, 0x70, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x50, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x0d, 0x67, 0x72, 0x70, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x40,
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x50, 0x62, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x3c, 0x0a, 0x0d, 0x63, 0x32, 0x43, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x50, 0x62, 0x43, 0x32, 0x43, 0x52,
	0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x0d, 0x63, 0x32, 0x43, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x58,
	0x0a, 0x16, 0x50, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x71, 0x22, 0x56, 0x0a, 0x18, 0x50, 0x62, 0x44, 0x69,
	0x73, 0x63, 0x75, 0x73, 0x73, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x55, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x55, 0x69, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x71,
	0x22, 0x64, 0x0a, 0x14, 0x50, 0x62, 0x43, 0x32, 0x43, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63,
	0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x79,
	0x6e, 0x63, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x69, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x55, 0x69, 0x6e,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61,
	0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x0f, 0x55, 0x69, 0x6e, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65,
	0x65, 0x72, 0x55, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x65, 0x65,
	0x72, 0x55, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x6d, 0x53,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x72, 0x6d, 0x53, 0x69, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x65, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x70, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x69, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61,
	0x69, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x54, 0x69, 0x6e, 0x79,
	0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x6f, 0x54, 0x69, 0x6e, 0x79,
	0x49, 0x64, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x50, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x1c,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x71, 0x22, 0x9d, 0x01, 0x0a,
	0x19, 0x50, 0x62, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x66, 0x55, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x66, 0x55, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65,
	0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x53, 0x65, 0x71, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x53, 0x65, 0x71, 0x22, 0x67, 0x0a, 0x15,
	0x50, 0x62, 0x43, 0x32, 0x43, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6f,
	0x6b, 0x69, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x6d, 0x73, 0x67,
}

var (
	file_report_proto_rawDescOnce sync.Once
	file_report_proto_rawDescData = file_report_proto_rawDesc
)

func file_report_proto_rawDescGZIP() []byte {
	file_report_proto_rawDescOnce.Do(func() {
		file_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_proto_rawDescData)
	})
	return file_report_proto_rawDescData
}

var file_report_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_report_proto_goTypes = []interface{}{
	(*PbMsgReadedReportReq)(nil),      // 0: PbMsgReadedReportReq
	(*PbMsgReadedReportResp)(nil),     // 1: PbMsgReadedReportResp
	(*PbGroupReadedReportReq)(nil),    // 2: PbGroupReadedReportReq
	(*PbDiscussReadedReportReq)(nil),  // 3: PbDiscussReadedReportReq
	(*PbC2CReadedReportReq)(nil),      // 4: PbC2CReadedReportReq
	(*UinPairReadInfo)(nil),           // 5: UinPairReadInfo
	(*PbGroupReadedReportResp)(nil),   // 6: PbGroupReadedReportResp
	(*PbDiscussReadedReportResp)(nil), // 7: PbDiscussReadedReportResp
	(*PbC2CReadedReportResp)(nil),     // 8: PbC2CReadedReportResp
}
var file_report_proto_depIdxs = []int32{
	2, // 0: PbMsgReadedReportReq.grpReadReport:type_name -> PbGroupReadedReportReq
	3, // 1: PbMsgReadedReportReq.disReadReport:type_name -> PbDiscussReadedReportReq
	4, // 2: PbMsgReadedReportReq.c2CReadReport:type_name -> PbC2CReadedReportReq
	6, // 3: PbMsgReadedReportResp.grpReadReport:type_name -> PbGroupReadedReportResp
	7, // 4: PbMsgReadedReportResp.disReadReport:type_name -> PbDiscussReadedReportResp
	8, // 5: PbMsgReadedReportResp.c2CReadReport:type_name -> PbC2CReadedReportResp
	5, // 6: PbC2CReadedReportReq.pairInfo:type_name -> UinPairReadInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_report_proto_init() }
func file_report_proto_init() {
	if File_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbMsgReadedReportReq); i {
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
		file_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbMsgReadedReportResp); i {
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
		file_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbGroupReadedReportReq); i {
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
		file_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbDiscussReadedReportReq); i {
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
		file_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbC2CReadedReportReq); i {
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
		file_report_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UinPairReadInfo); i {
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
		file_report_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbGroupReadedReportResp); i {
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
		file_report_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbDiscussReadedReportResp); i {
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
		file_report_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbC2CReadedReportResp); i {
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
			RawDescriptor: file_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_report_proto_goTypes,
		DependencyIndexes: file_report_proto_depIdxs,
		MessageInfos:      file_report_proto_msgTypes,
	}.Build()
	File_report_proto = out.File
	file_report_proto_rawDesc = nil
	file_report_proto_goTypes = nil
	file_report_proto_depIdxs = nil
}
