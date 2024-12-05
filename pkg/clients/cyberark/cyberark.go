// Description: This file contains the Cyberark client and its methods to authenticate, get secrets and passwords from Cyberark.
// TODO: Remove NHN specific code and make it generic for all Cyberark users.
package cyberark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type CyberarkClient struct {
	client       http.Client
	Url          string
	token        string
	validDomains []string
}

type CyberarkClientInterface interface {
	Authenticate(username, password string) error
}

type CyberarkAuthenticationResquest struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	ConcurrentSession bool   `json:"concurrentSession"`
	Verify            string `json:"verify"`
	Timeout           string `json:"timeout"`
}
type CyberarkSecretsResponse struct {
	Value []CyberarkSecret `json:"value"`
	Count int              `json:"count"`
}

type CyberarkSecret struct {
	CategoryModificationTime  int64  `json:"categoryModificationTime"`
	PlatformID                string `json:"platformId"`
	SafeName                  string `json:"safeName"`
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	Address                   string `json:"address"`
	UserName                  string `json:"userName"`
	SecretType                string `json:"secretType"`
	PlatformAccountProperties struct {
		LogonDomain string `json:"LogonDomain"`
		NHNtag      string `json:"NHNtag"`
	} `json:"platformAccountProperties"`
	SecretManagement struct {
		AutomaticManagementEnabled bool   `json:"automaticManagementEnabled"`
		Status                     string `json:"status"`
		LastModifiedTime           int64  `json:"lastModifiedTime"`
		LastReconciledTime         int64  `json:"lastReconciledTime"`
		LastVerifiedTime           int64  `json:"lastVerifiedTime"`
	} `json:"secretManagement"`
	CreatedTime int64 `json:"createdTime"`
	Displayname string
}

type CyberarkPasswordRequest struct {
	Reason string `json:"reason"`
}

func NewCyberarkClient(url string, validDomains ...string) (*CyberarkClient, error) {

	cyberarkclient := CyberarkClient{
		client: http.Client{},
		Url:    url,
		token:  "",
	}
	if len(validDomains) == 0 {
		return nil, fmt.Errorf("no valid domains provided")
	}
	cyberarkclient.validDomains = validDomains

	return &cyberarkclient, nil
}

func (c *CyberarkClient) Ping() bool {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(c.Url)
	if err != nil {
		rlog.Error("Could not ping Cyberark", err, rlog.String("url", c.Url))
	}
	return err == nil
}

func (c *CyberarkClient) SetToken(token string) {
	c.token = token
}

func (c *CyberarkClient) Authenticate(username, password string) (string, time.Time, error) {
	expires := time.Now()

	authreq := CyberarkAuthenticationResquest{
		Username:          username,
		Password:          password,
		ConcurrentSession: false,
		Verify:            "false",
		Timeout:           "7200",
	}
	reqjson, err := json.Marshal(authreq)
	if err != nil {
		return "", expires, err
	}

	res, err := c.client.Post(c.Url+"/PasswordVault/API/auth/RADIUS/Logon/", "application/json", bytes.NewBuffer(reqjson))
	if err != nil {
		return "", expires, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 || res.StatusCode < 200 {
		return "", expires, fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", expires, err
	}
	token := string(body)
	token, _ = strings.CutPrefix(token, "\"")
	token, _ = strings.CutSuffix(token, "\"")
	c.token = token

	timeout, err := strconv.ParseInt(authreq.Timeout, 10, 64)
	if err != nil {
		timeout = 0
	}

	expires = time.Now().Add(time.Duration(timeout) * time.Second)

	return c.token, expires, nil
}

func (c *CyberarkClient) GetSecret(id string) (*CyberarkSecret, error) {
	secrets, err := c.GetSecrets()
	if err != nil {
		return nil, err
	}
	user := strings.Split(id, "@")
	if len(user) != 2 {
		return nil, fmt.Errorf("not a valid PAM user")
	}

	for _, secret := range *secrets {
		if secret.Address == user[1] && secret.UserName == user[0] {
			return &secret, nil
		}
	}
	return nil, fmt.Errorf("could not find PAM user")
}

func (c *CyberarkClient) GetSecrets() (*[]CyberarkSecret, error) {
	var out CyberarkSecretsResponse
	var ret []CyberarkSecret
	req, err := http.NewRequest("GET", c.Url+"/PasswordVault/api/Accounts", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.token)
	req.Header.Set("Accept", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 399 || res.StatusCode < 200 {

		return nil, fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	if len(c.validDomains) > 0 {
		for _, secret := range out.Value {
			if slices.Contains(c.validDomains, secret.Address) {
				ret = append(ret, secret)
			}

		}
	} else {
		ret = out.Value
	}

	return &ret, nil
}

func (c *CyberarkClient) GetPassword(id string) (string, error) {

	request := CyberarkPasswordRequest{
		Reason: "ror-cli",
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.Url+"/PasswordVault/api/Accounts/"+id+"/Password/Retrieve", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode > 399 || res.StatusCode < 200 {
		return "", fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	ret := string(body)
	ret, _ = strings.CutPrefix(ret, "\"")
	ret, _ = strings.CutSuffix(ret, "\"")

	return ret, nil
}
