package apisocket

import (
	"container/list"
	"context"
	"gim/config"
	"gim/internal/logic/domain/device"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"sync"
	"time"

	"github.com/alberliu/gn"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const (
	CoonTypeTCP int8 = 1 // tcp连接
	ConnTypeWS  int8 = 2 // websocket连接
)

//var logicIntServer *api.LogicIntServer = &api.LogicIntServer{}

//var logicExtServer *api.LogicExtServer = &api.LogicExtServer{}

type Conn struct {
	CoonType int8            // 连接类型
	TCP      *gn.Conn        // tcp连接
	WSMutex  sync.Mutex      // WS写锁
	WS       *websocket.Conn // websocket连接
	UserId   int64           // 用户ID
	DeviceId int64           // 设备ID
	RoomId   int64           // 订阅的房间ID
	Element  *list.Element   // 链表节点

	LoginToken string //@ms:tmp 不应该保存，暂时为了方便
}

// Write 写入数据
func (c *Conn) Write(bytes []byte) error {
	if c.CoonType == CoonTypeTCP {
		return c.TCP.WriteWithEncoder(bytes)
	} else if c.CoonType == ConnTypeWS {
		return c.WriteToWS(bytes)
	}
	logger.Logger.Error("unknown conn type", zap.Any("conn", c))
	return nil
}

// WriteToWS 消息写入WebSocket
func (c *Conn) WriteToWS(bytes []byte) error {
	c.WSMutex.Lock()
	defer c.WSMutex.Unlock()

	err := c.WS.SetWriteDeadline(time.Now().Add(10 * time.Millisecond))
	if err != nil {
		return err
	}
	return c.WS.WriteMessage(websocket.BinaryMessage, bytes)
}

// Close 关闭
func (c *Conn) Close() error {
	// 取消设备和连接的对应关系
	if c.DeviceId != 0 {
		DeleteConn(c.DeviceId)
	}

	//// 取消订阅，需要异步出去，防止重复加锁造成死锁
	//go func() {
	//	SubscribedRoom(c, 0)
	//}()

	if c.DeviceId != 0 {
		//_, _ = rpc.GetLogicIntClient().Offline(context.TODO(), &pb.OfflineReq{
		//_, _ = logicIntServer.Offline(context.TODO(), &pb.OfflineReq{
		//UserId:     c.UserId,
		//	DeviceId:   c.DeviceId,
		//		ClientAddr: c.GetAddr(),

		_ = device.App.Offline(context.TODO(), c.DeviceId, c.GetAddr())

	}

	if c.CoonType == CoonTypeTCP {
		c.TCP.Close()
	} else if c.CoonType == ConnTypeWS {
		return c.WS.Close()
	}
	return nil
}

func (c *Conn) GetAddr() string {
	if c.CoonType == CoonTypeTCP {
		return c.TCP.GetAddr()
	} else if c.CoonType == ConnTypeWS {
		return c.WS.RemoteAddr().String()
	}
	return ""
}

// HandleMessage 消息处理
func (c *Conn) HandleMessage(bytes []byte) {
	var input = new(pb.Input)
	err := proto.Unmarshal(bytes, input)
	if err != nil {
		logger.Logger.Error("unmarshal error", zap.Error(err), zap.Int("len", len(bytes)))
		return
	}
	logger.Logger.Debug("HandleMessage", zap.Any("input", input))

	// 对未登录的用户进行拦截
	if input.Type != pb.PackageType_PT_SIGN_IN && c.UserId == 0 {
		logger.Logger.Error("没有登录", zap.Any("info", "no login"))

		return
	}

	switch input.Type {
	case pb.PackageType_PT_SIGN_IN:
		c.SignIn(input)
	//case pb.PackageType_PT_SYNC:
	//	c.Sync(input)
	case pb.PackageType_PT_HEARTBEAT:
		c.Heartbeat(input)
	//case pb.PackageType_PT_MESSAGE:
	//	c.MessageACK(input)
	//case pb.PackageType_PT_SUBSCRIBE_ROOM:
	//	c.SubscribedRoom(input)
	case pb.PackageType_PT_SEARCH_USER:

		c.Handle_SearchUser(input)
	case pb.PackageType_PT_GET_USER:
		c.Handle_GetUser(input)
	case pb.PackageType_PT_GET_USERS:
		c.Handle_GetUsers(input)
	case pb.PackageType_PT_UPDATE_USER:
		c.Handle_UpdateUser(input)

	case pb.PackageType_PT_FRIEND_ADD_FRIEND:

		c.Handle_AddFriend(input)

	case pb.PackageType_PT_FRIEND_SEND_MSG_TO_FRIEND:
		c.Handle_SendMessageToFriend(input)

		//会话列表
	case pb.PackageType_PT_GET_USER_CONVERSATIONS:
		c.Handle_GetUserConversations(input)

	case pb.PackageType_PT_GET_USER_MESSAGES: //得到某个会话历史消息
		c.Handle_GetUserMessages(input)
	default:
		logger.Logger.Error("handler switch other")
	}
}

// Send 下发消息
func (c *Conn) Send(pt pb.PackageType, requestId int64, message proto.Message, err error) {
	var output = pb.Output{
		Type:      pt,
		RequestId: requestId,
	}

	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			output.Code = int32(status.Code())
			output.Message = status.Message()
		} else {
			//@ms:有可能不是gwrap的grcp封装的错误
			output.Code = int32(-1)
			output.Message = "未知错误"
		}

	}

	if message != nil {
		msgBytes, err := proto.Marshal(message)
		if err != nil {
			logger.Sugar.Error(err)
			return
		}
		output.Data = msgBytes
	}

	logger.Logger.Debug("HandleMessage-Send", zap.Any("send", output))

	outputBytes, err := proto.Marshal(&output)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Logger.Debug("send byte len:", zap.Int("len", len(outputBytes)))
	err = c.Write(outputBytes)
	if err != nil {
		logger.Sugar.Error(err)
		c.Close()
		return
	}
}

