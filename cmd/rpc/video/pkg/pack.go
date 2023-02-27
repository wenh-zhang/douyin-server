package pkg

import (
	"douyin/cmd/rpc/video/model"
	"douyin/shared/kitex_gen/base"
)

func BatchPackVideo(videoList []*model.Video, videoInteractInfoList []*base.VideoInteractInfo, userList []*base.User) []*base.Video {
	videos := make([]*base.Video, 0)
	for id, video := range videoList {
		videos = append(videos, &base.Video{
			Id:            video.ID,
			Author:        userList[id],
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			Title:         video.Title,
			FavoriteCount: videoInteractInfoList[id].FavoriteCount,
			CommentCount:  videoInteractInfoList[id].CommentCount,
			IsFavorite:    videoInteractInfoList[id].IsFavorite,
		})
	}
	return videos
}
