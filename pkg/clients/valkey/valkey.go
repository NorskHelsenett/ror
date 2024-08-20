package valkey

import (
	"context"
	"errors"
	"time"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/databasecredhelper"
	"github.com/dotse/go-health"
	"github.com/valkey-io/valkey-go"
)

// This type implements the redis connection in ror

type valkeycon struct {
	Client      valkey.Client
	Credentials *databasecredhelper.VaultDBCredentials
	InitAddress []string
	Port        string
}

func New(dbc *databasecredhelper.VaultDBCredentials, address ...string) (*valkeycon, error) {
	vc := valkeycon{
		Credentials: dbc,
		InitAddress: address,
	}
	err := vc.connect()
	if err != nil {
		return nil, err
	}

	return &vc, nil
}
func (vc *valkeycon) GetClient() valkey.Client {
	return vc.Client
}
func (vc *valkeycon) connect() error {

	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: vc.InitAddress, DisableCache: true, AuthCredentialsFn: func(valkey.AuthCredentialsContext) (valkey.AuthCredentials, error) {
		ok := vc.Credentials.CheckAndRenew()
		if !ok {
			return valkey.AuthCredentials{}, errors.New("could not renew credentials")
		}
		return valkey.AuthCredentials{Username: vc.Credentials.Username, Password: vc.Credentials.Password}, nil
	}})
	if err != nil {
		return err
	}

	vc.Client = client

	//valkey.
	return nil
}

// CheckHealth checks the health of the redis connection and returns a health check
func (vc valkeycon) CheckHealth() []health.Check {
	c := health.Check{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ok, err := vc.Client.Do(ctx, vc.Client.B().Ping().Build()).AsBool()
	if err != nil {
		c.Status = health.StatusFail
		c.Output = err.Error()
		return []health.Check{c}
	}
	if !ok {
		c.Status = health.StatusFail
		c.Output = "Could not ping redis"
		return []health.Check{c}
	}
	return []health.Check{c}
}
