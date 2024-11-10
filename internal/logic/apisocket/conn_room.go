package apisocket

import (
	"context"
	"gim/config"
	"gim/internal/logic/domain/room"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
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
	//c.Send(pb.PackageType_PT_SUBSCRIBE_ROOM, input.RequestId, nil, nil)
	c.Send(pb.PackageType_PT_ROOM_JOIN_ROOM, input.RequestId, nil, nil)

	//_, err = rpc.GetLogicIntClient().SubscribeRoom(context.TODO(), &pb.SubscribeRoomReq{
	//	UserId:   c.UserId,
	//	DeviceId: c.DeviceId,
	//	RoomId:   subscribeRoom.RoomId,
	//	Seq:      subscribeRoom.Seq,
	//	ConnAddr: config.Config.ConnectLocalAddr,
	//})
	req := &pb.SubscribeRoomReq{
		UserId:   c.UserId,
		DeviceId: c.DeviceId,
		RoomId:   subscribeRoom.RoomId,
		Seq:      subscribeRoom.Seq,
		ConnAddr: config.Config.ConnectLocalAddr,
	}

	room.App.SubscribeRoom(context.TODO(), req)

	if err != nil {
		logger.Logger.Error("SubscribedRoom error", zap.Error(err))
	}
}
func Handle_GetRoomList(c *Conn, input *pb.Input) error {
	var req pb.GetRoomListReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_GetRoomList", zap.Error(err))
		c.Send(pb.PackageType_PT_ROOM_GET_ROOM_LIST, input.RequestId, nil, err)
		return err
	}
	logger.Logger.Info(" Handle_GetRoomList", zap.Any("req", req))

	deviceId, userId := c.DeviceId, c.UserId
	_ = deviceId
	pbRooms, err := room.App.GetChatRoomList(context.TODO(), userId) // ([]*pb.ChatRoom, error) {
	if err != nil {
		logger.Logger.Error("Handle_GetRoomList", zap.Error(err))
		c.Send(pb.PackageType_PT_ROOM_GET_ROOM_LIST, input.RequestId, nil, err)
		return err
	}
	resp := &pb.GetRoomListResp{
		Rooms: pbRooms,
	}

	c.Send(pb.PackageType_PT_ROOM_GET_ROOM_LIST, input.RequestId, resp, err)
	return nil
}

// 获取chatroom消息，类似 SubscribeRoom
func Handle_GetChatRoomMessages(c *Conn, input *pb.Input) {
	logger.Logger.Info("Handle_GetChatRoomMessages")
	var req pb.GetUserMessagesReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	//roomId := req.TargetId
	//

	resp, err := room.App.GetMessages(context.TODO(), req.OwnerUid, req.Seq, req.TargetId, req.Count)

	var message proto.Message
	if err == nil {
		message = &pb.GetUserMessagesResp{Messages: resp.Messages, HasMore: resp.HasMore}
	}
	c.Send(pb.PackageType_PT_GET_USER_MESSAGES, input.RequestId, message, err)

}
