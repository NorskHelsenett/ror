package mongodb

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/dotse/go-health"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var mongodb mongodbcon

type DatabaseCredentialHelper interface {
	GetCredentials() (string, string)
	CheckAndRenew() bool
}

// This type implements the mongodb connection in ror
type mongodbcon struct {
	Client      *mongo.Client
	Context     context.Context
	Credentials DatabaseCredentialHelper
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

// Initializes the mongodb client
func Init(cp DatabaseCredentialHelper, host string, port string, database string) {
	mongodb.init(cp, host, port, database)
	health.Register("mongodb", mongodb)
}

// CheckHealth checks the health of the redis connection and returns a health check
func (rc mongodbcon) CheckHealth() []health.Check {
	c := health.Check{}
	if !Ping() {
		c.Status = health.StatusFail
		c.Output = "Could not ping mongodb"
	}
	return []health.Check{c}
}

// Ping the mongodb database and returns the result as a bool
func Ping() bool {
	return mongodb.ping()
}

func (mdb mongodbcon) getConnectionstring() string {
	username, password := mdb.Credentials.GetCredentials()
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, mdb.Host, mdb.Port, mdb.Database)
}

func (mdb *mongodbcon) init(cp DatabaseCredentialHelper, host string, port string, database string) {
	mdb.Context = context.Background()
	mdb.Host = host
	mdb.Port = port
	mdb.Database = database
	mdb.Credentials = cp
	mdb.connect()
}

func (mdb mongodbcon) ping() bool {
	err := mdb.Client.Ping(context.Background(), nil)
	if err != nil {
		rlog.Debug(err.Error())
		return false
	}
	return true
}

func (mdb *mongodbcon) connect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().SetMonitor(otelmongo.NewMonitor()).ApplyURI(mdb.getConnectionstring()).SetServerAPIOptions(serverAPI)

	cli, err := mongo.Connect(mdb.Context, opts)
	if err != nil {
		rlog.Fatal("could not connect to Mongodb", err)
	}

	err = cli.Ping(mdb.Context, nil)
	if err != nil {
		rlog.Fatal("could not ping Mongodb", err)
	}

	mdb.Client = cli
}

func (mdb *mongodbcon) disconnect() {
	_ = mdb.Client.Disconnect(mdb.Context)
}

func (mdb *mongodbcon) getDbConnectionWithReconnect() *mongo.Client {
	if mdb.Client == nil {
		mdb.connect()
	}

	if mdb.Credentials.CheckAndRenew() {
		rlog.Info("reconnecting to mongodb")
		mdb.disconnect()
		mdb.connect()
	}

	return mdb.Client
}
