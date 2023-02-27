package dao

import (
	"context"
	"douyin/cmd/rpc/message/model"
	"gorm.io/gorm"
)

type Message struct {
	db *gorm.DB
}

func NewMessage(db *gorm.DB) *Message {
	return &Message{
		db: db,
	}
}

func (s *Message) CreateMessage(ctx context.Context, message *model.Message) error {
	return s.db.WithContext(ctx).Create(message).Error
}

func (s *Message) GetMessageListByUserId(ctx context.Context, fromUserId, toUserId, preMsgTime int64) ([]*model.Message, error) {
	messageList := make([]*model.Message, 0)
	if err := s.db.WithContext(ctx).Where("created_at > ?", preMsgTime).Where("(from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?)",
		fromUserId, toUserId, toUserId, fromUserId).Order("created_at").Find(&messageList).Error; err != nil {
		return nil, err
	}
	return messageList, nil
}

func (s *Message) GetLatestMsgByUserId(ctx context.Context, fromUserId, toUserId int64) (*model.Message, error) {
	message := new(model.Message)
	if err := s.db.WithContext(ctx).Where("(from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?)",
		fromUserId, toUserId, toUserId, fromUserId).Order("created_at desc").First(message).Error; err != nil {
		return message, err
	}
	return message, nil
}

func (s *Message) BatchGetLatestMsgByUserId(ctx context.Context, fromUserId int64, toUserIds []int64) ([]*model.Message, error) {
	messageList := make([]*model.Message, 0)
	for _, toUserId := range toUserIds {
		message, err := s.GetLatestMsgByUserId(ctx, fromUserId, toUserId)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		messageList = append(messageList, message)
	}
	return messageList, nil
}
