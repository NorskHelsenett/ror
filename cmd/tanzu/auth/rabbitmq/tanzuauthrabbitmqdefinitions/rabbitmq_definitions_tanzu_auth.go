package tanzuauthrabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/auth/tanzuauthconnections"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/rabbitmq/amqp091-go"
)

var (
	TanzuAuthQueueName = "tanzu-auth"
	TanzuAuthQueue     amqp091.Queue
)

// Exchange -> Queue:
//
//	        -------------------
//	        | Exchange: ROR |-
//	        ------------------  \
//			                     \
//			                      \  --------------------------------------
//			                       > | Queue: tanzu-auth |
//			                         --------------------------------------
//
// Name: tanzu-auth
//
//   - Durable: true
//   - Arguments: x-queue-type: quorum
//
// Bindings:
//
//   - Exchange: ror
//   - Routing Key: ""
//
// InitOrDie initializes the rabbitmq definitions and panics if it fails
func InitOrDie() {
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	var err error
	TanzuAuthQueue, err = tanzuauthconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		TanzuAuthQueueName, // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		queueArgs,          // arguments
	)
	if err != nil {
		args := [...]any{TanzuAuthQueue, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	routingKey := ""
	err = tanzuauthconnections.RabbitMQConnection.GetChannel().QueueBind(
		TanzuAuthQueueName,              // queue name
		routingKey,                      // routing key
		messagebuscontracts.ExchangeRor, // exchange
		false,
		nil,
	)
	if err != nil {
		args := [...]any{TanzuAuthQueue, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
