package service

import (
	"context"
	"douyin/dal"
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"os"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	id, err := NewUserService(context.Background()).Register(&core.DouyinUserRegisterRequest{
		Username: "zhangwenhua",
		Password: "password",
	})
	if err != nil {
		t.Errorf("error msg: %s", err.Error())
		return
	}
	users, err := db.MGetUsers(context.Background(), dal.GetNewConn(), []int64{id})
	if err != nil {
		t.Errorf("error msg: %s", err.Error())
		return
	}
	if len(users) != 1 {
		t.Errorf("user register fail")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
