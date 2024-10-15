package mstalosrabbitmqdefinitions

import (
	"fmt"
	mstalosconnections "github.com/NorskHelsenett/ror/cmd/talos/mstalosconnections"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/spf13/viper"

	"github.com/rabbitmq/amqp091-go"
)

var (
	MsTalosQueueName = "ms-talos"
	MsTalosQueue     amqp091.Queue
)

// InitOrDie Exchange
//
//		     -----------------
//			 | Exchange: Ror | -
//			 -----------------  \
//					             \
//				                  >  -------------------
//						             | Exchange: talos |
//						             -------------------
//	                            /
//	                           /
//	                          /
//	     ----------------------
//	    | Queue: ms-talos      |
//	     ----------------------
//
// Talos Exchange: 	- type: headers
//   - durable: true
//   - autoDelete: false
//   - internal: false
//   - noWait: false
//   - arguments: nil
//
// Bindings:
//   - ROR -> talos
//   - key: "provider.talos.#"
//   - noWait: false
//   - arguments: nil
//
// Queue:
//   - Name: ms-talos
//   - Durable: true
//   - Arguments: x-queue-type: quorum
//   - Bindings:
//   - Exchange: talos
//   - Routing Key: "resource.*"
//   - Arguments: x-match: all, apiVersion: general.ror.internal/v1alpha1, kind: ClusterOrder
//   - Bindings:
//   - Exchange: talos
//   - Routing Key: "resource.*"
//
// InitOrDie initializes the RabbitMQ definitions
// and panics if it fails
// It is called from the main function
// and it is blocking
func InitOrDie() {
	if viper.GetBool(configconsts.DEVELOPMENT) {
		rlog.Info("---------------- MS-Talos RabbitMQ Definitions ---------------- ")
		rlog.Info("Initializing RabbitMQ definitions for ms-talos, will panic if it fails.")
		rlog.Info("ROR-Api creates the ROR rabbitmq exchange, ")
		rlog.Info("so ms-talos will fail/reboot if exchange is not created.... ")
		rlog.Info("---------------- MS-Talos RabbitMQ Definitions ---------------- ")
	}
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	err := mstalosconnections.RabbitMQConnection.GetChannel().ExchangeDeclare(
		messagebuscontracts.ExchangeTalos, // name
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

	err = mstalosconnections.RabbitMQConnection.GetChannel().ExchangeBind(
		messagebuscontracts.ExchangeTalos, //destination
		"provider.talos.#",                // key
		messagebuscontracts.ExchangeRor,   // source
		false,                             // noWait
		nil,
	)
	if err != nil {
		panic(err)
	}

	MsTalosQueue, err = mstalosconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		MsTalosQueueName, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		queueArgs,        // arguments
	)
	if err != nil {
		args := [...]any{MsTalosQueueName, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	routingKey := messagebuscontracts.Route_ResourceCreated
	err = mstalosconnections.RabbitMQConnection.GetChannel().QueueBind(
		MsTalosQueueName,                         // queue name
		"resource.created",                       // routing key
		messagebuscontracts.ExchangeRorResources, // exchange
		false,
		amqp091.Table{
			"x-match":    "all",
			"apiVersion": "general.ror.internal/v1alpha1",
			"kind":       "ClusterOrder",
		},
	)
	if err != nil {
		args := [...]any{MsTalosQueueName, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
