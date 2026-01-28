package mocktransport

import (
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/info"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportinfo"
)

type RorMockTransport struct {
	infoClientV1 v1info.InfoInterface
}

func NewRorMockTransport() *RorMockTransport {
	t := &RorMockTransport{
		infoClientV1: mocktransportinfo.NewV1Client(),
	}
	return t
}

func (t *RorMockTransport) Info() v1info.InfoInterface {
	return t.infoClientV1
}
