package mongodb

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/helpers/credshelper"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
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

// Initializes the mongodb client
func Init(cp credshelper.CredHelperWithRenew, host string, port string, database string) {
	mongodb.init(cp, host, port, database)
	rorhealth.Register(context.TODO(), "mongodb", mongodb)
}

func GetMongodbConnection() *MongodbCon {
	return &mongodb
}

func (rc MongodbCon) GetMongoDb() *mongo.Database {
	mongoClient := rc.getDbConnectionWithReconnect().Database(mongodb.Database)
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
	mdb.connect()
}

func (mdb MongodbCon) ping(ctx context.Context) bool {
	err := mdb.Client.Ping(ctx, nil)
	if err != nil {
		rlog.Debug(err.Error())
		return false
	}
	return true
}

func (mdb *MongodbCon) connect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().SetMonitor(otelmongo.NewMonitor()).ApplyURI(mdb.getConnectionstring()).SetServerAPIOptions(serverAPI).SetMaxPoolSize(100).SetMinPoolSize(10)

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

func (mdb *MongodbCon) disconnect() {
	_ = mdb.Client.Disconnect(mdb.Context)
}

func (mdb *MongodbCon) getDbConnectionWithReconnect() *mongo.Client {
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
