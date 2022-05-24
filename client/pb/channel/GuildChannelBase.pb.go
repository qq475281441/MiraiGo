// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/channel/GuildChannelBase.proto

package channel

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type ChannelUserInfo struct {
	ClientIdentity *ClientIdentity        `protobuf:"bytes,1,opt"`
	MemberType     proto.Option[uint32]   `protobuf:"varint,2,opt"`
	Permission     *ChannelUserPermission `protobuf:"bytes,3,opt"`
	RoleGroups     []*BaseRoleGroupInfo   `protobuf:"bytes,4,rep"`
}

type ChannelUserPermission struct {
	AllowReadFeed  proto.Option[bool] `protobuf:"varint,1,opt"`
	AllowWriteFeed proto.Option[bool] `protobuf:"varint,2,opt"`
}

type ClientIdentity struct {
	ClientId proto.Option[uint32] `protobuf:"varint,1,opt"`
	Desc     proto.Option[string] `protobuf:"bytes,2,opt"`
}

type BaseGuildInfo struct {
	GuildId  proto.Option[uint64] `protobuf:"varint,1,opt"`
	Name     proto.Option[string] `protobuf:"bytes,2,opt"`
	JoinTime proto.Option[uint64] `protobuf:"varint,3,opt"`
}

type BaseRoleGroupInfo struct {
	RoleId proto.Option[uint64] `protobuf:"varint,1,opt"`
	Name   proto.Option[string] `protobuf:"bytes,2,opt"`
	Color  proto.Option[uint32] `protobuf:"varint,3,opt"`
}

type StChannelInfo struct {
	Sign    *StChannelSign       `protobuf:"bytes,1,opt"`
	Name    proto.Option[string] `protobuf:"bytes,2,opt"`
	IconUrl proto.Option[string] `protobuf:"bytes,3,opt"`
}

type StChannelSign struct {
	GuildId   proto.Option[uint64] `protobuf:"varint,1,opt"`
	ChannelId proto.Option[uint64] `protobuf:"varint,2,opt"`
}

type StEmotionReactionInfo struct {
	Id                proto.Option[string] `protobuf:"bytes,1,opt"`
	EmojiReactionList []*EmojiReaction     `protobuf:"bytes,2,rep"`
}

type StCommonExt struct {
	MapInfo      []*CommonEntry       `protobuf:"bytes,1,rep"`
	AttachInfo   proto.Option[string] `protobuf:"bytes,2,opt"`
	MapBytesInfo []*BytesEntry        `protobuf:"bytes,3,rep"`
}

type BytesEntry struct {
	Key   proto.Option[string] `protobuf:"bytes,1,opt"`
	Value []byte               `protobuf:"bytes,2,opt"`
}

type CommonEntry struct {
	Key   proto.Option[string] `protobuf:"bytes,1,opt"`
	Value proto.Option[string] `protobuf:"bytes,2,opt"`
}
