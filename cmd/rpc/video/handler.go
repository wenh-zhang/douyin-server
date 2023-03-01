package main

import (
	"context"
	"douyin/cmd/rpc/video/dao"
	"douyin/cmd/rpc/video/global"
	"douyin/cmd/rpc/video/model"
	"douyin/cmd/rpc/video/mq"
	"douyin/cmd/rpc/video/pkg"
	"douyin/shared/constant"
	"douyin/shared/errno"
	"douyin/shared/kitex_gen/base"
	code "douyin/shared/kitex_gen/errno"
	"douyin/shared/kitex_gen/interaction"
	"douyin/shared/kitex_gen/user"
	"douyin/shared/kitex_gen/video"
	"douyin/shared/util"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct {
	Dao       *dao.Video
	Publisher *mq.Publisher
}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	resp = new(video.DouyinFeedResponse)
	videoList, err := s.Dao.GetVideoListByLatestTime(ctx, req.LatestTime, constant.FeedLimit)
	if err != nil {
		klog.Errorf("feed error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("feed error"))
		return resp, nil
	}
	nextTime := time.Now().Unix()
	if len(videoList) > 0 {
		nextTime = videoList[len(videoList)-1].CreatedAt
	}

	videoIdList := make([]int64, 0)
	for _, video := range videoList {
		videoIdList = append(videoIdList, video.ID)
	}
	userIdList := make([]int64, 0)
	for _, video := range videoList {
		userIdList = append(userIdList, video.UserID)
	}
	videos, err := s.getVideoList(ctx, req.UserId, videoList)
	if err != nil {
		klog.Errorf("get video info error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("get video info error"))
		return resp, nil
	}
	resp.VideoList = videos
	resp.NextTime = nextTime
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
	resp = new(video.DouyinPublishActionResponse)
	sf, err := snowflake.NewSnowflake(constant.SnowFlakeDataCenterId, constant.VideoSnowFlakeWorkerId)
	if err != nil {
		klog.Errorf("snow flake generate error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("snow flake generate error"))
		return resp, nil
	}
	video := &model.Video{
		ID:       sf.NextVal(),
		UserID:   req.UserId,
		PlayURL:  req.PlayUrl,
		CoverURL: req.CoverUrl,
		Title:    req.Title,
	}
	if err = s.Publisher.Publish(video); err != nil {
		klog.Errorf("create video error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("create video error"))
		return resp, nil
	}
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

func (s *VideoServiceImpl) getVideoList(ctx context.Context, userId int64, videoList []*model.Video) ([]*base.Video, error) {
	videoIdList := make([]int64, 0)
	for _, video := range videoList {
		videoIdList = append(videoIdList, video.ID)
	}
	userIdList := make([]int64, 0)
	for _, video := range videoList {
		userIdList = append(userIdList, video.UserID)
	}

	interactionResp, err := global.InteractionClient.BatchGetVideoInteractInfo(ctx, &interaction.DouyinBatchGetVideoInteractInfoRequest{
		UserId:      userId,
		VideoIdList: videoIdList,
	})
	if err != nil {
		return nil, err
	}
	if interactionResp.BaseResp.StatusCode != int32(code.Code_Success) {
		return nil, errno.InteractionServerErr
	}
	userResp, err := global.UserClient.BatchGetUserInfo(ctx, &user.DouyinBatchGetUserRequest{
		LocalUserId:      userId,
		TargetUserIdList: userIdList,
	})
	if err != nil {
		return nil, err
	}
	if userResp.BaseResp.StatusCode != int32(code.Code_Success) {
		return nil, errno.UserServerErr
	}
	return pkg.BatchPackVideo(videoList, interactionResp.VideoInteractInfoList, userResp.UserList), nil
}

// GetPublishedVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishedVideoList(ctx context.Context, req *video.DouyinGetPublishListRequest) (resp *video.DouyinGetPublishListResponse, err error) {
	resp = new(video.DouyinGetPublishListResponse)
	videoList, err := s.Dao.GetVideoListByUserId(ctx, req.TargetUserId)
	if err != nil {
		klog.Errorf("get publish video list error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("get publish video list error"))
		return resp, nil
	}
	videos, err := s.getVideoList(ctx, req.LocalUserId, videoList)
	if err != nil {
		klog.Errorf("get video info error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("get video info error"))
		return resp, nil
	}
	resp.VideoList = videos
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoList(ctx context.Context, req *video.DouyinGetFavoriteListRequest) (resp *video.DouyinGetFavoriteListResponse, err error) {
	resp = new(video.DouyinGetFavoriteListResponse)
	videoIdResp, err := global.InteractionClient.GetFavoriteVideoIdList(ctx, &interaction.DouyinGetFavoriteVideoIdListRequest{
		UserId: req.TargetUserId,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCInteractionErr)
		return resp, nil
	}
	if videoIdResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = videoIdResp.BaseResp
		return resp, nil
	}
	videoList, err := s.Dao.BatchGetVideoListByVideoId(ctx, videoIdResp.VideoIdList)
	if err != nil {
		klog.Errorf("get video list error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("get video list error"))
		return resp, nil
	}

	videos, err := s.getVideoList(ctx, req.LocalUserId, videoList)
	if err != nil {
		klog.Errorf("get video info error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("get video info error"))
		return resp, nil
	}
	resp.VideoList = videos
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetPublishVideoIdList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishVideoIdList(ctx context.Context, req *video.DouyinGetPublishVideoIdListRequest) (resp *video.DouyinGetPublishVideoIdListResponse, err error) {
	resp = new(video.DouyinGetPublishVideoIdListResponse)
	videoIdList, err := s.Dao.GetVideoIdListByUserId(ctx, req.UserId)
	if err != nil {
		klog.Errorf("get video id list error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.VideoServerErr.WithMessage("get video id list error"))
		return resp, nil
	}
	resp.VideoIdList = videoIdList
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}
