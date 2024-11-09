package apisocket

import (
	"context"
	"gim/config"
	"gim/internal/logic/domain/room"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/rpc"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func Handle_SendMsgToRoom(c *Conn, input *pb.Input) error {
	var req pb.SendMessageReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_SendMsgToRoom", zap.Error(err))
		c.Send(pb.PackageType_PT_SEND_MESSAGE, input.RequestId, nil, err)
		return err
	}
	logger.Logger.Info(" Handle_SendMsgToRoom", zap.Any("req", req))

	deviceId, userId := c.DeviceId, c.UserId

	err = room.App.SendRoomMessage(context.TODO(), deviceId, userId, &req)
	return err
}

// SubscribedRoom 订阅房间
func Handle_SubscribedRoom(c *Conn, input *pb.Input) {
	var subscribeRoom pb.SubscribeRoomInput
	err := proto.Unmarshal(input.Data, &subscribeRoom)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	SubscribedRoom(c, subscribeRoom.RoomId)
	c.Send(pb.PackageType_PT_SUBSCRIBE_ROOM, input.RequestId, nil, nil)
	_, err = rpc.GetLogicIntClient().SubscribeRoom(context.TODO(), &pb.SubscribeRoomReq{
		UserId:   c.UserId,
		DeviceId: c.DeviceId,
		RoomId:   subscribeRoom.RoomId,
		Seq:      subscribeRoom.Seq,
		ConnAddr: config.Config.ConnectLocalAddr,
	})
	if err != nil {
		logger.Logger.Error("SubscribedRoom error", zap.Error(err))
	}
}
