package common

import (
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"douyin/kitex_gen/interact"
	"time"
)

func User(user *db.User) *core.User {
	if user == nil{
		return nil
	}
	return &core.User{
		Id:            int64(user.ID),
		Name:          user.UserName,
		FollowCount:   &user.FollowCount,
		FollowerCount: &user.FollowerCount,
		Avatar:        &user.Avatar,
	}
}

func Video(video *db.Video) *core.Video {
	if video == nil {
		return nil
	}
	return &core.Video{
		Id: int64(video.ID),
		Author: &core.User{
			Id: int64(video.UserID),
		},
		PlayUrl:       video.PlayURL,
		CoverUrl:      video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	}
}

func Videos(videos []*db.Video) ([]*core.Video, *int64) {
	var nextTime time.Time
	res := make([]*core.Video, 0, len(videos))
	for _, video := range videos {
		if v := Video(video); v != nil {
			res = append(res, v)
			nextTime = video.CreatedAt
		}
	}
	timeStamp := nextTime.Unix()
	return res, &timeStamp
}

func Comment(comment *db.Comment) *interact.Comment {
	if comment == nil{
		return nil
	}
	return &interact.Comment{
		Id: comment.ID,
		User: &core.User{
			Id: comment.UserID,
		},
		Content: comment.Content,
		CreateDate: comment.CreatedAt.String(),
	}
}

func Comments(comments []*db.Comment) []*interact.Comment {
	res := make([]*interact.Comment, 0, len(comments))
	for _, comment := range comments {
		if v := Comment(comment); v != nil {
			res = append(res, v)
		}
	}
	return res
}