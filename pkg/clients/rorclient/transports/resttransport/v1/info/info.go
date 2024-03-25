package info

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/info",
	}
}

func (c *V1Client) GetVersion() (string, error) {
	var versiondata struct {
		Version string `json:"version"`
	}

	err := c.Client.GetJSON(c.basePath+"/version", &versiondata, httpclient.HttpTransportClientParamsNoAuth)
	if err != nil {
		return "", err
	}

	return versiondata.Version, nil
}
