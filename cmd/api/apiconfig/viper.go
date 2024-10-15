package apiconfig

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/config/rorversion"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

var Version string = "0.1.0"
var Commit string = "FFFFFFF"

var RorVersion rorversion.RorVersion

func InitViper() {
	rlog.Info("initializing configuration")

	viper.AutomaticEnv()

	viper.SetDefault(configconsts.VERSION, Version)
	viper.SetDefault(configconsts.COMMIT, Commit)
	viper.SetDefault(configconsts.API_KEY_SALT, "")
	viper.SetDefault(configconsts.ROLE, "ror-api")

	// Remove we dont set env in variables.
	viper.SetDefault(configconsts.DEVELOPMENT, false)
	//Remove used in gin.go
	viper.SetDefault(configconsts.HTTP_PORT, "8080")
	viper.SetDefault(configconsts.HTTP_TIMEOUT, "15s")

	viper.SetDefault(configconsts.PROFILER_ENABLED, false)
	viper.SetDefault(configconsts.ENABLE_TRACING, true)
	viper.SetDefault(configconsts.TRACER_ID, "ror-api")

	viper.SetDefault(configconsts.OIDC_PROVIDER, "http://localhost:5556/dex")
	viper.SetDefault(configconsts.OIDC_CLIENT_ID, "ror.sky.test.nhn.no")
	viper.SetDefault(configconsts.OIDC_DEVICE_CLIENT_ID, "ror-cli")
	viper.SetDefault(configconsts.OIDC_SKIP_ISSUER_VERIFY, false)

	viper.SetDefault("VAULT_PREFIX", "http://")
	viper.SetDefault("VAULT_HOST", "localhost")
	viper.SetDefault("VAULT_PORT", "8200")

	viper.SetDefault(configconsts.VAULT_URL, fmt.Sprintf("%s%s:%s", viper.GetString("VAULT_PREFIX"), viper.GetString("VAULT_HOST"), viper.GetString("VAULT_PORT")))

	viper.SetDefault(configconsts.RABBITMQ_HOST, "localhost")
	viper.SetDefault(configconsts.RABBITMQ_PORT, "5672")
	viper.SetDefault(configconsts.RABBITMQ_BROADCAST_NAME, "nhn.ror.broadcast")

	viper.SetDefault(configconsts.REDIS_HOST, "localhost")
	viper.SetDefault(configconsts.REDIS_PORT, "6379")

	viper.SetDefault(configconsts.MONGODB_HOST, "localhost")
	viper.SetDefault(configconsts.MONGODB_PORT, "27017")
	viper.SetDefault(configconsts.MONGO_DATABASE, "nhn-ror")

	viper.SetDefault(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT, "opentelemetry-collector:4317")
	viper.SetDefault(configconsts.HELSEGITLAB_BASE_URL, "https://helsegitlab.nhn.no/api/v4/projects/")

	viper.SetDefault(configconsts.HEALTH_ENDPOINT, "0.0.0.0:9999")
}

func GetRorVersion() rorversion.RorVersion {
	return rorversion.NewRorVersion(viper.GetString(configconsts.VERSION), viper.GetString(configconsts.COMMIT))
}