// SignIn 登录
func (c *Conn) SignIn(input *pb.Input) {
	var signIn pb.SignInInput
	err := proto.Unmarshal(input.Data, &signIn)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	logger.Logger.Debug(" SignIn", zap.Any("signIn", signIn))

	//_, err = rpc.GetLogicIntClient().ConnSignIn(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.ConnSignInReq{\
	//_, err = logicIntServer.ConnSignIn(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.ConnSignInReq{
	//	UserId:     signIn.UserId,
	//	DeviceId:   signIn.DeviceId,
	//	Token:      signIn.Token,
	//	ConnAddr:   config.Config.ConnectLocalAddr,
	//	ClientAddr: c.GetAddr(),
	//})
	req := pb.ConnSignInReq{
		UserId:     signIn.UserId,
		DeviceId:   signIn.DeviceId,
		Token:      signIn.Token,
		ConnAddr:   config.Config.ConnectLocalAddr,
		ClientAddr: c.GetAddr(),
	}

	err = device.App.SignIn(context.TODO(), req.UserId, req.DeviceId, req.Token, req.ConnAddr, req.ClientAddr)

	logger.Logger.Debug(" SignIn", zap.Any("signIn", "ok,send pkg"))

	c.Send(pb.PackageType_PT_SIGN_IN, input.RequestId, nil, err)
	if err != nil {
		return
	}

	c.UserId = signIn.UserId
	c.DeviceId = signIn.DeviceId
	//@ms:add
	c.LoginToken = signIn.Token //TODO

	SetConn(signIn.DeviceId, c)
}

// Heartbeat 心跳
func (c *Conn) Heartbeat(input *pb.Input) {
	c.Send(pb.PackageType_PT_HEARTBEAT, input.RequestId, nil, nil)

	logger.Sugar.Infow("heartbeat", "device_id", c.DeviceId, "user_id", c.UserId)
}
