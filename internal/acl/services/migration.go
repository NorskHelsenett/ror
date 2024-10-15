package aclservice

import (
	"context"
	aclrepository "github.com/NorskHelsenett/ror/internal/acl/repositories"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

// TODO: remove when migration is complete
func MigrateAcl1toAcl2(ctx context.Context) error {

	aclListV1, err := aclrepository.GetAllACL1(ctx)
	if err != nil {
		rlog.Error("could not get ACL1 model", err)
	}

	aclListV2, err := aclrepository.GetAllACL2(ctx)
	if err != nil {
		rlog.Error("could not get ACL2 model", err)
	}

	_ = aclrepository.MigrateAcl1UpdateCreatreAcl2(ctx, aclListV1, aclListV2)
	aclrepository.MigrateAcl1DeleteRemovedAcl1(ctx, aclListV1, aclListV2)

	return nil
}
