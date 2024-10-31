package apisocket

import (
	ctx "context"
	"gim/internal/logic/domain/friend"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
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

	c.Send(pb.PackageType_PT_FRIEND_SEND_MSG_TO_FRIEND, input.RequestId, resp, err)
	return nil
}
