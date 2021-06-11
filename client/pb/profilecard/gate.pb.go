// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: gate.proto

package profilecard

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GateCommTaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid    *int32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	TaskData []byte `protobuf:"bytes,2,opt,name=taskData" json:"taskData,omitempty"`
}

func (x *GateCommTaskInfo) Reset() {
	*x = GateCommTaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateCommTaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateCommTaskInfo) ProtoMessage() {}

func (x *GateCommTaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateCommTaskInfo.ProtoReflect.Descriptor instead.
func (*GateCommTaskInfo) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{0}
}

func (x *GateCommTaskInfo) GetAppid() int32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *GateCommTaskInfo) GetTaskData() []byte {
	if x != nil {
		return x.TaskData
	}
	return nil
}

type GateGetGiftListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin *int32 `protobuf:"varint,1,opt,name=uin" json:"uin,omitempty"`
}

func (x *GateGetGiftListReq) Reset() {
	*x = GateGetGiftListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateGetGiftListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateGetGiftListReq) ProtoMessage() {}

func (x *GateGetGiftListReq) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateGetGiftListReq.ProtoReflect.Descriptor instead.
func (*GateGetGiftListReq) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{1}
}

func (x *GateGetGiftListReq) GetUin() int32 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

type GateGetGiftListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GiftUrl   []string `protobuf:"bytes,1,rep,name=giftUrl" json:"giftUrl,omitempty"`
	CustomUrl *string  `protobuf:"bytes,2,opt,name=customUrl" json:"customUrl,omitempty"`
	Desc      *string  `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	IsOn      *bool    `protobuf:"varint,4,opt,name=isOn" json:"isOn,omitempty"`
}

func (x *GateGetGiftListRsp) Reset() {
	*x = GateGetGiftListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateGetGiftListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateGetGiftListRsp) ProtoMessage() {}

func (x *GateGetGiftListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateGetGiftListRsp.ProtoReflect.Descriptor instead.
func (*GateGetGiftListRsp) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{2}
}

func (x *GateGetGiftListRsp) GetGiftUrl() []string {
	if x != nil {
		return x.GiftUrl
	}
	return nil
}

func (x *GateGetGiftListRsp) GetCustomUrl() string {
	if x != nil && x.CustomUrl != nil {
		return *x.CustomUrl
	}
	return ""
}

func (x *GateGetGiftListRsp) GetDesc() string {
	if x != nil && x.Desc != nil {
		return *x.Desc
	}
	return ""
}

func (x *GateGetGiftListRsp) GetIsOn() bool {
	if x != nil && x.IsOn != nil {
		return *x.IsOn
	}
	return false
}

type GateGetVipCareReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin *int64 `protobuf:"varint,1,opt,name=uin" json:"uin,omitempty"`
}

func (x *GateGetVipCareReq) Reset() {
	*x = GateGetVipCareReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateGetVipCareReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateGetVipCareReq) ProtoMessage() {}

func (x *GateGetVipCareReq) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateGetVipCareReq.ProtoReflect.Descriptor instead.
func (*GateGetVipCareReq) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{3}
}

func (x *GateGetVipCareReq) GetUin() int64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

type GateGetVipCareRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buss   *int32 `protobuf:"varint,1,opt,name=buss" json:"buss,omitempty"`
	Notice *int32 `protobuf:"varint,2,opt,name=notice" json:"notice,omitempty"`
}

func (x *GateGetVipCareRsp) Reset() {
	*x = GateGetVipCareRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateGetVipCareRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateGetVipCareRsp) ProtoMessage() {}

func (x *GateGetVipCareRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateGetVipCareRsp.ProtoReflect.Descriptor instead.
func (*GateGetVipCareRsp) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{4}
}

func (x *GateGetVipCareRsp) GetBuss() int32 {
	if x != nil && x.Buss != nil {
		return *x.Buss
	}
	return 0
}

func (x *GateGetVipCareRsp) GetNotice() int32 {
	if x != nil && x.Notice != nil {
		return *x.Notice
	}
	return 0
}

type GateOidbFlagInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fieled     *int32 `protobuf:"varint,1,opt,name=fieled" json:"fieled,omitempty"`
	ByetsValue []byte `protobuf:"bytes,2,opt,name=byetsValue" json:"byetsValue,omitempty"`
}

func (x *GateOidbFlagInfo) Reset() {
	*x = GateOidbFlagInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateOidbFlagInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateOidbFlagInfo) ProtoMessage() {}

func (x *GateOidbFlagInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateOidbFlagInfo.ProtoReflect.Descriptor instead.
func (*GateOidbFlagInfo) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{5}
}

func (x *GateOidbFlagInfo) GetFieled() int32 {
	if x != nil && x.Fieled != nil {
		return *x.Fieled
	}
	return 0
}

func (x *GateOidbFlagInfo) GetByetsValue() []byte {
	if x != nil {
		return x.ByetsValue
	}
	return nil
}

type GatePrivilegeBaseInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UReqUin *int64 `protobuf:"varint,1,opt,name=uReqUin" json:"uReqUin,omitempty"`
}

func (x *GatePrivilegeBaseInfoReq) Reset() {
	*x = GatePrivilegeBaseInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatePrivilegeBaseInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatePrivilegeBaseInfoReq) ProtoMessage() {}

func (x *GatePrivilegeBaseInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatePrivilegeBaseInfoReq.ProtoReflect.Descriptor instead.
func (*GatePrivilegeBaseInfoReq) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{6}
}

func (x *GatePrivilegeBaseInfoReq) GetUReqUin() int64 {
	if x != nil && x.UReqUin != nil {
		return *x.UReqUin
	}
	return 0
}

type GatePrivilegeBaseInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg        []byte               `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	JumpUrl    []byte               `protobuf:"bytes,2,opt,name=jumpUrl" json:"jumpUrl,omitempty"`
	VOpenPriv  []*GatePrivilegeInfo `protobuf:"bytes,3,rep,name=vOpenPriv" json:"vOpenPriv,omitempty"`
	VClosePriv []*GatePrivilegeInfo `protobuf:"bytes,4,rep,name=vClosePriv" json:"vClosePriv,omitempty"`
	UIsGrayUsr *int32               `protobuf:"varint,5,opt,name=uIsGrayUsr" json:"uIsGrayUsr,omitempty"`
}

