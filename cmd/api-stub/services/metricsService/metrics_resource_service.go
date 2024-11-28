package metricsservice

import (
	"context"
	"strconv"

	resourcesservice "github.com/NorskHelsenett/ror/cmd/api-stub/services/resourcesService"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	apimachinery "k8s.io/apimachinery/pkg/api/resource"

	metricsRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/metricsRepo"
)

// Enriches metricsreports and calls metricsRepo.WriteMetrics to save the data to mongodb
func ProcessMetricReport(ctx context.Context, resourceUpdate *apicontracts.MetricsReport) error {
	ownerref := apiresourcecontracts.ResourceOwnerReference{
		Scope:   resourceUpdate.Owner.Scope,
		Subject: string(resourceUpdate.Owner.Subject),
	}
	nodes, _ := resourcesservice.GetNodes(ctx, ownerref)

	for m, resource := range resourceUpdate.Nodes {
		node := nodes.GetByName(resource.Name)
		if len(node.Metadata.Name) > 0 {
			cpuAllocated, _ := strconv.ParseInt(node.Status.Capacity.Cpu, 10, 64)
			resourceUpdate.Nodes[m].CpuAllocated = cpuAllocated
			cpuPercentage := (float64(resource.CpuUsage) / (float64(cpuAllocated) * 1000)) * 100
			resourceUpdate.Nodes[m].CpuPercentage = cpuPercentage
			memoryAllocatedQuantity, _ := apimachinery.ParseQuantity(node.Status.Capacity.Memory)
			memoryAllocated, _ := memoryAllocatedQuantity.AsInt64()
			resourceUpdate.Nodes[m].MemoryAllocated = memoryAllocated
			memoryPercentage := (float64(resource.MemoryUsage) / float64(memoryAllocated)) * 100
			resourceUpdate.Nodes[m].MemoryPercentage = memoryPercentage
		}
	}
	metricsRepo.WriteMetrics(resourceUpdate, string(resourceUpdate.Owner.Subject), ctx)
	return nil
}
