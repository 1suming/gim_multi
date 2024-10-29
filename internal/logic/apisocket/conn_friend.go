package apisocket

import (
	"context"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
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

	_, err = logicExtServer.AddFriend(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &req)

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
