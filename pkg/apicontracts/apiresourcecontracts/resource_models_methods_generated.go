// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package apiresourcecontracts

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
)

// Function to return Namespace resource by name.
func (m ResourceListNamespaces) GetByName(search string) ResourceNamespace {
	for _, resource := range m.Namespaces {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceNamespace
	return emptyResponse
}

// Function to return Namespace resource by uid.
func (m ResourceListNamespaces) GetByUid(search string) ResourceNamespace {
	for _, res := range m.Namespaces {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceNamespace
	return emptyResponse
}

// Function to return Namespace resource by label.
func (m ResourceListNamespaces) GetByLabels(search map[string]string) []ResourceNamespace {
	var Response []ResourceNamespace
	for _, res := range m.Namespaces {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Node resource by name.
func (m ResourceListNodes) GetByName(search string) ResourceNode {
	for _, resource := range m.Nodes {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceNode
	return emptyResponse
}

// Function to return Node resource by uid.
func (m ResourceListNodes) GetByUid(search string) ResourceNode {
	for _, res := range m.Nodes {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceNode
	return emptyResponse
}

// Function to return Node resource by label.
func (m ResourceListNodes) GetByLabels(search map[string]string) []ResourceNode {
	var Response []ResourceNode
	for _, res := range m.Nodes {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return PersistentVolumeClaim resource by name.
func (m ResourceListPersistentvolumeclaims) GetByName(search string) ResourcePersistentVolumeClaim {
	for _, resource := range m.Persistentvolumeclaims {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourcePersistentVolumeClaim
	return emptyResponse
}

// Function to return PersistentVolumeClaim resource by uid.
func (m ResourceListPersistentvolumeclaims) GetByUid(search string) ResourcePersistentVolumeClaim {
	for _, res := range m.Persistentvolumeclaims {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourcePersistentVolumeClaim
	return emptyResponse
}

// Function to return PersistentVolumeClaim resource by label.
func (m ResourceListPersistentvolumeclaims) GetByLabels(search map[string]string) []ResourcePersistentVolumeClaim {
	var Response []ResourcePersistentVolumeClaim
	for _, res := range m.Persistentvolumeclaims {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Deployment resource by name.
func (m ResourceListDeployments) GetByName(search string) ResourceDeployment {
	for _, resource := range m.Deployments {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceDeployment
	return emptyResponse
}

// Function to return Deployment resource by namespace.
func (m ResourceListDeployments) GetByNamespace(search string) ResourceDeployment {
	for _, res := range m.Deployments {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceDeployment
	return emptyResponse
}

// Function to return Deployment resource by uid.
func (m ResourceListDeployments) GetByUid(search string) ResourceDeployment {
	for _, res := range m.Deployments {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceDeployment
	return emptyResponse
}

// Function to return Deployment resource by label.
func (m ResourceListDeployments) GetByLabels(search map[string]string) []ResourceDeployment {
	var Response []ResourceDeployment
	for _, res := range m.Deployments {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return StorageClass resource by name.
func (m ResourceListStorageclasses) GetByName(search string) ResourceStorageClass {
	for _, resource := range m.Storageclasses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceStorageClass
	return emptyResponse
}

// Function to return StorageClass resource by uid.
func (m ResourceListStorageclasses) GetByUid(search string) ResourceStorageClass {
	for _, res := range m.Storageclasses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceStorageClass
	return emptyResponse
}

// Function to return StorageClass resource by label.
func (m ResourceListStorageclasses) GetByLabels(search map[string]string) []ResourceStorageClass {
	var Response []ResourceStorageClass
	for _, res := range m.Storageclasses {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return PolicyReport resource by name.
func (m ResourceListPolicyreports) GetByName(search string) ResourcePolicyReport {
	for _, resource := range m.Policyreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourcePolicyReport
	return emptyResponse
}

// Function to return PolicyReport resource by namespace.
func (m ResourceListPolicyreports) GetByNamespace(search string) ResourcePolicyReport {
	for _, res := range m.Policyreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourcePolicyReport
	return emptyResponse
}

// Function to return PolicyReport resource by uid.
func (m ResourceListPolicyreports) GetByUid(search string) ResourcePolicyReport {
	for _, res := range m.Policyreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourcePolicyReport
	return emptyResponse
}

// Function to return PolicyReport resource by label.
func (m ResourceListPolicyreports) GetByLabels(search map[string]string) []ResourcePolicyReport {
	var Response []ResourcePolicyReport
	for _, res := range m.Policyreports {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Application resource by name.
func (m ResourceListApplications) GetByName(search string) ResourceApplication {
	for _, resource := range m.Applications {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceApplication
	return emptyResponse
}

// Function to return Application resource by namespace.
func (m ResourceListApplications) GetByNamespace(search string) ResourceApplication {
	for _, res := range m.Applications {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceApplication
	return emptyResponse
}

// Function to return Application resource by uid.
func (m ResourceListApplications) GetByUid(search string) ResourceApplication {
	for _, res := range m.Applications {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceApplication
	return emptyResponse
}

// Function to return Application resource by label.
func (m ResourceListApplications) GetByLabels(search map[string]string) []ResourceApplication {
	var Response []ResourceApplication
	for _, res := range m.Applications {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return AppProject resource by name.
func (m ResourceListAppprojects) GetByName(search string) ResourceAppProject {
	for _, resource := range m.Appprojects {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceAppProject
	return emptyResponse
}

// Function to return AppProject resource by namespace.
func (m ResourceListAppprojects) GetByNamespace(search string) ResourceAppProject {
	for _, res := range m.Appprojects {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceAppProject
	return emptyResponse
}

// Function to return AppProject resource by uid.
func (m ResourceListAppprojects) GetByUid(search string) ResourceAppProject {
	for _, res := range m.Appprojects {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceAppProject
	return emptyResponse
}

// Function to return AppProject resource by label.
func (m ResourceListAppprojects) GetByLabels(search map[string]string) []ResourceAppProject {
	var Response []ResourceAppProject
	for _, res := range m.Appprojects {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Certificate resource by name.
func (m ResourceListCertificates) GetByName(search string) ResourceCertificate {
	for _, resource := range m.Certificates {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceCertificate
	return emptyResponse
}

// Function to return Certificate resource by namespace.
func (m ResourceListCertificates) GetByNamespace(search string) ResourceCertificate {
	for _, res := range m.Certificates {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceCertificate
	return emptyResponse
}

// Function to return Certificate resource by uid.
func (m ResourceListCertificates) GetByUid(search string) ResourceCertificate {
	for _, res := range m.Certificates {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceCertificate
	return emptyResponse
}

// Function to return Certificate resource by label.
func (m ResourceListCertificates) GetByLabels(search map[string]string) []ResourceCertificate {
	var Response []ResourceCertificate
	for _, res := range m.Certificates {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Service resource by name.
func (m ResourceListServices) GetByName(search string) ResourceService {
	for _, resource := range m.Services {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceService
	return emptyResponse
}

// Function to return Service resource by namespace.
func (m ResourceListServices) GetByNamespace(search string) ResourceService {
	for _, res := range m.Services {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceService
	return emptyResponse
}

// Function to return Service resource by uid.
func (m ResourceListServices) GetByUid(search string) ResourceService {
	for _, res := range m.Services {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceService
	return emptyResponse
}

// Function to return Service resource by label.
func (m ResourceListServices) GetByLabels(search map[string]string) []ResourceService {
	var Response []ResourceService
	for _, res := range m.Services {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Pod resource by name.
func (m ResourceListPods) GetByName(search string) ResourcePod {
	for _, resource := range m.Pods {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourcePod
	return emptyResponse
}

// Function to return Pod resource by namespace.
func (m ResourceListPods) GetByNamespace(search string) ResourcePod {
	for _, res := range m.Pods {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourcePod
	return emptyResponse
}

// Function to return Pod resource by uid.
func (m ResourceListPods) GetByUid(search string) ResourcePod {
	for _, res := range m.Pods {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourcePod
	return emptyResponse
}

// Function to return Pod resource by label.
func (m ResourceListPods) GetByLabels(search map[string]string) []ResourcePod {
	var Response []ResourcePod
	for _, res := range m.Pods {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return ReplicaSet resource by name.
func (m ResourceListReplicasets) GetByName(search string) ResourceReplicaSet {
	for _, resource := range m.Replicasets {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceReplicaSet
	return emptyResponse
}

// Function to return ReplicaSet resource by namespace.
func (m ResourceListReplicasets) GetByNamespace(search string) ResourceReplicaSet {
	for _, res := range m.Replicasets {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceReplicaSet
	return emptyResponse
}

// Function to return ReplicaSet resource by uid.
func (m ResourceListReplicasets) GetByUid(search string) ResourceReplicaSet {
	for _, res := range m.Replicasets {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceReplicaSet
	return emptyResponse
}

// Function to return ReplicaSet resource by label.
func (m ResourceListReplicasets) GetByLabels(search map[string]string) []ResourceReplicaSet {
	var Response []ResourceReplicaSet
	for _, res := range m.Replicasets {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return StatefulSet resource by name.
func (m ResourceListStatefulsets) GetByName(search string) ResourceStatefulSet {
	for _, resource := range m.Statefulsets {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceStatefulSet
	return emptyResponse
}

// Function to return StatefulSet resource by namespace.
func (m ResourceListStatefulsets) GetByNamespace(search string) ResourceStatefulSet {
	for _, res := range m.Statefulsets {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceStatefulSet
	return emptyResponse
}

// Function to return StatefulSet resource by uid.
func (m ResourceListStatefulsets) GetByUid(search string) ResourceStatefulSet {
	for _, res := range m.Statefulsets {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceStatefulSet
	return emptyResponse
}

// Function to return StatefulSet resource by label.
func (m ResourceListStatefulsets) GetByLabels(search map[string]string) []ResourceStatefulSet {
	var Response []ResourceStatefulSet
	for _, res := range m.Statefulsets {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return DaemonSet resource by name.
func (m ResourceListDaemonsets) GetByName(search string) ResourceDaemonSet {
	for _, resource := range m.Daemonsets {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceDaemonSet
	return emptyResponse
}

// Function to return DaemonSet resource by namespace.
func (m ResourceListDaemonsets) GetByNamespace(search string) ResourceDaemonSet {
	for _, res := range m.Daemonsets {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceDaemonSet
	return emptyResponse
}

// Function to return DaemonSet resource by uid.
func (m ResourceListDaemonsets) GetByUid(search string) ResourceDaemonSet {
	for _, res := range m.Daemonsets {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceDaemonSet
	return emptyResponse
}

// Function to return DaemonSet resource by label.
func (m ResourceListDaemonsets) GetByLabels(search map[string]string) []ResourceDaemonSet {
	var Response []ResourceDaemonSet
	for _, res := range m.Daemonsets {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Ingress resource by name.
func (m ResourceListIngresses) GetByName(search string) ResourceIngress {
	for _, resource := range m.Ingresses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceIngress
	return emptyResponse
}

// Function to return Ingress resource by namespace.
func (m ResourceListIngresses) GetByNamespace(search string) ResourceIngress {
	for _, res := range m.Ingresses {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceIngress
	return emptyResponse
}

// Function to return Ingress resource by uid.
func (m ResourceListIngresses) GetByUid(search string) ResourceIngress {
	for _, res := range m.Ingresses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceIngress
	return emptyResponse
}

// Function to return Ingress resource by label.
func (m ResourceListIngresses) GetByLabels(search map[string]string) []ResourceIngress {
	var Response []ResourceIngress
	for _, res := range m.Ingresses {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return IngressClass resource by name.
func (m ResourceListIngressclasses) GetByName(search string) ResourceIngressClass {
	for _, resource := range m.Ingressclasses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceIngressClass
	return emptyResponse
}

// Function to return IngressClass resource by namespace.
func (m ResourceListIngressclasses) GetByNamespace(search string) ResourceIngressClass {
	for _, res := range m.Ingressclasses {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceIngressClass
	return emptyResponse
}

// Function to return IngressClass resource by uid.
func (m ResourceListIngressclasses) GetByUid(search string) ResourceIngressClass {
	for _, res := range m.Ingressclasses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceIngressClass
	return emptyResponse
}

// Function to return IngressClass resource by label.
func (m ResourceListIngressclasses) GetByLabels(search map[string]string) []ResourceIngressClass {
	var Response []ResourceIngressClass
	for _, res := range m.Ingressclasses {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return VulnerabilityReport resource by name.
func (m ResourceListVulnerabilityreports) GetByName(search string) ResourceVulnerabilityReport {
	for _, resource := range m.Vulnerabilityreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceVulnerabilityReport
	return emptyResponse
}

// Function to return VulnerabilityReport resource by namespace.
func (m ResourceListVulnerabilityreports) GetByNamespace(search string) ResourceVulnerabilityReport {
	for _, res := range m.Vulnerabilityreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceVulnerabilityReport
	return emptyResponse
}

// Function to return VulnerabilityReport resource by uid.
func (m ResourceListVulnerabilityreports) GetByUid(search string) ResourceVulnerabilityReport {
	for _, res := range m.Vulnerabilityreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceVulnerabilityReport
	return emptyResponse
}

// Function to return VulnerabilityReport resource by label.
func (m ResourceListVulnerabilityreports) GetByLabels(search map[string]string) []ResourceVulnerabilityReport {
	var Response []ResourceVulnerabilityReport
	for _, res := range m.Vulnerabilityreports {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return ExposedSecretReport resource by name.
func (m ResourceListExposedsecretreports) GetByName(search string) ResourceExposedSecretReport {
	for _, resource := range m.Exposedsecretreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceExposedSecretReport
	return emptyResponse
}

// Function to return ExposedSecretReport resource by namespace.
func (m ResourceListExposedsecretreports) GetByNamespace(search string) ResourceExposedSecretReport {
	for _, res := range m.Exposedsecretreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceExposedSecretReport
	return emptyResponse
}

// Function to return ExposedSecretReport resource by uid.
func (m ResourceListExposedsecretreports) GetByUid(search string) ResourceExposedSecretReport {
	for _, res := range m.Exposedsecretreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceExposedSecretReport
	return emptyResponse
}

// Function to return ExposedSecretReport resource by label.
func (m ResourceListExposedsecretreports) GetByLabels(search map[string]string) []ResourceExposedSecretReport {
	var Response []ResourceExposedSecretReport
	for _, res := range m.Exposedsecretreports {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return ConfigAuditReport resource by name.
func (m ResourceListConfigauditreports) GetByName(search string) ResourceConfigAuditReport {
	for _, resource := range m.Configauditreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceConfigAuditReport
	return emptyResponse
}

// Function to return ConfigAuditReport resource by namespace.
func (m ResourceListConfigauditreports) GetByNamespace(search string) ResourceConfigAuditReport {
	for _, res := range m.Configauditreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceConfigAuditReport
	return emptyResponse
}

// Function to return ConfigAuditReport resource by uid.
func (m ResourceListConfigauditreports) GetByUid(search string) ResourceConfigAuditReport {
	for _, res := range m.Configauditreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceConfigAuditReport
	return emptyResponse
}

// Function to return ConfigAuditReport resource by label.
func (m ResourceListConfigauditreports) GetByLabels(search map[string]string) []ResourceConfigAuditReport {
	var Response []ResourceConfigAuditReport
	for _, res := range m.Configauditreports {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return RbacAssessmentReport resource by name.
func (m ResourceListRbacassessmentreports) GetByName(search string) ResourceRbacAssessmentReport {
	for _, resource := range m.Rbacassessmentreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceRbacAssessmentReport
	return emptyResponse
}

// Function to return RbacAssessmentReport resource by namespace.
func (m ResourceListRbacassessmentreports) GetByNamespace(search string) ResourceRbacAssessmentReport {
	for _, res := range m.Rbacassessmentreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceRbacAssessmentReport
	return emptyResponse
}

// Function to return RbacAssessmentReport resource by uid.
func (m ResourceListRbacassessmentreports) GetByUid(search string) ResourceRbacAssessmentReport {
	for _, res := range m.Rbacassessmentreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceRbacAssessmentReport
	return emptyResponse
}

// Function to return RbacAssessmentReport resource by label.
func (m ResourceListRbacassessmentreports) GetByLabels(search map[string]string) []ResourceRbacAssessmentReport {
	var Response []ResourceRbacAssessmentReport
	for _, res := range m.Rbacassessmentreports {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return TanzuKubernetesCluster resource by name.
func (m ResourceListTanzukubernetesclusters) GetByName(search string) ResourceTanzuKubernetesCluster {
	for _, resource := range m.Tanzukubernetesclusters {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceTanzuKubernetesCluster
	return emptyResponse
}

// Function to return TanzuKubernetesCluster resource by namespace.
func (m ResourceListTanzukubernetesclusters) GetByNamespace(search string) ResourceTanzuKubernetesCluster {
	for _, res := range m.Tanzukubernetesclusters {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceTanzuKubernetesCluster
	return emptyResponse
}

// Function to return TanzuKubernetesCluster resource by uid.
func (m ResourceListTanzukubernetesclusters) GetByUid(search string) ResourceTanzuKubernetesCluster {
	for _, res := range m.Tanzukubernetesclusters {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceTanzuKubernetesCluster
	return emptyResponse
}

// Function to return TanzuKubernetesCluster resource by label.
func (m ResourceListTanzukubernetesclusters) GetByLabels(search map[string]string) []ResourceTanzuKubernetesCluster {
	var Response []ResourceTanzuKubernetesCluster
	for _, res := range m.Tanzukubernetesclusters {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return TanzuKubernetesRelease resource by name.
func (m ResourceListTanzukubernetesreleases) GetByName(search string) ResourceTanzuKubernetesRelease {
	for _, resource := range m.Tanzukubernetesreleases {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceTanzuKubernetesRelease
	return emptyResponse
}

// Function to return TanzuKubernetesRelease resource by uid.
func (m ResourceListTanzukubernetesreleases) GetByUid(search string) ResourceTanzuKubernetesRelease {
	for _, res := range m.Tanzukubernetesreleases {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceTanzuKubernetesRelease
	return emptyResponse
}

// Function to return TanzuKubernetesRelease resource by label.
func (m ResourceListTanzukubernetesreleases) GetByLabels(search map[string]string) []ResourceTanzuKubernetesRelease {
	var Response []ResourceTanzuKubernetesRelease
	for _, res := range m.Tanzukubernetesreleases {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return VirtualMachineClass resource by name.
func (m ResourceListVirtualmachineclasses) GetByName(search string) ResourceVirtualMachineClass {
	for _, resource := range m.Virtualmachineclasses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceVirtualMachineClass
	return emptyResponse
}

// Function to return VirtualMachineClass resource by uid.
func (m ResourceListVirtualmachineclasses) GetByUid(search string) ResourceVirtualMachineClass {
	for _, res := range m.Virtualmachineclasses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceVirtualMachineClass
	return emptyResponse
}

// Function to return VirtualMachineClass resource by label.
func (m ResourceListVirtualmachineclasses) GetByLabels(search map[string]string) []ResourceVirtualMachineClass {
	var Response []ResourceVirtualMachineClass
	for _, res := range m.Virtualmachineclasses {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return KubernetesCluster resource by name.
func (m ResourceListKubernetesclusters) GetByName(search string) ResourceKubernetesCluster {
	for _, resource := range m.Kubernetesclusters {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceKubernetesCluster
	return emptyResponse
}

// Function to return KubernetesCluster resource by namespace.
func (m ResourceListKubernetesclusters) GetByNamespace(search string) ResourceKubernetesCluster {
	for _, res := range m.Kubernetesclusters {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceKubernetesCluster
	return emptyResponse
}

// Function to return KubernetesCluster resource by uid.
func (m ResourceListKubernetesclusters) GetByUid(search string) ResourceKubernetesCluster {
	for _, res := range m.Kubernetesclusters {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceKubernetesCluster
	return emptyResponse
}

// Function to return KubernetesCluster resource by label.
func (m ResourceListKubernetesclusters) GetByLabels(search map[string]string) []ResourceKubernetesCluster {
	var Response []ResourceKubernetesCluster
	for _, res := range m.Kubernetesclusters {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return ClusterOrder resource by name.
func (m ResourceListClusterorders) GetByName(search string) ResourceClusterOrder {
	for _, resource := range m.Clusterorders {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceClusterOrder
	return emptyResponse
}

// Function to return ClusterOrder resource by namespace.
func (m ResourceListClusterorders) GetByNamespace(search string) ResourceClusterOrder {
	for _, res := range m.Clusterorders {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceClusterOrder
	return emptyResponse
}

// Function to return ClusterOrder resource by uid.
func (m ResourceListClusterorders) GetByUid(search string) ResourceClusterOrder {
	for _, res := range m.Clusterorders {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceClusterOrder
	return emptyResponse
}

// Function to return ClusterOrder resource by label.
func (m ResourceListClusterorders) GetByLabels(search map[string]string) []ResourceClusterOrder {
	var Response []ResourceClusterOrder
	for _, res := range m.Clusterorders {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Project resource by name.
func (m ResourceListProjects) GetByName(search string) ResourceProject {
	for _, resource := range m.Projects {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceProject
	return emptyResponse
}

// Function to return Project resource by namespace.
func (m ResourceListProjects) GetByNamespace(search string) ResourceProject {
	for _, res := range m.Projects {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceProject
	return emptyResponse
}

// Function to return Project resource by uid.
func (m ResourceListProjects) GetByUid(search string) ResourceProject {
	for _, res := range m.Projects {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceProject
	return emptyResponse
}

// Function to return Project resource by label.
func (m ResourceListProjects) GetByLabels(search map[string]string) []ResourceProject {
	var Response []ResourceProject
	for _, res := range m.Projects {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Configuration resource by name.
func (m ResourceListConfigurations) GetByName(search string) ResourceConfiguration {
	for _, resource := range m.Configurations {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceConfiguration
	return emptyResponse
}

// Function to return Configuration resource by namespace.
func (m ResourceListConfigurations) GetByNamespace(search string) ResourceConfiguration {
	for _, res := range m.Configurations {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceConfiguration
	return emptyResponse
}

// Function to return Configuration resource by uid.
func (m ResourceListConfigurations) GetByUid(search string) ResourceConfiguration {
	for _, res := range m.Configurations {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceConfiguration
	return emptyResponse
}

// Function to return Configuration resource by label.
func (m ResourceListConfigurations) GetByLabels(search map[string]string) []ResourceConfiguration {
	var Response []ResourceConfiguration
	for _, res := range m.Configurations {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return ClusterComplianceReport resource by name.
func (m ResourceListClustercompliancereports) GetByName(search string) ResourceClusterComplianceReport {
	for _, resource := range m.Clustercompliancereports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceClusterComplianceReport
	return emptyResponse
}

// Function to return ClusterComplianceReport resource by uid.
func (m ResourceListClustercompliancereports) GetByUid(search string) ResourceClusterComplianceReport {
	for _, res := range m.Clustercompliancereports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceClusterComplianceReport
	return emptyResponse
}

// Function to return ClusterComplianceReport resource by label.
func (m ResourceListClustercompliancereports) GetByLabels(search map[string]string) []ResourceClusterComplianceReport {
	var Response []ResourceClusterComplianceReport
	for _, res := range m.Clustercompliancereports {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return ClusterVulnerabilityReport resource by name.
func (m ResourceListClustervulnerabilityreports) GetByName(search string) ResourceClusterVulnerabilityReport {
	for _, resource := range m.Clustervulnerabilityreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceClusterVulnerabilityReport
	return emptyResponse
}

// Function to return ClusterVulnerabilityReport resource by uid.
func (m ResourceListClustervulnerabilityreports) GetByUid(search string) ResourceClusterVulnerabilityReport {
	for _, res := range m.Clustervulnerabilityreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceClusterVulnerabilityReport
	return emptyResponse
}

// Function to return ClusterVulnerabilityReport resource by label.
func (m ResourceListClustervulnerabilityreports) GetByLabels(search map[string]string) []ResourceClusterVulnerabilityReport {
	var Response []ResourceClusterVulnerabilityReport
	for _, res := range m.Clustervulnerabilityreports {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Route resource by name.
func (m ResourceListRoutes) GetByName(search string) ResourceRoute {
	for _, resource := range m.Routes {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceRoute
	return emptyResponse
}

// Function to return Route resource by uid.
func (m ResourceListRoutes) GetByUid(search string) ResourceRoute {
	for _, res := range m.Routes {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceRoute
	return emptyResponse
}

// Function to return Route resource by label.
func (m ResourceListRoutes) GetByLabels(search map[string]string) []ResourceRoute {
	var Response []ResourceRoute
	for _, res := range m.Routes {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return SlackMessage resource by name.
func (m ResourceListSlackmessages) GetByName(search string) ResourceSlackMessage {
	for _, resource := range m.Slackmessages {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceSlackMessage
	return emptyResponse
}

// Function to return SlackMessage resource by uid.
func (m ResourceListSlackmessages) GetByUid(search string) ResourceSlackMessage {
	for _, res := range m.Slackmessages {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceSlackMessage
	return emptyResponse
}

// Function to return SlackMessage resource by label.
func (m ResourceListSlackmessages) GetByLabels(search map[string]string) []ResourceSlackMessage {
	var Response []ResourceSlackMessage
	for _, res := range m.Slackmessages {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return VulnerabilityEvent resource by name.
func (m ResourceListVulnerabilityevents) GetByName(search string) ResourceVulnerabilityEvent {
	for _, resource := range m.Vulnerabilityevents {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceVulnerabilityEvent
	return emptyResponse
}

// Function to return VulnerabilityEvent resource by uid.
func (m ResourceListVulnerabilityevents) GetByUid(search string) ResourceVulnerabilityEvent {
	for _, res := range m.Vulnerabilityevents {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceVulnerabilityEvent
	return emptyResponse
}

// Function to return VulnerabilityEvent resource by label.
func (m ResourceListVulnerabilityevents) GetByLabels(search map[string]string) []ResourceVulnerabilityEvent {
	var Response []ResourceVulnerabilityEvent
	for _, res := range m.Vulnerabilityevents {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return VirtualMachine resource by name.
func (m ResourceListVirtualmachines) GetByName(search string) ResourceVirtualMachine {
	for _, resource := range m.Virtualmachines {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceVirtualMachine
	return emptyResponse
}

// Function to return VirtualMachine resource by uid.
func (m ResourceListVirtualmachines) GetByUid(search string) ResourceVirtualMachine {
	for _, res := range m.Virtualmachines {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceVirtualMachine
	return emptyResponse
}

// Function to return VirtualMachine resource by label.
func (m ResourceListVirtualmachines) GetByLabels(search map[string]string) []ResourceVirtualMachine {
	var Response []ResourceVirtualMachine
	for _, res := range m.Virtualmachines {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return Endpoints resource by name.
func (m ResourceListEndpoints) GetByName(search string) ResourceEndpoints {
	for _, resource := range m.Endpoints {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceEndpoints
	return emptyResponse
}

// Function to return Endpoints resource by namespace.
func (m ResourceListEndpoints) GetByNamespace(search string) ResourceEndpoints {
	for _, res := range m.Endpoints {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceEndpoints
	return emptyResponse
}

// Function to return Endpoints resource by uid.
func (m ResourceListEndpoints) GetByUid(search string) ResourceEndpoints {
	for _, res := range m.Endpoints {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceEndpoints
	return emptyResponse
}

// Function to return Endpoints resource by label.
func (m ResourceListEndpoints) GetByLabels(search map[string]string) []ResourceEndpoints {
	var Response []ResourceEndpoints
	for _, res := range m.Endpoints {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return NetworkPolicy resource by name.
func (m ResourceListNetworkpolicies) GetByName(search string) ResourceNetworkPolicy {
	for _, resource := range m.Networkpolicies {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceNetworkPolicy
	return emptyResponse
}

// Function to return NetworkPolicy resource by namespace.
func (m ResourceListNetworkpolicies) GetByNamespace(search string) ResourceNetworkPolicy {
	for _, res := range m.Networkpolicies {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceNetworkPolicy
	return emptyResponse
}

// Function to return NetworkPolicy resource by uid.
func (m ResourceListNetworkpolicies) GetByUid(search string) ResourceNetworkPolicy {
	for _, res := range m.Networkpolicies {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceNetworkPolicy
	return emptyResponse
}

// Function to return NetworkPolicy resource by label.
func (m ResourceListNetworkpolicies) GetByLabels(search map[string]string) []ResourceNetworkPolicy {
	var Response []ResourceNetworkPolicy
	for _, res := range m.Networkpolicies {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return BackupJob resource by name.
func (m ResourceListBackupjobs) GetByName(search string) ResourceBackupJob {
	for _, resource := range m.Backupjobs {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceBackupJob
	return emptyResponse
}

// Function to return BackupJob resource by uid.
func (m ResourceListBackupjobs) GetByUid(search string) ResourceBackupJob {
	for _, res := range m.Backupjobs {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceBackupJob
	return emptyResponse
}

// Function to return BackupJob resource by label.
func (m ResourceListBackupjobs) GetByLabels(search map[string]string) []ResourceBackupJob {
	var Response []ResourceBackupJob
	for _, res := range m.Backupjobs {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}
