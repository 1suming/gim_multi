// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: room.proto

package pb

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

type RoomCreateRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                            // 名称
	AvatarUrl    string `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // 头像
	Introduction string `protobuf:"bytes,3,opt,name=introduction,proto3" json:"introduction,omitempty"`            // 简介
	Extra        string `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`                          // 附加字段
}

func (x *RoomCreateRoomReq) Reset() {
	*x = RoomCreateRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreateRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreateRoomReq) ProtoMessage() {}

func (x *RoomCreateRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreateRoomReq.ProtoReflect.Descriptor instead.
func (*RoomCreateRoomReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{0}
}

func (x *RoomCreateRoomReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoomCreateRoomReq) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *RoomCreateRoomReq) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *RoomCreateRoomReq) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type RoomCreateRoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // 群组id
}

func (x *RoomCreateRoomResp) Reset() {
	*x = RoomCreateRoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomCreateRoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomCreateRoomResp) ProtoMessage() {}

func (x *RoomCreateRoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomCreateRoomResp.ProtoReflect.Descriptor instead.
func (*RoomCreateRoomResp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{1}
}

func (x *RoomCreateRoomResp) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type ChatRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatroomId   int64  `protobuf:"varint,1,opt,name=chatroom_id,json=chatroomId,proto3" json:"chatroom_id,omitempty"`   // 群组id
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                  // 名称
	AvatarUrl    string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`       // 头像
	Introduction string `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`                  // 简介
	MaxUserMum   int32  `protobuf:"varint,5,opt,name=max_user_mum,json=maxUserMum,proto3" json:"max_user_mum,omitempty"` // 用户数
	Extra        string `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`                                // 附加字段
	CreateTime   int64  `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`   // 创建时间
	UpdateTime   int64  `protobuf:"varint,8,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`   // 更新时间
}

func (x *ChatRoom) Reset() {
	*x = ChatRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRoom) ProtoMessage() {}

func (x *ChatRoom) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRoom.ProtoReflect.Descriptor instead.
func (*ChatRoom) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{2}
}

func (x *ChatRoom) GetChatroomId() int64 {
	if x != nil {
		return x.ChatroomId
	}
	return 0
}

func (x *ChatRoom) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatRoom) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *ChatRoom) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *ChatRoom) GetMaxUserMum() int32 {
	if x != nil {
		return x.MaxUserMum
	}
	return 0
}

func (x *ChatRoom) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *ChatRoom) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ChatRoom) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type GetChatRoomsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*ChatRoom `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
}

func (x *GetChatRoomsResp) Reset() {
	*x = GetChatRoomsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRoomsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRoomsResp) ProtoMessage() {}

func (x *GetChatRoomsResp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatRoomsResp.ProtoReflect.Descriptor instead.
func (*GetChatRoomsResp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{3}
}

func (x *GetChatRoomsResp) GetRooms() []*ChatRoom {
	if x != nil {
		return x.Rooms
	}
	return nil
}

//roomo
// 消息同步请求,package_type:2
type GetRoomListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRoomListReq) Reset() {
	*x = GetRoomListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomListReq) ProtoMessage() {}

func (x *GetRoomListReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomListReq.ProtoReflect.Descriptor instead.
func (*GetRoomListReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{4}
}

// 消息同步响应,package_type:2
type GetRoomListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms []*ChatRoom `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"` // 消息列表
}

func (x *GetRoomListResp) Reset() {
	*x = GetRoomListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomListResp) ProtoMessage() {}

func (x *GetRoomListResp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomListResp.ProtoReflect.Descriptor instead.
func (*GetRoomListResp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{5}
}

func (x *GetRoomListResp) GetRooms() []*ChatRoom {
	if x != nil {
		return x.Rooms
	}
	return nil
}

//
type QuitRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // 房间ID，如果为0，取消房间订阅
}

func (x *QuitRoomReq) Reset() {
	*x = QuitRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitRoomReq) ProtoMessage() {}

func (x *QuitRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitRoomReq.ProtoReflect.Descriptor instead.
func (*QuitRoomReq) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{6}
}

func (x *QuitRoomReq) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

type QuitRoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QuitRoomResp) Reset() {
	*x = QuitRoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_room_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitRoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitRoomResp) ProtoMessage() {}

func (x *QuitRoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_room_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitRoomResp.ProtoReflect.Descriptor instead.
func (*QuitRoomResp) Descriptor() ([]byte, []int) {
	return file_room_proto_rawDescGZIP(), []int{7}
}

var File_room_proto protoreflect.FileDescriptor

var file_room_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x80, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x22, 0x2d, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x22, 0xfc, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x35, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22,
	0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f,
	0x6d, 0x73, 0x22, 0x26, 0x0a, 0x0b, 0x51, 0x75, 0x69, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x51, 0x75,
	0x69, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x69,
	0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_room_proto_rawDescOnce sync.Once
	file_room_proto_rawDescData = file_room_proto_rawDesc
)

func file_room_proto_rawDescGZIP() []byte {
	file_room_proto_rawDescOnce.Do(func() {
		file_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_room_proto_rawDescData)
	})
	return file_room_proto_rawDescData
}

var file_room_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_room_proto_goTypes = []interface{}{
	(*RoomCreateRoomReq)(nil),  // 0: pb.RoomCreateRoomReq
	(*RoomCreateRoomResp)(nil), // 1: pb.RoomCreateRoomResp
	(*ChatRoom)(nil),           // 2: pb.ChatRoom
	(*GetChatRoomsResp)(nil),   // 3: pb.GetChatRoomsResp
	(*GetRoomListReq)(nil),     // 4: pb.GetRoomListReq
	(*GetRoomListResp)(nil),    // 5: pb.GetRoomListResp
	(*QuitRoomReq)(nil),        // 6: pb.QuitRoomReq
	(*QuitRoomResp)(nil),       // 7: pb.QuitRoomResp
}
var file_room_proto_depIdxs = []int32{
	2, // 0: pb.GetChatRoomsResp.rooms:type_name -> pb.ChatRoom
	2, // 1: pb.GetRoomListResp.rooms:type_name -> pb.ChatRoom
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_room_proto_init() }
func file_room_proto_init() {
	if File_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomCreateRoomReq); i {
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
		file_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomCreateRoomResp); i {
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
		file_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRoom); i {
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
		file_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatRoomsResp); i {
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
		file_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomListReq); i {
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
		file_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomListResp); i {
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
		file_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitRoomReq); i {
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
		file_room_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitRoomResp); i {
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
			RawDescriptor: file_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_room_proto_goTypes,
		DependencyIndexes: file_room_proto_depIdxs,
		MessageInfos:      file_room_proto_msgTypes,
	}.Build()
	File_room_proto = out.File
	file_room_proto_rawDesc = nil
	file_room_proto_goTypes = nil
	file_room_proto_depIdxs = nil
}
