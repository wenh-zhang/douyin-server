package main

import (
	"context"
	"douyin/cmd/rpc/user/dao"
	"douyin/cmd/rpc/user/global"
	"douyin/cmd/rpc/user/model"
	"douyin/cmd/rpc/user/mq"
	"douyin/cmd/rpc/user/pkg"
	"douyin/shared/constant"
	"douyin/shared/errno"
	code "douyin/shared/kitex_gen/errno"
	"douyin/shared/kitex_gen/interaction"
	"douyin/shared/kitex_gen/message"
	"douyin/shared/kitex_gen/sociality"
	user "douyin/shared/kitex_gen/user"
	"douyin/shared/util"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	Jwt       *util.JWT
	Dao       *dao.User
	Publisher *mq.Publisher
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp = new(user.DouyinUserRegisterResponse)
	sf, err := snowflake.NewSnowflake(constant.SnowFlakeDataCenterId, constant.UserSnowFlakeWorkerId)
	if err != nil {
		klog.Errorf("snow flake generate failed: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("generate user id failed"))
		return resp, nil
	}
	u := &model.User{
		ID:       sf.NextVal(),
		Name:     req.Username,
		Password: pkg.MD5(req.Password),
		// TODO: Set the avatar, backgroundImage and signature according to user needs
		Avatar:          "https://w.wallhaven.cc/full/3l/wallhaven-3lq8k3.png",
		BackGroundImage: "https://w.wallhaven.cc/full/9d/wallhaven-9ddqjw.png",
		Signature:       "hello world",
	}
	_, err = s.Dao.GetUserByName(ctx, u.Name)
	if err == nil {
		resp.BaseResp = util.BuildBaseResp(errno.UserAlreadyExistErr)
		return resp, nil
	} else if err != gorm.ErrRecordNotFound {
		klog.Errorf("create user error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("create user error"))
		return resp, nil
	}
	if err = s.Publisher.Publish(u); err != nil {
		klog.Errorf("create user error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("create user error"))
		return resp, nil
	}

	token, err := s.Jwt.GenerateToken(u.ID)
	if err != nil {
		klog.Errorf("generate token error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("generate token error"))
		return resp, nil
	}
	resp.Token = token
	resp.UserId = u.ID
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp = new(user.DouyinUserLoginResponse)
	u, err := s.Dao.GetUserByName(ctx, req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.BaseResp = util.BuildBaseResp(errno.UserNotFoundErr)
			return resp, nil
		}
		klog.Errorf("login error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("login error"))
		return resp, nil
	}

	if pkg.MD5(req.Password) != u.Password {
		resp.BaseResp = util.BuildBaseResp(errno.AuthorizeFailErr.WithMessage("wrong password"))
		return resp, nil
	}

	token, err := s.Jwt.GenerateToken(u.ID)
	if err != nil {
		klog.Errorf("generate token error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("generate token error"))
		return resp, nil
	}
	resp.Token = token
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	resp.UserId = u.ID
	return
}

// BatchGetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) BatchGetUserInfo(ctx context.Context, req *user.DouyinBatchGetUserRequest) (resp *user.DouyinBatchGetUserResponse, err error) {
	resp = new(user.DouyinBatchGetUserResponse)
	userList, err := s.Dao.BatchGetUserById(ctx, req.TargetUserIdList)
	if err != nil {
		klog.Errorf("batch get user info error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("batch get user info error"))
		return resp, nil
	}
	interactionResp, err := global.InteractionClient.BatchGetUserInteractInfo(ctx, &interaction.DouyinBatchGetUserInteractInfoRequest{
		TargetUserIdList: req.TargetUserIdList,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCInteractionErr)
		return resp, nil
	}
	if interactionResp.UserInteractInfoList == nil {
		resp.BaseResp = interactionResp.BaseResp
		return resp, nil
	}
	socialityResp, err := global.SocialityClient.BatchGetSocialInfo(ctx, &sociality.DouyinBatchGetSocialInfoRequest{
		LocalUserId:      req.LocalUserId,
		TargetUserIdList: req.TargetUserIdList,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCSocialityErr)
		return resp, nil
	}
	if socialityResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = socialityResp.BaseResp
		return resp, nil
	}
	resp.UserList = pkg.BatchPackUser(userList, interactionResp.UserInteractInfoList, socialityResp.SocialInfoList)
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetRelationFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetRelationFollowList(ctx context.Context, req *user.DouyinGetRelationFollowListRequest) (resp *user.DouyinGetRelationFollowListResponse, err error) {
	resp = new(user.DouyinGetRelationFollowListResponse)
	followIdResp, err := global.SocialityClient.GetRelationUserIdList(ctx, &sociality.DouyinGetRelationUserIdListRequest{
		UserId: req.TargetUserId,
		Option: constant.OptionFollow,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCSocialityErr)
		return resp, nil
	}
	if followIdResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = followIdResp.BaseResp
		return resp, nil
	}
	userResp, err := s.BatchGetUserInfo(ctx, &user.DouyinBatchGetUserRequest{
		LocalUserId:      req.LocalUserId,
		TargetUserIdList: followIdResp.UserIdList,
	})
	if err != nil {
		klog.Errorf("batch get user info error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("batch get user info error"))
		return resp, nil
	}
	if userResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}
	resp.UserList = userResp.UserList
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetRelationFollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetRelationFollowerList(ctx context.Context, req *user.DouyinGetRelationFollowerListRequest) (resp *user.DouyinGetRelationFollowerListResponse, err error) {
	resp = new(user.DouyinGetRelationFollowerListResponse)
	followerIdResp, err := global.SocialityClient.GetRelationUserIdList(ctx, &sociality.DouyinGetRelationUserIdListRequest{
		UserId: req.TargetUserId,
		Option: constant.OptionFollower,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCSocialityErr)
		return resp, nil
	}
	if followerIdResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = followerIdResp.BaseResp
		return resp, nil
	}
	userResp, err := s.BatchGetUserInfo(ctx, &user.DouyinBatchGetUserRequest{
		LocalUserId:      req.LocalUserId,
		TargetUserIdList: followerIdResp.UserIdList,
	})
	if err != nil {
		klog.Errorf("batch get user info error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("batch get user info error"))
		return resp, nil
	}
	if userResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}
	resp.UserList = userResp.UserList
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetRelationFriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetRelationFriendList(ctx context.Context, req *user.DouyinGetRelationFriendListRequest) (resp *user.DouyinGetRelationFriendListResponse, err error) {
	resp = new(user.DouyinGetRelationFriendListResponse)
	// 获取好友id列表
	friendIdResp, err := global.SocialityClient.GetRelationUserIdList(ctx, &sociality.DouyinGetRelationUserIdListRequest{
		UserId: req.TargetUserId,
		Option: constant.OptionFriend,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCSocialityErr)
		return resp, nil
	}
	if friendIdResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = friendIdResp.BaseResp
		return resp, nil
	}
	// 获取好友信息
	userResp, err := s.BatchGetUserInfo(ctx, &user.DouyinBatchGetUserRequest{
		LocalUserId:      req.LocalUserId,
		TargetUserIdList: friendIdResp.UserIdList,
	})
	if err != nil {
		klog.Errorf("batch get user info error: %s", err.Error())
		resp.BaseResp = util.BuildBaseResp(errno.UserServerErr.WithMessage("batch get user info error"))
		return resp, nil
	}
	if userResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = userResp.BaseResp
		return resp, nil
	}
	// 获取最新消息
	msgResp, err := global.MessageClient.BatchGetLatestMessage(ctx, &message.DouyinBatchGetLatestMessageRequest{
		LocalUserId:      req.LocalUserId,
		TargetUserIdList: friendIdResp.UserIdList,
	})
	if err != nil {
		resp.BaseResp = util.BuildBaseResp(errno.RPCMessageErr)
		return resp, nil
	}
	if msgResp.BaseResp.StatusCode != int32(code.Code_Success) {
		resp.BaseResp = msgResp.BaseResp
		return resp, nil
	}
	resp.UserList = pkg.BatchPackFriendUser(userResp.UserList, msgResp.LatestMsgList)
	resp.BaseResp = util.BuildBaseResp(errno.Success)
	return resp, nil
}