func (x *GatePrivilegeBaseInfoRsp) Reset() {
	*x = GatePrivilegeBaseInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatePrivilegeBaseInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatePrivilegeBaseInfoRsp) ProtoMessage() {}

func (x *GatePrivilegeBaseInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatePrivilegeBaseInfoRsp.ProtoReflect.Descriptor instead.
func (*GatePrivilegeBaseInfoRsp) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{7}
}

func (x *GatePrivilegeBaseInfoRsp) GetMsg() []byte {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *GatePrivilegeBaseInfoRsp) GetJumpUrl() []byte {
	if x != nil {
		return x.JumpUrl
	}
	return nil
}

func (x *GatePrivilegeBaseInfoRsp) GetVOpenPriv() []*GatePrivilegeInfo {
	if x != nil {
		return x.VOpenPriv
	}
	return nil
}

func (x *GatePrivilegeBaseInfoRsp) GetVClosePriv() []*GatePrivilegeInfo {
	if x != nil {
		return x.VClosePriv
	}
	return nil
}

func (x *GatePrivilegeBaseInfoRsp) GetUIsGrayUsr() int32 {
	if x != nil && x.UIsGrayUsr != nil {
		return *x.UIsGrayUsr
	}
	return 0
}

type GatePrivilegeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IType         *int32 `protobuf:"varint,1,opt,name=iType" json:"iType,omitempty"`
	ISort         *int32 `protobuf:"varint,2,opt,name=iSort" json:"iSort,omitempty"`
	IFeeType      *int32 `protobuf:"varint,3,opt,name=iFeeType" json:"iFeeType,omitempty"`
	ILevel        *int32 `protobuf:"varint,4,opt,name=iLevel" json:"iLevel,omitempty"`
	IFlag         *int32 `protobuf:"varint,5,opt,name=iFlag" json:"iFlag,omitempty"`
	IconUrl       []byte `protobuf:"bytes,6,opt,name=iconUrl" json:"iconUrl,omitempty"`
	DeluxeIconUrl []byte `protobuf:"bytes,7,opt,name=deluxeIconUrl" json:"deluxeIconUrl,omitempty"`
	JumpUrl       []byte `protobuf:"bytes,8,opt,name=jumpUrl" json:"jumpUrl,omitempty"`
	IIsBig        *int32 `protobuf:"varint,9,opt,name=iIsBig" json:"iIsBig,omitempty"`
}

