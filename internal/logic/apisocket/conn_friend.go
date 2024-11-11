package apisocket

import (
	ctx "context"
	"gim/internal/logic/domain/friend"
	"gim/internal/logic/domain/message"
	"gim/pkg/dto"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/util"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	recentContactService "gim/internal/logic/domain/recentcontact/service"
)

func Handle_AddFriend(c *Conn, input *pb.Input) error {
	var req pb.AddFriendReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_AddFriend", zap.Error(err))
		c.Send(pb.PackageType_PT_FRIEND_ADD_FRIEND, input.RequestId, nil, err)
		return err
	}
	logger.Logger.Info(" Handle_AddFriend", zap.Any("req", req))

	//deviceId, userId, token := c.DeviceId, c.UserId, c.LoginToken
	//
	//_, err = logicExtServer.AddFriend(grpclib.ContextWithUserInfo(context.TODO(), input.RequestId, deviceId, userId, token), &req)

	//var message proto.Message
	//if err == nil {
	//	message = &pb.SyncOutput{Messages: resp.Messages, HasMore: resp.HasMore}
	//}
	//c.Send(pb.PackageType_PT_SYNC, input.RequestId, message, err)
	//
	//var message proto.Message
	//if err == nil {
	//	message = &pb.SyncOutput{Messages: resp.Messages, HasMore: resp.HasMore}
	//}
	//c.Send(pb.PackageType_PT_SYNC, input.RequestId, message, err)

	userId := c.UserId
	resp := new(emptypb.Empty)
	err = friend.App.AddFriend(ctx.TODO(), userId, req.FriendId, req.Remarks, req.Description)

	c.Send(pb.PackageType_PT_FRIEND_ADD_FRIEND, input.RequestId, resp, err)
	return nil

}
func Handle_AgreeAddFriend(c *Conn, input *pb.Input) error {
	var req pb.AgreeAddFriendReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_AgreeAddFriend", zap.Error(err))
		c.Send(pb.PackageType_PT_FRIEND_AGREE_ADD_FRIEND, input.RequestId, nil, err)
		return err
	}
	logger.Logger.Info(" Handle_AgreeAddFriend", zap.Any("req", req))

	userId := c.UserId
	resp := new(emptypb.Empty)
	err = friend.App.HandleAddFriendRequest(ctx.TODO(), userId, req.UserId, req.Remarks, req.Status)

	c.Send(pb.PackageType_PT_FRIEND_AGREE_ADD_FRIEND, input.RequestId, resp, err)
	return nil

}

func Handle_GetFriendReqs(c *Conn, input *pb.Input) error {
	var req pb.GetFriendRequestsReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_GetFriendReqs", zap.Error(err))
		c.Send(pb.PackageType_PT_FRIEND_GET_FRIEND_REQUESTS, input.RequestId, nil, err)
		return err
	}
	logger.Logger.Info(" Handle_GetFriendReqs", zap.Any("req", nil))

	userId := c.UserId
	isSendFriend := false
	friendReqs, err := friend.App.GetFriendReqs(ctx.TODO(), userId, isSendFriend)
	resp := &pb.GetFriendRequestsResp{Requests: friendReqs}

	c.Send(pb.PackageType_PT_FRIEND_GET_FRIEND_REQUESTS, input.RequestId, resp, err)
	return nil

}

func Handle_GetFriends(c *Conn, input *pb.Input) error {
	//var req pb.AgreeAddFriendReq
	//err := proto.Unmarshal(input.Data, &req)
	//if err != nil {
	//	logger.Logger.Error("Handle_AgreeAddFriend", zap.Error(err))
	//	c.Send(pb.PackageType_PT_FRIEND_ADD_FRIEND, input.RequestId, nil, err)
	//	return err
	//}
	logger.Logger.Info(" Handle_GetFriends", zap.Any("req", nil))

	userId := c.UserId
	friends, err := friend.App.List(ctx.TODO(), userId)
	resp := &pb.GetFriendsResp{Friends: friends}

	c.Send(pb.PackageType_PT_FRIEND_GET_FRIENDS, input.RequestId, resp, err)
	return nil

}

// SendMessageToFriend 发送好友消息
func Handle_SendMessageToFriend(c *Conn, input *pb.Input) error {

	var req pb.SendMessageReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_SendMessageToFriend", zap.Error(err))
		c.Send(pb.PackageType_PT_SEND_MESSAGE, input.RequestId, nil, err)
		return err
	}
	logger.Logger.Info(" Handle_SendMessageToFriend", zap.Any("req", req))

	deviceId, userId := c.DeviceId, c.UserId
	seq, targetSeq, err := friend.App.SendToFriend(context.TODO(), deviceId, userId, &req)
	_ = targetSeq

	resp, err := &pb.SendMessageResp{Seq: seq}, nil

	sendTime := util.GetNowTime()
	saveOrUpdateRecentContactDTO := dto.SaveOrUpdateRecentContactDTO{
		ConversationType:   int8(pb.ChatType_SINGLE_CHAT),
		LastMessageContent: string(req.Content),
		LastMessageId:      seq,
		TargetId:           req.ReceiverId,
		LastTime:           sendTime,
		OwnerUid:           userId,
	}
	//保存或更新会话信息
	err = recentContactService.RecentConversationService.SaveOrUpdate(context.TODO(), &saveOrUpdateRecentContactDTO)
	if err != nil {
		logger.Logger.Error("Handle_SendMessageToFriend", zap.Error(err))
		c.Send(pb.PackageType_PT_SEND_MESSAGE, input.RequestId, nil, err)
		return err
	}
	c.Send(pb.PackageType_PT_SEND_MESSAGE, input.RequestId, resp, err)
	return nil
}

func Handle_GetUserMessages(c *Conn, input *pb.Input) {
	var req pb.GetUserMessagesReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	if req.ConversationType == pb.ChatType_CHAT_ROOM {
		Handle_GetChatRoomMessages(c, input)
		return
	}

	//
	//resp, err := rpc.GetLogicIntClient().Sync(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.SyncReq{
	//	UserId:   c.UserId,
	//	DeviceId: c.DeviceId,
	//	Seq:      sync.Seq,
	//})
	//
	//req := pb.GetUserConversationsReq{
	//	UserId:   c.UserId,
	//	DeviceId: c.DeviceId,
	//	Seq:      sync.Seq,
	//}

	resp, err := message.App.GetUserMessages(context.TODO(), req.OwnerUid, req.Seq, req.TargetId, req.Count)

	var message proto.Message
	if err == nil {
		message = &pb.GetUserMessagesResp{Messages: resp.Messages, HasMore: resp.HasMore}
	}
	c.Send(pb.PackageType_PT_GET_USER_MESSAGES, input.RequestId, message, err)
}
