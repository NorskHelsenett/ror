package mstanzurabbitmqhandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/mstanzuconnections"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/rabbitmq/mstanzurabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/rorclient"
	tanzukubernetesservice "github.com/NorskHelsenett/ror/cmd/tanzu/ms/tanzukubernetesService"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/tanzumsservice"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"
	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"github.com/rabbitmq/amqp091-go"
)

const (
	rorInternalApiVersion     = "general.ror.internal/v1alpha1"
	orderKind                 = "ClusterOrder"
	tanzuK8sClusterApiVersion = "run.tanzu.vmware.com/v1alpha2"
	tanzuK8sClusterKind       = "TanzuKubernetesCluster"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    mstanzuconnections.RabbitMQConnection,
			QueueName: mstanzurabbitmqdefinitions.MsTanzuQueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbitmqHandler := rabbitmqhandler.New(config, tanzuMessageHandler{})
		err := mstanzuconnections.RabbitMQConnection.RegisterHandler(rabbitmqHandler)
		if err != nil {
			rlog.Error("could not register handler", err)
		}
	}()
}

type tanzuMessageHandler struct {
}

func (tmh tanzuMessageHandler) HandleMessage(ctx context.Context, message amqp091.Delivery) error {
	var event apiresourcecontracts.ResourceUpdateModel

	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		rlog.Error("could not convert to json", err)
		return err
	}

	if event.Version == apiresourcecontracts.ResourceVersionV2 {
		errMsg := "resourcev2 is not supported"
		rlog.Warnc(ctx, errMsg)
		return errors.New(errMsg)
	}

	if event.ApiVersion == rorInternalApiVersion && event.Kind == orderKind {
		clusterOrder, err := rorclient.RorClient.Resources().GetClusterOrderByUid(event.Uid,
			aclmodels.Acl2Subject(event.Owner.Subject),
			event.Owner.Scope)
		if err != nil {
			rlog.Errorc(ctx, "could not get resource", err)
			return err
		}

		if clusterOrder.Spec.Provider != providers.ProviderTypeTanzu {
			err := errors.New("provider not supported by this micro service")
			rlog.Debugc(ctx, "wrong provider", rlog.Any("provider", clusterOrder.Spec.Provider))
			return err
		}
		if message.RoutingKey == messagebuscontracts.Route_ResourceCreated {
			handleClusterOrder(ctx, clusterOrder, message)
		}
	}

	if event.ApiVersion == tanzuK8sClusterApiVersion &&
		event.Kind == tanzuK8sClusterKind &&
		(message.RoutingKey == messagebuscontracts.Route_ResourceUpdated ||
			message.RoutingKey == messagebuscontracts.Route_ResourceCreated) {
		handleTanzuK8sClusterChanges(ctx, event, message)
	}

	return nil
}

func handleTanzuK8sClusterChanges(ctx context.Context, event apiresourcecontracts.ResourceUpdateModel, _ amqp091.Delivery) {
	if event.Owner.Scope == "" || event.Owner.Subject == "" {
		rlog.Errorc(ctx, "owner scope or subject is empty", nil)
		return
	}

	tanzuK8sCluster, err := rorclient.RorClient.Resources().GetTanzuKubernetesClusterByUid(event.Uid, event.Owner.Subject, event.Owner.Scope)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource tkc", err, rlog.String("uid", event.Uid))
		return
	}

	if tanzuK8sCluster == nil {
		rlog.Errorc(ctx, "tkc is nil", nil)
		return
	}

	clusterOrderOwner := rortypes.RorResourceOwnerReference{
		Scope:   aclmodels.Acl2ScopeRor,
		Subject: aclmodels.Acl2RorSubjectGlobal,
	}
	clusterOrders, err := rorclient.RorClient.Resources().GetClusterOrders(clusterOrderOwner.Subject, clusterOrderOwner.Scope)
	if err != nil {
		rlog.Errorc(ctx, "could not get cluster order resources", err)
		return
	}

	if len(clusterOrders) == 0 {
		rlog.Debugc(ctx, "no cluster orders found")
		return
	}

	// TODO - this is a hack, we need to find a better way to get the cluster order
	var clusterOrder apiresourcecontracts.ResourceClusterOrder
	found := false
	for _, co := range clusterOrders {
		if len(co.Spec.Cluster) == 0 {
			continue
		}
		if tanzuK8sCluster.Metadata.Name == co.Spec.Cluster {
			clusterOrder = *co
			found = true
			break
		}
	}

	if !found {
		return
	}

	rlog.Infoc(ctx, "tkc created or updated, syncing cluster order and tkc",
		rlog.String("tkc cluster name", tanzuK8sCluster.Metadata.Name),
		rlog.String("cluster order clustername", clusterOrder.Spec.Cluster))
	err = tanzukubernetesservice.UpdateClusterOrderByTanzuKubernetesClusterChanges(ctx, clusterOrder, tanzuK8sCluster)
	if err != nil {
		rlog.Errorc(ctx, "could not update cluster order", err)
	}
}

func handleClusterOrder(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder, message amqp091.Delivery) {
	rlog.Debugc(ctx, "Received message", rlog.Any("route", message.RoutingKey))
	switch clusterOrder.Spec.OrderType {
	case apiresourcecontracts.ResourceActionTypeCreate:
		err := tanzumsservice.ClusterOrderToClusterCreate(ctx, clusterOrder)
		if err != nil {
			rlog.Error("could not ack message", err)
		}
	case apiresourcecontracts.ResourceActionTypeUpdate:
		err := tanzumsservice.ClusterOrderToClusterUpdate(ctx, clusterOrder)
		if err != nil {
			rlog.Error("could not ack message", err)
		}
	case apiresourcecontracts.ResourceActionTypeDelete:
		err := tanzumsservice.ClusterOrderToClusterDelete(ctx, clusterOrder)
		if err != nil {
			rlog.Error("could not ack message", err)
		}

	}

}
