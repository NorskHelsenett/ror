package rabbitmqclient

import "github.com/NorskHelsenett/ror/pkg/rlog"

func failOnError(err error, msg string) {
	if err != nil {
		rlog.Fatal(msg, err)
	}
}
