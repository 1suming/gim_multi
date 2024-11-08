package connManage

import (
	"gim/internal/logic/apisocket"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"go.uber.org/zap"
	"sync"
)

var ConnsManager = sync.Map{}

// SetConn 存储
func SetConn(deviceId int64, conn *apisocket.Conn) {
	ConnsManager.Store(deviceId, conn)
}

// GetConn 获取
func GetConn(deviceId int64) *apisocket.Conn {
	logger.Logger.Debug("GetConn", zap.Any("deviceid:", deviceId))
	value, ok := ConnsManager.Load(deviceId)
	if ok {
		return value.(*apisocket.Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(deviceId int64) {
	ConnsManager.Delete(deviceId)
}

// PushAll 全服推送
func PushAll(message *pb.Message) {
	ConnsManager.Range(func(key, value interface{}) bool {
		conn := value.(*apisocket.Conn)
		logger.Logger.Info("pushall", zap.Any("userid:", conn.UserId), zap.Any("device", conn.DeviceId))
		conn.Send(pb.PackageType_PT_MESSAGE, 0, message, nil)
		return true
	})
}
