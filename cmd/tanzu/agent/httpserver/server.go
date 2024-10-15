package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/helpers/otel/httpserver"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

func InitHttpServer() {
	serveAddress := fmt.Sprintf("0.0.0.0:%s", viper.GetString(configconsts.TANZU_AGENT_HEALTH_PORT))
	err := httpserver.RunOtelHttpHealthServer(serveAddress, health)
	if err != nil {
		rlog.Fatal("could not start health server", err)
	}
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := TanzuHealtReport{
		Version:   viper.GetString(configconsts.VERSION),
		CommitSha: viper.GetString(configconsts.COMMIT),
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		rlog.Fatal("Could not encode health report", err)
	}
}

type TanzuHealtReport struct {
	Version   string `json:"version"`
	CommitSha string `json:"commitSha"`
}
