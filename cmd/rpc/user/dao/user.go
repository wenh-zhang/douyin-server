package dao

import (
	"context"
	"douyin/cmd/rpc/user/model"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (s *User) BatchGetUserById(ctx context.Context, userIds []int64) ([]*model.User, error) {
	if userIds == nil {
		return nil, nil
	}
	users := make([]*model.User, 0)
	for _, userId := range userIds {
		user := new(model.User)
		if err := s.db.WithContext(ctx).Where("id = ?", userId).First(&user).Error; err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *User) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user := new(model.User)
	if err := s.db.WithContext(ctx).Model(&model.User{}).Where("name = ?", name).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *User) CreateUser(ctx context.Context, user *model.User) error {
	return s.db.WithContext(ctx).Create(user).Error
}
