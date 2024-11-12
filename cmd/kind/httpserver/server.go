package httpserver

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/otel/httpserver"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func InitHttpServer() {
	serveAddress := ":18084"
	err := httpserver.RunOtelHttpHealthServer(serveAddress, httpserver.HealthHandlerTODO)
	if err != nil {
		rlog.Fatal("could not start health server", err)
	}
}
