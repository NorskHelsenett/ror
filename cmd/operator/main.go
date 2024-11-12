/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/config/rorclientconfig"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/cmd/agentv2/agentconfig"
	"github.com/NorskHelsenett/ror/cmd/operator/clients"
	"github.com/NorskHelsenett/ror/cmd/operator/initialservices"
	"github.com/NorskHelsenett/ror/cmd/operator/scheduler"
	"github.com/NorskHelsenett/ror/cmd/operator/variables"

	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	//

	"github.com/spf13/viper"
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	appsv1alpha1 "github.com/NorskHelsenett/ror/cmd/operator/api/v1alpha1"
	"github.com/NorskHelsenett/ror/cmd/operator/controllers"

	//+kubebuilder:scaffold:imports

	// https://blog.devgenius.io/know-gomaxprocs-before-deploying-your-go-app-to-kubernetes-7a458fb63af1
	"go.uber.org/automaxprocs/maxprocs"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_, _ = maxprocs.Set(maxprocs.Logger(rlog.Infof))
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(appsv1alpha1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme

	rlog.Info("Starting ROR-Operator ")
	variables.LoadSettings()
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan struct{})

	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Metrics: metricsserver.Options{
			BindAddress: metricsAddr,
		},
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "55497e69.ror",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

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

	k8sClient, err := clients.Kubernetes.GetKubernetesClientset()
	if err != nil {
		panic(err.Error())
	}

	metricsClient, err := clients.Kubernetes.GetMetricsClient()
	if err != nil {
		rlog.Error("failed to get metrics client", err)

	}

	err = initialservices.GetOrCreateNamespace(k8sClient)
	if err != nil {
		rlog.Fatal("could not get or create namespace", err)
	}

	// This will generate a new API key if it does not exist
	err = initialservices.ExtractApikeyOrDie(k8sClient, metricsClient)
	if err != nil {
		rlog.Fatal("could not get or create secret", err)
	}

	err = initialservices.FetchConfiguration()
	if err != nil {
		rlog.Fatal("Could not fetch configuration for cluster", err)
	}

	if err = (&controllers.TaskReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Task")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	go func() {
		rlog.Debug("Initializing operator")

		scheduler.SetUpScheduler(k8sClient)

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- struct{}{}
	}()

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

	<-done
	fmt.Println("Ror-Operator finishing")
}
