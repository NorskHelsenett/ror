package msnhnrabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/nhn/msnhnconnections"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/rabbitmq/amqp091-go"
)

var (
	MsNhnQueue amqp091.Queue
	QueueName  string = "ms-nhn"
)

//	       ---------------------------
//	       | Exchange: Ror.Resources | \
//		   ---------------------------  \
//		                                 \
//		                                  > ----------------------
//		                                    | Queue: ms-nhn       |
//		                                    -----------------------
//
//	       ---------------------------
//	       | Exchange: Ror            | \
//		   ---------------------------   \
//		                                  \
//		                                   > ----------------------
//		                                     | Queue: ms-nhn       |
//		                                     -----------------------
//
// Queue
// - Name: ms-nhn
// - Durable: false
// - Auto-Delete: false
// - Exclusive: false
// - No-Wait: false
// - Arguments: nil
//
// Bindings
// - Exchange: ror.resources
// - Routing Key: resource.#
// - Arguments: {"x-match":"all","apiVersion":"argoproj.io/v1alpha1","kind":"Application"}
//
// Bindings
// - Exchange: ror
// - Routing Key: cluster.created
// - Arguments: nil
//
// InitOrDie initializes the RabbitMQ definitions for the NHN microservice.
// If the definitions cannot be initialized, the application will exit.
func InitOrDie() {
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	var err error
	MsNhnQueue, err = msnhnconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		QueueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		queueArgs, // arguments
	)
	if err != nil {
		rlog.Fatal("could not declare queue", err, rlog.String("queue", QueueName))
	}

	routingKey := messagebuscontracts.Route_Cluster_Created
	err = msnhnconnections.RabbitMQConnection.GetChannel().QueueBind(
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

	err = msnhnconnections.RabbitMQConnection.GetChannel().QueueBind(
		QueueName,                                // queue name
		"resource.#",                             // routing key
		messagebuscontracts.ExchangeRorResources, // exchange
		false,
		amqp091.Table{
			"x-match":    "all",
			"apiVersion": "argoproj.io/v1alpha1",
			"kind":       "Application",
		},
	)
	if err != nil {
		args := [...]any{QueueName, "resources.#", err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
