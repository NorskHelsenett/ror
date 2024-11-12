package main

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/httpserver"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/rabbitmq/tanzuagentrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/rabbitmq/tanzuagentrabbitmqhandler"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/services/resourceupdatev2"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/settings"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/tanzuagentconnections"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/tanzuservice"
	"os"
	"os/signal"
	"time"

	"syscall"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	// https://blog.devgenius.io/know-gomaxprocs-before-deploying-your-go-app-to-kubernetes-7a458fb63af1
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(rlog.Infof))
}

func main() {
	rlog.Info("Tanzu Agent is starting", rlog.String("version", settings.AgentVersionNumber))
	sigs := make(chan os.Signal, 1)                                    // Create channel to receive os signals
	stop := make(chan struct{})                                        // Create channel to receive stop signal
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // Register the sigs channel to receieve SIGTERM

	settings.Load()

	tanzuAccess := viper.GetBool(configconsts.TANZU_AGENT_TANZU_ACCESS)

	tanzuagentconnections.InitConnections()

	tanzuagentrabbitmqdefinitions.InitOrDie()

	vaultData, err := tanzuagentconnections.VaultClient.GetSecret("secret/data/v1.0/ror/tanzu/agent")
	if err != nil {
		rlog.Fatal("could not get tanzu secrets", err)
	}

	data, ok := vaultData["data"].(map[string]interface{})
	if !ok {
		rlog.Error("could not", fmt.Errorf("data type assertion failed: %T %#v", vaultData["data"], vaultData["data"]))
		panic("could not get tanzu secrets")
	}

	tanzuUsername, _ := data["tanzuUsername"].(string)
	tanzuPassword, _ := data["tanzuPassword"].(string)

	if tanzuAccess && (tanzuPassword == "" || tanzuUsername == "") {
		rlog.Fatal("tanzu secrets are not set", nil)
	}

	viper.Set(configconsts.TANZU_AGENT_USERNAME, tanzuUsername)
	viper.Set(configconsts.TANZU_AGENT_PASSWORD, tanzuPassword)

	if len(viper.GetString(configconsts.API_KEY)) == 0 {
		rlog.Fatal("API_KEY is not set", nil)
	}

	kubeconfigPath := viper.GetString(configconsts.TANZU_AGENT_KUBECONFIG)
	if len(kubeconfigPath) == 0 {
		rlog.Fatal("KUBECONFIG is not set, please set absolute path to kubeconfig file", nil)
	}

	go func() {
		httpserver.InitHttpServer()
		sig := <-sigs
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	if tanzuAccess {
		go func() {
			err := resourceupdatev2.ResourceCache.Init()
			if err != nil {
				rlog.Fatal("could not get hashlist for clusterid", err)
			}
			tanzuservice.Run(sigs, stop)
			sig := <-sigs
			_, _ = fmt.Println()
			_, _ = fmt.Println(sig)
			stop <- struct{}{}
		}()
	}

	time.Sleep(2 * time.Second)
	tanzuagentrabbitmqhandler.StartListening()

	sig := <-stop
	rlog.Info("Caught signal", rlog.Any("signal", sig))
	// shutdown other goroutines gracefully
	// close other resources
}
