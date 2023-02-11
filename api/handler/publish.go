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

type VideoListResponse struct {
	Response
	VideoList []*core.Video `json:"video_list,omitempty"`
}

func PublishAction(ctx context.Context, c *app.RequestContext) {
	userID, _ := c.Get(constant.TokenUserIdentifyKey)
	title := c.PostForm(constant.Title)
	playURL, coverURL, err := service.Upload(userID.(int64), c)
	if err != nil {
		SendResponse(c, NewResponse(err))
	}

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
