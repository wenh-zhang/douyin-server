package pkg

import (
	hbase "douyin/cmd/api/biz/model/base"
	kbase "douyin/shared/kitex_gen/base"
)

func User(u *kbase.User) *hbase.User {
	if u == nil {
		return nil
	}
	return &hbase.User{
		ID:              u.Id,
		Name:            u.Name,
		FollowCount:     u.FollowCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        u.IsFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
}

func FriendUser(fu *kbase.FriendUser) *hbase.FriendUser {
	if fu == nil {
		return nil
	}
	return &hbase.FriendUser{
		ID:              fu.Id,
		Name:            fu.Name,
		FollowCount:     fu.FollowCount,
		FollowerCount:   fu.FollowerCount,
		IsFollow:        fu.IsFollow,
		BackgroundImage: fu.BackgroundImage,
		Avatar:          fu.Avatar,
		Signature:       fu.Signature,
		TotalFavorited:  fu.TotalFavorited,
		WorkCount:       fu.WorkCount,
		FavoriteCount:   fu.FavoriteCount,
		Message:         fu.Message,
		MsgType:         fu.MsgType,
	}
}

func Users(users []*kbase.User) []*hbase.User {
	us := make([]*hbase.User, 0)
	for _, ku := range users {
		if hu := User(ku); hu != nil {
			us = append(us, hu)
		}
	}
	return us
}

func FriendUsers(kfus []*kbase.FriendUser) []*hbase.FriendUser {
	hfus := make([]*hbase.FriendUser, 0)
	for _, kfu := range kfus {
		if hfu := FriendUser(kfu); hfu != nil {
			hfus = append(hfus, hfu)
		}
	}
	return hfus
}

func Comment(c *kbase.Comment) *hbase.Comment {
	if c == nil {
		return nil
	}
	return &hbase.Comment{
		ID:         c.Id,
		User:       User(c.User),
		Content:    c.Content,
		CreateDate: c.CreateDate,
	}
}

func Video(v *kbase.Video) *hbase.Video {
	if v == nil {
		return nil
	}
	return &hbase.Video{
		ID:            v.Id,
		Author:        User(v.Author),
		PlayURL:       v.PlayUrl,
		CoverURL:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
		Title:         v.Title,
	}
}

func Videos(videos []*kbase.Video) []*hbase.Video {
	vs := make([]*hbase.Video, 0)
	for _, video := range videos {
		if v := Video(video); v != nil {
			vs = append(vs, v)
		}
	}
	return vs
}

func Comments(comments []*kbase.Comment) []*hbase.Comment {
	cs := make([]*hbase.Comment, 0)
	for _, comment := range comments {
		if c := Comment(comment); c != nil {
			cs = append(cs, c)
		}
	}
	return cs
}

func Message(m *kbase.Message) *hbase.Message {
	if m == nil {
		return nil
	}
	return &hbase.Message{
		ID:         m.Id,
		ToUserID:   m.ToUserId,
		FromUserID: m.Id,
		Content:    m.Content,
		CreateTime: m.CreateTime,
	}
}

func Messages(kms []*kbase.Message) []*hbase.Message {
	hms := make([]*hbase.Message, 0)
	for _, km := range kms {
		if hm := Message(km); hm != nil {
			hms = append(hms, hm)
		}
	}
	return hms
}
