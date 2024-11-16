package rabbit

import cnfg "Notify-storage-service/internal/broker/rabbit/config"

type Consumer struct {
	QueueName  string
	RoutingKey string
}

type Config struct {
	Consumers []Consumer
}

func NewCfg() Config {
	config := Config{}

	consumer := Consumer{
		QueueName:  cnfg.HConsumerQueueName,
		RoutingKey: cnfg.HConsumerRoutingKey,
	}

	config.Consumers = append(config.Consumers, consumer)

	consumer = Consumer{
		QueueName:  cnfg.UConsumerQueueName,
		RoutingKey: cnfg.UConsumerRoutingKey,
	}

	config.Consumers = append(config.Consumers, consumer)

	return config
}
