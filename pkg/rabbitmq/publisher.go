package rabbitmq

import (
	"context"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher is an interface used for defining methods on a rabbitmq publisher.
// The Connection interface is embedded in this interface.
type Publisher interface {
	Connection
	Publish(context.Context, amqp.Publishing) error
}

// publisher struct used for representing a rabbitmq publisher. The connection
// struct is embedded in this struct.
type publisher struct {
	*connection
	exchangeName string
	exchangeType string
	routingKey   string
	args         amqp.Table
}

// NewPublisher creates a new Publisher with an underlying Connection. The publisher is configurable through environment variables or by passing and array of Option.
//
// This method returns an error if the underlying connection fails to connect.
//
// If no configuration is provided, a default connection and publisher is set up.
//
// Pass an option WithAuthenticator to use a custom Authenticator
func NewPublisher(opts ...Option) Publisher {
	// Create a default connection with a default authenticator and uuids for
	// producer names and keys.
	p := &publisher{
		connection:   newConnection(),
		exchangeName: uuid.NewString(),
		exchangeType: amqp.ExchangeDirect,
		routingKey:   uuid.NewString(),
	}

	// Apply consumer overrides from the environment and options passed in the
	// constructor. The options passed in the constructor take precedence.
	applyEnvOptions(p)
	for _, opt := range opts {
		opt(p)
	}

	err := p.connect()
	if err != nil {
		panic(err)
	}

	return p
}

// Publish declares the configured exchange and publishes a Publishing on it. If
// the underlying connection is in reconnect mode, this method will wait for the
// reconnect to finish.
//
// This method returns an error if the exchange declaration fails or the
// publishing fails.
func (p *publisher) Publish(ctx context.Context, publishing amqp.Publishing) error {
	// If the connection is in shutdown mode we return immediately.
	if p.connectionShutdown {
		p.logger.InfoContext(ctx, "connection shutdown started, not restarting consume loop")
		return nil
	}

	// We wait for the reconnection to finish.
	p.reconnect.Wait()

	err := p.setupPublishExchange()
	if err != nil {
		return err
	}

	err = p.amqpChannel.PublishWithContext(ctx, p.exchangeName, p.routingKey, false, false, publishing)
	if err != nil {
		return err
	}
	return nil
}

// setupPublishExchange is a convenience method for declaring the exchange with
// the configured settings.
func (p *publisher) setupPublishExchange() error {
	err := p.amqpChannel.ExchangeDeclare(p.exchangeName, p.exchangeType, true, false, false, false, p.args)
	if err != nil {
		return err
	}
	return nil
}

func (p *publisher) setExchangeType(t string) {
	p.exchangeType = t
}

func (p *publisher) setExchangeName(e string) {
	p.exchangeName = e
}

func (p *publisher) setRoutingKey(r string) {
	p.routingKey = r
}

func (p *publisher) setArgs(a amqp.Table) {
	p.args = a
}

// No-Op method to satisfy the option interface
func (p *publisher) setConsumerName(s string) {}

// No-Op method to satisfy the option interface
func (p *publisher) setQueueName(s string) {}
