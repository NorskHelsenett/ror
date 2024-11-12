package main

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/utils/switchboard"
	"github.com/NorskHelsenett/ror/cmd/api/webserver/sse"
	"os"
	"os/signal"
	"syscall"

	"github.com/NorskHelsenett/ror/cmd/api/apiconfig"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	mongodbseeding "github.com/NorskHelsenett/ror/cmd/api/databases/mongodb/seeding"
	"github.com/NorskHelsenett/ror/cmd/api/rabbitmq/apirabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/api/rabbitmq/apirabbitmqhandler"
	"github.com/NorskHelsenett/ror/cmd/api/utils"
	"github.com/NorskHelsenett/ror/cmd/api/webserver"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/databasecredhelper"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	healthserver "github.com/NorskHelsenett/ror/pkg/helpers/rorhealth/server"
	"github.com/spf13/viper"

	"go.uber.org/automaxprocs/maxprocs"
)

//	@title			Swagger ROR-API
//	@version		0.1
//	@description	ROR-API, need any help? Go to channel #drift-sdi-devops in norskhelsenett.slack.com slack workspace
//	@BasePath		/

//	@contact.name	Privat Sky
//	@contact.url	http://ror.sky.nhn.no

//	@securityDefinitions.apikey	AccessToken
//	@in							header
//	@name						Authorization
//	@securityDefinitions.apikey	ApiKey
//	@in							header
//	@name						X-API-KEY

func main() {
	// rebuild: 2
	ctx := context.Background()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})

	rlog.Infoc(ctx, "ROR by NHN Api startup ")
	rlog.Infof("Version: %s", viper.GetString(configconsts.VERSION))

	_, _ = maxprocs.Set(maxprocs.Logger(rlog.Infof))

	apiconfig.InitViper()

	apiconnections.InitConnections()

	utils.GetCredsFromVault()

	mongocredshelper := databasecredhelper.NewVaultDBCredentials(apiconnections.VaultClient, viper.GetString(configconsts.ROLE), "mongodb")
	mongodb.Init(mongocredshelper, viper.GetString(configconsts.MONGODB_HOST), viper.GetString(configconsts.MONGODB_PORT), viper.GetString(configconsts.MONGO_DATABASE))

	apirabbitmqdefinitions.InitOrDie()
	mongodbseeding.CheckAndSeed(ctx)
	sse.Init()

	if viper.GetBool(configconsts.OIDC_SKIP_ISSUER_VERIFY) {
		rlog.Error("skipping OIDC issuer verification. THIS IS UNSAFE IN PRODUCTION!!!", nil)
	}

	if viper.GetBool(configconsts.ENABLE_TRACING) {
		rlog.Infoc(ctx, "Connecting to open-telemetry")
		go func() {
			trace.ConnectTracer(done, viper.GetString(configconsts.TRACER_ID), viper.GetString(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT))
			sig := <-sigs
			_, _ = fmt.Println()
			_, _ = fmt.Println(sig)
			done <- struct{}{}
		}()
	}

	go func() {
		rlog.Infoc(ctx, "Initializing http server")
		webserver.InitHttpServer()
		sig := <-sigs
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		done <- struct{}{}
	}()

	rlog.Infoc(ctx, "Initializing health server")
	_ = healthserver.Start(healthserver.ServerString(viper.GetString(configconsts.HEALTH_ENDPOINT)))

	if apiconnections.RabbitMQConnection.Ping() {
		switchboard.PublishStarted(ctx)
	}

	apirabbitmqhandler.StartListening()

	<-done
	_, _ = fmt.Println("Ror-API finishing")
}
