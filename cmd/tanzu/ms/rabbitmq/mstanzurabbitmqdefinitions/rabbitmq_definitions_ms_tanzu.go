package mstanzurabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/ms/mstanzuconnections"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/rabbitmq/amqp091-go"
)

var (
	MsTanzuQueueName = "ms-tanzu"
	MsTanzuqueue     amqp091.Queue
)

// InitOrDie Exchange
//
//		     -----------------
//			 | Exchange: Ror | -
//			 -----------------  \
//					             \
//				                  >  -------------------
//						             | Exchange: Tanzu |
//						             -------------------
//	                            /
//	                           /
//	                          /
//	     ----------------------
//	    | Queue: ms-tanzu      |
//	     ----------------------
//
// Tanzu Exchange: 	- type: headers
//   - durable: true
//   - autoDelete: false
//   - internal: false
//   - noWait: false
//   - arguments: nil
//
// Bindings:
//   - ROR -> Tanzu
//   - key: "provider.tanzu.#"
//   - noWait: false
//   - arguments: nil
//
// Queue:
//   - Name: ms-tanzu
//   - Durable: true
//   - Arguments: x-queue-type: quorum
//   - Bindings:
//   - Exchange: Tanzu
//   - Routing Key: "resource.*"
//   - Arguments: x-match: all, apiVersion: general.ror.internal/v1alpha1, kind: ClusterOrder
//   - Bindings:
//   - Exchange: Tanzu
//   - Routing Key: "resource.*"
//   - Arguments: x-match: all, apiVersion: run.tanzu.vmware.com/v1alpha2, kind: TanzuKubernetesCluster
//
// InitOrDie initializes the RabbitMQ definitions
// and panics if it fails
// It is called from the main function
// and it is blocking
func InitOrDie() {
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	err := mstanzuconnections.RabbitMQConnection.GetChannel().ExchangeDeclare(
		messagebuscontracts.ExchangeTanzu, // name
		"headers",                         // kind
		true,                              // durable
		false,                             // autoDelete -> delete when unused
		false,                             // internal
		false,                             // no-wait
		nil,                               // arguments
	)
	if err != nil {
		args := [...]any{messagebuscontracts.ExchangeRorResources, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	err = mstanzuconnections.RabbitMQConnection.GetChannel().ExchangeBind(
		messagebuscontracts.ExchangeTanzu, //destination
		"provider.tanzu.#",                // key
		messagebuscontracts.ExchangeRor,   // source
		false,                             // noWait
		nil,
	)
	if err != nil {
		panic(err)
	}

	MsTanzuqueue, err = mstanzuconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		MsTanzuQueueName, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		queueArgs,        // arguments
	)
	if err != nil {
		args := [...]any{MsTanzuQueueName, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	routingKey := messagebuscontracts.Route_ResourceCreated
	err = mstanzuconnections.RabbitMQConnection.GetChannel().QueueBind(
		MsTanzuQueueName,                         // queue name
		"resource.#",                             // routing key
		messagebuscontracts.ExchangeRorResources, // exchange
		false,
		amqp091.Table{
			"x-match":    "all",
			"apiVersion": "general.ror.internal/v1alpha1",
			"kind":       "ClusterOrder",
		},
	)
	if err != nil {
		args := [...]any{MsTanzuQueueName, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}

	err = mstanzuconnections.RabbitMQConnection.GetChannel().QueueBind(
		MsTanzuQueueName,                         // queue name
		"resource.#",                             // routing key
		messagebuscontracts.ExchangeRorResources, // exchange
		false,
		amqp091.Table{
			"x-match":    "all",
			"apiVersion": "run.tanzu.vmware.com/v1alpha2",
			"kind":       "TanzuKubernetesCluster",
		},
	)
	if err != nil {
		args := [...]any{MsTanzuQueueName, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
