package mocktransportstream

import (
	v1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/stream"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) StartEventstream() (<-chan v1stream.RorEvent, error) {
	// Create a channel for mock events
	eventChan := make(chan v1stream.RorEvent, 10)

	// Send a mock event
	go func() {
		defer close(eventChan)
		mockEvent := v1stream.RorEvent{
			Type: "mock",
			Data: []byte(`{"message": "mock event"}`),
		}
		eventChan <- mockEvent
	}()

	return eventChan, nil
}

func (c *V1Client) StartEventstreamWithCallback(callbackfunc func(v1stream.RorEvent)) (<-chan struct{}, error) {
	// Create a done channel
	done := make(chan struct{})

	// Start a goroutine that sends mock events to the callback
	go func() {
		defer close(done)
		mockEvent := v1stream.RorEvent{
			Type: "mock",
			Data: []byte(`{"message": "mock event with callback"}`),
		}
		callbackfunc(mockEvent)
	}()

	return done, nil
}
