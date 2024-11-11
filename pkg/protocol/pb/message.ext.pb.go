// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: message.ext.proto

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
	PackageType_PT_SEARCH_USER                PackageType = 10001 //搜索用户
	PackageType_PT_GET_USER                   PackageType = 10002
	PackageType_PT_GET_USERS                  PackageType = 10003
	PackageType_PT_UPDATE_USER                PackageType = 1004
	PackageType_PT_FRIEND_ADD_FRIEND          PackageType = 10010 //添加好友
	PackageType_PT_SEND_MESSAGE               PackageType = 10011 //给
	PackageType_PT_FRIEND_AGREE_ADD_FRIEND    PackageType = 10012
	PackageType_PT_FRIEND_SET_FRIEND          PackageType = 10013
	PackageType_PT_FRIEND_GET_FRIENDS         PackageType = 10014
	PackageType_PT_FRIEND_GET_FRIEND_REQUESTS PackageType = 10015 //获取好友请求
	PackageType_PT_GET_USER_CONVERSATIONS     PackageType = 20001 //用户会话
	PackageType_PT_GET_USER_MESSAGES          PackageType = 20002 //得到某个会话消息
	PackageType_PT_MESSAGE_ACK                PackageType = 30001 //消息回执
	PackageType_PT_ROOM_GET_ROOM_LIST         PackageType = 40001 //room列表
	PackageType_PT_ROOM_CREATE_ROOM           PackageType = 40002 //创建room
	PackageType_PT_ROOM_DELETE_ROOM           PackageType = 40003 //解散room
	PackageType_PT_ROOM_JOIN_ROOM             PackageType = 40004 //加入room
	PackageType_PT_ROOM_QUITROOM              PackageType = 40005 //退出room
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
		10011: "PT_SEND_MESSAGE",
		10012: "PT_FRIEND_AGREE_ADD_FRIEND",
		10013: "PT_FRIEND_SET_FRIEND",
		10014: "PT_FRIEND_GET_FRIENDS",
		10015: "PT_FRIEND_GET_FRIEND_REQUESTS",
		20001: "PT_GET_USER_CONVERSATIONS",
		20002: "PT_GET_USER_MESSAGES",
		30001: "PT_MESSAGE_ACK",
		40001: "PT_ROOM_GET_ROOM_LIST",
		40002: "PT_ROOM_CREATE_ROOM",
		40003: "PT_ROOM_DELETE_ROOM",
		40004: "PT_ROOM_JOIN_ROOM",
		40005: "PT_ROOM_QUITROOM",
	}
	PackageType_value = map[string]int32{
		"PT_UNKNOWN":                    0,
		"PT_SIGN_IN":                    1,
		"PT_SYNC":                       2,
		"PT_HEARTBEAT":                  3,
		"PT_MESSAGE":                    4,
		"PT_SUBSCRIBE_ROOM":             5,
		"PT_SEARCH_USER":                10001,
		"PT_GET_USER":                   10002,
		"PT_GET_USERS":                  10003,
		"PT_UPDATE_USER":                1004,
		"PT_FRIEND_ADD_FRIEND":          10010,
		"PT_SEND_MESSAGE":               10011,
		"PT_FRIEND_AGREE_ADD_FRIEND":    10012,
		"PT_FRIEND_SET_FRIEND":          10013,
		"PT_FRIEND_GET_FRIENDS":         10014,
		"PT_FRIEND_GET_FRIEND_REQUESTS": 10015,
		"PT_GET_USER_CONVERSATIONS":     20001,
		"PT_GET_USER_MESSAGES":          20002,
		"PT_MESSAGE_ACK":                30001,
		"PT_ROOM_GET_ROOM_LIST":         40001,
		"PT_ROOM_CREATE_ROOM":           40002,
		"PT_ROOM_DELETE_ROOM":           40003,
		"PT_ROOM_JOIN_ROOM":             40004,
		"PT_ROOM_QUITROOM":              40005,
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
	return file_message_ext_proto_enumTypes[0].Descriptor()
}

func (PackageType) Type() protoreflect.EnumType {
	return &file_message_ext_proto_enumTypes[0]
}

