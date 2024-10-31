// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: connect.ext.proto

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

type PackageType int32

const (
	PackageType_PT_UNKNOWN        PackageType = 0 // 未知
	PackageType_PT_SIGN_IN        PackageType = 1 // 设备登录请求
	PackageType_PT_SYNC           PackageType = 2 // 消息同步触发
	PackageType_PT_HEARTBEAT      PackageType = 3 // 心跳
	PackageType_PT_MESSAGE        PackageType = 4 // 消息投递
	PackageType_PT_SUBSCRIBE_ROOM PackageType = 5 // 订阅房间
	//@ms:add
	PackageType_PT_SEARCH_USER               PackageType = 10001 //搜索用户
	PackageType_PT_GET_USER                  PackageType = 10002
	PackageType_PT_GET_USERS                 PackageType = 10003
	PackageType_PT_UPDATE_USER               PackageType = 1004
	PackageType_PT_FRIEND_ADD_FRIEND         PackageType = 10010 //添加好友
	PackageType_PT_FRIEND_SEND_MSG_TO_FRIEND PackageType = 10011 //给
)

// Enum value maps for PackageType.
var (
	PackageType_name = map[int32]string{
		0:     "PT_UNKNOWN",
		1:     "PT_SIGN_IN",
		2:     "PT_SYNC",
		3:     "PT_HEARTBEAT",
		4:     "PT_MESSAGE",
		5:     "PT_SUBSCRIBE_ROOM",
		10001: "PT_SEARCH_USER",
		10002: "PT_GET_USER",
		10003: "PT_GET_USERS",
		1004:  "PT_UPDATE_USER",
		10010: "PT_FRIEND_ADD_FRIEND",
		10011: "PT_FRIEND_SEND_MSG_TO_FRIEND",
	}
	PackageType_value = map[string]int32{
		"PT_UNKNOWN":                   0,
		"PT_SIGN_IN":                   1,
		"PT_SYNC":                      2,
		"PT_HEARTBEAT":                 3,
		"PT_MESSAGE":                   4,
		"PT_SUBSCRIBE_ROOM":            5,
		"PT_SEARCH_USER":               10001,
		"PT_GET_USER":                  10002,
		"PT_GET_USERS":                 10003,
		"PT_UPDATE_USER":               1004,
		"PT_FRIEND_ADD_FRIEND":         10010,
		"PT_FRIEND_SEND_MSG_TO_FRIEND": 10011,
	}
)

func (x PackageType) Enum() *PackageType {
	p := new(PackageType)
	*p = x
	return p
}

func (x PackageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageType) Descriptor() protoreflect.EnumDescriptor {
	return file_connect_ext_proto_enumTypes[0].Descriptor()
}

func (PackageType) Type() protoreflect.EnumType {
	return &file_connect_ext_proto_enumTypes[0]
}

func (x PackageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageType.Descriptor instead.
func (PackageType) EnumDescriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{0}
}

// 上行数据
type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      PackageType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PackageType" json:"type,omitempty"`        // 包的类型
	RequestId int64       `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"` // 请求id
	Data      []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                             // 数据
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetType() PackageType {
	if x != nil {
		return x.Type
	}
	return PackageType_PT_UNKNOWN
}

func (x *Input) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Input) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

//@ms:msg server给router发消息
type RouterMsgInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginInput *Input `protobuf:"bytes,1,opt,name=originInput,proto3" json:"originInput,omitempty"`      //原始数据
	UserId      int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
}

func (x *RouterMsgInput) Reset() {
	*x = RouterMsgInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouterMsgInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouterMsgInput) ProtoMessage() {}

func (x *RouterMsgInput) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouterMsgInput.ProtoReflect.Descriptor instead.
func (*RouterMsgInput) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{1}
}

func (x *RouterMsgInput) GetOriginInput() *Input {
	if x != nil {
		return x.OriginInput
	}
	return nil
}

func (x *RouterMsgInput) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 下行数据
type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      PackageType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.PackageType" json:"type,omitempty"`        // 包的类型
	RequestId int64       `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"` // 请求id
	Code      int32       `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`                            // 错误码
	Message   string      `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`                       // 错误信息
	Data      []byte      `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`                             // 数据
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{2}
}

func (x *Output) GetType() PackageType {
	if x != nil {
		return x.Type
	}
	return PackageType_PT_UNKNOWN
}

func (x *Output) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Output) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Output) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Output) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 设备登录,package_type:1
type SignInInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId int64  `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"` // 设备id
	UserId   int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // 用户id
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`                        // 秘钥
}

func (x *SignInInput) Reset() {
	*x = SignInInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInInput) ProtoMessage() {}

func (x *SignInInput) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInInput.ProtoReflect.Descriptor instead.
func (*SignInInput) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{3}
}

func (x *SignInInput) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *SignInInput) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SignInInput) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 消息同步请求,package_type:2
type SyncInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq int64 `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"` // 客户端已经同步的序列号
}

func (x *SyncInput) Reset() {
	*x = SyncInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncInput) ProtoMessage() {}

func (x *SyncInput) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncInput.ProtoReflect.Descriptor instead.
func (*SyncInput) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{4}
}

func (x *SyncInput) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

// 消息同步响应,package_type:2
type SyncOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`               // 消息列表
	HasMore  bool       `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"` // 是否有更多数据
}

