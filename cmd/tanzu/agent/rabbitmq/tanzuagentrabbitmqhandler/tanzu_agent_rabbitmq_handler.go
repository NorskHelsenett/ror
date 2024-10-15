package tanzuagentrabbitmqhandler

import (
	"context"
	"encoding/json"
	clusterservicev1alpha2 "github.com/NorskHelsenett/ror/cmd/tanzu/agent/clusterserviceV1alpha2"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/rabbitmq/tanzuagentrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/services/roroperatorservice"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/tanzuagentconnections"
	"github.com/NorskHelsenett/ror/internal/factories/datacenterfactory"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    tanzuagentconnections.RabbitMQConnection,
			QueueName: tanzuagentrabbitmqdefinitions.TanzuAgentQueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbithandler := rabbitmqhandler.New(config, tanzuagentmessagehandler{})
		_ = tanzuagentconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
	}()
}

type tanzuagentmessagehandler struct {
}

func (tamh tanzuagentmessagehandler) HandleMessage(ctx context.Context, message amqp091.Delivery) error {
	messageDatacenterName := message.Headers["datacenter"]
	datacenterName := datacenterfactory.DatacenterToTanzu(messageDatacenterName.(string))
	tanzuAgentDatacenterName := viper.GetString(configconsts.TANZU_AGENT_DATACENTER)
	if datacenterName != tanzuAgentDatacenterName {
		rlog.Errorc(ctx, "Received message, but not handling it ... ", nil, rlog.Any("route", message.RoutingKey))
		return nil
	}

	rlog.Infoc(ctx, "Received message", rlog.Any("route", message.RoutingKey))
	switch message.RoutingKey {
	case messagebuscontracts.Route_ProviderTanzuClusterCreate:
		return handleClusterCreate(ctx, message)
	case messagebuscontracts.Route_ProviderTanzuClusterModify:
		return handleClusterModify(ctx, message)
	case messagebuscontracts.Route_ProviderTanzuClusterDelete:
		return handleClusterDelete(ctx, message)
	case messagebuscontracts.Route_ProviderTanzuOperatorOrder:
		return handleOperatorOrder(ctx, message)
	default:
		rlog.Errorc(ctx, "Received message, but not handling it ... ", nil, rlog.Any("route", message.RoutingKey))
	}
	return nil
}

func handleClusterCreate(ctx context.Context, message amqp091.Delivery) error {
	rlog.Debugc(ctx, "Received message", rlog.Any("route", message.RoutingKey))

	var event messagebuscontracts.ClusterCreate
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		rlog.Error("could not convert to json", err)
	}

	stringhelper.PrettyprintStruct(event)

	err = clusterservicev1alpha2.CreateCluster(ctx, event.ClusterInput)
	if err != nil {
		rlog.Error("could not create cluster", err)
		return err
	}
	return nil
}

func handleClusterModify(ctx context.Context, message amqp091.Delivery) error {
	rlog.Debugc(ctx, "Received message", rlog.Any("route", message.RoutingKey))

	var event messagebuscontracts.ClusterModify
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		rlog.Error("could not convert to json", err)
		return err
	}

	stringhelper.PrettyprintStruct(event)
	return nil
}

func handleClusterDelete(ctx context.Context, message amqp091.Delivery) error {
	rlog.Debugc(ctx, "Received message", rlog.Any("route", message.RoutingKey))

	var event messagebuscontracts.ClusterDelete
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		rlog.Error("could not convert to json", err)
		return err
	}

	stringhelper.PrettyprintStruct(event)
	return nil
}

func handleOperatorOrder(ctx context.Context, message amqp091.Delivery) error {
	rlog.Debugc(ctx, "Received message", rlog.Any("route", message.RoutingKey))

	var event messagebuscontracts.OperatorOrder
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		rlog.Errorc(ctx, "could not convert to json", err)
		return err
	}

	stringhelper.PrettyprintStruct(event)

	if event.TanzuKubernetesCluster.Metadata.Namespace == "" || event.TanzuKubernetesCluster.Metadata.Name == "" {
		rlog.Errorc(ctx, "could not install ror operator, missing tkc name and tkc namespace", err)
		return err
	}

	err = roroperatorservice.InstallRorOperatorInCluster(context.Background(), event.TanzuKubernetesCluster.Metadata.Namespace, event.TanzuKubernetesCluster.Metadata.Name)
	if err != nil {
		rlog.Errorc(ctx, "could not install ror operator", err)
		return err
	}
	return nil
}
