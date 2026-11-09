package httpclient_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
)

func TestHTTPErrorContract(t *testing.T) {
	for _, method := range []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodHead} {
		for status := 100; status <= 599; status++ {
			t.Run(fmt.Sprintf("%s/%d", method, status), func(t *testing.T) {
				req, err := http.NewRequest(method, "https://example.test/resource", nil)
				if err != nil {
					t.Fatal(err)
				}
				client := &httpclient.HttpTransportClient{Status: httpclient.NewHttpTransportClientStatus()}
				res := &http.Response{
					StatusCode: status,
					Status:     fmt.Sprintf("%d %s", status, http.StatusText(status)),
					Request:    req,
					Header:     http.Header{"Retry-After": []string{"120"}},
				}
				before := time.Now()
				err = client.HandleNonOk(res)
				if status >= 200 && status <= 399 {
					if err != nil {
						t.Fatalf("accepted response returned error: %v", err)
					}
					return
				}
				var responseErr *httpclient.HTTPError
				if !errors.As(fmt.Errorf("wrapped: %w", err), &responseErr) {
					t.Fatalf("expected wrapped HTTPError, got %T: %v", err, err)
				}
				if responseErr.StatusCode != status || responseErr.Method != method || responseErr.URL != req.URL.String() {
					t.Fatalf("incorrect response metadata: %+v", responseErr)
				}
				message := fmt.Sprintf("http error: %s from %s", res.Status, req.URL)
				var delay time.Duration
				switch {
				case status == 429 || status == 503:
					delay = 120 * time.Second
					message += fmt.Sprintf(" (retry after: %s)", responseErr.RetryAfter.Format(time.RFC3339))
				case status == 401 && method != http.MethodHead:
					delay = httpclient.DefaultTimeout
					message += " (unauthorized, check your authentication) Connection throttled"
				case status == 403 && method != http.MethodHead:
					delay = httpclient.DefaultTimeout
					message += " (forbidden, check your permissions) Connection throttled"
				}
				if err.Error() != message {
					t.Fatalf("message changed: got %q, want %q", err.Error(), message)
				}
				if responseErr.RetryAfter != client.Status.RetryAfter {
					t.Fatal("error and client retry timestamps differ")
				}
				if delay == 0 {
					if !responseErr.RetryAfter.IsZero() {
						t.Fatal("unexpected retry delay")
					}
				} else if responseErr.RetryAfter.Before(before.Add(delay)) || responseErr.RetryAfter.After(time.Now().Add(delay)) {
					t.Fatalf("incorrect retry timestamp: %s", responseErr.RetryAfter)
				}
			})
		}
	}
}

func TestRequestMethodsReturnHTTPError(t *testing.T) {
	methods := map[string]func(*httpclient.HttpTransportClient) error{
		"GET": func(c *httpclient.HttpTransportClient) error {
			return c.GetJSON(context.Background(), "/resource", nil)
		},
		"POST": func(c *httpclient.HttpTransportClient) error {
			return c.PostJSON(context.Background(), "/resource", nil, nil)
		},
		"PUT": func(c *httpclient.HttpTransportClient) error {
			return c.PutJSON(context.Background(), "/resource", nil, nil)
		},
		"DELETE": func(c *httpclient.HttpTransportClient) error {
			return c.Delete(context.Background(), "/resource", nil)
		},
		"HEAD": func(c *httpclient.HttpTransportClient) error {
			_, _, err := c.Head(context.Background(), "/resource")
			return err
		},
		"GetRaw": func(c *httpclient.HttpTransportClient) error {
			_, err := c.GetRaw(context.Background(), "/resource")
			return err
		},
	}
	for name, request := range methods {
		for _, status := range []int{400, 401, 403, 404, 409, 422, 429, 500, 503} {
			t.Run(fmt.Sprintf("%s/%d", name, status), func(t *testing.T) {
				client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(status)
				})
				err := request(client)
				var responseErr *httpclient.HTTPError
				if !errors.As(err, &responseErr) || responseErr.StatusCode != status {
					t.Fatalf("expected HTTPError status %d, got %T: %v", status, err, err)
				}
			})
		}
	}
}

func TestNonResponseErrorsRetainIdentity(t *testing.T) {
	client := newRawClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("invalid json"))
	})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := client.GetJSON(ctx, "/resource", nil)
	var responseErr *httpclient.HTTPError
	if !errors.Is(err, context.Canceled) || errors.As(err, &responseErr) {
		t.Fatalf("expected original cancellation error, got %T: %v", err, err)
	}
	var result any
	err = client.GetJSON(context.Background(), "/resource", &result)
	var syntaxErr *json.SyntaxError
	if !errors.As(err, &syntaxErr) || errors.As(err, &responseErr) {
		t.Fatalf("expected JSON decoding error, got %T: %v", err, err)
	}
}

// testNewHttpTransportClientConfig tests that the constructor functions and
// it passes valid addresses and rejects any invalid addresses.
func TestNewHttpTransportClientConfig(t *testing.T) {

	authProvider := httpauthprovider.NewNoAuthprovider()
	role := "test"
	version := rorversion.GetRorVersion()
	urls := map[string]bool{
		"":                        false,
		"127.0.0.1":               false,
		"127.0.0.1:10000":         false,
		"http://localhost":        true,
		"http://localhost:10000":  true,
		"https://localhost":       true,
		"https://localhost:10000": true,
		"asfjniefsjfsdjdsf":       false,
		"::f1":                    false,
	}

	for url, expectedAction := range urls {
		_, err := httpclient.NewHttpTransportClientConfig(url, authProvider, role, version)
		hadError := err == nil
		if hadError != expectedAction {
			t.Errorf("failed on validation on url '%v', got %v, expected %v", url, hadError, expectedAction)
		} else {
			t.Logf("passed url validation on '%v', got %v, expected %v", url, hadError, expectedAction)
		}

	}
}
