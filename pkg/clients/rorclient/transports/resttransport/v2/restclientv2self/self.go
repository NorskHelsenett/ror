package restclientv2self

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"
)

type V2Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV2Client(client *httpclient.HttpTransportClient) *V2Client {
	return &V2Client{
		Client:   client,
		basePath: "/v2/self",
	}
}

func (c *V2Client) Get() (apicontractsv2self.SelfData, error) {
	var selfdata apicontractsv2self.SelfData
	err := c.Client.GetJSON(c.basePath, &selfdata)
	if err != nil {
		return apicontractsv2self.SelfData{}, err
	}

	return selfdata, nil
}
