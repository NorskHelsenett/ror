package stream

import "encoding/json"

type RorEvent struct {
	Type string
	Data []byte
}

type EventData struct {
	Id    string `json:"id"`
	Event string `json:"event"`
	Data  string `json:"data"`
}

type StreamInterface interface {
	StartEventstream() (<-chan RorEvent, error)
	StartEventstreamWithCallback(callbackfunc func(RorEvent)) (<-chan struct{}, error)
}

func NewRorEvent(eventType string, data string) RorEvent {
	event := EventData{
		Event: eventType,
		Data:  data,
	}
	jsonevent, err := json.Marshal(event)
	if err != nil {
		jsonevent = []byte("")
	}

	return RorEvent{
		Type: eventType,
		Data: jsonevent,
	}
}
