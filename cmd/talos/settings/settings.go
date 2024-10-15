package settings

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
)

var (
	VaultSecret *vault.Secret
	Role        string
)

func init() {

	viper.SetDefault(configconsts.VAULT_URL, "http://localhost:8200")
	viper.SetDefault("KUBECTL_BASE_URL", "https://127.0.0.1")
	viper.GetString(configconsts.ROLE)

	viper.AutomaticEnv()
}

func Load() {
	environment := viper.GetString(configconsts.ENVIRONMENT)
	rlog.Info("loaded environment", rlog.String("Environment", environment))

	_ = viper.WriteConfig()
}
