package repo

import (
	"errors"
	"fmt"
	"gim/internal/logic/domain/recentcontact/model"
	"gim/pkg/db"
	"gim/pkg/gerrors"
	"github.com/jinzhu/gorm"
)

type SRecentConversationRepo struct{}

var RecentContactRepo = new(SRecentConversationRepo)

func (*SRecentConversationRepo) tableName() string {
	return fmt.Sprintf("im_recent_conversation")
}

// Save 插入一
func (d *SRecentConversationRepo) Save(row *model.ImRecentConversation) error {
	err := db.DB.Table(d.tableName()).Create(row).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// Get 获取好友
func (*SRecentConversationRepo) Get(conversation_type int8, userId int64, targetId int64) (*model.ImRecentConversation, error) {
	row := model.ImRecentConversation{}
	err := db.DB.First(&row, "conversation_type = ? and owner_uid = ? and target_id=?", conversation_type, userId, targetId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &row, nil
}
func (*SRecentConversationRepo) Update(conversation_type int8, userId int64, targetId int64, updateMap map[string]interface{}) (int64, error) {
	db := db.DB.Model(&model.ImRecentConversation{}).Where("conversation_type = ? and owner_uid = ? and target_id=?", conversation_type, userId, targetId).
		Updates(updateMap)
	if db.Error != nil {
		return 0, gerrors.WrapError(db.Error)
	}
	return db.RowsAffected, nil
}

// ListBySeq 根据类型和id查询大于序号大于seq的消息
func (r *SRecentConversationRepo) QueryAll(userId int64) ([]model.ImRecentConversation, error) {
	DB := db.DB.Table(r.tableName()).
		Where(" owner_uid = ?  ", userId)

	//var count int64
	//err := DB.Count(&count).Error
	//if err != nil {
	//	return nil, gerrors.WrapError(err)
	//}
	//if count == 0 {
	//	return nil, nil
	//}

	var messages []model.ImRecentConversation
	err := DB.Find(&messages).Error
	if err != nil {
		return nil, gerrors.WrapError(err)
	}
	return messages, nil
}
