package vaultclient

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type StaticVaultCredsHelper struct {
	token string
}

func NewStaticVaultCredsHelper(token string) VaultCredsHelper {
	return &StaticVaultCredsHelper{
		token: token,
	}
}

func (sch StaticVaultCredsHelper) GetToken() string {
	return sch.token
}

func (sch *StaticVaultCredsHelper) Login(vc *VaultClient) error {
	ctx := vc.Context
	rlog.Warnc(ctx, "authenticating against Vault with a static token. This is not recomended in production!!!!")
	// TODO: Check if development or get static token from env.
	//sch.token = "S3cret!"

	err := vc.Client.SetToken(sch.token)
	if err != nil {
		return fmt.Errorf("could not set token: %w", err)
	}
	return nil
}
