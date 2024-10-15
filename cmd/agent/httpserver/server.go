package httpserver

import (
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/agent/clients"
	"github.com/NorskHelsenett/ror/cmd/agent/config"
	"net/http"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/helpers/otel/httpserver"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

var healthStatus HealthStatus

func InitHttpServer() {
	serveAddress := viper.GetString(configconsts.HEALTH_ENDPOINT)

	healthStatus = getHealthReportOrDie()

	err := httpserver.RunOtelHttpHealthServer(serveAddress, health)
	if err != nil {
		rlog.Fatal("could not start health server", err)
	}
}

func health(w http.ResponseWriter, req *http.Request) {
	healthStatus = getHealthReportOrDie()
	js, err := json.Marshal(healthStatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if strings.Contains(healthStatus.Status, "Crap") || strings.Contains(healthStatus.Status, "UnHealthy") || healthStatus.Report.ErrorCount > 0 || len(viper.GetString(configconsts.API_KEY)) == 0 {
		w.WriteHeader(500)
	}

	_, _ = w.Write(js)
}

func getHealthReportOrDie() (result HealthStatus) {
	gotAuthSetting := false
	if len(viper.GetString(configconsts.API_KEY)) != 0 {
		gotAuthSetting = true
	}

	hasK8sConfig := false
	if clients.Kubernetes.GetConfig() != nil {
		hasK8sConfig = true
	}

	status := "Crap ... ET phone home!"
	if hasK8sConfig && config.ErrorCount == 0 && gotAuthSetting {
		status = "Healthy"
	} else if hasK8sConfig && (!gotAuthSetting || config.ErrorCount > 0) {
		status = "UnHealthy"
	}

	healthStatus = HealthStatus{
		Status: status,
		Report: HealthReport{
			Kubernetes: K8sReport{
				HasConfig: hasK8sConfig,
			},
			RorApi: RorApiReport{
				GotToken: gotAuthSetting,
			},
			ErrorCount: config.ErrorCount,
		},
	}

	return healthStatus
}
