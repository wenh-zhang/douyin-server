package model

import "douyin/shared/constant"

type Follow struct {
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
}

func (*Follow) TableName() string {
	return constant.FollowTableName
}

