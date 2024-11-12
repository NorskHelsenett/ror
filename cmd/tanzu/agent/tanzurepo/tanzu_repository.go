package tanzurepo

import (
	"context"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/kubernetesclients"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func DeleteResource(gvr schema.GroupVersionResource, resource unstructured.Unstructured) error {
	dynamicClient := kubernetesclients.GetDynamicClientOrDie()

	err := dynamicClient.Resource(gvr).Namespace(resource.GetNamespace()).Delete(context.Background(), resource.GetName(), v1.DeleteOptions{})
	if err != nil {
		rlog.Error("failed to resource", err)
		return err
	}

	rlog.Info("Delete resource result")
	return nil
}

func UpdateResource(gvr schema.GroupVersionResource, resource unstructured.Unstructured) error {
	dynamicClient := kubernetesclients.GetDynamicClientOrDie()
	updateResult, err := dynamicClient.Resource(gvr).Namespace(resource.GetNamespace()).Update(context.Background(), &resource, v1.UpdateOptions{})
	if err != nil {
		rlog.Error("failed to update resource", err)
		return err
	}

	rlog.Info("Updated resource", rlog.Any("result", updateResult))
	return nil
}

func GetResource(gvr schema.GroupVersionResource, name string, namespace string) (*unstructured.Unstructured, error) {
	dynamicClient := kubernetesclients.GetDynamicClientOrDie()
	resource, err := dynamicClient.Resource(gvr).Namespace(namespace).Get(context.Background(), name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func CreateResource(gvr schema.GroupVersionResource, resource unstructured.Unstructured) (*unstructured.Unstructured, error) {
	dynamicClient := kubernetesclients.GetDynamicClientOrDie()
	result, err := dynamicClient.Resource(gvr).Namespace(resource.GetNamespace()).Create(context.Background(), &resource, v1.CreateOptions{})
	if err != nil {
		rlog.Error("failed to create resource", err)
		return nil, err
	}

	rlog.Info("Created resource")
	return result, nil
}
