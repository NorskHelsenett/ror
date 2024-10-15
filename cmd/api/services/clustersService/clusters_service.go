// Package clustersservice cluster package provides services to get and manipulate the cluster object
package clustersservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/cmd/api/apiconnections"
	"github.com/NorskHelsenett/ror/cmd/api/services"
	"github.com/NorskHelsenett/ror/cmd/api/webserver/sse"
	"github.com/NorskHelsenett/ror/internal/services/clusterservice"
	"github.com/NorskHelsenett/ror/internal/services/kubeconfigservice"
	"time"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/models"
	clustersRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/clustersRepo"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/cmd/api/models/ssemodels"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/messagebuscontracts"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetByClusterId(ctx context.Context, clusterId string) (*apicontracts.Cluster, error) {
	result, err := clustersRepo.GetByClusterId(ctx, clusterId)
	if err != nil {
		return nil, errors.New("could not get clusters")
	}

	return result, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[*apicontracts.Cluster], error) {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "clustersservice.GetByFilter")
	defer span.End()

	_, span1 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "getMongodb")
	defer span1.End()

	span1.End()

	_, span2 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "clustersRepo.GetByFilter")
	defer span2.End()

	result, err := clustersRepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not get clusters: %v", err)
	}

	span2.End()

	_, span4 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "fillClusterObject")
	defer span4.End()

	var data []*apicontracts.Cluster
	data = append(data, result.Data...)

	span4.End()
	result.Data = data

	return result, nil
}

func FindByName(ctx context.Context, clusterName string) (*apicontracts.Cluster, error) {
	result, err := clustersRepo.FindByName(ctx, clusterName)
	if err != nil {
		return nil, errors.New("could not get cluster")
	}

	return result, nil
}

func GetByWorkspace(ctx context.Context, filter *apicontracts.Filter, workspaceName string) (*apicontracts.PaginatedResult[apicontracts.Cluster], error) {
	result, err := clustersRepo.GetByWorkspace(ctx, filter, workspaceName)
	if err != nil {
		return nil, errors.New("could not get clusters")
	}

	return result, nil
}

func CreateOrUpdate(ctx context.Context, input *apicontracts.Cluster, clusterId string) error {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Service: clustersservice.CreateOrUpdate")
	defer span.End()
	_, span1 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Check if cluster exists")
	defer span1.End()
	existing, err := clustersRepo.GetByClusterId(ctx, clusterId)
	if err != nil {
		return fmt.Errorf("could not create or update cluster with id: %s", input.ClusterId)
	}
	span1.End()

	CpuMemPercentageCalc(input)
	FindMachineClass(ctx, input)

	if existing != nil {
		_, span1 = otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Run service update")
		defer span1.End()

		err := Update(ctx, input, existing)

		span1.End()
		return err
	} else {
		_, span1 = otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Run service create")
		defer span1.End()

		_, err := Create(ctx, input)
		span1.End()
		return err
	}
}

func Create(ctx context.Context, input *apicontracts.Cluster) (string, error) {
	clusterId, err := clusterservice.Create(ctx, input.ClusterName, input.Workspace.DatacenterID, input.WorkspaceId, input.Workspace.Name, input.Metadata.ProjectID)
	if err != nil {
		rlog.Errorc(ctx, "could not create cluster", err, rlog.String("clusterName", input.ClusterName))
		return "", fmt.Errorf("could not create cluster with id: %s", input.ClusterId)
	}

	event := messagebuscontracts.ClusterCreatedEvent{}
	event.ClusterId = clusterId
	event.WorkspaceName = input.Workspace.Name
	event.ClusterName = input.ClusterName
	err = apiconnections.RabbitMQConnection.SendMessage(ctx, event, messagebuscontracts.Route_Cluster_Created, nil)
	if err != nil {
		rlog.Errorc(ctx, "could not send cluster created event", err, rlog.String("clusterId", clusterId))
	}

	sse.Server.BroadcastMessage(ssemodels.SseMessage{SSEBase: ssemodels.SSEBase{Event: ssemodels.SseType_Cluster_Created}, Data: event})
	return clusterId, nil
}

