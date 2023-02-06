package db

import (
	"context"
	"douyin/dal"
	"douyin/pkg/constants"
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
	return constants.VideoTableName
}

// MGetVideos multiple get list of video info
func MGetVideos(ctx context.Context, videoIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(videoIDs) == 0 {
		return res, nil
	}

	if err := dal.DB.WithContext(ctx).Where("id in ?", videoIDs).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetVideosByUserID multiple get list of video info
func MGetVideosByUserID(ctx context.Context, userIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := dal.DB.WithContext(ctx).Where("user_id in ?", userIDs).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetVideosByTime multiple get list of video info
func MGetVideosByTime(ctx context.Context, lastTime time.Time, limit int) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := dal.DB.WithContext(ctx).Where("created_at < ?", lastTime).Order("created_at desc").Limit(limit).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, videos []*Video) error {
	return dal.DB.WithContext(ctx).Create(videos).Error
}

// UpdateVideo update video info
func UpdateVideo(ctx context.Context, videoID int64, playURL, coverURL, title *string, favoriteCount, commentCount int64) error {
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
	return dal.DB.WithContext(ctx).Model(&Video{}).Where("id = ?", videoID).Updates(params).Error
}

// DeleteVideo delete video info
func DeleteVideo(ctx context.Context, videoID int64) error {
	return dal.DB.WithContext(ctx).Where("id = ?", videoID).Delete(&Video{}).Error
}
