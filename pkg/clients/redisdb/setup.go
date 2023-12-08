package redisdb

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/dotse/go-health"
	"github.com/redis/go-redis/extra/redisotel/v9"
	goredis "github.com/redis/go-redis/v9"
)

type DatabaseCredentialHelper interface {
	GetCredentials() (string, string)
	CheckAndRenew() bool
}

type RedisDB interface {
	Get(ctx context.Context, key string, output *string) error
	Set(ctx context.Context, key string, value interface{}) error
	GetJSON(context.Context, string, string, interface{}) error
	SetJSON(ctx context.Context, key string, path string, value interface{}) error
	clients.CommonClient
}

var redisdb rediscon

// This type implements the redis connection in ror

type rediscon struct {
	Client      *goredis.Client
	Credentials DatabaseCredentialHelper
	Host        string
	Port        string
}

func New(cp DatabaseCredentialHelper, host string, port string) RedisDB {
	rc := rediscon{
		Credentials: cp,
		Host:        host,
		Port:        port,
	}

	rc.connect()

	return rc
}

// GetRedisClient function returns a pointer to the `goredis.Client` instance used to communicate with Redis server.
// The function simply returns the Redis client instance stored in a `redisdb` singleton object.
// This function is used to obtain the Redis client connection in other parts of the application.
func GetRedisClient() *goredis.Client {
	return redisdb.Client
}

// CheckHealth checks the health of the redis connection and returns a health check
func (rc rediscon) CheckHealth() []health.Check {
	c := health.Check{}
	if !rc.Ping() {
		c.Status = health.StatusFail
		c.Output = "Could not ping redis"
	}
	return []health.Check{c}
}

// Ping the redis connection
func Ping() bool {
	return redisdb.Ping()
}

func (rc rediscon) Ping() bool {
	_, err := rc.Client.Ping(context.Background()).Result()
	return err == nil
}

func (rc *rediscon) connect() {
	cli := goredis.NewClient(&goredis.Options{Addr: rc.getAddr(), CredentialsProvider: rc.Credentials.GetCredentials})
	rc.Client = cli
	if !rc.Ping() {
		rlog.Error("could not connect to redis", nil)
	}
	_ = redisotel.InstrumentMetrics(cli)
	_ = redisotel.InstrumentTracing(cli)
}

func (rc rediscon) getAddr() string {
	return fmt.Sprintf("%s:%s", rc.Host, rc.Port)
}
