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

func NewVaultClient(role string, url string) *VaultClient {
	vc := VaultClient{}
	vc.Context = context.Background()
	vc.Role = role
	vc.Exp = 0
	vc.Ttl = 86400
	vc.RenewThreshold = 3600
	vc.Url = url
	vc.initClient()
	vc.getInitialToken()
	return &vc
}
func (rc VaultClient) CheckHealth() []health.Check {
	c := health.Check{}
	if !rc.Ping() {
		c.Status = health.StatusFail
		c.Output = "Could not onnect to vault"
	}
	return []health.Check{c}
}

// TODO: Remove this function, it is only used to support legacy code
func GetInitiatedVaultClient() *VaultClient {
	return vaultClient
}

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

func (v VaultClient) GetClient() *vault.Client {
	v.renewTokenIfNeeded()
	return v.Client
}

func (v VaultClient) renewTokenIfNeeded() {
	if v.isExpired() {
		v.renewToken()
	}
}

func (v *VaultClient) renewToken() {
	rlog.Infof("Renewing vault token %s for %v", v.Role, v.Ttl)
	resp, err := v.Client.Auth.TokenRenew(v.Context, schema.TokenRenewRequest{Increment: fmt.Sprintf("%v", v.Ttl), Token: v.Token})
	if err != nil {
		rlog.Fatal("Could not renew vault token", err)
	}

	v.Token = resp.Auth.ClientToken
	v.Exp = time.Now().Unix() + int64(resp.Auth.LeaseDuration)

	err = v.Client.SetToken(v.Token)
	if err != nil {
		rlog.Error("Could not set token", err)
	}
}

func (v VaultClient) Ping() bool {
	if v.Client == nil {
		rlog.Error("Vault client is not initialized", fmt.Errorf(""))
		return false
	}
	_, err := v.Client.Auth.TokenLookUpSelf(v.Context)

	return err == nil
}

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
