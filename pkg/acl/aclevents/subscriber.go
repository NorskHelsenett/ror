package aclevents

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rabbitmqclient"
	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"

	"uuid"

	amqp "github.com/rabbitmq/amqp091-go"
)

// changeHandler adapts an onChange callback to the rabbitmq message handler.
type changeHandler struct {
	onChange func()
}

// HandleMessage invokes the callback for every delivered change signal.
func (h changeHandler) HandleMessage(_ context.Context, _ amqp.Delivery) error {
	h.onChange()
	return nil
}

// Subscriber consumes ACL-change signals and invokes a callback for each. Every
// process binds its own ephemeral, auto-deleting queue to the fanout exchange,
// so all processes receive every signal.
type Subscriber struct {
	conn       rabbitmqclient.RabbitMQConnection
	exchange   string
	routingKey string
	onChange   func()
}

// NewSubscriber creates a subscriber on the shared rabbitmq connection. onChange
// is invoked for every received signal (typically Refresher.Notify).
func NewSubscriber(conn rabbitmqclient.RabbitMQConnection, exchange, routingKey string, onChange func()) *Subscriber {
	return &Subscriber{conn: conn, exchange: exchange, routingKey: routingKey, onChange: onChange}
}

// Start registers the change listener with the shared connection. Each process
// binds its own unique, auto-deleting queue to the durable fanout exchange, so a
// single published signal reaches every process. The connection owns the consume
// goroutine and restarts it on reconnect.
func (s *Subscriber) Start() error {
	listener := rabbitmqhandler.New(rabbitmqhandler.RabbitMQListnerConfig{
		Client:             s.conn,
		QueueName:          "ror-acl-events-" + uuid.NewV4().String(),
		QueueAutoDelete:    true,
		Exclusive:          true,
		Exchange:           s.exchange,
		ExcahngeKind:       amqp.ExchangeFanout,
		ExcahngeDurable:    true,
		ExchangeAutoDelete: false,
		ExcahngeRoutingKey: s.routingKey,
	}, changeHandler{onChange: s.onChange})

	return s.conn.RegisterHandler(listener)
}
