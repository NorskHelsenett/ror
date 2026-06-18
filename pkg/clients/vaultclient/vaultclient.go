package vaultclient

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/hashicorp/vault-client-go"
)

var vaultClient *VaultClient = &VaultClient{}

const (
	// vaultInitialBackoff is the wait time before the first connection retry.
	vaultInitialBackoff = 1 * time.Second
	// vaultMaxBackoff caps the exponential backoff between connection retries.
	vaultMaxBackoff = 30 * time.Second
)

type VaultClient struct {
	Context context.Context
	Client  *vault.Client
	Url     string
}

type VaultCredsHelper interface {
	GetToken() string
	Login(vc *VaultClient) error
}

// creates a new vault client
//
// The connected client stores a long-lived context (context.Background) rather
// than the caller's ctx. VaultClient.Context is used for the lifetime of the
// client: runtime secret reads and the background token-renewal goroutine. The
// caller's ctx is typically a bounded startup context that is cancelled once
// connection setup completes (e.g. InitConnections' deferred cancel), so
// storing it would break every later operation with "context canceled". The
// startup connect/login timeout is instead enforced by the caller's retry loop
// and the per-request timeout configured on the client.
func New(ctx context.Context, credsHelper VaultCredsHelper, url string) (*VaultClient, error) {
	client, err := getClient(url)
	if err != nil {
		return nil, err
	}

	vaultClient := VaultClient{
		Context: context.Background(),
		Client:  client,
	}

	err = credsHelper.Login(&vaultClient)
	if err != nil {
		return nil, err
	}

	return &vaultClient, nil
}

// vaultStartupChecker is a health checker that tracks the vault connection
// state while it is being established. Before a connection succeeds it reports
// StatusFail so the health endpoint clearly shows vault as the dependency that
// is blocking startup. Once a client is set it delegates to the live client's
// ping so the check reflects the real connection state.
type vaultStartupChecker struct {
	mu     sync.RWMutex
	client *VaultClient
}

func (c *vaultStartupChecker) setClient(client *VaultClient) {
	c.mu.Lock()
	c.client = client
	c.mu.Unlock()
}

func (c *vaultStartupChecker) CheckHealth(ctx context.Context) []rorhealth.Check {
	c.mu.RLock()
	client := c.client
	c.mu.RUnlock()

	if client == nil {
		return []rorhealth.Check{{
			Status: rorhealth.StatusFail,
			Output: "Connecting to vault",
		}}
	}
	return client.CheckHealth(ctx)
}

