package scheduler

import (
	"context"
	"encoding/json"
	"github.com/NorskHelsenett/ror/cmd/agent/clients"
	"github.com/NorskHelsenett/ror/cmd/agent/config"
	"github.com/NorskHelsenett/ror/cmd/agent/services/authservice"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rlog"

	apimachinery "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
)

func MetricsReporting() error {
	k8sClient, err := clients.Kubernetes.GetKubernetesClientset()
	if err != nil {
		return err
	}
	var metricsReport apicontracts.MetricsReport

	metricsReportNodes, err := CreateNodeMetricsList(k8sClient)
	if err != nil {
		rlog.Error("error converting podmetrics", err)
		return err
	}
	ownerref := authservice.CreateOwnerref()

	metricsReport.Owner = apiresourcecontracts.ResourceOwnerReference{
		Scope:   aclmodels.Acl2Scope(ownerref.Scope),
		Subject: string(ownerref.Subject),
	}
	metricsReport.Nodes = metricsReportNodes

	err = sendMetricsToRor(metricsReport)

	return err
}

func sendMetricsToRor(metricsReport apicontracts.MetricsReport) error {
	rorClient, err := clients.GetOrCreateRorClient()
	if err != nil {
		rlog.Error("Could not get ror-api client", err)
		config.IncreaseErrorCount()
		return err
	}

	url := "/v1/metrics"
	response, err := rorClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(metricsReport).
		Post(url)
	if err != nil {
		rlog.Error("Could not send metrics data to ror-api", err)
		config.IncreaseErrorCount()
		return err
	}

	if response == nil {
		rlog.Error("Response is nil", err)
		config.IncreaseErrorCount()
		return err
	}

	if !response.IsSuccess() {
		config.IncreaseErrorCount()
		rlog.Error("Got unsuccessful status code from ror-api", err,
			rlog.Int("status code", response.StatusCode()),
			rlog.Int("error count", config.ErrorCount))
		return err
	} else {
		config.ResetErrorCount()
		rlog.Info("Metrics report sent to ror")

		byteReport, err := json.Marshal(metricsReport)
		if err == nil {
			rlog.Debug("", rlog.String("byte report", string(byteReport)))
		}
	}
	return nil
}

func CreateNodeMetricsList(k8sClient *kubernetes.Clientset) ([]apicontracts.NodeMetric, error) {
	var nodeMetricsList apicontracts.NodeMetricsList
	var metricsReportNodes []apicontracts.NodeMetric

	data, err := k8sClient.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.TODO())
	if err != nil {
		rlog.Error("error converting podmetrics", err)
		return metricsReportNodes, err
	}

	err = json.Unmarshal(data, &nodeMetricsList)
	if err != nil {
		rlog.Error("error unmarshaling podmetrics", err)
		return metricsReportNodes, err
	}

	for _, node := range nodeMetricsList.Items {

		metricsReportNode, err := CreateNodeMetrics(node)
		if err != nil {
			rlog.Error("error converting podmetrics", err)
			return metricsReportNodes, err
		}
		metricsReportNodes = append(metricsReportNodes, metricsReportNode)
	}

	return metricsReportNodes, nil

}

func CreateNodeMetrics(node apicontracts.NodeMetricsListItem) (apicontracts.NodeMetric, error) {
	var nodeMetric apicontracts.NodeMetric
	var timestamp time.Time = node.Timestamp

	nodeCpuRaw, err := apimachinery.ParseQuantity(node.Usage.CPU)
	if err != nil {
		rlog.Error("error converting nodemetrics", err)
		return nodeMetric, err
	}
	nodeCpu := nodeCpuRaw.MilliValue()

	nodeMemoryRaw, err := apimachinery.ParseQuantity(node.Usage.Memory)
	if err != nil {
		rlog.Error("error converting nodemetrics", err)
		return nodeMetric, err
	}
	nodeMemory, _ := nodeMemoryRaw.AsInt64()

	nodeMetric = apicontracts.NodeMetric{
		Name:        node.Metadata.Name,
		TimeStamp:   timestamp,
		CpuUsage:    nodeCpu,
		MemoryUsage: nodeMemory,
	}
	return nodeMetric, nil
}
func CreatePodMetricsList(k8sClient *kubernetes.Clientset) ([]apicontracts.PodMetric, error) {
	var podMetricsList apicontracts.PodMetricsList
	var metricsReportPods []apicontracts.PodMetric

	data, err := k8sClient.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw(context.TODO())
	if err != nil {
		rlog.Error("error unmarshaling podmetrics", err)
		return metricsReportPods, err
	}

	err = json.Unmarshal(data, &podMetricsList)
	if err != nil {
		rlog.Error("error unmarshaling podmetrics", err)
		return metricsReportPods, err
	}

	for _, pod := range podMetricsList.Items {
		metricsReportPod, err := CreatePodMetrics(pod)
		if err != nil {
			rlog.Error("error converting podmetrics", err)
			return metricsReportPods, err
		}
		metricsReportPods = append(metricsReportPods, metricsReportPod)
	}
	return metricsReportPods, nil
}

func CreatePodMetrics(pod apicontracts.PodMetricsListItem) (apicontracts.PodMetric, error) {
	var podMetric apicontracts.PodMetric
	var timestamp time.Time = pod.Timestamp
	var podCpuSum int64 = 0
	var podMemorySum int64 = 0

	for _, container := range pod.Containers {
		podCpu, err := apimachinery.ParseQuantity(container.Usage.CPU)
		if err != nil {
			rlog.Error("error converting podmetrics", err)
			return podMetric, err
		}
		podCpuSum = podCpuSum + podCpu.MilliValue()
		podMemoryObj, err := apimachinery.ParseQuantity(container.Usage.Memory)
		if err != nil {
			rlog.Error("error converting podmetrics", err)
			return podMetric, err
		}
		podMemory, _ := podMemoryObj.AsInt64()
		podMemorySum = podMemorySum + podMemory
	}
	podMetric = apicontracts.PodMetric{
		Name:        pod.Metadata.Name,
		Namespace:   pod.Metadata.Namespace,
		TimeStamp:   timestamp,
		CpuUsage:    podCpuSum,
		MemoryUsage: podMemorySum,
	}
	return podMetric, nil
}
