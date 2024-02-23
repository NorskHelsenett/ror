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
	vc.RenewThreshold = 10
	vc.Url = url
	vc.initClient()
	vc.getInitialToken()
	return &vc
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

// Starts a goroutine, that will set a timer for the next renew threshold and
// renew the token
func (v VaultClient) WaitForTokenRenewal(ctx context.Context, done chan interface{}) {
	rlog.Debugc(ctx, "started token refresher")
	//debu
	err := v.renewToken()
	if err != nil {
		rlog.Infoc(ctx, "failed to renew token", rlog.Any("error", err))
	}
	//\debug
	for {
		timer := time.NewTimer(time.Second * time.Duration(v.Ttl-int32(v.RenewThreshold)))
		rlog.Debugc(ctx, "new timer created", rlog.Any("seconds", v.Ttl-int32(v.RenewThreshold)))

		select {
		case <-done:
			break
		case <-timer.C:
			rlog.Debugc(ctx, "time to renew token")
			err := v.renewToken()
			if err != nil {
				rlog.Infoc(ctx, "failed to renew token", rlog.Any("error", err))
			}
		}
	}
}

func (v VaultClient) renewTokenIfNeeded() {
	if v.isExpired() {
		err := v.renewToken()
		if err != nil {
			rlog.Info("could not renew token", rlog.Any("error", err))
		}
	}
}

func (v *VaultClient) renewToken() error {
	rlog.Infof("Renewing vault token %s for %v", v.Role, v.Ttl)
	rlog.Info("token values", rlog.Any("exp", v.Exp), rlog.Any("now", time.Now().Unix()), rlog.Any("ttl", v.Ttl), rlog.Any("renewThreshold", v.RenewThreshold))

	resp, err := v.Client.Auth.TokenRenew(v.Context, schema.TokenRenewRequest{Token: v.Token})
	if err != nil {
		return fmt.Errorf("Could not renew token: %w", err)
	}

	v.Token = resp.Auth.ClientToken
	v.Exp = time.Now().Unix() + int64(resp.Auth.LeaseDuration)
	v.Ttl = int32(resp.Auth.LeaseDuration)

	err = v.Client.SetToken(v.Token)
	if err != nil {
		return fmt.Errorf("Could not set token: %w", err)
	}

	return nil
}

func (v VaultClient) Ping() (bool, error) {
	if v.Client == nil {
		err := fmt.Errorf("Vault client is not initialized")
		rlog.Error("could not ping vault", err)
		return false, err
	}
	_, err := v.Client.Auth.TokenLookUpSelf(v.Context)
	if err != nil {
		return false, err
	}

	return true, nil
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

		rlog.Info("checking if token is even renewable", rlog.Any("is token renewable", resp.Renewable), rlog.Any("leak the token for good meassure", resp.Auth.ClientToken))
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
