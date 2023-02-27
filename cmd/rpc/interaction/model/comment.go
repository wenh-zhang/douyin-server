package model

import "douyin/shared/constant"

type Comment struct {
	ID        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	VideoId   int64  `json:"video_id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

func (*Comment) TableName() string {
	return constant.CommentTableName
}
