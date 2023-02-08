package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"time"
)

func User(user *db.User) *core.User {
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
			Id:   int64(video.UserID),
			Name: "test_user",
		},
		PlayUrl:       video.PlayURL,
		CoverUrl:      video.CoverURL,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
	}
}

func Videos(videos []*db.Video) ([]*core.Video, *int64) {
	var nextTime time.Time
	res := make([]*core.Video, 0)
	for _, video := range videos {
		if v := Video(video); v != nil {
			res = append(res, v)
			nextTime = video.CreatedAt
		}
	}
	if len(res) == 0 {
		return nil, nil
	}
	timeStamp := nextTime.Unix()
	return res, &timeStamp
}
