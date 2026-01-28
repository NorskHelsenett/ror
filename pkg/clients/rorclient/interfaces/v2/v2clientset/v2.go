package v2clientset

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/clientinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/apikeys"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/rorclientv2self"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/token"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/v2stream"
)

type ClientSet struct {
	transport transportinterface.RorTransport

	apikeysClientV2   apikeys.ApiKeysInterface
	selfClientV2      rorclientv2self.SelfInterface
	resourcesClientV2 resources.ResourcesInterface
	streamClientV2    v2stream.StreamInterface
	tokenClientV2     token.TokenInterface
}

func NewV2ClientSet(transport transportinterface.RorTransport) clientinterface.RorCommonClientApiInterfaceV2 {
	return &ClientSet{
		apikeysClientV2:   transport.ApiKeysV2(),
		selfClientV2:      transport.Self(),
		resourcesClientV2: transport.ResourcesV2(),
		streamClientV2:    transport.StreamV2(),
		tokenClientV2:     transport.TokenV2(),
	}
}
func (c *ClientSet) ApiKeys() apikeys.ApiKeysInterface {
	return c.apikeysClientV2
}

func (c *ClientSet) Resources() resources.ResourcesInterface {
	return c.resourcesClientV2
}
func (c *ClientSet) Self() rorclientv2self.SelfInterface {
	return c.selfClientV2
}

func (c *ClientSet) Stream() v2stream.StreamInterface {
	return c.streamClientV2
}

func (c *ClientSet) Token() token.TokenInterface {
	return c.tokenClientV2

}
