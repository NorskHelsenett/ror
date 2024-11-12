package resourceupdatev2

import (
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/agent/clients"
	"github.com/NorskHelsenett/ror/cmd/agent/config"
	"github.com/NorskHelsenett/ror/cmd/agent/services/authservice"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/go-resty/resty/v2"
)

// the function sends the resource to the ror api. If recieving a non 2xx statuscode it will retun an error.
func sendResourceUpdateToRor(resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error {
	rorClient, err := clients.GetOrCreateRorClient()
	if err != nil {
		rlog.Error("Could not get ror-api client", err)
		config.IncreaseErrorCount()
		return err
	}
	var url string
	var response *resty.Response

	if resourceUpdate.Action == apiresourcecontracts.K8sActionAdd {
		url = "/v1/resources"
		response, err = rorClient.R().
			SetHeader("Content-Type", "application/json").
			SetBody(resourceUpdate).
			Post(url)

	} else if resourceUpdate.Action == apiresourcecontracts.K8sActionUpdate {
		url = "/v1/resources/uid/" + resourceUpdate.Uid
		response, err = rorClient.R().
			SetHeader("Content-Type", "application/json").
			SetBody(resourceUpdate).
			Put(url)
	} else if resourceUpdate.Action == apiresourcecontracts.K8sActionDelete {
		url = "/v1/resources/uid/" + resourceUpdate.Uid
		response, err = rorClient.R().
			SetHeader("Content-Type", "application/json").
			SetBody(resourceUpdate).
			Delete(url)
	}

	if err != nil {
		config.IncreaseErrorCount()
		rlog.Error("could not send data to ror-api", err,
			rlog.Int("error count", config.ErrorCount))
		return err
	}

	if response == nil {
		config.IncreaseErrorCount()
		rlog.Error("response is nil", err,
			rlog.Int("error count", config.ErrorCount))
		return err
	}

	if !response.IsSuccess() {
		config.IncreaseErrorCount()
		rlog.Info("got non 200 statuscode from ror-api", rlog.Int("status code", response.StatusCode()),
			rlog.Int("error count", config.ErrorCount))
		return err
	} else {
		config.ResetErrorCount()
		rlog.Debug("partial update sent to ror", rlog.String("api verson", resourceUpdate.ApiVersion), rlog.String("kind", resourceUpdate.Kind), rlog.String("uid", resourceUpdate.Uid))
	}
	return nil
}

// function to get the persisted list of hashes from the api. The function is called on startup to populate the internal hashlist.
// The function makes the agent able to catch up on changes that has happened when its offline exluding deletes.
// TODO: Create a check to remove objects that are deleted during downtime of the agent.
func getResourceHashList() (hashList, error) {
	var hashlist hashList

	rorClient, err := clients.GetOrCreateRorClient()
	if err != nil {
		config.IncreaseErrorCount()
		rlog.Error("could not get ror-api client", err,
			rlog.Int("error count", config.ErrorCount))
		return hashlist, err
	}
	url := "/v1/resources/hashes"

	ownerref := authservice.CreateOwnerref()

	response, err := rorClient.R().
		SetQueryParams(ownerref.GetQueryParams()).
		SetHeader("Content-Type", "application/json").
		Get(url)

	if err != nil {
		config.IncreaseErrorCount()
		rlog.Error("could not send data to ror-api", err,
			rlog.Int("error count", config.ErrorCount))
		return hashlist, err
	}

	if response == nil {
		config.IncreaseErrorCount()
		rlog.Error("response is nil", err,
			rlog.Int("error count", config.ErrorCount))
		return hashlist, err
	}

	if !response.IsSuccess() {
		config.IncreaseErrorCount()
		rlog.Info("got non 200 statuscode from ror-api", rlog.Int("status code", response.StatusCode()),
			rlog.Int("error count", config.ErrorCount))
		return hashlist, err
	} else {
		config.ResetErrorCount()
		rlog.Info("hashList fetched from ror-api")

		err = json.Unmarshal(response.Body(), &hashlist)
		if err != nil {
			rlog.Error("could not unmarshal reply", err)
		}
	}
	return hashlist, nil
}
