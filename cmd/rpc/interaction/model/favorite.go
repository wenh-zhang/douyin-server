package model

import "douyin/shared/constant"

type Favorite struct {
	UserId    int64 `json:"user_id"`
	VideoId   int64 `json:"video_id"`
	CreatedAt int64 `json:"created_at"`
}

func (*Favorite) TableName() string {
	return constant.FavoriteTableName
}
