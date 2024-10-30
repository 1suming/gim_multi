package apisocket

import (
	"context"
	app2 "gim/internal/logic/domain/user/app"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

//	func (c *Conn) Handle_SearchUser(input *pb.Input) {
//		logger.Logger.Info("Handle_SearchUser", zap.Any("input", input))
//		var req pb.SearchUserReq
//		err := proto.Unmarshal(input.Data, &req)
//		if err != nil {
//			logger.Logger.Error("handle_SearchUser", zap.Error(err))
//			return
//		}
//		deviceId, userId, token := c.DeviceId, c.UserId, c.LoginToken
//		resp, err := rpc.GetBusinessExtClient().SearchUser(grpclib.ContextWithUserInfo(context.TODO(), input.RequestId, deviceId, userId, token), &req)
//		if err != nil {
//			logger.Logger.Error("handle_SearchUser", zap.Error(err))
//		}
//		//return resp, err
//
//		logger.Logger.Info(" handle_SearchUser", zap.Any("resp", resp))
//
//		c.Send(pb.PackageType_PT_SEARCH_USER, input.RequestId, resp, err)
//
// }
func (c *Conn) Handle_SearchUser(input *pb.Input) {
	logger.Logger.Info("Handle_SearchUser", zap.Any("input", input))
	var req pb.SearchUserReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("handle_SearchUser", zap.Error(err))
		return
	}
	users, err := app2.UserApp.Search(context.TODO(), req.Key)
	resp, err := &pb.SearchUserResp{Users: users}, err

	logger.Logger.Info(" handle_SearchUser", zap.Any("resp", resp))

	c.Send(pb.PackageType_PT_SEARCH_USER, input.RequestId, resp, err)

}
