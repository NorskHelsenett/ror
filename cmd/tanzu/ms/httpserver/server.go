package httpserver

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/helpers/otel/httpserver"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

// InitHttpServer The function initializes the http server and starts listening on the configured port.
func InitHttpServer() {
	serveAddress := fmt.Sprintf(":%s", viper.GetString(configconsts.TANZU_AGENT_HEALTH_PORT))
	err := httpserver.RunOtelHttpHealthServer(serveAddress, httpserver.HealthHandlerTODO)
	if err != nil {
		rlog.Fatal("could not start health server", err)
	}
}
