package dynamicclient

import (
	"github.com/NorskHelsenett/ror/cmd/agentv2/services/resourceupdatev2"

	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func addResource(obj any) {
	rawData := obj.(*unstructured.Unstructured)
	resourceupdatev2.SendResource(rortypes.K8sActionAdd, rawData)
}

func deleteResource(obj any) {
	rawData := obj.(*unstructured.Unstructured)
	resourceupdatev2.SendResource(rortypes.K8sActionDelete, rawData)
}

func updateResource(_ any, obj any) {
	rawData := obj.(*unstructured.Unstructured)
	resourceupdatev2.SendResource(rortypes.K8sActionUpdate, rawData)
}
