package desiredversionservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/models"
	desiredversionrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/desiredversionRepo"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

func GetByKey(ctx context.Context, key string) (*apicontracts.DesiredVersion, error) {
	return desiredversionrepo.GetByKey(ctx, key)
}

func GetByID(ctx context.Context, id interface{}) (*apicontracts.DesiredVersion, error) {
	return desiredversionrepo.GetByID(ctx, id)
}
func GetAll(ctx context.Context) ([]apicontracts.DesiredVersion, error) {
	desiredVersions, err := desiredversionrepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get desired versions:  %v", err)
	}

	if desiredVersions == nil {
		return make([]apicontracts.DesiredVersion, 0), nil
	}

	return desiredVersions, nil
}

func Create(ctx context.Context, desiredversion apicontracts.DesiredVersion) (*apicontracts.DesiredVersion, error) {
	exists, err := GetByKey(ctx, desiredversion.Key)
	if err != nil {
		return nil, err
	}

	if exists.Key != "" {
		return nil, fmt.Errorf("desired version already exists")
	}

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	if !identity.IsUser() {
		errMsg := "must be a user to create"
		rlog.Errorc(ctx, errMsg, nil)
		return nil, fmt.Errorf(errMsg)
	}

	creationresult, err := desiredversionrepo.Create(ctx, desiredversion)
	if err != nil {
		return nil, err
	}

	_, err = auditlog.Create(ctx, "New desired version created", models.AuditCategoryConfiguration, models.AuditActionCreate, identity.User, creationresult, nil)
	if err != nil {
		rlog.Errorc(ctx, "failed to create auditlog", err)
		return nil, err
	}

	// validate creation
	creation, err := GetByID(ctx, creationresult.InsertedID)
	if err != nil {
		return nil, err
	}

	return creation, nil
}

func UpdateByKey(ctx context.Context, key string, desiredversion apicontracts.DesiredVersion) (*apicontracts.DesiredVersion, error) {
	// we dont want to create duplicates by changing the key to an existing key
	if key != desiredversion.Key {
		exists, err := GetByKey(ctx, desiredversion.Key)
		if err != nil {
			return nil, err
		}
		if exists.Key != "" {
			return nil, fmt.Errorf("another desired version with key: %s already exists", desiredversion.Key)
		}
	}

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	if !identity.IsUser() {
		return nil, errors.New("must be a user to modify")
	}

	updateResult, err := desiredversionrepo.UpdateByKey(ctx, key, desiredversion)
	if err != nil {
		return nil, err
	}

	_, err = auditlog.Create(ctx, "Desired version updated", models.AuditCategoryConfiguration, models.AuditActionUpdate, identity.User, updateResult, nil)
	if err != nil {
		rlog.Errorc(ctx, "failed to create auditlog", err)
		return nil, err
	}

	update, err := GetByKey(ctx, key)
	if err != nil {
		return nil, err
	}

	return update, nil
}

func DeleteByKey(ctx context.Context, key string) (int64, error) {
	deleteResult, err := desiredversionrepo.DeleteByKey(ctx, key)
	if err != nil {
		return 0, err
	}

	identity := rorcontext.GetIdentityFromRorContext(ctx)
	if !identity.IsUser() {
		return 0, errors.New("must be a user to modify")
	}

	_, err = auditlog.Create(ctx, "Desired version deleted", models.AuditCategoryConfiguration, models.AuditActionDelete, identity.User, key, nil)
	if err != nil {
		rlog.Errorc(ctx, "failed to create auditlog", err)
		return 0, err
	}

	if deleteResult.DeletedCount == 0 {
		return 0, fmt.Errorf("no desired version was deleted")
	}

	return deleteResult.DeletedCount, nil
}
