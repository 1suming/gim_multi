package proxy

import (
	"context"
	"gim/internal/logic/apisocket"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var MessageProxy messageProxy

type messageProxy interface {
	SendToUser(ctx context.Context, fromDeviceID, toUserID int64, message *pb.Message, isPersist bool) (int64, error)
}

func PushToUserBytes(ctx context.Context, toUserID int64, code int32, bytes []byte, isPersist bool) (int64, error) {
	message := pb.Message{
		Code:     code,
		Content:  bytes,
		SendTime: util.UnixMilliTime(time.Now()),
	}
	seq, err := MessageProxy.SendToUser(ctx, 0, toUserID, &message, isPersist)
	if err != nil {
		logger.Logger.Error("PushToUser", zap.Error(err))
		return 0, err
	}
	return seq, nil
}

func PushToUser(ctx context.Context, toUserID int64, code pb.PushCode, msg proto.Message, isPersist bool) (int64, error) {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		logger.Logger.Error("PushToUser", zap.Error(err))
		return 0, err
	}
	return PushToUserBytes(ctx, toUserID, int32(code), bytes, isPersist)
}

// DeliverMessage 投递消息
func DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}
	logger.Logger.Info("DeliverMessage func start", zap.Any("req", req))
	//// 获取设备对应的TCP连接
	conn := apisocket.GetConn(req.DeviceId)
	if conn == nil {
		logger.Logger.Warn("GetConn warn", zap.Int64("device_id", req.DeviceId))
		return resp, nil
	}

	if conn.DeviceId != req.DeviceId {
		logger.Logger.Warn("GetConn warn", zap.Int64("device_id", req.DeviceId))
		return resp, nil
	}
	logger.Logger.Info("devliveMsg: PackageType_PT_MESSAGE", zap.Any("req", req))
	conn.Send(pb.PackageType_PT_MESSAGE, grpclib.GetCtxRequestId(ctx), req.Message, nil)

	return resp, nil
}
