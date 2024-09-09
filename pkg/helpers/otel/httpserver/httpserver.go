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
