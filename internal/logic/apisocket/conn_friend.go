package apisocket

import (
	ctx "context"
	"gim/internal/logic/domain/friend"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"go.uber.org/zap"
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
