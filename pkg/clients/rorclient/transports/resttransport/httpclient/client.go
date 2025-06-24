package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

var (
	DefaultTimeout = 60 * time.Second
)

type HttpTransportClientParams struct {
	Key   HttpTransportClientOpts
	Value any
}

type HttpTransportClientOpts string

type HttpTransportClientStatus struct {
	Established bool      `json:"established"`
	ApiVersion  string    `json:"api_version"`
	LibVersion  string    `json:"lib_version"`
	RetryAfter  time.Time `json:"retry_after"`
}

const (
	HttpTransportClientOptsNoAuth  HttpTransportClientOpts = "NOAUTH"
	HttpTransportClientOptsHeaders HttpTransportClientOpts = "HEADERS"
	HttpTransportClientOptsQuery   HttpTransportClientOpts = "QUERY"
	HttpTransportClientTimeout     HttpTransportClientOpts = "TIMEOUT"
)

// HttpTransportClientConfig is the configuration for the HTTP transport client
type HttpTransportClientConfig struct {
	// BaseURL is the base URL for the API
	// Example: https://api.example.com
	BaseURL string
	// AuthProvider is the provider for the authentication
	AuthProvider HttpTransportAuthProvider
	// Role is the role of the client
	Role string
	// Version is the version of the client
	Version rorversion.RorVersion
}

type HttpTransportAuthProvider interface {
	AddAuthHeaders(req *http.Request)
}

type HttpTransportClient struct {
	Client *http.Client
	Config *HttpTransportClientConfig
	Status *HttpTransportClientStatus
}

func (t *HttpTransportClient) GetJSON(path string, out any, params ...HttpTransportClientParams) error {
	return t.GetJSONWithContext(context.TODO(), path, out, params...)
}

func (t *HttpTransportClient) GetJSONWithContext(ctx context.Context, path string, out any, params ...HttpTransportClientParams) error {
	if err := t.PreflightCheck(); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "GET", t.Config.BaseURL+path, nil)
	if err != nil {
		return err
	}

	t.AddCommonHeaders(req)
	t.ParseParams(req, params...)

	res, err := t.Client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode > 399 || res.StatusCode < 200 {
		return fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
	}
	defer res.Body.Close()

	err = t.handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

func (t *HttpTransportClient) PostJSON(path string, in any, out any, params ...HttpTransportClientParams) error {
	return t.PostJSONWithContext(context.TODO(), path, in, out, params...)
}

