package db

import (
	"context"
	"douyin/pkg/constant"
	"gorm.io/gorm"
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
	return constant.CommentTableName
}

// MGetComments multiple get list of comments info
func MGetComments(ctx context.Context, db *gorm.DB, commentIDs []int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := db.WithContext(ctx).Where("id = ?", commentIDs).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// MGetCommentsByVideoID multiple get list of comments info
func MGetCommentsByVideoID(ctx context.Context, db *gorm.DB, videoID int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := db.WithContext(ctx).Where("video_id = ?", videoID).Order("created_at desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateComment create comment info
func CreateComment(ctx context.Context, db *gorm.DB, comments []*Comment) error {
	return db.WithContext(ctx).Create(comments).Error
}

// DeleteComment delete comment info
func DeleteComment(ctx context.Context, db *gorm.DB, commentIDs []int64) error {
	return db.WithContext(ctx).Where("id in ?", commentIDs).Delete(&Comment{}).Error
}
