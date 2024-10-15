package applicationservice

import (
	"context"
	"errors"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	clustersrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/clustersRepo"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func ProcessResourceUpdatedEvent(ctx context.Context, event apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]) error {
	if event.Resource.Metadata.Name != "nhn-tooling" {
		return nil
	}

	application, err := getResource(ctx, event)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return errors.New("could not get resource")
	}

	err = updateCluster(ctx, event.Owner.Subject, application)
	if err != nil {
		rlog.Errorc(ctx, "could not update cluster", err)
		return errors.New("could not update cluster")
	}

	return nil
}

func getResource(ctx context.Context, event apiresourcecontracts.ResourceModel[apiresourcecontracts.ResourceApplication]) (apiresourcecontracts.ResourceApplication, error) {
	query := apiresourcecontracts.ResourceQuery{
		Owner:      event.Owner,
		ApiVersion: event.Resource.ApiVersion,
		Kind:       event.Resource.Kind,
		Uid:        event.Resource.Metadata.Uid,
		Internal:   true,
	}

	resource, err := resourcesservice.GetResource[apiresourcecontracts.ResourceApplication](ctx, query)
	if err != nil {
		rlog.Errorc(ctx, "could not get resource", err)
		return apiresourcecontracts.ResourceApplication{}, errors.New("could not get resource")
	}

	return resource, nil
}

func updateCluster(ctx context.Context, clusterId string, application apiresourcecontracts.ResourceApplication) error {
	existingCluster, err := clustersrepo.GetByClusterId(ctx, clusterId)
	if err != nil {
		rlog.Errorc(ctx, "could not get cluster", err)
		return errors.New("could not get cluster")
	}

	if existingCluster == nil {
		rlog.Errorc(ctx, "could not find cluster", err, rlog.String("clusterId", clusterId))
		return nil
	}

	updatedCluster := existingCluster

	if application.Spec.SyncPolicy.Automated == nil {
		updateConditionsRelatedToToolingSync(updatedCluster, apicontracts.ConditionStatusFalse, "Missing NHN-tooling autosync")
	} else {
		updateConditionsRelatedToToolingSync(updatedCluster, apicontracts.ConditionStatusTrue, "NHN-tooling autosync active")
	}

	err = clustersrepo.Update(ctx, updatedCluster)
	if err != nil {
		rlog.Errorc(ctx, "could not update cluster", err)
		return errors.New("could not update cluster")
	}

	rlog.Debugc(ctx, "cluster updated")

	return nil
}

func updateConditionsRelatedToToolingSync(cluster *apicontracts.Cluster, status apicontracts.ConditionStatus, message string) {
	condition := apicontracts.ClusterCondition{
		Type:    apicontracts.ConditionTypeToolingOk,
		Status:  status,
		Message: message,
		Created: time.Now(),
		Updated: time.Now(),
	}

	if cluster.Status.Conditions == nil || len(cluster.Status.Conditions) == 0 {
		cluster.Status.Conditions = append(cluster.Status.Conditions, condition)
		return
	}

	for index, c := range cluster.Status.Conditions {
		if c.Type == apicontracts.ConditionTypeToolingOk {
			c.Status = status
			c.Message = message
			c.Updated = time.Now()
			cluster.Status.Conditions[index] = c
			break
		}
	}
}
