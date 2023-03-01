package mq

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/streadway/amqp"
)

type CommentSubscriber struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
}

func NewCommentSubscriber(conn *amqp.Connection, exchange string) *CommentSubscriber {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	if err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
		panic(err)
	}
	return &CommentSubscriber{
		conn:     conn,
		ch:       ch,
		exchange: exchange,
	}
}

func (s *CommentSubscriber) Subscribe() (chan *CommentWithAction, func(), error) {
	closeConnFunc := func() {
		if err := s.conn.Close(); err != nil {
			hlog.Errorf("cannot close amqp connection: %s", err.Error())
		}
	}
	queue, err := s.ch.QueueDeclare("", true, true, false, false, nil)
	if err != nil {
		return nil, closeConnFunc, err
	}
	closeFunc := func() {
		_, err = s.ch.QueueDelete(queue.Name, false, false, false)
		if err != nil {
			hlog.Errorf("cannot delete queue: %s", err.Error())
		}
		closeConnFunc()
	}
	if err = s.ch.QueueBind(queue.Name, "", s.exchange, false, nil); err != nil {
		return nil, closeFunc, err
	}
	msgs, err := s.ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, closeFunc, err
	}
	commentChan := make(chan *CommentWithAction)
	go func() {
		for msg := range msgs {
			comment := new(CommentWithAction)
			if err = sonic.Unmarshal(msg.Body, comment); err != nil {
				hlog.Errorf("cannot unmarshal msg: %s", err.Error())
			} else {
				commentChan <- comment
			}
		}
		close(commentChan)
	}()
	return commentChan, closeFunc, nil
}
