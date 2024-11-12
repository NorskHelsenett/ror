package kubeconfigservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/spf13/viper"
)

var httpClient = http.Client{Timeout: 55 * time.Second}

func GetKubeconfig(cluster *apicontracts.Cluster, credentials apicontracts.KubeconfigCredentials) (string, error) {
	switch cluster.Workspace.Datacenter.Provider {
	case providers.ProviderTypeTanzu:
		return getKubeconfigForTanzuCluster(cluster, credentials)
	default:
		return "", fmt.Errorf("provider %s is not supported", cluster.Workspace.Datacenter.Provider)
	}
}

func GetKubeconfigForWorkspace(workspace *apicontracts.Workspace, credentials apicontracts.KubeconfigCredentials) (string, error) {
	switch workspace.Datacenter.Provider {
	case providers.ProviderTypeTanzu:
		return getKubeconfigForTanzuWorkspace(workspace, credentials)
	default:
		return "", fmt.Errorf("provider %s is not supported", workspace.Datacenter.Provider)
	}
}

func getKubeconfigForTanzuCluster(cluster *apicontracts.Cluster, credentials apicontracts.KubeconfigCredentials) (string, error) {
	creds := apicontracts.TanzuKubeConfigPayload{
		User:          credentials.Username,
		Password:      credentials.Password,
		DatacenterUrl: cluster.Workspace.Datacenter.APIEndpoint,
		WorkspaceName: cluster.Workspace.Name,
		ClusterName:   cluster.ClusterName,
		ClusterId:     cluster.ClusterId,
		WorkspaceOnly: false,
	}

	return getKubeconfig(creds)
}

func getKubeconfigForTanzuWorkspace(workspace *apicontracts.Workspace, credentials apicontracts.KubeconfigCredentials) (string, error) {
	creds := apicontracts.TanzuKubeConfigPayload{
		User:          credentials.Username,
		Password:      credentials.Password,
		DatacenterUrl: workspace.Datacenter.APIEndpoint,
		WorkspaceName: workspace.Name,
		ClusterName:   "",
		ClusterId:     "",
		WorkspaceOnly: true,
	}

	return getKubeconfig(creds)
}

func getKubeconfig(configPayload apicontracts.TanzuKubeConfigPayload) (string, error) {
	var payload bytes.Buffer
	err := json.NewEncoder(&payload).Encode(configPayload)
	if err != nil {
		rlog.Error("failed to encode payload", err)
		return "", err
	}

	serviceUrl := viper.GetString(configconsts.TANZU_AUTH_BASE_URL)
	httpposturl := fmt.Sprintf("%s/v1/kubeconfig", serviceUrl)
	request, err := http.NewRequest("POST", httpposturl, &payload)
	if err != nil {
		rlog.Error("failed to create request", err)
		return "", err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := httpClient.Do(request)
	if err != nil {
		rlog.Error("failed to get kubeconfig", err)
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to get kubeconfig, status code: %d", response.StatusCode)
		rlog.Error("error", err)
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		rlog.Error("failed to read response body", err)
		return "", err
	}

	return string(body), nil
}
