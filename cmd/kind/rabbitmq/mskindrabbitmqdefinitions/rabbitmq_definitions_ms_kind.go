package mskindrabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/kind/mskindconnections"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/spf13/viper"

	"github.com/rabbitmq/amqp091-go"
)

var (
	MsKindQueueName = "ms-kind"
	MsKindQueue     amqp091.Queue
)

// InitOrDie Exchange
//
//		     -----------------
//			 | Exchange: Ror | -
//			 -----------------  \
//					             \
//				                  >  -------------------
//						             | Exchange: kind |
//						             -------------------
//	                            /
//	                           /
//	                          /
//	     ----------------------
//	    | Queue: ms-kind      |
//	     ----------------------
//
// kind Exchange: 	- type: headers
//   - durable: true
//   - autoDelete: false
//   - internal: false
//   - noWait: false
//   - arguments: nil
//
// Bindings:
//   - ROR -> kind
//   - key: "provider.kind.#"
//   - noWait: false
//   - arguments: nil
//
// Queue:
//   - Name: ms-kind
//   - Durable: true
//   - Arguments: x-queue-type: quorum
//   - Bindings:
//   - Exchange: kind
//   - Routing Key: "resource.*"
//   - Arguments: x-match: all, apiVersion: general.ror.internal/v1alpha1, kind: ClusterOrder
//   - Bindings:
//   - Exchange: kind
//   - Routing Key: "resource.*"
//
// InitOrDie initializes the RabbitMQ definitions
// and panics if it fails
// It is called from the main function
// and it is blocking
func InitOrDie() {
	if viper.GetBool(configconsts.DEVELOPMENT) {
		rlog.Info("---------------- MS-Kind RabbitMQ Definitions ---------------- ")
		rlog.Info("Initializing RabbitMQ definitions for ms-kind, will panic if it fails.")
		rlog.Info("ROR-Api creates the ROR rabbitmq exchange, ")
		rlog.Info("so ms-kind will fail/reboot if exchange is not created.... ")
		rlog.Info("---------------- MS-Kind RabbitMQ Definitions ---------------- ")
	}
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	err := mskindconnections.RabbitMQConnection.GetChannel().ExchangeDeclare(
		messagebuscontracts.ExchangeKind, // name
		"headers",                        // kind
		true,                             // durable
		false,                            // autoDelete -> delete when unused
		false,                            // internal
		false,                            // no-wait
		nil,                              // arguments
	)
	if err != nil {
		args := [...]any{messagebuscontracts.ExchangeRorResources, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	err = mskindconnections.RabbitMQConnection.GetChannel().ExchangeBind(
		messagebuscontracts.ExchangeKind, //destination
		"provider.kind.#",                // key
		messagebuscontracts.ExchangeRor,  // source
		false,                            // noWait
		nil,
	)
	if err != nil {
		panic(err)
	}

	MsKindQueue, err = mskindconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		MsKindQueueName, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		queueArgs,       // arguments
	)
	if err != nil {
		args := [...]any{MsKindQueueName, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	routingKey := messagebuscontracts.Route_ResourceCreated
	err = mskindconnections.RabbitMQConnection.GetChannel().QueueBind(
		MsKindQueueName,                          // queue name
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
		args := [...]any{MsKindQueueName, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
