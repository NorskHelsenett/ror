package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/resourcecachehashlist"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

type ResourceInterface interface {
	Create(ctx context.Context, resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error
	Update(ctx context.Context, resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error
	Delete(ctx context.Context, uid string) error

	GetClusterOrderByUid(ctx context.Context, uid string, ownerSubject aclscope.Subject, scope aclscope.Scope) (*apiresourcecontracts.ResourceClusterOrder, error)
	GetClusterOrders(ctx context.Context, ownerSubject aclscope.Subject, scope aclscope.Scope) ([]*apiresourcecontracts.ResourceClusterOrder, error)
	UpdateClusterOrder(ctx context.Context, clusterOrder *apiresourcecontracts.ResourceUpdateModel) error
	GetHashList(ctx context.Context, ownerref rorresourceowner.RorResourceOwnerReference) (resourcecachehashlist.HashList, error)
	GetTanzuKubernetesClusterByUid(ctx context.Context, uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourceTanzuKubernetesCluster, error)

	GetApplicationByUid(ctx context.Context, uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourceApplication, error)
	GetPVCByUid(ctx context.Context, uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourcePersistentVolumeClaim, error)

	GetVulnerabilityReportByUid(ctx context.Context, uid, owner string, scope aclscope.Scope) (*apiresourcecontracts.ResourceVulnerabilityReport, error)
	GetVulnerabilityReportsByOwner(ctx context.Context, owner string, scope aclscope.Scope) ([]apiresourcecontracts.ResourceVulnerabilityReport, error)

	GetClusterVulnerabilityReportByUid(ctx context.Context, uid, owner string, scope aclscope.Scope) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error)
	CreateClusterVulnerabilityReport(ctx context.Context, report *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error)
	UpdateClusterVulnerabilityReportByUid(ctx context.Context, report *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error)

	GetRoutesByOwner(ctx context.Context, owner string, scope aclscope.Scope) ([]apiresourcecontracts.ResourceRoute, error)

	GetSlackMessageByUid(ctx context.Context, uid, owner string, scope aclscope.Scope) (*apiresourcecontracts.ResourceSlackMessage, error)
	CreateSlackMessage(ctx context.Context, sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error)
	UpdateSlackMessageByUid(ctx context.Context, sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error)

	GetVulnerabilityEventByUid(ctx context.Context, uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourceVulnerabilityEvent, error)
	CreateVulnerabilityEvent(ctx context.Context, u *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceVulnerabilityEvent, error)
}
