package tanzumsservice

import (
	"context"
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/mstanzuconnections"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/rorclient"
	"github.com/NorskHelsenett/ror/internal/factories/datacenterfactory"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/utils"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"
	"github.com/NorskHelsenett/ror/pkg/models/providers/tanzu"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

const (
	controlPlaneMachineClass = "best-effort-medium"
)

func ClusterOrderToClusterCreate(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	rlog.Debugc(ctx, "Create Cluster order to cluster specs")

	var providerConfigTanzu apiresourcecontracts.ResourceProviderConfigTanzu
	jsonString, _ := json.Marshal(clusterOrder.Spec.ProviderConfig)
	err := json.Unmarshal(jsonString, &providerConfigTanzu)
	if err != nil {
		rlog.Error("could not cast to tanzuProviderConfig", err)
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "could not convert to tanzuProviderConfig"
		UpdateClusterOrder(ctx, clusterOrder)
		rlog.Errorc(ctx, "could not cast to tanzuProviderConfig", err)
		return nil
	}

	workspace, err := rorclient.RorClient.Workspaces().GetById(providerConfigTanzu.NamespaceId)
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "workspace not found"
		UpdateClusterOrder(ctx, clusterOrder)
		rlog.Errorc(ctx, "workspace not found", err)
		return err
	}

	datacenter, err := rorclient.RorClient.Datacenters().GetById(providerConfigTanzu.DatacenterId)
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "datacenter not found"
		UpdateClusterOrder(ctx, clusterOrder)
		rlog.Errorc(ctx, "datacenter not found", err)
		return err
	}

	count := 1
	nodePools := make([]tanzu.NodePool, 0)
	for _, nodePool := range clusterOrder.Spec.NodePools {
		tanzuNodePool := tanzu.NodePool{
			Name:              nodePool.Name,
			KubernetesVersion: clusterOrder.Spec.K8sVersion,
			Replicas:          int64(nodePool.Count),
			VmClass:           nodePool.MachineClass,
		}
		nodePools = append(nodePools, tanzuNodePool)
		count = count + 1
	}

	clusterInput := tanzu.TanzuKubernetesClusterInput{
		Name:       clusterOrder.Spec.Cluster,
		Namespace:  workspace.Name,
		DataCenter: datacenter.Name,
		ControlPlane: tanzu.ControlPlane{
			HighAvailability:  clusterOrder.Spec.HighAvailability,
			KubernetesVersion: clusterOrder.Spec.K8sVersion, //todo fetch from spec
			VmClass:           controlPlaneMachineClass,     //todo fetch from spec
		},
		NodePools: nodePools,
	}

	payload := messagebuscontracts.ClusterCreate{
		ClusterInput: clusterInput,
	}

	extraHeaders := map[string]interface{}{"datacenter": datacenterfactory.DatacenterToTanzu(datacenter.Name)}
	err = mstanzuconnections.RabbitMQConnection.SendMessage(ctx, payload, messagebuscontracts.Route_ProviderTanzuClusterCreate, extraHeaders)
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "could not send message from ms-tanzu"
		UpdateClusterOrder(ctx, clusterOrder)
		rlog.Errorc(ctx, "could not send message from ms-tanzu", err)
		return err
	}

	clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseCreating
	clusterOrder.Status.Status = "Provisioning"
	UpdateClusterOrder(ctx, clusterOrder)

	return nil
}

func ClusterOrderToClusterUpdate(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	rlog.Debugc(ctx, "Create Cluster order to cluster specs")
	return nil
}

func ClusterOrderToClusterDelete(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	rlog.Debugc(ctx, "Delete Cluster order to cluster specs")
	return nil
}

func UpdateClusterOrder(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) {
	if clusterOrder.Metadata.Uid == "" {
		rlog.Error("cluster order uid is empty", nil)
		return
	}

	updateModel, err := utils.NewResourceUpdate(ctx, *clusterOrder)
	if err != nil {
		rlog.Error("failed to create update model", err)
		return
	}

	err = rorclient.RorClient.Resources().UpdateClusterOrder(&updateModel)
	if err != nil {
		rlog.Error("failed to update cluster order", err)
		return
	}
}
