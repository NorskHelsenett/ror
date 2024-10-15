package operatorconfigservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/helpers/mapping"
	"github.com/NorskHelsenett/ror/internal/models"
	"github.com/NorskHelsenett/ror/internal/mongodbrepo/mongoTypes"
	operatorconfigrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/operatorconfigRepo"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func GetAll(ctx context.Context) (*[]apicontracts.OperatorConfig, error) {
	configs, err := operatorconfigrepo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("could not get operatorconfigs")
	}

	return configs, nil
}

func GetById(ctx context.Context, id string) (*apicontracts.OperatorConfig, error) {
	config, err := operatorconfigrepo.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("could not get operator config by id")
	}

	return config, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[apicontracts.OperatorConfig], error) {
	result, err := operatorconfigrepo.GetByFilter(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not get operator configs: %v", err)
	}

	return result, nil
}

func Create(ctx context.Context, input *apicontracts.OperatorConfig) (*apicontracts.OperatorConfig, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	exists, err := operatorconfigrepo.FindOne(ctx, "kind", input.Kind)
	if err != nil {
		return exists, fmt.Errorf("could not check if operatorconfig exists: %v", err)
	}

	if exists != nil {
		return exists, fmt.Errorf("operatorconfig already exists")
	}

	var mappedInput mongoTypes.MongoOperatorConfig
	err = mapping.Map(input, &mappedInput)
	if err != nil {
		return nil, fmt.Errorf("could not map operatorconfig from apitype to mongotype: %v", err)
	}

	created, err := operatorconfigrepo.Create(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("could not create operatorconfig: %v", err)
	}

	_, err = auditlog.Create(ctx, "New operator config created", models.AuditCategoryConfiguration, models.AuditActionCreate, identity.User, created, nil)
	if err != nil {
		rlog.Error("failed to create auditlog", err)
	}

	return created, nil
}

func Update(ctx context.Context, id string, input *apicontracts.OperatorConfig) (*apicontracts.OperatorConfig, *apicontracts.OperatorConfig, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	var mongoOperatorConfig mongoTypes.MongoOperatorConfig
	err := mapping.Map(input, &mongoOperatorConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map from apitype to mongotype: %v", err)
	}

	newObject, oldObject, err := operatorconfigrepo.Update(ctx, id, mongoOperatorConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("could not update operatorconfig: %v", err)
	}

	_, err = auditlog.Create(ctx, "New operator config updated", models.AuditCategoryConfiguration, models.AuditActionUpdate, identity.User, newObject, oldObject)
	if err != nil {
		rlog.Error("failed to create auditlog", err)
	}

	var mappedNewObject apicontracts.OperatorConfig
	err = mapping.Map(newObject, &mappedNewObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map updated operatorconfig from apitype to mongotype: %v", err)
	}

	var mappedOldObject apicontracts.OperatorConfig
	err = mapping.Map(oldObject, &mappedOldObject)
	if err != nil {
		return nil, nil, fmt.Errorf("could not map original operatorconfig from apitype to mongotype: %v", err)
	}

	return &mappedNewObject, &mappedOldObject, nil
}

func Delete(ctx context.Context, id string) (bool, error) {
	identity := rorcontext.GetIdentityFromRorContext(ctx)
	deleted, err := operatorconfigrepo.Delete(ctx, id)
	if err != nil {
		return false, fmt.Errorf("could not delete operatorconfig: %v", err)
	}

	_, err = auditlog.Create(ctx, "New operator config updated", models.AuditCategoryConfiguration, models.AuditActionUpdate, identity.User, nil, deleted)
	if err != nil {
		rlog.Error("failed to create auditlog", err)
	}

	return deleted, nil
}
