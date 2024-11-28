package projectsservice

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	"github.com/NorskHelsenett/ror/internal/models"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	clustersservice "github.com/NorskHelsenett/ror/cmd/api-stub/services/clustersService"

	clustersRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/clustersRepo"
	projectsRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/projectsRepo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[apicontracts.Project], error) {
	projects, totalCount, err := projectsRepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error when getting projects by filter from repo: %v", err)
	}

	var results []apicontracts.Project
	for _, v := range projects {
		var project apicontracts.Project
		err := mapping.Map(v, &project)
		if err != nil {
			return nil, fmt.Errorf("could not map from mongotype to apitype: %v", err)
		}
		results = append(results, project)
	}

	paginatedResult := apicontracts.PaginatedResult[apicontracts.Project]{}

	paginatedResult.Data = results
	paginatedResult.DataCount = int64(len(results))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalCount)

	return &paginatedResult, nil
}

func GetById(ctx context.Context, ID string) (*apicontracts.Project, error) {
	object, err := projectsRepo.GetById(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("could not get object by ID from repository: %v", err)
	}

	var mapped apicontracts.Project
	err = mapping.Map(object, &mapped)
	if err != nil {
		return nil, fmt.Errorf("could not map object: %v", err)
	}

	return &mapped, nil
}

func GetClustersByProjectId(ctx context.Context, projectId string) ([]*apicontracts.ClusterInfo, error) {
	objects, err := clustersRepo.GetClusterIdByProjectId(ctx, projectId)
	if err != nil {
		return nil, fmt.Errorf("could not get object by ID from repository: %v", err)
	}

	var results []*apicontracts.ClusterInfo
	for _, v := range objects {
		var clusterInfo apicontracts.ClusterInfo
		err := mapping.Map(v, &clusterInfo)
		if err != nil {
			return nil, fmt.Errorf("could not map from mongotype to apitype: %v", err)
		}
		results = append(results, &clusterInfo)
	}

	return results, nil
}

func Create(ctx context.Context, projectInput *apicontracts.ProjectModel) (*apicontracts.Project, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	var mappedInput mongoTypes.MongoProject
	err := mapping.Map(projectInput, &mappedInput)
	if err != nil {
		return nil, fmt.Errorf("could not map project from apitype to mongotype: %v", err)
	}

	createdProject, err := projectsRepo.Create(ctx, &mappedInput)
	if err != nil {
		return nil, fmt.Errorf("could not create project: %v", err)
	}

	var mappedResult apicontracts.Project
	err = mapping.Map(createdProject, &mappedResult)
	if err != nil {
		return nil, fmt.Errorf("could not map project from mongotype to apitype: %v", err)
	}

	_, err = auditlog.Create(ctx, "Project created", models.AuditCategoryProject, models.AuditActionCreate, identity.User, createdProject, nil)
	if err != nil {
		return nil, fmt.Errorf("could not audit log delete action: %v", err)
	}

	return &mappedResult, nil
}

func Update(ctx context.Context, projectId string, input *apicontracts.ProjectModel) (*apicontracts.Project, *apicontracts.Project, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	var mongoObject mongoTypes.MongoProject
	err := mapping.Map(input, &mongoObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map from apitype to mongotype: %v", err)
	}

	mongoObject.ID = primitive.NilObjectID
	newObject, oldObject, err := projectsRepo.Update(ctx, mongoObject, projectId)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update object: %v", err)
	}

	var mappedNewObject apicontracts.Project
	err = mapping.Map(newObject, &mappedNewObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map updated object from apitype to mongotype: %v", err)
	}

	var mappedOldObject apicontracts.Project
	err = mapping.Map(oldObject, &mappedOldObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map original object from apitype to mongotype: %v", err)
	}

	_, err = auditlog.Create(ctx, "Project updated", models.AuditCategoryProject, models.AuditActionUpdate, identity.User, newObject, oldObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not audit log delete action: %v", err)
	}

	return &mappedNewObject, &mappedOldObject, nil
}

func Delete(ctx context.Context, projectId string) (bool, *apicontracts.Project, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	clustersUsingProject, err := clustersservice.GetClusterIdByProjectId(ctx, projectId)
	if err != nil {
		return false, nil, fmt.Errorf("could not delete object, could not check if project is used by others: %v", err)
	}

	if len(clustersUsingProject) > 0 {
		return false, nil, fmt.Errorf("could not delete project, it is in use")
	}

	deleted, deletedObject, err := projectsRepo.Delete(ctx, projectId)
	if err != nil {
		return false, nil, fmt.Errorf("could not delete object: %v", err)
	}

	_, err = auditlog.Create(ctx, "Project deleted", models.AuditCategoryProject, models.AuditActionDelete, identity.User, deleted, nil)
	if err != nil {
		return false, nil, fmt.Errorf("could not audit log delete action: %v", err)
	}

	var mappedDeletedObject apicontracts.Project
	err = mapping.Map(deletedObject, &mappedDeletedObject)
	if err != nil {
		return false, nil, fmt.Errorf("could not map deleted object from mongotype to apitype: %v", err)
	}

	return deleted, &mappedDeletedObject, nil
}
