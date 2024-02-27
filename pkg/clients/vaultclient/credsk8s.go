package vaultclient

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/hashicorp/vault-client-go/schema"
)

type KubernetesVaultCredsHelper struct {
	client         *VaultClient
	renewThreshold int32
	token          string
	role           string
	ttl            int32
}

func NewKubernetesVaultCredsHelper(vaultRole string, tokenRenewThreshold int32) VaultCredsHelper {
	return &KubernetesVaultCredsHelper{
		role:           vaultRole,
		renewThreshold: tokenRenewThreshold,
	}
}

func (kvch KubernetesVaultCredsHelper) GetToken() string {
	return kvch.token
}

func (kvch *KubernetesVaultCredsHelper) Login(vc *VaultClient) error {
	tokenFilePath := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	kvch.client = vc
	ctx := vc.Context

	if _, err := os.Stat(tokenFilePath); err != nil {
		return err
	}

	rlog.Infoc(ctx, "loging into vault using service account token")

	byteValue, _ := os.ReadFile(tokenFilePath)
	jwt := string(byteValue)
	resp, err := vc.Client.Auth.KubernetesLogin(ctx, schema.KubernetesLoginRequest{Jwt: jwt, Role: kvch.role})
	if err != nil {
		return fmt.Errorf("could not authenticate against vault: %w", err)
	}

	kvch.token = resp.Auth.ClientToken
	kvch.ttl = int32(resp.Auth.LeaseDuration)

	rlog.Infoc(ctx, "authenticated to vault", rlog.Int("ttl", resp.Auth.LeaseDuration))

	err = kvch.client.Client.SetToken(kvch.token)
	if err != nil {
		return fmt.Errorf("could not set token: %w", err)
	}

	go kvch.waitForTokenRenewal(ctx)
	return nil
}

func (kvch *KubernetesVaultCredsHelper) waitForTokenRenewal(ctx context.Context) {
	rlog.Debugc(ctx, "started vault token refresher")

	for {
		timer := time.NewTimer(time.Second * time.Duration(kvch.ttl-int32(kvch.renewThreshold)))

		<-timer.C
		rlog.Debugc(ctx, "attempting to renew vault acces token")
		err := kvch.renewToken(ctx)
		if err != nil {
			rlog.Errorc(ctx, "failed to renew token", err)
		}
		rlog.Debugc(ctx, "token renewed")

	}
}

func (kvch *KubernetesVaultCredsHelper) renewToken(ctx context.Context) error {
	resp, err := kvch.client.Client.Auth.TokenRenewSelf(ctx, schema.TokenRenewSelfRequest{})
	if err != nil {
		return fmt.Errorf("could not renew token: %w", err)
	}

	kvch.token = resp.Auth.ClientToken
	kvch.ttl = int32(resp.Auth.LeaseDuration)

	err = kvch.client.Client.SetToken(kvch.token)
	if err != nil {
		return fmt.Errorf("could not set token: %w", err)
	}

	return nil
}
