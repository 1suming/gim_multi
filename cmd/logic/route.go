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

	DeviceUniqueId string ` json:"device_unique_id"` //设备唯一id;@ms

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
	//&pb.RegisterDeviceReq{
	//	Type:           1,
	//	Brand:          "huawei",
	//	Model:          "huawei P30",
	//	SystemVersion:  "1.0.0",
	//	SdkVersion:     "1.0.0",
	//	DeviceUniqueId: "xx",
	//})
	resp, err := logicExt.RegisterDevice(ctx,
		&pb.RegisterDeviceReq{
			Type:           req.Type,
			Brand:          req.Brand,
			Model:          req.Model,
			SystemVersion:  req.SystemVersion,
			SdkVersion:     req.SdkVersion,
			DeviceUniqueId: req.DeviceUniqueId,
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
	PhoneNumber string `json:"phone_number,omitempty"` // 手机号
	Code        string `json:"code,omitempty"`         // 验证码
	DeviceId    int64  `json:"device_id,omitempty"`    // 设备id

	OperateType int32  `json:"operate_type,omitempty"` //操作类型
	Pwd         string `json:"pwd,omitempty"`          //密码@ms
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
		//我们在调用errors.New("")来返回一个错误时， 可以通过比较指针，来比较error是否相等， 实际上就是控制相同的错误我们只创建一个error对象。否则对象复制一下，在比较就是false了。
		if gerrors.ErrUserExisted.Error() == err.Error() {
			ctx.JSON(200, response.Errno(errs.ErrUserExisted))
			return
		}
		if err.Error() == gerrors.ErrUserNotFound.Error() || err.Error() == gerrors.ErrPasswordError.Error() {
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
