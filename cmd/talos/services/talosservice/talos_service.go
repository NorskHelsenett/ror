package talosservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/talos/rorclient"
	"github.com/NorskHelsenett/ror/internal/provider/clusterorder/utils"
	"math/rand/v2"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"gopkg.in/yaml.v3"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/spf13/viper"
	"k8s.io/client-go/tools/clientcmd"
)

func ClusterOrderToClusterCreate(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	_, _ = fmt.Println("\n🚀  Starting to create cluster 🚀")
	_, _ = fmt.Println("\n👮‍♂️  Checking if cluster with same name exists")

	clusterExist, err := ClusterExists(ctx, clusterOrder.Spec.Cluster, true)
	if err != nil {
		rlog.Error("failed to check if cluster exists", err)
		return err
	}

	clusterOrder.Status.Status = "About to create cluster"
	clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseCreating
	updateClusterOrder(ctx, clusterOrder)

	if clusterExist {
		msg := fmt.Sprintf("cluster already exists with name: %s", clusterOrder.Spec.Cluster)
		rlog.Errorc(ctx, msg, nil)

		clusterOrder.Status.Status = "Cluster already exists"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)

		return errors.New(msg)
	}

	_, _ = fmt.Println("\n🚀  Creating talos cluster 🚀")
	err = createTalosCluster(*clusterOrder)
	if err != nil {
		clusterOrder.Status.Status = "Failed to create cluster"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	configFilePath := fmt.Sprintf("%s/%s/kube.config", viper.GetString("CONFIG_FOLDER_PATH"), clusterOrder.Spec.Cluster)
	_, err = modifyKubeconfig(configFilePath, clusterOrder.Spec.Cluster)
	if err != nil {
		clusterOrder.Status.Status = "Failed to modify kubeconfig"
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		updateClusterOrder(ctx, clusterOrder)
		return err
	}

	_, _ = fmt.Println("\n🚀  Installing cluster agent 🚀")
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

	_, _ = fmt.Println("\n🎉 Done creating cluster and installing cluster agent 🎉")

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

func createTalosCluster(clusterOrder apiresourcecontracts.ResourceClusterOrder) error {
	controlPlaneCount := 1
	if clusterOrder.Spec.HighAvailability {
		controlPlaneCount = 3
	}

	workersCount := 0
	for _, nodePool := range clusterOrder.Spec.NodePools {
		workersCount += nodePool.Count
	}
	if workersCount == 0 {
		workersCount = 1
	}

	var k8sversion string
	if len(clusterOrder.Spec.K8sVersion) > 0 {
		k8sversion = strings.Replace(clusterOrder.Spec.K8sVersion, "v", "", 1)
	} else {
		k8sversion = "1.30.3"
	}

	k8sPort, err := GetFreePort()
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "Unable to get free port"
		updateClusterOrder(context.Background(), &clusterOrder)
		rlog.Error("failed to get free port", err)
		return err
	}

	talosPort, err := GetFreePort()
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "Unable to get free port"
		updateClusterOrder(context.Background(), &clusterOrder)
		rlog.Error("failed to get free port", err)
		return err
	}
	uniqueSubNet := getRandomInt(6, 254)
	cidr := fmt.Sprintf("10.%d.0.0/24", uniqueSubNet)
	endpointAddress := fmt.Sprintf("host.docker.internal:%d", talosPort)
	if viper.GetBool(configconsts.DEVELOPMENT) {
		endpointAddress = fmt.Sprintf("127.0.0.1:%d", talosPort)
	}
	clusterFolderPath := fmt.Sprintf("%s/%s", viper.GetString("CONFIG_FOLDER_PATH"), clusterOrder.Spec.Cluster)

	cmdCreateClusterFolder := exec.Command(
		"mkdir",
		clusterFolderPath,
	) // #nosec G204 (ignore shell injection warning) all arguments are sanitized
	_, _ = fmt.Println(cmdCreateClusterFolder)
	outputCreateFolder, err := cmdCreateClusterFolder.CombinedOutput()
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "Unable to create folder for cluster"
		updateClusterOrder(context.Background(), &clusterOrder)
		rlog.Error("\n🛑  Failed to run mkdir command", err,
			rlog.String("mkdir output", string(outputCreateFolder)))
		return err
	}

	patchFolder := viper.GetString("TALOS_PATCH_FOLDER")
	k8sEndpointAddress := fmt.Sprintf("https://host.docker.internal:%d", k8sPort)
	if viper.GetBool(configconsts.DEVELOPMENT) {
		k8sEndpointAddress = fmt.Sprintf("https://127.0.0.1:%d", k8sPort)
	}
	orginalPatchFile := fmt.Sprintf("%s/%s", patchFolder, "patch.yaml")
	patchPath, err := updatePatchFile(orginalPatchFile, clusterFolderPath, k8sEndpointAddress, uniqueSubNet)
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "Unable to update patch file"
		updateClusterOrder(context.Background(), &clusterOrder)
		rlog.Error("failed to update patch file", err)
		return err
	}

	cmdCreateTalosconfig := exec.Command(
		"talosctl",
		"gen",
		"config",
		clusterOrder.Spec.Cluster,
		fmt.Sprintf("https://%s", endpointAddress),
		"--with-examples=false",
		fmt.Sprintf("--config-patch=@%s", patchPath),
		fmt.Sprintf(`--talosconfig=%s/talosconfig`, clusterFolderPath),
	) // #nosec G204 (ignore shell injection warning) all arguments are sanitized

	cmdCreateTalosconfig.Dir = clusterFolderPath
	_, _ = fmt.Printf("\n%s", cmdCreateTalosconfig)

	outputGenConfig, err := cmdCreateTalosconfig.CombinedOutput()
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "Unable to generate config for cluster"
		updateClusterOrder(context.Background(), &clusterOrder)
		rlog.Error("\n🛑  Failed to run talosctl gen config command", err,
			rlog.String("talosctl gen config output", string(outputGenConfig)))
		return err
	}

	_, _ = fmt.Println("\n Config generation output: ", string(outputGenConfig))

	cmdCreateCluster := exec.Command(
		"talosctl",
		"cluster",
		"create",
		fmt.Sprintf(`--name=%s`, clusterOrder.Spec.Cluster),
		fmt.Sprintf(`--workers=%d`, workersCount),
		fmt.Sprintf(`--controlplanes=%d`, controlPlaneCount),
		fmt.Sprintf(`--kubernetes-version=%s`, k8sversion),
		fmt.Sprintf(`--cidr=%s`, cidr),
		fmt.Sprintf(`--exposed-ports=%d:50000/tcp,%d:6443/tcp`, talosPort, k8sPort),
		fmt.Sprintf(`--endpoint=https://%s`, endpointAddress),
		fmt.Sprintf(`--config-patch=@%s`, patchPath),
		fmt.Sprintf(`--talosconfig=%s/talosconfig`, clusterFolderPath),
		//fmt.Sprintf(`--wait-timeout=%q`, "10m"),
		"--wait=false",
	) // #nosec G204 (ignore shell injection warning) all arguments are sanitized
	cmdCreateCluster.Dir = clusterFolderPath
	kubeconfigName := fmt.Sprintf("%s/%s/kube.config", viper.GetString("CONFIG_FOLDER_PATH"), clusterOrder.Spec.Cluster)
	cmdCreateCluster.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", kubeconfigName))
	_, _ = fmt.Printf("\n%s", cmdCreateCluster)

	output, err := cmdCreateCluster.CombinedOutput()
	if err != nil {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		clusterOrder.Status.Status = "Unable to create cluster"
		updateClusterOrder(context.Background(), &clusterOrder)
		rlog.Error("\n🛑  Failed to run talosctl command", err,
			rlog.String("talosctl output", string(output)))
		return err
	}
	_, _ = fmt.Println("\n Cluster create output: ", string(output))

	_, _ = fmt.Println("")
	// extracting kubeconfig
	cmdExtractKubeconfig := exec.Command(
		"talosctl",
		"kubeconfig",
		kubeconfigName,
		fmt.Sprintf(`-n=%s`, fmt.Sprintf("10.%d.0.2", uniqueSubNet)),
		fmt.Sprintf(`--talosconfig=%s/talosconfig`, clusterFolderPath),
	) // #nosec G204 (ignore shell injection warning) all arguments are sanitized
	cmdExtractKubeconfig.Dir = clusterFolderPath
	cmdExtractKubeconfig.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", kubeconfigName))
	_, _ = fmt.Printf("\n%s", cmdExtractKubeconfig)

	outputkubeconfig, err := cmdExtractKubeconfig.CombinedOutput()
	if err != nil {
		rlog.Error("\n🛑  Failed to run talosctl command", err,
			rlog.String("talosctl output", string(outputkubeconfig)))
		return err
	}

	var i int
	responseTimeout := 10 * time.Minute
	deadline := time.Now().Add(responseTimeout)
	waitTime := 10 * time.Second
	totalWaitTime := 0 * time.Second
	_, _ = fmt.Printf("\n⏰ Waiting (%s) for cluster to be created", waitTime)
	for time.Now().Before(deadline) {
		hasConnection, _ := HasK8sConnection(context.Background(), &clusterOrder)
		if hasConnection {
			break
		}

		totalWaitTime += waitTime
		time.Sleep(waitTime)
		_, _ = fmt.Printf("\n 🔂  Retrying to connect to cluster in %s, total wait time: %s", waitTime, totalWaitTime)
		i++
	}

	if time.Now().After(deadline) {
		clusterOrder.Status.Phase = apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed
		rlog.Error("timeout waiting for cluster to be created", nil)
		clusterOrder.Status.Status = "Timeout waiting for cluster to be created"
		updateClusterOrder(context.Background(), &clusterOrder)
		return errors.New("\n🛑  Timeout waiting for cluster to be created, giving up")
	}

	return nil
}

