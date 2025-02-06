package rabbitmqclient

import (
	"context"
	"encoding/json"
	"errors"
	"maps"

	"github.com/NorskHelsenett/ror/pkg/telemetry/trace"

	amqp "github.com/rabbitmq/amqp091-go"
)

var mandatory = false

func (rc rabbitmqcon) SendMessage(ctx context.Context, message any, routing string, extraheaders map[string]interface{}) error {
	ctx, span := rc.Trace(ctx, "rabbitmqclient: SendMessage")
	defer span.End()

	if message == nil {
		return errors.New("could not send message")
	}

	_, span2 := rc.Trace(ctx, "Prepare payload")
	defer span2.End()
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return errors.New("could not marshal message")
	}
	span2.End()

	if len(messageBytes) == 0 {
		return errors.New("missing message content")
	}

	ctx, span5 := rc.Trace(ctx, "publish message")
	defer span5.End()

	mandatory = true
	headers := trace.InjectAMQPHeaders(ctx)
	maps.Copy(headers, extraheaders)

	err = rc.RabbitMqChannel.PublishWithContext(context.TODO(), rc.SenderQueName, routing, mandatory, false, amqp.Publishing{
		ContentType: "application/json",
		Headers:     headers,
		Body:        messageBytes,
	})
	if err != nil {
		return err
	}

	span5.End()
	span.End()

	return nil
}
