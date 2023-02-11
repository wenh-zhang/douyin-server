package db

import (
	"context"
	"douyin/pkg/constant"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID         int64     `json:"id"`
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

func (*Message) TableName() string {
	return constant.MessageTableName
}

// MGetMessages multiple get list of messages info
func MGetMessages(ctx context.Context, db *gorm.DB, fromUserID, toUserID int64) ([]*Message, error) {
	res := make([]*Message, 0)
	if err := db.WithContext(ctx).Where("(from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?)", fromUserID, toUserID, toUserID, fromUserID).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetAllMessages multiple get list of messages info
func MGetAllMessages(ctx context.Context, db *gorm.DB, userID int64) ([]*Message, error) {
	res := make([]*Message, 0)
	if err := db.WithContext(ctx).Where("from_user_id = ?  or to_user_id = ?", userID, userID).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateMessage create message info
func CreateMessage(ctx context.Context, db *gorm.DB, messages []*Message) error {
	return db.WithContext(ctx).Create(messages).Error
}

// DeleteMessage delete follow info
func DeleteMessage(ctx context.Context, db *gorm.DB, fromUserID, toUserID int64) error {
	return db.WithContext(ctx).Where("from_user_id = ? and to_user_id = ?", fromUserID, toUserID).Delete(&Message{}).Error
}