func updatePatchFile(patchPath string, clusterConfigFolder string, endpointAddress string, uniqueSubnet int) (string, error) {
	yamlContent, err := os.ReadFile(patchPath) // #nosec G304 (ignore file path injection warning) all arguments are sanitized
	if err != nil {
		rlog.Error("failed to read patch file", err)
		return "", err
	}

	var patchObject PatchObject
	err = yaml.Unmarshal(yamlContent, &patchObject)
	if err != nil {
		rlog.Error("failed to unmarshal yaml content", err)
		return "", err
	}

	patchObject.Cluster.ControlPlane.Endpoint = endpointAddress
	cpId := fmt.Sprintf("10.%d.0.2", uniqueSubnet)
	patchObject.Machine.CertSANs = []string{"127.0.0.1", "localhost", "host.docker.internal", cpId}

	patchContent, err := yaml.Marshal(patchObject)
	if err != nil {
		rlog.Error("failed to marshal yaml content", err)
		return "", err
	}

	patchResultPath := fmt.Sprintf("%s/patch.yaml", clusterConfigFolder)
	err = os.WriteFile(patchResultPath, patchContent, 0666) // #nosec G306 Will only run in docker, need to sett 666 to allow user read/write
	if err != nil {
		rlog.Error("failed to write patch file", err)
		return "", err
	}
	return patchResultPath, nil
}

