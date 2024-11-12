package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/auth/controllers/kubeconfigcontroller"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func InitHttpServer() {

	controllerPath := "/v1/kubeconfig"
	serverAddress := fmt.Sprintf("0.0.0.0:%s", viper.GetString(configconsts.TANZU_AUTH_HEALTH_PORT))

	mux := http.NewServeMux()
	mux.Handle("/health", http.HandlerFunc(health))
	mux.Handle(controllerPath, http.HandlerFunc(MiddlewareLogging(MiddlewareJsonContentType(kubeconfigcontroller.HandleKubeConfigRequest))))
	var handler http.Handler = mux

	httpServer := &http.Server{
		Addr:              serverAddress,
		Handler:           otelhttp.NewHandler(handler, "/"),
		ReadHeaderTimeout: 0,
	}

	rlog.Fatal("Could not run http server",
		httpServer.ListenAndServe())
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
