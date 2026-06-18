package redisdb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb/redisdblogadapter"
	"github.com/NorskHelsenett/ror/pkg/helpers/credshelper"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"

	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/redis/go-redis/extra/redisotel/v9"
	goredis "github.com/redis/go-redis/v9"
)

const (
	// redisInitialBackoff is the wait time before the first connection retry.
	redisInitialBackoff = 1 * time.Second
	// redisMaxBackoff caps the exponential backoff between connection retries.
	redisMaxBackoff = 30 * time.Second
)

type RedisDB interface {
	Get(ctx context.Context, key string, output *string) error
	Keys(ctx context.Context) ([]string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context, key string) error
	MGet(ctx context.Context, keys ...string) ([]interface{}, error)
	SetPipelined(ctx context.Context, items []SetItem) error
	clients.CommonClient
}

// SetItem represents a key-value pair with expiration for pipelined SET operations.
type SetItem struct {
	Key        string
	Value      interface{}
	Expiration time.Duration
}

var redisdb rediscon

// This type implements the redis connection in ror

type rediscon struct {
	Client      *goredis.Client
	Credentials credshelper.CredHelper
	Host        string
	Port        string
}

func New(dbc credshelper.CredHelper, host string, port string) RedisDB {
	return NewWithContext(context.Background(), dbc, host, port)
}

// NewWithContext creates a redis connection, retrying with exponential backoff
// until it succeeds or the context is cancelled. A health checker is registered
// up front so the health endpoint reports redis as unhealthy ("Connecting to
// redis") while the retry loop runs, instead of the dependency being invisible
// until it finally connects.
//
// If the context is cancelled before a connection is established it returns a
// non-nil, unconnected connection whose health check reports unhealthy, so
// callers never receive nil.
func NewWithContext(ctx context.Context, dbc credshelper.CredHelper, host string, port string) RedisDB {
	rc := newRediscon(dbc, host, port)

	checker := &redisStartupChecker{conn: rc}
	rorhealth.Register(ctx, "redis", checker)

	if err := rc.connectWithRetry(ctx); err == nil {
		checker.setConnected()
	}
	return rc
}

// MustNewWithContext behaves like NewWithContext but treats a cancelled context
// as fatal: it logs the failure and exits the process. Use this when a redis
// connection is a hard prerequisite and the process must not continue without it.
func MustNewWithContext(ctx context.Context, dbc credshelper.CredHelper, host string, port string) RedisDB {
	rc := newRediscon(dbc, host, port)

	checker := &redisStartupChecker{conn: rc}
	rorhealth.Register(ctx, "redis", checker)

	if err := rc.connectWithRetry(ctx); err != nil {
		rlog.Fatal("could not connect to redis within timeout, giving up", err,
			rlog.String("host", host),
			rlog.String("port", port))
	}
	checker.setConnected()
	return rc
}

// newRediscon builds an unconnected redis connection with the redigo log adapter
// installed. The actual network connection is established later by connect or
// connectWithRetry.
func newRediscon(dbc credshelper.CredHelper, host string, port string) *rediscon {
	goredis.SetLogger(redisdblogadapter.NewRlogAdapter())
	return &rediscon{
		Credentials: dbc,
		Host:        host,
		Port:        port,
	}
}

// redisStartupChecker is a health checker that tracks the redis connection state
// while it is being established. Before a connection succeeds it reports
// StatusFail so the health endpoint clearly shows redis as the dependency that
// is blocking startup. Once connected it delegates to the live connection's ping
// so the check reflects the real connection state.
type redisStartupChecker struct {
	mu        sync.RWMutex
	conn      *rediscon
	connected bool
}

func (c *redisStartupChecker) setConnected() {
	c.mu.Lock()
	c.connected = true
	c.mu.Unlock()
}

func (c *redisStartupChecker) CheckHealth(ctx context.Context) []rorhealth.Check {
	c.mu.RLock()
	connected := c.connected
	c.mu.RUnlock()

	if !connected {
		return []rorhealth.Check{{
			Status: rorhealth.StatusFail,
			Output: "Connecting to redis",
		}}
	}
	return c.conn.CheckHealth(ctx)
}

// NewFromClient creates a RedisDB from an existing go-redis client.
// Useful for testing with miniredis or other pre-configured clients.
func NewFromClient(client *goredis.Client) RedisDB {
	return &rediscon{Client: client}
}

// GetRedis function returns a pointer to the `goredis.Client` instance used to communicate with Redis server.
// The function simply returns the Redis client instance stored in a `redisdb` singleton object.
// This function is used to obtain the Redis client connection in other parts of the application.
func GetRedis() *goredis.Client {
	return redisdb.Client
}

// CheckHealth checks the health of the redis connection and returns a health check
func (rc rediscon) CheckHealth(_ context.Context) []rorhealth.Check {
	return rc.CheckHealthWithoutContext()
}

// CheckHealth checks the health of the redis connection and returns a health check
func (rc rediscon) CheckHealthWithoutContext() []rorhealth.Check {
	c := rorhealth.Check{}
	if !rc.Ping() {
		c.Status = rorhealth.StatusFail
		c.Output = "Could not ping redis"
	}
	return []rorhealth.Check{c}
}

// Ping the redis connection
func Ping() bool {
	return redisdb.Ping()
}

func (rc rediscon) Ping() bool {
	if rc.Client == nil {
		return false
	}
	_, err := rc.Client.Ping(context.Background()).Result()
	return err == nil
}

func (rc rediscon) PingWithContext(ctx context.Context) bool {
	if rc.Client == nil {
		return false
	}
	_, err := rc.Client.Ping(ctx).Result()
	return err == nil
}

// connect establishes a single connection to redis and verifies it with a ping.
// It returns an error instead of logging so callers can choose their own retry
// or failure policy.
func (rc *rediscon) connect() error {
	cli := goredis.NewClient(&goredis.Options{Addr: rc.getAddr(), CredentialsProvider: rc.Credentials.GetCredentials})
	rc.Client = cli
	if !rc.Ping() {
		// Close and clear the client so a failed attempt does not leak the
		// connection pool when the caller retries.
		_ = cli.Close()
		rc.Client = nil
		return fmt.Errorf("could not ping redis at %s", rc.getAddr())
	}
	_ = redisotel.InstrumentMetrics(cli)
	_ = redisotel.InstrumentTracing(cli)
	return nil
}

// connectWithRetry connects to redis, retrying with exponential backoff until it
// succeeds or the context is cancelled. Each failed attempt is logged with the
// host, port, attempt number and underlying error so failures are easy to
// troubleshoot.
func (rc *rediscon) connectWithRetry(ctx context.Context) error {
	backoff := redisInitialBackoff
	for attempt := 1; ; attempt++ {
		err := rc.connect()
		if err == nil {
			if attempt > 1 {
				rlog.Info("connected to redis",
					rlog.String("host", rc.Host),
					rlog.String("port", rc.Port),
					rlog.Int("attempts", attempt))
			}
			return nil
		}

		rlog.Error("could not connect to redis, retrying", err,
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
		if backoff > redisMaxBackoff {
			backoff = redisMaxBackoff
		}
	}
}

func (rc rediscon) getAddr() string {
	return fmt.Sprintf("%s:%s", rc.Host, rc.Port)
}
