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

	// API endpoint paths
	baseAPIPath             = "/PasswordVault/api"
	baseAuthPath            = "/PasswordVault/API/auth"
	accountsEndpoint        = baseAPIPath + "/Accounts"
	passwordRetrievePattern = accountsEndpoint + "/%s/Password/Retrieve"
	authLoginPattern        = baseAuthPath + "/%s/Logon"

	// Default settings
	defaultTimeout        = 7200      // Default timeout in seconds
	defaultPingTimeout    = 5         // Default ping timeout in seconds
	defaultPasswordReason = "ror-cli" // Default reason for password requests
)

// CyberArkClient is the main structure for interfacing with the CyberArk API.
// It handles authentication, session management, and API operations.
type CyberArkClient struct {
	client       http.Client        // HTTP client for making API requests
	baseURL      string             // Base URL of CyberArk API
	token        string             // Authentication token
	method       CyberArkAuthMethod // Authentication method to use
	validDomains []string           // List of domains to filter secrets by
}

// String returns the string representation of the CyberArkAuthMethod.
func (c *CyberArkAuthMethod) String() string {
	return string(*c)
}

// CyberArkClientInterface defines the authentication interface for a CyberArk client.
type CyberArkClientInterface interface {
	Authenticate(username, password string) (string, time.Time, error)
	GetSecrets() (*[]CyberArkSecret, error)
	GetSecret(id string) (*CyberArkSecret, error)
	GetPassword(id string) (string, error)
	Ping() bool
}

// CyberArkAuthenticationRequest represents the payload for authenticating with CyberArk.
type CyberArkAuthenticationRequest struct {
	Username          string `json:"username"`          // Username for authentication
	Password          string `json:"password"`          // Password for authentication
	ConcurrentSession bool   `json:"concurrentSession"` // Whether to allow concurrent sessions
	Verify            string `json:"verify"`            // Whether to verify the authentication
	Timeout           string `json:"timeout"`           // Session timeout in seconds
}

// CyberArkSecretsResponse represents the response structure when retrieving secrets from CyberArk.
type CyberArkSecretsResponse struct {
	Value []CyberArkSecret `json:"value"` // List of secrets
	Count int              `json:"count"` // Number of secrets returned
}

// CyberArkSecret represents a secret stored in the CyberArk vault.
// It contains account information and related metadata.
type CyberArkSecret struct {
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
	DisplayName string // Display name for the secret
}

// CyberArkPasswordRequest represents the request payload for retrieving a password.
type CyberArkPasswordRequest struct {
	Reason string `json:"reason"` // Reason for accessing the password
}

// NewCyberArkClient creates a new CyberArk client with RADIUS authentication method (default).
//
// Parameters:
//   - url: Base URL for the CyberArk API
//   - validDomains: List of domains to filter secrets by (at least one required)
//
// Returns:
//   - *CyberArkClient: New CyberArk client
//   - error: Error if no valid domains are provided
func NewCyberArkClient(url string, validDomains ...string) (*CyberArkClient, error) {
	return newCyberArkClientInternal(url, CyberArkAuthMethodRadius, validDomains...)
}

// NewCyberArkClientWithMethod creates a new CyberArk client with the specified authentication method.
//
// Parameters:
//   - url: Base URL for the CyberArk API
//   - method: Authentication method to use (CyberArkAuthMethodRadius, CyberArkAuthMethodLDAP, etc.)
//   - validDomains: List of domains to filter secrets by (at least one required)
//
// Returns:
//   - *CyberArkClient: New CyberArk client
//   - error: Error if no valid domains are provided
func NewCyberArkClientWithMethod(url string, method CyberArkAuthMethod, validDomains ...string) (*CyberArkClient, error) {
	return newCyberArkClientInternal(url, method, validDomains...)
}

// newCyberArkClientInternal is an internal helper function for creating CyberArk clients.
// It handles validation and initialization common to all constructors.
//
// Parameters:
//   - url: Base URL for the CyberArk API
//   - method: Authentication method to use
//   - validDomains: List of domains to filter secrets by
//
// Returns:
//   - *CyberArkClient: New CyberArk client
//   - error: Error if validation fails
func newCyberArkClientInternal(url string, method CyberArkAuthMethod, validDomains ...string) (*CyberArkClient, error) {
	// Validate URL
	if url == "" {
		return nil, fmt.Errorf("empty CyberArk API URL provided")
	}

	// Validate domains
	if len(validDomains) == 0 {
		return nil, fmt.Errorf("no valid domains provided")
	}

	// Create and return client
	client := &CyberArkClient{
		client:       http.Client{},
		baseURL:      url,
		token:        "",
		method:       method,
		validDomains: validDomains,
	}

	return client, nil
}

