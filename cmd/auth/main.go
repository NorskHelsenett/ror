package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"syscall"

	"github.com/NorskHelsenett/ror/cmd/auth/httpserver"
	"github.com/NorskHelsenett/ror/cmd/auth/msauthconnections"
	"github.com/NorskHelsenett/ror/cmd/auth/rabbitmq/msauthrabbitmqdefintions"
	"github.com/NorskHelsenett/ror/cmd/auth/rabbitmq/msauthrabbitmqhandler"
	"github.com/NorskHelsenett/ror/cmd/auth/settings"

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

	rlog.Info("Auth micro service starting")
	settings.LoadSettings()

	msauthconnections.InitConnections()
	msauthrabbitmqdefintions.InitOrDie()

	dexHost := viper.GetString(configconsts.DEX_HOST)
	if len(dexHost) == 0 {
		rlog.Fatal("failed to load config", fmt.Errorf("empty dex host is empty"))
	}

	go func() {
		httpserver.InitHttpServer()
	}()

	if viper.GetBool(configconsts.ENABLE_TRACING) {
		go func() {
			trace.ConnectTracer(stop, "ror-auth", viper.GetString(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT))
			sig := <-cancelChan
			_, _ = fmt.Println()
			_, _ = fmt.Println(sig)
			stop <- struct{}{}
		}()
	}

	msauthrabbitmqhandler.StartListening()

	sig := <-cancelChan
	rlog.Info("Caught signal", rlog.Any("singal", sig))
	// shutdown other goroutines gracefully
	// close other resources
}
