package main

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/httpserver"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/mstanzuconnections"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/rabbitmq/mstanzurabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/rabbitmq/mstanzurabbitmqhandler"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/rorclient"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/settings"
	"os"
	"os/signal"

	"syscall"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	"github.com/spf13/viper"

	// https://blog.devgenius.io/know-gomaxprocs-before-deploying-your-go-app-to-kubernetes-7a458fb63af1
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(rlog.Infof))
}
func main() {
	cancelChan := make(chan os.Signal, 1)
	stop := make(chan struct{})
	// catch SIGETRM or SIGINTERRUPT
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	rlog.Info("Tanzu micro service starting")
	settings.Load()

	role := viper.GetString(configconsts.ROLE)
	rorclient.SetupRORClient()
	mstanzuconnections.InitConnections()
	mstanzurabbitmqdefinitions.InitOrDie()

	go func() {
		httpserver.InitHttpServer()
		sig := <-cancelChan
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	if viper.GetBool(configconsts.ENABLE_TRACING) {
		go func() {
			trace.ConnectTracer(stop, role, viper.GetString(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT))
			sig := <-cancelChan
			_, _ = fmt.Println()
			_, _ = fmt.Println(sig)
			stop <- struct{}{}
		}()
	}

	mstanzurabbitmqhandler.StartListening()

	sig := <-cancelChan
	rlog.Info("Caught signal", rlog.Any("singal", sig))
	// shutdown other goroutines gracefully
	// close other resources
}
