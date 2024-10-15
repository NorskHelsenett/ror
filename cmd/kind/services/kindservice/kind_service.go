package kindservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/kind/rorclient"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/utils"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"

	"gopkg.in/yaml.v3"
)

func ClusterOrderToClusterCreate(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	cmdCheckKindClusterExists := exec.Command(
		"kind",
		"get",
		"clusters",
	)
	_, _ = fmt.Println(cmdCheckKindClusterExists)
	outputExists, err := cmdCheckKindClusterExists.CombinedOutput()
	if err != nil {
		rlog.Error("failed to run Kind command", err,
			rlog.String("kind output", string(outputExists)))
		return err
	}

	clusterOrder.Status.Status = "About to create cluster"
	clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseCreating
	updateClusterOrder(ctx, clusterOrder)

	if strings.Contains(string(outputExists), clusterOrder.Spec.Cluster) {
		msg := fmt.Sprintf("cluster already exists with name: %s", clusterOrder.Spec.Cluster)
		rlog.Errorc(ctx, msg, nil)

		clusterOrder.Status.Status = "Cluster already exists"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)

		return errors.New(msg)
	}

	err = createKindCluster(*clusterOrder)
	if err != nil {
		clusterOrder.Status.Status = "Failed to create cluster"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	configFilePath, err := extractKubeconfig(*clusterOrder)
	if err != nil {
		clusterOrder.Status.Status = "Failed to extract kubeconfig"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	configContent, err := modifyKubeconfig(configFilePath, fmt.Sprintf("kind-kind-%s", clusterOrder.Spec.Cluster))
	if err != nil {
		clusterOrder.Status.Status = "Failed to modify kubeconfig"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	configFilePath, err = writeKubeconfig(configContent, *clusterOrder)
	if err != nil {
		clusterOrder.Status.Status = "Failed to write kubeconfig"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	err = installKindExternalNameService(configFilePath)
	if err != nil {
		clusterOrder.Status.Status = "Failed to install external name service"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	err = installClusterAgent(configFilePath)
	if err != nil {
		clusterOrder.Status.Status = "Failed to install cluster agent"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	clusterOrder.Status.Status = "Cluster created"
	clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseCompleted
	updateClusterOrder(ctx, clusterOrder)

	return nil
}

func ClusterOrderToClusterUpdate(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	return nil
}

func ClusterOrderToClusterDelete(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	cmdDeletekindCluster := exec.Command(
		"kind",
		"delete",
		"cluster",
		stringhelper.EscapeString(clusterOrder.Spec.Cluster),
	) // #nosec G204 (ignore shell injection warning) All arguments are sanitized
	_, _ = fmt.Println(cmdDeletekindCluster)

	output, err := cmdDeletekindCluster.CombinedOutput()
	if err != nil {
		rlog.Error("failed to run kind command", err,
			rlog.String("kind output", string(output)))
		return err
	}

	return nil
}

func GetFreePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}

func createKindCluster(clusterOrder apiresourcecontracts.ResourceClusterOrder) error {
	image := viper.GetString("DEFAULT_KIND_IMAGE")
	if clusterOrder.Spec.K8sVersion != "" {
		image = clusterOrder.Spec.K8sVersion
	}

	nodes := []KindNode{
		{
			Role:  KindRoleControlPlane,
			Image: image,
		},
	}

	if clusterOrder.Spec.HighAvailability {
		nodes = append(nodes, KindNode{
			Role:  KindRoleControlPlane,
			Image: image,
		})
		nodes = append(nodes, KindNode{
			Role:  KindRoleControlPlane,
			Image: image,
		})
	}

	for range clusterOrder.Spec.NodePools {
		for i := 0; i < clusterOrder.Spec.NodePools[0].Count; i++ {
			nodes = append(nodes, KindNode{
				Role:  KindRoleWorker,
				Image: image,
			})
		}
	}

	kindConfig := KindConfig{
		Kind:       "Cluster",
		ApiVersion: "kind.x-k8s.io/v1alpha4",
		Nodes:      nodes,
	}

	kindConfigBytes, err := yaml.Marshal(kindConfig)
	if err != nil {
		return err
	}

	_, _ = fmt.Println(string(kindConfigBytes))
	cmdCreatekindCluster := exec.Command(
		"sh",
		"-c",
		"cat <<EOF | kind create cluster --name "+
			fmt.Sprintf("kind-%q", clusterOrder.Spec.Cluster)+
			" --config - "+
			fmt.Sprintf("\n%s", string(kindConfigBytes))+
			"EOF",
	) // #nosec G204 (ignore shell injection warning) All arguments are sanitized
	_, _ = fmt.Println(cmdCreatekindCluster)

	output, err := cmdCreatekindCluster.CombinedOutput()
	if err != nil {
		rlog.Error("failed to run kind command", err,
			rlog.String("kind output", string(output)))
		return err
	}
	return nil
}

func installKindExternalNameService(configFilePath string) error {
	cmdAddService := exec.Command(
		"sh",
		"-c",
		"cat <<EOF | kubectl --insecure-skip-tls-verify apply -f -"+
			`
kind: Service
apiVersion: v1
metadata:
  name: backend-ide
spec:
  type: ExternalName
  externalName: host.docker.internal`+
			"\nEOF",
	)
	cmdAddService.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", configFilePath))
	_, _ = fmt.Println(cmdAddService)

	outputService, err := cmdAddService.CombinedOutput()
	if err != nil {
		rlog.Error("failed to run kubectl command", err,
			rlog.String("kubectl output", string(outputService)))
		return err
	}
	return nil
}

func installClusterAgent(configFilePath string) error {
	cmdInstallRorOperator := exec.Command(
		"helm",
		"upgrade",
		"--install",
		"ror-operator",
		viper.GetString("ROR_OPERATOR_OCI_IMAGE"),
		"--version",
		viper.GetString("ROR_OPERATOR_OCI_IMAGE_VERSION"),
		"-n",
		viper.GetString(configconsts.ROR_OPERATOR_NAMESPACE),
		"--create-namespace",
		"--set",
		fmt.Sprintf("environments.rorApiUrl=%s", viper.GetString("KIND_"+configconsts.API_ENDPOINT)),
		"--set",
		fmt.Sprintf("environments.containerRegistryPrefix=%s", viper.GetString(configconsts.CONTAINER_REG_PREFIX)),
		"--set",
		fmt.Sprintf("image.repository=%s", viper.GetString("ROR_OPERATOR_IMAGE")),
		"--kube-insecure-skip-tls-verify",
	) // #nosec G204 (ignore shell injection warning) All arguments are from environment variables
	cmdInstallRorOperator.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", configFilePath))
	_, _ = fmt.Println(cmdInstallRorOperator)
	rorOpertorOutput, err := cmdInstallRorOperator.CombinedOutput()
	if err != nil {
		rlog.Error("failed to run kind command", err,
			rlog.String("kind output", string(rorOpertorOutput)))
		return err
	}
	return nil
}

func updateClusterOrder(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) {
	if clusterOrder.Metadata.Uid == "" {
		rlog.Error("cluster order uid is empty", nil)
		return
	}

	updateModel, err := utils.NewResourceUpdate(ctx, *clusterOrder)
	if err != nil {
		rlog.Error("failed to create update model", err)
		return
	}

	err = rorclient.RorClient.Resources().UpdateClusterOrder(&updateModel)
	if err != nil {
		rlog.Error("failed to update cluster order", err)
		return
	}
}

func extractKubeconfig(clusterOrder apiresourcecontracts.ResourceClusterOrder) (string, error) {
	cmdGetKubeconfig := exec.Command(
		"sh",
		"-c",
		"kind get kubeconfig --name "+
			fmt.Sprintf("kind-%q", clusterOrder.Spec.Cluster),
	) // #nosec G204 (ignore shell injection warning) All arguments are sanitized
	_, _ = fmt.Println(cmdGetKubeconfig)

	outputKubeconfig, err := cmdGetKubeconfig.CombinedOutput()
	if err != nil {
		rlog.Error("failed to run kind command", err,
			rlog.String("kind output", string(outputKubeconfig)))
		return "", err
	}

	configFilePath, err := writeKubeconfig(outputKubeconfig, clusterOrder)
	if err != nil {
		return "", err
	}

	return configFilePath, nil
}

func writeKubeconfig(kubeconfigContent []byte, clusterOrder apiresourcecontracts.ResourceClusterOrder) (string, error) {
	kubeconfigPath := fmt.Sprintf("%s/kind-%s.config", viper.GetString("CONFIG_FOLDER_PATH"), clusterOrder.Spec.Cluster)

	err := os.WriteFile(kubeconfigPath, kubeconfigContent, 0666) // #nosec G306 Will only run in docker, need to sett 666 to allow user read/write
	if err != nil {
		rlog.Error("failed to write kubeconfig", err)
		return "", err
	}

	return kubeconfigPath, nil
}

func modifyKubeconfig(kubeconfigPath string, contextName string) ([]byte, error) {
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
		serverArray := strings.Split(c.Server, ":")
		port := serverArray[len(serverArray)-1]
		c.Server = fmt.Sprintf("%s:%s", viper.GetString("KUBECTL_BASE_URL"), port)

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

type KindConfig struct {
	Kind       string         `json:"kind" yaml:"kind"`
	ApiVersion string         `json:"apiVersion" yaml:"apiVersion"`
	Nodes      []KindNode     `json:"nodes" yaml:"nodes"`
	Networking KindNetworking `json:"networking,omitempty" yaml:"networking,omitempty"`
}

type KindNode struct {
	Role              KindRole               `json:"role" yaml:"role"`
	ExtraPortMappings []KindExtraPortMapping `json:"extraPortMappings,omitempty" yaml:"extraPortMappings,omitempty"`
	Image             string                 `json:"image,omitempty" yaml:"image,omitempty"`
}

type KindExtraPortMapping struct {
	ContainerPort int    `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`
	HostPort      int    `json:"hostPort,omitempty" yaml:"hostPort,omitempty"`
	Protocol      string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}

type KindNetworking struct {
	ApiServerPort     int          `json:"apiServerPort" yaml:"apiServerPort"`
	ApiServerAddress  string       `json:"apiServerAddress" yaml:"apiServerAddress"`
	IpFamiliy         KindIpFamily `json:"ipFamiliy" yaml:"ipFamiliy"`
	DisableDefaultCNI bool         `json:"disableDefaultCNI" yaml:"disableDefaultCNI"`
}

type KindRole string

const (
	KindRoleControlPlane KindRole = "control-plane"
	KindRoleWorker       KindRole = "worker"
)

type KindIpFamily string

const (
	KindIpFamilyDualStack KindIpFamily = "dual"
	KindIpFamilyIPv4      KindIpFamily = "ipv4"
	KindIpFamilyIPv6      KindIpFamily = "ipv6"
)
