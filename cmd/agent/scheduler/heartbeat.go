package scheduler

import (
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/agent/clients"
	"github.com/NorskHelsenett/ror/cmd/agent/config"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/cmd/agent/services"
)

func HeartbeatReporting() error {
	clusterReport, err := services.GetHeartbeatReport()
	if err != nil {
		rlog.Error("error when getting heartbeat report", err)
		return err
	}

	err = sendReportToRor(clusterReport)
	return err
}

func sendReportToRor(clusterReport apicontracts.Cluster) error {
	rorClient, err := clients.GetOrCreateRorClient()
	if err != nil {
		config.IncreaseErrorCount()
		rlog.Error("could not get ror-api client", err,
			rlog.Int("error count", config.ErrorCount))
		return err
	}

	url := "/v1/cluster/heartbeat"
	response, err := rorClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(clusterReport).
		Post(url)
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
		rlog.Error("got unsuccessful status code from ror-api", err, rlog.Int("status code", response.StatusCode()))
		return err
	} else {
		config.ResetErrorCount()
		rlog.Info("heartbeat report sent to ror")

		byteReport, err := json.Marshal(clusterReport)
		if err == nil {
			rlog.Debug("", rlog.String("byte report", string(byteReport)))
		}
	}
	return nil
}
