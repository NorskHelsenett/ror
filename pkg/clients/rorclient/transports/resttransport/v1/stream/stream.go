package stream

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/stream"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
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
