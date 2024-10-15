package rorclient

import (
	"encoding/json"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/clients/rorapiclient"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
)

func GetWorkspaces() ([]apicontracts.Workspace, error) {
	rorClient, err := rorapiclient.GetOrCreateRorRestyClient()
	if err != nil {
		return nil, err
	}

	requestUrl := "v1/workspaces"
	response, err := rorClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-API-KEY", viper.GetString(configconsts.API_KEY)).
		Get(requestUrl)
	if err != nil {
		rlog.Error("Could not send data to ror-api", err)
		return nil, err
	}

	if response.StatusCode() > 299 {
		rlog.Fatal("Error calling ror api", fmt.Errorf("non 200 errorcode"),
			rlog.String("request", fmt.Sprintf("%s/%s", viper.GetString(configconsts.API_ENDPOINT), requestUrl)),
			rlog.Int("code", response.StatusCode()))
	}

	var workspaces []apicontracts.Workspace
	err = json.Unmarshal(response.Body(), &workspaces)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}
