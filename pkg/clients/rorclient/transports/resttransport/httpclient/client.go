package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
)

type HttpTransportClientConfig struct {
	BaseURL      string
	AuthProvider HttpTransportAuthProvider
	Role         string
	Version      rorversion.RorVersion
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return fmt.Errorf("empty response")
	}

	err = json.Unmarshal(body, out)
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

	res, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 || res.StatusCode < 200 {
		return fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
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

	if res.StatusCode > 399 || res.StatusCode < 200 {
		return fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
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
			}
		}
	}
	if !noAuth {
		t.AddAuthHeaders(req)
	}
}
