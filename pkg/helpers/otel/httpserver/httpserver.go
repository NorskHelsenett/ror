package httpserver

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func RunOtelHttpHealthServer(serverAddress string, healthHandler http.HandlerFunc) error {
	mux := http.NewServeMux()
	mux.Handle("/health", http.HandlerFunc(healthHandler))
	var handler http.Handler = mux

	httpServer := &http.Server{
		Addr:              serverAddress,
		Handler:           otelhttp.NewHandler(handler, "/"),
		ReadHeaderTimeout: 0,
	}

	return httpServer.ListenAndServe()

}

// The function writes a 200 status code and the string "Healthy" to the response writer.
func HealthHandlerTODO(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("Healthy"))
	if err != nil {
		panic(err)
	}
}
