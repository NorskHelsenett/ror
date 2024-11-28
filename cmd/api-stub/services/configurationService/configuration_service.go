package configurationservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/cmd/api-stub/apiconnections"
	clustersservice "github.com/NorskHelsenett/ror/cmd/api-stub/services/clustersService"
	"github.com/NorskHelsenett/ror/internal/clients/helsegitlab"
	"github.com/NorskHelsenett/ror/internal/configuration"
	"github.com/NorskHelsenett/ror/internal/factories/storagefactory"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/models/providers"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/spf13/viper"

	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type TaskName string

const (
	TASKNAME_UNKNOWN      TaskName = ""
	TASKNAME_ARGOCD       TaskName = "argocd-installer"
	TASKNAME_CLUSTERAGENT TaskName = "cluster-agent-installer"
	TASKNAME_NHNTOOLING   TaskName = "nhn-tooling-installer"
	CHART_VERSION                  = "0.1.*"
	ARGOCD_VERSION                 = "5.55.0"
)

func GetTaskConfigByClusterIdAndTaskName(ctx context.Context, task *apicontracts.Task, clusterId string) (apicontracts.OperatorJob, error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	cluster, err := clustersservice.GetByClusterId(ctx, clusterId)
	if err != nil {
		rlog.Error("could not get cluster", err)
	}

	operatorJob := apicontracts.OperatorJob{
		ImageName:        fmt.Sprintf("%s%s", viper.GetString(configconsts.CONTAINER_REG_IMAGE_PATH), task.Config.ImageName),
		ImageTag:         task.Config.Version,
		Cmd:              task.Config.Cmd,
		BackOffLimit:     task.Config.BackOffLimit,
		TimeOutInSeconds: task.Config.TimeOutInSeconds,
	}

	switch task.Name {
	case string(TASKNAME_ARGOCD):
		err = getArgoCdInstallerConfig(&operatorJob)
	case string(TASKNAME_CLUSTERAGENT):
		err = getClusterAgentInstallerConfig(&operatorJob, cluster)
	case string(TASKNAME_NHNTOOLING):
		err = getNHNToolingInstallerConfig(&operatorJob, cluster)
	default:
		return apicontracts.OperatorJob{}, errors.New("could not find taskSpec")
	}
	if err != nil {
		return apicontracts.OperatorJob{}, err
	}
	return operatorJob, nil
}

func getArgoCdInstallerConfig(job *apicontracts.OperatorJob) error {
	if job == nil {
		return errors.New("could not find taskSpec")
	}

	secretPath := "secret/data/v1.0/ror/config/common" // #nosec G101 Jest the path to the token file in the secrets engine
	seretKey := "argocdSdiPassword"
	jsonPath := "configs.credentialTemplates.helsegitlab-sdi-creds.password"

	generator := configuration.NewConfigurationsGenerator()

	generator.AddConfiguration(configuration.NewConfigurationLayer("values.yaml", 1, 1, configuration.NewHelsegitlabLoader(configuration.ParserTypeYaml, 428, "argocd/argominimal.yaml", "main", apiconnections.VaultClient)))
	generator.AddConfiguration(configuration.NewConfigurationLayer("helsegitlab-secret", 128, 1, configuration.NewSecretLoader(secretPath, seretKey, jsonPath, apiconnections.VaultClient)))

	data, err := generator.GenerateConfigYaml()
	if err != nil {
		return errors.New("error getting secret file content from helsegitlab for task")
	}

	entrypoint, err := helsegitlab.GetFileContent(428, "argocd/entrypoint.sh", "main", apiconnections.VaultClient)
	if err != nil {
		rlog.Error("Could not get file from helsegitlab", err)
		return nil
	}

	config2 := apicontracts.OperatorJobConfig{
		Name: "env",
		Type: apicontracts.OperatorJobConfigTypeEnv,
		Data: map[string]string{
			"ARGOCD_VERSION":   ARGOCD_VERSION,
			"VALUES_FILE_PATH": "/app/values.yaml",
		},
	}

	rolebindings, err := helsegitlab.GetFileContent(428, "argocd/rolebinding.yaml", "main", apiconnections.VaultClient)
	if err != nil {
		rlog.Error("Could not get file from helsegitlab", err)
		return nil
	}

	config1 := apicontracts.OperatorJobConfig{
		Name: "app",
		Type: apicontracts.OperatorJobConfigTypeFile,
		Path: "/app",
		Data: map[string]string{
			"values.yaml":      string(data),
			"entrypoint.sh":    string(entrypoint),
			"rolebinding.yaml": string(rolebindings),
		},
	}

	job.Configs = append(job.Configs, config1)
	job.Configs = append(job.Configs, config2)

	return nil
}

