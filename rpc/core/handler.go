package main

import (
	"context"
	core "douyin/kitex_gen/core"
	"douyin/pkg/errno"
	"douyin/rpc/core/service"
)

// CoreServiceImpl implements the last service interface defined in the IDL.
type CoreServiceImpl struct{}

// Feed implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) Feed(ctx context.Context, req *core.DouyinFeedRequest) (resp *core.DouyinFeedResponse, err error) {
	resp = new(core.DouyinFeedResponse)
	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	errNo := errno.Success
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.VideoList = videos
	resp.NextTime = nextTime
	return resp, nil
}

// UserRegister implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) UserRegister(ctx context.Context, req *core.DouyinUserRegisterRequest) (resp *core.DouyinUserRegisterResponse, err error) {
	resp = new(core.DouyinUserRegisterResponse)
	errNo := errno.Success
	if len(req.Username) == 0 || len(req.Username) == 0 {
		errNo = errno.ParamErr
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	userID, err := service.NewUserService(ctx).Register(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.UserId = userID
	return resp, nil
}

// UserLogin implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) UserLogin(ctx context.Context, req *core.DouyinUserLoginRequest) (resp *core.DouyinUserLoginResponse, err error) {
	resp = new(core.DouyinUserLoginResponse)
	errNo := errno.Success
	if len(req.Username) == 0 || len(req.Username) == 0 {
		errNo = errno.ParamErr
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	userID, err := service.NewUserService(ctx).Login(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.UserId = userID
	return resp, nil
}

// UserInfo implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) UserInfo(ctx context.Context, req *core.DouyinUserRequest) (resp *core.DouyinUserResponse, err error) {
	resp = new(core.DouyinUserResponse)
	errNo := errno.Success
	user, err := service.NewUserService(ctx).Info(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.User = user
	return resp, nil
}

// PublishAction implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) PublishAction(ctx context.Context, req *core.DouyinPublishActionRequest) (resp *core.DouyinPublishActionResponse, err error) {
	resp = new(core.DouyinPublishActionResponse)
	errNo := errno.Success
	if err = service.NewPublishService(ctx).PublishAction(req); err != nil {
		errNo = errno.ConvertErr(err)
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	return resp, nil
}

// PublishList implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) PublishList(ctx context.Context, req *core.DouyinPublishListRequest) (resp *core.DouyinPublishListResponse, err error) {
	resp = new(core.DouyinPublishListResponse)
	errNo := errno.Success
	videos, err := service.NewPublishService(ctx).PublishList(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.VideoList = videos
	return resp, nil
}
