package room

import (
	"context"
	"gim/internal/logic/domain/room/repo"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const MaxSyncBufLen = 65536 // 最大字节数组长度

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

	return Service.SendRoomMessage(ctx, fromDeviceID, fromUserID, req)
}

func (*app) GetChatRoomList(ctx context.Context, userId int64) ([]*pb.ChatRoom, error) {
	rooms, err := repo.ChatRoomRepo.QueryAll()
	if err != nil {
		return nil, err
	}

	pbChatRooms := make([]*pb.ChatRoom, len(rooms))
	for i := range rooms {
		pbChatRooms[i] = rooms[i].ToProto()
	}
	return pbChatRooms, nil
}

/*
参考 ：func (*messageService) GetUserMessages(ctx context.Context, userId, seq int64, targetId int64, count int64) (*pb.GetUserMessagesResp, error) {
*/
func (*app) GetMessages(ctx context.Context, userId, seq int64, roomId int64, count int64) (*pb.GetUserMessagesResp, error) {
	logger.Logger.Info("room GetMessages", zap.Any("userId", userId), zap.Any("seq", seq), zap.Any("targetId", roomId), zap.Any("count", count))

	pbMessages, err := repo.MessageRepo.List(roomId, seq)

	length := len(pbMessages)
	hasMore := true
	resp := &pb.GetUserMessagesResp{Messages: pbMessages, HasMore: hasMore}
	bytes, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}

	// 如果字节数组大于一个包的长度，需要减少字节数组
	for len(bytes) > MaxSyncBufLen {
		length = length * 2 / 3
		resp = &pb.GetUserMessagesResp{Messages: pbMessages[0:length], HasMore: true}
		bytes, err = proto.Marshal(resp)
		if err != nil {
			return nil, err
		}
	}
	logger.Logger.Info("room GetMessages result len:", zap.Any("len", length))

	return resp, nil
}