func (x *GatePrivilegeInfo) Reset() {
	*x = GatePrivilegeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatePrivilegeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatePrivilegeInfo) ProtoMessage() {}

func (x *GatePrivilegeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatePrivilegeInfo.ProtoReflect.Descriptor instead.
func (*GatePrivilegeInfo) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{8}
}

func (x *GatePrivilegeInfo) GetIType() int32 {
	if x != nil && x.IType != nil {
		return *x.IType
	}
	return 0
}

func (x *GatePrivilegeInfo) GetISort() int32 {
	if x != nil && x.ISort != nil {
		return *x.ISort
	}
	return 0
}

func (x *GatePrivilegeInfo) GetIFeeType() int32 {
	if x != nil && x.IFeeType != nil {
		return *x.IFeeType
	}
	return 0
}

func (x *GatePrivilegeInfo) GetILevel() int32 {
	if x != nil && x.ILevel != nil {
		return *x.ILevel
	}
	return 0
}

func (x *GatePrivilegeInfo) GetIFlag() int32 {
	if x != nil && x.IFlag != nil {
		return *x.IFlag
	}
	return 0
}

func (x *GatePrivilegeInfo) GetIconUrl() []byte {
	if x != nil {
		return x.IconUrl
	}
	return nil
}

func (x *GatePrivilegeInfo) GetDeluxeIconUrl() []byte {
	if x != nil {
		return x.DeluxeIconUrl
	}
	return nil
}

func (x *GatePrivilegeInfo) GetJumpUrl() []byte {
	if x != nil {
		return x.JumpUrl
	}
	return nil
}

func (x *GatePrivilegeInfo) GetIIsBig() int32 {
	if x != nil && x.IIsBig != nil {
		return *x.IIsBig
	}
	return 0
}

type GateVaProfileGateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UCmd           *int32                    `protobuf:"varint,1,opt,name=uCmd" json:"uCmd,omitempty"`
	StPrivilegeReq *GatePrivilegeBaseInfoReq `protobuf:"bytes,2,opt,name=stPrivilegeReq" json:"stPrivilegeReq,omitempty"`
	StGiftReq      *GateGetGiftListReq       `protobuf:"bytes,3,opt,name=stGiftReq" json:"stGiftReq,omitempty"`
	TaskItem       []*GateCommTaskInfo       `protobuf:"bytes,4,rep,name=taskItem" json:"taskItem,omitempty"`
	OidbFlag       []*GateOidbFlagInfo       `protobuf:"bytes,5,rep,name=oidbFlag" json:"oidbFlag,omitempty"`
	StVipCare      *GateGetVipCareReq        `protobuf:"bytes,6,opt,name=stVipCare" json:"stVipCare,omitempty"`
}

func (x *GateVaProfileGateReq) Reset() {
	*x = GateVaProfileGateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateVaProfileGateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateVaProfileGateReq) ProtoMessage() {}

func (x *GateVaProfileGateReq) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateVaProfileGateReq.ProtoReflect.Descriptor instead.
func (*GateVaProfileGateReq) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{9}
}

func (x *GateVaProfileGateReq) GetUCmd() int32 {
	if x != nil && x.UCmd != nil {
		return *x.UCmd
	}
	return 0
}

func (x *GateVaProfileGateReq) GetStPrivilegeReq() *GatePrivilegeBaseInfoReq {
	if x != nil {
		return x.StPrivilegeReq
	}
	return nil
}

func (x *GateVaProfileGateReq) GetStGiftReq() *GateGetGiftListReq {
	if x != nil {
		return x.StGiftReq
	}
	return nil
}

func (x *GateVaProfileGateReq) GetTaskItem() []*GateCommTaskInfo {
	if x != nil {
		return x.TaskItem
	}
	return nil
}

func (x *GateVaProfileGateReq) GetOidbFlag() []*GateOidbFlagInfo {
	if x != nil {
		return x.OidbFlag
	}
	return nil
}

func (x *GateVaProfileGateReq) GetStVipCare() *GateGetVipCareReq {
	if x != nil {
		return x.StVipCare
	}
	return nil
}

type GateQidInfoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qid     *string `protobuf:"bytes,1,opt,name=qid" json:"qid,omitempty"`
	Url     *string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Color   *string `protobuf:"bytes,3,opt,name=color" json:"color,omitempty"`
	LogoUrl *string `protobuf:"bytes,4,opt,name=logoUrl" json:"logoUrl,omitempty"`
}

