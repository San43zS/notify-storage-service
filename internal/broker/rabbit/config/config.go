package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	ProducerQueueName    = "purpleQueue"
	ProducerExchangeName = "test"
)

const (
	UConsumerQueueName   = "blueQueue"
	ConsumerExchangeName = "test"
	HConsumerQueueName   = "redQueue"
)

const (
	ProducerRoutingKey  = "purple"
	UConsumerRoutingKey = "blue"
	HConsumerRoutingKey = "red"
)

const (
	UserServiceConsumer   = "UserServiceConsumer"
	HandleServiceConsumer = "HandleServiceConsumer"
)

type Config struct {
	URL    string
	Driver string
}

type amqpParams struct {
	host     string
	port     string
	user     string
	password string
}

func getAMQPParams() *amqpParams {
	return &amqpParams{
		host:     viper.GetString("AMQP.HOST"),
		port:     viper.GetString("AMQP.PORT"),
		user:     viper.GetString("AMQP.USER"),
		password: viper.GetString("AMQP.PASSWORD"),
	}
}

func (amqp amqpParams) ParseURL() string {
	template := viper.GetString("AMQP.URLTEMPLATE")

	return fmt.Sprintf(template, amqp.user, amqp.password, amqp.host, amqp.port)
}

func NewConfig() *Config {
	return &Config{
		URL:    getAMQPParams().ParseURL(),
		Driver: viper.GetString("AMQP.DRIVER"),
	}
}

type Consumer struct {
	QueueName  string
	RoutingKey string
}

type Cnfg struct {
	Consumers []Consumer
}

func NewCfg() Cnfg {
	config := Cnfg{}

	consumer := Consumer{
		QueueName:  HConsumerQueueName,
		RoutingKey: HConsumerRoutingKey,
	}

	config.Consumers = append(config.Consumers, consumer)

	consumer = Consumer{
		QueueName:  UConsumerQueueName,
		RoutingKey: UConsumerRoutingKey,
	}

	config.Consumers = append(config.Consumers, consumer)

	return config
}
