// Package httpclient provides an HTTP transport implementation for interacting with ROR APIs.
// It offers various HTTP methods (GET, POST, PUT, DELETE, HEAD) with JSON serialization/deserialization,
// authentication handling, retry mechanisms, and common HTTP client features.
//
// The client handles authentication through an injected AuthProvider interface, manages request
// timeouts, and implements rate limiting through Retry-After response headers.
package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// DefaultTimeout defines the default HTTP request timeout duration
var DefaultTimeout = 60 * time.Second

// HttpTransportClientParams encapsulates parameters for HTTP client requests
// with a key-value structure to enable flexible configuration options.
type HttpTransportClientParams struct {
	Key   HttpTransportClientOpts
	Value any
}

// HttpTransportClientOpts represents option types for HTTP client configuration.
type HttpTransportClientOpts string

// HttpTransportClientStatus tracks the connection state and version information
// for the HTTP transport client, including retry timing information.
type HttpTransportClientStatus struct {
	Established bool      `json:"established"` // Whether a connection has been established
	ApiVersion  string    `json:"api_version"` // Version of the remote API
	LibVersion  string    `json:"lib_version"` // Version of the client library
	RetryAfter  time.Time `json:"retry_after"` // Time after which to retry if rate limited
}

// Available client configuration options
const (
	HttpTransportClientOptsNoAuth  HttpTransportClientOpts = "NOAUTH"  // Skip authentication for this request
	HttpTransportClientOptsHeaders HttpTransportClientOpts = "HEADERS" // Add custom headers
	HttpTransportClientOptsQuery   HttpTransportClientOpts = "QUERY"   // Add query parameters
	HttpTransportClientTimeout     HttpTransportClientOpts = "TIMEOUT" // Set request timeout
)

// HttpTransportClientConfig defines the configuration for the HTTP transport client.
type HttpTransportClientConfig struct {
	// BaseURL is the base URL for the API, e.g., "https://api.example.com"
	BaseURL string

	// AuthProvider handles request authentication
	AuthProvider HttpTransportAuthProvider

	// Role identifies the client's purpose or role
	Role string

	// Version identifies the client version
	Version rorversion.RorVersion
}

// HttpTransportAuthProvider is an interface that authentication providers must implement.
type HttpTransportAuthProvider interface {
	// AddAuthHeaders adds authentication headers to the request
	AddAuthHeaders(req *http.Request)

	// GetApiSecret returns the API secret for authentication
	GetApiSecret() string
}

// HttpTransportClient implements an HTTP client with authentication, status tracking,
// and standardized request/response handling.
type HttpTransportClient struct {
	Client *http.Client               // Underlying HTTP client
	Config *HttpTransportClientConfig // Client configuration
	Status *HttpTransportClientStatus // Client connection status
}

// NewHttpTransportClientStatus creates a new client status object with default values.
// The client starts as not established with unknown versions and no retry restrictions.
func NewHttpTransportClientStatus() *HttpTransportClientStatus {
	return &HttpTransportClientStatus{
		Established: false,
		ApiVersion:  "",
		LibVersion:  "",
		RetryAfter:  time.Time{},
	}
}


// NewHttpTransportClientConfig creates a new configuration object for the HTTP transport client
// The constructor allows for validation of parameters like BaseURL to stop some of the faulty configuration possibilities.
//
// # BaseURL is the base URL for the API
// Example: https://api.example.com
//
// # AuthProvider is the provider for the authentication
//
// # Role is the role of the client
//
// # Version is the version of the client.
func NewHttpTransportClientConfig(baseUrl string, authProvider HttpTransportAuthProvider, role string, version rorversion.RorVersion) (*HttpTransportClientConfig, error) {

	config := HttpTransportClientConfig{
		BaseURL:      baseUrl,
		AuthProvider: authProvider,
		Role:         role,
		Version:      version,
	}

	err := config.ValidateUrl()
	if err != nil {
		return nil, fmt.Errorf("failed to validate provided url. %v", err)
	}

	return &config, nil
}

// validateUrl valides that BaseUrl provided passes at minimum a valid URL.
func (h *HttpTransportClientConfig) ValidateUrl() error {
	return validateUrl(h.BaseURL)
}

func NewHttpTransportClient(client *http.Client, config *HttpTransportClientConfig, status *HttpTransportClientStatus) *HttpTransportClient {
	hClient := HttpTransportClient{
		Client: client,
		Config: config,
		Status: status,
	}

	return &hClient
}

// GetRole returns the configured role of the client.
func (t *HttpTransportClientConfig) GetRole() string {
	return t.Role
}

// GetJSON performs a GET request and unmarshals the JSON response into the out parameter.
// This is a convenience wrapper around GetJSONWithContext using the background context.
func (t *HttpTransportClient) GetJSON(path string, out any, params ...HttpTransportClientParams) error {
	return t.GetJSONWithContext(context.TODO(), path, out, params...)
}

