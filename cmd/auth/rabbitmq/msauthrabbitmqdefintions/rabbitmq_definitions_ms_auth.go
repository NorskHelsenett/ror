package msauthrabbitmqdefintions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/auth/msauthconnections"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/rabbitmq/amqp091-go"
)

var (
	QueueName   = "ms-auth"
	MsAuthQueue amqp091.Queue
)

func InitOrDie() {
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	var err error
	MsAuthQueue, err = msauthconnections.RabbitMQConnection.GetChannel().QueueDeclare(
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

	routingKey := messagebuscontracts.Route_Auth
	err = msauthconnections.RabbitMQConnection.GetChannel().QueueBind(
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
