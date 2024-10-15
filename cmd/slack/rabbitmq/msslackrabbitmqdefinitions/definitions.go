package msslackrabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/slack/slackconnections"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	QueueName    = messagebuscontracts.Workqueue_Slack_Message_Create
	MsSlackQueue amqp.Queue
)

func InitOrDie() {
	queueArgs := amqp.Table{
		amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		"kind":            "SlackMessage",
	}
	var err error
	MsSlackQueue, err = slackconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		QueueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		queueArgs, // arguments
	)
	if err != nil {
		args := [...]any{QueueName, err}
		msg := fmt.Sprintf("could not declare queue %s,", args)
		rlog.Fatal(msg, err)
	}

	routingKey := messagebuscontracts.Workqueue_Slack_Message_Create
	err = slackconnections.RabbitMQConnection.GetChannel().QueueBind(
		QueueName,                                // queue name
		routingKey,                               // routing key
		messagebuscontracts.ExchangeRorResources, // exchange
		false,
		amqp.Table{
			"kind": "SlackMessage",
		},
	)
	if err != nil {
		args := [...]any{QueueName, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
