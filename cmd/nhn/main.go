package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/databasecredhelper"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/cmd/nhn/httpserver"
	"github.com/NorskHelsenett/ror/cmd/nhn/msnhnconnections"
	"github.com/NorskHelsenett/ror/cmd/nhn/rabbitmq/msnhnrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/nhn/rabbitmq/msnhnrabbitmqhandler"
	"github.com/NorskHelsenett/ror/cmd/nhn/settings"
	"syscall"

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

	rlog.Info("NHN Tooling micro service starting")
	settings.LoadSettings()

	msnhnconnections.InitConnections()

	mongocredshelper := databasecredhelper.NewVaultDBCredentials(msnhnconnections.VaultClient, viper.GetString(configconsts.ROLE), "mongodb")
	mongodb.Init(mongocredshelper, viper.GetString(configconsts.MONGODB_HOST), viper.GetString(configconsts.MONGODB_PORT), viper.GetString(configconsts.MONGODB_DATABASE))

	msnhnrabbitmqdefinitions.InitOrDie()

	go func() {
		httpserver.InitHttpServer()
		sig := <-cancelChan
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	go func() {
		trace.ConnectTracer(stop, "ror-nhn", viper.GetString(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT))
		sig := <-cancelChan
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	msnhnrabbitmqhandler.StartListening()

	sig := <-cancelChan
	rlog.Info("caught signal", rlog.Any("signal", sig))
	// shutdown other goroutines gracefully
	// close other resources
}
