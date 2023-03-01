package mq

import (
	"douyin/cmd/rpc/interaction/model"
	"github.com/bytedance/sonic"
	"github.com/streadway/amqp"
)

type CommentWithAction struct {
	Comment    *model.Comment
	ActionType int8
}

type CommentPublisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewCommentPublisher(conn *amqp.Connection, exchange string) *CommentPublisher {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	if err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
		panic(err)
	}
	return &CommentPublisher{
		ch:       ch,
		exchange: exchange,
	}
}

func (s *CommentPublisher) Publish(comment *CommentWithAction) error {
	body, err := sonic.Marshal(comment)
	if err != nil {
		return err
	}
	return s.ch.Publish(s.exchange, "", false, false, amqp.Publishing{
		Body: body,
	})
}
