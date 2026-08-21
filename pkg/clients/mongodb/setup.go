package mongodb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/credshelper"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/v2/mongo/otelmongo"
)

const (
	// mongoInitialBackoff is the wait time before the first connection retry.
	mongoInitialBackoff = 1 * time.Second
	// mongoMaxBackoff caps the exponential backoff between connection retries.
	mongoMaxBackoff = 30 * time.Second
	// mongoMonitorInterval is how often the background monitor pings mongodb.
	mongoMonitorInterval = 15 * time.Second
	// mongoPingTimeout bounds a single background ping.
	mongoPingTimeout = 5 * time.Second
	// mongoDisconnectTimeout bounds disconnecting a superseded client.
	mongoDisconnectTimeout = 10 * time.Second
)

var (
	mongodb MongodbCon
	// connMu guards the mongodb client lifecycle (connect/reconnect/swap) so
	// concurrent callers cannot race on the *mongo.Client during a rebuild.
	connMu sync.Mutex
	// monitorOnce ensures a single background connection monitor is started.
	monitorOnce sync.Once
)

// This type implements the mongodb connection in ror
type MongodbCon struct {
	Client      *mongo.Client
	Context     context.Context
	Credentials credshelper.CredHelperWithRenew
	Host        string
	Port        string
	Database    string
}

// GetMongoDb function returns a pointer to the `mongo.Database` instance used to communicate with MongoDB server.
// The function simply returns the MongoDB client instance stored in a `mongodb` singleton object.
// This function is used to obtain the MongoDB client connection in other parts of the application.
func GetMongoDb() *mongo.Database {
	return getDbConnectionWithReconnect().Database(mongodb.Database)
}

func GetMongoClient() *mongo.Client {
	return getDbConnectionWithReconnect()
}

// Init initializes the mongodb client, retrying the connection with exponential
// backoff until it succeeds. It blocks until a connection is established.
//
// Kept for backwards compatibility; prefer InitWithContext or MustInitWithContext
// so the connection attempt can be bounded by a context.
func Init(cp credshelper.CredHelperWithRenew, host string, port string, database string) {
	_ = InitWithContext(context.Background(), cp, host, port, database)
}

// InitWithContext initializes the mongodb client, retrying the connection with
// exponential backoff until it succeeds or the context is cancelled. A health
// checker is registered up front so the health endpoint reports mongodb as
// unhealthy ("Connecting to mongodb") while the retry loop runs, instead of the
// dependency being invisible until it finally connects.
//
// It returns nil on success, or the context error if the context is cancelled
// before a connection is established.
func InitWithContext(ctx context.Context, cp credshelper.CredHelperWithRenew, host string, port string, database string) error {
	mongodb.init(cp, host, port, database)

	checker := &mongoStartupChecker{}
	rorhealth.Register(ctx, "mongodb", checker)

	if err := mongodb.connectWithRetry(ctx); err != nil {
		return err
	}
	checker.setConnected()
	startConnectionMonitor(ctx)
	return nil
}

// MustInitWithContext behaves like InitWithContext but treats a cancelled
// context as fatal: it logs the failure and exits the process. Use this when a
// mongodb connection is a hard prerequisite and the process must not continue
// without it.
func MustInitWithContext(ctx context.Context, cp credshelper.CredHelperWithRenew, host string, port string, database string) {
	if err := InitWithContext(ctx, cp, host, port, database); err != nil {
		rlog.Fatal("could not connect to mongodb within timeout, giving up", err,
			rlog.String("host", host),
			rlog.String("port", port))
	}
}

// mongoStartupChecker is a health checker that tracks the mongodb connection
// state while it is being established. Before a connection succeeds it reports
// StatusFail so the health endpoint clearly shows mongodb as the dependency that
// is blocking startup. Once connected it delegates to the live connection's ping
// so the check reflects the real connection state.
type mongoStartupChecker struct {
	mu        sync.RWMutex
	connected bool
}

func (c *mongoStartupChecker) setConnected() {
	c.mu.Lock()
	c.connected = true
	c.mu.Unlock()
}

func (c *mongoStartupChecker) CheckHealth(ctx context.Context) []rorhealth.Check {
	c.mu.RLock()
	connected := c.connected
	c.mu.RUnlock()

	if !connected {
		return []rorhealth.Check{{
			Status: rorhealth.StatusFail,
			Output: "Connecting to mongodb",
		}}
	}
	return mongodb.CheckHealth(ctx)
}

