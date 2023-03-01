package pkg

import (
	"context"
	"douyin/cmd/rpc/video/dao"
	"douyin/cmd/rpc/video/mq"

	"github.com/cloudwego/kitex/pkg/klog"
)

func SubscribeVideo(subscriber *mq.Subscriber, dao *dao.Video) error {
	videoChan, closeFunc, err := subscriber.Subscribe()
	defer closeFunc()
	if err != nil {
		return err
	}
	for video := range videoChan {
		if err = dao.CreateVideo(context.Background(), video); err != nil {
			klog.Errorf("create video error: %s", err.Error())
		}
	}
	return nil
}
