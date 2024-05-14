package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
)

type ResourceInterface interface {
	GetClusterOrderByUid(uid string, ownerSubject aclmodels.Acl2Subject, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceClusterOrder, error)
	GetClusterOrders(ownerSubject aclmodels.Acl2Subject, scope aclmodels.Acl2Scope) ([]*apiresourcecontracts.ResourceClusterOrder, error)
	UpdateClusterOrder(clusterOrder *apiresourcecontracts.ResourceUpdateModel) error
	GetHashList(ownerref rortypes.RorResourceOwnerReference) (apiresourcecontracts.HashList, error)
	GetTanzuKubernetesClusterByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceTanzuKubernetesCluster, error)

	GetApplicationByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceApplication, error)
	GetPVCByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourcePersistentVolumeClaim, error)
}
