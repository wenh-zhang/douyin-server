package handler

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/core"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []*core.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(ctx context.Context, c *app.RequestContext) {
	request := new(core.DouyinFeedRequest)
	latestTime := time.Now().Unix()
	var err error
	latestTimeStr := c.Param(constants.LatestTime)
	if len(latestTimeStr) != 0 {
		latestTime, err = strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			SendResponse(c, NewResponse(err))
			return
		}
	}
	request.LatestTime = &latestTime

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
	//SendResponse(c, response)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
