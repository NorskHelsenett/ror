package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
	"os"
)

// option is an interface for passing configuration options to the structs in
// this package.
type option interface {
	setEndpoint(string)
	setConsumerName(string)
	setQueueName(string)
	setExchangeName(string)
	setRoutingKey(string)
	setArgs(amqp.Table)
	setLogger(*slog.Logger)
	setAuthenticator(Authenticator)
	setExchangeType(string)
}

// Option is used in a functional options pattern to apply configuration options
// to the consumer and producer structs provided in this package.
type Option func(option)

func WithConsumerName(name string) Option {
	return func(o option) {
		o.setConsumerName(name)
	}
}

func WithEndpoint(endpoint string) Option {
	return func(o option) {
		o.setEndpoint(endpoint)
	}
}

func WithQueueName(name string) Option {
	return func(o option) {
		o.setQueueName(name)
	}
}

func WithRoutingKey(key string) Option {
	return func(o option) {
		o.setRoutingKey(key)
	}
}

func WithArgs(args amqp.Table) Option {
	return func(o option) {
		o.setArgs(args)
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(o option) {
		o.setLogger(logger)
	}
}

func WithAuthenticator(authenticator Authenticator) Option {
	return func(o option) {
		o.setAuthenticator(authenticator)
	}
}

func WithExchangeName(name string) Option {
	return func(o option) {
		o.setExchangeName(name)
	}
}

func WithExchangeType(t string) Option {
	return func(o option) {
		o.setExchangeType(t)
	}
}

// applyEnvOptions gets options from the environment and applies them to the
// provided option interface.
func applyEnvOptions(o option) {
	opts := getOptionsFromEnv()
	for _, opt := range opts {
		opt(o)
	}
}

// getOptionsFromEnv gets various configuration options from environment
// variables and returns an Option slice.
func getOptionsFromEnv() []Option {
	var opts []Option
	endpoint, ok := os.LookupEnv("ROR_RABBITMQ_ENDPOINT")
	if ok {
		opts = append(opts, WithEndpoint(endpoint))
	}
	consumerName, ok := os.LookupEnv("ROR_RABBITMQ_CONSUMER_NAME")
	if ok {
		opts = append(opts, WithConsumerName(consumerName))
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
	exchangeType, ok := os.LookupEnv("ROR_RABBITMQ_EXCHANGE_TYPE")
	if ok {
		opts = append(opts, WithExchangeType(exchangeType))
	}
	return opts
}
