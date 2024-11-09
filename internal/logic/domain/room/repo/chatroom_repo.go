package repo

import (
	"errors"
	"gim/internal/logic/domain/room/model"
	"gim/pkg/db"
	"gim/pkg/gerrors"
)
import (
	"github.com/jinzhu/gorm"
)

type SChatRoomRepo struct{}

var ChatRoomRepo = new(SChatRoomRepo)

// Get 获取群组信息
func (*SChatRoomRepo) Get(roomId int64) (*model.ChatRoom, error) {
	var room = model.ChatRoom{Id: roomId}
	err := db.DB.First(&room).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &room, nil
}

// Save 插入一条群组
func (*SChatRoomRepo) Save(room *model.ChatRoom) error {
	err := db.DB.Save(&room).Error
	if err != nil {
		return err
	}
	return nil
}

func (*SChatRoomRepo) QueryAll() ([]model.ChatRoom, error) {
	var rooms []model.ChatRoom
	err := db.DB.Find(&rooms).Error
	if err != nil {
		return nil, gerrors.WrapError(err)
	}
	return rooms, nil
}
