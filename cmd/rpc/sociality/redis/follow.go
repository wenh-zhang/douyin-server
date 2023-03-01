package redis

import (
	"context"
	"douyin/cmd/rpc/sociality/model"
	"douyin/shared/errno"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Follow struct {
	client *redis.Client
}

func NewFollow(client *redis.Client) *Follow {
	return &Follow{
		client: client,
	}
}

func (s *Follow) CreateFollow(ctx context.Context, follow *model.Follow) error {
	tp := s.client.TxPipeline()
	if err := tp.SAdd(ctx, "follow"+strconv.FormatInt(follow.FromUserID, 10), follow.ToUserID).Err(); err != nil {
		return err
	}
	if err := tp.SAdd(ctx, "follower"+strconv.FormatInt(follow.ToUserID, 10), follow.FromUserID).Err(); err != nil {
		return err
	}
	if _, err := tp.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Follow) DeleteFollow(ctx context.Context, fromUserId, toUserId int64) error {
	tp := s.client.TxPipeline()
	if err := tp.SRem(ctx, "follow"+strconv.FormatInt(fromUserId, 10), toUserId).Err(); err != nil {
		return err
	}
	if err := tp.SRem(ctx, "follower"+strconv.FormatInt(toUserId, 10), fromUserId).Err(); err != nil {
		return err
	}
	if _, err := tp.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// GetFollowIdList 获取关注用户id列表
func (s *Follow) GetFollowIdList(ctx context.Context, userId int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := s.client.SMembers(ctx, "follow"+strconv.FormatInt(userId, 10)).ScanSlice(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetFollowerIdList 获取粉丝用户id列表
func (s *Follow) GetFollowerIdList(ctx context.Context, userId int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := s.client.SMembers(ctx, "follower"+strconv.FormatInt(userId, 10)).ScanSlice(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetFriendIdList 获取朋友用户id列表
func (s *Follow) GetFriendIdList(ctx context.Context, userId int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := s.client.SInter(ctx, "follow"+strconv.FormatInt(userId, 10), "follower"+strconv.FormatInt(userId, 10)).
		ScanSlice(&res); err != nil {
		return nil, err
	}
	return res, nil
}

// BatchGetFollowCountByUserId 通过用户id列表批量获取关注数
func (s *Follow) BatchGetFollowCountByUserId(ctx context.Context, userIds []int64) ([]int64, error) {
	res := make([]int64, 0)
	for _, userId := range userIds {
		size, err := s.client.SCard(ctx, "follow"+strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			return nil, err
		}
		res = append(res, size)
	}
	return res, nil
}

// BatchGetFollowerCountByUserId 通过用户id列表批量获取粉丝数
func (s *Follow) BatchGetFollowerCountByUserId(ctx context.Context, userIds []int64) ([]int64, error) {
	res := make([]int64, 0)
	for _, userId := range userIds {
		size, err := s.client.SCard(ctx, "follower"+strconv.FormatInt(userId, 10)).Result()
		if err != nil {
			return nil, err
		}
		res = append(res, size)
	}
	return res, nil
}

func (s *Follow) GetFollowInfoByUserId(ctx context.Context, localUserId int64, targetUserId int64) (bool, error) {
	is1, err1 := s.client.SIsMember(ctx, "follow"+strconv.FormatInt(localUserId, 10), targetUserId).Result()
	if err1 != nil {
		return false, err1
	}
	is2, err2 := s.client.SIsMember(ctx, "follow"+strconv.FormatInt(targetUserId, 10), localUserId).Result()
	if err2 != nil {
		return false, err2
	}
	if is1 != is2 {
		return false, errno.SocialityServerErr.WithMessage("dirty data in redis")
	}
	return is1, nil
}

// BatchGetFollowInfoByUserId 通过用户id列表批量获取关注信息
func (s *Follow) BatchGetFollowInfoByUserId(ctx context.Context, localUserId int64, targetUserIds []int64) ([]bool, error) {
	res := make([]bool, 0)
	for _, targetUserId := range targetUserIds {
		is, err := s.GetFollowInfoByUserId(ctx, localUserId, targetUserId)
		if err != nil {
			return nil, err
		}
		res = append(res, is)
	}
	return res, nil
}