func GetMongodbConnection() *MongodbCon {
	return &mongodb
}

func (rc MongodbCon) GetMongoDb() *mongo.Database {
	mongoClient := getDbConnectionWithReconnect().Database(mongodb.Database)
	return mongoClient
}

// CheckHealthWithoutContext checks the health of the redis connection and returns a health check
func (rc MongodbCon) CheckHealthWithoutContext() []rorhealth.Check {
	c := rorhealth.Check{}
	if !Ping() {
		c.Status = rorhealth.StatusFail
		c.Output = "Could not ping mongodb"
	}
	return []rorhealth.Check{c}
}

func (rc MongodbCon) CheckHealth(ctx context.Context) []rorhealth.Check {
	c := rorhealth.Check{}
	if !PingWithContext(ctx) {
		c.Status = rorhealth.StatusFail
		c.Output = "Could not ping mongodb"
	}
	return []rorhealth.Check{c}
}

// Ping the mongodb database and returns the result as a bool
func Ping() bool {
	return pingCurrent(context.Background())
}

func PingWithContext(ctx context.Context) bool {
	return pingCurrent(ctx)
}

func (mdb MongodbCon) getConnectionstring() string {
	username, password := mdb.Credentials.GetCredentials()
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, mdb.Host, mdb.Port, mdb.Database)
}

func (mdb *MongodbCon) init(cp credshelper.CredHelperWithRenew, host string, port string, database string) {
	mdb.Context = context.Background()
	mdb.Host = host
	mdb.Port = port
	mdb.Database = database
	mdb.Credentials = cp
}

func (mdb MongodbCon) ping(ctx context.Context) bool {
	if mdb.Client == nil {
		rlog.Debug("mongodb client is not initialized")
		return false
	}
	err := mdb.Client.Ping(ctx, nil)
	if err != nil {
		rlog.Debug(err.Error())
		return false
	}
	return true
}

// connect establishes a single connection to mongodb and verifies it with a
// ping. It returns an error instead of exiting so callers can choose their own
// retry or failure policy.
func (mdb *MongodbCon) connect() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		SetMonitor(otelmongo.NewMonitor()).
		ApplyURI(mdb.getConnectionstring()).
		SetServerAPIOptions(serverAPI).
		SetMaxPoolSize(100).
		SetMinPoolSize(10).
		SetBSONOptions(&options.BSONOptions{
			ObjectIDAsHexString: true,
		})

	cli, err := mongo.Connect(opts)
	if err != nil {
		return fmt.Errorf("could not connect to mongodb: %w", err)
	}

	err = cli.Ping(mdb.Context, nil)
	if err != nil {
		// Disconnect the temporary client so a failed attempt does not leak
		// sockets and goroutines when the caller retries.
		_ = cli.Disconnect(mdb.Context)
		return fmt.Errorf("could not ping mongodb: %w", err)
	}

	mdb.Client = cli
	return nil
}

// connectWithRetry connects to mongodb, retrying with exponential backoff until
// it succeeds or the context is cancelled. Each failed attempt is logged with
// the host, port, attempt number and underlying error so failures are easy to
// troubleshoot.
func (mdb *MongodbCon) connectWithRetry(ctx context.Context) error {
	backoff := mongoInitialBackoff
	for attempt := 1; ; attempt++ {
		err := mdb.connect()
		if err == nil {
			if attempt > 1 {
				rlog.Info("connected to mongodb",
					rlog.String("host", mdb.Host),
					rlog.String("port", mdb.Port),
					rlog.Int("attempts", attempt))
			}
			return nil
		}

		rlog.Error("could not connect to mongodb, retrying", err,
			rlog.String("host", mdb.Host),
			rlog.String("port", mdb.Port),
			rlog.Int("attempt", attempt),
			rlog.String("retryIn", backoff.String()))

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(backoff):
		}

		backoff *= 2
		if backoff > mongoMaxBackoff {
			backoff = mongoMaxBackoff
		}
	}
}

// currentClient returns the active client under the connection lock so reads do
// not race with a concurrent reconnect.
func currentClient() *mongo.Client {
	connMu.Lock()
	defer connMu.Unlock()
	return mongodb.Client
}

