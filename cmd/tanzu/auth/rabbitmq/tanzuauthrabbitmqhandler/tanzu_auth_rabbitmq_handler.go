package tanzuauthrabbitmqhandler

import (
	"context"
	"github.com/NorskHelsenett/ror/cmd/tanzu/auth/rabbitmq/tanzuauthrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/tanzu/auth/tanzuauthconnections"

	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"

	"github.com/rabbitmq/amqp091-go"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    tanzuauthconnections.RabbitMQConnection,
			QueueName: tanzuauthrabbitmqdefinitions.TanzuAuthQueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbithandler := rabbitmqhandler.New(config, tanzuauthmessagehandler{})
		_ = tanzuauthconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
	}()
}

type tanzuauthmessagehandler struct {
}

func (tamh tanzuauthmessagehandler) HandleMessage(ctx context.Context, message amqp091.Delivery) error {
	return nil
}
