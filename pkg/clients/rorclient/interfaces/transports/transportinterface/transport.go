package transportinterface

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/clientinterface"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportstatusinterface"
)

type RorTransport interface {
	clientinterface.RorCommonClientApiInterface

	RorCommonTransportInterface

	RorCommonClientTransportInterface
}

type RorCommonClientTransportInterface interface {
	CheckConnection() error
	GetRole() string
	GetApiSecret() string
}

type RorCommonTransportInterface interface {
	Ping(ctx context.Context) bool

	GetTransportName() string
	Status() transportstatusinterface.RorTransportStatus
}

type RorCommonClientTransportSetterInterface interface {
	SetTransport(transport RorTransport)
}
