package friend

import (
	"context"
	"gim/pkg/protocol/pb"
	"time"
)

type app struct{}

var App = new(app)

// List 获取好友列表
func (s *app) List(ctx context.Context, userId int64) ([]*pb.Friend, error) {
	return Service.List(ctx, userId)
}

func (s *app) GetFriendReqs(ctx context.Context, userId int64, isSendFriend bool) ([]*pb.FriendReq, error) {
	return Service.GetFriendReqs(ctx, userId, isSendFriend)
}

// AddFriend 添加好友
func (*app) AddFriend(ctx context.Context, userId, friendId int64, remarks, description string) error {
	return Service.AddFriend(ctx, userId, friendId, remarks, description)
}

// AgreeAddFriend 同意添加好友
func (*app) AgreeAddFriend(ctx context.Context, userId, friendId int64, remarks string) error {
	return Service.AgreeAddFriend(ctx, userId, friendId, remarks, FriendStatusAgree)
}

// @ms:
func (*app) HandleAddFriendRequest(ctx context.Context, userId, friendId int64, remarks string, status pb.FriendReqStatus) error {
	if status == pb.FriendReqStatus_FRIEND_REQ_STATUS_AGREE {
		return Service.AgreeAddFriend(ctx, userId, friendId, remarks, FriendStatusAgree)
	} else {
		return Service.AgreeAddFriend(ctx, userId, friendId, remarks, FriendStatusRefuse)
	}
}

// SetFriend 设置好友信息
func (*app) SetFriend(ctx context.Context, userId int64, req *pb.SetFriendReq) error {
	friend, err := Repo.Get(userId, req.FriendId)
	if err != nil {
		return err
	}
	if friend == nil {
		return nil
	}

	friend.Remarks = req.Remarks
	friend.Extra = req.Extra
	friend.UpdateTime = time.Now()

	err = Repo.Save(friend)
	if err != nil {
		return err
	}
	return nil
}

// SendToFriend 消息发送至好友
func (*app) SendToFriend(ctx context.Context, fromDeviceID, fromUserID int64, req *pb.SendMessageReq) (int64, int64, error) {
	return Service.SendToFriend(ctx, fromDeviceID, fromUserID, req)
}
