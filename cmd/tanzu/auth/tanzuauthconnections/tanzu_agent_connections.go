package tanzuauthconnections

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/rabbitmqclient"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/rabbitmqcredhelper"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"

	"github.com/dotse/go-health"
	"github.com/spf13/viper"
)

var (
	VaultClient        *vaultclient.VaultClient
	RabbitMQConnection rabbitmqclient.RabbitMQConnection
)

func InitConnections() {
	VaultClient = vaultclient.NewVaultClient(viper.GetString(configconsts.ROLE), viper.GetString(configconsts.VAULT_URL))
	rmqcredhelper := rabbitmqcredhelper.NewVaultRMQCredentials(VaultClient, viper.GetString(configconsts.ROLE))
	RabbitMQConnection = rabbitmqclient.NewRabbitMQConnection(rmqcredhelper, viper.GetString(configconsts.RABBITMQ_HOST), viper.GetString(configconsts.RABBITMQ_PORT), viper.GetString(configconsts.RABBITMQ_BROADCAST_NAME))

	health.Register("vault", VaultClient)
	health.Register("rabbitmq", RabbitMQConnection)
}
