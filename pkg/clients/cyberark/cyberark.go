// Description: This file contains the Cyberark client and its methods to authenticate, get secrets and passwords from Cyberark.
// TODO: Remove NHN specific code and make it generic for all Cyberark users.
package cyberark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ror/cmd/cli/config"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CyberarkClient struct {
	client http.Client
	Url    string
	token  string
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

func NewCyberarkClient(url string) *CyberarkClient {

	cyberarkclient := CyberarkClient{
		client: http.Client{},
		Url:    url,
		token:  "",
	}

	return &cyberarkclient
}

func (c *CyberarkClient) Ping() bool {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(c.Url)
	return err == nil
}

func (c *CyberarkClient) SetToken(token string) {
	c.token = token
}

func (c *CyberarkClient) Authenticate(username, password string) error {

	authreq := CyberarkAuthenticationResquest{
		Username:          username,
		Password:          password,
		ConcurrentSession: false,
		Verify:            "false",
		Timeout:           "7200",
	}
	reqjson, err := json.Marshal(authreq)
	if err != nil {
		return err
	}

	res, err := c.client.Post(c.Url+"/PasswordVault/API/auth/RADIUS/Logon/", "application/json", bytes.NewBuffer(reqjson))
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
	token := string(body)
	token, _ = strings.CutPrefix(token, "\"")
	token, _ = strings.CutSuffix(token, "\"")
	c.token = token

	timeout, err := strconv.ParseInt(authreq.Timeout, 10, 64)
	if err != nil {
		timeout = 0
	}
	viper.Set(config.RorAuthCyberarkToken, c.token)
	viper.Set(config.RorAuthCyberarkExpires, time.Now().Add(time.Duration(timeout)*time.Second))
	viper.WriteConfig()

	return nil
}

func (c *CyberarkClient) GetSecret(id string) (*CyberarkSecret, error) {
	secrets, _ := c.GetSecrets()
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
	cobra.CheckErr(err)

	req.Header.Set("Authorization", c.token)
	req.Header.Set("Accept", "application/json")
	res, err := c.client.Do(req)
	cobra.CheckErr(err)

	if res.StatusCode > 399 || res.StatusCode < 200 {
		cobra.CheckErr(fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	cobra.CheckErr(err)

	err = json.Unmarshal(body, &out)
	cobra.CheckErr(err)

	for _, secret := range out.Value {
		if slices.Contains(config.CyberarkValidDomains, secret.Address) {
			ret = append(ret, secret)
		}

	}

	return &ret, nil
}

func (c *CyberarkClient) GetPassword(id string) (string, error) {

	request := CyberarkPasswordRequest{
		Reason: "ror-cli",
	}

	jsonData, err := json.Marshal(request)
	cobra.CheckErr(err)

	req, err := http.NewRequest("POST", c.Url+"/PasswordVault/api/Accounts/"+id+"/Password/Retrieve", bytes.NewBuffer(jsonData))
	cobra.CheckErr(err)

	req.Header.Set("Authorization", c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.client.Do(req)
	cobra.CheckErr(err)

	if res.StatusCode > 399 || res.StatusCode < 200 {
		cobra.CheckErr(fmt.Errorf("http error: %s from %s", res.Status, res.Request.URL))
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	cobra.CheckErr(err)

	ret := string(body)
	ret, _ = strings.CutPrefix(ret, "\"")
	ret, _ = strings.CutSuffix(ret, "\"")

	return ret, nil
}
