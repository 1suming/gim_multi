package proxy

import "gim/pkg/protocol/pb"

type IRoomDeliveMessageProxy interface {
	PushRoomMsg(roomId int64, message *pb.Message)
}

var RoomDeliveMessageProxy IRoomDeliveMessageProxy
