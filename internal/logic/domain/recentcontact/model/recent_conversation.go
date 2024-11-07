package model

import "time"

// ImRecentConversation undefined
type ImRecentConversation struct {
	ConversationType   int8      `json:"conversation_type" gorm:"conversation_type"` // 1:对个人；2 room 3. group
	OwnerUid           int64     `json:"owner_uid" gorm:"owner_uid"`
	TargetId           int64     `json:"target_id" gorm:"target_id"`
	LastMessageId      int64     `json:"last_message_id" gorm:"last_message_id"`
	LastMessageContent string    `json:"last_message_content" gorm:"last_message_content"`
	LastTime           time.Time `json:"last_time" gorm:"last_time"`

	UnreadCnt int64 `json:"unread_cnt" gorm:"unread_cnt"`
}

// TableName 表名称
func (*ImRecentConversation) TableName() string {
	return "im_recent_conversation"
}
