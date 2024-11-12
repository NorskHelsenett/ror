package apirabbitmqhandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/api/models/ssemodels"
	"github.com/NorskHelsenett/ror/cmd/api/webserver/sse"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/rabbitmq/amqp091-go"
)

func HandleClusterOrderResource(ctx context.Context, message amqp091.Delivery) error {
	var resourceUpdateModel apiresourcecontracts.ResourceUpdateModel
	err := json.Unmarshal(message.Body, &resourceUpdateModel)
	if err != nil {
		return err
	}

	payload := ssemodels.SseMessage{
		SSEBase: ssemodels.SSEBase{
			Event: ssemodels.SseType_ClusterOrder_Updated,
		},
		Data: resourceUpdateModel,
	}

	sse.Server.BroadcastMessage(payload)
	return nil
}

func HandleEvents(ctx context.Context, message amqp091.Delivery) error {
	if message.Body == nil {
		return errors.New("message.body is nil")
	}

	var sseEvent ssemodels.SseMessage
	err := json.Unmarshal(message.Body, &sseEvent)
	if err != nil {
		return err
	}

	sse.Server.BroadcastMessage(sseEvent)
	return nil
}
