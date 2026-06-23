package rabbitmqclient

import (
	"context"
	"fmt"
	"sync"
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

const (
	// rabbitmqInitialBackoff is the wait time before the first connection retry.
	rabbitmqInitialBackoff = 1 * time.Second
	// rabbitmqMaxBackoff caps the exponential backoff between connection retries.
	rabbitmqMaxBackoff = 30 * time.Second
)

type RabbitMQListnerInterface interface {
	Listen(chan *amqp.Error)
	ListenWithTTL(chan *amqp.Error, time.Duration)
}

type RabbitMQConnection interface {
	GetChannel() *amqp.Channel
	RegisterHandler(RabbitMQListnerInterface) error
	RegisterHandlerWithTTL(RabbitMQListnerInterface, time.Duration) error
	SendMessage(ctx context.Context, message any, routing string, extraheaders map[string]any) error
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

// NewRabbitMQConnection creates a rabbitmq connection, retrying with
// exponential backoff until it succeeds. It blocks until connected.
//
// Prefer NewRabbitMQConnectionWithContext or MustNewRabbitMQConnectionWithContext
// so the connection attempt can be bounded by a context.
func NewRabbitMQConnection(cp credshelper.CredHelper, host string, port string, broadcastName string) RabbitMQConnection {
	return NewRabbitMQConnectionWithContext(context.Background(), cp, host, port, broadcastName)
}

// NewRabbitMQConnectionWithContext creates a rabbitmq connection, retrying with
// exponential backoff until it succeeds or the context is cancelled. A health
// checker is registered up front so the health endpoint reports rabbitmq as
// unhealthy ("Connecting to rabbitmq") while the retry loop runs, instead of the
// dependency being invisible until it finally connects.
//
// If the context is cancelled before a connection is established it returns a
// non-nil, unconnected connection whose health check reports unhealthy, so
// callers never receive nil.
func NewRabbitMQConnectionWithContext(ctx context.Context, cp credshelper.CredHelper, host string, port string, broadcastName string) RabbitMQConnection {
	rc := getDefaultRabbitMQConnectionConfig()
	rc.applyOptions(
		OptionCredentialsProvider(cp),
		OptionHost(host),
		OptionPort(port),
		OptionBroadcastName(broadcastName),
	)

	checker := &rabbitmqStartupChecker{conn: rc}
	rorhealth.Register(ctx, "rabbitmq", checker)

	if err := rc.connectWithRetry(ctx); err == nil {
		checker.setConnected()
	}
	return rc
}

// MustNewRabbitMQConnectionWithContext behaves like
// NewRabbitMQConnectionWithContext but treats a cancelled context as fatal: it
// logs the failure and exits the process. Use this when a rabbitmq connection is
// a hard prerequisite and the process must not continue without it.
func MustNewRabbitMQConnectionWithContext(ctx context.Context, cp credshelper.CredHelper, host string, port string, broadcastName string) RabbitMQConnection {
	rc := getDefaultRabbitMQConnectionConfig()
	rc.applyOptions(
		OptionCredentialsProvider(cp),
		OptionHost(host),
		OptionPort(port),
		OptionBroadcastName(broadcastName),
	)

	checker := &rabbitmqStartupChecker{conn: rc}
	rorhealth.Register(ctx, "rabbitmq", checker)

	if err := rc.connectWithRetry(ctx); err != nil {
		rlog.Fatal("could not connect to rabbitmq within timeout, giving up", err,
			rlog.String("host", host),
			rlog.String("port", port))
	}
	checker.setConnected()
	return rc
}

func NewRabbitMQConnectionWithOptions(cfg ...RabbitMQConnectionOption) RabbitMQConnection {
	rc := getDefaultRabbitMQConnectionConfig()
	rc.applyOptions(cfg...)
	if err := rc.connect(); err != nil {
		rlog.Fatal("could not connect to rabbitmq", err)
	}
	return rc
}

func NewRabbitMQConnectionWithDefaults(cfg ...RabbitMQConnectionOption) RabbitMQConnection {
	rc := getDefaultRabbitMQConnectionConfig()
	rc.loadDefaultConfig()
	rc.applyOptions(cfg...)
	if err := rc.connect(); err != nil {
		rlog.Fatal("could not connect to rabbitmq", err)
	}
	return rc
}

// rabbitmqStartupChecker is a health checker that tracks the rabbitmq connection
// state while it is being established. Before a connection succeeds it reports
// StatusFail so the health endpoint clearly shows rabbitmq as the dependency
// that is blocking startup. Once connected it delegates to the live connection's
// ping so the check reflects the real connection state.
type rabbitmqStartupChecker struct {
	mu        sync.RWMutex
	conn      *rabbitmqcon
	connected bool
}

func (c *rabbitmqStartupChecker) setConnected() {
	c.mu.Lock()
	c.connected = true
	c.mu.Unlock()
}

func (c *rabbitmqStartupChecker) CheckHealth(ctx context.Context) []rorhealth.Check {
	c.mu.RLock()
	connected := c.connected
	c.mu.RUnlock()

	if !connected {
		return []rorhealth.Check{{
			Status: rorhealth.StatusFail,
			Output: "Connecting to rabbitmq",
		}}
	}
	return c.conn.CheckHealth(ctx)
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
		if err := rc.connect(); err != nil {
			rlog.Fatal("cannot connect to rabbitmq", err)
		}
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

// connect establishes a single connection to rabbitmq, opens a channel and
// starts the reconnect watcher. It returns an error instead of exiting so
// callers can choose their own retry or failure policy. The reconnect watcher
// spawned on success keeps the previous fatal-on-failure behaviour for runtime
// reconnects.
func (rc *rabbitmqcon) connect() error {
	err := rc.validateConfig()
	if err != nil {
		return fmt.Errorf("invalid rabbitmq configuration: %w", err)
	}
	rlog.Debug("Connecting", rlog.String("rabbitmq", rc.getConnectionstringLog()))

	connection, err := amqp.Dial(rc.getConnectionstring())
	if err != nil {
		return fmt.Errorf("cannot connect to rabbitmq: %w", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		_ = connection.Close()
		return fmt.Errorf("failed to open a rabbitmq channel: %w", err)
	}

	c := make(chan *amqp.Error)
	rc.CancelChannel = c
	rc.RabbitMqConnection = connection
	connection.NotifyClose(c)
	rc.RabbitMqChannel = channel
	rc.Connected = true
	rlog.Info("connected to RabbitMQ", rlog.Any("Connected", rc.Connected))

	// Runtime reconnect: on connection loss, retry once and keep the previous
	// fatal-on-failure behaviour for the steady-state path.
	go func(rc *rabbitmqcon) {
		err := <-rc.CancelChannel
		rlog.Error("reconnect", err)
		rc.Connected = false
		if rerr := rc.connect(); rerr != nil {
			rlog.Fatal("could not reconnect to rabbitmq", rerr)
		}
	}(rc)

	for _, listner := range rc.Listeners {
		go listner.Listen(c)
	}
	return nil
}

// connectWithRetry connects to rabbitmq, retrying with exponential backoff until
// it succeeds or the context is cancelled. Each failed attempt is logged with
// the host, port, attempt number and underlying error so failures are easy to
// troubleshoot.
func (rc *rabbitmqcon) connectWithRetry(ctx context.Context) error {
	backoff := rabbitmqInitialBackoff
	for attempt := 1; ; attempt++ {
		err := rc.connect()
		if err == nil {
			if attempt > 1 {
				rlog.Info("connected to rabbitmq",
					rlog.String("host", rc.Host),
					rlog.String("port", rc.Port),
					rlog.Int("attempts", attempt))
			}
			return nil
		}

		rlog.Error("could not connect to rabbitmq, retrying", err,
			rlog.String("host", rc.Host),
			rlog.String("port", rc.Port),
			rlog.Int("attempt", attempt),
			rlog.String("retryIn", backoff.String()))

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}

		backoff *= 2
		if backoff > rabbitmqMaxBackoff {
			backoff = rabbitmqMaxBackoff
		}
	}
}
