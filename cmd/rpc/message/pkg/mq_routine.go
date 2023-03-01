package pkg

import (
	"context"
	"douyin/cmd/rpc/message/dao"
	"github.com/cloudwego/kitex/pkg/klog"

	"douyin/cmd/rpc/message/mq"
)

func SubscribeMessage(subscriber *mq.Subscriber, dao *dao.Message) error {
	messageChan, closeFunc, err := subscriber.Subscribe()
	defer closeFunc()
	if err != nil {
		return err
	}
	for message := range messageChan {
		if err = dao.CreateMessage(context.Background(), message); err != nil {
			klog.Errorf("create message error: %s", err.Error())
		}
	}
	return nil
}
