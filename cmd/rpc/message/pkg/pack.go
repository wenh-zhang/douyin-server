package pkg

import (
	"douyin/cmd/rpc/message/model"
	"douyin/shared/constant"
	"douyin/shared/kitex_gen/base"
)

func Message(message *model.Message) *base.Message {
	return &base.Message{
		Id:         message.ID,
		FromUserId: message.FromUserId,
		ToUserId:   message.ToUserId,
		Content:    message.Content,
		CreateTime: message.CreatedAt,
	}
}

func BatchMessage(messageList []*model.Message) []*base.Message {
	res := make([]*base.Message, 0)
	for _, message := range messageList {
		res = append(res, Message(message))
	}
	return res
}

func LatestMsg(message *model.Message, userId int64) *base.LatestMsg {
	msgType := constant.MsgTypeReceive
	if message.FromUserId == userId {
		msgType = constant.MsgTypeSend
	}
	return &base.LatestMsg{
		Message: message.Content,
		MsgType: int64(msgType),
	}
}

func BatchLatestMsg(messageList []*model.Message, userId int64) []*base.LatestMsg {
	res := make([]*base.LatestMsg, 0)
	for _, message := range messageList {
		res = append(res, LatestMsg(message, userId))
	}
	return res
}
