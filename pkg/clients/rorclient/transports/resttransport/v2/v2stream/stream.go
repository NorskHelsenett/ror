package v2stream

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/v2stream"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/sseclient/v2sseclient"
)

type V2Client struct {
	Client   *v2sseclient.SSEClient
	basePath string
}

func NewV2Client(client *v2sseclient.SSEClient) *V2Client {
	return &V2Client{
		Client:   client,
		basePath: "/v2/events",
	}
}

func (c *V2Client) StartEventstream() (<-chan v2stream.RorEvent, error) {
	stream, err := c.Client.OpenSSEStream(c.basePath + "/listen")
	if err != nil {
		return nil, err
	}

	return stream, nil
}

func (c *V2Client) StartEventstreamWithCallback(callback func(v2stream.RorEvent)) (<-chan struct{}, error) {
	return c.Client.OpenSSEStreamWithCallback(callback, c.basePath+"/listen")
}

func (c *V2Client) BroadcastEvent(event v2stream.RorEvent) error {
	return c.Client.BroadcastEvent(c.basePath+"/send", event)
}
