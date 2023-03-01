package initialize

import (
	"douyin/cmd/rpc/interaction/global"
	"douyin/shared/constant"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/streadway/amqp"
)

func initAmqp() {
	config := global.AmqpConfig
	amqpConn, err := amqp.Dial(fmt.Sprintf(constant.AmqpURL, config.User, config.Password, config.Host, config.Port))
	if err != nil {
		hlog.Fatal("cannot dial amqp", err)
	}
	global.AmqpConn = amqpConn
}
