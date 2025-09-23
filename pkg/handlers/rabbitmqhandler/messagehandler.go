package rabbitmqhandler

import (
	"context"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/rabbitmqclient"
	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQMessageHandler interface {
	HandleMessage(ctx context.Context, message amqp.Delivery) error
}

type RabbitMQListner struct {
	Client rabbitmqclient.RabbitMQConnection

	isInitialized      bool
	queueName          string
	queueConsumer      string
	queueAutoAck       bool
	queueAutoDelete    bool
	queueExclusive     bool
	queueNoLocal       bool
	queueNoWait        bool
	queueArgs          amqp.Table
	Handler            RabbitMQMessageHandler
	exchange           string
	excahngeKind       string
	excahngeRoutingKey string
	exchangeAutoDelete bool
	excahngeDurable    bool
}

type RabbitMQListnerConfig struct {
	Client             rabbitmqclient.RabbitMQConnection
	QueueName          string
	Consumer           string
	AutoAck            bool
	QueueAutoDelete    bool
	Exclusive          bool
	NoLocal            bool
	NoWait             bool
	Args               amqp.Table
	Exchange           string
	ExcahngeKind       string
	ExcahngeDurable    bool
	ExcahngeRoutingKey string
	ExchangeAutoDelete bool
}

func New(config RabbitMQListnerConfig, handler RabbitMQMessageHandler) RabbitMQListner {
	return RabbitMQListner{
		Client:             config.Client,
		queueName:          config.QueueName,
		queueConsumer:      config.Consumer,
		queueAutoAck:       config.AutoAck,
		queueAutoDelete:    config.QueueAutoDelete,
		queueExclusive:     config.Exclusive,
		queueNoLocal:       config.NoLocal,
		queueNoWait:        config.NoWait,
		queueArgs:          config.Args,
		Handler:            handler,
		exchange:           config.Exchange,
		excahngeKind:       config.ExcahngeKind,
		excahngeRoutingKey: config.ExcahngeRoutingKey,
		exchangeAutoDelete: config.ExchangeAutoDelete,
		excahngeDurable:    config.ExcahngeDurable,
	}
}

// ListenWithTTL Convience method for setting up channel with TTL on messages
func (r RabbitMQListner) ListenWithTTL(hangup chan *amqp.Error, TTL time.Duration) {
	// create new args if listener is declared without args, otherwise override ttl
	if r.queueArgs == nil {
		r.queueArgs = amqp.Table{
			amqp.QueueMessageTTLArg: TTL.Milliseconds(),
		}
	} else {
		r.queueArgs[amqp.QueueMessageTTLArg] = TTL.Milliseconds()
	}

	r.Listen(hangup)
}

func (r RabbitMQListner) Listen(hangup chan *amqp.Error) {
	if (r.exchangeAutoDelete && r.exchange != "") || (!r.isInitialized && r.exchange != "") {
		// (re)Declare exchange if it is not declared or if it is set to auto delete{
		err := r.Client.GetChannel().ExchangeDeclare(
			r.exchange,           // name
			r.excahngeKind,       // kind
			r.excahngeDurable,    // durable
			r.exchangeAutoDelete, // autoDelete -> delete when unused
			false,                // internal
			false,                // no-wait
			nil,                  // arguments
		)
		if err != nil {
			rlog.Fatal("Could not declare excahnge", err)
		}

		err = r.Client.GetChannel().ExchangeBind(
			r.exchange,                      //destination
			r.excahngeRoutingKey,            // key
			messagebuscontracts.ExchangeRor, // source
			false,                           // noWait
			nil,                             // arguments
		)
		if err != nil {
			panic(err)
		}
	}

	queue, err := r.Client.GetChannel().QueueDeclare(
		r.queueName,       // name
		true,              // durable
		r.queueAutoDelete, // delete when unused
		r.queueExclusive,  // exclusive
		r.queueNoWait,     // no-wait
		r.queueArgs,       // arguments
	)
	if err != nil {
		rlog.Fatal("failed to declare a queue", err, rlog.String("queue", r.queueName))
	}

	if r.exchange != "" {
		err = r.Client.GetChannel().QueueBind(
			queue.Name,           // queue name
			r.excahngeRoutingKey, // routing key
			r.exchange,           // exchange
			r.queueNoWait,
			r.queueArgs,
		)
		if err != nil {
			rlog.Fatal("failed to bind queue to exchange", err, rlog.String("queue", r.queueName), rlog.String("exchange", r.exchange))
		}
	}

	messages, err := r.Client.GetChannel().Consume(
		queue.Name,       // queue
		r.queueConsumer,  // consumer
		r.queueAutoAck,   // auto-ack
		r.queueExclusive, // exclusive
		r.queueNoLocal,   // no-local
		r.queueNoWait,    // no-wait
		r.queueArgs,      // args
	)
	if err != nil {
		rlog.Fatal("failed to register a consumer on queue", err, rlog.String("queue", r.queueName))
	}

	// set as initialized
	r.isInitialized = true

	rlog.Info("listening on RabbitMQ queue", rlog.String("queue", r.queueName))

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
