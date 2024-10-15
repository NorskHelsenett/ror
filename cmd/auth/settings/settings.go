package settings

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
)

var (
	VaultSecret *vault.Secret
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault(rlog.LOG_LEVEL, "INFO")
	viper.SetDefault(configconsts.DEX_VAULT_PATH, "secret/data/v1.0/ror/dex/")
}

func LoadSettings() {
	environment := viper.GetString(configconsts.ENVIRONMENT)
	rlog.Info("environment loaded", rlog.String("environment", environment))

	dexHost := viper.GetString(configconsts.DEX_HOST)
	rlog.Info("dex host loaded", rlog.String("host", dexHost))

	dexPort := viper.GetString(configconsts.DEX_PORT)
	rlog.Info("dex port loaded", rlog.String("port", dexPort))

	dexGrpcPort := viper.GetString(configconsts.DEX_GRPC_PORT)
	rlog.Info("dex grpc port loaded", rlog.String("grpc port", dexGrpcPort))

	role := viper.GetString(configconsts.ROLE)
	rlog.Info("role loaded", rlog.String("role", role))
}
