package common

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/core"
	"douyin/kitex_gen/interact"
	"gorm.io/gorm"
)

// CheckIfLikeVideo check if user liked the videos
func CheckIfLikeVideo(ctx context.Context, tx *gorm.DB, videos []*core.Video, userID *int64) error {
	if userID == nil {
		return nil
	}
	for _, video := range videos {
		res, err := db.CheckIfFavorite(ctx, tx, video.Id, *userID)
		if err != nil {
			return err
		}
		video.IsFavorite = res
	}
	return nil
}

// CheckAuthorFromUserID query user info and return map with information of users
func CheckAuthorFromUserID(ctx context.Context, tx *gorm.DB, queryUserIDs []int64, userID *int64) (map[int64]*core.User, error) {
	//query all user info at once
	if len(queryUserIDs) == 0 {
		return nil, nil
	}
	users, err := db.MGetUsers(ctx, tx, queryUserIDs)
	if err != nil {
		return nil, err
	}
	//save user info in a map
	userMap := make(map[int64]*core.User, len(users))
	for _, user := range users {
		author := User(user)
		if userID != nil { // userID is acquired from token
			follow, err := db.CheckIfFollow(ctx, tx, *userID, author.Id)
			if err != nil {
				return nil, err
			}
			author.IsFollow = follow
		}
		userMap[int64(user.ID)] = author
	}
	return userMap, nil
}

// CheckAuthorOfVideo add author info into videos
func CheckAuthorOfVideo(ctx context.Context, tx *gorm.DB, videos []*core.Video, userID *int64) error {
	userIDs := make([]int64, 0, len(videos))
	for _, video := range videos {
		userIDs = append(userIDs, video.Author.Id)
	}

	userMap, err := CheckAuthorFromUserID(ctx, tx, userIDs, userID)
	if err != nil {
		return err
	}

	for _, video := range videos {
		video.Author = userMap[video.Author.Id]
	}
	return nil
}

// CheckAuthorOfComment add author info into comments
func CheckAuthorOfComment(ctx context.Context, tx *gorm.DB, comments []*interact.Comment, userID *int64) error {
	userIDs := make([]int64, 0, len(comments))
	for _, comment := range comments {
		userIDs = append(userIDs, comment.User.Id)
	}

	userMap, err := CheckAuthorFromUserID(ctx, tx, userIDs, userID)
	if err != nil {
		return err
	}

	for _, comment := range comments {
		comment.User = userMap[comment.User.Id]
	}
	return nil
}
