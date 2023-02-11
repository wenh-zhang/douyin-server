package service

import (
	"context"
	"douyin/dal"
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"douyin/pkg/constant"
	"douyin/rpc/common"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx}
}

func (s *FeedService) Feed(req *core.DouyinFeedRequest) ([]*core.Video, *int64, error) {
	conn := dal.GetConn()
	videos, err := db.MGetVideosByTime(s.ctx, conn, *req.LatestTime, constant.FeedLimit)
	if err != nil {
		return nil, nil, err
	}
	convertVideos, nextTime := common.Videos(videos)
	// check if user like the videos
	if err = common.CheckIfLikeVideo(s.ctx, conn, convertVideos, req.UserId); err != nil {
		return nil, nil, err
	}
	//check author info of each video
	if err = common.CheckAuthorOfVideo(s.ctx, conn, convertVideos, req.UserId); err != nil {
		return nil, nil, err
	}
	return convertVideos, nextTime, nil
}
