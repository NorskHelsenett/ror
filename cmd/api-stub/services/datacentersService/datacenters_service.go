package datacentersservice

import (
	"context"
	"errors"

	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/models"
	datacentersRepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/datacentersRepo"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

func GetAllByUser(ctx context.Context) (*[]apicontracts.Datacenter, error) {
	datacenters, err := datacentersRepo.GetAllByUser(ctx)
	if err != nil {
		return nil, errors.New("could not fetch datacenters")
	}

	return datacenters, nil
}

func GetById(ctx context.Context, datacenterId string) (*apicontracts.Datacenter, error) {
	datacenter, err := datacentersRepo.GetById(ctx, datacenterId)
	if err != nil {
		return nil, errors.New("could not fetch datacenter by id")
	}

	return datacenter, nil
}

func GetByName(ctx context.Context, datacenterName string) (*apicontracts.Datacenter, error) {
	datacenter, err := datacentersRepo.FindByName(ctx, datacenterName)
	if err != nil {
		return nil, errors.New("could not fetch datacenter by name")
	}

	return datacenter, nil
}

func Create(ctx context.Context, datacenterInput *apicontracts.DatacenterModel, user *identitymodels.User) (*apicontracts.Datacenter, error) {
	exists, err := datacentersRepo.FindByName(ctx, datacenterInput.Name)
	if err != nil {
		rlog.Error("could not create datacenter", err)
		return nil, errors.New("could not fetch datacenter")
	}

	if exists != nil {
		return nil, nil
	}

	datacenterResult, err := datacentersRepo.Create(ctx, datacenterInput, user)
	if err != nil {
		rlog.Error("could not create datacenter", err)
		return nil, errors.New("could not fetch datacenters")
	}

	_, err = auditlog.Create(ctx, "New datacenter created", models.AuditCategoryDatacenter, models.AuditActionCreate, user, datacenterResult, nil)
	if err != nil {
		rlog.Error("failed to create auditlog", err)
	}

	return datacenterResult, nil
}

func Update(ctx context.Context, datacenterId string, datacenterInput *apicontracts.DatacenterModel, user *identitymodels.User) (*apicontracts.Datacenter, error) {
	datacenter, err := datacentersRepo.GetById(ctx, datacenterId)
	if err != nil {
		rlog.Error("could not update datacenter", err)
		return nil, errors.New("could not update datacenter")
	}

	if datacenter == nil {
		return nil, errors.New("could not find datacenter")
	}

	updated, err := datacentersRepo.Update(ctx, datacenterInput, user)
	if err != nil {
		rlog.Error("could not update datacenter", err)
		return nil, errors.New("could not update datacenter")
	}

	_, err = auditlog.Create(ctx, "Datacenter updated", models.AuditCategoryDatacenter, models.AuditActionUpdate, user, updated, datacenter)
	if err != nil {
		rlog.Error("failed to create auditlog", err)
	}

	return updated, nil
}
