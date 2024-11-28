package workspacesservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	"github.com/NorskHelsenett/ror/internal/models"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"
	workspacesRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/workspacesRepo"
	"github.com/NorskHelsenett/ror/internal/services/kubeconfigservice"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAll(ctx context.Context) (*[]apicontracts.Workspace, error) {
	workspaces, err := workspacesRepo.GetAllByIdentity(ctx)
	if err != nil {
		return nil, errors.New("could not fetch workspaces")
	}

	return workspaces, nil
}

func GetByFilter(ctx context.Context, filter apicontracts.Filter) ([]*apicontracts.PaginatedResult[apicontracts.Workspace], error) {
	return nil, nil
}

func GetById(ctx context.Context, id string) (*apicontracts.Workspace, error) {
	object, err := workspacesRepo.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get object by ID from repository: %w", err)
	}

	if object == nil {
		return nil, nil
	}

	var mapped apicontracts.Workspace
	err = mapping.Map(object, &mapped)
	if err != nil {
		return nil, fmt.Errorf("could not map object: %w", err)
	}

	return &mapped, nil
}

func GetByName(ctx context.Context, workspaceName string) (*apicontracts.Workspace, error) {
	workspace, err := workspacesRepo.GetByName(ctx, workspaceName)
	if err != nil {
		return nil, errors.New("could not fetch workspace")
	}

	return workspace, nil
}

func Update(ctx context.Context, input *apicontracts.Workspace, id string) (*apicontracts.Workspace, *apicontracts.Workspace, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	var mappedInput mongoTypes.MongoWorkspace
	err := mapping.Map(input, &mappedInput)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map from api type to mongo type, %w", err)
	}

	if input.DatacenterID != "" {
		datacenterID, err := primitive.ObjectIDFromHex(input.DatacenterID)
		if err != nil {
			return nil, nil, fmt.Errorf("could not get datacenterID from input: %w", err)
		}
		mappedInput.DatacenterID = datacenterID
	}

	// if input.ProjectID != "" {
	// 	projectID, err := primitive.ObjectIDFromHex(input.ProjectID)
	// 	if err != nil {
	// 		return nil, nil, fmt.Errorf("could not get projectID from input: %w", err)
	// 	}
	// 	mappedInput.ProjectID = projectID
	// }

	updatedObject, originalObject, err := workspacesRepo.Update(ctx, &mappedInput, id)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update object: %w", err)
	}

	_, err = auditlog.Create(ctx, "New task created", models.AuditCategoryWorkspace, models.AuditActionUpdate, identity.User, updatedObject, originalObject)
	if err != nil {
		rlog.Error("failed to create auditlog", err)
	}

	var mappedUpdatedObject apicontracts.Workspace
	err = mapping.Map(updatedObject, &mappedUpdatedObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map updated object from api type to mongo type: %w", err)
	}

	var mappedOriginalObject apicontracts.Workspace
	err = mapping.Map(originalObject, &mappedOriginalObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map original object from api type to mongo type: %w", err)
	}

	return &mappedUpdatedObject, &mappedOriginalObject, nil
}

func GetKubeconfig(ctx context.Context, workspaceName string, credentials apicontracts.KubeconfigCredentials) (string, error) {
	if credentials.Username == "" || credentials.Password == "" {
		err := errors.New("username and password must be provided")
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("workspaceName", workspaceName))
		return "", err
	}

	if workspaceName == "" {
		err := errors.New("workspaceName must be provided")
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("workspaceName", workspaceName))
		return "", err
	}

	workspace, err := workspacesRepo.GetByName(ctx, workspaceName)
	if err != nil {
		err := fmt.Errorf("could not find workspace with name: %s", workspaceName)
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("workspaceName", workspaceName))
		return "", err
	}

	if workspace == nil {
		err := fmt.Errorf("could not find workspace with name: %s", workspaceName)
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("workspaceName", workspaceName))
		return "", err
	}

	kubeconfigstring, err := kubeconfigservice.GetKubeconfigForWorkspace(workspace, credentials)
	if err != nil {
		rlog.Errorc(ctx, "could not get kubeconfig", err, rlog.String("workspaceName", workspaceName))
		return "", err
	}

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	_, err = auditlog.Create(ctx, "Identity fetching kubeconfig for workspace",
		models.AuditCategoryKubeconfig,
		models.AuditActionRead,
		identity.User,
		fmt.Sprintf("identity type: '%s', id: '%s' fetching kubeconfig for workspace name: %s", identity.Type, identity.GetId(), workspaceName),
		nil)
	if err != nil {
		return "", fmt.Errorf("could not audit log fetch action: %v", err)
	}

	return kubeconfigstring, nil
}
