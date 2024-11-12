package controllers

import (
	"github.com/NorskHelsenett/ror/cmd/agent/services/resourceupdatev2"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
)

type DynamicController struct {
	dynInformer cache.SharedIndexInformer
	client      dynamic.Interface
}

func (c *DynamicController) Run(stop <-chan struct{}) {
	// Execute go function
	go c.dynInformer.Run(stop)
}

// Function creates a new dynamic controller to listen for api-changes in provided GroupVersionResource
func NewDynamicController(client dynamic.Interface, resource schema.GroupVersionResource) *DynamicController {
	dynWatcher := &DynamicController{}
	dynInformer := dynamicinformer.NewDynamicSharedInformerFactory(client, 0)
	informer := dynInformer.ForResource(resource).Informer()

	dynWatcher.client = client
	dynWatcher.dynInformer = informer

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    addResource,
		UpdateFunc: updateResource,
		DeleteFunc: deleteResource,
	})

	if err != nil {
		rlog.Error("Error adding event handler", err)
	}
	return dynWatcher
}

func addResource(obj any) {
	rawData := obj.(*unstructured.Unstructured)
	resourceupdatev2.SendResource(apiresourcecontracts.K8sActionAdd, rawData)
}

func deleteResource(obj any) {
	rawData := obj.(*unstructured.Unstructured)
	resourceupdatev2.SendResource(apiresourcecontracts.K8sActionDelete, rawData)
}

func updateResource(_ any, obj any) {
	rawData := obj.(*unstructured.Unstructured)
	resourceupdatev2.SendResource(apiresourcecontracts.K8sActionUpdate, rawData)
}
