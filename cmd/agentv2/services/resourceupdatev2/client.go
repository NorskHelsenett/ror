package resourceupdatev2

import (
	"github.com/NorskHelsenett/ror/cmd/agentv2/agentconfig"
	"github.com/NorskHelsenett/ror/cmd/agentv2/clients"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/go-resty/resty/v2"
)

// the function sends the resource to the ror api. If recieving a non 2xx statuscode it will retun an error.
func sendResourceUpdateToRor(resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error {
	rorClient, err := clients.GetOrCreateRorClient()
	if err != nil {
		rlog.Error("Could not get ror-api client", err)
		agentconfig.IncreaseErrorCount()
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
		agentconfig.IncreaseErrorCount()
		rlog.Error("could not send data to ror-api", err,
			rlog.Int("error count", agentconfig.ErrorCount))
		return err
	}

	if response == nil {
		agentconfig.IncreaseErrorCount()
		rlog.Error("response is nil", err,
			rlog.Int("error count", agentconfig.ErrorCount))
		return err
	}

	if !response.IsSuccess() {
		agentconfig.IncreaseErrorCount()
		rlog.Info("got non 200 statuscode from ror-api", rlog.Int("status code", response.StatusCode()),
			rlog.Int("error count", agentconfig.ErrorCount))
		return err
	} else {
		agentconfig.ResetErrorCount()
		rlog.Debug("partial update sent to ror", rlog.String("api verson", resourceUpdate.ApiVersion), rlog.String("kind", resourceUpdate.Kind), rlog.String("uid", resourceUpdate.Uid))
	}
	return nil
}
