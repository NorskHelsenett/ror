package msauditrabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/audit/msauditconnections"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/rabbitmq/amqp091-go"
)

var (
	QueueName    = "ms-audit"
	MsAuditQueue amqp091.Queue
)

func InitOrDie() {
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	var err error
	MsAuditQueue, err = msauditconnections.RabbitMQConnection.GetChannel().QueueDeclare(
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

	routingKey := messagebuscontracts.Route_Acl_Update
	err = msauditconnections.RabbitMQConnection.GetChannel().QueueBind(
		QueueName,                       // queue name
		routingKey,                      // routing key
		messagebuscontracts.ExchangeRor, // exchange
		false,
		nil,
	)
	if err != nil {
		args := [...]any{QueueName, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
