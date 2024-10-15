package clusterservicev1alpha2

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/tanzurepo"
	"github.com/NorskHelsenett/ror/internal/factories/storagefactory"

	"github.com/NorskHelsenett/ror/pkg/models/providers/tanzu"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	resourceDefinition = schema.GroupVersionResource{
		Group:    "run.tanzu.vmware.com",
		Version:  "v1alpha2",
		Resource: "tanzukubernetesclusters",
	}
)

func CreateCluster(ctx context.Context, clusterInput tanzu.TanzuKubernetesClusterInput) error {
	storageClass := storagefactory.GetStorageClassByDatacenter(clusterInput.DataCenter)

	if storageClass == "" {
		rlog.Errorc(ctx, "storage class not found", nil)

		return fmt.Errorf("storage class not found")
	}

	controlPlaneReplicas := 1
	if clusterInput.ControlPlane.HighAvailability {
		controlPlaneReplicas = 3
	}

	nodePools := make([]map[string]any, 0)
	for i := 0; i < len(clusterInput.NodePools); i++ {
		nodepoolInput := clusterInput.NodePools[i]

		nodepool := map[string]any{
			"name":         nodepoolInput.Name,
			"replicas":     nodepoolInput.Replicas,
			"storageClass": storageClass,
			"tkr": map[string]any{
				"reference": map[string]any{
					"name": nodepoolInput.KubernetesVersion,
				},
			},
			"vmClass": nodepoolInput.VmClass,
			"volumes": []map[string]any{
				{
					"capacity": map[string]any{
						"storage": "32Gi",
					},
					"mountPath": "/var/lib/containerd",
					"name":      "containerd",
				},
			},
		}

		nodePools = append(nodePools, nodepool)
	}

	if len(nodePools) == 0 {
		rlog.Errorc(ctx, "no nodepools defined", nil)
		return fmt.Errorf("no nodepools defined")
	}

	clusterDefinition := unstructured.Unstructured{
		Object: map[string]any{
			"apiVersion": "run.tanzu.vmware.com/v1alpha2",
			"kind":       "TanzuKubernetesCluster",
			"metadata": map[string]any{
				"name":      clusterInput.Name,
				"namespace": clusterInput.Namespace,
			},
			"spec": map[string]any{
				"settings": map[string]any{
					"storage": map[string]any{
						"defaultClass": storageClass,
					},
					"network": map[string]any{
						"cni": map[string]string{
							"name": "antrea",
						},
						"pods": map[string]any{
							"cidrBlocks": []string{
								"193.0.2.0/16",
							},
						},
						"serviceDomain": "cluster.local",
						"services": map[string]any{
							"cidrBlocks": []string{
								"195.51.100.0/12",
							},
						},
					},
				},
				"topology": map[string]any{
					"controlPlane": map[string]any{
						"replicas":     controlPlaneReplicas,
						"storageClass": storageClass,
						"tkr": map[string]any{
							"reference": map[string]any{
								"name": clusterInput.ControlPlane.KubernetesVersion,
							},
						},
						"vmClass": clusterInput.ControlPlane.VmClass,
					},
					"nodePools": nodePools,
				},
			},
		},
	}

	result, err := tanzurepo.CreateResource(resourceDefinition, clusterDefinition)
	if err != nil {
		rlog.Errorc(ctx, "failed to create cluster", err)
		return err
	}

	rlog.Infoc(ctx, "Created cluster", rlog.Any("result", result))
	return nil
}

func DeleteCluster(ctx context.Context, name, namespace string) error {
	resource, err := tanzurepo.GetResource(resourceDefinition, name, namespace)
	if err != nil {
		rlog.Errorc(ctx, "failed to fetch cluster resource", err)
		return err
	}

	err = tanzurepo.DeleteResource(resourceDefinition, *resource)
	if err != nil {
		rlog.Errorc(ctx, "failed to delete cluster", err)
		return err
	}

	return nil
}

func SetClusterControlPlaneToHA(ctx context.Context, name, namespace string) error {
	resource, err := tanzurepo.GetResource(resourceDefinition, name, namespace)
	if err != nil {
		rlog.Errorc(ctx, "failed to get cluster", err)
		return err
	}

	if err := unstructured.SetNestedField(resource.Object, int64(3), "spec", "topology", "controlPlane", "replicas"); err != nil {
		rlog.Errorc(ctx, "failed to set controlPlane replicas", err)
		return err
	}

	err = tanzurepo.UpdateResource(resourceDefinition, *resource)
	if err != nil {
		rlog.Errorc(ctx, "failed to update cluster", err)
		return err
	}

	return nil
}

func SetClusterNodepoolReplica(ctx context.Context, resourceName, resourceNamespace, nodepoolName string, replica int64) error {
	resource, err := tanzurepo.GetResource(resourceDefinition, resourceName, resourceNamespace)
	if err != nil {
		rlog.Errorc(ctx, "failed to get cluster", err)
		return err
	}

	nodepools, found, err := unstructured.NestedSlice(resource.Object, "spec", "topology", "nodePools")
	if err != nil || !found || nodepools == nil {
		rlog.Errorc(ctx, "failed to get nodepools from resource", err)
		return err
	}

	index := -1
	for i := 0; i < len(nodepools); i++ {
		if nodepools[i].(map[string]any)["name"] == nodepoolName {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("nodepool %s not found", nodepoolName)
	}

	if err := unstructured.SetNestedField(nodepools[index].(map[string]any), replica, "replicas"); err != nil {
		rlog.Errorc(ctx, "failed to set nested fields (replica)", err)
		return err
	}

	if err := unstructured.SetNestedField(resource.Object, nodepools, "spec", "topology", "nodePools"); err != nil {
		rlog.Errorc(ctx, "failed to set nested fields (nodePools)", err)
		return err
	}

	err = tanzurepo.UpdateResource(resourceDefinition, *resource)
	if err != nil {
		rlog.Errorc(ctx, "failed to update resource", err)
		return err
	}

	return nil
}
