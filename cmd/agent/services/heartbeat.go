package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/agent/clients"
	"github.com/NorskHelsenett/ror/cmd/agent/models"
	"github.com/NorskHelsenett/ror/cmd/agent/utils"
	"github.com/NorskHelsenett/ror/internal/kubernetes/k8smodels"
	"github.com/NorskHelsenett/ror/internal/kubernetes/nodeservice"
	"strings"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

var MissingConst = "Missing ..."

func GetHeartbeatReport() (apicontracts.Cluster, error) {

	k8sClient, err := clients.Kubernetes.GetKubernetesClientset()
	if err != nil {
		return apicontracts.Cluster{}, err
	}

	dynamicClient, err := clients.Kubernetes.GetDynamicClient()
	if err != nil {
		return apicontracts.Cluster{}, err
	}
	metricsClient, err := clients.Kubernetes.GetMetricsClient()
	if err != nil {
		return apicontracts.Cluster{}, err
	}

	clusterName := "localhost"
	workspaceName := "localhost"
	datacenterName := "local"
	provider := providers.ProviderTypeUnknown

	nhnToolingMetadata, err := getNhnToolingMetadata(k8sClient, dynamicClient)
	if err != nil {
		rlog.Warn("NHN-Tooling is not installed?!")
	}

	k8sVersion, err := k8sClient.ServerVersion()
	if err != nil {
		rlog.Error("could not get kubernetes server version", err)
	}

	kubernetesVersion := MissingConst
	if k8sVersion != nil {
		kubernetesVersion = k8sVersion.String()
	}

	k8sVersionArray := strings.Split(kubernetesVersion, "+")
	if len(k8sVersionArray) > 1 {
		kubernetesVersion = k8sVersionArray[0]
	}

	nodes, err := nodeservice.GetNodes(k8sClient, metricsClient)
	if err != nil {
		rlog.Error("error getting nodes", err)
	}

	var k8sControlPlaneEndpoint string
	var controlPlane = apicontracts.ControlPlane{}
	nodePools := make([]apicontracts.NodePool, 0)
	if len(nodes) > 0 {
		for _, node := range nodes {
			if _, ok := node.Labels["node-role.kubernetes.io/control-plane"]; ok {
				appendNodeToControlePlane(&node, &controlPlane)
				if node.Provider == "tanzu" {
					k8sControlPlaneEndpoint, _ = getControlPlaneEndpoint(k8sClient)
				}
			} else {
				appendNodeToNodePools(&nodePools, &node)
			}
		}
		if err != nil {
			rlog.Error("could not extract controlplane and workers from nodes", err)
		}
	}

	if len(nodes) > 0 {
		firstNode := nodes[0]
		clusterName = firstNode.ClusterName
		workspaceName = firstNode.Workspace
		datacenterName = firstNode.Datacenter
		provider = firstNode.Provider
	}

	ingresses, err := getIngresses(k8sClient)
	if err != nil {
		rlog.Error("could not get ingresses", err)
	}

	nodeCount := int64(0)
	cpuSum := int64(0)
	cpuConsumedSum := int64(0)
	memorySum := int64(0)
	memoryConsumedSum := int64(0)
	for i := 0; i < len(nodePools); i++ {
		nodepool := nodePools[i]
		nodeCount = nodeCount + nodepool.Metrics.NodeCount
		cpuSum = cpuSum + nodepool.Metrics.Cpu
		cpuConsumedSum = cpuConsumedSum + nodepool.Metrics.CpuConsumed
		memorySum = memorySum + nodepool.Metrics.Memory
		memoryConsumedSum = memoryConsumedSum + nodepool.Metrics.MemoryConsumed
	}

	agentVersion := viper.GetString(configconsts.VERSION)
	agentSha := viper.GetString(configconsts.COMMIT)

	var created time.Time
	kubeSystem := "kube-system"
	kubeSystemNamespace, err := k8sClient.CoreV1().Namespaces().Get(context.Background(), kubeSystem, v1.GetOptions{})
	if err != nil {
		rlog.Error("could not fetch namespace", err, rlog.String("namespace", kubeSystem))
	} else {
		created = kubeSystemNamespace.CreationTimestamp.Time
	}

	report := apicontracts.Cluster{
		ACL: apicontracts.AccessControlList{
			AccessGroups: nhnToolingMetadata.AccessGroups,
		},
		Environment: nhnToolingMetadata.Environment,
		ClusterId:   viper.GetString(configconsts.CLUSTER_ID),
		ClusterName: clusterName,
		Ingresses:   ingresses,
		Created:     created,
		Topology: apicontracts.Topology{
			ControlPlaneEndpoint: k8sControlPlaneEndpoint,
			EgressIp:             EgressIp,
			ControlPlane:         controlPlane,
			NodePools:            nodePools,
		},
		Versions: apicontracts.Versions{
			Kubernetes: kubernetesVersion,
			NhnTooling: apicontracts.NhnTooling{
				Version:     nhnToolingMetadata.Version,
				Branch:      nhnToolingMetadata.Branch,
				Environment: nhnToolingMetadata.Environment,
			},
			Agent: apicontracts.Agent{
				Version: agentVersion,
				Sha:     agentSha,
			},
		},
		Metrics: apicontracts.Metrics{
			NodeCount:      nodeCount,
			NodePoolCount:  int64(len(nodePools)),
			Cpu:            cpuSum,
			CpuConsumed:    cpuConsumedSum,
			Memory:         memorySum,
			MemoryConsumed: memoryConsumedSum,
			ClusterCount:   1,
		},
		Workspace: apicontracts.Workspace{
			Name: workspaceName,
			Datacenter: apicontracts.Datacenter{
				Name:     datacenterName,
				Provider: provider,
			},
		},
	}
	return report, nil
}

