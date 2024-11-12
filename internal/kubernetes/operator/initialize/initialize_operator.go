package initialize

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/NorskHelsenett/ror/internal/kubernetes/nodeservice"
	"github.com/NorskHelsenett/ror/internal/models/operatormodels"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// Deprecated: GetApikey is deprecated use something like cmd\agentv2\clients\rorconfig.go instead
func GetApikey(clusterInfo *operatormodels.ClusterInfo, rorUrl string) (string, error) {
	if clusterInfo == nil {
		return "", errors.New("identifier is empty, cannot fetch api key")
	}

	client := http.Client{Timeout: time.Duration(20) * time.Second}
	jsonData, _ := json.Marshal(apicontracts.AgentApiKeyModel{
		Identifier:     clusterInfo.ClusterName,
		DatacenterName: clusterInfo.DatacenterName,
		WorkspaceName:  clusterInfo.WorkspaceName,
		Provider:       clusterInfo.Provider,
		Type:           "Cluster",
	})

	requestUrl := fmt.Sprintf("%s/v1/clusters/register", rorUrl)
	response, err := client.Post(requestUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		rlog.Error("error when getting api key", err)
		return "", errors.New("could not send data to API")
	}

	if response.StatusCode > 299 {
		bodyByte, err := io.ReadAll(response.Body)
		if err != nil {
			rlog.Fatal("Response body", err, rlog.ByteString("body", bodyByte))
		}
		rlog.Error("could not get api key from API", fmt.Errorf("non 200 response code"),
			rlog.String("identifier", clusterInfo.ClusterName),
			rlog.Int("response code", response.StatusCode),
			rlog.String("url", rorUrl))

		return "", fmt.Errorf("could not get api key from API, identifier: %s, rorUrl: %s", clusterInfo.ClusterName, rorUrl)
	}

	body := response.Body
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		rlog.Error("could not read body", err)
		return "", errors.New("could not read data from response")
	}

	apikey := string(bodyByte)
	return apikey, nil
}

// Deprecated: GetClusterInfoFromNode is deprecated use kubernetes interegator instead
func GetClusterInfoFromNode(k8sClient *kubernetes.Clientset, metricsClient *metrics.Clientset) (*operatormodels.ClusterInfo, error) {
	nodes, err := nodeservice.GetNodes(k8sClient, metricsClient)
	if err != nil {
		rlog.Error("could not get nodes", err)
		return nil, errors.New("could not get clusterId")
	}

	if len(nodes) < 1 {
		rlog.Error("nodes list is empty", nil)
		return nil, errors.New("nodes list is empty")
	}

	firstNode := nodes[0]
	clusterName := firstNode.ClusterName
	workspaceName := firstNode.Workspace

	clusterId := fmt.Sprintf("%s.%s", clusterName, workspaceName)
	clusterInfo := &operatormodels.ClusterInfo{
		Id:             clusterId,
		ClusterName:    clusterName,
		DatacenterName: firstNode.Datacenter,
		WorkspaceName:  firstNode.Workspace,
		Provider:       firstNode.Provider,
	}

	return clusterInfo, nil
}

func GetOwnClusterId() (string, error) {
	apikey := viper.GetString(configconsts.API_KEY)
	rorUrl := viper.GetString(configconsts.API_ENDPOINT)

	client := http.Client{Timeout: time.Duration(10) * time.Second}
	request, err := http.NewRequest("GET", rorUrl+"/v1/clusters/self", nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", apikey)

	response, err := client.Do(request)
	if err != nil {
		rlog.Error("error when cluster self info", err)
		return "", errors.New("could not get data from ROR-API")
	}

	if response.StatusCode > 299 {
		bodyByte, err := io.ReadAll(response.Body)
		if err != nil {
			rlog.Fatal("response body: ", err, rlog.ByteString("bytes", bodyByte))
		}

		return "", fmt.Errorf("could not get cluster self data from API, rorUrl: %s", rorUrl)
	}

	body := response.Body
	bodyByte, err := io.ReadAll(body)
	if err != nil {
		rlog.Error("could not read body", err)
		return "", errors.New("could not read data from response")
	}

	var clusterSelf apicontracts.ClusterSelf
	err = json.Unmarshal(bodyByte, &clusterSelf)
	if err != nil {
		rlog.Error("could not unmarshal body", err)
		rlog.Fatal("Could not fetch secret", err)
	}

	return clusterSelf.ClusterId, nil
}
