package handler

import (
	"context"
	"douyin/api/rpc"
	"douyin/kitex_gen/core"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"time"
)

// FeedResponse The returned json format of feed videos
type FeedResponse struct {
	VideoListResponse
	NextTime int64 `json:"next_time,omitempty"`
}

// Feed Handler of request for applying videos information
// token may be empty, because users who are not logged in can also watch videos
// the return videos should be ordered by time desc
// this method has no authentication middleware before it
func Feed(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinFeedRequest)
	latestTime := time.Now().Unix()
	var err error
	latestTimeStr := c.Query(constant.LatestTime)
	if len(latestTimeStr) != 0 {
		latestTime, err = strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			SendResponse(c, NewResponse(err))
			return
		}
	}
	request.LatestTime = &latestTime

	// parse the token to find out whether there is user id in it
	token := c.Query(constant.Token)
	if len(token) != 0 {
		claims, err := util.ParseToken(token)
		if err == nil {
			request.UserId = &claims.UserID
		}
	}

	// call RPC service
	videos, nextTime, err := rpc.Feed(context.Background(), request)
	if err != nil {
		SendResponse(c, NewResponse(err))
		return
	}
	var response FeedResponse
	response.Response = NewResponse(errno.Success)
	if videos != nil {
		response.VideoList = videos
	}
	if nextTime != nil {
		response.NextTime = *nextTime
	}
	SendResponse(c, response)
}
