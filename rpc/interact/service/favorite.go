package service

import (
	"context"
	"douyin/dal"
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"douyin/kitex_gen/interact"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/rpc/common"
)

type FavoriteService struct {
	ctx context.Context
}

func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{ctx}
}

// FavoriteAction insert data into table favorite and update video info at the same time
// check the table favorite first in case of the same record be inserted
func (s *FavoriteService) FavoriteAction(req *interact.DouyinFavoriteActionRequest) error {
	conn := dal.GetConn()
	videos, err := db.MGetVideos(s.ctx, conn, []int64{req.VideoId})
	if err != nil {
		return err
	}
	if len(videos) == 0 {
		return errno.ServiceErr
	}
	video := videos[0]
	favoriteCount := video.FavoriteCount

	tx := conn.Begin()
	if req.ActionType == constant.ActionTypeFavorite {
		isFavorite, err := db.CheckIfFavorite(s.ctx, tx, req.VideoId, req.UserId)
		if err != nil {
			tx.Rollback()
			return err
		}
		if isFavorite {
			tx.Rollback()
			return nil
		}
		if err := db.CreateFavorite(s.ctx, tx, []*db.Favorite{&db.Favorite{
			UserID:  req.UserId,
			VideoID: req.VideoId,
		}}); err != nil {
			tx.Rollback()
			return err
		}
		favoriteCount++
	} else if req.ActionType == constant.ActionTypeCancelFavorite {
		if err := db.DeleteFavorite(s.ctx, tx, req.VideoId, req.UserId); err != nil {
			tx.Rollback()
			return err
		}
		if favoriteCount > 0 {
			favoriteCount--
		}
	} else {
		tx.Rollback()
		return errno.ParamErr
	}
	if favoriteCount != video.FavoriteCount {
		if err = db.UpdateVideo(s.ctx, tx, req.VideoId, nil, nil, nil, favoriteCount, video.CommentCount); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (s *FavoriteService) FavoriteList(req *interact.DouyinFavoriteListRequest) ([]*core.Video, error) {
	conn := dal.GetConn()
	videoIDs, err := db.MGetFavoriteVideoIDsByUserID(s.ctx, conn, req.QueryUserId)
	if err != nil {
		return nil, err
	}
	videos, err := db.MGetVideos(s.ctx, conn, videoIDs)
	if err != nil {
		return nil, err
	}
	convertVideos, _ := common.Videos(videos)
	// check if user like the videos
	if err = common.CheckIfLikeVideo(s.ctx, conn, convertVideos, &req.UserId); err != nil {
		return nil, err
	}
	//check author info of each video
	if err = common.CheckAuthorOfVideo(s.ctx, conn, convertVideos, &req.UserId); err != nil {
		return nil, err
	}
	return convertVideos, nil
}
