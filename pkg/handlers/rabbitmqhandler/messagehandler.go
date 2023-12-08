package rabbitmqhandler

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rabbitmqclient"

	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQMessageHandler interface {
	HandleMessage(ctx context.Context, message amqp.Delivery) error
}

type RabbitMQListner struct {
	Client    rabbitmqclient.RabbitMQConnection
	QueueName string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
	Handler   RabbitMQMessageHandler
}

type RabbitMQListnerConfig struct {
	Client    rabbitmqclient.RabbitMQConnection
	QueueName string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func New(config RabbitMQListnerConfig, handler RabbitMQMessageHandler) RabbitMQListner {
	return RabbitMQListner{
		Client:    config.Client,
		QueueName: config.QueueName,
		Consumer:  config.Consumer,
		AutoAck:   config.AutoAck,
		Exclusive: config.Exclusive,
		NoLocal:   config.NoLocal,
		NoWait:    config.NoWait,
		Args:      config.Args,
		Handler:   handler,
	}
}

func (r RabbitMQListner) Listen(hangup chan *amqp.Error) {
	messages, err := r.Client.GetChannel().Consume(
		r.QueueName, // queue
		r.Consumer,  // consumer
		r.AutoAck,   // auto-ack
		r.Exclusive, // exclusive
		r.NoLocal,   // no-local
		r.NoWait,    // no-wait
		r.Args,      // args
	)
	if err != nil {
		rlog.Fatal("failed to register a consumer on queue", err, rlog.String("queue", r.QueueName))
	}

	rlog.Info("listening on RabbitMQ queue", rlog.String("queue", r.QueueName))

	go func() {
		for message := range messages {
			ctx := trace.ExtractAMQPHeaders(context.Background(), message.Headers)

			err := r.Handler.HandleMessage(ctx, message)
			if err != nil {
				rlog.Errorc(ctx, "Could not handle message", err, rlog.Any("event", message))
			} else {
				err = message.Ack(true)
				if err != nil {
					rlog.Errorc(ctx, "Could not ack message", err, rlog.Any("event", message))

				}
			}
		}
	}()
	<-hangup
}
