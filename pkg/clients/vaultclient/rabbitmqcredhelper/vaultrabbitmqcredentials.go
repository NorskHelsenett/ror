package rabbitmqcredhelper

import (
	"errors"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/hashicorp/vault-client-go"
)

type VaultRMQCredentials struct {
	VaultClient *vaultclient.VaultClient
	Username    string
	Password    string
	VaultPath   string
	Exp         int64
}

// NewVaultRMQCredentials creates a new VaultRMQCredentialhelper
// If vaultpath is empty, the role will be used as path
func NewVaultRMQCredentials(vcli *vaultclient.VaultClient, vaultpath string) *VaultRMQCredentials {
	if vaultpath == "" {
		vaultpath = vcli.Role
	}
	vc := VaultRMQCredentials{
		VaultClient: vcli,
		VaultPath:   vaultpath,
		Exp:         0,
	}
	return &vc
}

func (rmc *VaultRMQCredentials) GetCredentials() (string, string) {
	rmc.checkAndRenew()
	return rmc.Username, rmc.Password
}

func (rmc *VaultRMQCredentials) isExpired() bool {
	return rmc.Exp-10 < time.Now().Unix()
}

func (rmc *VaultRMQCredentials) checkAndRenew() bool {
	if rmc.isExpired() {
		msg := fmt.Sprintf("Renewing lease for RabbitMQ rabbitmq/credentials/%s", rmc.VaultPath)
		rlog.Debug(msg)
		_ = rmc.updateCreds()
		return true
	}
	return false
}

func (rmc *VaultRMQCredentials) updateCreds() error {
	if rmc.VaultPath == "" {
		return errors.New("secret path is nil or empty")
	}

	data, err := rmc.VaultClient.Client.Secrets.RabbitMqRequestCredentials(rmc.VaultClient.Context, rmc.VaultPath, vault.WithMountPath("rabbitmq"))
	if err != nil {
		msg := "could not get secret, are you sure that the token is valid? "
		rlog.Error(msg, err)
		return errors.New(msg)
	}

	if data != nil {
		rmc.Username = data.Data["username"].(string)
		rmc.Password = data.Data["password"].(string)
		rmc.Exp = time.Now().Local().Unix() + int64(data.LeaseDuration)
		msg := fmt.Sprintf("New user: %s for path %s", rmc.Username, rmc.VaultPath)
		rlog.Debug(msg)
		return nil
	}

	return errors.New("no credential found")
}