func (x PackageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageType.Descriptor instead.
func (PackageType) EnumDescriptor() ([]byte, []int) {
	return file_message_ext_proto_rawDescGZIP(), []int{0}
}

type ChatType int32

const (
	ChatType_DEFAULT     ChatType = 0
	ChatType_SINGLE_CHAT ChatType = 1
	ChatType_GROUP_CHAT  ChatType = 2
	ChatType_CHAT_ROOM   ChatType = 3 //    // 会话类型：单聊、群聊和聊天室分别为 `singleChat`、`groupChat` 和 `chatRoom`，默认为单聊。
)

// Enum value maps for ChatType.
var (
	ChatType_name = map[int32]string{
		0: "DEFAULT",
		1: "SINGLE_CHAT",
		2: "GROUP_CHAT",
		3: "CHAT_ROOM",
	}
	ChatType_value = map[string]int32{
		"DEFAULT":     0,
		"SINGLE_CHAT": 1,
		"GROUP_CHAT":  2,
		"CHAT_ROOM":   3,
	}
)

func (x ChatType) Enum() *ChatType {
	p := new(ChatType)
	*p = x
	return p
}

func (x ChatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_ext_proto_enumTypes[1].Descriptor()
}

func (ChatType) Type() protoreflect.EnumType {
	return &file_message_ext_proto_enumTypes[1]
}

func (x ChatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatType.Descriptor instead.
func (ChatType) EnumDescriptor() ([]byte, []int) {
	return file_message_ext_proto_rawDescGZIP(), []int{1}
}

type MessageContentType int32

const (
	MessageContentType_MCT_TEXT         MessageContentType = 0 //普通文本消息
	MessageContentType_MCT_NOTIFICATION MessageContentType = 1 //通知类消息
	MessageContentType_MCT_IMAGE        MessageContentType = 2 //图片类
	MessageContentType_MCT_FILE         MessageContentType = 3
)

// Enum value maps for MessageContentType.
var (
	MessageContentType_name = map[int32]string{
		0: "MCT_TEXT",
		1: "MCT_NOTIFICATION",
		2: "MCT_IMAGE",
		3: "MCT_FILE",
	}
	MessageContentType_value = map[string]int32{
		"MCT_TEXT":         0,
		"MCT_NOTIFICATION": 1,
		"MCT_IMAGE":        2,
		"MCT_FILE":         3,
	}
)

func (x MessageContentType) Enum() *MessageContentType {
	p := new(MessageContentType)
	*p = x
	return p
}

func (x MessageContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_ext_proto_enumTypes[2].Descriptor()
}

func (MessageContentType) Type() protoreflect.EnumType {
	return &file_message_ext_proto_enumTypes[2]
}

func (x MessageContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageContentType.Descriptor instead.
func (MessageContentType) EnumDescriptor() ([]byte, []int) {
	return file_message_ext_proto_rawDescGZIP(), []int{2}
}

type MessageStatus int32

const (
	MessageStatus_MS_NORMAL MessageStatus = 0 // 正常的
	MessageStatus_MS_RECALL MessageStatus = 1 // 撤回
)

// Enum value maps for MessageStatus.
var (
	MessageStatus_name = map[int32]string{
		0: "MS_NORMAL",
		1: "MS_RECALL",
	}
	MessageStatus_value = map[string]int32{
		"MS_NORMAL": 0,
		"MS_RECALL": 1,
	}
)

func (x MessageStatus) Enum() *MessageStatus {
	p := new(MessageStatus)
	*p = x
	return p
}

func (x MessageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_message_ext_proto_enumTypes[3].Descriptor()
}

func (MessageStatus) Type() protoreflect.EnumType {
	return &file_message_ext_proto_enumTypes[3]
}

func (x MessageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageStatus.Descriptor instead.
func (MessageStatus) EnumDescriptor() ([]byte, []int) {
	return file_message_ext_proto_rawDescGZIP(), []int{3}
}

// 单条消息投递内容（估算大约100个字节）,todo 通知栏提醒
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             int32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`                                                                  // 推送码
	Content          []byte             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                                                             // 推送内容
	Seq              int64              `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`                                                                    // 用户消息发送序列号
	SendTime         int64              `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`                                          // 消息发送时间戳，精确到毫秒
	Status           MessageStatus      `protobuf:"varint,5,opt,name=status,proto3,enum=pb.MessageStatus" json:"status,omitempty"`                                        // 消息状态
	TargetId         int64              `protobuf:"varint,6,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`                                          //目标用户id
	SenderId         int64              `protobuf:"varint,7,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`                                          //来自于用户id
	ConversationType ChatType           `protobuf:"varint,8,opt,name=conversation_type,json=conversationType,proto3,enum=pb.ChatType" json:"conversation_type,omitempty"` //会话类型
	MsgContentType   MessageContentType `protobuf:"varint,9,opt,name=msg_content_type,json=msgContentType,proto3,enum=pb.MessageContentType" json:"msg_content_type,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_ext_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetSeq() int64 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Message) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *Message) GetStatus() MessageStatus {
	if x != nil {
		return x.Status
	}
	return MessageStatus_MS_NORMAL
}

func (x *Message) GetTargetId() int64 {
	if x != nil {
		return x.TargetId
	}
	return 0
}

func (x *Message) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *Message) GetConversationType() ChatType {
	if x != nil {
		return x.ConversationType
	}
	return ChatType_DEFAULT
}

func (x *Message) GetMsgContentType() MessageContentType {
	if x != nil {
		return x.MsgContentType
	}
	return MessageContentType_MCT_TEXT
}

var File_message_ext_proto protoreflect.FileDescriptor

var file_message_ext_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xc8, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x73, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x40, 0x0a, 0x10, 0x6d, 0x73, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0e, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2a, 0xc6, 0x04, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x49, 0x4e,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x54, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10,
	0x04, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x54, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42,
	0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x53,
	0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x91, 0x4e, 0x12, 0x10, 0x0a,
	0x0b, 0x50, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x92, 0x4e, 0x12,
	0x11, 0x0a, 0x0c, 0x50, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x53, 0x10,
	0x93, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x10, 0xec, 0x07, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x54, 0x5f, 0x46, 0x52,
	0x49, 0x45, 0x4e, 0x44, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10,
	0x9a, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x50, 0x54, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x9b, 0x4e, 0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x54, 0x5f, 0x46,
	0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x41, 0x47, 0x52, 0x45, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f,
	0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x9c, 0x4e, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x54, 0x5f,
	0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x10, 0x9d, 0x4e, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e,
	0x44, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x53, 0x10, 0x9e, 0x4e,
	0x12, 0x22, 0x0a, 0x1d, 0x50, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x47, 0x45,
	0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x53, 0x10, 0x9f, 0x4e, 0x12, 0x1f, 0x0a, 0x19, 0x50, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x53, 0x10, 0xa1, 0x9c, 0x01, 0x12, 0x1a, 0x0a, 0x14, 0x50, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x5f,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x53, 0x10, 0xa2, 0x9c,
	0x01, 0x12, 0x14, 0x0a, 0x0e, 0x50, 0x54, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x41, 0x43, 0x4b, 0x10, 0xb1, 0xea, 0x01, 0x12, 0x1b, 0x0a, 0x15, 0x50, 0x54, 0x5f, 0x52, 0x4f,
	0x4f, 0x4d, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4c, 0x49, 0x53, 0x54,
	0x10, 0xc1, 0xb8, 0x02, 0x12, 0x19, 0x0a, 0x13, 0x50, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xc2, 0xb8, 0x02, 0x12,
	0x19, 0x0a, 0x13, 0x50, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xc3, 0xb8, 0x02, 0x12, 0x17, 0x0a, 0x11, 0x50, 0x54,
	0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10,
	0xc4, 0xb8, 0x02, 0x12, 0x16, 0x0a, 0x10, 0x50, 0x54, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x5f, 0x51,
	0x55, 0x49, 0x54, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0xc5, 0xb8, 0x02, 0x2a, 0x47, 0x0a, 0x08, 0x43,
	0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x5f, 0x43,
	0x48, 0x41, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x43,
	0x48, 0x41, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x52, 0x4f,
	0x4f, 0x4d, 0x10, 0x03, 0x2a, 0x55, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x43,
	0x54, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x43, 0x54, 0x5f,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x4d, 0x43, 0x54, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x43, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x03, 0x2a, 0x2d, 0x0a, 0x0d, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09,
	0x4d, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d,
	0x53, 0x5f, 0x52, 0x45, 0x43, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x69,
	0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_ext_proto_rawDescOnce sync.Once
	file_message_ext_proto_rawDescData = file_message_ext_proto_rawDesc
)

func file_message_ext_proto_rawDescGZIP() []byte {
	file_message_ext_proto_rawDescOnce.Do(func() {
		file_message_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_ext_proto_rawDescData)
	})
	return file_message_ext_proto_rawDescData
}

var file_message_ext_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_message_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_ext_proto_goTypes = []interface{}{
	(PackageType)(0),        // 0: pb.PackageType
	(ChatType)(0),           // 1: pb.ChatType
	(MessageContentType)(0), // 2: pb.MessageContentType
	(MessageStatus)(0),      // 3: pb.MessageStatus
	(*Message)(nil),         // 4: pb.Message
}
var file_message_ext_proto_depIdxs = []int32{
	3, // 0: pb.Message.status:type_name -> pb.MessageStatus
	1, // 1: pb.Message.conversation_type:type_name -> pb.ChatType
	2, // 2: pb.Message.msg_content_type:type_name -> pb.MessageContentType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_ext_proto_init() }
func file_message_ext_proto_init() {
	if File_message_ext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_message_ext_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_ext_proto_goTypes,
		DependencyIndexes: file_message_ext_proto_depIdxs,
		EnumInfos:         file_message_ext_proto_enumTypes,
		MessageInfos:      file_message_ext_proto_msgTypes,
	}.Build()
	File_message_ext_proto = out.File
	file_message_ext_proto_rawDesc = nil
	file_message_ext_proto_goTypes = nil
	file_message_ext_proto_depIdxs = nil
}
