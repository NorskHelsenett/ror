package settings

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault(configconsts.ENVIRONMENT, "production")
	viper.SetDefault(rlog.LOG_LEVEL, "info")
	viper.SetDefault(configconsts.ROLE, "ror-ms-slack")
}

func Load() {
	environment := viper.GetString(configconsts.ENVIRONMENT)
	rlog.Info("loaded environment", rlog.String("Environment", environment))

	_ = viper.WriteConfig()
}