func (x *GateQidInfoItem) Reset() {
	*x = GateQidInfoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateQidInfoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateQidInfoItem) ProtoMessage() {}

func (x *GateQidInfoItem) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateQidInfoItem.ProtoReflect.Descriptor instead.
func (*GateQidInfoItem) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{10}
}

func (x *GateQidInfoItem) GetQid() string {
	if x != nil && x.Qid != nil {
		return *x.Qid
	}
	return ""
}

func (x *GateQidInfoItem) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *GateQidInfoItem) GetColor() string {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return ""
}

func (x *GateQidInfoItem) GetLogoUrl() string {
	if x != nil && x.LogoUrl != nil {
		return *x.LogoUrl
	}
	return ""
}

type GateVaProfileGateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IRetCode       *int32                    `protobuf:"varint,1,opt,name=iRetCode" json:"iRetCode,omitempty"`
	SRetMsg        []byte                    `protobuf:"bytes,2,opt,name=sRetMsg" json:"sRetMsg,omitempty"`
	StPrivilegeRsp *GatePrivilegeBaseInfoRsp `protobuf:"bytes,3,opt,name=stPrivilegeRsp" json:"stPrivilegeRsp,omitempty"`
	StGiftRsp      *GateGetGiftListRsp       `protobuf:"bytes,4,opt,name=stGiftRsp" json:"stGiftRsp,omitempty"`
	TaskItem       []*GateCommTaskInfo       `protobuf:"bytes,5,rep,name=taskItem" json:"taskItem,omitempty"`
	OidbFlag       []*GateOidbFlagInfo       `protobuf:"bytes,6,rep,name=oidbFlag" json:"oidbFlag,omitempty"`
	StVipCare      *GateGetVipCareRsp        `protobuf:"bytes,7,opt,name=stVipCare" json:"stVipCare,omitempty"`
	QidInfo        *GateQidInfoItem          `protobuf:"bytes,9,opt,name=qidInfo" json:"qidInfo,omitempty"`
}

func (x *GateVaProfileGateRsp) Reset() {
	*x = GateVaProfileGateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gate_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateVaProfileGateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateVaProfileGateRsp) ProtoMessage() {}

