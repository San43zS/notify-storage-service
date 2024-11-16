package consumer

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	Consume(ctx context.Context, q string) ([]byte, error)
}

type consumer struct {
	dial   *amqp.Channel
}

func New(dial *amqp.Channel) Consumer {
	return &consumer{
		dial:   dial,
	}
}

func (c consumer) Consume(ctx context.Context, q string) ([]byte, error) {

	msgs, err := c.dial.Consume(
		q,                          // queue
		"", // consumer
		false,                      // auto-ack
		false,                      // exclusive
		false,                      // no-local
		true,                       // no-wait
		nil,                        // args
	)
	if err != nil {
		return nil, err
	}

	for msg := range msgs {
		return msg.Body, nil
	}

	return nil, nil
}

