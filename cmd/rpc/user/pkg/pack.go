package pkg

import (
	"douyin/cmd/rpc/user/model"
	"douyin/shared/kitex_gen/base"
)

func BatchPackUser(userList []*model.User, userInteractInfoList []*base.UserInteractInfo, socialInfoList []*base.SocialInfo) []*base.User {
	users := make([]*base.User, 0)
	for id, user := range userList {
		users = append(users, &base.User{
			Id:              user.ID,
			Name:            user.Name,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackGroundImage,
			Signature:       user.Signature,
			FollowCount:     socialInfoList[id].FollowCount,
			FollowerCount:   socialInfoList[id].FollowerCount,
			IsFollow:        socialInfoList[id].IsFollow,
			TotalFavorited:  userInteractInfoList[id].TotalFavorited,
			WorkCount:       userInteractInfoList[id].WorkCount,
			FavoriteCount:   userInteractInfoList[id].FavoriteCount,
		})
	}
	return users
}

func BatchPackFriendUser(userList []*base.User, latestMsgList []*base.LatestMsg) []*base.FriendUser {
	friends := make([]*base.FriendUser, 0)
	for id, user := range userList {
		friends = append(friends, &base.FriendUser{
			Id:              user.Id,
			Name:            user.Name,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
			Message:         latestMsgList[id].Message,
			MsgType:         latestMsgList[id].MsgType,
		})
	}
	return friends
}
