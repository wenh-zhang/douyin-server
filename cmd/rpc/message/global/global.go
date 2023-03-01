package global

import (
	"douyin/cmd/rpc/message/config"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	AmqpConn *amqp.Connection

	EtcdConfig  *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RPCConfig   *config.RPCConfig
	AmqpConfig  *config.AmqpConfig
)