// Ping checks if the CyberArk service is available.
//
// Returns:
//   - bool: true if service is reachable, false otherwise
func (c *CyberArkClient) Ping() bool {
	client := http.Client{
		Timeout: time.Duration(defaultPingTimeout) * time.Second,
	}

	_, err := client.Get(c.baseURL)
	if err != nil {
		rlog.Error("Could not ping CyberArk API", err, rlog.String("url", c.baseURL))
		return false
	}

	return true
}

// createRequest is a helper method that creates an HTTP request with common headers.
//
// Parameters:
//   - method: HTTP method (GET, POST, etc.)
//   - path: API path relative to base URL
//   - body: Request body (can be nil)
//
// Returns:
//   - *http.Request: Prepared HTTP request
//   - error: Error if request creation fails
func (c *CyberArkClient) createRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+path, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set common headers
	if c.token != "" {
		req.Header.Set("Authorization", c.token)
	}
	req.Header.Set("Accept", "application/json")

	// Set content type if body is provided
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// executeRequest is a helper method that executes an HTTP request and handles common error cases.
//
// Parameters:
//   - req: HTTP request to execute
//
// Returns:
//   - *http.Response: HTTP response
//   - []byte: Response body
//   - error: Error if request execution or response processing fails
func (c *CyberArkClient) executeRequest(req *http.Request) (*http.Response, []byte, error) {
	// Execute the request
	res, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			rlog.Errorc(req.Context(), "error in closing response body", err)
		}
	}()

	// Check response status
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		return res, nil, fmt.Errorf("http error: %s from %s", res.Status, req.URL)
	}

	// Read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return res, body, nil
}

