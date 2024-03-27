package databasecredhelper

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/hashicorp/vault-client-go"
)

type VaultDBCredentials struct {
	lock        sync.RWMutex
	VaultClient *vaultclient.VaultClient
	Username    string
	Password    string
	VaultPath   string
	MountPath   string
	Exp         int64
}

func NewVaultDBCredentials(vcli *vaultclient.VaultClient, vaultpath string, mountpath string) *VaultDBCredentials {

	vc := VaultDBCredentials{
		VaultClient: vcli,
		VaultPath:   vaultpath,
		Exp:         0,
	}
	if mountpath != "" {
		vc.MountPath = mountpath
	} else {
		vc.MountPath = "database"
	}

	return &vc
}
func (dbc *VaultDBCredentials) IsExpired() bool {
	dbc.lock.RLock()
	defer dbc.lock.RUnlock()

	return dbc.Exp-10 < time.Now().Unix()
}

func (dbc *VaultDBCredentials) CheckAndRenew() bool {
	if dbc.IsExpired() {
		dbc.lock.Lock()
		defer dbc.lock.Unlock()

		msg := fmt.Sprintf("Renewing lease for database %s/credentials/%s", dbc.MountPath, dbc.VaultPath)
		rlog.Debug(msg)
		_ = dbc.updateCreds()
		return true
	}
	return false
}

// Deprecated: Use GetCredentials instead
func (dbc *VaultDBCredentials) GetUsernamePassword() (string, string) {
	return dbc.GetCredentials()
}

func (dbc *VaultDBCredentials) GetCredentials() (string, string) {
	dbc.CheckAndRenew()
	return dbc.Username, dbc.Password
}
func (dbc *VaultDBCredentials) updateCreds() error {

	if dbc.VaultPath == "" {
		return errors.New("secret path is nil or empty")
	}

	data, err := dbc.VaultClient.Client.Secrets.DatabaseGenerateCredentials(dbc.VaultClient.Context, dbc.VaultPath, vault.WithMountPath(dbc.MountPath))
	if err != nil {
		msg := "could not get secret, are you sure that the token is valid? "
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	if data != nil {
		dbc.Username = data.Data["username"].(string)
		dbc.Password = data.Data["password"].(string)
		dbc.Exp = time.Now().Unix() + int64(data.LeaseDuration)
		msg := fmt.Sprintf("new user: %s for path %s", dbc.Username, dbc.VaultPath)
		rlog.Debug(msg)
		return nil
	}
	return errors.New("no credential found")
}
