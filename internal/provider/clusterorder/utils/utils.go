// The utils package is a helper package that provides utility functions for the clusterorder package.
// It should not contain any provider specific logic.
package utils

import (
	"context"
	"crypto/md5" // #nosec G501 - MD5 is used for hash calculation only
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	clustersservice "github.com/NorskHelsenett/ror/cmd/api/services/clustersService"
	projectservice "github.com/NorskHelsenett/ror/cmd/api/services/projectsService"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func NewClusterOrderResource(ctx context.Context, order apiresourcecontracts.ResourceClusterOrderSpec) (apiresourcecontracts.ResourceClusterOrder, error) {
	universalId := GenerateUUID().String()
	apiVersion := fmt.Sprintf("%s/%s", "general.ror.internal", "v1alpha1")
	kind := "ClusterOrder"
	orderres := apiresourcecontracts.ResourceClusterOrder{
		ApiVersion: apiVersion,
		Kind:       kind,
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: universalId,
		},
		Spec: order,
		Status: apiresourcecontracts.ResourceClusterOrderStatus{
			Phase:            "Recieved",
			CreatedTime:      time.Now().Format(time.RFC3339),
			LastObservedTime: time.Now().Format(time.RFC3339),
			UpdatedTime:      time.Now().Format(time.RFC3339),
		},
	}
	return orderres, nil
}

func NewResourceUpdate(ctx context.Context, order apiresourcecontracts.ResourceClusterOrder) (apiresourcecontracts.ResourceUpdateModel, error) {

	resource := apiresourcecontracts.ResourceUpdateModel{
		Owner: apiresourcecontracts.ResourceOwnerReference{
			Scope:   aclmodels.Acl2ScopeRor,
			Subject: string(aclmodels.Acl2RorSubjectGlobal),
		},
		ApiVersion: order.ApiVersion,
		Kind:       order.Kind,
		Uid:        order.Metadata.Uid,
		Action:     apiresourcecontracts.K8sActionAdd,
		Hash:       "",
		Resource:   order,
	}

	hash, err := calculateHash(resource.Resource)
	if err != nil {
		rlog.Errorc(ctx, "error calculating hash", err)
		return apiresourcecontracts.ResourceUpdateModel{}, err
	}
	resource.Hash = hash
	return resource, nil

}

func CreateResource(ctx context.Context, resource apiresourcecontracts.ResourceUpdateModel) error {

	err := resourcesservice.ResourceNewCreateService(ctx, resource)
	if err != nil {
		rlog.Errorc(ctx, "error creating cluster order", err)
		return err
	}

	return nil
}
func UpdateStatus(ctx context.Context, uid string, status apiresourcecontracts.ResourceClusterOrderStatus) error {
	patch := bson.M{}
	if status.Phase != "" {
		patch["resource.status.phase"] = status.Phase
	}
	if status.Status != "" {
		patch["resource.status.status"] = status.Status
	}
	patch["resource.status.updatedtime"] = time.Now().UTC().String()
	rlog.Debug("Patching clusterorder", rlog.Any("patch", patch))
	_, err := resourcesservice.PatchResource(ctx, uid, patch)
	return err
}
func GenerateUUID() uuid.UUID {
	uniqueId, _ := uuid.NewRandom()
	return uniqueId
}

func ValidateOrder(ctx context.Context, order apiresourcecontracts.ResourceClusterOrderSpec) error {

	switch order.OrderType {
	case apiresourcecontracts.ResourceActionTypeCreate:
		err := validateCreateOrder(ctx, order)
		if err != nil {
			return err
		}
	default:
		return errors.New("orderType not supported")
	}

	// TODO: Find a way to diferenciate between different owners
	// TODO: Find a way to allow actions on two clusters with the same name simultaneously
	owner := apiresourcecontracts.ResourceOwnerReference{
		Scope:   aclmodels.Acl2ScopeRor,
		Subject: string(aclmodels.Acl2RorSubjectGlobal),
	}
	clusterOrders, err := resourcesservice.GetClusterorders(ctx, owner)
	if err != nil {
		return err
	}

	for _, clusterOrder := range clusterOrders.Clusterorders {
		specClusterName := strings.ToLower(clusterOrder.Spec.Cluster)
		clusterName := strings.ToLower(order.Cluster)
		if specClusterName == clusterName &&
			!(clusterOrder.Status.Phase == apiresourcecontracts.ResourceClusterOrderStatusPhaseCompleted ||
				clusterOrder.Status.Phase == apiresourcecontracts.ResourceClusterOrderStatusPhaseFailed) {
			return errors.New("clusterOrder with clusterName is already running")
		}
	}

	return nil
}

func validateCreateOrder(ctx context.Context, order apiresourcecontracts.ResourceClusterOrderSpec) error {

	if order.Cluster == "" {
		return errors.New("clusterName is required")
	}
	if order.ProjectId == "" {
		return errors.New("projectId is required")
	}
	if order.OrderBy == "" {
		return errors.New("orderBy is required")
	}
	if apiresourcecontracts.EnvironmentType(order.Environment) == 0 {
		return errors.New("environment is required")
	}
	if apiresourcecontracts.CriticalityLevel(order.Criticality) == 0 {
		return errors.New("criticality is required")
	}
	if apiresourcecontracts.SensitivityLevel(order.Sensitivity) == 0 {
		return errors.New("sensitivity is required")
	}

	if !checkUniqueNodePoolNames(order.NodePools) {
		return errors.New("nodePools names must be unique")
	}
	if len(order.NodePools) == 0 {
		return errors.New("at least one nodePools is required")
	}
	if len(order.OwnerGroup) < 4 {
		return errors.New("ownerGroup is required")
	}

	project, err := projectservice.GetById(ctx, order.ProjectId)
	if err != nil {
		return err
	}

	if project == nil {
		return errors.New("project not found")
	}
	// TODO: Allow two clusters with the same name in same project
	clusterResult, err := clustersservice.GetByFilter(ctx, &apicontracts.Filter{
		Filters: []apicontracts.FilterMetadata{
			{
				Field:     "clustername",
				MatchMode: apicontracts.MatchModeEquals,
				Value:     order.Cluster,
			}, {
				Field:     "projectid",
				MatchMode: apicontracts.MatchModeEquals,
				Value:     order.ProjectId,
			},
		},
	})
	if err != nil {
		return err
	}

	if clusterResult.TotalCount > 0 {
		err := fmt.Errorf("clusterName %s already exists in project %s", order.Cluster, order.ProjectId)
		return err
	}

	return nil

}

func checkUniqueNodePoolNames(pools []apiresourcecontracts.ResourceClusterOrderSpecNodePool) bool {
	check := make(map[string]bool)
	for _, nodePool := range pools {
		if check[nodePool.Name] {
			return false
		}
		check[nodePool.Name] = true
	}
	return true
}

// TODO: Move to a common package, the ResourceFramework package wil be a good candidate
func calculateHash(input any) (string, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	resourceHash := fmt.Sprintf("%x", md5.Sum(bytes)) // #nosec G401 - MD5 is used for hash calculation only
	return resourceHash, nil
}
