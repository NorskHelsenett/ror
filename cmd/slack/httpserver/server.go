package httpserver

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// InitHttpServer The function initializes the http server and starts listening on the configured port.
func InitHttpServer() {
	http.Handle("/health", otelhttp.NewHandler(http.HandlerFunc(health), "/health"))
	serveAddress := fmt.Sprintf(":%s", "8080")

	err := http.ListenAndServe(serveAddress, nil)
	if err != nil {
		panic(err)
	}
}

// The function writes a 200 status code and the string "Healthy" to the response writer.
func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("Healthy"))
	if err != nil {
		panic(err)
	}
}
