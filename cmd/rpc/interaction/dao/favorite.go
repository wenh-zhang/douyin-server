package dao

import (
	"context"
	"douyin/cmd/rpc/interaction/model"
	"gorm.io/gorm"
)

type Favorite struct {
	db *gorm.DB
}

func NewFavorite(db *gorm.DB) *Favorite {
	return &Favorite{
		db: db,
	}
}

func (s *Favorite) CreateFavorite(ctx context.Context, favorite *model.Favorite) error {
	return s.db.WithContext(ctx).Create(favorite).Error
}

func (s *Favorite) DeleteFavorite(ctx context.Context, userId, videoId int64) error {
	return s.db.WithContext(ctx).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&model.Favorite{}).Error
}

// BatchGetFavoriteInfoByVideoId 获取用户对视频列表的点赞信息
func (s *Favorite) BatchGetFavoriteInfoByVideoId(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	fvVideoIdList := make([]int64, 0)
	if err := s.db.WithContext(ctx).Model(&model.Favorite{}).
		Where("user_id = ? and video_id in ?", userId, videoIds).Select("video_id").Find(&fvVideoIdList).Error; err != nil {
		return nil, err
	}
	umap := make(map[int64]struct{})
	for _, id := range fvVideoIdList {
		umap[id] = struct{}{}
	}
	res := make([]bool, 0)
	for _, id := range videoIds {
		if _, ok := umap[id]; ok {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res, nil
}

// BatchGetFavoriteCountByVideoId 批量获取视频的点赞数
func (s *Favorite) BatchGetFavoriteCountByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	counts := make([]int64, 0)
	var count int64
	for _, videoId := range videoIds {
		if err := s.db.WithContext(ctx).Model(&model.Favorite{}).
			Where("video_id = ?", videoId).Count(&count).Error; err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}

// BatchGetFavoriteCountSumByVideoId 批量获取视频的总点赞数
func (s *Favorite) BatchGetFavoriteCountSumByVideoId(ctx context.Context, videoIds []int64) (int64, error) {
	var count int64
	if err := s.db.WithContext(ctx).Model(&model.Favorite{}).
		Where("video_id in ?", videoIds).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// GetFavoriteVideoIdListByUserId 根据用户id获取其点赞的视频id列表
func (s *Favorite) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	videoIds := make([]int64, 0)
	if err := s.db.WithContext(ctx).Model(&model.Favorite{}).Where("user_id = ?", userId).
		Select("video_id").Order("created_at desc").Find(&videoIds).Error; err != nil {
		return nil, err
	}
	return videoIds, nil
}

func (s *Favorite) BatchGetFavoriteCountByUserId(ctx context.Context, userIds []int64) ([]int64, error) {
	counts := make([]int64, 0)
	var count int64
	for _, userId := range userIds {
		if err := s.db.WithContext(ctx).Model(&model.Favorite{}).
			Where("user_id = ?", userId).Count(&count).Error; err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}
