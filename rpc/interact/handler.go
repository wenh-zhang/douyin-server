package main

import (
	"context"
	interact "douyin/kitex_gen/interact"
	"douyin/pkg/errno"
	"douyin/rpc/interact/service"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// FavoriteAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteAction(ctx context.Context, req *interact.DouyinFavoriteActionRequest) (resp *interact.DouyinFavoriteActionResponse, err error) {
	resp = new(interact.DouyinFavoriteActionResponse)
	errNo := errno.Success
	err = service.NewFavoriteService(ctx).FavoriteAction(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	return resp, nil
}

// FavoriteList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) FavoriteList(ctx context.Context, req *interact.DouyinFavoriteListRequest) (resp *interact.DouyinFavoriteListResponse, err error) {
	resp = new(interact.DouyinFavoriteListResponse)
	errNo := errno.Success
	videos, err := service.NewFavoriteService(ctx).FavoriteList(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.VideoList = videos
	return resp, nil
}

// CommentAction implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentAction(ctx context.Context, req *interact.DouyinCommentActionRequest) (resp *interact.DouyinCommentActionResponse, err error) {
	resp = new(interact.DouyinCommentActionResponse)
	errNo := errno.Success
	comment, err := service.NewCommentService(ctx).CommentAction(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.Comment = comment
	return resp, nil
}

// CommentList implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CommentList(ctx context.Context, req *interact.DouyinCommentListRequest) (resp *interact.DouyinCommentListResponse, err error) {
	resp = new(interact.DouyinCommentListResponse)
	errNo := errno.Success
	comments, err := service.NewCommentService(ctx).CommentList(req)
	if err != nil {
		errNo = errno.ConvertErr(err)
		resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
		return resp, nil
	}
	resp.StatusCode, resp.StatusMsg = errNo.ErrCode, &errNo.ErrMsg
	resp.CommentList = comments
	return resp, nil
}
