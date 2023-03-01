package pkg

import (
	"context"
	"douyin/cmd/rpc/user/dao"
	"douyin/cmd/rpc/user/mq"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SubscribeUser(subscriber *mq.Subscriber, dao *dao.User) error {
	userChan, closeFunc, err := subscriber.Subscribe()
	defer closeFunc()
	if err != nil {
		return err
	}
	for user := range userChan {
		if err = dao.CreateUser(context.Background(), user); err != nil {
			klog.Errorf("create user error: %s", err.Error())
		}
	}
	return nil
}
