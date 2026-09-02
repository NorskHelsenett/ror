package aclevents

import (
	"context"
	"time"

	aclstore "github.com/NorskHelsenett/ror/pkg/acl/aclstore/v2"
	"github.com/NorskHelsenett/ror/pkg/clients/rabbitmqclient"
)

const (
	// DefaultExchangeName is the fanout exchange ACL-change signals are delivered
	// to (via a binding from the main ror exchange).
	DefaultExchangeName = "ror.acl.events"
	// DefaultRoutingKey is the routing key used on the main ror exchange to reach
	// the ACL fanout exchange.
	DefaultRoutingKey = "ror.acl.changed"
)

// changeSignal is the opaque payload broadcast on every ACL change. Subscribers
// react by reloading their whole snapshot, so it carries no entry detail.
type changeSignal struct {
	Timestamp time.Time `json:"timestamp"`
}

// Publisher broadcasts ACL-change signals over the shared rabbitmq connection so
// every subscribing process refreshes its snapshot.
type Publisher struct {
	conn       rabbitmqclient.RabbitMQConnection
	routingKey string
}

// compile-time assurance that Publisher satisfies the store's ChangePublisher.
var _ aclstore.ChangePublisher = (*Publisher)(nil)

// NewPublisher creates a publisher on the shared rabbitmq connection using the
// given routing key (see DefaultRoutingKey).
func NewPublisher(conn rabbitmqclient.RabbitMQConnection, routingKey string) *Publisher {
	return &Publisher{conn: conn, routingKey: routingKey}
}

// PublishChange broadcasts an opaque "something changed" signal on the ror
// exchange with the configured routing key.
func (p *Publisher) PublishChange(ctx context.Context) error {
	return p.conn.SendMessage(ctx, changeSignal{Timestamp: time.Now()}, p.routingKey, nil)
}
