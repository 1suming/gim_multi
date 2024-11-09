package model

import (
	"gim/pkg/protocol/pb"
	"gim/pkg/util"
	"time"
)

// Message 消息
type Message struct {
	Id         int64     // 自增主键
	UserId     int64     // 所属类型id
	RequestId  int64     // 请求id
	Code       int32     // 推送码
	Content    []byte    // 推送内容
	Seq        int64     // 消息同步序列
	SendTime   time.Time // 消息发送时间
	Status     int32     // 创建时间
	TargetId   int64     //对方Id //@ms;
	SenderId   int64     //发送者id
	StrContent string    //消息内容string格式
}

func (m *Message) MessageToPB() *pb.Message {
	return &pb.Message{
		Code:     m.Code,
		Content:  m.Content,
		Seq:      m.Seq,
		SendTime: util.UnixMilliTime(m.SendTime),
		Status:   pb.MessageStatus(m.Status),

		TargetId: m.TargetId,
		SenderId: m.SenderId,

		//int64 from_user_id=6;//来自于用户id
		//int64 to_user_id=7;//目标用户id
		//
		//MessageConversationType conversation_type=8;//会话类型
		//
	}
}

func MessagesToPB(messages []Message) []*pb.Message {
	pbMessages := make([]*pb.Message, 0, len(messages))
	for i := range messages {
		pbMessage := messages[i].MessageToPB()
		if pbMessages != nil {
			pbMessages = append(pbMessages, pbMessage)
		}
	}
	return pbMessages
}
