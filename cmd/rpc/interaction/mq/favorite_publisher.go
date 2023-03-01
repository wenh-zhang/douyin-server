package mq

import (
	"douyin/cmd/rpc/interaction/model"
	"github.com/bytedance/sonic"
	"github.com/streadway/amqp"
)

type FavoriteWithAction struct {
	Favorite   *model.Favorite
	ActionType int8
}

type FavoritePublisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewFavoritePublisher(conn *amqp.Connection, exchange string) *FavoritePublisher {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	if err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
		panic(err)
	}
	return &FavoritePublisher{
		ch:       ch,
		exchange: exchange,
	}
}

func (s *FavoritePublisher) Publish(favorite *FavoriteWithAction) error {
	body, err := sonic.Marshal(favorite)
	if err != nil {
		return err
	}
	return s.ch.Publish(s.exchange, "", false, false, amqp.Publishing{
		Body: body,
	})
}
