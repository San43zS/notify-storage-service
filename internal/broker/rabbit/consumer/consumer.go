package consumer

import (
	"Notify-storage-service/internal/broker/rabbit/config"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer interface {
	UConsume(ctx context.Context) ([]byte, error)
	HConsume(ctx context.Context) ([]byte, error)
}

type consumer struct {
	dial *amqp.Channel
}

func New(dial *amqp.Channel) Consumer {
	return &consumer{
		dial: dial,
	}
}

func (c consumer) UConsume(ctx context.Context) ([]byte, error) {

	msgs, err := c.dial.Consume(
		config.UConsumerQueueName,  // queue
		config.UserServiceConsumer, // consumer
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

func (c consumer) HConsume(ctx context.Context) ([]byte, error) {
	msgs, err := c.dial.Consume(
		config.HConsumerQueueName,    // queue
		config.HandleServiceConsumer, // consumer
		false,                        // auto-ack
		false,                        // exclusive
		false,                        // no-local
		true,                         // no-wait
		nil,                          // args
	)
	if err != nil {
		return nil, err
	}

	for msg := range msgs {
		return msg.Body, nil
	}

	return nil, nil
}
