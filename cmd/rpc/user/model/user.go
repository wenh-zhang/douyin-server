package model

import "douyin/shared/constant"

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackGroundImage string `json:"back_ground_image"`
	Signature       string `json:"signature"`
}

func (*User) TableName() string {
	return constant.UserTableName
}
