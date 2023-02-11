package handler

import (
	"context"
	"douyin/api/rpc"
	"douyin/kitex_gen/interact"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

type CommentActionResponse struct {
	Response
	Comment *interact.Comment `json:"comment,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []*interact.Comment `json:"comment_list"`
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	userID, _ := c.Get(constant.TokenUserIdentifyKey)
	videoIDStr := c.Query(constant.VideoIdentityKey)
	videoID, err := strconv.ParseInt(videoIDStr, 10, 64)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	actionTypeStr := c.Query(constant.ActionType)
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 32)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	var commentID *int64
	var commentText *string
	commentIDStr := c.Query(constant.CommentIdentityKey)
	if len(commentIDStr) != 0 {
		id, err := strconv.ParseInt(commentIDStr, 10, 64)
		if err == nil {
			commentID = &id
		}
	}
	commentTextStr := c.Query(constant.CommentText)
	if len(commentTextStr) != 0 {
		commentText = &commentTextStr
	}

	comment, err := rpc.CommentAction(context.Background(), &interact.DouyinCommentActionRequest{
		UserId:      userID.(int64),
		VideoId:     videoID,
		ActionType:  int32(actionType),
		CommentText: commentText,
		CommentId:   commentID,
	})
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, CommentActionResponse{
		Response: NewResponse(errno.Success),
		Comment:  comment,
	})
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	request := new(interact.DouyinCommentListRequest)
	videoIDStr := c.Query(constant.VideoIdentityKey)
	videoID, err := strconv.ParseInt(videoIDStr, 10, 64)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	request.VideoId = videoID

	token := c.Query(constant.Token)
	if len(token) != 0 {
		claims, err := util.ParseToken(token)
		if err == nil {
			request.UserId = claims.UserID
		}
	}
	comments, err := rpc.CommentList(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, CommentListResponse{
		Response:    NewResponse(errno.Success),
		CommentList: comments,
	})
}
