package main

import (
	"context"
	"douyin/cmd/rpc/interaction/dao"
	"douyin/cmd/rpc/interaction/global"
	"douyin/cmd/rpc/interaction/model"
	"douyin/cmd/rpc/interaction/pkg"
	"douyin/cmd/rpc/interaction/redis"
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

// InteractionServerImpl implements the last service interface defined in the IDL.
type InteractionServerImpl struct {
	FavoriteDao      *dao.Favorite
	CommentDao       *dao.Comment
	FavoriteRedisDao *redis.Favorite
	CommentRedisDao  *redis.Comment
}

// Favorite implements the InteractionServerImpl interface.
func (s *InteractionServerImpl) Favorite(ctx context.Context, req *interaction.DouyinFavoriteActionRequest) (resp *interaction.DouyinFavoriteActionResponse, err error) {
	resp = new(interaction.DouyinFavoriteActionResponse)
	favorite := &model.Favorite{
		UserId:    req.UserId,
		VideoId:   req.VideoId,
		CreatedAt: time.Now().Unix(),
	}
	if req.ActionType == constant.ActionTypeFavorite {
		if err = s.FavoriteDao.CreateFavorite(ctx, favorite); err != nil {
			klog.Errorf("create favorite error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("create favorite error"))
			return resp, nil
		}
		if err = s.FavoriteRedisDao.Like(ctx, favorite); err != nil {
			klog.Errorf("create favorite by redis error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("create favorite error"))
			return resp, nil
		}
	} else if req.ActionType == constant.ActionTypeCancelFavorite {
		if err = s.FavoriteDao.DeleteFavorite(ctx, req.UserId, req.VideoId); err != nil {
			klog.Errorf("cancel favorite error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("cancel favorite error"))
			return resp, nil
		}
		if err = s.FavoriteRedisDao.Unlike(ctx, req.UserId, req.VideoId); err != nil {
			klog.Errorf("cancel favorite error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("cancel favorite error"))
			return resp, nil
		}
	} else {
		resp.BaseResp = util.BuildBaseResp(errno.ParamErr.WithMessage("wrong action type"))
		return resp, nil
	}
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteVideoIdList implements the InteractionServerImpl interface.
func (s *InteractionServerImpl) GetFavoriteVideoIdList(ctx context.Context, req *interaction.DouyinGetFavoriteVideoIdListRequest) (resp *interaction.DouyinGetFavoriteVideoIdListResponse, err error) {
	resp = new(interaction.DouyinGetFavoriteVideoIdListResponse)
	videoIds, err := s.FavoriteRedisDao.GetFavoriteVideoIdListByUserId(ctx, req.UserId)
	if err != nil {
		klog.Errorf("get favorite video id list by redis error: %s", err.Error())
		videoIds, err = s.FavoriteDao.GetFavoriteVideoIdListByUserId(ctx, req.UserId)
		if err != nil {
			klog.Errorf("get favorite video id list error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("get favorite video id list error"))
			return resp, nil
		}
	}
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	resp.VideoIdList = videoIds
	return resp, nil
}

// Comment implements the InteractionServerImpl interface.
func (s *InteractionServerImpl) Comment(ctx context.Context, req *interaction.DouyinCommentActionRequest) (resp *interaction.DouyinCommentActionResponse, err error) {
	resp = new(interaction.DouyinCommentActionResponse)
	sf, err := snowflake.NewSnowflake(constant.SnowFlakeDataCenterId, constant.CommentSnowFlakeWorkerId)
	if err != nil {
		klog.Errorf("snow flake generate failed: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("generate user id failed"))
		return resp, nil
	}
	comment := &model.Comment{
		ID:        sf.NextVal(),
		VideoId:   req.VideoId,
		UserId:    req.UserId,
		Content:   req.CommentText,
		CreatedAt: time.Now().Unix(),
	}
	if req.ActionType == constant.ActionTypeComment {
		// 需要实时返回，后续不可放入mq中
		if err = s.CommentDao.CreateComment(ctx, comment); err != nil {
			klog.Errorf("create comment error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("create comment error"))
			return resp, nil
		}
		if err = s.CommentRedisDao.CreateComment(ctx, comment); err != nil {
			klog.Errorf("create comment by redis error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("create comment error"))
			return resp, nil
		}
		userResp, err := global.UserClient.BatchGetUserInfo(ctx, &user.DouyinBatchGetUserRequest{
			LocalUserId:      req.UserId,
			TargetUserIdList: []int64{req.UserId},
		})
		if err != nil {
			resp.BaseResp = util.BuildBaseResp(errno.RPCUserErr)
			return resp, nil
		}
		if userResp.BaseResp.StatusCode != int32(code.Code_Success) {
			resp.BaseResp = userResp.BaseResp
			return resp, nil
		}
		if len(userResp.UserList) == 0 {
			resp.BaseResp = util.BuildBaseResp(errno.RPCUserErr.WithMessage("get user info error"))
			return resp, nil
		}
		resp.Comment = pkg.PackComment(comment, userResp.UserList[0])
		resp.BaseResp = util.BuildBaseResp(errno.Success)
		return resp, nil
	} else if req.ActionType == constant.ActionTypeDeleteComment {
		if err = s.CommentDao.DeleteComment(ctx, req.VideoId, req.CommentId); err != nil {
			klog.Errorf("delete comment error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("delete comment error"))
			return resp, nil
		}
		if err = s.CommentRedisDao.DeleteComment(ctx, req.VideoId, req.CommentId); err != nil {
			klog.Errorf("delete comment by redis error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("delete comment error"))
			return resp, nil
		}
		resp.BaseResp = util.BuildBaseResp(errno.Success)
		return resp, nil
	} else {
		resp.BaseResp = util.BuildBaseResp(errno.ParamErr.WithMessage("action type error"))
		return resp, nil
	}
}

// GetCommentList implements the InteractionServerImpl interface.
func (s *InteractionServerImpl) GetCommentList(ctx context.Context, req *interaction.DouyinGetCommentListRequest) (resp *interaction.DouyinGetCommentListResponse, err error) {
	resp = new(interaction.DouyinGetCommentListResponse)
	commentList, err := s.CommentDao.GetCommentListByVideoId(ctx, req.VideoId)
	if err != nil {
		klog.Errorf("get comment list error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("get comment list error"))
		return resp, nil
	}
	userIdList := make([]int64, 0)
	for _, comment := range commentList {
		userIdList = append(userIdList, comment.UserId)
	}
	userResp, err := global.UserClient.BatchGetUserInfo(ctx, &user.DouyinBatchGetUserRequest{
		LocalUserId:      req.UserId,
		TargetUserIdList: userIdList,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCUserErr)
		return resp, nil
	}
	if userResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}
	resp.CommentList = pkg.BatchPackComment(commentList, userResp.UserList)
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// BatchGetVideoInteractInfo implements the InteractionServerImpl interface.
func (s *InteractionServerImpl) BatchGetVideoInteractInfo(ctx context.Context, req *interaction.DouyinBatchGetVideoInteractInfoRequest) (resp *interaction.DouyinBatchGetVideoInteractInfoResponse, err error) {
	resp = new(interaction.DouyinBatchGetVideoInteractInfoResponse)
	favoriteCounts, err := s.FavoriteRedisDao.BatchGetFavoriteCountByVideoId(ctx, req.VideoIdList)
	if err != nil {
		klog.Errorf("get favorite count by redis error: %s", err.Error())
		favoriteCounts, err = s.FavoriteDao.BatchGetFavoriteCountByVideoId(ctx, req.VideoIdList)
		if err != nil {
			klog.Errorf("get favorite count error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("get favorite count error"))
			return resp, nil
		}
	}

	commentCounts, err := s.CommentRedisDao.BatchGetCommentCountByVideoId(ctx, req.VideoIdList)
	if err != nil {
		klog.Errorf("get comment count by redis error: %s", err.Error())
		commentCounts, err = s.CommentDao.BatchGetCommentCountByVideoId(ctx, req.VideoIdList)
		if err != nil {
			klog.Errorf("get comment count error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("get comment count error"))
			return resp, nil
		}
	}

	isFavorites, err := s.FavoriteRedisDao.BatchGetFavoriteInfoByVideoId(ctx, req.UserId, req.VideoIdList)
	if err != nil {
		klog.Errorf("check if favorite by redis error: %s", err.Error())
		isFavorites, err = s.FavoriteDao.BatchGetFavoriteInfoByVideoId(ctx, req.UserId, req.VideoIdList)
		if err != nil {
			klog.Errorf("check if favorite error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("check if favorite error"))
			return resp, nil
		}
	}

	interactInfos := make([]*base.VideoInteractInfo, 0)
	for i := range req.VideoIdList {
		interactInfo := &base.VideoInteractInfo{
			FavoriteCount: favoriteCounts[i],
			CommentCount:  commentCounts[i],
			IsFavorite:    isFavorites[i],
		}
		interactInfos = append(interactInfos, interactInfo)
	}
	resp.VideoInteractInfoList = interactInfos
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// BatchGetUserInteractInfo implements the InteractionServerImpl interface.
func (s *InteractionServerImpl) BatchGetUserInteractInfo(ctx context.Context, req *interaction.DouyinBatchGetUserInteractInfoRequest) (resp *interaction.DouyinBatchGetUserInteractInfoResponse, err error) {
	resp = new(interaction.DouyinBatchGetUserInteractInfoResponse)
	favoriteCounts, err := s.FavoriteRedisDao.BatchGetFavoriteCountByUserId(ctx, req.TargetUserIdList)
	if err != nil {
		klog.Errorf("get user's favorite count by redis error: %s", err.Error())
		favoriteCounts, err = s.FavoriteDao.BatchGetFavoriteCountByUserId(ctx, req.TargetUserIdList)
		if err != nil {
			klog.Errorf("get user's favorite count error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("get user's favorite count err"))
			return resp, nil
		}
	}

	// search user's videos and favorite count of video
	userInteractInfoList := make([]*base.UserInteractInfo, 0)
	for id, userId := range req.TargetUserIdList {
		videoResp, err := global.VideoClient.GetPublishVideoIdList(ctx, &video.DouyinGetPublishVideoIdListRequest{
			UserId: userId,
		})
		if err != nil {
			resp.BaseResp = util.BuildBaseResp(errno.RPCVideoErr)
			return resp, nil
		}
		if videoResp.BaseResp.StatusCode != int32(code.Code_Success) {
			resp.BaseResp = videoResp.BaseResp
			return resp, nil
		}
		getFavorite, err := s.FavoriteRedisDao.BatchGetFavoriteCountSumByVideoId(ctx, videoResp.VideoIdList)
		if err != nil {
			klog.Errorf("get favorite count of user by redis error: %s", err.Error())
			getFavorite, err = s.FavoriteDao.BatchGetFavoriteCountSumByVideoId(ctx, videoResp.VideoIdList)
			if err != nil {
				klog.Errorf("get favorite count of user error: %s", err.Error())
				resp.BaseResp = util.BuildBaseResp(errno.InteractionServerErr.WithMessage("get favorite count of user error"))
				return resp, nil
			}
		}

		userInteractInfoList = append(userInteractInfoList, &base.UserInteractInfo{
			TotalFavorited: getFavorite,
			WorkCount:      int64(len(videoResp.VideoIdList)),
			FavoriteCount:  favoriteCounts[id],
		})
	}
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	resp.UserInteractInfoList = userInteractInfoList
	return resp, nil
}
