package service

import (
	"context"
	"gim/internal/business/domain/user/model"
	"gim/internal/business/domain/user/repo"
	"gim/pkg/gerrors"
	"gim/pkg/protocol/pb"
	"gim/pkg/rpc"
	"gim/pkg/util"
	"time"
)

type authService struct{}

var AuthService = new(authService)

const (
	LOGIN_OPERATE_TYPE_REGISTER int32 = 1
	LOGIN_OPERATE_TYPE_LOGIN    int32 = 2
)

// SignIn 登录
func (*authService) SignIn(ctx context.Context, phoneNumber, code string, deviceId int64, operate_type int32, pwd string) (bool, int64, string, error) {
	user, err := repo.UserRepo.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return false, 0, "", err
	}

	var isNew = false

	if operate_type == LOGIN_OPERATE_TYPE_REGISTER {
		if user != nil {
			return false, 0, "", gerrors.ErrUserExisted
		}
		if !Verify(phoneNumber, code) {
			return false, 0, "", gerrors.ErrBadCode
		}

		if user == nil {
			user = &model.User{
				PhoneNumber: phoneNumber,
				CreateTime:  time.Now(),
				UpdateTime:  time.Now(),
				Password:    util.Md5(pwd),
			}
			err := repo.UserRepo.Save(user)
			if err != nil {
				return false, 0, "", err
			}
			isNew = true
		}

	} else if operate_type == LOGIN_OPERATE_TYPE_LOGIN {
		if user == nil {
			return false, 0, "", gerrors.ErrUserNotFound
		}
		if user.Password != util.Md5(pwd) {
			return false, 0, "", gerrors.ErrPasswordError
		}
	}

	resp, err := rpc.GetLogicIntClient().GetDevice(ctx, &pb.GetDeviceReq{DeviceId: deviceId})
	if err != nil {
		return false, 0, "", err
	}

	// 方便测试
	token := "1"
	//token := util.RandString(40)
	err = repo.AuthRepo.Set(user.Id, resp.Device.DeviceId, model.Device{
		Type:   resp.Device.Type,
		Token:  token,
		Expire: time.Now().AddDate(0, 3, 0).Unix(),
	})
	if err != nil {
		return false, 0, "", err
	}

	return isNew, user.Id, token, nil
}

func Verify(phoneNumber, code string) bool {
	// 假装他成功了
	return true
}

// Auth 验证用户是否登录
func (*authService) Auth(ctx context.Context, userId, deviceId int64, token string) error {
	device, err := repo.AuthRepo.Get(userId, deviceId)
	if err != nil {
		return err
	}

	if device == nil {
		return gerrors.ErrUnauthorized
	}

	if device.Expire < time.Now().Unix() {
		return gerrors.ErrUnauthorized
	}

	if device.Token != token {
		return gerrors.ErrUnauthorized
	}
	return nil
}
