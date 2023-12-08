package vaultclient

import (
	"errors"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type VaultCredentials struct {
	Username  string
	Password  string
	VaultPath string
	Exp       int64
}

func (vc *VaultCredentials) Init(vaultpath string) {
	vc.VaultPath = vaultpath
	vc.Exp = 0
}

func (vc *VaultCredentials) IsExpired() bool {
	return vc.Exp-10 < time.Now().Unix()
}

func (vc *VaultCredentials) CheckAndRenew() bool {
	if vc.IsExpired() {
		msg := fmt.Sprintf("Renewing lease for %s", vc.VaultPath)
		rlog.Debug(msg)
		_ = vc.updateCreds()
		return true
	}
	return false
}

func (vc *VaultCredentials) GetUsernamePassword() (string, string) {
	vc.CheckAndRenew()
	return vc.Username, vc.Password
}

func (vc *VaultCredentials) updateCreds() error {

	if vc.VaultPath == "" {
		return errors.New("secret path is nil or empty")
	}

	vaultClient, err := getVaultClient()
	if err != nil {
		msg := "could not get secret, problems with vault client: "
		rlog.Error(msg, err)
		return errors.New(msg)
	}

	data, err := vaultClient.Logical().Read(vc.VaultPath)
	if err != nil {
		msg := "could not get secret, are you sure that the token is valid? "
		rlog.Error(msg, err)
		return errors.New(msg)
	}
	if data != nil {
		vc.Username = data.Data["username"].(string)
		vc.Password = data.Data["password"].(string)
		vc.Exp = time.Now().Unix() + int64(data.LeaseDuration)
		msg := fmt.Sprintf("New user: %s for path %s", vc.Username, vc.VaultPath)
		rlog.Debug(msg)
		return nil
	}
	return errors.New("no credential found")
}
