// package dynamic watcher provides a way to watch for changes in kubernetes resources
// the function Start takes a kubernetes clientset and a list of DynamicWatcherConfig
// and starts a watcher for each configuration
package dynamicwatcher

import (
	"fmt"

	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
)

type WatcherConfig struct {
	ResourceDef schema.GroupVersionResource
	Add         func(obj any, isInInitialList bool)
	Delete      func(obj any)
	Update      func(oldObj any, newObj any)
}

type DynamicWatcherConfig interface {
	WatcherConfig() WatcherConfig
}

// Function starts dynamic watchers for provided configurations
func Start(k *kubernetesclient.K8sClientsets, configs []DynamicWatcherConfig, stop chan struct{}) error {
	rlog.Info("Starting dynamic watchers")
	discoveryClient, err := k.GetDiscoveryClient()
	if err != nil {
		rlog.Error("Could not initialize discovery client", err)
		return err
	}
	dynamicClient, err := k.GetDynamicClient()
	if err != nil {
		rlog.Error("Could not initialize dynamic client", err)
		return err
	}

	for _, config := range configs {
		check, err := discovery.IsResourceEnabled(discoveryClient, config.WatcherConfig().ResourceDef)
		if err != nil {
			rlog.Error("Could not query resources from cluster", err)
		}
		if check {
			controller := newDynamicWatcher(dynamicClient, config)
			go func() {
				controller.dynInformer.Run(stop)
			}()
		} else {
			errmsg := fmt.Sprintf("Could not register resource %s", config.WatcherConfig().ResourceDef)
			rlog.Info(errmsg)
		}
	}
	return nil
}

type DynamicWatcher struct {
	dynInformer cache.SharedIndexInformer
	client      dynamic.Interface
}

func (c *DynamicWatcher) Run(stop <-chan struct{}) {
	// Execute go function
	go c.dynInformer.Run(stop)
}

// Function creates a new dynamic controller to listen for api-changes in provided GroupVersionResource
func newDynamicWatcher(client dynamic.Interface, config DynamicWatcherConfig) *DynamicWatcher {
	dynWatcher := &DynamicWatcher{}
	dynInformer := dynamicinformer.NewDynamicSharedInformerFactory(client, 0)
	informer := dynInformer.ForResource(config.WatcherConfig().ResourceDef).Informer()

	dynWatcher.client = client
	dynWatcher.dynInformer = informer

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerDetailedFuncs{
		AddFunc:    config.WatcherConfig().Add,
		UpdateFunc: config.WatcherConfig().Update,
		DeleteFunc: config.WatcherConfig().Delete,
	})
	if err != nil {
		rlog.Error("Error adding event handler", err)
	}
	//rlog.Debug("Watcher", rlog.Any("has synced", watcher.HasSynced()))
	//config.Register(watcher)

	return dynWatcher
}
