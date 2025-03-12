package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Connection interface for handling rabbitmq connections.
type Connection interface {
	Ping(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

// connection struct for handling rabbitmq connections. The Authenticator
// interface is embedded in this struct.
type connection struct {
	endpoint            string
	amqpConnection      *amqp.Connection
	amqpChannel         *amqp.Channel
	healthExchange      string
	healthQueue         string
	healthRoutingKey    string
	logger              *slog.Logger
	connectionShutdown  bool
	connectionCloseChan chan *amqp.Error
	channelCloseChan    chan *amqp.Error
	channelCancelChan   chan string
	reconnect           sync.WaitGroup
	Authenticator
}

// connect sets up the amqpConnection and amqpChannel, runs setupHealthQueue, and
// runs handleReconnect in a separate goroutine. It also creates the
// connectionCloseChan, channelCloseChan, and channelCancelChan and registers
// them as listeners on the amqpConnection and amqpChannel respectively.
func (c *connection) connect() error {
	var err error
	c.amqpConnection, err = amqp.Dial(c.getConnectionString())
	if err != nil {
		c.logger.Error("failed to connect to RabbitMQ", "error", err)
		return err
	}

	c.connectionCloseChan = make(chan *amqp.Error)
	c.amqpConnection.NotifyClose(c.connectionCloseChan)

	c.amqpChannel, err = c.amqpConnection.Channel()
	if err != nil {
		c.logger.Error("failed to open a channel", "error", err)
		return err
	}

	c.channelCloseChan = make(chan *amqp.Error)
	c.amqpChannel.NotifyClose(c.channelCloseChan)
	c.channelCancelChan = make(chan string)
	c.amqpChannel.NotifyCancel(c.channelCancelChan)

	err = c.setupHealthQueue()
	if err != nil {
		c.logger.Error("failed to setup health queue", "error", err)
		return err
	}

	go c.handleReconnect()

	c.logger.Info("connected successfully to rabbitmq")
	return nil
}

// handleReconnect listens for amqpConnection close, amqpChannel close, or
// amqpChannel cancel. If a signal is received, handleReconnect closes all
// remaining connections and runs the connect function.
func (c *connection) handleReconnect() {
	if c.connectionShutdown {
		return
	}
	select {
	case <-c.connectionCloseChan:
		if c.connectionShutdown {
			return
		}
		c.reconnect.Add(1)
		c.logger.Info("connection closed, reconnecting")
		err := c.connect()
		if err != nil {
			c.logger.Error("failed to reconnect", "error", err)
		}
		c.reconnect.Done()
	case <-c.channelCloseChan:
		if c.connectionShutdown {
			return
		}
		c.reconnect.Add(1)
		c.logger.Info("channel closed, reconnecting")
		err := c.amqpConnection.Close()
		if err != nil {
			c.logger.Error("failed to close connection", "error", err)
		}
		err = c.connect()
		if err != nil {
			c.logger.Error("failed to reconnect", "error", err)
		}
		c.reconnect.Done()
	case <-c.channelCancelChan:
		if c.connectionShutdown {
			return
		}
		c.reconnect.Add(1)
		c.logger.Info("channel cancelled, reconnecting")
		err := c.amqpChannel.Close()
		if err != nil {
			c.logger.Error("failed to close channel", "error", err)
		}
		err = c.amqpConnection.Close()
		if err != nil {
			c.logger.Error("failed to close connection", "error", err)
		}
		err = c.connect()
		if err != nil {
			c.logger.Error("failed to reconnect", "error", err)
		}
		c.reconnect.Done()
	}
}

// Shutdown puts the connection in shutdown mode and closes the amqp channel and
// connection.
func (c *connection) Shutdown(ctx context.Context) error {
	c.logger.Info("shutting rabbitmq down connection")
	c.connectionShutdown = true

	err := c.amqpChannel.Close()
	if err != nil {
		c.logger.Error("failed to close channel", "error", err)
		return err
	}

	// use context deadline when shutting down amqp connection if it exists
	deadline, ok := ctx.Deadline()
	if !ok {
		err = c.amqpConnection.Close()
	} else {
		err = c.amqpConnection.CloseDeadline(deadline)
	}
	if err != nil {
		c.logger.Error("failed to close connection", "error", err)
		return err
	}

	c.logger.Info("rabbitmq connection shut down")

	return nil
}

// setupHealthQueue declares an exchange and a queue with unique names and
// routing key that is used in the Ping function.
func (c *connection) setupHealthQueue() error {
	c.healthExchange = uuid.NewString()
	c.healthQueue = uuid.NewString()
	c.healthRoutingKey = uuid.NewString()

	err := c.amqpChannel.ExchangeDeclare(
		c.healthExchange,
		amqp.ExchangeDirect,
		false,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	_, err = c.amqpChannel.QueueDeclare(
		c.healthQueue,
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = c.amqpChannel.QueueBind(
		c.healthQueue,
		c.healthRoutingKey,
		c.healthExchange,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	c.logger.Debug("health queue setup completed")

	return nil
}

// Ping publishes to the connections health exchange and gets the delivery from
// the health queue synchronously. This method can be used to check the health of
// the connection.
func (c *connection) Ping(ctx context.Context) error {
	// we return immediately if the connection is in shutdown mode
	if c.connectionShutdown {
		return nil
	}

	// we wait for either the context to be done or for the reconnect to finish
	waitChan := make(chan struct{})
	go func() {
		defer close(waitChan)
		c.reconnect.Wait()
	}()
	select {
	case <-waitChan:
		break
	case <-ctx.Done():
		return ctx.Err()
	}

	err := c.amqpChannel.PublishWithContext(ctx, c.healthExchange, c.healthRoutingKey, true, false, amqp.Publishing{})
	if err != nil {
		c.logger.Error("failed to publish to health exchange", "error", err)
		return err
	}

	d, ok, err := c.amqpChannel.Get(c.healthQueue, false)
	if err != nil {
		c.logger.Error("failed to get from health queue", "error", err)
		return err
	}
	if !ok {
		c.logger.Error("no delivery on health channel")
		return errors.New("no delivery on health channel")
	}
	err = d.Ack(false)
	if err != nil {
		c.logger.Error("failed to ack delivery on health channel", "error", err)
		return err
	}

	return nil
}

// getConnectionString creates a connectionString from the connections endpoint
// and the credentials provided from the registered Authenticator.
func (c *connection) getConnectionString() string {
	username, password := c.GetCredentials()
	return fmt.Sprintf("amqp://%s:%s@%s", username, password, c.endpoint)
}
