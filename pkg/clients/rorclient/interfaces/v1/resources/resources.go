package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/resourcecachehashlist"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

type ResourceInterface interface {
	Create(resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error
	Update(resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error
	Delete(uid string) error

	GetClusterOrderByUid(uid string, ownerSubject aclscope.Subject, scope aclscope.Scope) (*apiresourcecontracts.ResourceClusterOrder, error)
	GetClusterOrders(ownerSubject aclscope.Subject, scope aclscope.Scope) ([]*apiresourcecontracts.ResourceClusterOrder, error)
	UpdateClusterOrder(clusterOrder *apiresourcecontracts.ResourceUpdateModel) error
	GetHashList(ownerref rorresourceowner.RorResourceOwnerReference) (resourcecachehashlist.HashList, error)
	GetTanzuKubernetesClusterByUid(uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourceTanzuKubernetesCluster, error)

	GetApplicationByUid(uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourceApplication, error)
	GetPVCByUid(uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourcePersistentVolumeClaim, error)

	GetVulnerabilityReportByUid(uid, owner string, scope aclscope.Scope) (*apiresourcecontracts.ResourceVulnerabilityReport, error)
	GetVulnerabilityReportsByOwner(owner string, scope aclscope.Scope) ([]apiresourcecontracts.ResourceVulnerabilityReport, error)

	GetClusterVulnerabilityReportByUid(uid, owner string, scope aclscope.Scope) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error)
	CreateClusterVulnerabilityReport(report *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error)
	UpdateClusterVulnerabilityReportByUid(report *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error)

	GetRoutesByOwner(owner string, scope aclscope.Scope) ([]apiresourcecontracts.ResourceRoute, error)

	GetSlackMessageByUid(uid, owner string, scope aclscope.Scope) (*apiresourcecontracts.ResourceSlackMessage, error)
	CreateSlackMessage(sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error)
	UpdateSlackMessageByUid(sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error)

	GetVulnerabilityEventByUid(uid, ownerSubject string, scope aclscope.Scope) (*apiresourcecontracts.ResourceVulnerabilityEvent, error)
	CreateVulnerabilityEvent(u *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceVulnerabilityEvent, error)
}
