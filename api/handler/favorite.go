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

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
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
	if err = rpc.FavoriteAction(context.Background(), &interact.DouyinFavoriteActionRequest{
		UserId:     userID.(int64),
		VideoId:    videoID,
		ActionType: int32(actionType),
	}); err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, NewResponse(errno.Success))
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	request := new(interact.DouyinFavoriteListRequest)
	queryUserIDStr := c.Query(constant.UserIdentityKey)
	queryUserID, err := strconv.ParseInt(queryUserIDStr, 10, 64)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	request.QueryUserId = queryUserID

	token := c.Query(constant.Token)
	if len(token) != 0{
		claims, err := util.ParseToken(token)
		if err == nil{
			request.UserId = claims.UserID
		}
	}

	videos, err := rpc.FavoriteList(context.Background(), request)
	if err != nil{
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, VideoListResponse{
		Response: NewResponse(errno.Success),
		VideoList: videos,
	})
}