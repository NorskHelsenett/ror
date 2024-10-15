package main

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/audit/httpserver"
	"github.com/NorskHelsenett/ror/cmd/audit/msauditconnections"
	"github.com/NorskHelsenett/ror/cmd/audit/rabbitmq/msauditrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/audit/rabbitmq/msauditrabbitmqhandler"
	"github.com/NorskHelsenett/ror/cmd/audit/settings"
	"os"
	"os/signal"

	"syscall"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/databasecredhelper"

	"github.com/NorskHelsenett/ror/pkg/rlog"

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

	rlog.Info("Audit micro service starting")
	settings.Load()

	msauditconnections.InitConnections()

	mongocredshelper := databasecredhelper.NewVaultDBCredentials(msauditconnections.VaultClient, viper.GetString(configconsts.ROLE), "mongodb")
	mongodb.Init(mongocredshelper, viper.GetString(configconsts.MONGODB_HOST), viper.GetString(configconsts.MONGODB_PORT), viper.GetString(configconsts.MONGODB_DATABASE))

	msauditrabbitmqdefinitions.InitOrDie()

	go func() {
		httpserver.InitHttpServer()
	}()

	go func() {
		trace.ConnectTracer(stop, viper.GetString(configconsts.ROLE), viper.GetString(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT))
		sig := <-cancelChan
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	msauditrabbitmqhandler.StartListening()

	sig := <-cancelChan
	rlog.Info("Caught signal", rlog.Any("singal", sig))
	// shutdown other goroutines gracefully
	// close other resources
}
