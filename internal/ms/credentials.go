package ms

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

type rabbitMQCredentials struct {
	Port     int16
	Host     string
	Username string
	Password string
}

func (r rabbitMQCredentials) String() string {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d", r.Username, r.Password, r.Host, r.Port)
	return connectionString
}

// GetCredsFromVault NOTE: Should be deprecated when all microservices are ported to the new microservice api
func GetCredsFromVault(role string, vaultClient *vaultclient.VaultClient) {
	GetRabbitMQCreds(role, vaultClient)
}

// GetRabbitMQCreds NOTE: Should be deprecated when all microservices are ported to the new microservice api
func GetRabbitMQCreds(role string, vaultClient *vaultclient.VaultClient) {
	secretPath := fmt.Sprintf("rabbitmq/creds/%s", role)
	//TODO: Replave with method from vaultclient
	rabbitMqCreds, errRabbitMq := vaultClient.GetSecret(secretPath)
	if errRabbitMq != nil {
		return
	}

	rabbitMQUser := rabbitMqCreds["username"].(string)
	rabbitMQPwd := rabbitMqCreds["password"].(string)
	rabbitmqHost := viper.GetString(configconsts.RABBITMQ_HOST)
	rabbitmqPort := viper.GetString(configconsts.RABBITMQ_PORT)
	rabbitMQConnectionstring := fmt.Sprintf("amqp://%s:%s@%s:%s", rabbitMQUser, rabbitMQPwd, rabbitmqHost, rabbitmqPort)
	viper.Set(configconsts.RABBITMQ_CONNECTIONSTRING, rabbitMQConnectionstring)
}

func GetMongodbCreds(role string, vaultClient *vaultclient.VaultClient) {
	secretPath := fmt.Sprintf("mongodb/creds/%s", role)
	//TODO: Replave with method from vaultclient
	mongoCreds, errMongo := vaultClient.GetSecret(secretPath)
	if errMongo != nil {
		rlog.Errorc(context.Background(), errMongo.Error(), nil)
		return
	}

	mongoDBUser := mongoCreds["username"].(string)
	mongoDBPwd := mongoCreds["password"].(string)
	mongoDbConnectionstring := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		mongoDBUser,
		mongoDBPwd,
		viper.GetString(configconsts.MONGODB_HOST),
		viper.GetString(configconsts.MONGODB_PORT),
		viper.GetString(configconsts.MONGODB_DATABASE))
	viper.Set(configconsts.MONGODB_URL, mongoDbConnectionstring)
}

func newGetRabbitMQCreds(role string, vaultClient *vaultclient.VaultClient) (*rabbitMQCredentials, error) {
	creds := new(rabbitMQCredentials)

	secretPath := fmt.Sprintf("rabbitmq/creds/%s", role)
	//TODO: Replave with method from vaultclient
	vaultResponse, err := vaultClient.GetSecret(secretPath)
	if err != nil {
		return nil, fmt.Errorf("ccould not ask vault for credentials: %v", err)
	}

	creds.Username = vaultResponse["username"].(string)
	creds.Password = vaultResponse["password"].(string)

	creds.Port = int16(viper.GetInt32(configconsts.RABBITMQ_PORT))
	creds.Host = viper.GetString(configconsts.RABBITMQ_HOST)

	return creds, nil
}
