package main

import (
	"context"
	"douyin/cmd/rpc/message/dao"
	"douyin/cmd/rpc/message/model"
	"douyin/cmd/rpc/message/pkg"
	"douyin/shared/errno"
	"douyin/shared/kitex_gen/message"
	"douyin/shared/util"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct {
	Dao *dao.Message
}

// SendMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMessage(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	resp = new(message.DouyinMessageActionResponse)
	if err = s.Dao.CreateMessage(ctx, &model.Message{
		FromUserId: req.LocalUserId,
		ToUserId:   req.TargetUserId,
		Content:    req.Content,
		CreatedAt:  time.Now().Unix(),
	}); err != nil {
		klog.Errorf("create message error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.MessageServerErr.WithMessage("create message error"))
		return resp, nil
	}
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetMessageHistory implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetMessageHistory(ctx context.Context, req *message.DouyinGetMessageChatRequest) (resp *message.DouyinGetMessageChatResponse, err error) {
	resp = new(message.DouyinGetMessageChatResponse)
	messageList, err := s.Dao.GetMessageListByUserId(ctx, req.LocalUserId, req.TargetUserId, req.PreMsgTime)
	if err != nil {
		klog.Errorf("get message history error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.MessageServerErr.WithMessage("get message history error"))
		return resp, nil
	}
	resp.MessageList = pkg.BatchMessage(messageList)
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// BatchGetLatestMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) BatchGetLatestMessage(ctx context.Context, req *message.DouyinBatchGetLatestMessageRequest) (resp *message.DouyinBatchGetLatestMessageResponse, err error) {
	resp = new(message.DouyinBatchGetLatestMessageResponse)
	messageList, err := s.Dao.BatchGetLatestMsgByUserId(ctx, req.LocalUserId, req.TargetUserIdList)
	if err != nil && err != gorm.ErrRecordNotFound {
		klog.Errorf("get latest message error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.MessageServerErr.WithMessage("get latest message error"))
		return resp, nil
	}
	resp.LatestMsgList = pkg.BatchLatestMsg(messageList, req.LocalUserId)
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}
