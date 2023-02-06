package main

import (
	"context"
	core "douyin/kitex_gen/core"
)

// CoreServiceImpl implements the last service interface defined in the IDL.
type CoreServiceImpl struct{}

// Feed implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) Feed(ctx context.Context, req *core.DouyinFeedRequest) (resp *core.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) UserRegister(ctx context.Context, req *core.DouyinUserRegisterRequest) (resp *core.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) UserLogin(ctx context.Context, req *core.DouyinUserLoginRequest) (resp *core.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) UserInfo(ctx context.Context, req *core.DouyinUserRequest) (resp *core.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) PublishAction(ctx context.Context, req *core.DouyinPublishActionRequest) (resp *core.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the CoreServiceImpl interface.
func (s *CoreServiceImpl) PublishList(ctx context.Context, req *core.DouyinPublishListRequest) (resp *core.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}
