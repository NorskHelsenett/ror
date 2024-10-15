package config

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

var (
	version    string = "1.1.0"
	commit     string = "FFFFF"
	ErrorCount int
)

func Init() {
	rlog.InitializeRlog()
	rlog.Info("Configuration initializing ...")
	viper.SetDefault(configconsts.VERSION, version)
	viper.SetDefault(configconsts.COMMIT, commit)
	viper.SetDefault(configconsts.HEALTH_ENDPOINT, ":8100")
	viper.SetDefault(configconsts.POD_NAMESPACE, "ror")
	viper.SetDefault(configconsts.API_KEY_SECRET, "ror-apikey")

	viper.AutomaticEnv()
}

func IncreaseErrorCount() {
	ErrorCount++
}
func ResetErrorCount() {
	ErrorCount = 0
}
