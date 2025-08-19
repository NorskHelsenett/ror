package mocktransportstreamv2

import (
	v2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/v2stream"
)

type V2Client struct{}

func NewV2Client() *V2Client {
	return &V2Client{}
}

func (c *V2Client) StartEventstream() (<-chan v2stream.RorEvent, error) {
	// Create a channel for mock events
	eventChan := make(chan v2stream.RorEvent, 10)

	// Send a mock event
	go func() {
		defer close(eventChan)
		mockEvent := v2stream.RorEvent{
			Type: "mock",
			Data: []byte(`{"message": "mock v2 event"}`),
		}
		eventChan <- mockEvent
	}()

	return eventChan, nil
}

func (c *V2Client) StartEventstreamWithCallback(callbackfunc func(v2stream.RorEvent)) (<-chan struct{}, error) {
	// Create a done channel
	done := make(chan struct{})

	// Start a goroutine that sends mock events to the callback
	go func() {
		defer close(done)
		mockEvent := v2stream.RorEvent{
			Type: "mock",
			Data: []byte(`{"message": "mock v2 event with callback"}`),
		}
		callbackfunc(mockEvent)
	}()

	return done, nil
}

func (c *V2Client) BroadcastEvent(event v2stream.RorEvent) error {
	// Mock implementation - just return nil to simulate success
	return nil
}
