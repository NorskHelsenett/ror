package token

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/token",
	}
}

func (c *V1Client) GetJWKS(_ context.Context) (string, error) {

	return "not implemented", nil
}
