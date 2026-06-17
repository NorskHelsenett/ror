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
)

var mongodb MongodbCon

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
	mongoClient := mongodb.getDbConnectionWithReconnect().Database(mongodb.Database)
	return mongoClient
}

func GetMongoClient() *mongo.Client {
	mongoClient := mongodb.getDbConnectionWithReconnect()
	return mongoClient
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
	mongoClient := rc.getDbConnectionWithReconnect().Database(rc.Database)
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
	return mongodb.ping(context.Background())
}
func PingWithContext(_ context.Context) bool {
	return Ping()
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

func (mdb *MongodbCon) disconnect() {
	_ = mdb.Client.Disconnect(mdb.Context)
}

func (mdb *MongodbCon) getDbConnectionWithReconnect() *mongo.Client {
	if mdb.Client == nil {
		if err := mdb.connect(); err != nil {
			rlog.Fatal("could not connect to Mongodb", err)
		}
	}

	if mdb.Credentials.CheckAndRenew() {
		rlog.Info("reconnecting to mongodb")
		mdb.disconnect()
		if err := mdb.connect(); err != nil {
			rlog.Fatal("could not reconnect to Mongodb", err)
		}
	}

	return mdb.Client
}

// DEPRECATED: This function is wrongly placed, its a local function for ror-api and is implemented
// in the ror-api repo. It should not be used by other services, and is removed from this package.
// it will panic if used, to prevent usage. Please use the implementation in ror-api instead.
// GenerateAggregateQuery generates a MongoDB aggregation query based on the provided ResourceQuery.
func (mdb *MongodbCon) GenerateAggregateQuery(rorResourceQuery *rorresources.ResourceQuery) []bson.M {
	panic("not implemented, use local implementation")
}