func getRandomInt(minValue int, maxValue int) int {
	cidr := rand.IntN(maxValue-minValue+1) + minValue // #nosec G404 no need to use cryptographically secure random number
	return cidr
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
		fmt.Sprintf("environments.rorApiUrl=%s", viper.GetString("TALOS_"+configconsts.API_ENDPOINT)),
		"--set",
		fmt.Sprintf("environments.containerRegistryPrefix=%s", viper.GetString(configconsts.CONTAINER_REG_PREFIX)),
		"--set",
		fmt.Sprintf("image.repository=%s", viper.GetString("ROR_OPERATOR_IMAGE")),
		fmt.Sprintf("--kubeconfig=%s", configFilePath),
		"--kube-insecure-skip-tls-verify",
	) // #nosec G204 (ignore shell injection warning) all arguments are sanitized
	_, _ = fmt.Printf("\n%s", cmdInstallRorOperator)
	rorOpertorOutput, err := cmdInstallRorOperator.CombinedOutput()
	if err != nil {
		rlog.Error("\n🛑  Failed to run helm command", err,
			rlog.String("helm output", string(rorOpertorOutput)))
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

	//apiConfig := filterConfig(apiconfig, contextName)

	err = clientcmd.Validate(*apiconfig)
	if err != nil {
		rlog.Error("Failed to validate kubeconfig after filtering", err)
		return nil, err
	}

	configArray, err := clientcmd.Write(*apiconfig)
	if err != nil {
		rlog.Error("Failed to write kubeconfig", err)
		return nil, err
	}

	return configArray, nil
}

