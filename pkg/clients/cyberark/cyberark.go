// Package cyberark provides functionality for interacting with CyberArk's password vault API.
// It enables authentication, retrieval of secrets, and password management for user accounts.
//
// The package filters secrets based on domains provided during client initialization and
// supports multiple authentication methods including RADIUS (default), LDAP, Windows, and CyberArk.
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

// CyberArkAuthMethod represents an authentication method for CyberArk.
type CyberArkAuthMethod string

// Authentication method constants for CyberArk authentication.
const (
	CyberArkAuthMethodCyberArk CyberArkAuthMethod = "Cyberark" // Native CyberArk authentication
	CyberArkAuthMethodWindows  CyberArkAuthMethod = "Windows"  // Windows authentication
	CyberArkAuthMethodRadius   CyberArkAuthMethod = "RADIUS"   // RADIUS authentication (default)
	CyberArkAuthMethodLDAP     CyberArkAuthMethod = "LDAP"     // LDAP authentication
)

// CyberarkClient is the main structure for interfacing with the CyberArk API.
// It handles authentication, session management, and API operations.
type CyberarkClient struct {
	client       http.Client        // HTTP client for making API requests
	Url          string             // Base URL of CyberArk API
	token        string             // Authentication token
	method       CyberArkAuthMethod // Authentication method to use
	validDomains []string           // List of domains to filter secrets by
}

// String returns the string representation of the CyberArkAuthMethod.
func (c *CyberArkAuthMethod) String() string {
	return string(*c)
}

// CyberarkClientInterface defines the authentication interface for a CyberArk client.
type CyberarkClientInterface interface {
	Authenticate(username, password string) error
}

// CyberarkAuthenticationResquest represents the payload for authenticating with CyberArk.
type CyberarkAuthenticationResquest struct {
	Username          string `json:"username"`          // Username for authentication
	Password          string `json:"password"`          // Password for authentication
	ConcurrentSession bool   `json:"concurrentSession"` // Whether to allow concurrent sessions
	Verify            string `json:"verify"`            // Whether to verify the authentication
	Timeout           string `json:"timeout"`           // Session timeout in seconds
}

// CyberarkSecretsResponse represents the response structure when retrieving secrets from CyberArk.
type CyberarkSecretsResponse struct {
	Value []CyberarkSecret `json:"value"` // List of secrets
	Count int              `json:"count"` // Number of secrets returned
}

// CyberarkSecret represents a secret stored in the CyberArk vault.
// It contains account information and related metadata.
type CyberarkSecret struct {
	CategoryModificationTime  int64  `json:"categoryModificationTime"` // Time when category was last modified
	PlatformID                string `json:"platformId"`               // Platform identifier
	SafeName                  string `json:"safeName"`                 // Name of the safe containing the secret
	ID                        string `json:"id"`                       // Unique identifier for the secret
	Name                      string `json:"name"`                     // Display name of the secret
	Address                   string `json:"address"`                  // Domain address associated with the secret
	UserName                  string `json:"userName"`                 // Username for the account
	SecretType                string `json:"secretType"`               // Type of the secret
	PlatformAccountProperties struct {
		LogonDomain string `json:"LogonDomain"` // Domain for logon
		NHNtag      string `json:"NHNtag"`      // NHN-specific tag
	} `json:"platformAccountProperties"`
	SecretManagement struct {
		AutomaticManagementEnabled bool   `json:"automaticManagementEnabled"` // Whether automatic management is enabled
		Status                     string `json:"status"`                     // Status of secret management
		LastModifiedTime           int64  `json:"lastModifiedTime"`           // Last time secret was modified
		LastReconciledTime         int64  `json:"lastReconciledTime"`         // Last time secret was reconciled
		LastVerifiedTime           int64  `json:"lastVerifiedTime"`           // Last time secret was verified
	} `json:"secretManagement"`
	CreatedTime int64  `json:"createdTime"` // Time when secret was created
	Displayname string // Display name for the secret
}

// CyberarkPasswordRequest represents the request payload for retrieving a password.
type CyberarkPasswordRequest struct {
	Reason string `json:"reason"` // Reason for accessing the password
}

// NewCyberarkClient creates a new CyberArk client with RADIUS authentication method (default).
//
// Parameters:
//   - url: Base URL for the CyberArk API
//   - validDomains: List of domains to filter secrets by (at least one required)
//
// Returns:
//   - *CyberarkClient: New CyberArk client
//   - error: Error if no valid domains are provided
func NewCyberarkClient(url string, validDomains ...string) (*CyberarkClient, error) {
	cyberarkclient := CyberarkClient{
		client: http.Client{},
		Url:    url,
		token:  "",
		method: CyberArkAuthMethodRadius,
	}
	if len(validDomains) == 0 {
		return nil, fmt.Errorf("no valid domains provided")
	}
	cyberarkclient.validDomains = validDomains

	return &cyberarkclient, nil
}

// NewCyberarkClientWithMethod creates a new CyberArk client with the specified authentication method.
//
// Parameters:
//   - url: Base URL for the CyberArk API
//   - method: Authentication method to use (CyberArkAuthMethodRadius, CyberArkAuthMethodLDAP, etc.)
//   - validDomains: List of domains to filter secrets by (at least one required)
//
// Returns:
//   - *CyberarkClient: New CyberArk client
//   - error: Error if no valid domains are provided
func NewCyberarkClientWithMethod(url string, method CyberArkAuthMethod, validDomains ...string) (*CyberarkClient, error) {
	cyberarkclient := CyberarkClient{
		client: http.Client{},
		Url:    url,
		token:  "",
		method: method,
	}
	if len(validDomains) == 0 {
		return nil, fmt.Errorf("no valid domains provided")
	}
	cyberarkclient.validDomains = validDomains

	return &cyberarkclient, nil
}

// Ping checks if the CyberArk service is available.
//
// Returns:
//   - bool: true if service is reachable, false otherwise
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

// SetToken sets the authentication token for the CyberArk client.
//
// Parameters:
//   - token: Authentication token to set
func (c *CyberarkClient) SetToken(token string) {
	c.token = token
}

// Authenticate performs authentication against the CyberArk API using the configured method.
//
// Parameters:
//   - username: Username for authentication
//   - password: Password for authentication
//
// Returns:
//   - string: Authentication token if successful
//   - time.Time: Token expiration time
//   - error: Error if authentication fails
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

	res, err := c.client.Post(c.Url+"/PasswordVault/API/auth/"+c.method.String()+"/Logon/", "application/json", bytes.NewBuffer(reqjson))
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

// GetSecret retrieves a specific secret by its ID.
//
// The ID must be in the format "username@domain". The function searches through
// all available secrets and returns the one matching both username and domain.
//
// Parameters:
//   - id: Secret identifier in the format "username@domain"
//
// Returns:
//   - *CyberarkSecret: Pointer to the matching secret if found
//   - error: Error if secret not found or retrieval fails
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

// GetSecrets retrieves all secrets available to the authenticated user,
// filtered by the valid domains specified during client initialization.
//
// Returns:
//   - *[]CyberarkSecret: Pointer to the list of secrets if successful
//   - error: Error if retrieval fails
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

// GetPassword retrieves the password for a specific account ID from CyberArk.
//
// Parameters:
//   - id: The unique identifier of the account to retrieve the password for
//
// Returns:
//   - string: The password if successful
//   - error: Error if retrieval fails
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
