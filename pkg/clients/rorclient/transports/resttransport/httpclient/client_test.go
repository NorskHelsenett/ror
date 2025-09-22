package httpclient_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
)

// testNewHttpTransportClientConfig tests that the constructor functions and
// it passes valid addresses and rejects any invalid addresses.
func TestNewHttpTransportClientConfig(t *testing.T) {

	authProvider := httpauthprovider.NewNoAuthprovider()
	role := "test"
	version := rorversion.GetRorVersion()
	urls := map[string]bool{
		"":                        false,
		"127.0.0.1":               false,
		"127.0.0.1:10000":         false,
		"http://localhost":        true,
		"http://localhost:10000":  true,
		"https://localhost":       true,
		"https://localhost:10000": true,
		"asfjniefsjfsdjdsf":       false,
	}

	for url, expectedAction := range urls {
		config, err := httpclient.NewHttpTransportClientConfig(url, authProvider, role, version)
		errStatus := err != nil
		if errStatus != expectedAction {
			t.Errorf("failed on validation on url %v, got %v, expected %v", url, errStatus, expectedAction)
		}

		if config.GetRole() != role {
			t.Errorf("failed to get expected role %v, got %v", role, config.GetRole())
		}

		if config.Version != version {
			t.Errorf("failed to get expected version  %v, got %v", config.Version, version)
		}
	}
}
