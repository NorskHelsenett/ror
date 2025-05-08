package rabbitmqclient

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"

	"github.com/dotse/go-health"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQListnerInterface interface {
	Listen(chan *amqp.Error)
}

type RabbitMQConnection interface {
	GetChannel() *amqp.Channel
	RegisterHandler(RabbitMQListnerInterface) error
	SendMessage(ctx context.Context, message any, routing string, extraheaders map[string]interface{}) error
	clients.CommonClient
}

type RabbitMQCredentialProvider interface {
	GetCredentials() (string, string)
}

type rabbitmqcon struct {
	Context            context.Context
	RabbitMqConnection *amqp.Connection
	RabbitMqChannel    *amqp.Channel
	Credentials        RabbitMQCredentialProvider
	Host               string
	Port               string
	BroadcastName      string
	Connected          bool
	Listeners          []RabbitMQListnerInterface
	CancelChannel      chan *amqp.Error
	TracerID           string
	SenderQueName      string
}

func NewRabbitMQConnection(cp RabbitMQCredentialProvider, host string, port string, broadcastName string) RabbitMQConnection {
	rc := rabbitmqcon{
		Credentials:   cp,
		Host:          host,
		Port:          port,
		BroadcastName: broadcastName,
		Connected:     false,
		Context:       context.Background(),
		TracerID:      "ror-unset-tracer-id",
		SenderQueName: "ror",
	}
	rc.connect()
	return &rc
}

func (rc *rabbitmqcon) SetTracerID(tracerID string) {
	rc.TracerID = tracerID
}
func (rc *rabbitmqcon) SetSenderQue(quename string) {
	rc.SenderQueName = quename
}
func (rc *rabbitmqcon) Trace(ctx context.Context, spanname string) (context.Context, trace.Span) {
	if rc.TracerID != "ror-unset-tracer-id" {
		return otel.GetTracerProvider().Tracer(rc.TracerID).Start(ctx, spanname)
	}
	return noop.NewTracerProvider().Tracer("noop").Start(ctx, spanname)
}
func (rc *rabbitmqcon) RegisterHandler(listner RabbitMQListnerInterface) error {
	rc.Listeners = append(rc.Listeners, listner)
	if rc.Connected {
		go listner.Listen(rc.CancelChannel)
	}
	return nil
}

func (rc rabbitmqcon) GetChannel() *amqp.Channel {
	if !rc.Connected {
		rc.connect()
	}
	return rc.RabbitMqChannel
}

// CheckHealth checks the health of the rabbitmq connection and returns a health check
func (rc *rabbitmqcon) CheckHealth() []health.Check {
	c := health.Check{}
	if !rc.Ping() {
		c.Status = health.StatusFail
		c.Output = "Could not ping rabbitmq"
	}
	return []health.Check{c}
}

func (rc *rabbitmqcon) Ping() bool {
	return rc.Connected
}

func (rc rabbitmqcon) getConnectionstring() string {
	username, password := rc.Credentials.GetCredentials()
	return fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, rc.Host, rc.Port)
}

func (rc *rabbitmqcon) connect() {
	rlog.Debug("Connecting", rlog.String("rabbitmq", rc.getConnectionstring()))
	c := make(chan *amqp.Error)
	rc.CancelChannel = c
	go func(rc *rabbitmqcon) {
		err := <-rc.CancelChannel
		rlog.Error("reconnect", err)
		rc.Connected = false
		rc.connect()
	}(rc)
	connection, err := amqp.Dial(rc.getConnectionstring())
	if err != nil {
		rlog.Fatal("cannot connect to rabbitmq", err)
	}

	rc.RabbitMqConnection = connection
	connection.NotifyClose(c)

	failOnError(err, "Failed to connect to RabbitMQ")

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	rc.RabbitMqChannel = channel
	rc.Connected = true
	rlog.Info("connected to RabbitMQ", rlog.Any("Conected", rc.Connected))

	for _, listner := range rc.Listeners {
		go listner.Listen(c)
	}
}