func getClusterAgentInstallerConfig(job *apicontracts.OperatorJob, cluster *apicontracts.Cluster) error {
	if job == nil {
		return errors.New("could not find taskSpec")
	}

	entrypoint, err := helsegitlab.GetFileContent(428, "ror-agent/entrypoint.sh", "main", apiconnections.VaultClient)
	if err != nil {
		rlog.Error("Could not get file from helsegitlab", err)
		return nil
	}

	config1 := apicontracts.OperatorJobConfig{
		Name: "app",
		Type: apicontracts.OperatorJobConfigTypeFile,
		Path: "/app",
		Data: map[string]string{
			"entrypoint.sh": string(entrypoint),
		},
	}

	data := map[string]string{
		"NAMESPACE":            "nhn-ror",
		"CHART_VERSION":        CHART_VERSION,
		"OCI_URL":              "oci://registry-1.docker.io/nhnhelm/cluster-agent",
		"ROR_URL":              viper.GetString(configconsts.LOCAL_KUBERNETES_ROR_BASE_URL),
		"CONTAINER_REG_PREFIX": "docker.io/",
	}

	if cluster.Workspace.Datacenter.Provider == providers.ProviderTypeK3d || cluster.Workspace.Datacenter.Provider == providers.ProviderTypeKind {
		data["NAMESPACE"] = "ror"
		data["OCI_URL"] = "oci://docker.io/nhnhelm/cluster-agent"
		data["CONTAINER_REG_PREFIX"] = "docker.io/"
		data["MORE_SETS"] = "--set image.repository=docker.io/nhnsdi/cluster-agent --set api=" + viper.GetString(configconsts.LOCAL_KUBERNETES_ROR_BASE_URL)
	}

	config2 := apicontracts.OperatorJobConfig{
		Name: "env",
		Type: apicontracts.OperatorJobConfigTypeEnv,
		Data: data,
	}
	job.Configs = append(job.Configs, config1)
	job.Configs = append(job.Configs, config2)

	return nil
}

func getNHNToolingInstallerConfig(job *apicontracts.OperatorJob, cluster *apicontracts.Cluster) error {
	if job == nil {
		return errors.New("could not find OperatorJob")
	}

	if cluster == nil {
		return errors.New("cluster not provided")
	}
	generator := configuration.NewConfigurationsGenerator()

	if cluster.Environment == "" {
		cluster.Environment = "dev"
	}

	jsonmap := map[string]string{
		"nhn.clusterName":                  cluster.ClusterName,
		"nhn.supervisorCluster":            cluster.Workspace.Name,
		"nhn.environment":                  cluster.Environment,
		"nhn.environmentnameoverride":      cluster.Environment,
		"nhn.cluster.storage.storageClass": storagefactory.GetStorageClassByDatacenter(cluster.Workspace.Datacenter.Name),
		"nhn.splunkConnect.clusterName":    cluster.ClusterName,
	}

	if cluster.Versions.NhnTooling.Branch == "" {
		cluster.Versions.NhnTooling.Branch = "v1"
	}

	var secretMap []configuration.SecretStruct
	secretMap = append(secretMap, configuration.SecretStruct{VaultPath: "secret/data/v1.0/ror/config/common", VaultKey: "argocdSdiPassword", JsonPath: "nhn.argocd.helsegitlabpassword"})
	secretMap = append(secretMap, configuration.SecretStruct{VaultPath: fmt.Sprintf("secret/data/v1.0/ror/dex/%s", cluster.ClusterId), VaultKey: "dexSecret", JsonPath: "nhn.argocd.argooicdsecret"})
	secretMap = append(secretMap, configuration.SecretStruct{VaultPath: "secret/data/v1.0/ror/config/common", VaultKey: "splunkHecToken", JsonPath: "nhn.splunkConnect.token"})

	// TODO: Legg til et valg for å sette i gui
	// toolingBranch := "staging"
	// if cluster.Versions.NhnTooling.Branch != "" && cluster.Versions.NhnTooling.Branch != "1.*" {
	// 	toolingBranch = cluster.Versions.NhnTooling.Branch
	// }

	generator.AddConfiguration(configuration.NewConfigurationLayer("json nhn", 254, 1, configuration.NewMapStringLoader(jsonmap)))
	generator.AddConfiguration(configuration.NewConfigurationLayer("secrets", 255, 1, configuration.NewSecretMapLoader(secretMap, apiconnections.VaultClient)))

	valuesdata, err := generator.GenerateConfigYaml()
	if err != nil {
		return errors.New("error generating config")
	}

	entrypoint, err := helsegitlab.GetFileContent(428, "nhn-tooling/entrypoint.sh", "main", apiconnections.VaultClient)
	if err != nil {
		rlog.Error("Could not get file from helsegitlab", err)
		return nil
	}

	config1 := apicontracts.OperatorJobConfig{
		Name: "app",
		Type: apicontracts.OperatorJobConfigTypeFile,
		Path: "/app",
		Data: map[string]string{
			"tooling.yaml":  string(valuesdata),
			"entrypoint.sh": string(entrypoint),
		},
	}

	data := map[string]string{
		"VERSION":    "1.*",
		"OCI_PATH":   "oci://ncr.sky.nhn.no/nhn/nhn-tooling-application",
		"CLUSTER_ID": cluster.ClusterId,
	}

	config2 := apicontracts.OperatorJobConfig{
		Name: "env",
		Type: apicontracts.OperatorJobConfigTypeEnv,
		Data: data,
	}

	job.Configs = append(job.Configs, config1)
	job.Configs = append(job.Configs, config2)

	return nil
}
