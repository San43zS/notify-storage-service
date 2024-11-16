package brokerNotification

import (
	"Notify-storage-service/internal/broker"
	"Notify-storage-service/internal/broker/rabbit/consumer"
	"Notify-storage-service/internal/broker/rabbit/producer"
	msg2 "Notify-storage-service/internal/handler/model/msg"
	"context"
	"fmt"
	"log"
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

	newMsg, err := msg2.New().Unparse(msg)
	if err != nil {
		log.Println("Failed to add notification: %w", err)
		return fmt.Errorf("failed to add notification: %w", err)
	}

	test := string(newMsg)
	log.Println(test)

	err = s.p.Produce(ctx, newMsg)
	if err != nil {
		log.Println("Failed to add notification: ", err)
		return fmt.Errorf("failed to add notification: %w", err)
	}

	return nil
}

func (s RespCons) GetOld(ctx context.Context) ([]byte, error) {
	consume, err := s.c.Consume(ctx)
	if err != nil {
		err = fmt.Errorf("Failed to get notifications without ttl:  %w", err)
		log.Println(err.Error())

		return nil, err
	}

	return consume, nil
}

func (s RespCons) GetCurrent(ctx context.Context) ([]byte, error) {
	consume, err := s.c.Consume(ctx)
	if err != nil {
		log.Println("Failed to get notifications with ttl: ", err)
		return nil, fmt.Errorf("failed to get notifications with ttl: %w", err)
	}

	return consume, nil
}
