package pkg

import (
	"context"
	"douyin/cmd/rpc/interaction/dao"
	"douyin/cmd/rpc/interaction/mq"
	"douyin/shared/constant"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SubscribeFavorite(subscriber *mq.FavoriteSubscriber, dao *dao.Favorite) error {
	favoriteChan, closeFunc, err := subscriber.Subscribe()
	defer closeFunc()
	if err != nil {
		return err
	}
	for favorite := range favoriteChan {
		if favorite.ActionType == constant.ActionTypeFavorite {
			if err = dao.CreateFavorite(context.Background(), favorite.Favorite); err != nil {
				klog.Errorf("create favorite error: %s", err.Error())
			}
		} else if favorite.ActionType == constant.ActionTypeCancelFavorite {
			if err = dao.DeleteFavorite(context.Background(), favorite.Favorite.UserId, favorite.Favorite.VideoId); err != nil {
				klog.Errorf("cancel favorite error: %s", err.Error())
			}
		} else {
			klog.Errorf("action type error")
		}
	}
	return nil
}

func SubscribeComment(subscriber *mq.CommentSubscriber, dao *dao.Comment) error {
	commentChan, closeFunc, err := subscriber.Subscribe()
	defer closeFunc()
	if err != nil {
		return err
	}
	for comment := range commentChan {
		if comment.ActionType == constant.ActionTypeComment {
			if err = dao.CreateComment(context.Background(), comment.Comment); err != nil {
				klog.Errorf("create comment error: %s", err.Error())
			}
		} else if comment.ActionType == constant.ActionTypeDeleteComment {
			if err = dao.DeleteComment(context.Background(), comment.Comment.VideoId, comment.Comment.ID); err != nil {
				klog.Errorf("delete comment error: %s", err.Error())
			}
		} else {
			klog.Errorf("action type error")
		}
	}
	return nil
}
