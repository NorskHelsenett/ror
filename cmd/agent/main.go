package main

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/agent/clients"
	"github.com/NorskHelsenett/ror/cmd/agent/config"
	"github.com/NorskHelsenett/ror/cmd/agent/controllers"
	"github.com/NorskHelsenett/ror/cmd/agent/httpserver"
	"github.com/NorskHelsenett/ror/cmd/agent/scheduler"
	"github.com/NorskHelsenett/ror/cmd/agent/services"
	"github.com/NorskHelsenett/ror/cmd/agent/services/resourceupdatev2"
	"os"
	"os/signal"
	"time"

	"github.com/NorskHelsenett/ror/internal/checks/initialchecks"
	"github.com/NorskHelsenett/ror/internal/kubernetes/operator/initialize"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"syscall"

	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"

	// https://blog.devgenius.io/know-gomaxprocs-before-deploying-your-go-app-to-kubernetes-7a458fb63af1

	"go.uber.org/automaxprocs/maxprocs"
)

// init
func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(rlog.Infof))
	config.Init()
}

func main() {
	_ = "rebuild 6"
	rlog.Info("Agent is starting", rlog.String("version", viper.GetString(configconsts.VERSION)))
	sigs := make(chan os.Signal, 1)                                    // Create channel to receive os signals
	stop := make(chan struct{})                                        // Create channel to receive stop signal
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // Register the sigs channel to receieve SIGTERM

	go func() {
		services.GetEgressIp()
		sig := <-sigs
		_, _ = fmt.Println()
		_, _ = fmt.Print(sig)
		stop <- struct{}{}
	}()

	clients.Initialize()

	k8sClient, err := clients.Kubernetes.GetKubernetesClientset()
	if err != nil {
		panic(err.Error())
	}

	discoveryClient, err := clients.Kubernetes.GetDiscoveryClient()
	if err != nil {
		rlog.Error("failed to get discovery client", err)
	}

	dynamicClient, err := clients.Kubernetes.GetDynamicClient()
	if err != nil {
		rlog.Error("failed to get dynamic client", err)
	}

	ns := viper.GetString(configconsts.POD_NAMESPACE)
	if ns == "" {
		rlog.Fatal("POD_NAMESPACE is not set", nil)
	}

	_, err = k8sClient.CoreV1().Namespaces().Get(context.Background(), ns, metav1.GetOptions{})
	if err != nil {
		rlog.Fatal("could not get namespace", err)
	}

	err = initialchecks.HasSuccessfullRorApiConnection()
	if err != nil {
		rlog.Fatal("could not connect to ror-api", err)
	}

	err = services.ExtractApikeyOrDie()
	if err != nil {
		rlog.Fatal("could not get or create secret", err)
	}

	clusterId, err := initialize.GetOwnClusterId()
	if err != nil {
		rlog.Fatal("could not fetch clusterid from ror-api", err)
	}
	viper.Set(configconsts.CLUSTER_ID, clusterId)

	err = resourceupdatev2.ResourceCache.Init()
	if err != nil {
		rlog.Fatal("could not get hashlist for clusterid", err)
	}

	err = scheduler.HeartbeatReporting()
	if err != nil {
		rlog.Fatal("could not send heartbeat to api", err)
	}

	// waiting for ip check to finish :)
	time.Sleep(time.Second * 1)

	schemas := clients.InitSchema()

	for _, schema := range schemas {
		check, err := discovery.IsResourceEnabled(discoveryClient, schema)
		if err != nil {
			rlog.Error("Could not query resources from cluster", err)
		}
		if check {
			controller := controllers.NewDynamicController(dynamicClient, schema)

			go func() {
				controller.Run(stop)
				sig := <-sigs
				_, _ = fmt.Println()
				_, _ = fmt.Println(sig)
				stop <- struct{}{}
			}()
		} else {
			errmsg := fmt.Sprintf("Could not register resource %s", schema.Resource)
			rlog.Info(errmsg)
		}
	}

	go func() {
		httpserver.InitHttpServer()
		sig := <-sigs
		_, _ = fmt.Println()
		_, _ = fmt.Println(sig)
		stop <- struct{}{}
	}()

	scheduler.SetUpScheduler()

	<-stop
	rlog.Info("Shutting down...")
}
