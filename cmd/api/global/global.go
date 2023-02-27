package global

import (
	"douyin/cmd/api/config"
	interaction "douyin/shared/kitex_gen/interaction/interactionserver"
	message "douyin/shared/kitex_gen/message/messageservice"
	sociality "douyin/shared/kitex_gen/sociality/socialityservice"
	user "douyin/shared/kitex_gen/user/userservice"
	video "douyin/shared/kitex_gen/video/videoservice"
)

var(
	UserClient user.Client
	VideoClient video.Client
	MessageClient message.Client
	InteractionClient interaction.Client
	SocialityClient sociality.Client

	EtcdConfig *config.EtcdConfig
)
