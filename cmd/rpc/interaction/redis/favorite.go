package redis

import (
	"context"
	"douyin/cmd/rpc/interaction/model"
	"douyin/shared/errno"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Favorite struct {
	client *redis.Client
}

func NewFavorite(client *redis.Client) *Favorite {
	return &Favorite{
		client: client,
	}
}

func (s *Favorite) Like(ctx context.Context, favorite *model.Favorite) error {
	tp := s.client.TxPipeline()
	if err := tp.ZAdd(ctx, "user"+strconv.FormatInt(favorite.UserId, 10), &redis.Z{
		Score:  float64(favorite.CreatedAt),
		Member: favorite.VideoId,
	}).Err(); err != nil {
		return err
	}
	if err := tp.ZAdd(ctx, "video"+strconv.FormatInt(favorite.VideoId, 10), &redis.Z{
		Score:  float64(favorite.CreatedAt),
		Member: favorite.UserId,
	}).Err(); err != nil {
		return err
	}
	if _, err := tp.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Favorite) Unlike(ctx context.Context, userId, videoId int64) error {
	tp := s.client.TxPipeline()
	if err := tp.ZRem(ctx, "user"+strconv.FormatInt(userId, 10), videoId).Err(); err != nil {
		return err
	}
	if err := tp.ZRem(ctx, "video"+strconv.FormatInt(videoId, 10), userId).Err(); err != nil {
		return err
	}
	if _, err := tp.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Favorite) GetFavoriteInfoByVideoId(ctx context.Context, userId, videoId int64) (bool, error) {
	_, err1 := s.client.ZScore(ctx, "user"+strconv.FormatInt(userId, 10), strconv.FormatInt(videoId, 10)).Result()
	if err1 != nil && err1 != redis.Nil {
		return false, err1
	}
	_, err2 := s.client.ZScore(ctx, "video"+strconv.FormatInt(videoId, 10), strconv.FormatInt(userId, 10)).Result()
	if err2 != nil && err2 != redis.Nil {
		return false, err2
	}
	if err1 != err2 {
		return false, errno.InteractionServerErr.WithMessage("dirty data in redis")
	}
	return err1 == nil, nil
}

// BatchGetFavoriteInfoByVideoId 获取用户对视频列表的点赞信息
func (s *Favorite) BatchGetFavoriteInfoByVideoId(ctx context.Context, userId int64, videoIds []int64) ([]bool, error) {
	res := make([]bool, 0)
	for _, videoId := range videoIds {
		is, err := s.GetFavoriteInfoByVideoId(ctx, userId, videoId)
		if err != nil {
			return nil, err
		}
		res = append(res, is)
	}
	return res, nil
}

// BatchGetFavoriteCountByVideoId 批量获取视频的点赞数
func (s *Favorite) BatchGetFavoriteCountByVideoId(ctx context.Context, videoIds []int64) ([]int64, error) {
	res := make([]int64, 0)
	for _, videoId := range videoIds {
		size, err := s.client.ZCard(ctx, "video"+strconv.FormatInt(videoId, 10)).Result()
		if err != nil {
			return nil, err
		}
		res = append(res, size)
	}
	return res, nil
}

// BatchGetFavoriteCountSumByVideoId 批量获取视频的总点赞数
func (s *Favorite) BatchGetFavoriteCountSumByVideoId(ctx context.Context, videoIds []int64) (int64, error) {
	counts, err := s.BatchGetFavoriteCountByVideoId(ctx, videoIds)
	if err != nil {
		return 0, err
	}
	var sum int64 = 0
	for _, count := range counts {
		sum += count
	}
	return sum, nil
}

// GetFavoriteVideoIdListByUserId 根据用户id获取其点赞的视频id列表
func (s *Favorite) GetFavoriteVideoIdListByUserId(ctx context.Context, userId int64) ([]int64, error) {
	res := make([]int64, 0)
	err := s.client.ZRevRange(ctx, "user"+strconv.FormatInt(userId, 10), 0, -1).ScanSlice(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// BatchGetFavoriteCountByUserId 根据用户id列表批量获取用户点赞数
func (s *Favorite) BatchGetFavoriteCountByUserId(ctx context.Context, userIds []int64) ([]int64, error) {
	res := make([]int64, 0)
	for _, userId := range userIds {
		size, err := s.client.ZCard(ctx, "user"+strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			return nil, err
		}
		res = append(res, size)
	}
	return res, nil
}
