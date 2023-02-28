package redis

import (
	"context"
	"douyin/cmd/rpc/interaction/model"
	"github.com/go-redis/redis/v8"
	"strconv"
)

// Comment redis缓存在此处仅作计数使用
type Comment struct {
	client *redis.Client
}

func NewComment(client *redis.Client) *Comment {
	return &Comment{
		client: client,
	}
}

func (s *Comment) CreateComment(ctx context.Context, comment *model.Comment) error {
	if err := s.client.SAdd(ctx, strconv.FormatInt(comment.VideoId, 10), comment.ID).
		Err(); err != nil {
		return err
	}
	return nil
}

func (s *Comment) DeleteComment(ctx context.Context, videoId, commentId int64) error {
	if err := s.client.SRem(ctx, strconv.FormatInt(videoId, 10), commentId).Err(); err != nil {
		return err
	}
	return nil
}

// BatchGetCommentCountByVideoId 通过视频id列表批量获取视频评论数列表
func (s *Comment) BatchGetCommentCountByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	res := make([]int64, 0)
	for _, videoId := range videoIds {
		size, err := s.client.SCard(ctx, strconv.FormatInt(videoId, 10)).Result()
		if err != nil {
			return nil, err
		}
		res = append(res, size)
	}
	return res, nil
}
