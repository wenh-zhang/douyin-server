package db

import (
	"context"
	"douyin/pkg/constant"
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	ID        int64     `json:"id"`
	VideoID   int64     `json:"video_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (*Favorite) TableName() string {
	return constant.FavoriteTableName
}

// MGetFavoriteVideoIDsByUserID multiple get favorite video ids by user id
func MGetFavoriteVideoIDsByUserID(ctx context.Context, db *gorm.DB, userID int64) ([]int64, error) {
	videoIDs := make([]int64, 0)
	if err := db.WithContext(ctx).Model(&Favorite{}).Where("user_id = ?", userID).Select("video_id").Order("created_at desc").Find(&videoIDs).Error; err != nil {
		return nil, err
	}
	return videoIDs, nil
}

// CheckIfFavorite check whether the video had been liked by user
func CheckIfFavorite(ctx context.Context, db *gorm.DB, videoID, userID int64) (bool, error) {
	var count int64
	if err := db.WithContext(ctx).Model(&Favorite{}).Where("video_id = ? and user_id = ?", videoID, userID).Count(&count).Error; err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// CreateFavorite create favorite info
func CreateFavorite(ctx context.Context, db *gorm.DB, favorites []*Favorite) error {
	return db.WithContext(ctx).Create(favorites).Error
}

// DeleteFavorite delete favorite info
func DeleteFavorite(ctx context.Context, db *gorm.DB, videoID, userID int64) error {
	return db.WithContext(ctx).Where("video_id = ? and user_id = ?", videoID, userID).Delete(&Favorite{}).Error
}

// DeleteFavoriteOfVideo delete favorite info of a video
func DeleteFavoriteOfVideo(ctx context.Context, db *gorm.DB, videoID int64) error {
	return db.WithContext(ctx).Where("video_id = ?", videoID).Delete(&Favorite{}).Error
}

// DeleteFavoriteOfUser delete favorite info of a user
func DeleteFavoriteOfUser(ctx context.Context, db *gorm.DB, userID int64) error {
	return db.WithContext(ctx).Where("user_id = ?", userID).Delete(&Favorite{}).Error
}
