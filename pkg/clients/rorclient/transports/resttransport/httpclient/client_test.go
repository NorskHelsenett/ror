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
		"::f1":                    false,
	}

	for url, expectedAction := range urls {
		_, err := httpclient.NewHttpTransportClientConfig(url, authProvider, role, version)
		hadError := err == nil
		if hadError != expectedAction {
			t.Errorf("failed on validation on url '%v', got %v, expected %v", url, hadError, expectedAction)
		} else {
			t.Logf("passed url validation on '%v', got %v, expected %v", url, hadError, expectedAction)
		}

	}
}
