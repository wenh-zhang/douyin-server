package social

import (
	"context"
	social "douyin/kitex_gen/social"
)

// SocialServiceImpl implements the last service interface defined in the IDL.
type SocialServiceImpl struct{}

// RelationAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) RelationAction(ctx context.Context, req *social.DouyinRelationActionRequest) (resp *social.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) RelationFollowList(ctx context.Context, req *social.DouyinRelationFollowListRequest) (resp *social.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) RelationFollowerList(ctx context.Context, req *social.DouyinRelationFollowerListRequest) (resp *social.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFriendList implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) RelationFriendList(ctx context.Context, req *social.DouyinRelationFriendListRequest) (resp *social.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageAction implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) MessageAction(ctx context.Context, req *social.DouyinMessageActionRequest) (resp *social.DouyinMessageActionResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the SocialServiceImpl interface.
func (s *SocialServiceImpl) MessageChat(ctx context.Context, req *social.DouyinMessageChatRequest) (resp *social.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
