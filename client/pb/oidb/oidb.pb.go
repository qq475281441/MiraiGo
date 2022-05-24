// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type OIDBSSOPkg struct {
	Command       int32  `protobuf:"varint,1,opt"`
	ServiceType   int32  `protobuf:"varint,2,opt"`
	Result        int32  `protobuf:"varint,3,opt"`
	Bodybuffer    []byte `protobuf:"bytes,4,opt"`
	ErrorMsg      string `protobuf:"bytes,5,opt"`
	ClientVersion string `protobuf:"bytes,6,opt"`
}

type D8A0RspBody struct {
	OptUint64GroupCode int64             `protobuf:"varint,1,opt"`
	MsgKickResult      []*D8A0KickResult `protobuf:"bytes,2,rep"`
}

type D8A0KickResult struct {
	OptUint32Result    int32 `protobuf:"varint,1,opt"`
	OptUint64MemberUin int64 `protobuf:"varint,2,opt"`
}

type D8A0KickMemberInfo struct {
	OptUint32Operate   int32  `protobuf:"varint,1,opt"`
	OptUint64MemberUin int64  `protobuf:"varint,2,opt"`
	OptUint32Flag      int32  `protobuf:"varint,3,opt"`
	OptBytesMsg        []byte `protobuf:"bytes,4,opt"`
}

type D8A0ReqBody struct {
	OptUint64GroupCode int64                 `protobuf:"varint,1,opt"`
	MsgKickList        []*D8A0KickMemberInfo `protobuf:"bytes,2,rep"`
	KickList           []int64               `protobuf:"varint,3,rep"`
	KickFlag           int32                 `protobuf:"varint,4,opt"`
	KickMsg            []byte                `protobuf:"bytes,5,opt"`
}

type D89AReqBody struct {
	GroupCode           int64          `protobuf:"varint,1,opt"`
	StGroupInfo         *D89AGroupinfo `protobuf:"bytes,2,opt"`
	OriginalOperatorUin int64          `protobuf:"varint,3,opt"`
	ReqGroupOpenAppid   int32          `protobuf:"varint,4,opt"`
}

type D89AGroupinfo struct {
	GroupExtAdmNum         int32                       `protobuf:"varint,1,opt"`
	Flag                   int32                       `protobuf:"varint,2,opt"`
	IngGroupName           []byte                      `protobuf:"bytes,3,opt"`
	IngGroupMemo           []byte                      `protobuf:"bytes,4,opt"`
	IngGroupFingerMemo     []byte                      `protobuf:"bytes,5,opt"`
	IngGroupAioSkinUrl     []byte                      `protobuf:"bytes,6,opt"`
	IngGroupBoardSkinUrl   []byte                      `protobuf:"bytes,7,opt"`
	IngGroupCoverSkinUrl   []byte                      `protobuf:"bytes,8,opt"`
	GroupGrade             int32                       `protobuf:"varint,9,opt"`
	ActiveMemberNum        int32                       `protobuf:"varint,10,opt"`
	CertificationType      int32                       `protobuf:"varint,11,opt"`
	IngCertificationText   []byte                      `protobuf:"bytes,12,opt"`
	IngGroupRichFingerMemo []byte                      `protobuf:"bytes,13,opt"`
	StGroupNewguidelines   *D89AGroupNewGuidelinesInfo `protobuf:"bytes,14,opt"`
	GroupFace              int32                       `protobuf:"varint,15,opt"`
	AddOption              int32                       `protobuf:"varint,16,opt"`
	ShutupTime             proto.Option[int32]         `protobuf:"varint,17,opt"`
	GroupTypeFlag          int32                       `protobuf:"varint,18,opt"`
	StringGroupTag         []byte                      `protobuf:"bytes,19,opt"`
	MsgGroupGeoInfo        *D89AGroupGeoInfo           `protobuf:"bytes,20,opt"`
	GroupClassExt          int32                       `protobuf:"varint,21,opt"`
	IngGroupClassText      []byte                      `protobuf:"bytes,22,opt"`
	AppPrivilegeFlag       int32                       `protobuf:"varint,23,opt"`
	AppPrivilegeMask       int32                       `protobuf:"varint,24,opt"`
	StGroupExInfo          *D89AGroupExInfoOnly        `protobuf:"bytes,25,opt"`
	GroupSecLevel          int32                       `protobuf:"varint,26,opt"`
	GroupSecLevelInfo      int32                       `protobuf:"varint,27,opt"`
	SubscriptionUin        int64                       `protobuf:"varint,28,opt"`
	AllowMemberInvite      int32                       `protobuf:"varint,29,opt"`
	IngGroupQuestion       []byte                      `protobuf:"bytes,30,opt"`
	IngGroupAnswer         []byte                      `protobuf:"bytes,31,opt"`
	GroupFlagext3          int32                       `protobuf:"varint,32,opt"`
	GroupFlagext3Mask      int32                       `protobuf:"varint,33,opt"`
	GroupOpenAppid         int32                       `protobuf:"varint,34,opt"`
	NoFingerOpenFlag       int32                       `protobuf:"varint,35,opt"`
	NoCodeFingerOpenFlag   int32                       `protobuf:"varint,36,opt"`
	RootId                 int64                       `protobuf:"varint,37,opt"`
	MsgLimitFrequency      int32                       `protobuf:"varint,38,opt"`
}

type D89AGroupNewGuidelinesInfo struct {
	BoolEnabled bool   `protobuf:"varint,1,opt"`
	IngContent  []byte `protobuf:"bytes,2,opt"`
}

type D89AGroupExInfoOnly struct {
	TribeId          int32 `protobuf:"varint,1,opt"`
	MoneyForAddGroup int32 `protobuf:"varint,2,opt"`
}

type D89AGroupGeoInfo struct {
	CityId        int32  `protobuf:"varint,1,opt"`
	Longtitude    int64  `protobuf:"varint,2,opt"`
	Latitude      int64  `protobuf:"varint,3,opt"`
	IngGeoContent []byte `protobuf:"bytes,4,opt"`
	PoiId         int64  `protobuf:"varint,5,opt"`
}

type DED3ReqBody struct {
	ToUin     int64 `protobuf:"varint,1,opt"`
	GroupCode int64 `protobuf:"varint,2,opt"`
	MsgSeq    int32 `protobuf:"varint,3,opt"`
	MsgRand   int32 `protobuf:"varint,4,opt"`
	AioUin    int64 `protobuf:"varint,5,opt"`
}
