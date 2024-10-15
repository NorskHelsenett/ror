package settings

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
)

var (
	VaultClient *vaultclient.VaultClient
	VaultSecret *vault.Secret
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault(configconsts.ENVIRONMENT, "production")
	viper.SetDefault(rlog.LOG_LEVEL, "info")
	viper.SetDefault(configconsts.ROLE, "ror-ms-switchboard")
}

func Load() {
	environment := viper.GetString(configconsts.ENVIRONMENT)
	rlog.Info("loaded environment", rlog.String("Environment", environment))

	_ = viper.WriteConfig()
}
