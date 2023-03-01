package global

import (
	"douyin/cmd/rpc/user/config"
	interaction "douyin/shared/kitex_gen/interaction/interactionserver"
	message "douyin/shared/kitex_gen/message/messageservice"
	sociality "douyin/shared/kitex_gen/sociality/socialityservice"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	DB                *gorm.DB
	InteractionClient interaction.Client
	SocialityClient   sociality.Client
	MessageClient     message.Client
	AmqpConn          *amqp.Connection

	EtcdConfig  *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RPCConfig   *config.RPCConfig
	AmqpConfig  *config.AmqpConfig
)
