package friend

import (
	"context"
	"gim/internal/logic/proxy"
	"gim/pkg/gerrors"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/rpc"
	"go.uber.org/zap"
	"time"
)

type service struct{}

var Service = new(service)

// List 获取好友列表
func (s *service) List(ctx context.Context, userId int64) ([]*pb.Friend, error) {
	friends, err := Repo.List(userId, FriendStatusAgree)
	if err != nil {
		return nil, err
	}

	userIds := make(map[int64]int32, len(friends))
	for i := range friends {
		userIds[friends[i].FriendId] = 0
	}
	resp, err := rpc.GetBusinessIntClient().GetUsers(ctx, &pb.GetUsersReq{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	var infos = make([]*pb.Friend, len(friends))
	for i := range friends {
		friend := pb.Friend{
			UserId:  friends[i].FriendId,
			Remarks: friends[i].Remarks,
			Extra:   friends[i].Extra,
		}

		user, ok := resp.Users[friends[i].FriendId]
		if ok {
			friend.Nickname = user.Nickname
			friend.Sex = user.Sex
			friend.AvatarUrl = user.AvatarUrl
			friend.UserExtra = user.Extra
		}
		infos[i] = &friend
	}

	return infos, nil
}

// AddFriend 添加好友
func (*service) AddFriend(ctx context.Context, userId, friendId int64, remarks, description string) error {
	friend, err := Repo.Get(userId, friendId)
	if err != nil {
		return err
	}
	if friend != nil {
		if friend.Status == FriendStatusApply {
			return nil
		}
		if friend.Status == FriendStatusAgree {
			return gerrors.ErrAlreadyIsFriend
		}
	}

	now := time.Now()
	err = Repo.Save(&Friend{
		UserId:     userId,
		FriendId:   friendId,
		Remarks:    remarks,
		Status:     FriendStatusApply,
		CreateTime: now,
		UpdateTime: now,
	})
	if err != nil {
		return err
	}

	resp, err := rpc.GetBusinessIntClient().GetUser(ctx, &pb.GetUserReq{UserId: userId})
	if err != nil {
		return err
	}

	_, err = proxy.PushToUser(ctx, friendId, pb.PushCode_PC_ADD_FRIEND, &pb.AddFriendPush{
		FriendId:    userId,
		Nickname:    resp.User.Nickname,
		AvatarUrl:   resp.User.AvatarUrl,
		Description: description,
	}, true)
	if err != nil {
		return err
	}
	return nil
}

// AgreeAddFriend 同意添加好友
func (*service) AgreeAddFriend(ctx context.Context, userId, friendId int64, remarks string, status int) error {
	friend, err := Repo.Get(friendId, userId)
	if err != nil {
		return err
	}
	if friend == nil {
		return gerrors.ErrBadRequest
	}
	if status == FriendStatusRefuse { //@ms:如果是拒绝，修改状态即可
		friend.Status = FriendStatusRefuse
		err = Repo.Save(friend)
		if err != nil {
			return err
		}
		return nil
	}
	if friend.Status == FriendStatusAgree {
		return nil
	}
	friend.Status = FriendStatusAgree
	err = Repo.Save(friend)
	if err != nil {
		return err
	}

	now := time.Now()
	err = Repo.Save(&Friend{
		UserId:     userId,
		FriendId:   friendId,
		Remarks:    remarks,
		Status:     FriendStatusAgree,
		CreateTime: now,
		UpdateTime: now,
	})
	if err != nil {
		return err
	}
	//@ms:不要
	//resp, err := rpc.GetBusinessIntClient().GetUser(ctx, &pb.GetUserReq{UserId: userId})
	//if err != nil {
	//	return err
	//}
	//
	//_, err = proxy.PushToUser(ctx, friendId, pb.PushCode_PC_AGREE_ADD_FRIEND, &pb.AgreeAddFriendPush{
	//	FriendId:  userId,
	//	Nickname:  resp.User.Nickname,
	//	AvatarUrl: resp.User.AvatarUrl,
	//}, true)
	//if err != nil {
	//	return err
	//}
	return nil
}

// SendToFriend 消息发送至好友
func (*service) SendToFriend(ctx context.Context, fromDeviceID, fromUserID int64, req *pb.SendMessageReq) (int64, int64, error) {
	senderSeq := int64(0)
	targetSeq := int64(0)
	//sender, err := rpc.GetSender(fromDeviceID, fromUserID)
	//if err != nil {
	//	return senderSeq, targetSeq, err
	//}
	//_ = sender
	// 发给发送者
	//push := pb.UserMessagePush{
	//	Sender:     sender,
	//	ReceiverId: req.ReceiverId,
	//	Content:    req.Content,
	//}
	//bytes, err := proto.Marshal(&push)

	bytes := req.Content //content本来就是bytes
	//if err != nil {
	//	return senderSeq, targetSeq, err
	//}

	msg := &pb.Message{
		Code:     int32(pb.PushCode_PC_USER_MESSAGE),
		Content:  bytes,
		SendTime: req.SendTime,

		SenderId:         fromUserID,     //来自于用户id
		TargetId:         req.ReceiverId, //目标用户id
		ConversationType: pb.ChatType_SINGLE_CHAT,
		MsgContentType:   pb.MessageContentType_MCT_TEXT,
	}
	logger.Logger.Info("SendToFriend", zap.Any("自身------开始", fromUserID))

	seq, err := proxy.MessageProxy.SendToUser(ctx, fromDeviceID, fromUserID, msg, true)
	if err != nil {
		return senderSeq, targetSeq, err
	}
	logger.Logger.Info("SendToFriend", zap.Any("自身------结束", fromUserID))

	logger.Logger.Info("SendToFriend", zap.Any("对方------开始", fromUserID))

	// 发给接收者
	targetSeq, err = proxy.MessageProxy.SendToUser(ctx, fromDeviceID, req.ReceiverId, msg, true)
	if err != nil {
		return senderSeq, targetSeq, err
	}
	logger.Logger.Info("SendToFriend", zap.Any("对方------结束", fromUserID))

	return seq, targetSeq, nil
}

/*
*
isSendFriend *  获取发送出去的好友请求
获取收到的好友请求
*/
func (s *service) GetFriendReqs(ctx context.Context, userId int64, isSendFriend bool) ([]*pb.FriendReq, error) {
	//friends, err := Repo.List(userId, FriendStatusAgree)
	var friends []Friend
	var err error

	friends, err = Repo.GetFriendReqs(userId, FriendStatusApply, isSendFriend)

	if err != nil {
		return nil, err
	}

	userIds := make(map[int64]int32, len(friends))
	for i := range friends {
		userIds[friends[i].FriendId] = 0
	}
	//resp, err := rpc.GetBusinessIntClient().GetUsers(ctx, &pb.GetUsersReq{UserIds: userIds})
	//if err != nil {
	//	return nil, err
	//}

	var infos = make([]*pb.FriendReq, len(friends))
	for i := range friends {

		createTime_unixMilli := friends[i].CreateTime.UnixMilli()
		friendReq := pb.FriendReq{

			Remarks:    friends[i].Remarks,
			CreateTime: createTime_unixMilli,
		}
		if isSendFriend == true {
			friendReq.FriendId = friends[i].FriendId
		} else {
			friendReq.FriendId = friends[i].UserId
		}

		//user, ok := resp.Users[friends[i].FriendId]
		//if ok {
		//	friend.Nickname = user.Nickname
		//	friend.Sex = user.Sex
		//	friend.AvatarUrl = user.AvatarUrl
		//	friend.UserExtra = user.Extra
		//}
		infos[i] = &friendReq
	}

	return infos, nil
}
