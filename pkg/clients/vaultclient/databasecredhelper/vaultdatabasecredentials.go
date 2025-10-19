package databasecredhelper

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/helpers/credshelper"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/hashicorp/vault-client-go"
)

type VaultDBCredentials struct {
	lock        sync.RWMutex
	VaultClient *vaultclient.VaultClient
	Username    string
	Password    string
	VaultRole   string
	MountPath   string
	Exp         int64
}

// NewVaultDBCredentials creates a new VaultDBCredentials object
// vcli is the vault client
// vaultRole is the role to request credentials for
// mountpath is the path to the database mount in vault
func NewVaultDBCredentials(vcli *vaultclient.VaultClient, vaultRole string, mountpath string) *credshelper.SimpleWrapperWithRenew {

	vc := VaultDBCredentials{
		VaultClient: vcli,
		VaultRole:   vaultRole,
		Exp:         0,
	}
	if mountpath != "" {
		vc.MountPath = mountpath
	} else {
		vc.MountPath = "database"
	}
	return credshelper.WrapSimpleCredsHelperWithRenew(&vc)
}
func (dbc *VaultDBCredentials) isExpired() bool {
	dbc.lock.RLock()
	defer dbc.lock.RUnlock()

	return dbc.Exp-10 < time.Now().Unix()
}

func (dbc *VaultDBCredentials) CheckAndRenew() bool {
	if dbc.isExpired() {
		dbc.lock.Lock()
		defer dbc.lock.Unlock()

		msg := fmt.Sprintf("Renewing lease for database %s/credentials/%s", dbc.MountPath, dbc.VaultRole)
		rlog.Debug(msg)
		_ = dbc.updateCreds()
		return true
	}
	return false
}

func (dbc *VaultDBCredentials) GetCredentials() (string, string) {
	dbc.CheckAndRenew()
	return dbc.Username, dbc.Password
}
func (dbc *VaultDBCredentials) updateCreds() error {

	if dbc.VaultRole == "" {
		return errors.New("secret path is nil or empty")
	}

	data, err := dbc.VaultClient.Client.Secrets.DatabaseGenerateCredentials(dbc.VaultClient.Context, dbc.VaultRole, vault.WithMountPath(dbc.MountPath))
	if err != nil {
		msg := "could not get secret, are you sure that the token is valid? "
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	if data != nil {
		dbc.Username = data.Data["username"].(string)
		dbc.Password = data.Data["password"].(string)
		dbc.Exp = time.Now().Unix() + int64(data.LeaseDuration)
		msg := fmt.Sprintf("new user: %s for path %s", dbc.Username, dbc.VaultRole)
		rlog.Debug(msg)
		return nil
	}
	return errors.New("no credential found")
}
