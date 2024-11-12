package main

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/auth/httpserver"
	"github.com/NorskHelsenett/ror/cmd/tanzu/auth/settings"
	"os"
	"os/signal"

	"syscall"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	// https://blog.devgenius.io/know-gomaxprocs-before-deploying-your-go-app-to-kubernetes-7a458fb63af1

	"go.uber.org/automaxprocs/maxprocs"
)

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(rlog.Infof))
}

func main() {
	rlog.Info("MS Tanzu Auth is starting", rlog.String("version", settings.TanzuAuthVersionNumber))
	sigs := make(chan os.Signal, 1)                                    // Create channel to receive os signals
	stop := make(chan struct{})                                        // Create channel to receive stop signal
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // Register the sigs channel to receieve SIGTERM

	settings.Load()

	//role := viper.GetString(configconsts.ROLE)
	//vaultclient.Init(role, viper.GetString(configconsts.VAULT_URL))
	//tanzuauthconnections.InitConnections()
	//tanzuauthrabbitmqdefinitions.InitOrDie()

	// if len(viper.GetString(configconsts.API_KEY)) == 0 {
	// 	rlog.Fatal("API_KEY is not set", nil)
	// }

	go func() {
		httpserver.InitHttpServer()
		sig := <-sigs
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	//tanzuauthrabbitmqhandler.StartListening()

	sig := <-stop
	rlog.Info("Caught signal", rlog.Any("signal", sig))
	// shutdown other goroutines gracefully
	// close other resources
}
