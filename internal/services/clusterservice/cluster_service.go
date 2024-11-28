package clusterservice

import (
	"context"
	"fmt"
	"strings"
	"time"

	resourcesservice "github.com/NorskHelsenett/ror/cmd/api-stub/services/resourcesService"

	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	clustersRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/clustersRepo"
	datacentersRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/datacentersRepo"
	projectsrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/projectsRepo"
	workspacesRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/workspacesRepo"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// Create creates a new cluster
//
// Parameters:
//
// - clusterName: name of the cluster
//
// - workspaceId: id of the workspace
//
// - workspaceName: name of the workspace, or what the name of the new workspace should be
//
// - datacenterId: id of the datacenter
func Create(ctx context.Context, clusterName, datacenterId, workspaceId, workspaceName, projectId string) (string, error) {
	if datacenterId == "" {
		return "", fmt.Errorf("datacenterId must be set")
	}

	clusterInput := apicontracts.Cluster{
		ClusterName: clusterName,
		Metadata: apicontracts.ClusterMetadata{
			ProjectID: projectId,
		},
	}
	datacenter, _ := datacentersRepo.GetById(ctx, datacenterId)
	if datacenter != nil {
		clusterInput.Workspace.DatacenterID = datacenter.ID
	} else {
		return "", fmt.Errorf("could not find datacenter with id: %s", datacenterId)
	}

	workspace, _ := workspacesRepo.FindByName(ctx, workspaceName)
	if workspace != nil {
		clusterInput.Workspace = *workspace
		clusterInput.WorkspaceId = workspace.ID
	} else {
		w, err := workspacesRepo.Create(ctx, &apicontracts.Workspace{
			Name:         workspaceName,
			DatacenterID: datacenter.ID,
		})
		if err != nil {
			rlog.Errorc(ctx, "could not create workspace", err, rlog.String("name", workspaceName))
			return "", fmt.Errorf("could not create workspace with name: %s", workspaceName)
		}
		clusterInput.Workspace.ID = w.ID
		clusterInput.WorkspaceId = w.ID
	}

	now := time.Now()
	clusterInput.FirstObserved = now
	clusterInput.LastObserved = now
	clusterInput.Identifier = GetClusterIdentifier(clusterName)
	clusterInput.ClusterId = clusterInput.Identifier

	err := clustersRepo.Create(ctx, &clusterInput)
	if err != nil {
		return "", fmt.Errorf("could not create cluster with id: %s", clusterInput.ClusterId)
	}

	err = postSetupCluster(ctx, clusterInput.ClusterId)
	if err != nil {
		rlog.Warnc(ctx, "error post setup cluster", rlog.Any("error", err))
	}

	return clusterInput.ClusterId, nil
}

// GetClusterIdentifier returns a cluster identifier
//
// Parameters:
//
// - clusterName: name of the cluster
func GetClusterIdentifier(clusterName string) string {
	idpostfix := stringhelper.RandomString(4, stringhelper.StringTypeClusterId)
	identifier := fmt.Sprintf("%s-%s", clusterName, idpostfix)
	return identifier
}

// postSetupCluster sets up the cluster after it has been created
func postSetupCluster(ctx context.Context, clusterId string) error {
	if clusterId == "" {
		return fmt.Errorf("clusterId must be set")
	}

	cluster, err := clustersRepo.FindByClusterId(ctx, clusterId)
	if err != nil {
		rlog.Errorc(ctx, "could not find cluster", err, rlog.String("clusterId", clusterId))
		return fmt.Errorf("could not find cluster with id: %s", clusterId)
	}

	// TODO: need to be filtered
	clusterOrderOwner := apiresourcecontracts.ResourceOwnerReference{
		Scope:   aclmodels.Acl2ScopeRor,
		Subject: string(aclmodels.Acl2RorSubjectGlobal),
	}
	clusterOrders, err := resourcesservice.GetClusterorders(ctx, clusterOrderOwner)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return err
	}

	// TODO - this is a hack, we need to find a better way to get the cluster order
	var clusterOrder apiresourcecontracts.ResourceClusterOrder
	found := false
	for _, co := range clusterOrders.Clusterorders {
		if strings.Contains(cluster.ClusterName, co.Spec.Cluster) {
			clusterOrder = co
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("could not find cluster order for cluster: %s", cluster.ClusterName)
	}

	mongoProject, err := projectsrepo.GetById(ctx, clusterOrder.Spec.ProjectId)
	if err != nil {
		rlog.Errorc(ctx, "error getting project", err)
		return err
	}

	var project apicontracts.Project
	err = mapping.Map(project, &mongoProject)
	if err != nil {
		return fmt.Errorf("could not map object: %v", err)
	}

	switch clusterOrder.Spec.Environment {
	case apiresourcecontracts.EnvironmentDevelopment:
		cluster.Environment = "dev"
	case apiresourcecontracts.EnvironmentQA:
		cluster.Environment = "qa"
	case apiresourcecontracts.EnvironmentTesting:
		cluster.Environment = "test"
	case apiresourcecontracts.EnvironmentProduction:
		cluster.Environment = "prod"
	}

	groups := make([]string, 0)
	groups = append(groups, clusterOrder.Spec.OwnerGroup)
	cluster.ACL.AccessGroups = groups
	cluster.Updated = time.Now()

	err = clustersRepo.Update(ctx, cluster)
	if err != nil {
		rlog.Errorc(ctx, "error updating cluster", err)
		return err
	}

	metadata := apicontracts.ClusterMetadataModel{
		ProjectID:   clusterOrder.Spec.ProjectId,
		Criticality: apicontracts.CriticalityLevel(clusterOrder.Spec.Criticality),
		Sensitivity: apicontracts.SensitivityLevel(clusterOrder.Spec.Sensitivity),
		ServiceTags: clusterOrder.Spec.ServiceTags,
		Description: project.Description,
		Billing: apicontracts.BillingModel{
			Workorder: project.ProjectMetadata.Billing.Workorder,
		},
		Roles: project.ProjectMetadata.Roles,
	}

	err = clustersRepo.UpdateMetadata(ctx, &metadata, cluster)
	if err != nil {
		rlog.Errorc(ctx, "error updating cluster metadata", err)
		return err
	}

	return nil
}
