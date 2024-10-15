package controllers

import (
	"context"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/services/resourceupdatev2"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
)

var (
	stopCh chan struct{}
)

type DynamicController struct {
	dynInformer cache.SharedIndexInformer
	client      dynamic.Interface
}

func (c *DynamicController) Run(stop chan struct{}) {
	// Execute go function
	stopCh = stop

	loginEveryMinute := viper.GetInt(configconsts.TANZU_AGENT_LOGIN_EVERY_MINUTE)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(loginEveryMinute)*time.Minute)

	defer close(stop)
	defer func() {
		rlog.Debug("Stopping informer")
		stopCh <- struct{}{}
		cancel()
	}()
	c.dynInformer.Run(ctx.Done())

	rlog.Debug("Canceling informer")
	cancel()
}

// Function creates a new dynamic controller to listen for api-changes in provided GroupVersionResource
func NewDynamicController(client dynamic.Interface, resource schema.GroupVersionResource, namespace apicontracts.Workspace) *DynamicController {
	dynWatcher := &DynamicController{}
	dynInformer := dynamicinformer.NewFilteredDynamicSharedInformerFactory(client, 0, namespace.Name, nil)
	informer := dynInformer.ForResource(resource).Informer()

	err := informer.SetWatchErrorHandler(setCustomErrorHandler)
	if err != nil {
		rlog.Error("error setting watch error handler", err)
		stopCh <- struct{}{}
	}

	dynWatcher.client = client
	dynWatcher.dynInformer = informer

	_, err = informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    addResource,
		UpdateFunc: updateResource,
		DeleteFunc: deleteResource,
	})
	if err != nil {
		rlog.Error("error adding event handler", err)
	}

	return dynWatcher
}

func setCustomErrorHandler(r *cache.Reflector, err error) {
	if errors.IsUnauthorized(err) {
		rlog.Error("unauthorized config", err)
		stopCh <- struct{}{}
	} else {
		cache.DefaultWatchErrorHandler(r, err)
	}
}

func addResource(obj any) {
	rawData := obj.(*unstructured.Unstructured)
	if !isValid(rawData) {
		return
	}

	resourceupdatev2.SendResource(apiresourcecontracts.K8sActionAdd, rawData)
}

func deleteResource(obj any) {
	rawData := obj.(*unstructured.Unstructured)
	if !isValid(rawData) {
		return
	}
	resourceupdatev2.SendResource(apiresourcecontracts.K8sActionDelete, rawData)
}

func updateResource(_ any, updated any) {
	rawData := updated.(*unstructured.Unstructured)
	if !isValid(rawData) {
		return
	}

	resourceupdatev2.SendResource(apiresourcecontracts.K8sActionUpdate, rawData)
}

func isValid(obj *unstructured.Unstructured) bool {
	if obj == nil {
		return false
	}
	if obj == nil {
		return false
	}

	if len(obj.GetName()) == 0 {
		return false
	}

	if len(obj.GetAPIVersion()) == 0 || len(obj.GetKind()) == 0 {
		return false
	}

	return true
}
