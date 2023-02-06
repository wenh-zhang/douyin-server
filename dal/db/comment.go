package db

import (
	"context"
	"douyin/dal"
	"douyin/pkg/constants"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	VideoID   int64     `json:"video_id"`
	UserID    int64     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (*Comment) TableName() string {
	return constants.CommentTableName
}

// MGetComments multiple get list of comments info
func MGetComments(ctx context.Context, commentIDs []int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := dal.DB.WithContext(ctx).Where("id = ?", commentIDs).Order("created_at desc").Find(&Comment{}).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetCommentsByVideoID multiple get list of comments info
func MGetCommentsByVideoID(ctx context.Context, videoID int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := dal.DB.WithContext(ctx).Where("video_id = ?", videoID).Order("created_at desc").Find(&Comment{}).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateComment create comment info
func CreateComment(ctx context.Context, comments []*Comment) error {
	return dal.DB.WithContext(ctx).Create(comments).Error
}

// DeleteComment delete comment info
func DeleteComment(ctx context.Context, favoriteID int64) error {
	return dal.DB.WithContext(ctx).Where("id = ?", favoriteID).Delete(&Comment{}).Error
}
