package transportinterface

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/clientinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportstatusinterface"
)

type RorTransport interface {
	clientinterface.RorCommonClientInterface

	GetTransportName() string
	Status() transportstatusinterface.RorTransportStatus

	Ping(ctx context.Context) bool

	RorCommonTransport
}

type RorCommonTransport interface {
	CheckConnection() error
	GetRole() string
	GetApiSecret() string
}
