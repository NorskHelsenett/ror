package main

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/slack/httpserver"
	"github.com/NorskHelsenett/ror/cmd/slack/rabbitmq/msslackrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/slack/rabbitmq/msslackrabbitmqhandler"
	"github.com/NorskHelsenett/ror/cmd/slack/ror"
	"github.com/NorskHelsenett/ror/cmd/slack/settings"
	"github.com/NorskHelsenett/ror/cmd/slack/slackconnections"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"
	"github.com/slack-go/slack"

	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

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

	rlog.Info("Slack micro service starting")
	settings.Load()

	slackconnections.InitConnections()
	ror.SetupRORClient()
	msslackrabbitmqdefinitions.InitOrDie()

	httpClient := http.Client{}

	if viper.GetBool(configconsts.ENABLE_TRACING) {
		go func() {
			trace.ConnectTracer(stop, "ror-slack", viper.GetString(configconsts.OPENTELEMETRY_COLLECTOR_ENDPOINT))
			sig := <-cancelChan
			_, _ = fmt.Println()
			_, _ = fmt.Println(sig)
			stop <- struct{}{}
		}()
		httpClient.Transport = otelhttp.NewTransport(http.DefaultTransport)
	}

	go func() {
		httpserver.InitHttpServer()
		sig := <-cancelChan
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	slackClient := slack.New(
		viper.GetString(configconsts.SLACK_BOT_TOKEN),
		slack.OptionHTTPClient(&httpClient),
		slack.OptionAppLevelToken(viper.GetString(configconsts.SLACK_APP_TOKEN)),
	)

	msslackrabbitmqhandler.StartListening(slackClient)

	sig := <-cancelChan
	rlog.Info("caught signal", rlog.Any("signal", sig))
}
