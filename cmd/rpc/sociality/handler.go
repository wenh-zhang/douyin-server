package main

import (
	"context"
	"douyin/cmd/rpc/sociality/dao"
	"douyin/cmd/rpc/sociality/model"
	"douyin/cmd/rpc/sociality/mq"
	"douyin/cmd/rpc/sociality/redis"
	"douyin/shared/constant"
	"douyin/shared/errno"
	"douyin/shared/kitex_gen/base"
	"douyin/shared/kitex_gen/sociality"
	"douyin/shared/util"
	"github.com/cloudwego/kitex/pkg/klog"
)

// SocialityServiceImpl implements the last service interface defined in the IDL.
type SocialityServiceImpl struct {
	Dao            *dao.Follow
	FollowRedisDao *redis.Follow
	Publisher      *mq.Publisher
}

// Relation implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) Relation(ctx context.Context, req *sociality.DouyinRelationActionRequest) (resp *sociality.DouyinRelationActionResponse, err error) {
	resp = new(sociality.DouyinRelationActionResponse)
	follow := &model.Follow{
		FromUserID: req.LocalUserId,
		ToUserID:   req.TargetUserId,
	}
	if err = s.Publisher.Publish(&mq.FollowWithAction{
		Follow:     follow,
		ActionType: req.ActionType,
	}); err != nil {
		klog.Errorf("follow error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("follow error"))
		return resp, nil
	}
	if req.ActionType == constant.ActionTypeFollow {
		if err = s.FollowRedisDao.CreateFollow(ctx, follow); err != nil {
			klog.Errorf("follow by redis error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("follow error"))
			return resp, nil
		}
	} else if req.ActionType == constant.ActionTypeCancelFollow {
		if err = s.FollowRedisDao.DeleteFollow(ctx, req.LocalUserId, req.TargetUserId); err != nil {
			klog.Errorf("cancel follow by redis error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("cancel follow error"))
			return resp, nil
		}
	} else {
		resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("action type error"))
		return resp, nil
	}
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetRelationUserIdList implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) GetRelationUserIdList(ctx context.Context, req *sociality.DouyinGetRelationUserIdListRequest) (resp *sociality.DouyinGetRelationUserIdListResponse, err error) {
	resp = new(sociality.DouyinGetRelationUserIdListResponse)
	if req.Option == constant.OptionFollow {
		resp.UserIdList, err = s.FollowRedisDao.GetFollowIdList(ctx, req.UserId)
		if err != nil {
			klog.Errorf("get follow user id list by redis error: %s", err.Error())
			resp.UserIdList, err = s.Dao.GetFollowIdList(ctx, req.UserId)
			if err != nil {
				klog.Errorf("get follow user id list error: %s", err.Error())
				resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("get follow user id list error"))
				return resp, nil
			}
		}
	} else if req.Option == constant.OptionFollower {
		resp.UserIdList, err = s.FollowRedisDao.GetFollowerIdList(ctx, req.UserId)
		if err != nil {
			klog.Errorf("get follower user id list by redis error: %s", err.Error())
			resp.UserIdList, err = s.Dao.GetFollowerIdList(ctx, req.UserId)
			if err != nil {
				klog.Errorf("get follower user id list error: %s", err.Error())
				resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("get follower user id list error"))
				return resp, nil
			}
		}
	} else if req.Option == constant.OptionFriend {
		resp.UserIdList, err = s.FollowRedisDao.GetFriendIdList(ctx, req.UserId)
		if err != nil {
			klog.Errorf("get friend user id list by redis error: %s", err.Error())
			resp.UserIdList, err = s.Dao.GetFriendIdList(ctx, req.UserId)
			if err != nil {
				klog.Errorf("get friend user id list error: %s", err.Error())
				resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("get friend user id list error"))
				return resp, nil
			}
		}
	} else {
		resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("option type error"))
		return resp, nil
	}
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// BatchGetSocialInfo implements the SocialityServiceImpl interface.
func (s *SocialityServiceImpl) BatchGetSocialInfo(ctx context.Context, req *sociality.DouyinBatchGetSocialInfoRequest) (resp *sociality.DouyinBatchGetSocialInfoResponse, err error) {
	resp = new(sociality.DouyinBatchGetSocialInfoResponse)
	followCounts, err := s.FollowRedisDao.BatchGetFollowCountByUserId(ctx, req.TargetUserIdList)
	if err != nil {
		klog.Errorf("get follow count by redis error: %s", err.Error())
		followCounts, err = s.Dao.BatchGetFollowCountByUserId(ctx, req.TargetUserIdList)
		if err != nil {
			klog.Errorf("get follow count error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("get follow count error"))
			return resp, nil
		}
	}

	followerCounts, err := s.FollowRedisDao.BatchGetFollowerCountByUserId(ctx, req.TargetUserIdList)
	if err != nil {
		klog.Errorf("get follower count by redis error: %s", err.Error())
		followerCounts, err = s.Dao.BatchGetFollowerCountByUserId(ctx, req.TargetUserIdList)
		if err != nil {
			klog.Errorf("get follower count error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("get follower count error"))
			return resp, nil
		}
	}

	isFollows, err := s.FollowRedisDao.BatchGetFollowInfoByUserId(ctx, req.LocalUserId, req.TargetUserIdList)
	if err != nil {
		klog.Errorf("check if  by redis error: %s", err.Error())
		isFollows, err = s.Dao.BatchGetFollowInfoByUserId(ctx, req.LocalUserId, req.TargetUserIdList)
		if err != nil {
			klog.Errorf("check if follow error: %s", err.Error())
			resp.BaseResp = util.BuildBaseResp(errno.SocialityServerErr.WithMessage("check if follow error"))
			return resp, nil
		}
	}
	socialInfos := make([]*base.SocialInfo, 0)
	for i := range req.TargetUserIdList {
		socialInfo := &base.SocialInfo{
			FollowCount:   followCounts[i],
			FollowerCount: followerCounts[i],
			IsFollow:      isFollows[i],
		}
		socialInfos = append(socialInfos, socialInfo)
	}
	resp.SocialInfoList = socialInfos
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}
