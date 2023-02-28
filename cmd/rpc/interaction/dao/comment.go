package dao

import (
	"context"
	"douyin/cmd/rpc/interaction/model"
	"gorm.io/gorm"
)

type Comment struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) *Comment {
	return &Comment{
		db: db,
	}
}

func (s *Comment) CreateComment(ctx context.Context, comment *model.Comment) error {
	return s.db.WithContext(ctx).Create(comment).Error
}

func (s *Comment) DeleteComment(ctx context.Context, videoId, commentId int64) error {
	return s.db.WithContext(ctx).Delete(&model.Comment{
		ID: commentId,
	}).Error
}

// GetCommentListByVideoId 通过视频id获取评论列表
func (s *Comment) GetCommentListByVideoId(ctx context.Context, videoId int64) ([]*model.Comment, error) {
	comments := make([]*model.Comment, 0)
	if err := s.db.WithContext(ctx).Where("video_id = ?", videoId).
		Order("created_at desc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

// BatchGetCommentCountByVideoId 通过视频id列表批量获取视频评论数列表
func (s *Comment) BatchGetCommentCountByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	counts := make([]int64, 0)
	var count int64
	for _, videoId := range videoIds {
		if err := s.db.WithContext(ctx).Model(&model.Comment{}).
			Where("video_id = ?", videoId).Count(&count).Error; err != nil {
			return nil, err
		}
		counts = append(counts, count)
	}
	return counts, nil
}
