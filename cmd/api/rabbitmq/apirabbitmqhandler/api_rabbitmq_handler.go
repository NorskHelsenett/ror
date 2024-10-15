package apirabbitmqhandler

import (
	"context"
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"github.com/NorskHelsenett/ror/cmd/api/rabbitmq/apirabbitmqdefinitions"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/rabbitmq/amqp091-go"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    apiconnections.RabbitMQConnection,
			QueueName: apirabbitmqdefinitions.ApiEventsQueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbithandler := rabbitmqhandler.New(config, apimessagehandler{})
		_ = apiconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
	}()
}

type apimessagehandler struct {
}

func (amh apimessagehandler) HandleMessage(ctx context.Context, message amqp091.Delivery) error {
	switch message.RoutingKey {
	case messagebuscontracts.Route_ResourceCreated:
		var event apiresourcecontracts.ResourceUpdateModel
		err := json.Unmarshal(message.Body, &event)
		if err != nil {
			rlog.Error("could not convert to json", err)
			return err
		}

		if event.ApiVersion == "general.ror.internal/v1alpha1" && event.Kind == "ClusterOrder" {
			err := HandleClusterOrderResource(ctx, message)
			if err != nil {
				return err
			}
		}
		return nil
	case messagebuscontracts.Event_Broadcast:
		err := HandleEvents(ctx, message)
		if err != nil {
			rlog.Error("could not handle event", err)
			return err
		}
	default:
		rlog.Debugc(ctx, "could not handle message")
	}

	return nil
}
