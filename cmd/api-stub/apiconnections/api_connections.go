package apiconnections

import (
	"encoding/json"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/auth/userauth"
	"github.com/NorskHelsenett/ror/pkg/clients/rabbitmqclient"

	"github.com/NorskHelsenett/ror/pkg/clients/redisdb"
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/databasecredhelper"
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/rabbitmqcredhelper"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	newhealth "github.com/dotse/go-health"
	"github.com/spf13/viper"
)

var (
	VaultClient        *vaultclient.VaultClient
	RedisDB            redisdb.RedisDB
	RabbitMQConnection rabbitmqclient.RabbitMQConnection
	DomainResolvers    *userauth.DomainResolvers
)

func InitConnections() {
	VaultClient = vaultclient.NewVaultClient(viper.GetString(configconsts.ROLE), viper.GetString(configconsts.VAULT_URL))

	redisdatabasecredhelper := databasecredhelper.NewVaultDBCredentials(VaultClient, fmt.Sprintf("redis-%v-role", viper.GetString(configconsts.ROLE)), "")
	RedisDB = redisdb.New(redisdatabasecredhelper, viper.GetString(configconsts.REDIS_HOST), viper.GetString(configconsts.REDIS_PORT))

	rmqcredhelper := rabbitmqcredhelper.NewVaultRMQCredentials(VaultClient, viper.GetString(configconsts.ROLE))
	RabbitMQConnection = rabbitmqclient.NewRabbitMQConnection(rmqcredhelper, viper.GetString(configconsts.RABBITMQ_HOST), viper.GetString(configconsts.RABBITMQ_PORT), viper.GetString(configconsts.RABBITMQ_BROADCAST_NAME))

	var err error
	DomainResolvers, err = LoadDomainResolvers()
	if err != nil {
		rlog.Error("Failed to load domain resolvers", err)
	}
	DomainResolvers.RegisterHealthChecks()
	newhealth.Register("vault", VaultClient)
	newhealth.Register("redis", RedisDB)
	newhealth.Register("rabbitmq", RabbitMQConnection)
}

func LoadDomainResolvers() (*userauth.DomainResolvers, error) {
	vaultconfig, err := VaultClient.GetSecret("secret/data/v1.0/ror/config/auth")
	if err != nil {
		rorerror := rorerror.NewRorError(500, "error getting domain resolvers config from secret provider", err)
		return nil, rorerror
	}
	data := vaultconfig["data"]
	drconfig, err := json.Marshal(data)
	if err != nil {
		rorerror := rorerror.NewRorError(500, "error marshaling secret value to json", err)
		return nil, rorerror
	}

	return userauth.NewDomainResolversFromJson(drconfig)
}
