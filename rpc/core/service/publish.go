package service

import (
	"context"
	"douyin/dal"
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"douyin/rpc/common"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{ctx}
}

func (s *PublishService) PublishAction(req *core.DouyinPublishActionRequest) error {
	video := &db.Video{
		UserID:  req.UserId,
		PlayURL: req.PlayUrl,
		CoverURL: req.CoverUrl,
		Title:   req.Title,
	}
	if err := db.CreateVideo(s.ctx, dal.GetConn(), []*db.Video{video}); err != nil {
		return err
	}
	return nil
}

func (s *PublishService) PublishList(req *core.DouyinPublishListRequest) ([]*core.Video, error) {
	conn := dal.GetConn()
	videos, err := db.MGetVideosByUserID(s.ctx, conn, []int64{req.QueryUserId})
	if err != nil {
		return nil, err
	}
	convertVideos, _ := common.Videos(videos)
	// check if user like the videos
	if err = common.CheckIfLikeVideo(s.ctx, conn, convertVideos, &req.UserId); err != nil{
		return nil, err
	}
	// check author information, only need to query once because all videos have same author
	if err = common.CheckAuthorOfVideo(s.ctx, conn, convertVideos, &req.UserId); err != nil{
		return nil, err
	}
	return convertVideos, nil
}