func getClient(vaultUrl string) (*vault.Client, error) {
	client, err := vault.New(
		vault.WithAddress(vaultUrl),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (rc VaultClient) CheckHealth(_ context.Context) []rorhealth.Check {
	return rc.CheckHealthWithoutContext()
}

func (rc VaultClient) CheckHealthWithoutContext() []rorhealth.Check {
	c := rorhealth.Check{
		Status: rorhealth.StatusPass,
		Output: "Vault connection ok",
	}
	ok, err := rc.Ping()
	if !ok {
		rlog.Error("vault ping returned error", err)
		c.Status = rorhealth.StatusFail
		c.Output = "Could not connect to vault"
	}

	return []rorhealth.Check{c}
}

func (v VaultClient) Ping() (bool, error) {
	if v.Client == nil {
		err := fmt.Errorf("vault client is not initialized")
		rlog.Error("could not ping vault", err)
		return false, err
	}
	_, err := v.Client.Auth.TokenLookUpSelf(v.Context)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (v VaultClient) PingWithContext(_ context.Context) (bool, error) {
	if v.Client == nil {
		err := fmt.Errorf("vault client is not initialized")
		rlog.Error("could not ping vault", err)
		return false, err
	}
	_, err := v.Client.Auth.TokenLookUpSelf(v.Context)
	if err != nil {
		return false, err
	}

	return true, nil
}

// This is a opinionated way to create a new vault client
// Might be better to migrate to New() a factory function that takes a VaultCredsHelper
// and a url as arguments.
// Migth deprecate this function in the future
//
// NewVaultClient retries the connection with exponential backoff until it
// succeeds, blocking until a usable client is returned. It never returns nil.
func NewVaultClient(role string, url string) *VaultClient {
	return NewVaultClientWithContext(context.Background(), role, url)
}

// NewVaultClientWithContext creates a vault client, retrying the connection and
// login with exponential backoff until they succeed or the context is
// cancelled. Each failed attempt is logged with the vault url, attempt number
// and underlying error so failures are easy to troubleshoot.
//
// On success it returns the connected client. If the context is cancelled
// before a connection is established it returns a non-nil, unconnected client
// whose health checks report unhealthy, so callers never receive nil.
func NewVaultClientWithContext(ctx context.Context, role string, url string) *VaultClient {
	credsHelper := newVaultCredsHelper(role)

	// Register a health checker up front so the health endpoint reports vault
	// as unhealthy ("Connecting to vault") while the retry loop runs, instead of
	// the dependency being invisible until it finally connects.
	checker := &vaultStartupChecker{}
	rorhealth.Register(ctx, "vault", checker)

	backoff := vaultInitialBackoff
	for attempt := 1; ; attempt++ {
		client, err := New(ctx, credsHelper, url)
		if err == nil {
			if attempt > 1 {
				rlog.Info("connected to vault",
					rlog.String("url", url),
					rlog.Int("attempts", attempt))
			}
			checker.setClient(client)
			return client
		}

		rlog.Error("could not connect to vault, retrying", err,
			rlog.String("url", url),
			rlog.Int("attempt", attempt),
			rlog.String("retryIn", backoff.String()))

		select {
		case <-ctx.Done():
			rlog.Error("giving up connecting to vault: context cancelled", ctx.Err(),
				rlog.String("url", url),
				rlog.Int("attempts", attempt))
			unconnected := &VaultClient{Context: context.Background(), Url: url}
			checker.setClient(unconnected)
			return unconnected
		case <-time.After(backoff):
		}

		backoff *= 2
		if backoff > vaultMaxBackoff {
			backoff = vaultMaxBackoff
		}
	}
}

// MustNewVaultClientWithContext behaves like NewVaultClientWithContext but
// treats a cancelled context as fatal: instead of returning an unconnected
// client it logs the failure and exits the process. Use this when a vault
// connection is a hard prerequisite and the process must not continue without
// it.
func MustNewVaultClientWithContext(ctx context.Context, role string, url string) *VaultClient {
	credsHelper := newVaultCredsHelper(role)

	// Register a health checker up front so the health endpoint reports vault
	// as unhealthy ("Connecting to vault") while the retry loop runs, instead of
	// the dependency being invisible until it finally connects.
	checker := &vaultStartupChecker{}
	rorhealth.Register(ctx, "vault", checker)

	backoff := vaultInitialBackoff
	for attempt := 1; ; attempt++ {
		client, err := New(ctx, credsHelper, url)
		if err == nil {
			if attempt > 1 {
				rlog.Info("connected to vault",
					rlog.String("url", url),
					rlog.Int("attempts", attempt))
			}
			checker.setClient(client)
			return client
		}

		rlog.Error("could not connect to vault, retrying", err,
			rlog.String("url", url),
			rlog.Int("attempt", attempt),
			rlog.String("retryIn", backoff.String()))

		select {
		case <-ctx.Done():
			rlog.Fatal("could not connect to vault within timeout, giving up", ctx.Err(),
				rlog.String("url", url),
				rlog.Int("attempts", attempt))
		case <-time.After(backoff):
		}

		backoff *= 2
		if backoff > vaultMaxBackoff {
			backoff = vaultMaxBackoff
		}
	}
}

// newVaultCredsHelper selects the credentials helper based on the runtime
// environment: a Kubernetes service account token when running in-cluster,
// otherwise the VAULT_TOKEN env var (falling back to a development token).
func newVaultCredsHelper(role string) VaultCredsHelper {
	tokenFilePath := "/var/run/secrets/kubernetes.io/serviceaccount/token" // #nosec G101 Jest the path to the token file in the secrets engine
	if _, err := os.Stat(tokenFilePath); err == nil {
		return NewKubernetesVaultCredsHelper(role, 3600)
	}

	envtoken := rorconfig.GetString("VAULT_TOKEN")
	if envtoken == "" {
		return NewStaticVaultCredsHelper("S3cret!")
	}
	return NewStaticVaultCredsHelper(envtoken)
}
