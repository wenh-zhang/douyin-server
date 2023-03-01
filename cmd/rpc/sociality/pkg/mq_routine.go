package pkg

import (
	"context"
	"douyin/cmd/rpc/sociality/dao"
	"douyin/cmd/rpc/sociality/mq"
	"douyin/shared/constant"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SubscribeFollow(subscriber *mq.Subscriber, dao *dao.Follow) error {
	followChan, closeFunc, err := subscriber.Subscribe()
	defer closeFunc()
	if err != nil {
		return err
	}
	for follow := range followChan {
		if follow.ActionType == constant.ActionTypeFollow {
			if err = dao.CreateFollow(context.Background(), follow.Follow); err != nil {
				klog.Errorf("create follow error: %s", err.Error())
			}
		} else if follow.ActionType == constant.ActionTypeCancelFollow {
			if err = dao.DeleteFollow(context.Background(), follow.Follow.FromUserID, follow.Follow.ToUserID); err != nil {
				klog.Errorf("cancel follow error: %s", err.Error())
			}
		} else {
			klog.Errorf("action type error")
		}
	}
	return nil
}
