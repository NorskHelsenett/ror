package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
)

type ResourceInterface interface {
	GetClusterOrderByUid(uid, ownerSubject, kind, apiversion string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceClusterOrder, error)
	GetClusterOrders(ownerSubject, kind, apiversion string, scope aclmodels.Acl2Scope) ([]*apiresourcecontracts.ResourceClusterOrder, error)
	UpdateClusterOrder(clusterOrder *apiresourcecontracts.ResourceUpdateModel) error

	GetTanzuKubernetesClusterByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceTanzuKubernetesCluster, error)

	GetApplicationByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceApplication, error)
	GetPVCByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourcePersistentVolumeClaim, error)
}
