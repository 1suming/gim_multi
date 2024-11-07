package dto

import (
	"gim/pkg/protocol/pb"
	"time"
)

type SaveOrUpdateRecentContactDTO struct {
	ConversationType   int8
	OwnerUid           int64
	TargetId           int64
	LastMessageId      int64
	LastMessageContent string
	LastTime           time.Time

	LastTargetMessageId int64 //@ms:target的lastMessageId
}

type UserRecentConversationSingle struct {
	ConversationType   int8
	OwnerUid           int64
	TargetId           int64
	LastMessageId      int64
	LastMessageContent string
	LastTime           time.Time
	UnreadCnt          int64
}

type UserRecentConversationAll struct {
	TotalUnread   int64
	Conversations []*UserRecentConversationSingle
}

func (m *UserRecentConversationSingle) MessageToPB() *pb.UserRecentConversationSingle {
	return &pb.UserRecentConversationSingle{
		ConversationType:   int32(m.ConversationType),
		OwnerUid:           m.OwnerUid,
		TargetId:           m.TargetId,
		LastMessageId:      m.LastMessageId,
		LastMessageContent: m.LastMessageContent,
		LastTime:           m.LastTime.Unix(),
		UnreadCnt:          m.UnreadCnt,
	}
}

func UserRecentConversationsToPB(messages []*UserRecentConversationSingle) []*pb.UserRecentConversationSingle {
	pbMessages := make([]*pb.UserRecentConversationSingle, 0, len(messages))
	for i := range messages {
		pbMessage := messages[i].MessageToPB()
		if pbMessages != nil {
			pbMessages = append(pbMessages, pbMessage)
		}
	}
	return pbMessages
}
