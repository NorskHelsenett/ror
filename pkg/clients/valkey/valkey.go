package valkey

import (
	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient/databasecredhelper"
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

	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: vc.InitAddress, AuthCredentialsFn: func(valkey.AuthCredentialsContext) (valkey.AuthCredentials, error) {
		return valkey.AuthCredentials{Username: vc.Credentials.Username, Password: vc.Credentials.Password}, nil
	}})
	if err != nil {
		return err
	}

	vc.Client = client

	//valkey.
	return nil
}
