package tanzuagentrabbitmqdefinitions

import (
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/tanzuagentconnections"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

var (
	TanzuAgentQueueName = fmt.Sprintf("%s.%s", "tanzu-agent", viper.GetString(configconsts.TANZU_AGENT_DATACENTER))
	TanzuAgentQueue     amqp091.Queue
)

// Exchange -> Queue:
//
//	        -------------------
//	        | Exchange: Tanzu |-
//	        ------------------  \
//			                     \
//			                      \  --------------------------------------
//			                       > | Queue: tanzu-agent.<datacenternam> |
//			                         --------------------------------------
//
// Name: tanzu-agent.<datacenternam>
//
//   - Durable: true
//   - Arguments: x-queue-type: quorum
//
// Bindings:
//
//   - Exchange: tanzu
//   - Routing Key: ""
//   - Arguments: x-match: all, datacenter: <datacenternam>
//
// InitOrDie initializes the rabbitmq definitions and panics if it fails
func InitOrDie() {
	queueArgs := amqp091.Table{
		amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
	}
	var err error
	TanzuAgentQueue, err = tanzuagentconnections.RabbitMQConnection.GetChannel().QueueDeclare(
		TanzuAgentQueueName, // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		queueArgs,           // arguments
	)
	if err != nil {
		args := [...]any{TanzuAgentQueue, err}
		msg := fmt.Sprintf("could not declare exchange  %s,", args)
		rlog.Fatal(msg, err)
	}

	datacenter := viper.GetString(configconsts.TANZU_AGENT_DATACENTER)
	routingKey := ""
	err = tanzuagentconnections.RabbitMQConnection.GetChannel().QueueBind(
		TanzuAgentQueueName,               // queue name
		routingKey,                        // routing key
		messagebuscontracts.ExchangeTanzu, // exchange
		false,
		amqp091.Table{
			"x-match":    "all",
			"datacenter": datacenter,
		},
	)
	if err != nil {
		args := [...]any{TanzuAgentQueue, routingKey, err}
		msg := fmt.Sprintf("could not bind queue  %s,", args)
		rlog.Fatal(msg, err)
	}
}