func (x *GateVaProfileGateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_gate_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateVaProfileGateRsp.ProtoReflect.Descriptor instead.
func (*GateVaProfileGateRsp) Descriptor() ([]byte, []int) {
	return file_gate_proto_rawDescGZIP(), []int{11}
}

func (x *GateVaProfileGateRsp) GetIRetCode() int32 {
	if x != nil && x.IRetCode != nil {
		return *x.IRetCode
	}
	return 0
}

func (x *GateVaProfileGateRsp) GetSRetMsg() []byte {
	if x != nil {
		return x.SRetMsg
	}
	return nil
}

func (x *GateVaProfileGateRsp) GetStPrivilegeRsp() *GatePrivilegeBaseInfoRsp {
	if x != nil {
		return x.StPrivilegeRsp
	}
	return nil
}

func (x *GateVaProfileGateRsp) GetStGiftRsp() *GateGetGiftListRsp {
	if x != nil {
		return x.StGiftRsp
	}
	return nil
}

func (x *GateVaProfileGateRsp) GetTaskItem() []*GateCommTaskInfo {
	if x != nil {
		return x.TaskItem
	}
	return nil
}

func (x *GateVaProfileGateRsp) GetOidbFlag() []*GateOidbFlagInfo {
	if x != nil {
		return x.OidbFlag
	}
	return nil
}

func (x *GateVaProfileGateRsp) GetStVipCare() *GateGetVipCareRsp {
	if x != nil {
		return x.StVipCare
	}
	return nil
}

func (x *GateVaProfileGateRsp) GetQidInfo() *GateQidInfoItem {
	if x != nil {
		return x.QidInfo
	}
	return nil
}

var File_gate_proto protoreflect.FileDescriptor

var file_gate_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10,
	0x47, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x26, 0x0a, 0x12, 0x47, 0x61, 0x74, 0x65, 0x47, 0x65, 0x74, 0x47, 0x69, 0x66,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x22, 0x74, 0x0a, 0x12, 0x47, 0x61,
	0x74, 0x65, 0x47, 0x65, 0x74, 0x47, 0x69, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x69, 0x66, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x69, 0x66, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x73, 0x4f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e,
	0x22, 0x25, 0x0a, 0x11, 0x47, 0x61, 0x74, 0x65, 0x47, 0x65, 0x74, 0x56, 0x69, 0x70, 0x43, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x61, 0x74, 0x65, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x70, 0x43, 0x61, 0x72, 0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x75, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x75, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x10, 0x47, 0x61, 0x74, 0x65,
	0x4f, 0x69, 0x64, 0x62, 0x46, 0x6c, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x79, 0x65, 0x74, 0x73, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x65, 0x74, 0x73, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x34, 0x0a, 0x18, 0x47, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x52, 0x65, 0x71, 0x55, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x75, 0x52, 0x65, 0x71, 0x55, 0x69, 0x6e, 0x22, 0xcc, 0x01, 0x0a, 0x18, 0x47,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x75, 0x6d,
	0x70, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6a, 0x75, 0x6d, 0x70,
	0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x09, 0x76, 0x4f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x76,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x76, 0x4f, 0x70, 0x65,
	0x6e, 0x50, 0x72, 0x69, 0x76, 0x12, 0x32, 0x0a, 0x0a, 0x76, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50,
	0x72, 0x69, 0x76, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x76,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x76, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x49, 0x73,
	0x47, 0x72, 0x61, 0x79, 0x55, 0x73, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75,
	0x49, 0x73, 0x47, 0x72, 0x61, 0x79, 0x55, 0x73, 0x72, 0x22, 0xfb, 0x01, 0x0a, 0x11, 0x47, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x69, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x46, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69,
	0x46, 0x65, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x69, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12,
	0x24, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x75, 0x78, 0x65, 0x49, 0x63, 0x6f, 0x6e, 0x55, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x75, 0x78, 0x65, 0x49, 0x63,
	0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6a, 0x75, 0x6d, 0x70, 0x55, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x49, 0x73, 0x42, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x69, 0x49, 0x73, 0x42, 0x69, 0x67, 0x22, 0xb0, 0x02, 0x0a, 0x14, 0x47, 0x61, 0x74, 0x65,
	0x56, 0x61, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x47, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x43, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x75, 0x43, 0x6d, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x52, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x52, 0x0e, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x74, 0x47, 0x69, 0x66,
	0x74, 0x52, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x47, 0x65, 0x74, 0x47, 0x69, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x52,
	0x09, 0x73, 0x74, 0x47, 0x69, 0x66, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2d, 0x0a, 0x08, 0x6f, 0x69, 0x64,
	0x62, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x4f, 0x69, 0x64, 0x62, 0x46, 0x6c, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x6f, 0x69, 0x64, 0x62, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x74, 0x56, 0x69,
	0x70, 0x43, 0x61, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x47, 0x65, 0x74, 0x56, 0x69, 0x70, 0x43, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x52,
	0x09, 0x73, 0x74, 0x56, 0x69, 0x70, 0x43, 0x61, 0x72, 0x65, 0x22, 0x65, 0x0a, 0x0f, 0x47, 0x61,
	0x74, 0x65, 0x51, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x71, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72,
	0x6c, 0x22, 0xfe, 0x02, 0x0a, 0x14, 0x47, 0x61, 0x74, 0x65, 0x56, 0x61, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x47, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x52,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x52,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x52, 0x65, 0x74, 0x4d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x52, 0x65, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x41, 0x0a, 0x0e, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52,
	0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x73, 0x70, 0x52, 0x0e, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x52, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x09, 0x73, 0x74, 0x47, 0x69, 0x66, 0x74, 0x52, 0x73, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x47, 0x65, 0x74,
	0x47, 0x69, 0x66, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x52, 0x09, 0x73, 0x74, 0x47,
	0x69, 0x66, 0x74, 0x52, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x74,
	0x65, 0x6d, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2d, 0x0a, 0x08, 0x6f, 0x69, 0x64, 0x62, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x4f, 0x69,
	0x64, 0x62, 0x46, 0x6c, 0x61, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6f, 0x69, 0x64, 0x62,
	0x46, 0x6c, 0x61, 0x67, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x74, 0x56, 0x69, 0x70, 0x43, 0x61, 0x72,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x70, 0x43, 0x61, 0x72, 0x65, 0x52, 0x73, 0x70, 0x52, 0x09, 0x73, 0x74, 0x56,
	0x69, 0x70, 0x43, 0x61, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x71, 0x69, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x51, 0x69,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x71, 0x69, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x63, 0x61, 0x72, 0x64,
}

