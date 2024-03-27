package mocktransport

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportinfo"
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/info"
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
