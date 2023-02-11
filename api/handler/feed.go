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

type FeedResponse struct {
	VideoListResponse
	NextTime  int64         `json:"next_time,omitempty"`
}

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

	token := c.Query(constant.Token)
	if len(token) != 0{
		claims, err := util.ParseToken(token)
		if err == nil{
			request.UserId = &claims.UserID
		}
	}

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
