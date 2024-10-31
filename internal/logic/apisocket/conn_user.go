package apisocket

import (
	"context"
	app2 "gim/internal/logic/domain/user/app"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
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
func (c *Conn) Handle_SearchUser(input *pb.Input) error {
	logger.Logger.Info("Handle_SearchUser", zap.Any("input", input))
	var req pb.SearchUserReq
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("handle_SearchUser", zap.Error(err))
		c.Send(pb.PackageType_PT_SEARCH_USER, input.RequestId, nil, err)
		return err
	}
	users, err := app2.UserApp.Search(context.TODO(), req.Key)
	resp, err := &pb.SearchUserResp{Users: users}, err

	logger.Logger.Info(" handle_SearchUser", zap.Any("resp", resp))

	c.Send(pb.PackageType_PT_SEARCH_USER, input.RequestId, resp, err)
	return nil
}

func (c *Conn) Handle_GetUser(input *pb.Input) error {

	var req pb.GetUserReq
	logger.Logger.Info("GetUser", zap.Any("input", input))
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("handle_SearchUser", zap.Error(err))
		c.Send(pb.PackageType_PT_GET_USER, input.RequestId, nil, err)
		return err
	}

	//userId, _, err := grpclib.GetCtxData(ctx)
	////if err != nil {
	////	return nil, err
	////}
	user, err := app2.UserApp.Get(context.TODO(), req.UserId)
	resp, err := &pb.GetUserResp{User: user}, err

	c.Send(pb.PackageType_PT_GET_USER, input.RequestId, resp, err)
	return nil
}
func (c *Conn) Handle_GetUsers(input *pb.Input) error {

	var req pb.GetUserIdsReq
	logger.Logger.Info("GetUser", zap.Any("input", input))
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("handle_SearchUser", zap.Error(err))
		c.Send(pb.PackageType_PT_GET_USERS, input.RequestId, nil, err)
		return err
	}

	users, err := app2.UserApp.GetByIds(context.TODO(), req.UserIds)
	resp, err := &pb.GetUsersResp{Users: users}, err
	c.Send(pb.PackageType_PT_GET_USERS, input.RequestId, resp, err)
	return nil
}

func (c *Conn) Handle_UpdateUser(input *pb.Input) error {
	//userId, _, err := grpclib.GetCtxData(ctx)
	//if err != nil {
	//	return nil, err
	//}
	userId := c.UserId
	var req pb.UpdateUserReq
	logger.Logger.Info("GetUser", zap.Any("input", input))
	err := proto.Unmarshal(input.Data, &req)
	if err != nil {
		logger.Logger.Error("handle_SearchUser", zap.Error(err))
		c.Send(pb.PackageType_PT_UPDATE_USER, input.RequestId, nil, err)
		return err
	}
	resp := new(emptypb.Empty)
	err = app2.UserApp.Update(context.TODO(), userId, &req)
	// new(emptypb.Empty), app2.UserApp.Update(ctx, userId, req)
	c.Send(pb.PackageType_PT_UPDATE_USER, input.RequestId, resp, err)
	return nil
}