var (
	file_gate_proto_rawDescOnce sync.Once
	file_gate_proto_rawDescData = file_gate_proto_rawDesc
)

func file_gate_proto_rawDescGZIP() []byte {
	file_gate_proto_rawDescOnce.Do(func() {
		file_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_gate_proto_rawDescData)
	})
	return file_gate_proto_rawDescData
}

var file_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_gate_proto_goTypes = []interface{}{
	(*GateCommTaskInfo)(nil),         // 0: GateCommTaskInfo
	(*GateGetGiftListReq)(nil),       // 1: GateGetGiftListReq
	(*GateGetGiftListRsp)(nil),       // 2: GateGetGiftListRsp
	(*GateGetVipCareReq)(nil),        // 3: GateGetVipCareReq
	(*GateGetVipCareRsp)(nil),        // 4: GateGetVipCareRsp
	(*GateOidbFlagInfo)(nil),         // 5: GateOidbFlagInfo
	(*GatePrivilegeBaseInfoReq)(nil), // 6: GatePrivilegeBaseInfoReq
	(*GatePrivilegeBaseInfoRsp)(nil), // 7: GatePrivilegeBaseInfoRsp
	(*GatePrivilegeInfo)(nil),        // 8: GatePrivilegeInfo
	(*GateVaProfileGateReq)(nil),     // 9: GateVaProfileGateReq
	(*GateQidInfoItem)(nil),          // 10: GateQidInfoItem
	(*GateVaProfileGateRsp)(nil),     // 11: GateVaProfileGateRsp
}
var file_gate_proto_depIdxs = []int32{
	8,  // 0: GatePrivilegeBaseInfoRsp.vOpenPriv:type_name -> GatePrivilegeInfo
	8,  // 1: GatePrivilegeBaseInfoRsp.vClosePriv:type_name -> GatePrivilegeInfo
	6,  // 2: GateVaProfileGateReq.stPrivilegeReq:type_name -> GatePrivilegeBaseInfoReq
	1,  // 3: GateVaProfileGateReq.stGiftReq:type_name -> GateGetGiftListReq
	0,  // 4: GateVaProfileGateReq.taskItem:type_name -> GateCommTaskInfo
	5,  // 5: GateVaProfileGateReq.oidbFlag:type_name -> GateOidbFlagInfo
	3,  // 6: GateVaProfileGateReq.stVipCare:type_name -> GateGetVipCareReq
	7,  // 7: GateVaProfileGateRsp.stPrivilegeRsp:type_name -> GatePrivilegeBaseInfoRsp
	2,  // 8: GateVaProfileGateRsp.stGiftRsp:type_name -> GateGetGiftListRsp
	0,  // 9: GateVaProfileGateRsp.taskItem:type_name -> GateCommTaskInfo
	5,  // 10: GateVaProfileGateRsp.oidbFlag:type_name -> GateOidbFlagInfo
	4,  // 11: GateVaProfileGateRsp.stVipCare:type_name -> GateGetVipCareRsp
	10, // 12: GateVaProfileGateRsp.qidInfo:type_name -> GateQidInfoItem
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_gate_proto_init() }
func file_gate_proto_init() {
	if File_gate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateCommTaskInfo); i {
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
		file_gate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateGetGiftListReq); i {
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
		file_gate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateGetGiftListRsp); i {
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
		file_gate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateGetVipCareReq); i {
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
		file_gate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateGetVipCareRsp); i {
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
		file_gate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateOidbFlagInfo); i {
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
		file_gate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatePrivilegeBaseInfoReq); i {
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
		file_gate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatePrivilegeBaseInfoRsp); i {
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
		file_gate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatePrivilegeInfo); i {
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
		file_gate_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateVaProfileGateReq); i {
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
		file_gate_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateQidInfoItem); i {
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
		file_gate_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateVaProfileGateRsp); i {
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
			RawDescriptor: file_gate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gate_proto_goTypes,
		DependencyIndexes: file_gate_proto_depIdxs,
		MessageInfos:      file_gate_proto_msgTypes,
	}.Build()
	File_gate_proto = out.File
	file_gate_proto_rawDesc = nil
	file_gate_proto_goTypes = nil
	file_gate_proto_depIdxs = nil
}
