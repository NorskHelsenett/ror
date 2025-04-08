package rabbitmq

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Consumer interface used for defining a rabbitmq consumer. The Connection
// interface is embedded in this interface.
type Consumer interface {
	Connection
	Consume(context.Context) error
}

// consumer struct used for representing a rabbitmq consumer. The connection
// struct is embedded in this struct. The handler in this struct handles
// messages consumed from rabbitmq.
type consumer struct {
	*connection
	consumerName string
	queueName    string
	exchangeName string
	routingKey   string
	args         amqp.Table
	handler      func(context.Context, amqp.Delivery) error
}

// NewConsumer creates a new Consumer with an underlying Connection. The consumer
// is configurable through environment variables or by passing in an array of
// Option.
//
// This method returns an error if the underlying connection fails to connect.
//
// If no configuration is provided, a default connection is set up.
//
// Pass an option WithAuthenticator to use a custom Authenticator.
func NewConsumer(handler func(context.Context, amqp.Delivery) error, opts ...Option) (Consumer, error) {
	// Create a default connection with a default authenticator and uuids for
	// consumer names and keys.
	c := &consumer{
		connection:   newConnection(),
		consumerName: uuid.NewString(),
		queueName:    uuid.NewString(),
		exchangeName: uuid.NewString(),
		routingKey:   uuid.NewString(),
	}

	// Apply consumer overrides from the environment and options passed in the
	// constructor. The options passed in the constructor take precedence.
	applyEnvOptions(c)
	for _, opt := range opts {
		opt(c)
	}

	// Register the handler provided in the constructor.
	if handler == nil {
		c.logger.Warn("starting consumer with default No-Op handler")
		c.handler = defaultHandler
	} else {
		c.handler = handler
	}

	err := c.connect()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Consume declares and binds the queue for this consumer and starts
// the consume loop. It handles the consumed messages using the registered
// handler. If a reconnect happens the consume loop is restarted.
//
// This method returns an error if the queue setup fails.
//
// This is a blocking method. If the provided context is cancelled this method
// will return.
func (c *consumer) Consume(ctx context.Context) error {
	// If the connection is in shutdown mode we return immediately.
	if c.connectionShutdown {
		c.logger.InfoContext(ctx, "connection shutdown started, not restarting consume loop")
		return nil
	}

	// We wait for the reconnection to finish.
	c.reconnect.Wait()

	err := c.setupConsumeQueue()
	if err != nil {
		return err
	}

	deliveryChan, err := c.amqpChannel.ConsumeWithContext(ctx, c.queueName, c.consumerName, false, false, // the noLocal flag is not supported but rabbitmq, so we just set to false.
		false, false, c.args)
	if err != nil {
		return err
	}

	c.logger.Info("consumer starting")

consumeLoop:
	for {
		select {
		case delivery, ok := <-deliveryChan:
			// If the delivery channel is cancelled or closed we break the consume loop.
			if !ok {
				break consumeLoop
			}

			// We extract trace headers from the delivery. If no trace headers are found this
			// is a No-Op.
			ctx = trace.ExtractAMQPHeaders(ctx, delivery.Headers)
			err = c.handler(ctx, delivery)
			if err != nil {
				c.logger.ErrorContext(ctx, "failed to handle message", "error", err)
			}

			// We only ack the delivery if the handler error is nil.
			if err == nil {
				err = c.amqpChannel.Ack(delivery.DeliveryTag, false)
				if err != nil {
					c.logger.ErrorContext(ctx, "failed to ack delivery", "error", err)
				}
			}

		// We need to handle connection close since the delivery channel will not be
		// closed if the connection is closed.
		case <-c.connectionCloseChan:
			break consumeLoop

		case <-ctx.Done():
			return nil
		}
	}

	// We restart the consume loop if the consumeLoop is broken
	return c.Consume(ctx)
}

// setupConsumeQueue declares the queue used by this consumer, and then binds the
// queue to an exchange. If any of the operations fails this method returns an
// error.
func (c *consumer) setupConsumeQueue() error {
	_, err := c.amqpChannel.QueueDeclare(c.queueName, false, true, false, false, c.args)
	if err != nil {
		c.logger.Error("failed to declare consumer queue", "error", err)
		return err
	}

	err = c.amqpChannel.QueueBind(c.queueName, c.routingKey, c.exchangeName, false, c.args)
	if err != nil {
		c.logger.Error("failed to bind consumer queue to exchange", "error", err)
		return err
	}

	return nil
}

// defaultHandler returns a No-Op for handling messages.
func defaultHandler(context.Context, amqp.Delivery) error {
	return nil
}

func (c *consumer) setArgs(a amqp.Table) {
	c.args = a
}

func (c *consumer) setQueueName(q string) {
	c.queueName = q
}

func (c *consumer) setRoutingKey(r string) {
	c.routingKey = r
}

func (c *consumer) setExchangeName(e string) {
	c.exchangeName = e
}

func (c *consumer) setConsumerName(n string) {
	c.consumerName = n
}

// No-Op method to satisfy the option interface
func (c *consumer) setExchangeType(s string) {}
