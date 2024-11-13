package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

const (
	ProducerQueueName    = "purpleQueue"
	ProducerExchangeName = "test"
	ContextTimeOut       = 5 * time.Second
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
