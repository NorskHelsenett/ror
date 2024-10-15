package mskindrabbitmqhandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/kind/mskindconnections"
	"github.com/NorskHelsenett/ror/cmd/kind/rabbitmq/mskindrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/kind/rorclient"
	"github.com/NorskHelsenett/ror/cmd/kind/services/kindservice"

	"strings"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"
	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
	"github.com/rabbitmq/amqp091-go"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

const (
	rorInternalApiVersion = "general.ror.internal/v1alpha1"
	orderKind             = "ClusterOrder"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    mskindconnections.RabbitMQConnection,
			QueueName: mskindrabbitmqdefinitions.MsKindQueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbithandler := rabbitmqhandler.New(config, mskindmessagehandler{})
		err := mskindconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
		if err != nil {
			rlog.Fatal("could not register handler", err)
		}
	}()
}

type mskindmessagehandler struct {
}

func (kmh mskindmessagehandler) HandleMessage(ctx context.Context, message amqp091.Delivery) error {
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
			if strings.Contains(err.Error(), "404") {
				rlog.Errorc(ctx, "resource not found", err)
				return nil
			}
			return err
		}

		if clusterOrder.Spec.Provider != providers.ProviderTypeKind {
			err := errors.New("provider not supported by this micro service")
			rlog.Errorc(ctx, "wrong provider", err, rlog.Any("provider", clusterOrder.Spec.Provider))
			return err
		}
		if message.RoutingKey == messagebuscontracts.Route_ResourceCreated {
			handleClusterOrder(ctx, clusterOrder, message)
		}
	}

	return nil
}

func handleClusterOrder(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder, message amqp091.Delivery) {
	rlog.Debugc(ctx, "Received message", rlog.Any("route", message.RoutingKey))
	if clusterOrder == nil {
		rlog.Errorc(ctx, "clusterOrder is nil", nil)
		return
	}
	switch clusterOrder.Spec.OrderType {
	case apiresourcecontracts.ResourceActionTypeCreate:
		err := kindservice.ClusterOrderToClusterCreate(ctx, clusterOrder)
		if err != nil {
			rlog.Error("could not ack message", err)
		}
	case apiresourcecontracts.ResourceActionTypeUpdate:
		err := kindservice.ClusterOrderToClusterUpdate(ctx, clusterOrder)
		if err != nil {
			rlog.Error("could not ack message", err)
		}
	case apiresourcecontracts.ResourceActionTypeDelete:
		err := kindservice.ClusterOrderToClusterDelete(ctx, clusterOrder)
		if err != nil {
			rlog.Error("could not ack message", err)
		}
	}
}