func Update(ctx context.Context, input *apicontracts.Cluster, existing *apicontracts.Cluster) error {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Service: clustersservice.Update")
	defer span.End()
	if existing == nil {
		return fmt.Errorf("could not find existing update cluster with id: %s", input.ClusterId)
	}

	input.LastObserved = time.Now()
	input.FirstObserved = existing.FirstObserved
	input.Metadata = existing.Metadata
	input.Config = existing.Config
	input.WorkspaceId = existing.WorkspaceId
	input.Status = existing.Status
	input.Identifier = existing.Identifier
	if len(input.Identifier) == 0 {
		input.Identifier = clusterservice.GetClusterIdentifier(input.ClusterName)
	}

	_, span1 := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "Run repository update")
	defer span1.End()
	err := clustersRepo.Update(ctx, input)
	if err != nil {
		return fmt.Errorf("could not update cluster with id: %s", input.ClusterId)
	}
	span1.End()
	span.End()
	return nil
}

func Exists(ctx context.Context, clusterId string) (bool, error) {
	cluster, err := clustersRepo.FindByClusterId(ctx, clusterId)
	if err != nil {
		return false, err
	}

	if cluster == nil || cluster.ClusterName == "" {
		return false, nil
	}

	return true, nil
}

func GetClusterIdByProjectId(ctx context.Context, projectId string) ([]*apicontracts.ClusterInfo, error) {
	clusters, err := clustersRepo.GetClusterIdByProjectId(ctx, projectId)
	if err != nil {
		return make([]*apicontracts.ClusterInfo, 0), err
	}

	return clusters, nil
}

func CpuMemPercentageCalc(cluster *apicontracts.Cluster) {
	cluster.Metrics.CpuPercentage = services.CpuPercentage(cluster.Metrics.Cpu, cluster.Metrics.CpuConsumed)
	cluster.Metrics.MemoryPercentage = services.MemoryPercentage(cluster.Metrics.Memory, cluster.Metrics.MemoryConsumed)

	cluster.Topology.ControlPlane.Metrics.CpuPercentage = services.CpuPercentage(cluster.Topology.ControlPlane.Metrics.Cpu, cluster.Topology.ControlPlane.Metrics.CpuConsumed)
	cluster.Topology.ControlPlane.Metrics.MemoryPercentage = services.MemoryPercentage(cluster.Topology.ControlPlane.Metrics.Memory, cluster.Topology.ControlPlane.Metrics.MemoryConsumed)
	controlPlaneNodesLength := len(cluster.Topology.ControlPlane.Nodes)

	for i := 0; i < controlPlaneNodesLength; i++ {
		node := cluster.Topology.ControlPlane.Nodes[i]
		(&cluster.Topology.ControlPlane.Nodes[i]).Metrics.CpuPercentage = services.CpuPercentage(node.Metrics.Cpu, node.Metrics.CpuConsumed)
		(&cluster.Topology.ControlPlane.Nodes[i]).Metrics.MemoryPercentage = services.MemoryPercentage(node.Metrics.Memory, node.Metrics.MemoryConsumed)
	}

	lengthNodePools := len(cluster.Topology.NodePools)
	for j := 0; j < lengthNodePools; j++ {
		nodePool := cluster.Topology.NodePools[j]
		(&cluster.Topology.NodePools[j]).Metrics.CpuPercentage = services.CpuPercentage(nodePool.Metrics.Cpu, nodePool.Metrics.CpuConsumed)
		(&cluster.Topology.NodePools[j]).Metrics.MemoryPercentage = services.MemoryPercentage(nodePool.Metrics.Memory, nodePool.Metrics.MemoryConsumed)

		lengthNodePoolNodes := len(nodePool.Nodes)
		for k := 0; k < lengthNodePoolNodes; k++ {
			node := nodePool.Nodes[k]
			(&nodePool.Nodes[k]).Metrics.CpuPercentage = services.CpuPercentage(node.Metrics.Cpu, node.Metrics.CpuConsumed)
			(&nodePool.Nodes[k]).Metrics.MemoryPercentage = services.MemoryPercentage(node.Metrics.Memory, node.Metrics.MemoryConsumed)
		}
	}
}

