package db

import (
	"context"
	"douyin/dal"
	"douyin/pkg/constants"
	"time"
)

type Follow struct {
	ID         int64     `json:"id"`
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
	CreatedAt  time.Time `json:"created_at"`
}

func (*Follow) TableName() string {
	return constants.FollowTableName
}

// MGetFollows multiple get list of follows info
func MGetFollows(ctx context.Context, userID int64) ([]*Follow, error) {
	res := make([]*Follow, 0)
	if err := dal.DB.WithContext(ctx).Where("from_user_id = ?", userID).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateFollow create follow info
func CreateFollow(ctx context.Context, follows []*Follow) error {
	return dal.DB.WithContext(ctx).Create(follows).Error
}

// DeleteFollow delete follow info
func DeleteFollow(ctx context.Context, fromUserID, toUserID int64) error {
	return dal.DB.WithContext(ctx).Where("from_user_id = ? and to_user_id = ?", fromUserID, toUserID).Delete(&Follow{}).Error
}
