package msnhnrabbitmqhandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/nhn/msnhnconnections"
	"github.com/NorskHelsenett/ror/cmd/nhn/rabbitmq/msnhnrabbitmqdefinitions"

	"github.com/NorskHelsenett/ror/pkg/context/mscontext"

	"github.com/NorskHelsenett/ror/cmd/nhn/services/applicationservice"
	"github.com/NorskHelsenett/ror/cmd/nhn/settings"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.opentelemetry.io/otel"
)

// StartListening starts listening to rabbitmq queue
func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    msnhnconnections.RabbitMQConnection,
			QueueName: msnhnrabbitmqdefinitions.QueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args:      nil,
		}
		rabbithandler := rabbitmqhandler.New(config, nhnmessagehandler{})
		_ = msnhnconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
	}()
}

type nhnmessagehandler struct {
}

func (nmh nhnmessagehandler) HandleMessage(ctx context.Context, message amqp.Delivery) error {
	headers := message.Headers
	apiversion := headers["apiVersion"]
	kind := headers["kind"]

	if message.RoutingKey == messagebuscontracts.Route_Cluster_Created {
		err := handleClusterCreated(ctx, message)
		if err != nil {
			rlog.Errorc(ctx, "Could not handle cluster created", err, rlog.Any("event", message))
		}
	} else if apiversion == "argoproj.io/v1alpha1" && kind == "Application" {
		err := handleApplicationChanges(context.Background(), message)
		if err != nil {
			rlog.Errorc(ctx, "Could not handle application changes", err, rlog.Any("event", message))
		}
	}

	return nil
}

func handleClusterCreated(ctx context.Context, message amqp.Delivery) error {
	tracer := otel.Tracer("listener")
	ctx, span := tracer.Start(ctx, "Processed cluster created message")
	defer span.End()

	var payload messagebuscontracts.ClusterCreatedEvent
	err := json.Unmarshal(message.Body, &payload)
	if err != nil {
		rlog.Error("Could not convert to json", err)
		return err
	}

	rlog.Infoc(ctx, "Cluster created", rlog.String("cluster id", payload.ClusterId))

	serviceIdentityCtx, _ := mscontext.GetRorContextFromServiceContext(&ctx, settings.ServiceName)
	err = msnhnconnections.RabbitMQConnection.SendMessage(serviceIdentityCtx, payload, messagebuscontracts.Route_Auth, nil)
	if err != nil {
		rlog.Error("Could not send event to ror.auth", err)
	}

	span.End()
	return nil
}

func handleApplicationChanges(ctx context.Context, message amqp.Delivery) error {
	tracer := otel.Tracer("listener")
	ctx, span := tracer.Start(ctx, "Processed message")
	defer span.End()

	var app apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]
	err := json.Unmarshal(message.Body, &app)
	if err != nil {
		rlog.Error("Could not convert to json", err)
		return err
	}

	if app.Version == apiresourcecontracts.ResourceVersionV2 {
		errMsg := "resourcev2 is not supported"
		rlog.Warnc(ctx, errMsg)
		return errors.New(errMsg)
	}

	if app.Resource.Metadata.Name != "argocd" {
		return nil
	}

	serviceIdentityCtx, _ := mscontext.GetRorContextFromServiceContext(&ctx, settings.ServiceName)
	err = applicationservice.ProcessResourceUpdatedEvent(serviceIdentityCtx, app)
	if err != nil {
		return err
	}

	span.End()
	return nil
}