func ClusterOrderToClusterUpdate(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	return nil
}

func ClusterOrderToClusterDelete(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) error {
	cmdDeleteCluster := exec.Command(
		"talosctl",
		"cluster",
		"destroy",
		"--name",
		clusterOrder.Spec.Cluster,
	) // #nosec G204 (ignore shell injection warning) all arguments are sanitized
	_, _ = fmt.Printf("\n%s", cmdDeleteCluster)

	output, err := cmdDeleteCluster.CombinedOutput()
	if err != nil {
		rlog.Error("\n🛑  Failed to run talosctl command", err,
			rlog.String("talosctl output", string(output)))
		return err
	}

	return nil
}

func ClusterExists(ctx context.Context, clusterName string, firstCheck bool) (bool, error) {
	clusterFolderPath := fmt.Sprintf("%s/%s", viper.GetString("CONFIG_FOLDER_PATH"), clusterName)
	cmdCheckClusterExists := exec.Command(
		"talosctl",
		"cluster",
		"show",
	)
	if !firstCheck {
		cmdCheckClusterExists.Args = append(cmdCheckClusterExists.Args, fmt.Sprintf("--name=%s", clusterName))
		cmdCheckClusterExists.Args = append(cmdCheckClusterExists.Args, fmt.Sprintf("--talosconfig=%s/talosconfig", clusterFolderPath))
	}

	outputExists, err := cmdCheckClusterExists.CombinedOutput()
	if err != nil {
		rlog.Error("\n🛑  Failed to run Talos command", err,
			rlog.String("talosctl output", string(outputExists)))
		return false, err
	}

	_, _ = fmt.Println("\n Cluster exists output: ", string(outputExists))

	if strings.Contains(string(outputExists), clusterName) {
		return true, nil
	}

	return false, nil
}

func HasK8sConnection(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceClusterOrder) (bool, error) {
	configFilePath := fmt.Sprintf("%s/%s/kube.config", viper.GetString("CONFIG_FOLDER_PATH"), clusterOrder.Spec.Cluster)
	cmdCheckK8sConnection := exec.Command(
		"kubectl",
		"get",
		"nodes",
		fmt.Sprintf("--kubeconfig=%s", configFilePath),
		fmt.Sprintf("--request-timeout=%s", "2s"),
	) // #nosec G204 (ignore shell injection warning) all arguments are sanitized
	//_, _ = fmt.Printf("\n%s", cmdCheckK8sConnection)

	_, err := cmdCheckK8sConnection.CombinedOutput()
	if err != nil {
		if !strings.Contains(err.Error(), "exit status 1") {
			rlog.Error("\n🛑  Failed to run kubectl command", err)
		}
		return false, err
	}

	return true, nil
}

type PatchObject struct {
	Cluster struct {
		ControlPlane struct {
			Endpoint string `json:"endpoint" yaml:"endpoint"`
		} `json:"controlPlane" yaml:"controlPlane"`
		InlineManifests []struct {
			Contents string `json:"contents" yaml:"contents"`
			Name     string `json:"name" yaml:"name"`
		} `json:"inlineManifests" yaml:"inlineManifests"`
	} `json:"cluster" yaml:"cluster"`
	Machine struct {
		Features struct {
			HostDNS struct {
				Enabled              bool `json:"enabled" yaml:"enabled"`
				ForwardKubeDnsToHost bool `json:"forwardKubeDNSToHost" yaml:"forwardKubeDNSToHost"`
			} `json:"hostDNS" yaml:"hostDNS"`
		} `json:"features" yaml:"features"`
		CertSANs []string `json:"certSANs" yaml:"certSANs"`
	} `json:"machine" yaml:"machine"`
}
