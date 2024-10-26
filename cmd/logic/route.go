package main

import (
	"fmt"
	"gim/internal/logic/api"
	"gim/internal/logic/domain/device"
	"gim/pkg/errs"
	"gim/pkg/gerrors"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"
	"gim/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	RETCODE_FAIL int32 = -1
)

type S_RegisterDeviceReq struct {
	Type          int32  ` json:"type,omitempty"`           // 设备类型
	Brand         string ` json:"brand,omitempty"`          // 厂商
	Model         string ` json:"model,omitempty"`          // 机型
	SystemVersion string ` json:"system_version,omitempty"` // 系统版本
	SdkVersion    string ` json:"sdk_version,omitempty"`    // sdk版本号
}
type S_RegisterDeviceResp struct {
	DeviceId int64 `json:"device_id"` // 设备id

}

var (
	logicExt *api.LogicExtServer = &api.LogicExtServer{}
)

func RegisterDevice(ctx *gin.Context) {

	var req S_RegisterDeviceReq
	var httpResp S_RegisterDeviceResp
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, response.Errno(errs.ErrParam))
		return
	}

	resp, err := logicExt.RegisterDevice(ctx,
		&pb.RegisterDeviceReq{
			Type:          1,
			Brand:         "huawei",
			Model:         "huawei P30",
			SystemVersion: "1.0.0",
			SdkVersion:    "1.0.0",
		})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(200, response.Errno(errs.ErrParam))
		return
	}

	httpResp.DeviceId = resp.DeviceId

	ctx.JSON(200, response.Ok(resp))
}

type S_SignInReq struct {
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"` // 手机号
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`                                  // 验证码
	DeviceId    int64  `protobuf:"varint,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`         // 设备id

	OperateType int32  `protobuf:"varint,4,opt,name=operate_type,json=operateType,proto3" json:"operate_type,omitempty"` //操作类型
	Pwd         string `protobuf:"bytes,5,opt,name=pwd,proto3" json:"pwd,omitempty"`                                     //密码@ms
}

type S_SignInResp struct {
	IsNew  bool   ` json:"is_new,omitempty"` // 是否是新用户
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // token

}

// 登录获取token
func GetToken(ctx *gin.Context) {

	var req S_SignInReq
	var httpResp S_SignInResp
	if err := ctx.ShouldBind(&req); err != nil {
		logger.Logger.Info("GetToken err", zap.Error(err))
		ctx.JSON(400, response.Errno(errs.ErrParam))
		return
	}

	resp, err := device.Service.GetToken(ctx, req.PhoneNumber, req.Code, req.DeviceId, req.OperateType, req.Pwd)
	if err != nil {
		logger.Logger.Info("GetToken err", zap.Error(err))
		if err == gerrors.ErrUserExisted {
			ctx.JSON(200, response.Errno(errs.ErrUserExisted))
			return
		}
		if err == gerrors.ErrUserNotFound || err == gerrors.ErrPasswordError {
			ctx.JSON(200, response.Errno(errs.ErrAccountOrPasswordIncorrect))
			return
		}
		ctx.JSON(200, response.Errno(errs.ErrParam))
		return
	}
	httpResp.IsNew = resp.IsNew
	httpResp.Token = resp.Token
	httpResp.UserId = resp.UserId

	ctx.JSON(200, response.Ok(resp))
}
