package service

import (
	"context"
	"douyin/dal"
	"douyin/dal/db"
	"douyin/kitex_gen/interact"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/rpc/common"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx}
}

func (s *CommentService) CommentAction(req *interact.DouyinCommentActionRequest) (*interact.Comment, error) {
	conn := dal.GetConn()
	videos, err := db.MGetVideos(s.ctx, conn, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errno.ServiceErr
	}
	video := videos[0]
	commentCount := video.CommentCount

	tx := conn.Begin()
	var convertComment *interact.Comment
	if req.ActionType == constant.ActionTypeComment {
		if req.CommentText == nil {
			tx.Rollback()
			return nil, errno.ParamErr
		}
		comment := &db.Comment{UserID: req.UserId, VideoID: req.VideoId, Content: *req.CommentText}
		if err = db.CreateComment(s.ctx, tx, []*db.Comment{comment}); err != nil {
			tx.Rollback()
			return nil, err
		}
		convertComment = common.Comment(comment)
		if err = common.CheckAuthorOfComment(s.ctx, tx, []*interact.Comment{convertComment}, &req.UserId); err != nil {
			tx.Rollback()
			return nil, err
		}
		commentCount++
	} else if req.ActionType == constant.ActionTypeDeleteComment {
		if req.CommentId == nil {
			tx.Rollback()
			return nil, errno.ParamErr
		}
		if err = db.DeleteComment(s.ctx, tx, []int64{*req.CommentId}); err != nil {
			tx.Rollback()
			return nil, err
		}
		if commentCount > 0 {
			commentCount--
		}
	} else {
		tx.Rollback()
		return nil, errno.ParamErr
	}
	if commentCount != video.CommentCount{
		if err = db.UpdateVideo(s.ctx, tx, req.VideoId, nil, nil, nil, video.FavoriteCount, commentCount); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return convertComment, nil
}

func (s *CommentService) CommentList(req *interact.DouyinCommentListRequest) ([]*interact.Comment, error) {
	conn := dal.GetConn()
	comments, err := db.MGetCommentsByVideoID(s.ctx, conn, req.VideoId)
	if err != nil {
		return nil, err
	}
	convertComments := common.Comments(comments)

	//check author info of each comment
	if err = common.CheckAuthorOfComment(s.ctx, conn, convertComments, &req.UserId); err != nil {
		return nil, err
	}
	return convertComments, nil
}
