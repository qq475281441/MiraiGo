// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: oidb0xec4.proto

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Comment  *string `protobuf:"bytes,2,opt,name=comment" json:"comment,omitempty"`
	Time     *uint64 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	FromUin  *uint64 `protobuf:"varint,4,opt,name=fromUin" json:"fromUin,omitempty"`
	ToUin    *uint64 `protobuf:"varint,5,opt,name=toUin" json:"toUin,omitempty"`
	ReplyId  *string `protobuf:"bytes,6,opt,name=replyId" json:"replyId,omitempty"`
	FromNick *string `protobuf:"bytes,7,opt,name=fromNick" json:"fromNick,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xec4_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xec4_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_oidb0xec4_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Comment) GetComment() string {
	if x != nil && x.Comment != nil {
		return *x.Comment
	}
	return ""
}

func (x *Comment) GetTime() uint64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Comment) GetFromUin() uint64 {
	if x != nil && x.FromUin != nil {
		return *x.FromUin
	}
	return 0
}

func (x *Comment) GetToUin() uint64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *Comment) GetReplyId() string {
	if x != nil && x.ReplyId != nil {
		return *x.ReplyId
	}
	return ""
}

func (x *Comment) GetFromNick() string {
	if x != nil && x.FromNick != nil {
		return *x.FromNick
	}
	return ""
}

type Praise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUin  *uint64 `protobuf:"varint,1,opt,name=fromUin" json:"fromUin,omitempty"`
	ToUin    *uint64 `protobuf:"varint,2,opt,name=toUin" json:"toUin,omitempty"`
	Time     *uint64 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	FromNick *string `protobuf:"bytes,4,opt,name=fromNick" json:"fromNick,omitempty"`
}

func (x *Praise) Reset() {
	*x = Praise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xec4_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Praise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Praise) ProtoMessage() {}

func (x *Praise) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xec4_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Praise.ProtoReflect.Descriptor instead.
func (*Praise) Descriptor() ([]byte, []int) {
	return file_oidb0xec4_proto_rawDescGZIP(), []int{1}
}

func (x *Praise) GetFromUin() uint64 {
	if x != nil && x.FromUin != nil {
		return *x.FromUin
	}
	return 0
}

func (x *Praise) GetToUin() uint64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *Praise) GetTime() uint64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Praise) GetFromNick() string {
	if x != nil && x.FromNick != nil {
		return *x.FromNick
	}
	return ""
}

type Quest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Quest       *string    `protobuf:"bytes,2,opt,name=quest" json:"quest,omitempty"`
	QuestUin    *uint64    `protobuf:"varint,3,opt,name=questUin" json:"questUin,omitempty"`
	Time        *uint64    `protobuf:"varint,4,opt,name=time" json:"time,omitempty"`
	Ans         *string    `protobuf:"bytes,5,opt,name=ans" json:"ans,omitempty"`
	AnsTime     *uint64    `protobuf:"varint,6,opt,name=ansTime" json:"ansTime,omitempty"`
	Comment     []*Comment `protobuf:"bytes,7,rep,name=comment" json:"comment,omitempty"`
	Praise      []*Praise  `protobuf:"bytes,8,rep,name=praise" json:"praise,omitempty"`
	PraiseNum   *uint64    `protobuf:"varint,9,opt,name=praiseNum" json:"praiseNum,omitempty"`
	LikeKey     *string    `protobuf:"bytes,10,opt,name=likeKey" json:"likeKey,omitempty"`
	SystemId    *uint64    `protobuf:"varint,11,opt,name=systemId" json:"systemId,omitempty"`
	CommentNum  *uint64    `protobuf:"varint,12,opt,name=commentNum" json:"commentNum,omitempty"`
	ShowType    *uint64    `protobuf:"varint,13,opt,name=showType" json:"showType,omitempty"`
	ShowTimes   *uint64    `protobuf:"varint,14,opt,name=showTimes" json:"showTimes,omitempty"`
	BeenPraised *uint64    `protobuf:"varint,15,opt,name=beenPraised" json:"beenPraised,omitempty"`
	QuestRead   *bool      `protobuf:"varint,16,opt,name=questRead" json:"questRead,omitempty"`
	AnsShowType *uint64    `protobuf:"varint,17,opt,name=ansShowType" json:"ansShowType,omitempty"`
}

func (x *Quest) Reset() {
	*x = Quest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xec4_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quest) ProtoMessage() {}

func (x *Quest) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xec4_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quest.ProtoReflect.Descriptor instead.
func (*Quest) Descriptor() ([]byte, []int) {
	return file_oidb0xec4_proto_rawDescGZIP(), []int{2}
}

func (x *Quest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Quest) GetQuest() string {
	if x != nil && x.Quest != nil {
		return *x.Quest
	}
	return ""
}

func (x *Quest) GetQuestUin() uint64 {
	if x != nil && x.QuestUin != nil {
		return *x.QuestUin
	}
	return 0
}

func (x *Quest) GetTime() uint64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *Quest) GetAns() string {
	if x != nil && x.Ans != nil {
		return *x.Ans
	}
	return ""
}

func (x *Quest) GetAnsTime() uint64 {
	if x != nil && x.AnsTime != nil {
		return *x.AnsTime
	}
	return 0
}

func (x *Quest) GetComment() []*Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *Quest) GetPraise() []*Praise {
	if x != nil {
		return x.Praise
	}
	return nil
}

func (x *Quest) GetPraiseNum() uint64 {
	if x != nil && x.PraiseNum != nil {
		return *x.PraiseNum
	}
	return 0
}

func (x *Quest) GetLikeKey() string {
	if x != nil && x.LikeKey != nil {
		return *x.LikeKey
	}
	return ""
}

func (x *Quest) GetSystemId() uint64 {
	if x != nil && x.SystemId != nil {
		return *x.SystemId
	}
	return 0
}

func (x *Quest) GetCommentNum() uint64 {
	if x != nil && x.CommentNum != nil {
		return *x.CommentNum
	}
	return 0
}

func (x *Quest) GetShowType() uint64 {
	if x != nil && x.ShowType != nil {
		return *x.ShowType
	}
	return 0
}

func (x *Quest) GetShowTimes() uint64 {
	if x != nil && x.ShowTimes != nil {
		return *x.ShowTimes
	}
	return 0
}

func (x *Quest) GetBeenPraised() uint64 {
	if x != nil && x.BeenPraised != nil {
		return *x.BeenPraised
	}
	return 0
}

func (x *Quest) GetQuestRead() bool {
	if x != nil && x.QuestRead != nil {
		return *x.QuestRead
	}
	return false
}

func (x *Quest) GetAnsShowType() uint64 {
	if x != nil && x.AnsShowType != nil {
		return *x.AnsShowType
	}
	return 0
}

type DEC4ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin        *uint64 `protobuf:"varint,1,opt,name=uin" json:"uin,omitempty"`
	QuestNum   *uint64 `protobuf:"varint,2,opt,name=questNum" json:"questNum,omitempty"`
	CommentNum *uint64 `protobuf:"varint,3,opt,name=commentNum" json:"commentNum,omitempty"`
	Cookie     []byte  `protobuf:"bytes,4,opt,name=cookie" json:"cookie,omitempty"`
	FetchType  *uint32 `protobuf:"varint,5,opt,name=fetchType" json:"fetchType,omitempty"`
}

func (x *DEC4ReqBody) Reset() {
	*x = DEC4ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xec4_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DEC4ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DEC4ReqBody) ProtoMessage() {}

func (x *DEC4ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xec4_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DEC4ReqBody.ProtoReflect.Descriptor instead.
func (*DEC4ReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0xec4_proto_rawDescGZIP(), []int{3}
}

func (x *DEC4ReqBody) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *DEC4ReqBody) GetQuestNum() uint64 {
	if x != nil && x.QuestNum != nil {
		return *x.QuestNum
	}
	return 0
}

func (x *DEC4ReqBody) GetCommentNum() uint64 {
	if x != nil && x.CommentNum != nil {
		return *x.CommentNum
	}
	return 0
}

func (x *DEC4ReqBody) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *DEC4ReqBody) GetFetchType() uint32 {
	if x != nil && x.FetchType != nil {
		return *x.FetchType
	}
	return 0
}

type DEC4RspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quest            []*Quest `protobuf:"bytes,1,rep,name=quest" json:"quest,omitempty"`
	IsFetchOver      *bool    `protobuf:"varint,2,opt,name=isFetchOver" json:"isFetchOver,omitempty"`
	TotalQuestNum    *uint32  `protobuf:"varint,3,opt,name=totalQuestNum" json:"totalQuestNum,omitempty"`
	Cookie           []byte   `protobuf:"bytes,4,opt,name=cookie" json:"cookie,omitempty"`
	Ret              *uint32  `protobuf:"varint,5,opt,name=ret" json:"ret,omitempty"`
	AnsweredQuestNum *uint32  `protobuf:"varint,6,opt,name=answeredQuestNum" json:"answeredQuestNum,omitempty"`
}

func (x *DEC4RspBody) Reset() {
	*x = DEC4RspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0xec4_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DEC4RspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DEC4RspBody) ProtoMessage() {}

func (x *DEC4RspBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0xec4_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DEC4RspBody.ProtoReflect.Descriptor instead.
func (*DEC4RspBody) Descriptor() ([]byte, []int) {
	return file_oidb0xec4_proto_rawDescGZIP(), []int{4}
}

func (x *DEC4RspBody) GetQuest() []*Quest {
	if x != nil {
		return x.Quest
	}
	return nil
}

func (x *DEC4RspBody) GetIsFetchOver() bool {
	if x != nil && x.IsFetchOver != nil {
		return *x.IsFetchOver
	}
	return false
}

func (x *DEC4RspBody) GetTotalQuestNum() uint32 {
	if x != nil && x.TotalQuestNum != nil {
		return *x.TotalQuestNum
	}
	return 0
}

func (x *DEC4RspBody) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *DEC4RspBody) GetRet() uint32 {
	if x != nil && x.Ret != nil {
		return *x.Ret
	}
	return 0
}

func (x *DEC4RspBody) GetAnsweredQuestNum() uint32 {
	if x != nil && x.AnsweredQuestNum != nil {
		return *x.AnsweredQuestNum
	}
	return 0
}

var File_oidb0xec4_proto protoreflect.FileDescriptor

var file_oidb0xec4_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x65, 0x63, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x66, 0x72,
	0x6f, 0x6d, 0x55, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63,
	0x6b, 0x22, 0x68, 0x0a, 0x06, 0x50, 0x72, 0x61, 0x69, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x66, 0x72,
	0x6f, 0x6d, 0x55, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x4e, 0x69, 0x63, 0x6b, 0x22, 0xde, 0x03, 0x0a, 0x05,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x55, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6e, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x6e, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x61, 0x6e, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x70,
	0x72, 0x61, 0x69, 0x73, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x72,
	0x61, 0x69, 0x73, 0x65, 0x52, 0x06, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x70, 0x72, 0x61, 0x69, 0x73, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69,
	0x6b, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x6b,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x73, 0x68, 0x6f, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65,
	0x65, 0x6e, 0x50, 0x72, 0x61, 0x69, 0x73, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x62, 0x65, 0x65, 0x6e, 0x50, 0x72, 0x61, 0x69, 0x73, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e,
	0x73, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x61, 0x6e, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x22, 0x91, 0x01, 0x0a,
	0x0b, 0x44, 0x45, 0x43, 0x34, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x44, 0x45, 0x43, 0x34, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x1c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x69, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4f, 0x76, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x72, 0x65, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x65, 0x64, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x61, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x65, 0x64, 0x51, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x3b, 0x6f, 0x69, 0x64, 0x62,
}

var (
	file_oidb0xec4_proto_rawDescOnce sync.Once
	file_oidb0xec4_proto_rawDescData = file_oidb0xec4_proto_rawDesc
)

func file_oidb0xec4_proto_rawDescGZIP() []byte {
	file_oidb0xec4_proto_rawDescOnce.Do(func() {
		file_oidb0xec4_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0xec4_proto_rawDescData)
	})
	return file_oidb0xec4_proto_rawDescData
}

var file_oidb0xec4_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_oidb0xec4_proto_goTypes = []interface{}{
	(*Comment)(nil),     // 0: Comment
	(*Praise)(nil),      // 1: Praise
	(*Quest)(nil),       // 2: Quest
	(*DEC4ReqBody)(nil), // 3: DEC4ReqBody
	(*DEC4RspBody)(nil), // 4: DEC4RspBody
}
var file_oidb0xec4_proto_depIdxs = []int32{
	0, // 0: Quest.comment:type_name -> Comment
	1, // 1: Quest.praise:type_name -> Praise
	2, // 2: DEC4RspBody.quest:type_name -> Quest
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_oidb0xec4_proto_init() }
func file_oidb0xec4_proto_init() {
	if File_oidb0xec4_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0xec4_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_oidb0xec4_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Praise); i {
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
		file_oidb0xec4_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quest); i {
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
		file_oidb0xec4_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DEC4ReqBody); i {
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
		file_oidb0xec4_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DEC4RspBody); i {
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
			RawDescriptor: file_oidb0xec4_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0xec4_proto_goTypes,
		DependencyIndexes: file_oidb0xec4_proto_depIdxs,
		MessageInfos:      file_oidb0xec4_proto_msgTypes,
	}.Build()
	File_oidb0xec4_proto = out.File
	file_oidb0xec4_proto_rawDesc = nil
	file_oidb0xec4_proto_goTypes = nil
	file_oidb0xec4_proto_depIdxs = nil
}
