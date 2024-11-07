package apisocket

import (
	ctx "context"
	"gim/internal/logic/domain/friend"
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

func (c *Conn) Handle_AddFriend(input *pb.Input) error {
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

// SendMessageToFriend 发送好友消息
func (c *Conn) Handle_SendMessageToFriend(input *pb.Input) error {

	var req pb.SendMessageReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_SendMessageToFriend", zap.Error(err))
		c.Send(pb.PackageType_PT_FRIEND_SEND_MSG_TO_FRIEND, input.RequestId, nil, err)
		return err
	}
	logger.Logger.Info(" Handle_SendMessageToFriend", zap.Any("req", req))

	deviceId, userId := c.DeviceId, c.UserId
	seq, err := friend.App.SendToFriend(context.TODO(), deviceId, userId, &req)

	resp, err := &pb.SendMessageResp{Seq: seq}, nil

	sendTime := util.GetNowTime()
	saveOrUpdateRecentContactDTO := dto.SaveOrUpdateRecentContactDTO{
		ConversationType:   int8(pb.MessageConversationType_FRIEND),
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
		c.Send(pb.PackageType_PT_FRIEND_SEND_MSG_TO_FRIEND, input.RequestId, nil, err)
		return err
	}
	c.Send(pb.PackageType_PT_FRIEND_SEND_MSG_TO_FRIEND, input.RequestId, resp, err)
	return nil
}
