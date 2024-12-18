package service

import (
	"context"
	"errors"
	"gim/internal/logic/domain/recentcontact/model"
	"gim/internal/logic/domain/recentcontact/repo"
	"gim/pkg/commondefine"
	"gim/pkg/db"
	"gim/pkg/dto"
	"gim/pkg/gerrors"
	"gim/pkg/logger"
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"strconv"
)

type SRecentConversationService struct{}

var RecentConversationService = new(SRecentConversationService)

func (r *SRecentConversationService) SaveOrUpdate(ctx context.Context, dataDto *dto.SaveOrUpdateRecentContactDTO) error {
	logger.Logger.Info("SRecentConversationService func start")
	err := r._saveOrUpdateSingle(ctx, dataDto, dataDto.OwnerUid, dataDto.TargetId, dataDto.LastMessageId)
	if err != nil {
		return err
	}
	err = r._saveOrUpdateSingle(ctx, dataDto, dataDto.TargetId, dataDto.OwnerUid, dataDto.LastTargetMessageId)
	if err != nil {
		return err
	}
	sourceId := dataDto.TargetId //收件人
	targetId := dataDto.OwnerUid

	logger.Logger.Info("update REDIS_KEY_CONVERSAION_UNREAD_CNT", zap.Any("userId:", sourceId), zap.Any("targetId", targetId))

	err = db.RedisUtil.GetRedisClient().HIncrBy(commondefine.REDIS_KEY_CONVERSAION_UNREAD_CNT+"_"+strconv.FormatInt(sourceId, 10), strconv.FormatInt(targetId, 10), 1).Err()
	if err != nil {
		logger.Logger.Error("redis error", zap.Error(err))
		return err
	}
	totalUnreadKey := commondefine.REDIS_KEY_CONVERSAION_UNREAD_TOTAL_CNT + "_" + strconv.FormatInt(sourceId, 10)
	err = db.RedisUtil.GetRedisClient().Incr(totalUnreadKey).Err()
	if err != nil {
		logger.Logger.Error("redis error", zap.Error(err))
		return err
	}

	return err

}

func (r *SRecentConversationService) _saveOrUpdateSingle(ctx context.Context, dataDto *dto.SaveOrUpdateRecentContactDTO, sourceId int64, targetId int64, lastMessageId int64) error {
	var recentConversationModel *model.ImRecentConversation

	recentConversationModel, err := repo.RecentContactRepo.Get(dataDto.ConversationType, sourceId, targetId)
	if err != nil {
		logger.Logger.Error("db error", zap.Error(err))
		return err
	}
	if recentConversationModel == nil {
		recentConversationModel = &model.ImRecentConversation{
			ConversationType:   dataDto.ConversationType,
			OwnerUid:           sourceId,
			LastMessageContent: dataDto.LastMessageContent,
			LastMessageId:      lastMessageId,
			TargetId:           targetId,
			LastTime:           dataDto.LastTime,
		}
		err := repo.RecentContactRepo.Save(recentConversationModel)
		if err != nil {
			logger.Logger.Error("db  error", zap.Error(err))
			return gerrors.WrapError(err)
		}
	} else {
		updateMap := make(map[string]interface{})

		updateMap["last_message_content"] = dataDto.LastMessageContent
		updateMap["last_message_id"] = dataDto.LastMessageId
		updateMap["last_time"] = dataDto.LastTime

		_, err := repo.RecentContactRepo.Update(dataDto.ConversationType, sourceId, targetId, updateMap)
		if err != nil {
			logger.Logger.Error("db  error", zap.Error(err))
			return gerrors.WrapError(err)
		}
	}
	//
	///**更未读更新 */
	//redisTemplate.opsForValue().increment(recipientUid + "_T", 1); //加总未读
	//redisTemplate.opsForHash().increment(recipientUid + "_C", senderUid, 1); //加会话未读
	//hmap

	return nil

}

func (r *SRecentConversationService) GetUserRecentConversations(ctx context.Context, userId int64, userRecentConverationAll *dto.UserRecentConversationAll) error {
	//var userRecentConverationAll dto.UserRecentConversationAll
	logger.Logger.Info("GetUserRecentConversations func start")
	var totalUnread int64
	totalUnread = 0
	err := db.RedisUtil.Get(commondefine.REDIS_KEY_CONVERSAION_UNREAD_TOTAL_CNT+"_"+strconv.FormatInt(userId, 10), &totalUnread)
	if err != nil && !errors.Is(err, redis.Nil) {
		return gerrors.WrapError(err)
	}
	userRecentConverationAll.TotalUnread = totalUnread

	userRecentConverationAll.Conversations = make([]*dto.UserRecentConversationSingle, 0)
	records, err := repo.RecentContactRepo.QueryAll(userId)
	if err != nil {
		return err
	}
	for _, record := range records {
		var conversationSingle dto.UserRecentConversationSingle
		copier.Copy(&conversationSingle, &record)

		targetId := record.TargetId
		val, err := db.RedisUtil.GetRedisClient().HGet(commondefine.REDIS_KEY_CONVERSAION_UNREAD_CNT+"_"+strconv.FormatInt(userId, 10), strconv.FormatInt(targetId, 10)).Result()
		if err != nil {
			// 如果返回的错误是key不存在
			if errors.Is(err, redis.Nil) {
				logger.Logger.Info("redis is nil", zap.Error(err))
				conversationSingle.UnreadCnt = 0
			} else {
				return gerrors.WrapError(err)
			}
		} else {
			unread, err := strconv.Atoi(val)
			if err != nil {
				logger.Logger.Error("redis error", zap.Error(err))
				return err
			}
			logger.Logger.Info("REDIS_KEY_CONVERSAION_UNREAD_CNT", zap.Any("userId:", userId), zap.Any("targetId", targetId), zap.Any("unread", unread))
			conversationSingle.UnreadCnt = int64(unread)
		}

		userRecentConverationAll.Conversations = append(userRecentConverationAll.Conversations, &conversationSingle)

		//userRecentConverationAll.Conversations = append(userRecentConverationAll.Conversations, &dto.UserRecentConversationSingle{
		//	ConversationType: record.ConversationType,
		//	LastMessageContent: record.LastMessageContent,
		//	LastMessageId: record.LastMessageId,
		//	LastTime: record.LastTime,
		//	OwnerUid: record.OwnerUid,
		//	TargetId: record.TargetId,
		//})
	}
	return nil
}
