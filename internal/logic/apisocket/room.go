package apisocket

import (
	"container/list"
	"context"
	"gim/internal/logic/domain/room"
	userRepo "gim/internal/logic/domain/user/repo"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/util"
	"go.uber.org/zap"
	"sync"
	"time"
)

var RoomsManager sync.Map

type SRoomApp struct{}

var RoomApp = new(SRoomApp)

// SubscribedRoom 订阅房间
func SubscribedRoom(conn *Conn, roomId int64, isSendMessage bool) error {
	logger.Logger.Info("SubscribedRoom", zap.Any("userid", conn.UserId), zap.Any("deviceID", conn.DeviceId))
	if roomId == conn.RoomId {
		return nil
	}

	oldRoomId := conn.RoomId
	// 取消订阅
	if oldRoomId != 0 {
		value, ok := RoomsManager.Load(oldRoomId)
		if !ok {
			return nil
		}
		room := value.(*Room)
		room.Unsubscribe(conn)

		if room.Conns.Front() == nil {
			RoomsManager.Delete(oldRoomId)
		}
		//@ms:这里必须要return return
	}

	// 订阅
	if roomId != 0 {
		value, ok := RoomsManager.Load(roomId)
		var room *Room
		if !ok {
			room = NewRoom(roomId)
			RoomsManager.Store(roomId, room)
		} else {
			room = value.(*Room)
		}
		room.Subscribe(conn)

	}

	if isSendMessage {

		//发送进房消息
		deviceId, userId := conn.DeviceId, conn.UserId

		userInfo, err := userRepo.UserRepo.Get(userId)
		if err != nil {
			logger.Logger.Error("Get err", zap.Error(err))

			return err
		}
		joinRoomContent := "欢迎" + " " + userInfo.Nickname + " 加入聊天室"
		sendMessageReq := pb.SendMessageReq{
			ChatType:   pb.ChatType_CHAT_ROOM,
			ReceiverId: roomId,

			Content: []byte(joinRoomContent),

			SendTime:       util.UnixMilliTime(time.Now()),
			MsgContentType: pb.MessageContentType_MCT_NOTIFICATION, //发送通知类消息
		}
		userId = 0 //0代表系统
		room.App.SendRoomMessage(context.TODO(), deviceId, userId, &sendMessageReq)
	}
	return nil
}
func UnSubscribedRoom(conn *Conn, isSendMessage bool) error {
	logger.Logger.Info("UnSubscribedRoom", zap.Any("userid", conn.UserId), zap.Any("deviceID", conn.DeviceId))
	oldRoomId := conn.RoomId
	if oldRoomId == 0 {
		return nil
	}
	// 取消订阅
	if oldRoomId != 0 {
		value, ok := RoomsManager.Load(oldRoomId)
		if !ok {
			return nil
		}
		room := value.(*Room)
		room.Unsubscribe(conn)

		if room.Conns.Front() == nil {
			RoomsManager.Delete(oldRoomId)
		}
		//@ms:这里必须要return return
	}

	//发送进房消息
	deviceId, userId := conn.DeviceId, conn.UserId

	userInfo, err := userRepo.UserRepo.Get(userId)
	if err != nil {
		logger.Logger.Error("Get err", zap.Error(err))
		//conn.Send(pb.PackageType_PT_ROOM_QUITROOM, nil, nil, err)
		return nil
	}
	joinRoomContent := userInfo.Nickname + " 离开聊天室"
	sendMessageReq := pb.SendMessageReq{
		ChatType:   pb.ChatType_CHAT_ROOM,
		ReceiverId: oldRoomId,

		Content: []byte(joinRoomContent),

		SendTime:       util.UnixMilliTime(time.Now()),
		MsgContentType: pb.MessageContentType_MCT_NOTIFICATION, //发送通知类消息
	}
	userId = 0 //0代表系统
	room.App.SendRoomMessage(context.TODO(), deviceId, userId, &sendMessageReq)

	return nil
}

// PushRoom 房间消息推送
func PushRoom(roomId int64, message *pb.Message) {
	value, ok := RoomsManager.Load(roomId)
	if !ok {
		return
	}

	value.(*Room).Push(message)
}

// @ms:
func (*SRoomApp) PushRoomMsg(roomId int64, message *pb.Message) {
	value, ok := RoomsManager.Load(roomId)
	if !ok {
		return
	}

	value.(*Room).Push(message)
}

type Room struct {
	RoomId int64      // 房间ID
	Conns  *list.List // 订阅房间消息的连接
	lock   sync.RWMutex
}

func NewRoom(roomId int64) *Room {
	return &Room{
		RoomId: roomId,
		Conns:  list.New(),
	}
}

// Subscribe 订阅房间
func (r *Room) Subscribe(conn *Conn) {
	r.lock.Lock()
	defer r.lock.Unlock()

	conn.Element = r.Conns.PushBack(conn)
	conn.RoomId = r.RoomId
}

// Unsubscribe 取消订阅
func (r *Room) Unsubscribe(conn *Conn) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.Conns.Remove(conn.Element)
	conn.Element = nil
	conn.RoomId = 0
}

// Push 推送消息到房间
func (r *Room) Push(message *pb.Message) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	element := r.Conns.Front()
	for {
		conn := element.Value.(*Conn)
		logger.Logger.Debug("推送房间消息Push", zap.Any("userid:", conn.UserId), zap.Any("deviceID:", conn.DeviceId))
		conn.Send(pb.PackageType_PT_MESSAGE, 0, message, nil)

		element = element.Next()
		if element == nil {
			break
		}
	}
}
