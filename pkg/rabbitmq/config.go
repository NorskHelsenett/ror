package rabbitmq

import (
	"log/slog"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ConsumerOption is an interface used for passing and applying configuration to
// a Consumer.
type ConsumerOption func(*consumer)

func WithEndpoint(endpoint string) ConsumerOption {
	return func(c *consumer) {
		c.endpoint = endpoint
	}
}

func WithName(name string) ConsumerOption {
	return func(c *consumer) {
		c.consumerName = name
	}
}

func WithQueueName(name string) ConsumerOption {
	return func(c *consumer) {
		c.queueName = name
	}
}

func WithExchangeName(name string) ConsumerOption {
	return func(c *consumer) {
		c.exchangeName = name
	}
}

func WithRoutingKey(routingKey string) ConsumerOption {
	return func(c *consumer) {
		c.routingKey = routingKey
	}
}
func WithArgs(args amqp.Table) ConsumerOption {
	return func(c *consumer) {
		c.args = args
	}
}

func WithLogger(logger *slog.Logger) ConsumerOption {
	return func(c *consumer) {
		c.logger = logger
	}
}

func applyEnvOptions(c *consumer) {
	opts := getOptionsFromEnv()
	for _, opt := range opts {
		opt(c)
	}
}

func getOptionsFromEnv() []ConsumerOption {
	var opts []ConsumerOption
	endpoint, ok := os.LookupEnv("ROR_RABBITMQ_ENDPOINT")
	if ok {
		opts = append(opts, WithEndpoint(endpoint))
	}
	consumerName, ok := os.LookupEnv("ROR_RABBITMQ_CONSUMER_NAME")
	if ok {
		opts = append(opts, WithName(consumerName))
	}
	queueName, ok := os.LookupEnv("ROR_RABBITMQ_QUEUE_NAME")
	if ok {
		opts = append(opts, WithQueueName(queueName))
	}
	exchangeName, ok := os.LookupEnv("ROR_RABBITMQ_EXCHANGE_NAME")
	if ok {
		opts = append(opts, WithExchangeName(exchangeName))
	}
	routingKey, ok := os.LookupEnv("ROR_RABBITMQ_ROUTING_KEY")
	if ok {
		opts = append(opts, WithRoutingKey(routingKey))
	}
	return opts
}
