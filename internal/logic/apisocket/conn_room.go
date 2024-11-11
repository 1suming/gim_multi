package apisocket

import (
	"context"
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
		c.Send(pb.PackageType_PT_ROOM_JOIN_ROOM, input.RequestId, nil, err)
		return
	}
	//roomId := subscribeRoom.RoomId
	isSendMsg := true
	err = SubscribedRoom(c, subscribeRoom.RoomId, isSendMsg)
	//c.Send(pb.PackageType_PT_SUBSCRIBE_ROOM, input.RequestId, nil, nil)
	c.Send(pb.PackageType_PT_ROOM_JOIN_ROOM, input.RequestId, nil, err)

	////发送进房消息
	//deviceId, userId := c.DeviceId, c.UserId
	//
	//userInfo, err := userRepo.UserRepo.Get(userId)
	//if err != nil {
	//	logger.Logger.Error("Get err", zap.Error(err))
	//	c.Send(pb.PackageType_PT_ROOM_JOIN_ROOM, input.RequestId, nil, err)
	//	return
	//}
	//joinRoomContent := "欢迎" + " " + userInfo.Nickname + " 加入聊天室"
	//sendMessageReq := pb.SendMessageReq{
	//	ChatType:   pb.ChatType_CHAT_ROOM,
	//	ReceiverId: roomId,
	//
	//	Content: []byte(joinRoomContent),
	//
	//	SendTime:       util.UnixMilliTime(time.Now()),
	//	MsgContentType: pb.MessageContentType_MCT_NOTIFICATION, //发送通知类消息
	//}
	//userId = 0 //0代表系统
	//room.App.SendRoomMessage(context.TODO(), deviceId, userId, &sendMessageReq)

	//_, err = rpc.GetLogicIntClient().SubscribeRoom(context.TODO(), &pb.SubscribeRoomReq{
	//	UserId:   c.UserId,
	//	DeviceId: c.DeviceId,
	//	RoomId:   subscribeRoom.RoomId,
	//	Seq:      subscribeRoom.Seq,
	//	ConnAddr: config.Config.ConnectLocalAddr,
	//})
	//req := &pb.SubscribeRoomReq{
	//	UserId:   c.UserId,
	//	DeviceId: c.DeviceId,
	//	RoomId:   subscribeRoom.RoomId,
	//	Seq:      subscribeRoom.Seq,
	//	ConnAddr: config.Config.ConnectLocalAddr,
	//}
	//下面这个函数 仅仅是发送历史消息，
	//room.App.SubscribeRoom(context.TODO(), req)

	//if err != nil {
	//	logger.Logger.Error("SubscribedRoom error", zap.Error(err))
	//}
}
func Handle_QuitRoom(c *Conn, input *pb.Input) {
	var roomReq pb.QuitRoomReq
	err := proto.Unmarshal(input.Data, &roomReq)
	if err != nil {
		logger.Sugar.Error(err)
		c.Send(pb.PackageType_PT_ROOM_QUITROOM, input.RequestId, nil, err)
		return
	}
	roomId := roomReq.RoomId
	if c.RoomId != roomId {
		logger.Logger.Error("Handle_QuitRoom", zap.Any("conn.room id !=req.roomid", c.RoomId))
		//return
	}
	isSendMsg := true
	err = UnSubscribedRoom(c, isSendMsg)
	c.Send(pb.PackageType_PT_ROOM_QUITROOM, input.RequestId, nil, err)

	//_, err = rpc.GetLogicIntClient().SubscribeRoom(context.TODO(), &pb.SubscribeRoomReq{
	//	UserId:   c.UserId,
	//	DeviceId: c.DeviceId,
	//	RoomId:   subscribeRoom.RoomId,
	//	Seq:      subscribeRoom.Seq,
	//	ConnAddr: config.Config.ConnectLocalAddr,
	//})
	//req := &pb.SubscribeRoomReq{
	//	UserId:   c.UserId,
	//	DeviceId: c.DeviceId,
	//	RoomId:   subscribeRoom.RoomId,
	//	Seq:      subscribeRoom.Seq,
	//	ConnAddr: config.Config.ConnectLocalAddr,
	//}
	//下面这个函数 仅仅是发送历史消息，
	//room.App.SubscribeRoom(context.TODO(), req)

	//if err != nil {
	//	logger.Logger.Error("SubscribedRoom error", zap.Error(err))
	//}
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
