package agentconfig

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/config/rorversion"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

var (
	Version string = "1.1.0"
	Commit  string = "FFFFF"
)

var (
	ErrorCount int
)

func Init() {
	rlog.InitializeRlog()
	rlog.Info("Configuration initializing ...")
	viper.SetDefault(configconsts.VERSION, Version)
	viper.SetDefault(configconsts.ROLE, "ClusterAgent")
	viper.SetDefault(configconsts.COMMIT, Commit)
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

func GetRorVersion() rorversion.RorVersion {
	return rorversion.NewRorVersion(viper.GetString(configconsts.VERSION), viper.GetString(configconsts.COMMIT))
}
