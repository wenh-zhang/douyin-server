package service

import (
	"context"
	"crypto/md5"
	"douyin/dal"
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"douyin/pkg/errno"
	"douyin/rpc/common"
	"fmt"
	"io"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx}
}

func (s *UserService) Register(req *core.DouyinUserRegisterRequest) (int64, error) {
	conn := dal.GetConn()
	users, err := db.QueryUser(s.ctx, conn, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	users = []*db.User{&db.User{
		UserName: req.Username,
		Password: password,
	}}
	err = db.CreateUser(s.ctx, conn, users)
	if err != nil {
		return 0, err
	}
	return int64(users[0].ID), nil
}

func (s *UserService) Login(req *core.DouyinUserLoginRequest) (int64, error) {
	conn := dal.GetConn()
	users, err := db.QueryUser(s.ctx, conn, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	if password != users[0].Password {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(users[0].ID), nil
}

func (s *UserService) Info(req *core.DouyinUserRequest) (*core.User, error) {
	conn := dal.GetConn()
	userMap, err := common.CheckAuthorFromUserID(s.ctx, conn, []int64{req.QueryUserId}, &req.UserId)
	if err != nil {
		return nil, err
	}
	return userMap[req.QueryUserId], nil
}