// GetJSONWithContext performs a GET request with the provided context and unmarshals
// the JSON response into the out parameter.
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

// PostJSON performs a POST request with the provided input and unmarshals the JSON response into the out parameter.
// This is a convenience wrapper around PostJSONWithContext using the background context.
func (t *HttpTransportClient) PostJSON(path string, in any, out any, params ...HttpTransportClientParams) error {
	return t.PostJSONWithContext(context.TODO(), path, in, out, params...)
}

// PostJSONWithContext performs a POST request with the provided context and input,
// then unmarshals the JSON response into the out parameter.
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

// PutJSON performs a PUT request with the provided input and unmarshals the JSON response into the out parameter.
// This is a convenience wrapper around PutJSONWithContext using the background context.
func (t *HttpTransportClient) PutJSON(path string, in any, out any, params ...HttpTransportClientParams) error {
	return t.PutJSONWithContext(context.TODO(), path, in, out, params...)
}

// PutJSONWithContext performs a PUT request with the provided context and input,
// then unmarshals the JSON response into the out parameter.
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

// Delete performs a DELETE request and unmarshals the JSON response into the out parameter.
// This is a convenience wrapper around DeleteWithContext using the background context.
func (t *HttpTransportClient) Delete(path string, out any, params ...HttpTransportClientParams) error {
	return t.DeleteWithContext(context.TODO(), path, out, params...)
}

// DeleteWithContext performs a DELETE request with the provided context and
// unmarshals the JSON response into the out parameter.
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
// It returns only the header and status code from the result, as HEAD requests have no response body.
// This is a convenience wrapper around HeadWithContext using the background context.
func (t *HttpTransportClient) Head(path string, params ...HttpTransportClientParams) (http.Header, int, error) {
	return t.HeadWithContext(context.TODO(), path, params...)
}

// HeadWithContext performs a HEAD request with the provided context and returns
// the response headers and status code.
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

// PreflightCheck verifies if the client is ready to make a request by checking
// if rate limiting is in effect via the RetryAfter timestamp.
func (t *HttpTransportClient) PreflightCheck() error {
	// if t.Status == nil {
	// 	t.Status = NewHttpTransportClientStatus()
	// }
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

// HandleNonOk processes HTTP responses with non-successful status codes and returns appropriate errors.
// It handles different HTTP error status codes with specific behavior:
//
// - For 503 (Service Unavailable) and 429 (Too Many Requests): Sets retry-after time from response headers
// - For 401 (Unauthorized): Sets default retry timeout and returns authentication error message
// - For 403 (Forbidden): Sets default retry timeout and returns permission error message
// - For other 4xx/5xx errors: Returns generic HTTP error
//
// The function updates the client's Status.RetryAfter field for rate limiting and throttling scenarios.
// Returns nil if the response status code indicates success (200-399).
//
// Parameters:
//   - res: HTTP response to process
//
// Returns:
//   - error: Formatted error message for non-successful responses, nil for successful responses
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
	//TODO: add constants for header names
	t.Status.ApiVersion = res.Header.Get("x-ror-version")
	t.Status.LibVersion = res.Header.Get("x-ror-libver")
	// If the response is successful, reset the retry after status
	t.Status.RetryAfter = time.Time{}
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
		if v.Kind() != reflect.Pointer || v.IsNil() {
			return fmt.Errorf("out must be a pointer and not nil")
		}
		if v.Elem().Kind() != reflect.String {
			rlog.Infof("something went wrong, server returned text/plain (%s) but we expected a %s", string(body), v.Elem().Kind().String())
			return fmt.Errorf("out must be a pointer to a string as the content type is text/plain")
		}
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

func (t *HttpTransportClientStatus) IsEstablished() bool {
	return t.Established
}
func (t *HttpTransportClientStatus) GetApiVersion() string {
	return t.ApiVersion
}
func (t *HttpTransportClientStatus) GetLibVersion() string {
	return t.LibVersion
}

func validateUrl(baseUrl string) error {
	_, err := url.ParseRequestURI(baseUrl)
	return err
}

// Example of creating and using the HTTP transport client:
//
// config := &httpclient.HttpTransportClientConfig{
//     BaseURL:      "https://api.example.com",
//     AuthProvider: myAuthProvider,
//     Role:         "api-client",
//     Version:      myRorVersion,
// }
//
// client := &httpclient.HttpTransportClient{
//     Client: &http.Client{Timeout: httpclient.DefaultTimeout},
//     Config: config,
//     Status: httpclient.NewHttpTransportClientStatus(),
// }
//
// var response MyResponseType
// err := client.GetJSON("/endpoint", &response, httpclient.HttpTransportClientParams{
//     Key:   httpclient.HttpTransportClientOptsHeaders,
//     Value: map[string]string{"X-Custom-Header": "value"},
//
