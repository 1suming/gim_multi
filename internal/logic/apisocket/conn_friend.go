package apisocket

import (
	"context"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/rpc"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (c *Conn) Handle_AddFriend(input *pb.Input) {
	var req pb.AddFriendReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_AddFriend", zap.Error(err))
		return
	}
	logger.Logger.Info(" Handle_AddFriend", zap.Any("req", req))

	deviceId, userId, token := c.DeviceId, c.UserId, c.LoginToken

	_, err = logicExtServer.AddFriend(grpclib.ContextWithUserInfo(context.TODO(), input.RequestId, deviceId, userId, token), &req)

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

	c.Send(pb.PackageType_PT_FRIEND_ADD_FRIEND, input.RequestId, nil, err)

}
func (c *Conn) Handle_SearchUser(input *pb.Input) {
	logger.Logger.Info("Handle_SearchUser", zap.Any("input", input))
	var req pb.SearchUserReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("handle_SearchUser", zap.Error(err))
		return
	}
	deviceId, userId, token := c.DeviceId, c.UserId, c.LoginToken
	resp, err := rpc.GetBusinessExtClient().SearchUser(grpclib.ContextWithUserInfo(context.TODO(), input.RequestId, deviceId, userId, token), &req)
	if err != nil {
		logger.Logger.Error("handle_SearchUser", zap.Error(err))
	}
	//return resp, err

	logger.Logger.Info(" handle_SearchUser", zap.Any("resp", resp))

	c.Send(pb.PackageType_PT_SEARCH_USER, input.RequestId, resp, err)

}
