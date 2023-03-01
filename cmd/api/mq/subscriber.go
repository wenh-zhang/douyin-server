package mq

import (
	"douyin/cmd/api/service"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/streadway/amqp"
)

type Subscriber struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
}

func NewSubscriber(conn *amqp.Connection, exchange string) *Subscriber {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	if err = ch.ExchangeDeclare(exchange, "fanout", true, false, false, false, nil); err != nil {
		panic(err)
	}
	return &Subscriber{
		conn:     conn,
		ch:       ch,
		exchange: exchange,
	}
}

func (s *Subscriber) Subscribe() (chan *service.UploadInfo, func(), error) {
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
	uploadInfoChan := make(chan *service.UploadInfo)
	go func() {
		for msg := range msgs {
			info := new(service.UploadInfo)
			if err = sonic.Unmarshal(msg.Body, info); err != nil {
				hlog.Errorf("cannot unmarshal msg: %s", err.Error())
			} else {
				uploadInfoChan <- info
			}
		}
		close(uploadInfoChan)
	}()
	return uploadInfoChan, closeFunc, nil
}
