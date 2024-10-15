package helsegitlabclient

import (
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/audit/msauditconnections"
	"github.com/NorskHelsenett/ror/cmd/audit/msauditmodels"
	"math"
	"strconv"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var client *resty.Client

// Pushes markdown to main branch (https://helsegitlab.nhn.no/sdi/SDI-Infrastruktur/authorizationsregister)
func PushAclToRepo(markdownContent string) error {
	gitlabToken, projectNumber, branch, err := getTokenAndProjectNumberFromVault()
	if err != nil {
		return err
	}

	helseGitlabBaseUrl := viper.GetString(configconsts.HELSEGITLAB_BASE_URL)
	url := fmt.Sprintf("%s%d/repository/commits", helseGitlabBaseUrl, projectNumber)

	if client == nil {
		client = resty.New()
	}

	commitPostMessage := msauditmodels.CommitPostMessage{
		Branch:        branch,
		CommitMessage: "ms audit - acl update",
		Actions: []msauditmodels.CommitAction{
			{
				Action:   "update",
				FilePath: "README.md",
				Content:  markdownContent,
			},
		},
	}
	response, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Encoding", "gzip, deflate, br").
		SetHeader("PRIVATE-TOKEN", gitlabToken).
		SetBody(commitPostMessage).
		Post(url)
	if err != nil {
		rlog.Error("could not get a response from helsegitlab client", err)
		return errors.New("could not get a response from helsegitlab client")
	}

	statusCode := response.StatusCode()
	if statusCode > 299 {
		messages, err := stringhelper.JsonToMap(response.String())
		if err != nil {
			return fmt.Errorf("could not get configuration from helsegitlab, status code: %d, error: %s", statusCode, response.String())
		}

		return fmt.Errorf("could not get configuration from helsegitlab, status code: %d, error: %s", statusCode, messages)
	}

	return nil
}

func getTokenAndProjectNumberFromVault() (string, int, string, error) {
	secretPath := "secret/data/v1.0/ror/config/ms-audit" // #nosec G101 Jest the path to the token file in the secrets engine
	vaultData, err := msauditconnections.VaultClient.GetSecret(secretPath)
	if err != nil {
		return "", math.MinInt, "", errors.New("could not extract gitlab access token from vault")
	}

	commonConfig, ok := vaultData["data"].(map[string]interface{})
	if !ok {
		rlog.Error("", fmt.Errorf("data type assertion failed: %T %#v", vaultData["data"], vaultData["data"]))
		return "", math.MinInt, "", errors.New("could not extract gitlab access token from vault.data")
	}

	token := commonConfig["helsegitlabAclToken"].(string)
	projectNumberString := commonConfig["aclRepoNumber"].(string)
	branch := commonConfig["branch"].(string)

	projectNumber, err := strconv.Atoi(projectNumberString)
	if err != nil {
		return "", math.MinInt, "", errors.New("could not convert projectnumber string to int")
	}

	return token, projectNumber, branch, nil
}
