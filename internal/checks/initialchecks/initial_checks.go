package initialchecks

import (
	"fmt"
	"github.com/NorskHelsenett/ror/internal/clients/rorapiclient"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

// HasSuccessfullRorApiConnection checks whether there is a successful connection to the ROR API.
//
// - If the ROR_URL is not set, it returns an error.
// Returns:
// - (bool): True if the connection to the ROR API is successful, false otherwise.
func HasSuccessfullRorApiConnection() error {
	if viper.GetString(configconsts.API_ENDPOINT) == "" {
		return fmt.Errorf("ROR_URL is not set")
	}

	rorClient, err := rorapiclient.GetOrCreateRorRestyClientNonAuth()
	if err != nil {
		rlog.Error("could not get ror-api client", err)
		return err
	}

	url := "/v1/info/version"
	response, err := rorClient.R().
		SetHeader("Content-Type", "application/json").
		Get(url)
	if err != nil {
		rlog.Error("could not get data from ror-api", err)
		return err
	}

	if response.StatusCode() > 299 {
		return fmt.Errorf("unsuccessful connection to ror-api: status code %d", response.StatusCode())
	}

	return nil
}
