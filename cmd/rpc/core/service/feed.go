package service

import (
	"context"
	"douyin/cmd/rpc/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"douyin/pkg/constants"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx}
}

func (s *FeedService) Feed(req *core.DouyinFeedRequest) ([]*core.Video, *int64, error) {
	videos, err := db.MGetVideosByTime(s.ctx, *req.LatestTime, constants.FeedLimit)
	if err != nil {
		return nil, nil, err
	}
	convertVideos, nextTime := pack.Videos(videos)
	return convertVideos, nextTime, nil
}