func FindMachineClass(ctx context.Context, cluster *apicontracts.Cluster) {
	db := mongodb.GetMongoDb()
	collectionName := "prices"
	mongoCollection := db.Collection(collectionName)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	var prices = make([]apicontracts.Price, 0)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "cpu", Value: 1}, {Key: "memory", Value: 1}, {Key: "machineclass", Value: 1}, {Key: "id", Value: 1}})
	results, err := mongoCollection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		rlog.Errorc(ctx, "could not get price collection", err)
		return
	}

	//reading from the db in an optimal way
	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)
	if results.RemainingBatchLength() == 0 {
		rlog.Errorc(ctx, "", fmt.Errorf("no prices"))
		return
	}
	for results.Next(ctx) {
		var price apicontracts.Price
		if err = results.Decode(&price); err != nil {
			rlog.Errorc(ctx, "could not decode price", err)
		}

		prices = append(prices, price)
	}

	months := int64(12)
	provider := cluster.Workspace.Datacenter.Provider

	var aggPrice int64
	lengthNodePools := len(cluster.Topology.NodePools)
	for j := 0; j < lengthNodePools; j++ {
		nodePool := cluster.Topology.NodePools[j]
		lengthNodePoolNodes := len(nodePool.Nodes)
		var nodePoolPrice int64
		var machineClassName string
		for k := 0; k < lengthNodePoolNodes; k++ {
			node := nodePool.Nodes[k]
			var price int64
			price, machineClassName = services.FindMachineClass(node.Metrics.Memory, node.Metrics.Cpu, provider, prices)
			nodePool.Nodes[k].Metrics.PriceMonth = price
			nodePool.Nodes[k].Metrics.PriceYear = price * months
			nodePool.Nodes[k].MachineClass = machineClassName
			nodePoolPrice = price + nodePoolPrice
		}

		cluster.Topology.NodePools[j].Metrics.PriceMonth = nodePoolPrice
		cluster.Topology.NodePools[j].Metrics.PriceYear = nodePoolPrice * months
		cluster.Topology.NodePools[j].MachineClass = machineClassName
		aggPrice = aggPrice + nodePoolPrice
	}
	cluster.Metrics.PriceMonth = aggPrice
	cluster.Metrics.PriceYear = aggPrice * months
}

func GetMetadata(ctx context.Context) (map[string][]string, error) {
	metadata, err := clustersRepo.GetMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get metadata from cluster repository: %v", err)
	}
	return metadata, nil
}

func UpdateMetadata(ctx context.Context, input *apicontracts.ClusterMetadataModel, existing *apicontracts.Cluster) error {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	if !identity.IsUser() {
		return errors.New("must be a user to update")
	}

	err := clustersRepo.UpdateMetadata(ctx, input, existing)
	if err != nil {
		return fmt.Errorf("could not update cluster with id: %s", existing.ClusterId)
	}

	_, err = auditlog.Create(ctx, "New taskcollection deleted", models.AuditCategoryClusterMetadata, models.AuditActionUpdate, identity.User, input, existing.Metadata)
	if err != nil {
		return fmt.Errorf("could not audit log delete action: %v", err)
	}

	return nil
}

func GetControlPlanesMetadata(ctx context.Context) ([]apicontracts.ClusterControlPlaneMetadata, error) {
	controlPlanes, err := clustersRepo.GetControlPlaneMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get control planes from cluster repository: %v", err)
	}
	return controlPlanes, nil
}

func GetKubeconfig(ctx context.Context, clusterId string, credentials apicontracts.KubeconfigCredentials) (string, error) {
	if credentials.Username == "" || credentials.Password == "" {
		err := errors.New("username and password must be provided")
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("clusterId", clusterId))
		return "", err
	}

	if clusterId == "" {
		err := errors.New("clusterId must be provided")
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("clusterId", clusterId))
		return "", err
	}

	cluster, err := clustersRepo.GetByClusterId(ctx, clusterId)
	if err != nil {
		err := fmt.Errorf("could not find cluster with id: %s", clusterId)
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("clusterId", clusterId))
		return "", err
	}

	if cluster == nil {
		err := fmt.Errorf("could not find cluster with id: %s", clusterId)
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("clusterId", clusterId))
		return "", err
	}

	now := time.Now()
	kubeconfigString, err := kubeconfigservice.GetKubeconfig(cluster, credentials)
	if err != nil {
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("clusterId", clusterId))
		return "", err
	}
	end := time.Now()
	duration := end.Sub(now)
	rlog.Infoc(ctx, "kubeconfig fetched", rlog.String("clusterId", clusterId), rlog.String("duration", fmt.Sprintf("%s", duration)))

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	_, err = auditlog.Create(ctx, "Identity fetching kubeconfig for workspace",
		models.AuditCategoryKubeconfig,
		models.AuditActionRead,
		identity.User,
		fmt.Sprintf("identity type: '%s', id: '%s' fetching kubeconfig for clusterId: %s", identity.Type, identity.GetId(), clusterId),
		nil)
	if err != nil {
		return "", fmt.Errorf("could not audit log fetch action: %v", err)
	}

	return kubeconfigString, nil
}
