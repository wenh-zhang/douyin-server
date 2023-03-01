package model

import "douyin/shared/constant"

type Message struct {
	ID         int64  `json:"id"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreatedAt  int64  `json:"created_at"`
}

func (*Message) TableName() string {
	return constant.MessageTableName
}