func getIngresses(k8sClient *kubernetes.Clientset) ([]apicontracts.Ingress, error) {
	var ingressList []apicontracts.Ingress
	nsList, err := k8sClient.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		rlog.Error("could not fetch namespaces", err)
		return ingressList, errors.New("could not fetch namespaces from cluster")
	}

	for _, namespace := range nsList.Items {
		ing := k8sClient.NetworkingV1().Ingresses(namespace.Name)
		ingresses, err := ing.List(context.TODO(), v1.ListOptions{})
		if err != nil {
			rlog.Error("could not list ingress in namespace", err, rlog.String("namespace", namespace.Name))
			continue
		}
		for _, ingress := range ingresses.Items {

			richIngress, err := utils.GetIngressDetails(&ingress)
			if err != nil {
				rlog.Error("could not enrich ingress", err,
					rlog.String("ingress", ingress.Name),
					rlog.String("namespace", namespace.Name))
				continue
			} else {
				ingressList = append(ingressList, *richIngress)
			}
		}
	}

	return ingressList, nil
}

func getNhnToolingMetadata(k8sClient *kubernetes.Clientset, dynamicClient dynamic.Interface) (k8smodels.NhnTooling, error) {
	var accessGroups []string
	result := k8smodels.NhnTooling{
		Version:      MissingConst,
		Branch:       MissingConst,
		AccessGroups: []string{},
		Environment:  "dev",
	}

	namespace := viper.GetString(configconsts.POD_NAMESPACE)
	nhnToolingConfigMap, err := k8sClient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), "nhn-tooling", v1.GetOptions{
		TypeMeta:        v1.TypeMeta{},
		ResourceVersion: "",
	})

	if err != nil {
		return result, fmt.Errorf("could not find config map %s for ror in namespace %s", "nhn-tooling", namespace)
	}

	if nhnToolingConfigMap.Data == nil {
		return result, errors.New("no data in config map for ror")
	}

	environment := nhnToolingConfigMap.Data["environment"]
	toolingVersion := nhnToolingConfigMap.Data["toolingVersion"]
	accessGroupsValue := nhnToolingConfigMap.Data["accessGroups"]
	if accessGroupsValue != "" {
		accessGroups = strings.Split(accessGroupsValue, ";")
	}

	if environment == "" {
		environment = "dev"
	}

	if toolingVersion == "" {
		toolingVersion = MissingConst
	}

	branch := MissingConst
	nhnToolingApp, err := getNhnToolingInfo(dynamicClient)
	if err != nil {
		rlog.Error("could not get nhn-tooling application", err)
	} else {
		branch = nhnToolingApp.Spec.Source.TargetRevision
		if len(nhnToolingApp.Status.Sync.Revision) < 20 {
			toolingVersion = nhnToolingApp.Status.Sync.Revision
		}
	}

	result.Version = toolingVersion
	result.Environment = environment
	result.AccessGroups = accessGroups
	result.Branch = branch

	return result, nil
}

