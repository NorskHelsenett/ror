package kubeconfigcontroller

import (
	"encoding/json"
	"errors"
	"github.com/NorskHelsenett/ror/cmd/tanzu/auth/auth"
	"net/http"
	"os"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorerror"

	"github.com/NorskHelsenett/ror/pkg/rlog"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func HandleKubeConfigRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		rlog.Error("Method not allowed", nil)
		return
	}

	var credentials apicontracts.TanzuKubeConfigPayload
	err := json.NewDecoder(req.Body).Decode(&credentials)
	if err != nil {
		rlog.Error("Failed to decode request body", err)
		handleError(http.StatusInternalServerError, err, w)
		return
	}

	kubeconfigPath, err := auth.Login(credentials)
	if err != nil {
		rlog.Error("Failed to login", err)
		handleError(http.StatusInternalServerError, err, w)
		return
	}

	if len(kubeconfigPath) == 0 {
		rlog.Error("Kubeconfig path is empty", nil)
		handleError(http.StatusNotFound, err, w)
		return
	}

	var extractName string
	if credentials.WorkspaceOnly {
		extractName = credentials.WorkspaceName
	} else {
		extractName = credentials.ClusterName
	}

	if len(extractName) == 0 {
		rlog.Error("context name is missing", nil)
		handleError(http.StatusBadRequest, err, w)
		return
	}

	config, err := extractKubeConfig(kubeconfigPath, extractName)
	if err != nil {
		rlog.Error("Failed to extract kubeconfig", err)
		handleError(http.StatusInternalServerError, err, w)
		return
	}

	err = deleteConfig(kubeconfigPath)
	if err != nil {
		rlog.Error("Failed to delete kubeconfig", err)
	}

	_, err = w.Write(config)
	if err != nil {
		handleError(http.StatusInternalServerError, err, w)
		return
	}
}

func handleError(statusCode int, err error, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	error := rorerror.RorError{
		Status:  statusCode,
		Message: err.Error(),
	}

	errorString, err := json.Marshal(error)
	if err != nil {
		rlog.Error("Failed to marshal error", err)
		return
	}

	_, _ = w.Write([]byte(errorString))
}

func extractKubeConfig(kubeconfigPath string, contextName string) ([]byte, error) {
	if len(kubeconfigPath) == 0 {
		rlog.Error("kubeconfig path is empty", nil)
		return nil, errors.New("kubeconfig path is empty")
	}

	if len(contextName) == 0 {
		rlog.Error("context name is empty", nil)
		return nil, errors.New("context name is empty")
	}

	apiconfig, err := clientcmd.LoadFromFile(kubeconfigPath)
	if err != nil {
		rlog.Error("Failed to load kubeconfig", err)
		return nil, err
	}

	err = clientcmd.Validate(*apiconfig)
	if err != nil {
		rlog.Error("Failed to validate kubeconfig", err)
		return nil, err
	}

	apiConfig := filterConfig(apiconfig, contextName)

	err = clientcmd.Validate(*apiconfig)
	if err != nil {
		rlog.Error("Failed to validate kubeconfig after filtering", err)
		return nil, err
	}

	configArray, err := clientcmd.Write(*apiConfig)
	if err != nil {
		rlog.Error("Failed to write kubeconfig", err)
		return nil, err
	}

	return configArray, nil
}

func filterConfig(apiconfig *api.Config, contextName string) *api.Config {
	contexts := make(map[string]*api.Context)
	var context *api.Context
	var user string
	apiConfig := api.NewConfig()
	apiConfig.CurrentContext = apiconfig.CurrentContext
	for cName, c := range apiconfig.Contexts {
		if cName == contextName {
			context = c
			user = string(c.AuthInfo)
			contexts[cName] = context
			break
		}
	}
	apiConfig.Contexts = contexts

	for cName, c := range apiconfig.Clusters {
		if cName == context.Cluster {
			apiConfig.Clusters[cName] = c
			break
		}
	}

	for uName, u := range apiconfig.AuthInfos {
		if uName == user {
			apiConfig.AuthInfos[uName] = u
			break
		}
	}

	apiConfig.CurrentContext = contextName
	return apiConfig
}

func deleteConfig(kubeconfigPath string) error {
	err := os.Remove(kubeconfigPath)
	if err != nil {
		rlog.Error("Failed to delete kubeconfig", err)
		return err
	}
	return nil
}
