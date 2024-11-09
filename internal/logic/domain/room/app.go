package room

import (
	"context"
	"gim/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

// Push 推送房间消息
func (s *app) Push(ctx context.Context, req *pb.PushRoomReq) error {
	return Service.Push(ctx, req)
}

// SubscribeRoom 订阅房间
func (s *app) SubscribeRoom(ctx context.Context, req *pb.SubscribeRoomReq) error {
	return Service.SubscribeRoom(ctx, req)
}

// @ms:
// SendMessage 发送群组消息
func (s *app) SendRoomMessage(ctx context.Context, fromDeviceID, fromUserID int64, req *pb.SendMessageReq) error {

	return service.SendRoomMessage(ctx, fromDeviceID, fromUserID, req)
}