func getNhnToolingInfo(dynamicClient dynamic.Interface) (models.Application, error) {
	result := models.Application{}
	applications, err := dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "argoproj.io",
		Version:  "v1alpha1",
		Resource: "applications",
	}).
		Namespace("argocd").
		Get(context.TODO(), "nhn-tooling", v1.GetOptions{})
	if err != nil {
		rlog.Error("could not get nhn-tooling application", err)
		return result, err
	}

	appByteArray, err := applications.MarshalJSON()
	if err != nil {
		rlog.Error("could not marshal application", err)
		return result, err
	}

	var nhnTooling models.Application
	err = json.Unmarshal(appByteArray, &nhnTooling)
	if err != nil {
		rlog.Error("could not marshal applications", err)
		return result, err
	}

	if nhnTooling.Metadata.Name == "" {
		return result, errors.New("could not find nhn-tooling application")
	}

	return nhnTooling, nil
}

func appendNodeToControlePlane(node *k8smodels.Node, controlPlane *apicontracts.ControlPlane) {
	apiNode := apicontracts.Node{
		Name:                    node.Name,
		Role:                    "control-plane",
		Created:                 node.Created,
		OsImage:                 node.OsImage,
		MachineName:             node.MachineName,
		Architecture:            node.Architecture,
		ContainerRuntimeVersion: node.ContainerRuntimeVersion,
		KernelVersion:           node.KernelVersion,
		KubeProxyVersion:        node.KubeProxyVersion,
		KubeletVersion:          node.KubeletVersion,
		OperatingSystem:         node.OperatingSystem,
		Metrics: apicontracts.Metrics{
			Cpu:            node.Resources.Allocated.Cpu,
			Memory:         node.Resources.Allocated.MemoryBytes,
			CpuConsumed:    node.Resources.Consumed.CpuMilliValue,
			MemoryConsumed: node.Resources.Consumed.MemoryBytes,
		},
	}

	controlPlane.Nodes = append(controlPlane.Nodes, apiNode)

	controlPlane.Metrics.NodeCount = int64(len(controlPlane.Nodes))
	controlPlane.Metrics.Cpu = controlPlane.Metrics.Cpu + apiNode.Metrics.Cpu
	controlPlane.Metrics.Memory = controlPlane.Metrics.Memory + apiNode.Metrics.Memory
	controlPlane.Metrics.CpuConsumed = controlPlane.Metrics.CpuConsumed + apiNode.Metrics.CpuConsumed
	controlPlane.Metrics.MemoryConsumed = controlPlane.Metrics.MemoryConsumed + apiNode.Metrics.MemoryConsumed
}

