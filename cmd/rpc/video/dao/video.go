package dao

import (
	"context"
	"douyin/cmd/rpc/video/model"
	"gorm.io/gorm"
)

type Video struct {
	db *gorm.DB
}

func NewVideo(db *gorm.DB) *Video {
	return &Video{
		db: db,
	}
}

func (s *Video) CreateVideo(ctx context.Context, video *model.Video) error {
	return s.db.WithContext(ctx).Create(video).Error
}

func(s *Video) GetVideoListByLatestTime(ctx context.Context, latestTime int64, limit int)([]*model.Video, error){
	videos := make([]*model.Video, 0)
	if err := s.db.WithContext(ctx).Where("created_at < ?", latestTime).
		Order("created_at desc").Limit(limit).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func(s *Video)BatchGetVideoListByVideoId(ctx context.Context, videoIds []int64)([]*model.Video, error){
	videos := make([]*model.Video, 0)
	if err := s.db.WithContext(ctx).Where("id in ?", videoIds).
		Order("created_at desc").Find(&videos).Error; err != nil{
			return nil, err
	}
	return videos, nil
}

func(s *Video)GetVideoListByUserId(ctx context.Context, userId int64)([]*model.Video, error){
	videos := make([]*model.Video, 0)
	if err := s.db.WithContext(ctx).Where("user_id = ?", userId).
		Order("created_at desc").Find(&videos).Error; err != nil{
		return nil, err
	}
	return videos, nil
}

func(s *Video)GetVideoIdListByUserId(ctx context.Context, userId int64)([]int64, error){
	videoIdList := make([]int64, 0)
	if err := s.db.WithContext(ctx).Model(&model.Video{}).Where("user_id = ?", userId).
		Order("created_at desc").Select("id").Find(&videoIdList).Error;err != nil{
			return nil, err
	}
	return videoIdList, nil
}