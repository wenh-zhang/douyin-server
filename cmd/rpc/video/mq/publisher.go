package mq

import (
	"douyin/cmd/rpc/video/model"
	"github.com/bytedance/sonic"
	"github.com/streadway/amqp"
)

type Publisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewPublisher(conn *amqp.Connection, exchange string) *Publisher {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	if err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
		panic(err)
	}
	return &Publisher{
		ch:       ch,
		exchange: exchange,
	}
}

func (s *Publisher) Publish(video *model.Video) error {
	body, err := sonic.Marshal(video)
	if err != nil {
		return err
	}
	return s.ch.Publish(s.exchange, "", false, false, amqp.Publishing{
		Body: body,
	})
}
