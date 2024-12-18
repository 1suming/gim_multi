package repo

import (
	"fmt"
	"gim/internal/logic/domain/message/model"
	"gim/pkg/db"
	"gim/pkg/gerrors"
	"gim/pkg/logger"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type messageRepo struct{}

var MessageRepo = new(messageRepo)

func (*messageRepo) tableName(userId int64) string {
	return fmt.Sprintf("message")
}

// Save 插入一条消息
func (d *messageRepo) Save(message model.Message) error {
	err := db.DB.Table(d.tableName(message.UserId)).Create(&message).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// ListBySeq 根据类型和id查询大于序号大于seq的消息
func (d *messageRepo) ListBySeq(userId, seq, limit int64) ([]model.Message, bool, error) {
	DB := db.DB.Table(d.tableName(userId)).
		Where("user_id = ? and seq > ?", userId, seq).Order("seq ASC")

	var count int64
	err := DB.Count(&count).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	if count == 0 {
		return nil, false, nil
	}

	var messages []model.Message
	err = DB.Limit(limit).Find(&messages).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	return messages, count > limit, nil
}

// ListBySeq 根据类型和id查询大于序号大于seq的消息
func (d *messageRepo) ListBySeqAndTargetId(userId, seq, limit int64, targetId int64) ([]model.Message, bool, error) {
	//DB := db.DB.Table(d.tableName(userId)).
	//	Where("user_id = ? and seq > ? and target_id=?", userId, seq, targetId).Order("seq ASC")
	logger.Logger.Info("ListBySeqAndTargetId", zap.Any("userId:", userId), zap.Any("limit:", limit), zap.Any("targetId:", targetId))
	var DB *gorm.DB
	if seq == 0 {
		//这里我们改变原有的含有，原来为0表示从来没有拉取过消息，要从最开始拉取。
		//这里修改为拉取最新的，这样可以避免每次都拉取全部消息。
		DB = db.DB.Table(d.tableName(userId)).
			Where("user_id = ? and target_id=?", userId, targetId).Order("seq DESC") //逆序排序
	} else {
		DB = db.DB.Table(d.tableName(userId)).
			Where("user_id = ? and seq < ? and target_id=?", userId, seq, targetId).Order("seq DESC")
	}

	var totalCount int64
	err := DB.Count(&totalCount).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	if totalCount == 0 {
		return nil, false, nil
	}

	var messages []model.Message
	err = DB.Limit(limit).Find(&messages).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	return messages, totalCount > limit, nil
}
