package model

import (
	"douyin/shared/constant"
)

type Video struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	PlayURL   string `json:"play_url"`
	CoverURL  string `json:"cover_url"`
	Title     string `json:"title"`
	CreatedAt int64 `json:"created_at"`
}

func (*Video) TableName() string {
	return constant.VideoTableName
}
