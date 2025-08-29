package rabbitmqclient

import (
	"fmt"
	"strings"
)

type RabbitMQConnectionOption interface {
	apply(*rabbitmqcon)
}

type optionFunc func(*rabbitmqcon)

func (of optionFunc) apply(cfg *rabbitmqcon) { of(cfg) }

func ServerString(serverstring string) RabbitMQConnectionOption {
	return optionFunc(func(cfg *rabbitmqcon) {
		var err error
		serverparts := strings.SplitN(serverstring, ":", 2)
		if len(serverparts) == 2 {
			cfg.Host = serverparts[0]
			cfg.Port = serverparts[1]
		} else {
			cfg.Host = serverparts[0]
			cfg.Port = "5672" // default RabbitMQ port
		}
		if cfg.Host == "" || cfg.Port == "" {
			fmt.Println("Error parsing server string: ", err)

		}
	})
}

func Host(host string) RabbitMQConnectionOption {
	return optionFunc(func(cfg *rabbitmqcon) {
		cfg.Host = host
	})
}

func Port(port string) RabbitMQConnectionOption {
	return optionFunc(func(cfg *rabbitmqcon) {
		cfg.Port = port
	})
}
func BroadcastName(broadcastname string) RabbitMQConnectionOption {
	return optionFunc(func(cfg *rabbitmqcon) {
		cfg.BroadcastName = broadcastname
	})
}

func CredentialsProvider(cp RabbitMQCredentialProvider) RabbitMQConnectionOption {
	return optionFunc(func(cfg *rabbitmqcon) {
		cfg.Credentials = cp
	})
}
