package variables

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

var OperatorVersionNumber string = "0.0.1"
var OperatorCommit string = "local-dev"

func init() {
	viper.AutomaticEnv()
	viper.SetDefault(rlog.LOG_LEVEL, "INFO")
	viper.SetDefault(configconsts.POD_NAMESPACE, "ror")
	viper.SetDefault(configconsts.OPERATOR_APPLOG_SECRET_NAME, "ror-app")
	viper.SetDefault(configconsts.CONTAINER_REG_PREFIX, "ncr.sky.nhn.no/")

	viper.SetDefault(configconsts.OPERATOR_BACKOFF_LIMIT, 3)
	viper.SetDefault(configconsts.OPERATOR_DEADLINE_SECONDS, 180)
	viper.SetDefault(configconsts.OPERATOR_JOB_SERVICE_ACCOUNT, "ror-operator")

	viper.SetDefault(configconsts.API_KEY_SECRET, "ror-apikey")
}

func LoadSettings() {
	rlog.Info("Configuration initializing ...")

}

type Version struct {
	Major     int    `json:"major"`
	Minor     int    `json:"minor"`
	Patch     int    `json:"patch"`
	CommitSha string `json:"commitSha"`
}
