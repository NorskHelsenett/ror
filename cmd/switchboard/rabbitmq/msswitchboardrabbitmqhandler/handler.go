package msswitchboardrabbitmqhandler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/cmd/switchboard/rabbitmq/msswitchboardrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/switchboard/ror"
	"github.com/NorskHelsenett/ror/cmd/switchboard/slack"
	"github.com/NorskHelsenett/ror/cmd/switchboard/switchboardconnections"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	amqp "github.com/rabbitmq/amqp091-go"
)

func StartListening() {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    switchboardconnections.RabbitMQConnection,
			QueueName: msswitchboardrabbitmqdefinitions.QueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args: amqp.Table{
				"kind": "VulnerabilityEvent",
			},
		}
		rabbithandler := rabbitmqhandler.New(config, switchboardhandler{})
		err := switchboardconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
		if err != nil {
			rlog.Fatal("could not register handler", err)
		}
	}()
}

type switchboardhandler struct {
}

func (h switchboardhandler) HandleMessage(ctx context.Context, message amqp.Delivery) error {
	var resource apiresourcecontracts.ResourceUpdateModel
	err := json.Unmarshal(message.Body, &resource)
	if err != nil {
		return err
	}

	if resource.Version == apiresourcecontracts.ResourceVersionV1 {
		errMsg := "resourcev1 is not supported"
		rlog.Warnc(ctx, errMsg)
		return errors.New(errMsg)
	}

	if resource.Kind == "VulnerabilityEvent" {
		owner := rortypes.RorResourceOwnerReference{
			Subject: aclmodels.Acl2Subject(resource.Owner.Subject),
			Scope:   resource.Owner.Scope,
		}
		rlog.Debugc(ctx, "getting routes from ror api")
		routes, err := getRoutes(ctx, owner)
		if err != nil || routes == nil {
			rlog.Errorc(ctx, "unable to get routes from ror api, aborting messaging", err)
			return err
		}
		vulnerabilityEvent, err := getVulnerabilityEvent(ctx, resource.Uid, owner)
		if err != nil || vulnerabilityEvent == nil {
			rlog.Errorc(ctx, "unable to get vulnerability event from ror api, aborting messaging", err)
			return err
		}
		rlog.Debugc(ctx, fmt.Sprintf("number of routes found: %d", len(routes)))
		for _, route := range routes {
			if route.Spec.MessageType.ApiVersion == "general.ror.internal/v1alpha1" && route.Spec.MessageType.Kind == "VulnerabilityEvent" {
				for _, slackReceiver := range route.Spec.Receivers.Slack {
					rlog.Infoc(ctx, fmt.Sprintf("creating slack message for owner %v", owner))
					err = slack.CreateSlackMessage(ctx, slackReceiver.ChannelId, vulnerabilityEvent.Spec.Message, owner)
					if err != nil {
						rlog.Errorc(ctx, "unable to send slack message to ror api", err)
					}
				}
			}
		}

	}

	return nil
}

func getRoutes(ctx context.Context, owner rortypes.RorResourceOwnerReference) ([]rortypes.ResourceRoute, error) {
	query := rorresources.NewResourceQuery()
	query.VersionKind.Kind = "Route"
	query.VersionKind.Version = "general.ror.internal/v1alpha1"
	query.OwnerRefs = make([]rortypes.RorResourceOwnerReference, 0)
	query.OwnerRefs = append(query.OwnerRefs, owner)
	resourceSet, err := ror.Client.ResourceV2().Get(ctx, *query)
	if err != nil {
		return nil, err
	}
	routes := []rortypes.ResourceRoute{}
	resources := resourceSet.GetAll()
	for _, r := range resources {
		routes = append(routes, *r.Route().Get())
	}
	return routes, nil
}

func getVulnerabilityEvent(ctx context.Context, uid string, owner rortypes.RorResourceOwnerReference) (*rortypes.ResourceVulnerabilityEvent, error) {
	query := rorresources.NewResourceQuery()
	query.VersionKind.Kind = "VulnerabilityEvent"
	query.VersionKind.Version = "general.ror.internal/v1alpha1"
	query.OwnerRefs = make([]rortypes.RorResourceOwnerReference, 0)
	query.OwnerRefs = append(query.OwnerRefs, owner)
	query.WithUID(uid)
	resourceSet, err := ror.Client.ResourceV2().Get(ctx, *query)
	if err != nil {
		return nil, err
	}
	vulnerabilityEvent := resourceSet.Get().VulnerabilityEvent().Get()
	return vulnerabilityEvent, nil
}
