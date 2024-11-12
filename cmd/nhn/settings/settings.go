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

const (
	ServiceName = "nhn"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault(rlog.LOG_LEVEL, "INFO")
	viper.SetDefault(configconsts.MONGO_DATABASE, "nhn-ror")
}

func LoadSettings() {
	environment := viper.GetString(configconsts.ENVIRONMENT)
	rlog.Info("environment loaded", rlog.String("environment", environment))

	role := viper.GetString(configconsts.ROLE)
	rlog.Info("role loaded", rlog.String("role", role))
}
