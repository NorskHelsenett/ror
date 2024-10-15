package msauthrabbitmqhandler

import (
	"context"
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/auth/dex"
	"github.com/NorskHelsenett/ror/cmd/auth/msauthconnections"
	"github.com/NorskHelsenett/ror/cmd/auth/rabbitmq/msauthrabbitmqdefintions"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    msauthconnections.RabbitMQConnection,
			QueueName: msauthrabbitmqdefintions.QueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbithandler := rabbitmqhandler.New(config, authmessagehandler{})
		err := msauthconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
		if err != nil {
			rlog.Error("could not register handler", err)
		}
	}()
}

type authmessagehandler struct{}

func (amh authmessagehandler) HandleMessage(ctx context.Context, message amqp.Delivery) error {
	var event messagebuscontracts.ClusterCreatedEvent
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		rlog.Error("could not convert to json", err)
		return err
	}

	rlog.Info("event: cluster created: clusterId", rlog.String("cluster id", event.ClusterId))

	err = dex.HandleClusterCreated(ctx, &message)
	if err != nil {
		rlog.Error("failed handling message", err)
		return err
	}
	return nil
}
