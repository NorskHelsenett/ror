package tanzukubernetesservice

import (
	"context"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/mstanzuconnections"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/rorclient"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/tanzumsservice"
	"github.com/NorskHelsenett/ror/internal/factories/datacenterfactory"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func UpdateClusterOrderByTanzuKubernetesClusterChanges(ctx context.Context, clusterOrder apiresourcecontracts.ResourceClusterOrder, tanzuk8sCluster *apiresourcecontracts.ResourceTanzuKubernetesCluster) error {
	if tanzuk8sCluster.Status.Phase == "running" && clusterOrder.Status.Phase == "Done" && clusterOrder.Status.Status == "Running" {
		return nil
	} else if tanzuk8sCluster.Status.Phase == "failed" && clusterOrder.Status.Phase == "Error" {
		return nil
	}

	co, err := clusterorder.NewClusterOrderFromResource(ctx, clusterOrder)
	if err != nil {
		return err
	}

	if tanzuk8sCluster.Status.Phase == "running" {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseCreating
		clusterOrder.Status.Status = "Installing operator"

		providerConfig := co.GetProviderConfig()
		rlog.Infoc(ctx, "Getting providerconfig", rlog.Any("providerconfig", providerConfig))
		providerConfigTanzu, ok := providerConfig.(apiresourcecontracts.ResourceProviderConfigTanzu)
		if !ok {
			clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
			clusterOrder.Status.Status = "Invalid providerconfig"
			tanzumsservice.UpdateClusterOrder(ctx, &clusterOrder)
			err := errors.New("providerconfig is nil")
			rlog.Errorc(ctx, "providerconfig is nil", err)
			return err
		}

		err := createRorCluster(ctx, clusterOrder, tanzuk8sCluster, providerConfigTanzu)
		if err != nil {
			clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
			clusterOrder.Status.Status = "could not create cluster"
			tanzumsservice.UpdateClusterOrder(ctx, &clusterOrder)
			rlog.Errorc(ctx, "could not create ror cluster", err)
			return err
		}

		err = sendOperatorOrder(ctx, tanzuk8sCluster, clusterOrder, providerConfigTanzu)
		if err != nil {
			rlog.Errorc(ctx, "could not send operator order", err)
		}
	} else if tanzuk8sCluster.Status.Phase == "failed" {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "tanzu kubernetes cluster failed"
		tanzumsservice.UpdateClusterOrder(ctx, &clusterOrder)
	} else {
		rlog.Debugc(ctx, "tanzu kubernetes cluster status is not running or failed", rlog.Any("status", tanzuk8sCluster.Status.Phase))
		return nil
	}

	tanzumsservice.UpdateClusterOrder(ctx, &clusterOrder)
	extraheaders := map[string]interface{}{"apiVersion": clusterOrder.ApiVersion, "kind": clusterOrder.Kind}
	err = mstanzuconnections.RabbitMQConnection.SendMessage(ctx, clusterOrder, messagebuscontracts.Event_ClusterOrderUpdated, extraheaders)
	if err != nil {
		rlog.Errorc(ctx, "could not send message from ms-tanzu", err)
	}

	return nil
}

func sendOperatorOrder(ctx context.Context, tanzuk8sCluster *apiresourcecontracts.ResourceTanzuKubernetesCluster, clusterOrder apiresourcecontracts.ResourceClusterOrder,
	providerConfig apiresourcecontracts.ResourceProviderConfigTanzu) error {
	rlog.Debugc(ctx, "Cluster order:")
	stringhelper.PrettyprintStruct(clusterOrder)
	rlog.Debugc(ctx, "Tkc:")
	stringhelper.PrettyprintStruct(tanzuk8sCluster)

	datacenter, err := rorclient.RorClient.Datacenters().GetById(providerConfig.DatacenterId)
	if err != nil {
		rlog.Errorc(ctx, "could not get datacenter", err)
		return err
	}

	if datacenter == nil {
		err := errors.New("datacenter is nil")
		rlog.Errorc(ctx, "datacenter is nil", err)
		return nil
	}

	stringhelper.PrettyprintStruct(datacenter)

	extraheaders := map[string]interface{}{"datacenter": datacenterfactory.DatacenterToTanzu(datacenter.Name)}
	payload := messagebuscontracts.OperatorOrder{
		TanzuKubernetesCluster: *tanzuk8sCluster,
	}
	err = mstanzuconnections.RabbitMQConnection.SendMessage(ctx, payload, messagebuscontracts.Route_ProviderTanzuOperatorOrder, extraheaders)
	if err != nil {
		rlog.Errorc(ctx, "could not send message from ms-tanzu", err)
		return err
	}

	return nil
}

func createRorCluster(ctx context.Context,
	clusterOrder apiresourcecontracts.ResourceClusterOrder,
	tanzuk8sCluster *apiresourcecontracts.ResourceTanzuKubernetesCluster,
	providerConfig apiresourcecontracts.ResourceProviderConfigTanzu) error {

	rlog.Infoc(ctx, "Creating cluster",
		rlog.Any("clusterName", clusterOrder.Spec.Cluster),
		rlog.Any("datacenterId", providerConfig.DatacenterId),
		rlog.Any("namespace", tanzuk8sCluster.Metadata.Namespace))
	clusterId, err := rorclient.RorClient.Clusters().Create(apicontracts.Cluster{
		ClusterName: clusterOrder.Spec.Cluster,
		WorkspaceId: providerConfig.NamespaceId,
		Workspace: apicontracts.Workspace{
			DatacenterID: providerConfig.DatacenterId,
			Name:         tanzuk8sCluster.Metadata.Namespace,
			Datacenter: apicontracts.Datacenter{
				Provider: providers.ProviderTypeTanzu,
			},
		},
		Metadata: apicontracts.ClusterMetadata{
			ProjectID: clusterOrder.Spec.ProjectId,
		},
		Topology: apicontracts.Topology{
			NodePools: []apicontracts.NodePool{
				{
					Name: "worker",
				},
			},
		},
	})
	if err != nil {
		rlog.Errorc(ctx, "could not create cluster", err)
		return err
	}

	event := messagebuscontracts.ClusterCreatedEvent{}
	event.ClusterId = clusterId
	event.WorkspaceName = tanzuk8sCluster.Metadata.Namespace
	event.ClusterName = clusterOrder.Spec.Cluster

	err = mstanzuconnections.RabbitMQConnection.SendMessage(ctx, event, messagebuscontracts.Route_Cluster_Created, nil)
	if err != nil {
		rlog.Errorc(ctx, "could not send cluster created event", err, rlog.String("clusterId", clusterId))
	}

	return nil
}
