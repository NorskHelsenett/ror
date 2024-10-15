package msauditrabbitmqhandler

import (
	"context"
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/audit/auditservice"
	"github.com/NorskHelsenett/ror/cmd/audit/msauditconnections"
	"github.com/NorskHelsenett/ror/cmd/audit/rabbitmq/msauditrabbitmqdefinitions"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    msauditconnections.RabbitMQConnection,
			QueueName: msauditrabbitmqdefinitions.QueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbithandler := rabbitmqhandler.New(config, auditmessagehandler{})
		err := msauditconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
		if err != nil {
			rlog.Fatal("could not register handler", err)
		}
	}()
}

type auditmessagehandler struct {
}

func (amh auditmessagehandler) HandleMessage(ctx context.Context, message amqp.Delivery) error {
	var event messagebuscontracts.AclUpdateEvent
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		rlog.Error("could not convert to json", err)
	}

	auditservice.CreateAndCommitAclList(ctx, event)
	return nil
}