func appendNodeToNodePools(nodePools *[]apicontracts.NodePool, node *k8smodels.Node) {
	rlog.Debug("", rlog.String("Clustername", node.ClusterName))
	clusterNameSplit := strings.Split(node.ClusterName, "-")
	machineNameSplit := strings.Split(node.MachineName, "-")
	rlog.Debug("", rlog.Strings("machine name split", machineNameSplit))
	var workerName string
	if node.Provider == providers.ProviderTypeTalos {
		workerName = node.Name
	} else if node.Provider != providers.ProviderTypeAks {
		workerName = machineNameSplit[len(clusterNameSplit)]
	} else {
		workerName = machineNameSplit[1]
	}

	rlog.Debug("", rlog.String("worker name", workerName))

	apiNode := apicontracts.Node{
		Role:                    "worker",
		Name:                    node.Name,
		Created:                 node.Created,
		OsImage:                 node.OsImage,
		MachineName:             node.MachineName,
		Architecture:            node.Architecture,
		ContainerRuntimeVersion: node.ContainerRuntimeVersion,
		KernelVersion:           node.KernelVersion,
		KubeProxyVersion:        node.KubeProxyVersion,
		KubeletVersion:          node.KubeletVersion,
		OperatingSystem:         node.OperatingSystem,

		Metrics: apicontracts.Metrics{
			Cpu:            node.Resources.Allocated.Cpu,
			CpuConsumed:    node.Resources.Consumed.CpuMilliValue,
			Memory:         node.Resources.Allocated.MemoryBytes,
			MemoryConsumed: node.Resources.Consumed.MemoryBytes,
		},
	}

	var nodePool *apicontracts.NodePool = nil
	var index int
	for i := 0; i < len(*nodePools); i++ {
		nodepool := (*nodePools)[i]
		if nodepool.Name == workerName {
			index = i
			nodePool = &nodepool
		}
	}

	if nodePool == nil {
		list := []apicontracts.Node{apiNode}
		np := apicontracts.NodePool{
			Name:  workerName,
			Nodes: list,
			Metrics: apicontracts.Metrics{
				NodeCount:      int64(len(list)),
				Cpu:            apiNode.Metrics.Cpu,
				Memory:         apiNode.Metrics.Memory,
				CpuConsumed:    apiNode.Metrics.CpuConsumed,
				MemoryConsumed: apiNode.Metrics.MemoryConsumed,
			},
		}
		*nodePools = append(*nodePools, np)
	} else {
		nodelist := append(nodePool.Nodes, apiNode)
		nodePool.Nodes = nodelist
		nodePool.Metrics.Cpu = nodePool.Metrics.Cpu + apiNode.Metrics.Cpu
		nodePool.Metrics.Memory = nodePool.Metrics.Memory + apiNode.Metrics.Memory
		nodePool.Metrics.CpuConsumed = nodePool.Metrics.CpuConsumed + apiNode.Metrics.CpuConsumed
		nodePool.Metrics.MemoryConsumed = nodePool.Metrics.MemoryConsumed + apiNode.Metrics.MemoryConsumed
		nodePool.Metrics.NodeCount = int64(len(nodelist))
		(*nodePools)[index] = *nodePool
	}
}

func getControlPlaneEndpoint(clientset *kubernetes.Clientset) (string, error) {
	kubeadmConfigMap, err := clientset.CoreV1().ConfigMaps("kube-system").Get(context.TODO(), "kubeadm-config", v1.GetOptions{})
	if err != nil {
		errMsg := "getControlPlaneEndpoint: Could not get cluster config from kube-system/kubeadm-config, check rbac"
		return "", errors.New(errMsg)
	}

	if kubeadmConfigMap == nil {
		errMsg := "getControlPlaneEndpoint: get value 'ControlPlaneEndpoint' from yaml"
		return "", errors.New(errMsg)
	}

	kubeadmClusterConfiguration := kubeadmConfigMap.Data["ClusterConfiguration"]

	var clusterConfigurationValues K8sClusterConfiguration
	err = yaml.Unmarshal([]byte(kubeadmClusterConfiguration), &clusterConfigurationValues)
	if err != nil {
		errMsg := "getControlPlaneEndpoint: Could not parse yaml string to stuct"
		rlog.Error(errMsg, err)
		return "", errors.New(errMsg)
	}

	return clusterConfigurationValues.ControlPlaneEndpoint, nil
}

type K8sClusterConfiguration struct {
	ControlPlaneEndpoint string `yaml:"controlPlaneEndpoint"`
}
type NhnToolingValues struct {
	Cluster NhnToolingCluster `yaml:"cluster"`
	NHN     NHN               `yaml:"nhn"`
}

type NhnToolingCluster struct {
	AccessGroups []string `yaml:"accessGroups"`
}

type NHN struct {
	AccessGroups   []string `yaml:"accessGroups"`
	ToolingVersion string   `yaml:"toolingVersion"`
	Environment    string   `yaml:"environment"`
}
