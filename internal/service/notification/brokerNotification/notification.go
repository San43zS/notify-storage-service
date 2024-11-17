package brokerNotification

import (
	"Notify-storage-service/internal/broker"
	"Notify-storage-service/internal/broker/rabbit/consumer"
	"Notify-storage-service/internal/broker/rabbit/producer"
	msg2 "Notify-storage-service/internal/handler/model/msg"
	"Notify-storage-service/internal/handler/model/msg/parser/msgParser"
	"context"
	"fmt"
)

type RespCons struct {
	p producer.Producer
	c consumer.Consumer
}

func New(broker broker.Broker) RespCons {
	return RespCons{
		p: broker.RabbitMQ.Producer(),
		c: broker.RabbitMQ.Consumer(),
	}
}

func (s RespCons) Add(ctx context.Context, msg msg2.MSG) error {

	newMsg, err := msgParser.New().Unparse(msg)
	if err != nil {
		return fmt.Errorf("failed to add notification: %w", err)
	}

	err = s.p.Produce(ctx, newMsg)
	if err != nil {
		return fmt.Errorf("failed to add notification: %w", err)
	}

	return nil
}

func (s RespCons) Send(ctx context.Context, msg []byte) error {
	err := s.p.Produce(ctx, msg)
	if err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	return nil
}
