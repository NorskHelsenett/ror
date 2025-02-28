package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
)

type HttpTransportClientParams struct {
	Key   HttpTransportClientOpts
	Value any
}

type HttpTransportClientOpts string

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
}

func (t *HttpTransportClient) GetJSON(path string, out any, params ...HttpTransportClientParams) error {
	req, err := http.NewRequest("GET", t.Config.BaseURL+path, nil)
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

	err = handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

func (t *HttpTransportClient) PostJSON(path string, in any, out any, params ...HttpTransportClientParams) error {
	jsonData, err := json.Marshal(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", t.Config.BaseURL+path, bytes.NewBuffer(jsonData))
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

	err = handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

func (t *HttpTransportClient) PutJSON(path string, in any, out any, params ...HttpTransportClientParams) error {
	jsonData, err := json.Marshal(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", t.Config.BaseURL+path, bytes.NewBuffer(jsonData))
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

	err = handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

func (t *HttpTransportClient) Delete(path string, out any, params ...HttpTransportClientParams) error {
	req, err := http.NewRequest("DELETE", t.Config.BaseURL+path, nil)
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

	err = handleResponse(res, out)
	if err != nil {
		return err
	}

	return nil
}

// Head makes a HEAD request with the given path and params.
// It returns only the header and status code from the result, as it expects no body in return.
func (t *HttpTransportClient) Head(path string, params ...HttpTransportClientParams) (http.Header, int, error) {
	req, err := http.NewRequest("HEAD", t.Config.BaseURL+path, nil)
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

	return res.Header, res.StatusCode, nil
}

func handleResponse(res *http.Response, out any) error {

	if res.StatusCode > 399 || res.StatusCode < 200 {
		return fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
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
