package v2sseclient

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/v2stream"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type SSEClient struct {
	mu            sync.RWMutex
	client        *httpclient.HttpTransportClient
	request       *http.Request
	callback      func(v2stream.RorEvent)
	isRetry       bool
	isClosing     bool
	retries       int
	retryLimit    int
	url           string
	lastEvetnID   string
	retyrInterval int
}

func NewSSEClient(client *httpclient.HttpTransportClient) *SSEClient {
	return &SSEClient{
		client: client,
	}
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
	time.Sleep(time.Second * time.Duration(s.retyrInterval))
	s.retries++
	s.callback(v2stream.NewRorEvent("info", fmt.Sprintf("Retrying, attempt %d of %d", s.retries, s.retryLimit)))
	return s.retries < s.retryLimit
}

func (s *SSEClient) SetLastEventID(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastEvetnID = id
}

func (s *SSEClient) SetRetryLimit(limit int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if limit < 0 {
		s.retryLimit = 1
		return
	}

	if limit > 30 {
		s.retryLimit = 30
		return
	}
	s.retryLimit = limit
}

func (s *SSEClient) UnSetRetry() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.retries = 0
	s.isRetry = false
}

func (s *SSEClient) Listen() (<-chan v2stream.RorEvent, error) {
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

	if strings.Split(resp.Header.Get("Content-Type"), ";")[0] != "text/event-stream" {
		return nil, fmt.Errorf("error: invalid Content-Type == %s", resp.Header.Get("Content-Type"))
	}

	events := make(chan v2stream.RorEvent)
	reader := bufio.NewReader(resp.Body)

	//events <- stream.NewRorEvent("debug", fmt.Sprintf("OpenURL resp: %+v", resp))

	go loop(s, reader, events)
	return events, nil
}

func (sse *SSEClient) OpenSSEStreamWithCallback(callback func(v2stream.RorEvent), path string) (<-chan struct{}, error) {
	var err error
	var retryinterval int = 5
	cancelCh := make(chan struct{})

	sseClient := &SSEClient{
		client:        sse.client,
		url:           sse.client.Config.BaseURL + path,
		retryLimit:    20,
		callback:      callback,
		retyrInterval: 1,
	}
	rorEvents := make(<-chan v2stream.RorEvent)
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
						callback(v2stream.NewRorEvent("error", "retying failed, closing channel"))
						close(cancelCh)
						return
					} else {
						callback(v2stream.NewRorEvent("error", "retying failed, retrying in 5 seconds"))
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

func (sse *SSEClient) OpenSSEStream(path string) (<-chan v2stream.RorEvent, error) {
	var err error

	sseClient := &SSEClient{
		client:        sse.client,
		url:           sse.client.Config.BaseURL + path,
		retyrInterval: 1,
	}
	events := make(chan v2stream.RorEvent)
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

func loop(client *SSEClient, reader *bufio.Reader, events chan v2stream.RorEvent) {
	eventId := ""
	eventdata := []byte{}
	eventtype := "unknown"

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			events <- v2stream.NewRorEvent("error", fmt.Sprintf("error during resp.Body read:%s", err))
			close(events)
			return
		}

		if hasPrefix(line, "\n") {
			// Empty line, dispatch message
			if len(eventdata) != 0 {
				events <- v2stream.RorEvent{Type: eventtype, Data: removeNewlineFromBytes(eventdata)}
				client.SetLastEventID(eventId)
			}
			eventId = ""
			eventdata = []byte{}
			eventtype = "unknown"

			continue
		}
		if hasPrefix(line, ":") {
			// Comment, do nothing
			continue
		}

		key, value := readLine(line)
		switch key {
		case "retry":
			// Retry, set retry interval
			retryInterval, err := strconv.Atoi(string(value))
			if err == nil {
				client.retyrInterval = retryInterval
			}
		case "id":
			eventId = string(value)
		case "data":
			eventdata = append(eventdata, []byte(value)...)
		case "event":
			eventtype = string(removeNewlineFromBytes(value))
		}
	}
}

// readLine reads a line from the reader and returns the key and value
// If the line does not contain a colon, the whole line is returned as key
func readLine(line []byte) (string, []byte) {
	input := string(line)
	if strings.Contains(input, ":") {
		splits := strings.Split(input, ":")
		if len(splits) > 2 {
			// Split on first colon, join the rest
			return splits[0], []byte(strings.Join(splits[1:], ":"))
		}
		return splits[0], []byte(splits[1])
	}
	// No colon, return input as key
	return input, []byte{}
}

func hasPrefix(s []byte, prefix string) bool {
	return bytes.HasPrefix(s, []byte(prefix))
}

func removeNewlineFromBytes(s []byte) []byte {
	return bytes.TrimSuffix(s, []byte("\n"))
}
