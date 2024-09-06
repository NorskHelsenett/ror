package httpclient

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/stream"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type SSEClient struct {
	mu         sync.RWMutex
	client     *HttpTransportClient
	request    *http.Request
	callback   func(stream.RorEvent)
	isRetry    bool
	isClosing  bool
	retries    int
	retryLimit int
	url        string
}

func (s *SSEClient) createRequest() error {
	var err error
	s.request, err = http.NewRequest("GET", s.url, nil)
	if err != nil {
		return err
	}

	s.client.Config.AuthProvider.AddAuthHeaders(s.request)
	s.request.Header.Set("User-Agent", fmt.Sprintf("%s - v%s (%s)", s.client.Config.Role, s.client.Config.Version.GetVersion(), s.client.Config.Version.GetCommit()))
	s.request.Header.Set("Cache-Control", "no-cache")
	s.request.Header.Set("Accept", "text/event-stream")
	s.request.Header.Set("Connection", "keep-alive")

	return nil

}

func (s *SSEClient) CheckRetry() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.isRetry {
		s.isRetry = true
		return true
	}
	s.retries++
	s.callback(stream.NewRorEvent("info", fmt.Sprintf("Retrying, attempt %d of %d", s.retries, s.retryLimit)))
	return s.retries < s.retryLimit
}

func (s *SSEClient) UnSetRetry() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.retries = 0
	s.isRetry = false
}

func (s *SSEClient) Listen() (<-chan stream.RorEvent, error) {
	err := s.createRequest()
	if err != nil {
		return nil, err
	}

	client := s.client.Client
	resp, err := client.Do(s.request)
	if err != nil {
		return nil, err
	}

	rlog.Debug(fmt.Sprintf("OpenURL resp: %+v", resp))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: resp.StatusCode == %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "text/event-stream" {
		return nil, fmt.Errorf("error: invalid Content-Type == %s", resp.Header.Get("Content-Type"))
	}

	events := make(chan stream.RorEvent)
	reader := bufio.NewReader(resp.Body)

	//events <- stream.NewRorEvent("debug", fmt.Sprintf("OpenURL resp: %+v", resp))

	go loop(reader, events)
	return events, nil
}

func (t *HttpTransportClient) OpenSSEStreamWithCallback(callback func(stream.RorEvent), path string) (<-chan struct{}, error) {
	var err error
	var retryinterval int = 5
	cancelCh := make(chan struct{})

	sseClient := &SSEClient{
		client:     t,
		url:        t.Config.BaseURL + path,
		retryLimit: 20,
		callback:   callback,
	}
	rorEvents := make(<-chan stream.RorEvent)
	rorEvents, err = sseClient.Listen()
	go func() {
		for {
			if sseClient.isClosing {
				close(cancelCh)
				return
			}
			if sseClient.isRetry {
				rorEvents, err = sseClient.Listen()

				if err != nil {
					if !sseClient.CheckRetry() {
						callback(stream.NewRorEvent("error", "retying failed, closing channel"))
						close(cancelCh)
						return
					} else {
						time.Sleep(time.Second * time.Duration(retryinterval))
						continue
					}
				}
				if err == nil {
					sseClient.UnSetRetry()
				}
				continue
			}
			event, open := <-rorEvents
			if !open && sseClient.CheckRetry() {
				continue
			}
			callback(event)
		}
	}()

	return cancelCh, nil
}

func (t *HttpTransportClient) OpenSSEStream(path string) (<-chan stream.RorEvent, error) {
	var err error

	sseClient := &SSEClient{
		client: t,
		url:    t.Config.BaseURL + path,
	}
	events := make(chan stream.RorEvent)
	rorEvents, err := sseClient.Listen()
	if err != nil {
		return nil, err
	}
	go func() {
		for rorEvent := range rorEvents {
			events <- rorEvent
		}
	}()

	return events, nil
}

func loop(reader *bufio.Reader, events chan stream.RorEvent) {

	for {

		line, err := reader.ReadBytes('\n')
		if err != nil {
			//fmt.Fprintf(os.Stderr, "error during resp.Body read:%s\n", err)
			events <- stream.NewRorEvent("error", fmt.Sprintf("error during resp.Body read:%s", err))
			close(events)
			return
		}
		eventdata := []byte{}
		eventtype := "unknown"

		switch {
		case hasPrefix(line, "\n"):
			// Empty line, do nothing
		case hasPrefix(line, ":"):
			// Comment, do nothing
		case hasPrefix(line, "retry:"):
			// Retry, do nothing for now
		case hasPrefix(line, "data: "):
			eventdata = line[6:]
		case hasPrefix(line, "data:"):
			eventdata = line[5:]

		// name of event
		case hasPrefix(line, "event: "):
			eventdata = line[7 : len(line)-1]
		case hasPrefix(line, "event:"):
			eventdata = line[6 : len(line)-1]
		default:
			events <- stream.NewRorEvent("error", fmt.Sprintf("Error: len:%d\n%s", len(line), line))
			close(events)
		}
		if len(eventdata) == 0 {
			continue
		}
		if hasPostfix(eventdata, "\n") {
			eventdata = trimPostfix(eventdata, "\n")
		}
		if contains(eventdata, "event") {
			event := apicontracts.SSEMessage{}
			err := json.Unmarshal(eventdata, &event)
			if err != nil {
				continue
			}
			eventtype = event.Event
		}

		events <- stream.RorEvent{Type: string(eventtype), Data: eventdata}
	}

}

func hasPrefix(s []byte, prefix string) bool {
	return bytes.HasPrefix(s, []byte(prefix))
}
func hasPostfix(s []byte, postfix string) bool {
	return bytes.HasSuffix(s, []byte(postfix))
}

func trimPostfix(s []byte, postfix string) []byte {
	return bytes.TrimSuffix(s, []byte(postfix))
}

func contains(s []byte, contaninsstring string) bool {
	return bytes.Contains(s, []byte(contaninsstring))
}
