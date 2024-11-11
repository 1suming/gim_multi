package friend

import (
	"errors"
	"gim/pkg/db"
	"gim/pkg/gerrors"

	"github.com/jinzhu/gorm"
)

type repo struct{}

var Repo = new(repo)

// Get 获取好友
func (*repo) Get(userId, friendId int64) (*Friend, error) {
	friend := Friend{}
	err := db.DB.First(&friend, "user_id = ? and friend_id = ?", userId, friendId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &friend, nil
}

// Save 添加好友
func (*repo) Save(friend *Friend) error {
	return gerrors.WrapError(db.DB.Save(&friend).Error)
}

// List 获取好友列表
func (*repo) List(userId int64, status int) ([]Friend, error) {
	var friends []Friend
	err := db.DB.Where("user_id = ? and status = ?", userId, status).Find(&friends).Error
	return friends, gerrors.WrapError(err)
}
func (*repo) GetFriendReqs(friend_id int64, status int, isSendFriend bool) ([]Friend, error) {
	var friends []Friend
	var err error
	if isSendFriend == true {
		err = db.DB.Where("user_id = ? and status = ?", friend_id, status).Find(&friends).Error
	} else {
		err = db.DB.Where("friend_id = ? and status = ?", friend_id, status).Find(&friends).Error
	}
	return friends, gerrors.WrapError(err)
}
