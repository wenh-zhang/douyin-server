package db

import (
	"context"
	"douyin/dal"
	"douyin/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	Avatar        string `json:"avatar"`
}

func (*User) TableName() string {
	return constants.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := dal.DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return dal.DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := dal.DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateUser update user info
func UpdateUser(ctx context.Context, userID int64, userName, password *string, followCount, followerCount int64, avatar *string) error {
	params := map[string]interface{}{}
	if userName != nil {
		params["user_name"] = *userName
	}
	if password != nil {
		params["password"] = *password
	}
	if avatar != nil {
		params["avatar"] = *avatar
	}
	params["follow_count"] = followCount
	params["follower_count"] = followerCount
	return dal.DB.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Updates(params).Error
}

// DeleteUser delete user info
func DeleteUser(ctx context.Context, userID int64) error {
	return dal.DB.WithContext(ctx).Where("id = ?", userID).Delete(&User{}).Error
}
