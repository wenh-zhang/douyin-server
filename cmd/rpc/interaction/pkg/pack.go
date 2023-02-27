package pkg

import (
	"douyin/cmd/rpc/interaction/model"
	"douyin/shared/kitex_gen/base"
	"time"
)

func PackComment(comment *model.Comment, user *base.User)*base.Comment{
	return &base.Comment{
		Id:         comment.ID,
		User:       user,
		Content:    comment.Content,
		CreateDate: time.Unix(comment.CreatedAt, 0).Format("2006.01.02 15:04:05"),
	}
}

func BatchPackComment(commentList []*model.Comment, userList []*base.User) []*base.Comment {
	comments := make([]*base.Comment, 0)
	for id, comment := range commentList {
		comments = append(comments, PackComment(comment, userList[id]))
	}
	return comments
}
