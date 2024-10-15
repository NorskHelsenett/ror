package msslackrabbitmqhandler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/slack/rabbitmq/msslackrabbitmqdefinitions"
	"github.com/NorskHelsenett/ror/cmd/slack/ror"
	"github.com/NorskHelsenett/ror/cmd/slack/slackconnections"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/handlers/rabbitmqhandler"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/slack-go/slack"
)

func StartListening(slackClient *slack.Client) {
	go func() {
		config := rabbitmqhandler.RabbitMQListnerConfig{
			Client:    slackconnections.RabbitMQConnection,
			QueueName: msslackrabbitmqdefinitions.QueueName,
			Consumer:  "",
			AutoAck:   false,
			Exclusive: false,
			NoLocal:   false,
			NoWait:    false,
			Args: amqp.Table{
				"kind": "SlackMessage",
			},
		}
		slackhandler := slackhandler{slackClient: slackClient}
		rabbithandler := rabbitmqhandler.New(config, slackhandler)
		err := slackconnections.RabbitMQConnection.RegisterHandler(rabbithandler)
		if err != nil {
			rlog.Fatal("could not register handler", err)
		}
	}()
}

type slackhandler struct {
	slackClient *slack.Client
}

func (h slackhandler) HandleMessage(ctx context.Context, message amqp.Delivery) error {
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

	// this should be guaranteed by rabbitmq setup
	if resource.Kind == "SlackMessage" {
		owner := rortypes.RorResourceOwnerReference{
			Subject: aclmodels.Acl2Subject(resource.Owner.Subject),
			Scope:   resource.Owner.Scope,
		}
		sm, err := getSlackMessage(resource.Uid, owner)
		if err != nil {
			rlog.Errorc(ctx, "unable to get message from ror api, aborting", err)
			return err
		}
		_, _, err = h.slackClient.PostMessageContext(ctx, sm.Spec.ChannelId, slack.MsgOptionText(sm.Spec.Message, false), slack.MsgOptionAsUser(true))
		if err != nil {
			rlog.Errorc(ctx, "unable to post message to slack", err)
		}
	}

	return nil
}

func getSlackMessage(uid string, owner rortypes.RorResourceOwnerReference) (*rortypes.ResourceSlackMessage, error) {
	query := rorresources.NewResourceQuery()
	query.VersionKind.Kind = "SlackMessage"
	query.VersionKind.Version = "general.ror.internal/v1alpha1"
	query.OwnerRefs = make([]rortypes.RorResourceOwnerReference, 0)
	query.OwnerRefs = append(query.OwnerRefs, owner)
	query.WithUID(uid)
	resourceSet, err := ror.Client.ResourceV2().Get(*query)
	if err != nil {
		return nil, err
	}
	slackMessage := resourceSet.Get().SlackMessage().Get()
	return slackMessage, nil
}
