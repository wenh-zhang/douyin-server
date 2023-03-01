package global

import (
	"douyin/cmd/api/config"
	"douyin/cmd/api/service"
	interaction "douyin/shared/kitex_gen/interaction/interactionserver"
	message "douyin/shared/kitex_gen/message/messageservice"
	sociality "douyin/shared/kitex_gen/sociality/socialityservice"
	user "douyin/shared/kitex_gen/user/userservice"
	video "douyin/shared/kitex_gen/video/videoservice"
	"github.com/minio/minio-go/v7"
	"github.com/streadway/amqp"
)

var (
	UserClient        user.Client
	VideoClient       video.Client
	MessageClient     message.Client
	InteractionClient interaction.Client
	SocialityClient   sociality.Client

	MinioClient   *minio.Client
	AmqpConn      *amqp.Connection
	UploadService *service.Upload

	EtcdConfig  *config.EtcdConfig
	MinioConfig *config.MinioConfig
	AmqpConfig  *config.AmqpConfig
)
