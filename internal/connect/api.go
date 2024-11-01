package connect

import (
	"context"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConnIntServer struct {
	pb.UnsafeConnectIntServer
}

// DeliverMessage 投递消息
func (s *ConnIntServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}
	logger.Logger.Info("DeliverMessage func start", zap.Any("req", req))
	// 获取设备对应的TCP连接
	conn := GetConn(req.DeviceId)
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

	PushAll(req.Message) //TODO
	return resp, nil
}

//func (s *ConnIntServer) HandleTransferProto(ctx context.Context, req *pb.RouterMsgInput) (*emptypb.Empty, error) {
//
//}
