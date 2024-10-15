// TODO: Replace with go-gitlab compatible library
package helsegitlab

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var client *resty.Client
var request *resty.Request

func GetFileContent(projectId int, filePath string, branch string, vaultClient *vaultclient.VaultClient) ([]byte, error) {
	gitlabRequest, err := getGitlabRequest(vaultClient)
	if err != nil {
		rlog.Error("could not get gitlab client", err)
		return nil, errors.New("could not get gitlab client")
	}

	urlencodeFilePath := url.QueryEscape(filePath)
	urlencodeBranch := url.QueryEscape(branch)

	repourl := fmt.Sprintf("%s%d/repository/files/%s/raw?ref=%s", viper.GetString(configconsts.HELSEGITLAB_BASE_URL), projectId, urlencodeFilePath, urlencodeBranch)
	response, err := gitlabRequest.Get(repourl)
	if err != nil {
		return nil, errors.New("could not get a response from helsegitlab client")
	}

	statusCode := response.StatusCode()
	if statusCode > 299 {
		messages, err := stringhelper.JsonToMap(response.String())
		if err != nil {
			return nil, fmt.Errorf("could not get configuration from ror-api, status code: %d, error: %s", statusCode, response.String())
		}

		return nil, fmt.Errorf("could not get configuration from ror-api, status code: %d, error: %s", statusCode, messages)
	}

	return response.Body(), nil
}

func getGitlabRequest(vaultClient *vaultclient.VaultClient) (*resty.Request, error) {
	if request != nil {
		return request, nil
	}

	secretPath := "secret/data/v1.0/ror/config/common" // #nosec G101 Jest the path to the token file in the secrets engine
	vaultData, err := vaultClient.GetSecret(secretPath)
	if err != nil {
		return nil, errors.New("could not extract gitlab access token from vault")
	}

	commonConfig, ok := vaultData["data"].(map[string]interface{})
	if !ok {
		rlog.Error("", fmt.Errorf("data type assertion failed: %T %#v", vaultData["data"], vaultData["data"]))
		return nil, errors.New("could not extract gitlab access token from vault-data")
	}

	token := commonConfig["helsegitlabToken"].(string)

	if len(token) < 1 {
		rlog.Error("could not get gitlab access token from vault", fmt.Errorf("token is empty"))
		return nil, errors.New("could not get gitlab access token from vault")
	}

	client = resty.New()
	req := client.R()
	req.SetHeader("PRIVATE-TOKEN", token)

	return req, nil
}
