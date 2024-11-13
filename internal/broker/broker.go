package broker

import (
	"Notify-storage-service/internal/broker/rabbit"
	"fmt"
)

type Broker struct {
	RabbitMQ rabbit.Service
}

func New() (Broker, error) {
	rabbitMQ, err := rabbit.New()
	if err != nil {

		return Broker{}, fmt.Errorf("failed to create RabbitMQ broker: %w", err)
	}

	broker := Broker{
		RabbitMQ: rabbitMQ,
	}

	return broker, nil
}
