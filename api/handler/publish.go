package handler

import (
	"context"
	"douyin/api/rpc"
	"douyin/api/service"
	"douyin/kitex_gen/core"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// VideoListResponse The returned json format of video list request
// including request of favorite list, publish list and feed list
type VideoListResponse struct {
	Response
	VideoList []*core.Video `json:"video_list,omitempty"`
}

// PublishAction The handler of request for user to publish a video
// user is supposed to log in before sending this request, which means user id can be parsed from token
func PublishAction(ctx context.Context, c *app.RequestContext) {
	userID, _ := c.Get(constant.TokenUserIdentifyKey)
	title := c.PostForm(constant.Title)
	playURL, coverURL, err := service.Upload(userID.(int64), c)
	if err != nil {
		SendResponse(c, NewResponse(err))
	}

	// call RPC service
	if err = rpc.PublishAction(context.Background(), &core.DouyinPublishActionRequest{
		UserId:   userID.(int64),
		PlayUrl:  playURL,
		CoverUrl: coverURL,
		Title:    title,
	}); err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, NewResponse(errno.Success))
}

// PublishList The handler of request for user to view a list of videos which are published by another user
// token may be empty, because users who are not logged in can also view the list
// this method has no authentication middleware before it
func PublishList(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinPublishListRequest)
	queryUserIDStr := c.Query(constant.UserIdentityKey)
	queryUserID, err := strconv.ParseInt(queryUserIDStr, 10, 64)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	request.QueryUserId = queryUserID

	token := c.Query(constant.Token)
	if len(token) != 0 {
		claims, err := util.ParseToken(token)
		if err == nil {
			request.UserId = claims.UserID
		}
	}

	videos, err := rpc.PublishList(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, VideoListResponse{
		Response:  NewResponse(errno.Success),
		VideoList: videos,
	})
}
