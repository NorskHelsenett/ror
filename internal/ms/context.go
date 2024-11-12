package ms

import (
	"context"
	"time"

	"github.com/NorskHelsenett/ror/pkg/context/mscontext"

	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Context struct {
	Ctx context.Context

	Delivery amqp.Delivery
	Channel  *amqp.Channel
}

func newContext(gctx context.Context, delivery amqp.Delivery, channel *amqp.Channel, role string) (*Context, error) {
	ctx := new(Context)

	ctx.Ctx = mscontext.GetRorContextFromServiceContextWithoutCancel(gctx, role)
	ctx.Delivery = delivery
	ctx.Channel = channel

	return ctx, nil
}

func (c *Context) Body() []byte {
	return c.Delivery.Body
}

func (c *Context) Ack(multiple bool) error {
	return c.Delivery.Ack(multiple)
}

func (c *Context) Publish(workqueue string, message []byte) error {
	q, err := c.Channel.QueueDeclare(
		workqueue,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	headers := trace.InjectAMQPHeaders(c)

	err = c.Channel.PublishWithContext(
		c,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
			Headers:     headers,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// Forward all functions from Context.Ctx (golang context) so we don't have to refer into the struct. For cleanliness.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.Ctx.Done()
}

func (c *Context) Err() error {
	return c.Ctx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	return c.Ctx.Value(key)
}