// pingCurrent pings the active client, honoring the supplied context.
func pingCurrent(ctx context.Context) bool {
	cli := currentClient()
	if cli == nil {
		rlog.Debug("mongodb client is not initialized")
		return false
	}
	if err := cli.Ping(ctx, nil); err != nil {
		rlog.Debug(err.Error())
		return false
	}
	return true
}

// getDbConnectionWithReconnect returns a live client, rebuilding it on scheduled
// Vault credential rotation. It is safe for concurrent use. On a rebuild failure
// it keeps the existing client so the background monitor can retry instead of
// crashing the process or returning a nil client.
func getDbConnectionWithReconnect() *mongo.Client {
	connMu.Lock()
	defer connMu.Unlock()

	if mongodb.Client == nil {
		// Reached only if queries run before Init succeeded; treat as a fatal
		// misconfiguration to preserve the previous hard-prerequisite contract.
		if err := mongodb.connect(); err != nil {
			rlog.Fatal("could not connect to Mongodb", err)
		}
		return mongodb.Client
	}

	if mongodb.Credentials.CheckAndRenew() {
		reconnectLocked("credential rotation")
	}
	return mongodb.Client
}

// reconnectLocked rebuilds the client. connMu must be held. The existing client
// is only replaced and disconnected once a new connection is established, so a
// failed rebuild leaves the previous client in place for the next retry.
func reconnectLocked(reason string) {
	rlog.Info("reconnecting to mongodb", rlog.String("reason", reason))
	old := mongodb.Client
	if err := mongodb.connect(); err != nil {
		rlog.Error("could not reconnect to mongodb, keeping existing connection", err,
			rlog.String("reason", reason))
		return
	}
	if old != nil {
		// Disconnect the superseded client asynchronously with a bounded context
		// so a stalled Disconnect cannot hold connMu and block all mongodb callers.
		go func(c *mongo.Client) {
			ctx, cancel := context.WithTimeout(context.Background(), mongoDisconnectTimeout)
			defer cancel()
			_ = c.Disconnect(ctx)
		}(old)
	}
}

// forceReconnect fetches fresh credentials (bypassing the local expiry timer)
// and rebuilds the connection. Used when an actual failure is detected, e.g. the
// dynamic database user was revoked server-side before its local lease expired.
func forceReconnect(reason string) {
	connMu.Lock()
	defer connMu.Unlock()

	if fr, ok := mongodb.Credentials.(credshelper.ForceRenewer); ok {
		if err := fr.ForceRenew(); err != nil {
			rlog.Error("mongodb: forced credential renewal failed", err, rlog.String("reason", reason))
		}
	} else {
		mongodb.Credentials.CheckAndRenew()
	}
	reconnectLocked(reason)
}

// startConnectionMonitor launches a single background goroutine that pings
// mongodb periodically and rebuilds the connection when it becomes unusable.
// This decouples recovery from incoming query traffic, which stops once a pod is
// marked NotReady, so a wedged connection can heal on its own.
func startConnectionMonitor(ctx context.Context) {
	monitorOnce.Do(func() {
		// Callers often pass a context that only bounds the initial connect (with
		// defer cancel()); using it directly would stop the monitor right after
		// Init. Keep the context values (e.g. tracing) but ignore its cancellation
		// so the monitor lives for the process lifetime.
		monitorCtx := context.WithoutCancel(ctx)
		go connectionMonitor(monitorCtx)
	})
}

func connectionMonitor(ctx context.Context) {
	ticker := time.NewTicker(mongoMonitorInterval)
	defer ticker.Stop()

	failures := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pingCtx, cancel := context.WithTimeout(ctx, mongoPingTimeout)
			ok := pingCurrent(pingCtx)
			cancel()
			if ok {
				failures = 0
				continue
			}
			failures++
			rlog.Warn("mongodb ping failed, rebuilding connection",
				rlog.Int("consecutiveFailures", failures))
			forceReconnect("ping failure")
		}
	}
}

// DEPRECATED: This function is wrongly placed, its a local function for ror-api and is implemented
// in the ror-api repo. It should not be used by other services, and is removed from this package.
// it will panic if used, to prevent usage. Please use the implementation in ror-api instead.
// GenerateAggregateQuery generates a MongoDB aggregation query based on the provided ResourceQuery.
func (mdb *MongodbCon) GenerateAggregateQuery(rorResourceQuery *rorresources.ResourceQuery) []bson.M {
	panic("not implemented, use local implementation")
}
