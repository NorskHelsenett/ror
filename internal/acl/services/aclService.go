package aclservice

import (
	"context"
	"fmt"
	"time"

	aclrepository "github.com/NorskHelsenett/ror/internal/acl/repositories"
	"github.com/NorskHelsenett/ror/internal/auditlog"
	"github.com/NorskHelsenett/ror/internal/models"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
)

func GetById(ctx context.Context, id string) (*aclmodels.AclV2ListItem, error) {
	object, err := aclrepository.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get object by Id from repository: %v", err)
	}

	return object, nil
}

func GetByFilter(ctx context.Context, filter *apicontracts.Filter) (*apicontracts.PaginatedResult[aclmodels.AclV2ListItem], error) {
	acl, totalCount, err := aclrepository.GetByFilter(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error when getting acl by filter from repo: %v", err)
	}

	paginatedResult := apicontracts.PaginatedResult[aclmodels.AclV2ListItem]{}

	paginatedResult.Data = acl
	paginatedResult.DataCount = int64(len(acl))
	paginatedResult.Offset = int64(filter.Skip)
	paginatedResult.TotalCount = int64(totalCount)

	return &paginatedResult, nil
}

func Create(ctx context.Context, aclModel *aclmodels.AclV2ListItem, identity *identitymodels.Identity) (*aclmodels.AclV2ListItem, error) {
	aclModel.Created = time.Now()
	object, err := aclrepository.Create(ctx, aclModel)
	if err != nil {
		return nil, fmt.Errorf("could not get object by ID from repository: %v", err)
	}

	_, err = auditlog.Create(ctx, "ACL created", models.AuditCategoryAcl, models.AuditActionCreate, identity.User, object, nil)
	if err != nil {
		return nil, fmt.Errorf("could not audit log create action: %v", err)
	}

	return object, nil
}

func Update(ctx context.Context, aclId string, aclModel *aclmodels.AclV2ListItem, identity *identitymodels.Identity) (*aclmodels.AclV2ListItem, error) {
	object, oldObject, err := aclrepository.Update(ctx, aclId, aclModel)
	if err != nil {
		return nil, fmt.Errorf("could not get object by id from repository: %v", err)
	}

	_, err = auditlog.Create(ctx, "ACL updated", models.AuditCategoryAcl, models.AuditActionUpdate, identity.User, object, oldObject)
	if err != nil {
		return nil, fmt.Errorf("could not audit log: %v", err)
	}

	return object, nil
}

func Delete(ctx context.Context, aclId string, identity *identitymodels.Identity) (bool, *aclmodels.AclV2ListItem, error) {
	if !identity.IsUser() {
		return false, nil, fmt.Errorf("could not delete object, must be delete by a user")
	}

	deleted, deletedObject, err := aclrepository.Delete(ctx, aclId)
	if err != nil {
		return false, nil, fmt.Errorf("could not delete object: %v", err)
	}

	_, err = auditlog.Create(ctx, "Acl deleted", models.AuditCategoryAcl, models.AuditActionDelete, identity.User, deleted, nil)
	if err != nil {
		return false, nil, fmt.Errorf("could not audit log delete action: %v", err)
	}

	return deleted, deletedObject, nil
}

// Gets ACL2 Access model for user/scope/subject returns aclmodels.AclV2ListItemAccess
func CheckAccessByContextScopeSubject(ctx context.Context, scope any, subject any) aclmodels.AclV2ListItemAccess {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "aclService.CheckAccessByContextScopeSubject")
	defer span.End()
	_, ok := ctx.(*gin.Context)
	if ok {
		rlog.Fatal("not expecting a gin.Context", nil)
	}
	aclModel := aclmodels.NewAclV2QueryAccessScopeSubject(scope, subject)

	return aclrepository.CheckAcl2ByIdentityQuery(ctx, aclModel)
}
func CheckAccessByContextAclQuery(ctx context.Context, query aclmodels.AclV2QueryAccessScopeSubject) aclmodels.AclV2ListItemAccess {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "aclService.CheckAccessByContextScopeSubject")
	defer span.End()
	_, ok := ctx.(*gin.Context)
	if ok {
		rlog.Fatal("not expecting a gin.Context", nil)
	}

	if !query.IsValid() {
		return aclmodels.AclV2ListItemAccess{}
	}

	return aclrepository.CheckAcl2ByIdentityQuery(ctx, query)
}

// Deprecated: use CheckAccessByRorOwnerref
// Gets ACL2 Access model for user/scope/subject returns aclmodels.AclV2ListItemAccess
func CheckAccessByOwnerref(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) aclmodels.AclV2ListItemAccess {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "aclService.CheckAccessByContextScopeSubject")
	defer span.End()
	_, ok := ctx.(*gin.Context)
	if ok {
		rlog.Fatal("not expecting a gin.Context", nil)
	}

	aclModel := aclmodels.NewAclV2QueryAccessScopeSubject(ownerref.Scope, ownerref.Subject)

	return aclrepository.CheckAcl2ByIdentityQuery(ctx, aclModel)
}

func CheckAccessByRorOwnerref(ctx context.Context, ownerref rortypes.RorResourceOwnerReference) aclmodels.AclV2ListItemAccess {
	ctx, span := otel.GetTracerProvider().Tracer(viper.GetString(configconsts.TRACER_ID)).Start(ctx, "aclService.CheckAccessByContextScopeSubject")
	defer span.End()
	_, ok := ctx.(*gin.Context)
	if ok {
		rlog.Fatal("not expecting a gin.Context", nil)
	}

	aclModel := aclmodels.NewAclV2QueryAccessScopeSubject(ownerref.Scope, ownerref.Subject)

	return aclrepository.CheckAcl2ByIdentityQuery(ctx, aclModel)
}
