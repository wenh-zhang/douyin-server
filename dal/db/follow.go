package db

import (
	"context"
	"douyin/pkg/constant"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	ID         int64     `json:"id"`
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (*Follow) TableName() string {
	return constant.FollowTableName
}

// MGetFollows multiple get list of follows info
func MGetFollows(ctx context.Context, db *gorm.DB, userID int64) ([]*Follow, error) {
	res := make([]*Follow, 0)
	if err := db.WithContext(ctx).Where("from_user_id = ?", userID).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CheckIfFollow check whether the user had been followed by anther user
func CheckIfFollow(ctx context.Context, db *gorm.DB, fromUserID, toUserID int64) (bool, error) {
	var count int64
	if err := db.WithContext(ctx).Model(&Follow{}).Where("from_user_id = ? and to_user_id = ?", fromUserID, toUserID).Count(&count).Error; err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// CreateFollow create follow info
func CreateFollow(ctx context.Context, db *gorm.DB, follows []*Follow) error {
	return db.WithContext(ctx).Create(follows).Error
}

// DeleteFollow delete follow info
func DeleteFollow(ctx context.Context, db *gorm.DB, fromUserID, toUserID int64) error {
	return db.WithContext(ctx).Where("from_user_id = ? and to_user_id = ?", fromUserID, toUserID).Delete(&Follow{}).Error
}

// DeleteFollowOfUser delete follow relationship when a user is deleted
func DeleteFollowOfUser(ctx context.Context, db *gorm.DB, fromUserID int64) error {
	return db.WithContext(ctx).Where("from_user_id = ?", fromUserID).Delete(&Follow{}).Error
}
