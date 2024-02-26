package vaultclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"errors"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/dotse/go-health"
	"github.com/hashicorp/vault-client-go/schema"
	"github.com/hashicorp/vault/api"

	"github.com/hashicorp/vault-client-go"
)

var vaultClient *VaultClient = &VaultClient{}

type VaultClient struct {
	Context        context.Context
	Client         *vault.Client
	Url            string
	Role           string
	Token          string
	Exp            int64
	Ttl            int32
	RenewThreshold int64
}

func Init(role string, url string) {
	vaultClient.Context = context.Background()
	vaultClient.Role = role
	vaultClient.Exp = 0
	vaultClient.Ttl = 86400
	vaultClient.RenewThreshold = 3600
	vaultClient.Url = url
	vaultClient.initClient()
	vaultClient.getInitialToken()
	health.Register("vault", vaultClient)
}

// creates a new vault client
func New(ctx context.Context, role string, url string, renewThreshold int64) (*VaultClient, error) {
	client, err := getClient(url)
	if err != nil {
		return nil, err
	}

	vaultClient := VaultClient{
		Context:        ctx,
		Client:         client,
		Role:           role,
		RenewThreshold: renewThreshold,
	}

	return &vaultClient, nil
}

// logs into a vault using a k8s authentication path and gets a new token
func (v *VaultClient) K8sLogin(ctx context.Context) error {
	tokenFilePath := "/var/run/secrets/kubernetes.io/serviceaccount/token"

	if _, err := os.Stat(tokenFilePath); err == nil {
		rlog.Infoc(ctx, "loging into vault using service account token")

		byteValue, _ := os.ReadFile(tokenFilePath)
		jwt := string(byteValue)
		resp, err := v.Client.Auth.KubernetesLogin(v.Context, schema.KubernetesLoginRequest{Jwt: jwt, Role: v.Role})
		if err != nil {
			return fmt.Errorf("could not authenticate against vault: %w", err)
		}

		v.Token = resp.Auth.ClientToken
		v.Exp = time.Now().Unix() + int64(resp.Auth.LeaseDuration)
		v.Ttl = int32(resp.Auth.LeaseDuration)

		rlog.Infoc(ctx, "authenticated to vault", rlog.Int("ttl", resp.Auth.LeaseDuration))

	} else {
		rlog.Warnc(ctx, "authenticating against Vault with a static token. This is not recomended in production!!!!")
		// TODO: Check if development or get static token from env.
		v.Token = "S3cret!"
		v.Exp = time.Now().Unix() + int64(365*24*3600)
	}

	err := v.Client.SetToken(v.Token)
	if err != nil {
		return fmt.Errorf("could not set token: %w", err)
	}

	return nil
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

// Starts a goroutine, that will set a timer for the next renew threshold and
// renew the token, it will run until it recieves a done signal. It does not
// handle errors well at the moment
func (v VaultClient) WaitForTokenRenewal(ctx context.Context, doneStream chan interface{}, errorStream chan error) {
	rlog.Debugc(ctx, "started vault token refresher")

loop:
	for {
		timer := time.NewTimer(time.Second * time.Duration(v.Ttl-int32(v.RenewThreshold)))

		select {
		case <-doneStream:
			break loop
		case <-timer.C:
			rlog.Debugc(ctx, "attempting to renew vault acces token")
			err := v.renewToken(ctx)
			if err != nil {
				//rlog.Errorc(ctx, "failed to renew token", err)
				errorStream <- err
			}
			rlog.Debugc(ctx, "token renewed")
		}
	}
}

func (rc VaultClient) CheckHealth() []health.Check {
	c := health.Check{
		Status: health.StatusPass,
		Output: "Vault connection ok",
	}
	ok, err := rc.Ping()
	if !ok {
		rlog.Error("vault ping returned error", err)
		c.Status = health.StatusFail
		c.Output = "Could not connect to vault"
	}

	return []health.Check{c}
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

func (v *VaultClient) renewToken(ctx context.Context) error {
	resp, err := v.Client.Auth.TokenRenewSelf(v.Context, schema.TokenRenewSelfRequest{})
	if err != nil {
		return fmt.Errorf("could not renew token: %w", err)
	}

	v.Token = resp.Auth.ClientToken
	v.Exp = time.Now().Unix() + int64(resp.Auth.LeaseDuration)
	v.Ttl = int32(resp.Auth.LeaseDuration)

	err = v.Client.SetToken(v.Token)
	if err != nil {
		return fmt.Errorf("could not set token: %w", err)
	}

	return nil
}

func (v *VaultClient) isExpired() bool {
	return v.Exp-v.RenewThreshold < time.Now().Unix()
}

func getVaultClient() (*api.Client, error) {
	client, err := api.NewClient(&api.Config{Address: vaultClient.Url, HttpClient: httpClient})
	if err != nil {
		msg := "could not get secret, problems with vault client"
		rlog.Error(msg, err)
		return nil, errors.New(msg)
	}

	vaultClient.renewTokenIfNeeded()
	client.AddHeader("X-Vault-Token", vaultClient.Token)
	return client, nil
}

// DEPRECATED. we dont want to have to rely on calling this withing the lease
// time of the tokens to be able to renew them. We use WaitForTokenRenewal go
// routine instead
func (v VaultClient) GetClient() *vault.Client {
	v.renewTokenIfNeeded()
	return v.Client
}

// DEPRECATED. Use New() instead
func NewVaultClient(role string, url string) *VaultClient {
	vc := VaultClient{}
	vc.Context = context.Background()
	vc.Role = role
	vc.Exp = 0
	vc.Ttl = 86400
	vc.RenewThreshold = 10
	vc.Url = url
	vc.initClient()
	vc.getInitialToken()
	return &vc
}

// DEPRECATED. Remove this function, it is only used to support legacy code
func GetInitiatedVaultClient() *VaultClient {
	return vaultClient
}

// DEPRECATED. use getClient instead
func (v *VaultClient) initClient() {
	var err error
	v.Client, err = vault.New(
		vault.WithAddress(v.Url),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
}

// DEPRECATED. use WaitForTokenRenewal instead
func (v VaultClient) renewTokenIfNeeded() {
	if v.isExpired() {
		ctx := context.Background()
		err := v.renewToken(ctx)
		if err != nil {
			rlog.Info("could not renew token", rlog.Any("error", err))
		}
	}
}

// DEPRECATED. use K8sLogin instead
func (v *VaultClient) getInitialToken() {
	tokenFilePath := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	if _, err := os.Stat(tokenFilePath); err == nil {
		rlog.Info("Staring vault login using Kubernetes logon method")

		byteValue, _ := os.ReadFile(tokenFilePath)
		jwt := string(byteValue)
		resp, err := v.Client.Auth.KubernetesLogin(v.Context, schema.KubernetesLoginRequest{Jwt: jwt, Role: v.Role})
		if err != nil {
			rlog.Fatal("Could not authenticate against vault", err)
		}

		v.Token = resp.Auth.ClientToken
		v.Exp = time.Now().Unix() + int64(resp.Auth.LeaseDuration)
		v.Ttl = int32(resp.Auth.LeaseDuration)

		rlog.Infof("Authenticated to vault, ttl: %v", resp.Auth.LeaseDuration)
	} else {
		rlog.Warn("Authenticating against Vault with a static token. This is not recomended in production!!!!")
		// TODO: Check if development or get static token from env.
		v.Token = "S3cret!"
		v.Exp = time.Now().Unix() + int64(365*24*3600)
	}
	err := v.Client.SetToken(v.Token)
	if err != nil {
		log.Fatal(err)
	}
}
