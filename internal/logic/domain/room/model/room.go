package model

import (
	"gim/pkg/protocol/pb"
	"time"
)

// Group 群组
type ChatRoom struct {
	Id           int64     // 群组id
	Name         string    // 组名
	AvatarUrl    string    // 头像
	Introduction string    // 群简介
	MaxUserNum   int32     // 群组人数
	Extra        string    // 附加字段
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
}

// TableName 表名称
func (*ChatRoom) TableName() string {
	return "chatroom"
}

func (g *ChatRoom) ToProto() *pb.ChatRoom {
	if g == nil {
		return nil
	}

	return &pb.ChatRoom{
		ChatroomId:   g.Id,
		Name:         g.Name,
		AvatarUrl:    g.AvatarUrl,
		Introduction: g.Introduction,
		MaxUserMum:   g.MaxUserNum,
		Extra:        g.Extra,
		CreateTime:   g.CreateTime.Unix(),
		UpdateTime:   g.UpdateTime.Unix(),
	}
}
