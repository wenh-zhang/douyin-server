package global

import (
	"douyin/cmd/rpc/video/config"
	interaction "douyin/shared/kitex_gen/interaction/interactionserver"
	user "douyin/shared/kitex_gen/user/userservice"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	DB                *gorm.DB
	InteractionClient interaction.Client
	UserClient        user.Client
	AmqpConn          *amqp.Connection

	EtcdConfig  *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RPCConfig   *config.RPCConfig
	AmqpConfig  *config.AmqpConfig
)
