package db

import (
	"context"
	"douyin/pkg/constant"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	gorm.Model
	UserID        int64  `json:"user_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title"`
}

func (*Video) TableName() string {
	return constant.VideoTableName
}

// MGetVideos multiple get list of video info
func MGetVideos(ctx context.Context, db *gorm.DB, videoIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(videoIDs) == 0 {
		return res, nil
	}

	if err := db.WithContext(ctx).Where("id in ?", videoIDs).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetVideosByUserID multiple get list of video info
func MGetVideosByUserID(ctx context.Context, db *gorm.DB, userIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := db.WithContext(ctx).Where("user_id in ?", userIDs).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetVideosByTime multiple get list of video info
func MGetVideosByTime(ctx context.Context, db *gorm.DB, latestTime int64, limit int) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := db.WithContext(ctx).Where("created_at < ?", time.Unix(latestTime, 0).Format(constant.TimestampFormatStr)).Order("created_at desc").Limit(limit).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, db *gorm.DB, videos []*Video) error {
	return db.WithContext(ctx).Create(videos).Error
}

// UpdateVideo update video info
func UpdateVideo(ctx context.Context, db *gorm.DB, videoID int64, playURL, coverURL, title *string, favoriteCount, commentCount int64) error {
	params := map[string]interface{}{}
	if playURL != nil {
		params["play_url"] = *playURL
	}
	if coverURL != nil {
		params["cover_url"] = *coverURL
	}
	if title != nil {
		params["title"] = *title
	}
	params["favorite_count"] = favoriteCount
	params["comment_count"] = commentCount
	return db.WithContext(ctx).Model(&Video{}).Where("id = ?", videoID).Updates(params).Error
}

// DeleteVideo delete video info
func DeleteVideo(ctx context.Context, db *gorm.DB, videoIDs []int64) error {
	return db.WithContext(ctx).Where("id in ?", videoIDs).Delete(&Video{}).Error
}
