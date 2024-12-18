package apisocket

import (
	"container/list"
	"context"
	"gim/config"
	"gim/internal/logic/domain/device"
	"gim/internal/logic/domain/message"
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
	logger.Logger.Info("Socket close:", zap.Any("userid", c.UserId), zap.Any("deviceID", c.DeviceId))
	// 取消设备和连接的对应关系
	if c.DeviceId != 0 {
		DeleteConn(c.DeviceId)
	}

	//// 取消订阅，需要异步出去，防止重复加锁造成死锁
	go func() {
		isSendMsg := true
		UnSubscribedRoom(c, isSendMsg)
	}()

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
	//	MessageACK(input)
	case pb.PackageType_PT_MESSAGE_ACK: //消息回执
		Handle_MessageACK(c, input)
	//case pb.PackageType_PT_SUBSCRIBE_ROOM:
	//	Handle_SubscribedRoom(c, input)
	//

	case pb.PackageType_PT_SEARCH_USER:

		Handle_SearchUser(c, input)
	case pb.PackageType_PT_GET_USER:
		Handle_GetUser(c, input)
	case pb.PackageType_PT_GET_USERS:
		Handle_GetUsers(c, input)
	case pb.PackageType_PT_UPDATE_USER:
		Handle_UpdateUser(c, input)

	case pb.PackageType_PT_SEND_MESSAGE:
		Handle_SendMessage(c, input)

		//会话列表
	case pb.PackageType_PT_GET_USER_CONVERSATIONS:
		Handle_GetUserConversations(c, input)

	case pb.PackageType_PT_GET_USER_MESSAGES: //得到某个会话历史消息
		Handle_GetUserMessages(c, input)
	//chatroom
	case pb.PackageType_PT_ROOM_GET_ROOM_LIST:
		Handle_GetRoomList(c, input)
		//PT_ROOM_GET_ROOM_LIST = 40001;//room列表
		//PT_ROOM_CREATE_ROOM =40002;//创建room
		//PT_ROOM_DELETE_ROOM =40003;//解散room
		//PT_ROOM_JOIN_ROOM = 40004;//加入room

	case pb.PackageType_PT_ROOM_JOIN_ROOM: // pb.PackageType_PT_SUBSCRIBE_ROOM:
		Handle_SubscribedRoom(c, input)
	case pb.PackageType_PT_ROOM_QUITROOM: // pb.PackageType_PT_SUBSCRIBE_ROOM:
		Handle_QuitRoom(c, input)

	//friend
	case pb.PackageType_PT_FRIEND_ADD_FRIEND:
		Handle_AddFriend(c, input)
	case pb.PackageType_PT_FRIEND_AGREE_ADD_FRIEND:
		Handle_AgreeAddFriend(c, input)
	case pb.PackageType_PT_FRIEND_GET_FRIENDS:
		Handle_GetFriends(c, input)
	case pb.PackageType_PT_FRIEND_GET_FRIEND_REQUESTS:
		Handle_GetFriendReqs(c, input)
	default:
		logger.Logger.Error("handler switch other")
	}
}
func Handle_SendMessage(c *Conn, input *pb.Input) error {

	var req pb.SendMessageReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("Handle_SendMessageToFriend", zap.Error(err))
		return err
	}
	//会话类型：单聊、群聊和聊天室分别为 `singleChat`、`groupChat` 和 `chatRoom`，默认为单聊。
	if req.ChatType == pb.ChatType_SINGLE_CHAT {
		Handle_SendMessageToFriend(c, input)
	} else if req.ChatType == pb.ChatType_GROUP_CHAT {
		//
	} else if req.ChatType == pb.ChatType_CHAT_ROOM {
		Handle_SendMsgToRoom(c, input)
	}

	return nil
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

// MessageACK 消息收到回执
func Handle_MessageACK(c *Conn, input *pb.Input) {
	var messageACK pb.MessageACK
	err := proto.Unmarshal(input.Data, &messageACK)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}
	//只有单聊才能ack，chatroom不能，

	//_, _ = rpc.GetLogicIntClient().MessageACK(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.MessageACKReq{
	//	UserId:      c.UserId,
	//	DeviceId:    c.DeviceId,
	//	DeviceAck:   messageACK.DeviceAck,
	//	ReceiveTime: messageACK.ReceiveTime,
	//})
	userId, deviceId := c.UserId, c.DeviceId

	err = message.App.MessageAck(context.TODO(), userId, deviceId, messageACK.DeviceAck)
	if err != nil {
		logger.Logger.Info("handle", zap.Any("err", err))
	}

	c.Send(pb.PackageType_PT_MESSAGE_ACK, input.RequestId, nil, err)
	return

}