// SetToken sets the authentication token for the CyberArk client.
//
// Parameters:
//   - token: Authentication token to set
func (c *CyberArkClient) SetToken(token string) {
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
func (c *CyberArkClient) Authenticate(username, password string) (string, time.Time, error) {
	// Initialize expiration time to current time (will be modified on success)
	expires := time.Now()

	// Validate inputs
	if username == "" || password == "" {
		return "", expires, fmt.Errorf("username and password must not be empty")
	}

	// Prepare authentication request
	authReq := CyberArkAuthenticationRequest{
		Username:          username,
		Password:          password,
		ConcurrentSession: false,
		Verify:            "false",
		Timeout:           strconv.Itoa(defaultTimeout),
	}

	// Convert request to JSON
	reqJSON, err := json.Marshal(authReq)
	if err != nil {
		return "", expires, fmt.Errorf("failed to marshal authentication request: %w", err)
	}

	// Construct the authentication URL and request
	authURL := fmt.Sprintf(authLoginPattern, c.method.String())
	req, err := c.createRequest(http.MethodPost, authURL, bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", expires, fmt.Errorf("failed to create authentication request: %w", err)
	}

	// Execute the request
	_, body, err := c.executeRequest(req)
	if err != nil {
		// For authentication errors, provide more context
		return "", expires, fmt.Errorf("authentication failed: %w", err)
	}

	// Extract token (remove surrounding quotes if present)
	token := strings.Trim(string(body), "\"")
	c.token = token

	// Calculate token expiration time
	timeout, err := strconv.ParseInt(authReq.Timeout, 10, 64)
	if err != nil {
		// Default to 2 hours if timeout parsing fails
		rlog.Warn("Failed to parse timeout value, using default", rlog.String("error", err.Error()))
		// Set default timeout
		timeout = defaultTimeout
	}

	// Set expiration time
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
//   - *CyberArkSecret: Pointer to the matching secret if found
//   - error: Error if secret not found or retrieval fails
func (c *CyberArkClient) GetSecret(id string) (*CyberArkSecret, error) {
	// Validate authentication state
	if c.token == "" {
		return nil, fmt.Errorf("not authenticated: token is empty")
	}

	// Validate the input format
	if id == "" {
		return nil, fmt.Errorf("empty secret identifier provided")
	}

	// Parse username and domain from the identifier
	user := strings.Split(id, "@")
	if len(user) != 2 {
		return nil, fmt.Errorf("invalid secret identifier format: expected 'username@domain', got '%s'", id)
	}

	username, domain := user[0], user[1]

	// Check if the domain is in the list of valid domains
	if !slices.Contains(c.validDomains, domain) {
		return nil, fmt.Errorf("domain '%s' is not in the list of valid domains", domain)
	}

	// Fetch all secrets
	secrets, err := c.GetSecrets()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve secrets: %w", err)
	}

	if secrets == nil || len(*secrets) == 0 {
		return nil, fmt.Errorf("no secrets found")
	}

	// Find the matching secret
	return c.findSecretByUsernameAndDomain(secrets, username, domain)
}

// findSecretByUsernameAndDomain is a helper method that searches for a secret
// with the specified username and domain in a collection of secrets.
//
// Parameters:
//   - secrets: Pointer to slice of secrets to search in
//   - username: Username to match
//   - domain: Domain to match
//
// Returns:
//   - *CyberArkSecret: Pointer to the matching secret if found
//   - error: Error if secret not found
func (c *CyberArkClient) findSecretByUsernameAndDomain(secrets *[]CyberArkSecret, username, domain string) (*CyberArkSecret, error) {
	for _, secret := range *secrets {
		if secret.Address == domain && secret.UserName == username {
			// Create a copy to avoid issues with the loop variable
			matchedSecret := secret
			return &matchedSecret, nil
		}
	}

	return nil, fmt.Errorf("secret not found for user '%s' in domain '%s'", username, domain)
}

// GetSecrets retrieves all secrets available to the authenticated user,
// filtered by the valid domains specified during client initialization.
//
// Returns:
//   - *[]CyberArkSecret: Pointer to the list of secrets if successful
//   - error: Error if retrieval fails or if not authenticated
func (c *CyberArkClient) GetSecrets() (*[]CyberArkSecret, error) {
	// Validate authentication state
	if c.token == "" {
		return nil, fmt.Errorf("not authenticated: token is empty")
	}

	// Create request to fetch accounts
	req, err := c.createRequest(http.MethodGet, accountsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// Execute the request
	_, body, err := c.executeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get secrets: %w", err)
	}

	// Parse response
	var response CyberArkSecretsResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Filter secrets by valid domains if necessary
	return c.filterSecretsByDomain(response.Value), nil
}

// filterSecretsByDomain filters the secrets based on the client's valid domains.
// This is a helper method extracted from GetSecrets to improve readability.
//
// Parameters:
//   - secrets: Slice of CyberArkSecret from the API response
//
// Returns:
//   - *[]CyberArkSecret: Pointer to filtered slice of secrets
func (c *CyberArkClient) filterSecretsByDomain(secrets []CyberArkSecret) *[]CyberArkSecret {
	// If no valid domains specified, return all secrets
	if len(c.validDomains) == 0 {
		return &secrets
	}

	// Filter secrets based on domain
	var filteredSecrets []CyberArkSecret
	for _, secret := range secrets {
		if slices.Contains(c.validDomains, secret.Address) {
			filteredSecrets = append(filteredSecrets, secret)
		}
	}

	return &filteredSecrets
}

// GetPassword retrieves the password for a specific account ID from CyberArk.
//
// Parameters:
//   - id: The unique identifier of the account to retrieve the password for
//
// Returns:
//   - string: The password if successful
//   - error: Error if retrieval fails or if not authenticated
func (c *CyberArkClient) GetPassword(id string) (string, error) {
	// Validate authentication state
	if c.token == "" {
		return "", fmt.Errorf("not authenticated: token is empty")
	}

	// Validate input
	if id == "" {
		return "", fmt.Errorf("empty account ID provided")
	}

	// Prepare request payload
	request := CyberArkPasswordRequest{
		Reason: defaultPasswordReason,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal password request: %w", err)
	}

	// Create request using helper method
	passwordPath := fmt.Sprintf(passwordRetrievePattern, id)
	req, err := c.createRequest(http.MethodPost, passwordPath, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// Execute the request
	_, body, err := c.executeRequest(req)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve password: %w", err)
	}

	// Process the response - trim surrounding quotes
	password := strings.Trim(string(body), "\"")

	return password, nil
}
