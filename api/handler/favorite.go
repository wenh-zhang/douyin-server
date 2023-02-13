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

// FavoriteAction Handler of request for a user to like the video
// user is supposed to log in before sending this request, which means user id can be parsed from token
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

	// call RPC service
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

// FavoriteList Handler of request for user to view a list of videos which are liked by another user
// token may be empty, because users who are not logged in can also view the list
// this method has no authentication middleware before it
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	request := new(interact.DouyinFavoriteListRequest)
	queryUserIDStr := c.Query(constant.UserIdentityKey)
	queryUserID, err := strconv.ParseInt(queryUserIDStr, 10, 64)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	request.QueryUserId = queryUserID

	// parse the token to find out whether there is user id in it
	token := c.Query(constant.Token)
	if len(token) != 0 {
		claims, err := util.ParseToken(token)
		if err == nil {
			request.UserId = claims.UserID
		}
	}

	// call RPC service
	videos, err := rpc.FavoriteList(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	SendResponse(c, VideoListResponse{
		Response:  NewResponse(errno.Success),
		VideoList: videos,
	})
}
