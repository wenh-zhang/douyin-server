package dao

import (
	"context"
	"douyin/cmd/rpc/sociality/model"
	"gorm.io/gorm"
)

type Follow struct {
	db *gorm.DB
}

func NewFollow(db *gorm.DB) *Follow {
	return &Follow{
		db: db,
	}
}

func (s *Follow) CreateFollow(ctx context.Context, follow *model.Follow) error {
	return s.db.WithContext(ctx).Create(follow).Error
}

func (s *Follow) DeleteFollow(ctx context.Context, fromUserId, toUserId int64) error {
	return s.db.WithContext(ctx).Where("from_user_id = ? and to_user_id = ?", fromUserId, toUserId).
		Delete(&model.Follow{}).Error
}

func (s *Follow) GetFollowIdList(ctx context.Context, userId int64) ([]int64, error) {
	follows := make([]int64, 0)
	if err := s.db.WithContext(ctx).Model(&model.Follow{}).Where("from_user_id = ?", userId).
		Select("to_user_id").Find(&follows).Error; err != nil {
		return nil, err
	}
	return follows, nil
}

func (s *Follow) GetFollowerIdList(ctx context.Context, userId int64) ([]int64, error) {
	followers := make([]int64, 0)
	if err := s.db.WithContext(ctx).Model(&model.Follow{}).Where("to_user_id = ?", userId).
		Select("from_user_id").Find(&followers).Error; err != nil {
		return nil, err
	}
	return followers, nil
}

func (s *Follow) GetFriendIdList(ctx context.Context, userId int64) ([]int64, error) {
	follows := make([]int64, 0)
	friends := make([]int64, 0)
	if err := s.db.WithContext(ctx).Model(&model.Follow{}).Where("from_user_id = ?", userId).
		Select("to_user_id").Find(&follows).Error; err != nil {
		return nil, err
	}
	if err := s.db.WithContext(ctx).Model(&model.Follow{}).Where("from_user_id in ? and to_user_id = ?", follows, userId).
		Select("from_user_id").Find(&friends).Error; err != nil {
		return nil, err
	}
	return friends, nil
}

func (s *Follow) BatchGetFollowCountByUserId(ctx context.Context, userIds []int64) ([]int64, error) {
	counts := make([]int64, 0)
	var count int64
	for _, userId := range userIds {
		if err := s.db.WithContext(ctx).Model(&model.Follow{}).Where("from_user_id = ?", userId).Count(&count).Error; err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}

func (s *Follow) BatchGetFollowerCountByUserId(ctx context.Context, userIds []int64) ([]int64, error) {
	counts := make([]int64, 0)
	var count int64
	for _, userId := range userIds {
		if err := s.db.WithContext(ctx).Model(&model.Follow{}).Where("to_user_id = ?", userId).Count(&count).Error; err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}

func (s *Follow) BatchGetFollowInfoByUserId(ctx context.Context, localUserId int64, targetUserIds []int64) ([]bool, error) {
	userIdList := make([]int64, 0)
	if err := s.db.WithContext(ctx).Model(&model.Follow{}).
		Where("from_user_id = ? and to_user_id in ?", localUserId, targetUserIds).Select("to_user_id").Find(&userIdList).Error; err != nil {
		return nil, err
	}
	umap := make(map[int64]struct{})
	for _, id := range userIdList {
		umap[id] = struct{}{}
	}
	res := make([]bool, 0)
	for _, id := range targetUserIds {
		if _, ok := umap[id]; ok {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res, nil
}
