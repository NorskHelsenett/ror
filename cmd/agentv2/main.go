package main

import (
	"github.com/NorskHelsenett/ror/cmd/agentv2/agentconfig"
	"github.com/NorskHelsenett/ror/cmd/agentv2/clients"
	"github.com/NorskHelsenett/ror/cmd/agentv2/clients/dynamicclient"
	"github.com/NorskHelsenett/ror/cmd/agentv2/scheduler"
	"github.com/NorskHelsenett/ror/cmd/agentv2/services/resourceupdatev2"
	"os"
	"os/signal"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/config/rorclientconfig"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"syscall"

	"github.com/spf13/viper"

	// https://blog.devgenius.io/know-gomaxprocs-before-deploying-your-go-app-to-kubernetes-7a458fb63af1
	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	_ = "rebuild 6"
	_, _ = maxprocs.Set(maxprocs.Logger(rlog.Infof))
	agentconfig.Init()
	rlog.Info("Agent is starting", rlog.String("version", viper.GetString(configconsts.VERSION)), rlog.String("commit", viper.GetString(configconsts.COMMIT)))
	sigs := make(chan os.Signal, 1)                                    // Create channel to receive os signals
	stop := make(chan struct{})                                        // Create channel to receive stop signal
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // Register the sigs channel to receieve SIGTERM

	clientConfig := rorclientconfig.ClientConfig{
		Role:                     viper.GetString(configconsts.ROLE),
		Namespace:                viper.GetString(configconsts.POD_NAMESPACE),
		ApiKeySecret:             viper.GetString(configconsts.API_KEY_SECRET),
		ApiKey:                   viper.GetString(configconsts.API_KEY),
		ApiEndpoint:              viper.GetString(configconsts.API_ENDPOINT),
		RorVersion:               agentconfig.GetRorVersion(),
		MustInitializeKubernetes: true,
	}

	clients.InitClients(clientConfig)

	err := resourceupdatev2.ResourceCache.Init()
	if err != nil {
		rlog.Fatal("could not get hashlist for clusterid", err)
	}

	err = dynamicclient.Start(clients.Kubernetes, stop, sigs)
	if err != nil {
		rlog.Fatal("could not start dynamic client", err)
	}

	scheduler.SetUpScheduler()

	<-stop
	rlog.Info("Shutting down...")
}
