package rorserviceconfig

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
	healthserver "github.com/NorskHelsenett/ror/pkg/helpers/rorhealth/server"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/services/rorservice/rorserviceinterface"
	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"
)

type RorServiceConfig struct {
	Role string

	cancelChan chan os.Signal
	stopChan   chan struct{}
}

func MustInit() rorserviceinterface.RorService {

	service := &RorServiceConfig{
		cancelChan: make(chan os.Signal, 1),
		stopChan:   make(chan struct{}),
	}

	service.loadRole()
	signal.Notify(service.cancelChan, syscall.SIGTERM, syscall.SIGINT)

	rlog.Infof("%s service starting. Version: %s", service.Role, rorversion.GetRorVersion().GetVersionWithCommit())

	healthserver.MustStartWithDefaults()
	service.enableTracing()

	return service
}

func (c *RorServiceConfig) loadRole() {
	role := rorconfig.GetString(rorconfig.ROLE)
	if role == "" {
		role = "default"
	}
	c.Role = role
}

func (c *RorServiceConfig) enableTracing() {
	if rorconfig.GetString(rorconfig.ENABLE_TRACING) == "true" {

		otpendpoint := rorconfig.GetString(rorconfig.OPENTELEMETRY_COLLECTOR_ENDPOINT)
		if otpendpoint == "" {
			rlog.Fatal("OpenTelemetry collector endpoint is not set and tracing is enabled", errors.New("OpenTelemetry collector endpoint is not set in env"))
		}

		trace.StartTracing(c.GetStopChan(), c.GetCancelChan(), c.Role, otpendpoint)
	}
}

func (c *RorServiceConfig) Cancel() {
	close(c.cancelChan)
}

func (c *RorServiceConfig) Stop() {
	close(c.stopChan)
}

func (c *RorServiceConfig) GetCancelChan() chan os.Signal {
	return c.cancelChan
}

func (c *RorServiceConfig) GetStopChan() chan struct{} {
	return c.stopChan
}

func (c *RorServiceConfig) Wait() {
	sig := <-c.cancelChan
	<-c.stopChan
	rlog.Info("Caught signal", rlog.Any("signal", sig))
}
