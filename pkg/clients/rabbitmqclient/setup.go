package rabbitmqclient

import (
	"context"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"
	"github.com/NorskHelsenett/ror/pkg/helpers/credshelper"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQListnerInterface interface {
	Listen(chan *amqp.Error)
	ListenWithTTL(chan *amqp.Error, time.Duration)
}

type RabbitMQConnection interface {
	GetChannel() *amqp.Channel
	RegisterHandler(RabbitMQListnerInterface) error
	RegisterHandlerWithTTL(RabbitMQListnerInterface, time.Duration) error
	SendMessage(ctx context.Context, message any, routing string, extraheaders map[string]interface{}) error
	clients.CommonClient
}

type rabbitmqcon struct {
	Context            context.Context
	RabbitMqConnection *amqp.Connection
	RabbitMqChannel    *amqp.Channel
	Credentials        credshelper.CredHelper
	Host               string
	Port               string
	BroadcastName      string
	Connected          bool
	Listeners          []RabbitMQListnerInterface
	CancelChannel      chan *amqp.Error
	TracerID           string
	SenderQueName      string
}

func NewRabbitMQConnection(cp credshelper.CredHelper, host string, port string, broadcastName string) RabbitMQConnection {
	rc := getDefaultRabbitMQConnectionConfig()
	options := []RabbitMQConnectionOption{
		OptionCredentialsProvider(cp),
		OptionHost(host),
		OptionPort(port),
		OptionBroadcastName(broadcastName),
	}
	rc.applyOptions(options...)
	rc.connect()
	return rc
}

func NewRabbitMQConnectionWithOptions(cfg ...RabbitMQConnectionOption) RabbitMQConnection {
	rc := getDefaultRabbitMQConnectionConfig()
	rc.applyOptions(cfg...)
	rc.connect()
	return rc
}

func NewRabbitMQConnectionWithDefaults(cfg ...RabbitMQConnectionOption) RabbitMQConnection {
	rc := getDefaultRabbitMQConnectionConfig()
	rc.loadDefaultConfig()
	rc.applyOptions(cfg...)
	rc.connect()
	return rc
}

func getDefaultRabbitMQConnectionConfig() *rabbitmqcon {
	rc := &rabbitmqcon{}
	rc.applyDefaults()
	return rc
}

func (rc *rabbitmqcon) applyDefaults() {
	rc.Context = context.Background()
	rc.TracerID = "ror-unset-tracer-id"
	rc.SenderQueName = "ror"
}
func (rc *rabbitmqcon) loadDefaultConfig() {
	if rorconfig.GetString(rorconfig.RABBITMQ_HOST) != "" {
		rc.Host = rorconfig.GetString(rorconfig.RABBITMQ_HOST)
	}
	if rorconfig.GetString(rorconfig.RABBITMQ_PORT) != "" {
		rc.Port = rorconfig.GetString(rorconfig.RABBITMQ_PORT)
	}
	if rorconfig.GetString(rorconfig.RABBITMQ_BROADCAST_NAME) != "" {
		rc.BroadcastName = rorconfig.GetString(rorconfig.RABBITMQ_BROADCAST_NAME)
	}
}
func (rc *rabbitmqcon) applyOptions(options ...RabbitMQConnectionOption) {
	for _, opt := range options {
		opt.apply(rc)
	}
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

// RegisterHandlerWithTTL Convience method for registering a handler with a defined message TTL
func (rc *rabbitmqcon) RegisterHandlerWithTTL(listener RabbitMQListnerInterface, TTL time.Duration) error {
	rc.Listeners = append(rc.Listeners, listener)
	if rc.Connected {
		go listener.ListenWithTTL(rc.CancelChannel, TTL)
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
func (rc *rabbitmqcon) CheckHealth(_ context.Context) []rorhealth.Check {
	return rc.CheckHealthWithoutContext()
}

func (rc *rabbitmqcon) CheckHealthWithoutContext() []rorhealth.Check {
	c := rorhealth.Check{}
	if !rc.Ping() {
		c.Status = rorhealth.StatusFail
		c.Output = "Could not ping rabbitmq"
	}
	return []rorhealth.Check{c}
}

func (rc *rabbitmqcon) Ping() bool {
	return rc.Connected
}

func (rc *rabbitmqcon) PingWithContext(_ context.Context) bool {
	return rc.Connected
}

func (rc rabbitmqcon) getConnectionstring() string {
	username, password := rc.Credentials.GetCredentials()
	return fmt.Sprintf("amqp://%s:%s@%s:%s", username, password, rc.Host, rc.Port)
}

func (rc rabbitmqcon) getConnectionstringLog() string {
	username, _ := rc.Credentials.GetCredentials()
	return fmt.Sprintf("amqp://%s:%s@%s:%s", username, "******", rc.Host, rc.Port)
}

func (rc *rabbitmqcon) validateConfig() error {
	if rc.Host == "" || rc.Port == "" {
		return fmt.Errorf("invalid rabbitmq configuration: host=%q port=%q", rc.Host, rc.Port)
	}
	if rc.Credentials == nil {
		return fmt.Errorf("invalid rabbitmq configuration: credentials are required")
	}

	return nil
}

func (rc *rabbitmqcon) connect() {
	err := rc.validateConfig()
	if err != nil {
		rlog.Fatal("invalid rabbitmq configuration", err)
	}
	rlog.Debug("Connecting", rlog.String("rabbitmq", rc.getConnectionstringLog()))
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
