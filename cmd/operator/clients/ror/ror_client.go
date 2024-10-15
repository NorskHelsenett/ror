package ror

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/operator/clients"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/spf13/viper"
)

func FetchConfiguration(ctx context.Context) ([]byte, error) {
	rorClient, err := clients.GetOrCreateRorClient()
	if err != nil {
		return nil, errors.New("could not fetch ror client")
	}

	url := fmt.Sprintf("%s/v1/configs/operator", viper.GetString(configconsts.API_ENDPOINT))
	response, err := rorClient.R().
		SetContext(ctx).
		Get(url)
	if err != nil {
		return nil, errors.New("could not get configuration from ror-api")
	}

	return response.Body(), nil
}

func FetchTaskConfiguration(ctx context.Context, operatorTask apicontracts.OperatorTask) (*apicontracts.OperatorJob, error) {
	rorClient, err := clients.GetOrCreateRorClient()
	if err != nil {
		return nil, errors.New("could not fetch ror client")
	}

	url := fmt.Sprintf("%s/v1/clusters/%s/configs/%s", viper.GetString(configconsts.API_ENDPOINT), viper.GetString(configconsts.CLUSTER_ID), operatorTask.Name)
	response, err := rorClient.R().
		SetContext(ctx).
		Get(url)
	if err != nil {
		return nil, errors.New("could not get configuration from ror-api")
	}

	statusCode := response.StatusCode()
	if statusCode > 299 {
		messages, err := stringhelper.JsonToMap(response.String())
		if err != nil {
			return nil, fmt.Errorf("could not get configuration from ror-api, status code: %d, error: %s", statusCode, response.String())
		}

		msg := messages["message"]
		return nil, fmt.Errorf("could not get configuration from ror-api, status code: %d, error: %s", statusCode, msg)
	}

	var operatorJob apicontracts.OperatorJob
	err = json.Unmarshal(response.Body(), &operatorJob)
	if err != nil {
		return nil, errors.New("could not marshal task config from api")
	}

	return &operatorJob, nil
}