func (x *SyncOutput) Reset() {
	*x = SyncOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncOutput) ProtoMessage() {}

func (x *SyncOutput) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncOutput.ProtoReflect.Descriptor instead.
func (*SyncOutput) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{5}
}

func (x *SyncOutput) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *SyncOutput) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

// 订阅房间请求
type SubscribeRoomInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` // 房间ID，如果为0，取消房间订阅
	Seq    int64 `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`                     // 消息消息序列号，
}

func (x *SubscribeRoomInput) Reset() {
	*x = SubscribeRoomInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRoomInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRoomInput) ProtoMessage() {}

func (x *SubscribeRoomInput) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRoomInput.ProtoReflect.Descriptor instead.
func (*SubscribeRoomInput) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeRoomInput) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *SubscribeRoomInput) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

// 投递消息回执,package_type:4
type MessageACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceAck   int64 `protobuf:"varint,2,opt,name=device_ack,json=deviceAck,proto3" json:"device_ack,omitempty"`       // 设备收到消息的确认号
	ReceiveTime int64 `protobuf:"varint,3,opt,name=receive_time,json=receiveTime,proto3" json:"receive_time,omitempty"` // 消息接收时间戳，精确到毫秒
}

func (x *MessageACK) Reset() {
	*x = MessageACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connect_ext_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageACK) ProtoMessage() {}

func (x *MessageACK) ProtoReflect() protoreflect.Message {
	mi := &file_connect_ext_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageACK.ProtoReflect.Descriptor instead.
func (*MessageACK) Descriptor() ([]byte, []int) {
	return file_connect_ext_proto_rawDescGZIP(), []int{7}
}

func (x *MessageACK) GetDeviceAck() int64 {
	if x != nil {
		return x.DeviceAck
	}
	return 0
}

func (x *MessageACK) GetReceiveTime() int64 {
	if x != nil {
		return x.ReceiveTime
	}
	return 0
}

var File_connect_ext_proto protoreflect.FileDescriptor

var file_connect_ext_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x05, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x0e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a,
	0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0b, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x23,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x1d, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x71, 0x22, 0x50,
	0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65,
	0x22, 0x3f, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65,
	0x71, 0x22, 0x4e, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x43, 0x4b, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x2a, 0x80, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x50, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x03,
	0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x04,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45,
	0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x91, 0x4e, 0x12, 0x10, 0x0a, 0x0b,
	0x50, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x92, 0x4e, 0x12, 0x11,
	0x0a, 0x0c, 0x50, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x10, 0x93,
	0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x53, 0x45, 0x52, 0x10, 0xec, 0x07, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x54, 0x5f, 0x46, 0x52, 0x49,
	0x45, 0x4e, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x9a,
	0x4e, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x53,
	0x45, 0x4e, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x10, 0x9b, 0x4e, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_connect_ext_proto_rawDescOnce sync.Once
	file_connect_ext_proto_rawDescData = file_connect_ext_proto_rawDesc
)

func file_connect_ext_proto_rawDescGZIP() []byte {
	file_connect_ext_proto_rawDescOnce.Do(func() {
		file_connect_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_connect_ext_proto_rawDescData)
	})
	return file_connect_ext_proto_rawDescData
}

var file_connect_ext_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_connect_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_connect_ext_proto_goTypes = []interface{}{
	(PackageType)(0),           // 0: pb.PackageType
	(*Input)(nil),              // 1: pb.Input
	(*RouterMsgInput)(nil),     // 2: pb.RouterMsgInput
	(*Output)(nil),             // 3: pb.Output
	(*SignInInput)(nil),        // 4: pb.SignInInput
	(*SyncInput)(nil),          // 5: pb.SyncInput
	(*SyncOutput)(nil),         // 6: pb.SyncOutput
	(*SubscribeRoomInput)(nil), // 7: pb.SubscribeRoomInput
	(*MessageACK)(nil),         // 8: pb.MessageACK
	(*Message)(nil),            // 9: pb.Message
}
var file_connect_ext_proto_depIdxs = []int32{
	0, // 0: pb.Input.type:type_name -> pb.PackageType
	1, // 1: pb.RouterMsgInput.originInput:type_name -> pb.Input
	0, // 2: pb.Output.type:type_name -> pb.PackageType
	9, // 3: pb.SyncOutput.messages:type_name -> pb.Message
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_connect_ext_proto_init() }
func file_connect_ext_proto_init() {
	if File_connect_ext_proto != nil {
		return
	}
	file_message_ext_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_connect_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_connect_ext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouterMsgInput); i {
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
		file_connect_ext_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_connect_ext_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInInput); i {
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
		file_connect_ext_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncInput); i {
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
		file_connect_ext_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncOutput); i {
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
		file_connect_ext_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRoomInput); i {
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
		file_connect_ext_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageACK); i {
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
			RawDescriptor: file_connect_ext_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connect_ext_proto_goTypes,
		DependencyIndexes: file_connect_ext_proto_depIdxs,
		EnumInfos:         file_connect_ext_proto_enumTypes,
		MessageInfos:      file_connect_ext_proto_msgTypes,
	}.Build()
	File_connect_ext_proto = out.File
	file_connect_ext_proto_rawDesc = nil
	file_connect_ext_proto_goTypes = nil
	file_connect_ext_proto_depIdxs = nil
}