func (t *HttpTransportClient) PostJSONWithContext(ctx context.Context, path string, in any, out any, params ...HttpTransportClientParams) error {
	if err := t.PreflightCheck(); err != nil {
		return err
	}

	jsonData, err := json.Marshal(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", t.Config.BaseURL+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	t.AddCommonHeaders(req)
	t.ParseParams(req, params...)
	t.Client.Timeout = time.Second * 60
	res, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = t.handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

func (t *HttpTransportClient) PutJSON(path string, in any, out any, params ...HttpTransportClientParams) error {
	return t.PutJSONWithContext(context.TODO(), path, in, out, params...)
}

func (t *HttpTransportClient) PutJSONWithContext(ctx context.Context, path string, in any, out any, params ...HttpTransportClientParams) error {
	if err := t.PreflightCheck(); err != nil {
		return err
	}
	jsonData, err := json.Marshal(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", t.Config.BaseURL+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	t.AddCommonHeaders(req)
	t.ParseParams(req, params...)

	res, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = t.handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

func (t *HttpTransportClient) Delete(path string, out any, params ...HttpTransportClientParams) error {
	return t.DeleteWithContext(context.TODO(), path, out, params...)
}

func (t *HttpTransportClient) DeleteWithContext(ctx context.Context, path string, out any, params ...HttpTransportClientParams) error {
	if err := t.PreflightCheck(); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "DELETE", t.Config.BaseURL+path, nil)
	if err != nil {
		return err
	}

	t.AddCommonHeaders(req)
	t.ParseParams(req, params...)

	res, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = t.handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

// Head makes a HEAD request with the given path and params.
// It returns only the header and status code from the result, as it expects no body in return.
func (t *HttpTransportClient) Head(path string, params ...HttpTransportClientParams) (http.Header, int, error) {
	return t.HeadWithContext(context.TODO(), path, params...)
}

func (t *HttpTransportClient) HeadWithContext(ctx context.Context, path string, params ...HttpTransportClientParams) (http.Header, int, error) {
	if err := t.PreflightCheck(); err != nil {
		return nil, -1, err
	}
	req, err := http.NewRequestWithContext(ctx, "HEAD", t.Config.BaseURL+path, nil)
	if err != nil {
		return nil, -1, err
	}

	t.AddCommonHeaders(req)
	t.ParseParams(req, params...)

	res, err := t.Client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer res.Body.Close()

	err = t.handleResponse(res, nil)
	if err != nil {
		return nil, -1, err
	}

	return res.Header, res.StatusCode, nil
}

// PreflightCheck by checking if retry-after is set client.Status.RetryAfter
func (t *HttpTransportClient) PreflightCheck() error {
	if !t.Status.RetryAfter.IsZero() {
		if t.Status.RetryAfter.After(time.Now()) {
			return fmt.Errorf("preflight failed, retry after is set and not expired: %s", t.Status.RetryAfter.Format(time.RFC3339))
		}
		t.Status.RetryAfter = time.Time{} // Reset retry after if it has expired
	}
	return nil
}

// getRetryAfterHeader retrieves the "Retry-After" header from the response.
// It returns a time.Time value indicating when the client should retry the request.
// The header can be in either RFC1123 format or as an integer representing seconds.
// If the header is not present, it defaults to the variable DefaultTimeout (60 seconds from now).
// This is useful for handling rate limiting or service unavailability scenarios.
func (t *HttpTransportClient) getRetryAfterHeader(res *http.Response) time.Time {
	retryAfter := res.Header.Get("Retry-After")

	if retryAfter == "" {
		// Return default offset of 60 seconds if no Retry-After header found
		return time.Now().Add(DefaultTimeout)
	}

	// Try to parse as RFC1123 time format first
	retryAfterTime, err := time.Parse(time.RFC1123, retryAfter)
	if err == nil {
		return retryAfterTime
	}

	// Try to parse as seconds (integer)
	seconds, err := strconv.Atoi(retryAfter)
	if err != nil {
		// If both parsing attempts fail, return default offset
		return time.Now().Add(DefaultTimeout)
	}

	// Convert seconds to time offset from now
	return time.Now().Add(time.Duration(seconds) * time.Second)
}

func (t *HttpTransportClient) HandleNonOk(res *http.Response) error {
	if res.StatusCode > 399 || res.StatusCode < 200 {

		if res.StatusCode == http.StatusServiceUnavailable || res.StatusCode == http.StatusTooManyRequests {
			t.Status.RetryAfter = t.getRetryAfterHeader(res)
			return fmt.Errorf("http error: %s from %s (retry after: %s)", res.Status, res.Request.URL, t.Status.RetryAfter.Format(time.RFC3339))
		}
		if res.StatusCode == http.StatusUnauthorized {
			t.Status.RetryAfter = time.Now().Add(DefaultTimeout) // Default retry after 60 seconds for unauthorized
			return fmt.Errorf("http error: %s from %s (unauthorized, check your authentication) Connection throttled", res.Status, res.Request.URL)
		}
		if res.StatusCode == http.StatusForbidden {
			t.Status.RetryAfter = time.Now().Add(DefaultTimeout) // Default retry after 60 seconds for forbidden
			return fmt.Errorf("http error: %s from %s (forbidden, check your permissions) Connection throttled", res.Status, res.Request.URL)
		}
		return fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
	}
	return nil
}

// postflightCheck is a placeholder for any postflight checks that might be needed after a request.
func (t *HttpTransportClient) postflightCheck(res *http.Response) error {
	if err := t.HandleNonOk(res); err != nil {
		return err
	}
	t.Status.ApiVersion = res.Header.Get("x-ror-version")
	if t.Status.ApiVersion == "" {
		rlog.Warn("no x-ror-version header found in response")
	}
	t.Status.LibVersion = res.Header.Get("x-ror-libver")
	if t.Status.LibVersion == "" {
		rlog.Warn("no x-ror-libver header found in response")
	}
	// If the response is successful, reset the retry after status
	t.Status.RetryAfter = time.Time{} // Reset retry after if the request was successful
	t.Status.Established = true

	return nil
}

// handleResponse processes the HTTP response, checking for errors and unmarshalling the body into the provided output variable.
// It handles both JSON and plain text responses, ensuring the output variable is a pointer.
// If the response is successful (2xx status code), it reads the body and unmarshals it into the provided output variable.
// If the response is not successful, it checks for errors and returns an appropriate error message.
func (t *HttpTransportClient) handleResponse(res *http.Response, out any) error {

	if err := t.postflightCheck(res); err != nil {
		return err
	}

	// If no output is expected, return early
	if out == nil && res.StatusCode == http.StatusNoContent {
		return nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.Header.Get("Content-Type") == "text/plain" {
		v := reflect.ValueOf(out)
		if v.Kind() != reflect.Ptr || v.IsNil() {
			return fmt.Errorf("out must be a pointer and not nil")
		}
		// this might panic
		v.Elem().Set(reflect.ValueOf(string(body)))
		return nil
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}

func (t *HttpTransportClient) AddAuthHeaders(req *http.Request) {
	t.Config.AuthProvider.AddAuthHeaders(req)
}

// AddCommonHeaders adds common headers to the request
func (t *HttpTransportClient) AddCommonHeaders(req *http.Request) {
	req.Header.Set("User-Agent", fmt.Sprintf("%s - v%s (%s)", t.Config.Role, t.Config.Version.GetVersion(), t.Config.Version.GetCommit()))
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Content-Type", `application/json`)
}

func (t *HttpTransportClient) ParseParams(req *http.Request, params ...HttpTransportClientParams) {
	var noAuth bool
	if len(params) != 0 {
		for _, param := range params {
			switch param.Key {
			case HttpTransportClientOptsNoAuth:
				noAuth = true
			case HttpTransportClientOptsHeaders:
				for key, value := range param.Value.(map[string]string) {
					req.Header.Add(key, value)
				}
			case HttpTransportClientOptsQuery:
				q := req.URL.Query()
				for key, value := range param.Value.(map[string]string) {
					q.Add(key, value)
				}
				req.URL.RawQuery = q.Encode()
			case HttpTransportClientTimeout:
				t.Client.Timeout = param.Value.(time.Duration)
			}

		}
	}
	if !noAuth {
		t.AddAuthHeaders(req)
	}
}
