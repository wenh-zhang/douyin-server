package main

import (
	"context"
	interact "douyin/kitex_gen/interact"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// FavoriteAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *interact.DouyinFavoriteActionRequest) (resp *interact.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *interact.DouyinFavoriteListRequest) (resp *interact.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.DouyinCommentActionRequest) (resp *interact.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.DouyinCommentListRequest) (resp *interact.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
