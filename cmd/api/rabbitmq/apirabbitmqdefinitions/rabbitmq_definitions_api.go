package apirabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"github.com/rabbitmq/amqp091-go"
)

var (
	ApiEventsQueueName string = "sse-events"
	ApiEventsqueue     amqp091.Queue
)

func init() {
	ApiEventsQueueName = fmt.Sprintf("%s-%s", ApiEventsQueueName, uuid.New().String())

}

//		Exchanges:
//		 -----------------
//		 | Ror           | -
//		 -----------------  \
//		           |         \
//		           |          \   -----------------
//		           |           >  |  ROR Resources |
//		           |              -----------------
//		           |
//				-----------------
//				| ROR Events    |
//				-----------------
//	            |
//	            |
//	         ----------------------
//	         | Queue: sse-events  |
//	         ----------------------
//
// Ror Exchange: 	- type: topic
//   - durable: true
//   - autoDelete: false
//   - internal: false
//   - noWait: false
//   - arguments: nil
//
// Ror.Resources Exchange:
//   - type: headers
//   - durable: true
//   - autoDelete: false
//   - internal: false
//   - noWait: false
//   - arguments: nil
//
// Bindings:
//   - ROR -> ROR Resources
//   - key: resources.#
//   - noWait: false
//   - arguments: nil
//
// Ror.Events Exchange:
//   - type: headers
//   - durable: true
//   - autoDelete: false
//   - internal: false
//   - noWait: false
//   - arguments: nil
//
// Bindings:
//   - ROR -> ROR Events
//   - key: event.#
//   - noWait: false
//   - arguments: nil
//
// Queue:
// Bindings:
//   - ROR -> Tanzu
//   - key: "provider.tanzu.#"
//   - noWait: false
//   - arguments: nil
//
// Queue:
//   - Name: api-events
//   - Durable: true
//   - Arguments: x-queue-type: quorum
//   - Bindings:
//   - Exchange: ror.events
//   - Routing Key: "resource.*"
//
// InitOrDie initializes the RabbitMQ definitions
// and panics if it fails
// It is called from the main function
// and it is blocking
func InitOrDie() {
	queueArgs := amqp091.Table{}

	err := apiconnections.RabbitMQConnection.GetChannel().ExchangeDeclare(
		messagebuscontracts.ExchangeRor, // name
		"topic",                         // kind
		true,                            // durable
		false,                           // autoDelete -> delete when unused
		false,                           // internal
		false,                           // no-wait
		nil,                             // arguments
	)
	if err != nil {
		args := [...]any{messagebuscontracts.ExchangeRor, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	err = apiconnections.RabbitMQConnection.GetChannel().ExchangeDeclare(
		messagebuscontracts.ExchangeRorResources, // name
		"headers",                                // kind
		true,                                     // durable
		false,                                    // autoDelete -> delete when unused
		false,                                    // internal
		false,                                    // no-wait
		nil,                                      // arguments
	)
	if err != nil {
		args := [...]any{messagebuscontracts.ExchangeRorResources, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	err = apiconnections.RabbitMQConnection.GetChannel().ExchangeBind(
		messagebuscontracts.ExchangeRorResources, //destination
		"resource.#",                             // key
		messagebuscontracts.ExchangeRor,          // source
		false,                                    // noWait
		nil,                                      // arguments
	)
	if err != nil {
		panic(err)
	}

	err = apiconnections.RabbitMQConnection.GetChannel().ExchangeDeclare(
		messagebuscontracts.ExchangeRorEvents, // name
		"fanout",                              // kind
		true,                                  // durable
		false,                                 // autoDelete -> delete when unused
		false,                                 // internal
		false,                                 // no-wait
		nil,                                   // arguments
	)
	if err != nil {
		args := [...]any{messagebuscontracts.ExchangeRorEvents, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	err = apiconnections.RabbitMQConnection.GetChannel().ExchangeBind(
		messagebuscontracts.ExchangeRorEvents, //destination
		"event.#",                             // key
		messagebuscontracts.ExchangeRor,       // source
		false,                                 // noWait
		nil,                                   // arguments
	)
	if err != nil {
		panic(err)
	}

	automaticallyDelete := !viper.GetBool(configconsts.DEVELOPMENT)
	ApiEventsqueue, err = apiconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		ApiEventsQueueName,  // name
		true,                // durable
		automaticallyDelete, // delete when unused
		false,               // exclusive
		false,               // no-wait
		queueArgs,           // arguments, non quorum queue
	)
	if err != nil {
		args := [...]any{ApiEventsQueueName, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	err = apiconnections.RabbitMQConnection.GetChannel().QueueBind(
		ApiEventsQueueName,                    // queue name
		"",                                    // routing key
		messagebuscontracts.ExchangeRorEvents, // exchange
		false,
		nil,
	)
	if err != nil {
		args := [...]any{ApiEventsQueueName, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
