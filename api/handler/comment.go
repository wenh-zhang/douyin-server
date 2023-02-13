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

// CommentActionResponse The returned json format of comment action request
type CommentActionResponse struct {
	Response
	Comment *interact.Comment `json:"comment,omitempty"`
}

// CommentListResponse  The returned json format of comment list request
type CommentListResponse struct {
	Response
	CommentList []*interact.Comment `json:"comment_list"`
}

// CommentAction  Handler of request for user to post a comment
// user is supposed to log in before sending this request, which means user id can be parsed from token
func CommentAction(ctx context.Context, c *app.RequestContext) {
	// get parameters from request context
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

	// call RPC service
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

// CommentList Handler of request for user to view list of video comments
// token may be empty, because users who are not logged in can also view the list
// this method has no authentication middleware before it
func CommentList(ctx context.Context, c *app.RequestContext) {
	request := new(interact.DouyinCommentListRequest)
	videoIDStr := c.Query(constant.VideoIdentityKey)
	videoID, err := strconv.ParseInt(videoIDStr, 10, 64)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	request.VideoId = videoID

	// parse the token to find out whether there is user id in it
	token := c.Query(constant.Token)
	if len(token) != 0 {
		claims, err := util.ParseToken(token)
		if err == nil {
			request.UserId = claims.UserID
		}
	}

	// call RPC service
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
