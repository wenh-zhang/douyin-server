package db

import (
	"context"
	"douyin/dal"
	"douyin/pkg/constants"
	"time"
)

type Favorite struct {
	ID        int64     `json:"id"`
	VideoID   int64     `json:"video_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (*Favorite) TableName() string {
	return constants.FavoriteTableName
}

// MGetFavorites multiple get list of favorites info
func MGetFavorites(ctx context.Context, videoID, userID *int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	params := make(map[string]interface{})
	if videoID != nil {
		params["video_id"] = *videoID
	}
	if userID != nil {
		params["user_id"] = *userID
	}
	if err := dal.DB.WithContext(ctx).Where(params).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateFavorite create favorite info
func CreateFavorite(ctx context.Context, favorites []*Favorite) error {
	return dal.DB.WithContext(ctx).Create(favorites).Error
}

// DeleteFavorite delete favorite info
func DeleteFavorite(ctx context.Context, videoID, userID int64) error {
	return dal.DB.WithContext(ctx).Where("video_id = ? and user_id = ?", videoID, userID).Delete(&Favorite{}).Error
}
