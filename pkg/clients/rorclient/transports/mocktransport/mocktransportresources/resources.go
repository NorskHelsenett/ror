package mocktransportresources

import (
	"errors"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) GetClusterOrderByUid(uid string, ownerSubject aclmodels.Acl2Subject, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceClusterOrder, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	clusterOrder := &apiresourcecontracts.ResourceClusterOrder{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: uid,
		},
	}
	return clusterOrder, nil
}

func (c *V1Client) GetClusterOrders(ownerSubject aclmodels.Acl2Subject, scope aclmodels.Acl2Scope) ([]*apiresourcecontracts.ResourceClusterOrder, error) {
	clusterOrders := []*apiresourcecontracts.ResourceClusterOrder{
		{
			Metadata: apiresourcecontracts.ResourceMetadata{
				Uid: "mock-cluster-order-001",
			},
		},
		{
			Metadata: apiresourcecontracts.ResourceMetadata{
				Uid: "mock-cluster-order-002",
			},
		},
	}
	return clusterOrders, nil
}

func (c *V1Client) UpdateClusterOrder(clusterOrder *apiresourcecontracts.ResourceUpdateModel) error {
	if clusterOrder == nil {
		return errors.New("cluster order cannot be nil")
	}
	return nil
}

func (c *V1Client) GetHashList(ownerref rorresourceowner.RorResourceOwnerReference) (apiresourcecontracts.HashList, error) {
	hashList := apiresourcecontracts.HashList{
		Items: []apiresourcecontracts.HashItem{
			{
				Uid:  "mock-hash-001",
				Hash: "mock-hash-value-001",
			},
		},
	}
	return hashList, nil
}

func (c *V1Client) GetTanzuKubernetesClusterByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceTanzuKubernetesCluster, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	cluster := &apiresourcecontracts.ResourceTanzuKubernetesCluster{
		Metadata: apiresourcecontracts.ResourceTanuzKuberntesMetadata{
			Uid: uid,
		},
	}
	return cluster, nil
}

func (c *V1Client) GetApplicationByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceApplication, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	application := &apiresourcecontracts.ResourceApplication{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: uid,
		},
	}
	return application, nil
}

func (c *V1Client) GetPVCByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourcePersistentVolumeClaim, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	pvc := &apiresourcecontracts.ResourcePersistentVolumeClaim{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: uid,
		},
	}
	return pvc, nil
}

func (c *V1Client) GetVulnerabilityReportByUid(uid, owner string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceVulnerabilityReport, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	report := &apiresourcecontracts.ResourceVulnerabilityReport{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: uid,
		},
	}
	return report, nil
}

func (c *V1Client) GetVulnerabilityReportsByOwner(owner string, scope aclmodels.Acl2Scope) ([]apiresourcecontracts.ResourceVulnerabilityReport, error) {
	reports := []apiresourcecontracts.ResourceVulnerabilityReport{
		{
			Metadata: apiresourcecontracts.ResourceMetadata{
				Uid: "mock-vuln-report-001",
			},
		},
		{
			Metadata: apiresourcecontracts.ResourceMetadata{
				Uid: "mock-vuln-report-002",
			},
		},
	}
	return reports, nil
}

func (c *V1Client) GetClusterVulnerabilityReportByUid(uid, owner string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	report := &apiresourcecontracts.ResourceClusterVulnerabilityReport{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: uid,
		},
	}
	return report, nil
}

func (c *V1Client) CreateClusterVulnerabilityReport(report *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error) {
	if report == nil {
		return nil, errors.New("report cannot be nil")
	}

	result := &apiresourcecontracts.ResourceClusterVulnerabilityReport{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: "mock-cluster-vuln-report-new-001",
		},
	}
	return result, nil
}

func (c *V1Client) UpdateClusterVulnerabilityReportByUid(report *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceClusterVulnerabilityReport, error) {
	if report == nil {
		return nil, errors.New("report cannot be nil")
	}

	result := &apiresourcecontracts.ResourceClusterVulnerabilityReport{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: report.Uid,
		},
	}
	return result, nil
}

func (c *V1Client) GetRoutesByOwner(owner string, scope aclmodels.Acl2Scope) ([]apiresourcecontracts.ResourceRoute, error) {
	routes := []apiresourcecontracts.ResourceRoute{
		{
			Metadata: apiresourcecontracts.ResourceMetadata{
				Uid: "mock-route-001",
			},
		},
		{
			Metadata: apiresourcecontracts.ResourceMetadata{
				Uid: "mock-route-002",
			},
		},
	}
	return routes, nil
}

func (c *V1Client) GetSlackMessageByUid(uid, owner string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceSlackMessage, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	message := &apiresourcecontracts.ResourceSlackMessage{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: uid,
		},
	}
	return message, nil
}

func (c *V1Client) CreateSlackMessage(sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error) {
	if sm == nil {
		return nil, errors.New("slack message cannot be nil")
	}

	result := &apiresourcecontracts.ResourceSlackMessage{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: "mock-slack-message-new-001",
		},
	}
	return result, nil
}

func (c *V1Client) UpdateSlackMessageByUid(sm *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceSlackMessage, error) {
	if sm == nil {
		return nil, errors.New("slack message cannot be nil")
	}

	result := &apiresourcecontracts.ResourceSlackMessage{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: sm.Uid,
		},
	}
	return result, nil
}

func (c *V1Client) GetVulnerabilityEventByUid(uid, ownerSubject string, scope aclmodels.Acl2Scope) (*apiresourcecontracts.ResourceVulnerabilityEvent, error) {
	if uid == "" {
		return nil, errors.New("UID cannot be empty")
	}

	event := &apiresourcecontracts.ResourceVulnerabilityEvent{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: uid,
		},
	}
	return event, nil
}

func (c *V1Client) CreateVulnerabilityEvent(u *apiresourcecontracts.ResourceUpdateModel) (*apiresourcecontracts.ResourceVulnerabilityEvent, error) {
	if u == nil {
		return nil, errors.New("update model cannot be nil")
	}

	result := &apiresourcecontracts.ResourceVulnerabilityEvent{
		Metadata: apiresourcecontracts.ResourceMetadata{
			Uid: "mock-vuln-event-new-001",
		},
	}
	return result, nil
}
