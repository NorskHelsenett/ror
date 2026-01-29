package v1stream

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/stream"
	v1sseclient "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/sseclient/v1sseclient"
)

type V1Client struct {
	Client   *v1sseclient.SSEClient
	basePath string
}

func NewV1Client(client *v1sseclient.SSEClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/events",
	}
}

func (c *V1Client) StartEventstream() (<-chan stream.RorEvent, error) {
	stream, err := c.Client.OpenSSEStream(c.basePath + "/listen")
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *V1Client) StartEventstreamWithCallback(callback func(stream.RorEvent)) (<-chan struct{}, error) {
	cancelCh, err := c.Client.OpenSSEStreamWithCallback(callback, c.basePath+"/listen")
	if err != nil {
		return nil, err
	}
	return cancelCh, nil
}
