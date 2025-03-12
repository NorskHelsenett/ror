package rabbitmq

import (
	"log/slog"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ConsumerOption is an interface used for passing and applying configuration to
// a Consumer.
type ConsumerOption interface {
	apply(*consumer) *consumer
}

type endpointOption string

func WithEndpoint(endpoint string) ConsumerOption {
	return endpointOption(endpoint)
}

func (o endpointOption) apply(c *consumer) *consumer {
	c.endpoint = string(o)
	return c
}

type nameOption string

func WithName(name string) ConsumerOption {
	return nameOption(name)
}

func (o nameOption) apply(c *consumer) *consumer {
	c.consumerName = string(o)
	return c
}

type queueNameOption string

func WithQueueName(name string) ConsumerOption {
	return queueNameOption(name)
}

func (o queueNameOption) apply(c *consumer) *consumer {
	c.queueName = string(o)
	return c
}

type exchangeNameOption string

func WithExchangeName(name string) ConsumerOption {
	return exchangeNameOption(name)
}

func (o exchangeNameOption) apply(c *consumer) *consumer {
	c.exchangeName = string(o)
	return c
}

type routingKeyOption string

func WithRoutingKey(routingKey string) ConsumerOption {
	return routingKeyOption(routingKey)
}

func (o routingKeyOption) apply(c *consumer) *consumer {
	c.routingKey = string(o)
	return c
}

type argsOption amqp.Table

func WithArgs(args amqp.Table) ConsumerOption {
	return argsOption(args)
}

func (o argsOption) apply(c *consumer) *consumer {
	c.args = amqp.Table(o)
	return c
}

type loggerOption slog.Logger

func WithLogger(logger *slog.Logger) ConsumerOption {
	return loggerOption(*logger)
}

func (o loggerOption) apply(c *consumer) *consumer {
	logger := slog.Logger(o)
	c.logger = &logger
	return c
}

func applyEnvOptions(c *consumer) *consumer {
	opts := getOptionsFromEnv()
	for _, opt := range opts {
		c = opt.apply(c)
	}
	return c
}

func getOptionsFromEnv() []ConsumerOption {
	var opts []ConsumerOption
	endpoint, ok := os.LookupEnv("ROR_RABBITMQ_ENDPOINT")
	if ok {
		opts = append(opts, endpointOption(endpoint))
	}
	consumerName, ok := os.LookupEnv("ROR_RABBITMQ_CONSUMER_NAME")
	if ok {
		opts = append(opts, nameOption(consumerName))
	}
	queueName, ok := os.LookupEnv("ROR_RABBITMQ_QUEUE_NAME")
	if ok {
		opts = append(opts, queueNameOption(queueName))
	}
	exchangeName, ok := os.LookupEnv("ROR_RABBITMQ_EXCHANGE_NAME")
	if ok {
		opts = append(opts, exchangeNameOption(exchangeName))
	}
	routingKey, ok := os.LookupEnv("ROR_RABBITMQ_ROUTING_KEY")
	if ok {
		opts = append(opts, routingKeyOption(routingKey))
	}
	return opts
}
