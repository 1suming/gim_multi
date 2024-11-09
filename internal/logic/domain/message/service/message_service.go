package service

import (
	"context"
	"gim/internal/logic/domain/message/model"
	"gim/internal/logic/domain/message/repo"
	"gim/internal/logic/proxy"
	"gim/pkg/grpclib"
	"gim/pkg/grpclib/picker"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/rpc"
	"gim/pkg/util"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const MessageLimit = 50 // 最大消息同步数量

const MaxSyncBufLen = 65536 // 最大字节数组长度

type messageService struct {
}

var MessageService = new(messageService)

// Sync 消息同步
func (*messageService) Sync(ctx context.Context, userId, seq int64) (*pb.SyncResp, error) {
	messages, hasMore, err := MessageService.ListByUserIdAndSeq(ctx, userId, seq)
	if err != nil {
		return nil, err
	}
	pbMessages := model.MessagesToPB(messages)
	length := len(pbMessages)

	resp := &pb.SyncResp{Messages: pbMessages, HasMore: hasMore}
	bytes, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}

	// 如果字节数组大于一个包的长度，需要减少字节数组
	for len(bytes) > MaxSyncBufLen {
		length = length * 2 / 3
		resp = &pb.SyncResp{Messages: pbMessages[0:length], HasMore: true}
		bytes, err = proto.Marshal(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// Sync 消息同步
func (*messageService) GetUserMessages(ctx context.Context, userId, seq int64, targetId int64, count int64) (*pb.GetUserMessagesResp, error) {
	messages, hasMore, err := MessageService.ListByUserIdAndSeqAndTargetId(ctx, userId, seq, targetId, count)
	if err != nil {
		return nil, err
	}
	pbMessages := model.MessagesToPB(messages)
	length := len(pbMessages)

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

	return resp, nil
}

// ListByUserIdAndSeq 查询消息
func (*messageService) ListByUserIdAndSeq(ctx context.Context, userId, seq int64) ([]model.Message, bool, error) {
	var err error
	if seq == 0 {
		seq, err = DeviceAckService.GetMaxByUserId(ctx, userId)
		if err != nil {
			return nil, false, err
		}
	}
	return repo.MessageRepo.ListBySeq(userId, seq, MessageLimit)
}

func (*messageService) ListByUserIdAndSeqAndTargetId(ctx context.Context, userId, seq int64, targetId int64, count int64) ([]model.Message, bool, error) {
	var err error
	if seq == 0 {
		seq, err = DeviceAckService.GetMaxByUserId(ctx, userId)
		if err != nil {
			return nil, false, err
		}
	}
	//return repo.MessageRepo.ListBySeqAndTargetId(userId, seq, MessageLimit, targetId)

	return repo.MessageRepo.ListBySeqAndTargetId(userId, seq, count, targetId)
}

/*
sendToUser 会调用2次，一次是发给自己，一个是发给target
seq, err := proxy.MessageProxy.SendToUser(ctx, fromDeviceID, fromUserID, msg, true)

// 发给接收者
targetSeq, err = proxy.MessageProxy.SendToUser(ctx, fromDeviceID, req.ReceiverId, msg, true)

*/
// SendToUser 将消息发送给用户
func (*messageService) SendToUser(ctx context.Context, fromDeviceID, toUserID int64, message *pb.Message, isPersist bool) (int64, error) {
	logger.Logger.Debug("SendToUser",
		zap.Int64("request_id", grpclib.GetCtxRequestId(ctx)),
		zap.Int64("to_user_id", toUserID))
	var (
		seq int64 = 0
		err error
	)
	var targetId int64 //另一方ID
	if message.SenderId == toUserID {
		//是发送给自己
		targetId = message.TargetId

	} else {
		targetId = message.SenderId
	}

	if isPersist {
		seq, err = SeqService.GetUserNext(ctx, toUserID)
		if err != nil {
			return 0, err
		}
		message.Seq = seq

		selfMessage := model.Message{
			UserId:    toUserID,
			RequestId: grpclib.GetCtxRequestId(ctx),
			Code:      message.Code,
			Content:   message.Content,
			Seq:       seq,
			SendTime:  util.UnunixMilliTime(message.SendTime),
			Status:    int32(pb.MessageStatus_MS_NORMAL),

			TargetId:   targetId, //@ms:
			SenderId:   message.SenderId,
			StrContent: string(message.Content),
		}
		err = repo.MessageRepo.Save(selfMessage)
		if err != nil {
			logger.Sugar.Error(err)
			return 0, err
		}
	}

	// 查询用户在线设备
	devices, err := proxy.DeviceProxy.ListOnlineByUserId(ctx, toUserID)
	if err != nil {
		logger.Logger.Info("SendToUser", zap.Any("查找用于在线设备error", toUserID))
		logger.Sugar.Error(err)
		return 0, err
	}
	logger.Logger.Info("SendToUser", zap.Any("在线设备信息列表", devices))
	for i := range devices {
		// 消息不需要投递给发送消息的设备
		if fromDeviceID == devices[i].DeviceId {
			continue
		}
		logger.Logger.Info("SendToUser", zap.Any("开始发送 devices[i]", devices[i]))
		err = MessageService.SendToDevice(ctx, devices[i], message)
		if err != nil {
			logger.Sugar.Error(err, zap.Any("SendToUser error", devices[i]), zap.Error(err))
		}
	}
	return seq, nil
}

// SendToDevice 将消息发送给设备
func (m *messageService) SendToDevice(ctx context.Context, device *pb.Device, message *pb.Message) error {
	//_, err := rpc.GetConnectIntClient().DeliverMessage(picker.ContextWithAddr(ctx, device.ConnAddr), &pb.DeliverMessageReq{
	//	DeviceId: device.DeviceId,
	//	Message:  message,
	//})

	_, err := proxy.DeliveMessageProxy.DeliverMessage(picker.ContextWithAddr(ctx, device.ConnAddr), &pb.DeliverMessageReq{
		DeviceId: device.DeviceId,
		Message:  message,
	})

	if err != nil {
		logger.Logger.Error("SendToDevice error", zap.Error(err))
		return err
	}

	// todo 其他推送厂商
	return nil
}

func (*messageService) AddSenderInfo(sender *pb.Sender) {
	user, err := rpc.GetBusinessIntClient().GetUser(context.TODO(), &pb.GetUserReq{UserId: sender.UserId})
	if err == nil && user != nil {
		sender.AvatarUrl = user.User.AvatarUrl
		sender.Nickname = user.User.Nickname
		sender.Extra = user.User.Extra
	}
}
