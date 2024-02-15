// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package apiresourcecontracts

import (
	"github.com/NorskHelsenett/ror/pkg/helpers/stringhelper"
)

// Function to return Namespace resource by name.
func (m ResourceNamespaces) GetByName(search string) ResourceNamespace {
	for _, resource := range m.Namespaces {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceNamespace
	return emptyResponse
}

// Function to return Namespace resource by uid.
func (m ResourceNamespaces) GetByUid(search string) ResourceNamespace {
	for _, res := range m.Namespaces {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceNamespace
	return emptyResponse
}

// Function to return Namespace resource by label.
func (m ResourceNamespaces) GetByLabels(search map[string]string) []ResourceNamespace {
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
func (m ResourceNodes) GetByName(search string) ResourceNode {
	for _, resource := range m.Nodes {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceNode
	return emptyResponse
}

// Function to return Node resource by uid.
func (m ResourceNodes) GetByUid(search string) ResourceNode {
	for _, res := range m.Nodes {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceNode
	return emptyResponse
}

// Function to return Node resource by label.
func (m ResourceNodes) GetByLabels(search map[string]string) []ResourceNode {
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
func (m ResourcePersistentvolumeclaims) GetByName(search string) ResourcePersistentVolumeClaim {
	for _, resource := range m.Persistentvolumeclaims {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourcePersistentVolumeClaim
	return emptyResponse
}

// Function to return PersistentVolumeClaim resource by namespace.
func (m ResourcePersistentvolumeclaims) GetByNamespace(search string) ResourcePersistentVolumeClaim {
	for _, res := range m.Persistentvolumeclaims {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourcePersistentVolumeClaim
	return emptyResponse
}

// Function to return PersistentVolumeClaim resource by uid.
func (m ResourcePersistentvolumeclaims) GetByUid(search string) ResourcePersistentVolumeClaim {
	for _, res := range m.Persistentvolumeclaims {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourcePersistentVolumeClaim
	return emptyResponse
}

// Function to return PersistentVolumeClaim resource by label.
func (m ResourcePersistentvolumeclaims) GetByLabels(search map[string]string) []ResourcePersistentVolumeClaim {
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
func (m ResourceDeployments) GetByName(search string) ResourceDeployment {
	for _, resource := range m.Deployments {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceDeployment
	return emptyResponse
}

// Function to return Deployment resource by namespace.
func (m ResourceDeployments) GetByNamespace(search string) ResourceDeployment {
	for _, res := range m.Deployments {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceDeployment
	return emptyResponse
}

// Function to return Deployment resource by uid.
func (m ResourceDeployments) GetByUid(search string) ResourceDeployment {
	for _, res := range m.Deployments {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceDeployment
	return emptyResponse
}

// Function to return Deployment resource by label.
func (m ResourceDeployments) GetByLabels(search map[string]string) []ResourceDeployment {
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
func (m ResourceStorageclasses) GetByName(search string) ResourceStorageClass {
	for _, resource := range m.Storageclasses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceStorageClass
	return emptyResponse
}

// Function to return StorageClass resource by uid.
func (m ResourceStorageclasses) GetByUid(search string) ResourceStorageClass {
	for _, res := range m.Storageclasses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceStorageClass
	return emptyResponse
}

// Function to return StorageClass resource by label.
func (m ResourceStorageclasses) GetByLabels(search map[string]string) []ResourceStorageClass {
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
func (m ResourcePolicyreports) GetByName(search string) ResourcePolicyReport {
	for _, resource := range m.Policyreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourcePolicyReport
	return emptyResponse
}

// Function to return PolicyReport resource by namespace.
func (m ResourcePolicyreports) GetByNamespace(search string) ResourcePolicyReport {
	for _, res := range m.Policyreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourcePolicyReport
	return emptyResponse
}

// Function to return PolicyReport resource by uid.
func (m ResourcePolicyreports) GetByUid(search string) ResourcePolicyReport {
	for _, res := range m.Policyreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourcePolicyReport
	return emptyResponse
}

// Function to return PolicyReport resource by label.
func (m ResourcePolicyreports) GetByLabels(search map[string]string) []ResourcePolicyReport {
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
func (m ResourceApplications) GetByName(search string) ResourceApplication {
	for _, resource := range m.Applications {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceApplication
	return emptyResponse
}

// Function to return Application resource by namespace.
func (m ResourceApplications) GetByNamespace(search string) ResourceApplication {
	for _, res := range m.Applications {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceApplication
	return emptyResponse
}

// Function to return Application resource by uid.
func (m ResourceApplications) GetByUid(search string) ResourceApplication {
	for _, res := range m.Applications {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceApplication
	return emptyResponse
}

// Function to return Application resource by label.
func (m ResourceApplications) GetByLabels(search map[string]string) []ResourceApplication {
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
func (m ResourceAppprojects) GetByName(search string) ResourceAppProject {
	for _, resource := range m.Appprojects {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceAppProject
	return emptyResponse
}

// Function to return AppProject resource by namespace.
func (m ResourceAppprojects) GetByNamespace(search string) ResourceAppProject {
	for _, res := range m.Appprojects {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceAppProject
	return emptyResponse
}

// Function to return AppProject resource by uid.
func (m ResourceAppprojects) GetByUid(search string) ResourceAppProject {
	for _, res := range m.Appprojects {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceAppProject
	return emptyResponse
}

// Function to return AppProject resource by label.
func (m ResourceAppprojects) GetByLabels(search map[string]string) []ResourceAppProject {
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
func (m ResourceCertificates) GetByName(search string) ResourceCertificate {
	for _, resource := range m.Certificates {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceCertificate
	return emptyResponse
}

// Function to return Certificate resource by namespace.
func (m ResourceCertificates) GetByNamespace(search string) ResourceCertificate {
	for _, res := range m.Certificates {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceCertificate
	return emptyResponse
}

// Function to return Certificate resource by uid.
func (m ResourceCertificates) GetByUid(search string) ResourceCertificate {
	for _, res := range m.Certificates {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceCertificate
	return emptyResponse
}

// Function to return Certificate resource by label.
func (m ResourceCertificates) GetByLabels(search map[string]string) []ResourceCertificate {
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
func (m ResourceServices) GetByName(search string) ResourceService {
	for _, resource := range m.Services {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceService
	return emptyResponse
}

// Function to return Service resource by namespace.
func (m ResourceServices) GetByNamespace(search string) ResourceService {
	for _, res := range m.Services {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceService
	return emptyResponse
}

// Function to return Service resource by uid.
func (m ResourceServices) GetByUid(search string) ResourceService {
	for _, res := range m.Services {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceService
	return emptyResponse
}

// Function to return Service resource by label.
func (m ResourceServices) GetByLabels(search map[string]string) []ResourceService {
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
func (m ResourcePods) GetByName(search string) ResourcePod {
	for _, resource := range m.Pods {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourcePod
	return emptyResponse
}

// Function to return Pod resource by namespace.
func (m ResourcePods) GetByNamespace(search string) ResourcePod {
	for _, res := range m.Pods {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourcePod
	return emptyResponse
}

// Function to return Pod resource by uid.
func (m ResourcePods) GetByUid(search string) ResourcePod {
	for _, res := range m.Pods {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourcePod
	return emptyResponse
}

// Function to return Pod resource by label.
func (m ResourcePods) GetByLabels(search map[string]string) []ResourcePod {
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
func (m ResourceReplicasets) GetByName(search string) ResourceReplicaSet {
	for _, resource := range m.Replicasets {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceReplicaSet
	return emptyResponse
}

// Function to return ReplicaSet resource by namespace.
func (m ResourceReplicasets) GetByNamespace(search string) ResourceReplicaSet {
	for _, res := range m.Replicasets {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceReplicaSet
	return emptyResponse
}

// Function to return ReplicaSet resource by uid.
func (m ResourceReplicasets) GetByUid(search string) ResourceReplicaSet {
	for _, res := range m.Replicasets {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceReplicaSet
	return emptyResponse
}

// Function to return ReplicaSet resource by label.
func (m ResourceReplicasets) GetByLabels(search map[string]string) []ResourceReplicaSet {
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
func (m ResourceStatefulsets) GetByName(search string) ResourceStatefulSet {
	for _, resource := range m.Statefulsets {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceStatefulSet
	return emptyResponse
}

// Function to return StatefulSet resource by namespace.
func (m ResourceStatefulsets) GetByNamespace(search string) ResourceStatefulSet {
	for _, res := range m.Statefulsets {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceStatefulSet
	return emptyResponse
}

// Function to return StatefulSet resource by uid.
func (m ResourceStatefulsets) GetByUid(search string) ResourceStatefulSet {
	for _, res := range m.Statefulsets {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceStatefulSet
	return emptyResponse
}

// Function to return StatefulSet resource by label.
func (m ResourceStatefulsets) GetByLabels(search map[string]string) []ResourceStatefulSet {
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
func (m ResourceDaemonsets) GetByName(search string) ResourceDaemonSet {
	for _, resource := range m.Daemonsets {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceDaemonSet
	return emptyResponse
}

// Function to return DaemonSet resource by namespace.
func (m ResourceDaemonsets) GetByNamespace(search string) ResourceDaemonSet {
	for _, res := range m.Daemonsets {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceDaemonSet
	return emptyResponse
}

// Function to return DaemonSet resource by uid.
func (m ResourceDaemonsets) GetByUid(search string) ResourceDaemonSet {
	for _, res := range m.Daemonsets {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceDaemonSet
	return emptyResponse
}

// Function to return DaemonSet resource by label.
func (m ResourceDaemonsets) GetByLabels(search map[string]string) []ResourceDaemonSet {
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
func (m ResourceIngresses) GetByName(search string) ResourceIngress {
	for _, resource := range m.Ingresses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceIngress
	return emptyResponse
}

// Function to return Ingress resource by namespace.
func (m ResourceIngresses) GetByNamespace(search string) ResourceIngress {
	for _, res := range m.Ingresses {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceIngress
	return emptyResponse
}

// Function to return Ingress resource by uid.
func (m ResourceIngresses) GetByUid(search string) ResourceIngress {
	for _, res := range m.Ingresses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceIngress
	return emptyResponse
}

// Function to return Ingress resource by label.
func (m ResourceIngresses) GetByLabels(search map[string]string) []ResourceIngress {
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
func (m ResourceIngressclasses) GetByName(search string) ResourceIngressClass {
	for _, resource := range m.Ingressclasses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceIngressClass
	return emptyResponse
}

// Function to return IngressClass resource by uid.
func (m ResourceIngressclasses) GetByUid(search string) ResourceIngressClass {
	for _, res := range m.Ingressclasses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceIngressClass
	return emptyResponse
}

// Function to return IngressClass resource by label.
func (m ResourceIngressclasses) GetByLabels(search map[string]string) []ResourceIngressClass {
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
func (m ResourceVulnerabilityreports) GetByName(search string) ResourceVulnerabilityReport {
	for _, resource := range m.Vulnerabilityreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceVulnerabilityReport
	return emptyResponse
}

// Function to return VulnerabilityReport resource by namespace.
func (m ResourceVulnerabilityreports) GetByNamespace(search string) ResourceVulnerabilityReport {
	for _, res := range m.Vulnerabilityreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceVulnerabilityReport
	return emptyResponse
}

// Function to return VulnerabilityReport resource by uid.
func (m ResourceVulnerabilityreports) GetByUid(search string) ResourceVulnerabilityReport {
	for _, res := range m.Vulnerabilityreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceVulnerabilityReport
	return emptyResponse
}

// Function to return VulnerabilityReport resource by label.
func (m ResourceVulnerabilityreports) GetByLabels(search map[string]string) []ResourceVulnerabilityReport {
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
func (m ResourceExposedsecretreports) GetByName(search string) ResourceExposedSecretReport {
	for _, resource := range m.Exposedsecretreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceExposedSecretReport
	return emptyResponse
}

// Function to return ExposedSecretReport resource by namespace.
func (m ResourceExposedsecretreports) GetByNamespace(search string) ResourceExposedSecretReport {
	for _, res := range m.Exposedsecretreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceExposedSecretReport
	return emptyResponse
}

// Function to return ExposedSecretReport resource by uid.
func (m ResourceExposedsecretreports) GetByUid(search string) ResourceExposedSecretReport {
	for _, res := range m.Exposedsecretreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceExposedSecretReport
	return emptyResponse
}

// Function to return ExposedSecretReport resource by label.
func (m ResourceExposedsecretreports) GetByLabels(search map[string]string) []ResourceExposedSecretReport {
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
func (m ResourceConfigauditreports) GetByName(search string) ResourceConfigAuditReport {
	for _, resource := range m.Configauditreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceConfigAuditReport
	return emptyResponse
}

// Function to return ConfigAuditReport resource by namespace.
func (m ResourceConfigauditreports) GetByNamespace(search string) ResourceConfigAuditReport {
	for _, res := range m.Configauditreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceConfigAuditReport
	return emptyResponse
}

// Function to return ConfigAuditReport resource by uid.
func (m ResourceConfigauditreports) GetByUid(search string) ResourceConfigAuditReport {
	for _, res := range m.Configauditreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceConfigAuditReport
	return emptyResponse
}

// Function to return ConfigAuditReport resource by label.
func (m ResourceConfigauditreports) GetByLabels(search map[string]string) []ResourceConfigAuditReport {
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
func (m ResourceRbacassessmentreports) GetByName(search string) ResourceRbacAssessmentReport {
	for _, resource := range m.Rbacassessmentreports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceRbacAssessmentReport
	return emptyResponse
}

// Function to return RbacAssessmentReport resource by namespace.
func (m ResourceRbacassessmentreports) GetByNamespace(search string) ResourceRbacAssessmentReport {
	for _, res := range m.Rbacassessmentreports {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceRbacAssessmentReport
	return emptyResponse
}

// Function to return RbacAssessmentReport resource by uid.
func (m ResourceRbacassessmentreports) GetByUid(search string) ResourceRbacAssessmentReport {
	for _, res := range m.Rbacassessmentreports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceRbacAssessmentReport
	return emptyResponse
}

// Function to return RbacAssessmentReport resource by label.
func (m ResourceRbacassessmentreports) GetByLabels(search map[string]string) []ResourceRbacAssessmentReport {
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
func (m ResourceTanzukubernetesclusters) GetByName(search string) ResourceTanzuKubernetesCluster {
	for _, resource := range m.Tanzukubernetesclusters {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceTanzuKubernetesCluster
	return emptyResponse
}

// Function to return TanzuKubernetesCluster resource by namespace.
func (m ResourceTanzukubernetesclusters) GetByNamespace(search string) ResourceTanzuKubernetesCluster {
	for _, res := range m.Tanzukubernetesclusters {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceTanzuKubernetesCluster
	return emptyResponse
}

// Function to return TanzuKubernetesCluster resource by uid.
func (m ResourceTanzukubernetesclusters) GetByUid(search string) ResourceTanzuKubernetesCluster {
	for _, res := range m.Tanzukubernetesclusters {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceTanzuKubernetesCluster
	return emptyResponse
}

// Function to return TanzuKubernetesCluster resource by label.
func (m ResourceTanzukubernetesclusters) GetByLabels(search map[string]string) []ResourceTanzuKubernetesCluster {
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
func (m ResourceTanzukubernetesreleases) GetByName(search string) ResourceTanzuKubernetesRelease {
	for _, resource := range m.Tanzukubernetesreleases {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceTanzuKubernetesRelease
	return emptyResponse
}

// Function to return TanzuKubernetesRelease resource by uid.
func (m ResourceTanzukubernetesreleases) GetByUid(search string) ResourceTanzuKubernetesRelease {
	for _, res := range m.Tanzukubernetesreleases {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceTanzuKubernetesRelease
	return emptyResponse
}

// Function to return TanzuKubernetesRelease resource by label.
func (m ResourceTanzukubernetesreleases) GetByLabels(search map[string]string) []ResourceTanzuKubernetesRelease {
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
func (m ResourceVirtualmachineclasses) GetByName(search string) ResourceVirtualMachineClass {
	for _, resource := range m.Virtualmachineclasses {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceVirtualMachineClass
	return emptyResponse
}

// Function to return VirtualMachineClass resource by uid.
func (m ResourceVirtualmachineclasses) GetByUid(search string) ResourceVirtualMachineClass {
	for _, res := range m.Virtualmachineclasses {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceVirtualMachineClass
	return emptyResponse
}

// Function to return VirtualMachineClass resource by label.
func (m ResourceVirtualmachineclasses) GetByLabels(search map[string]string) []ResourceVirtualMachineClass {
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

// Function to return VirtualMachineClassBinding resource by name.
func (m ResourceVirtualmachineclassbindings) GetByName(search string) ResourceVirtualMachineClassBinding {
	for _, resource := range m.Virtualmachineclassbindings {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceVirtualMachineClassBinding
	return emptyResponse
}

// Function to return VirtualMachineClassBinding resource by namespace.
func (m ResourceVirtualmachineclassbindings) GetByNamespace(search string) ResourceVirtualMachineClassBinding {
	for _, res := range m.Virtualmachineclassbindings {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceVirtualMachineClassBinding
	return emptyResponse
}

// Function to return VirtualMachineClassBinding resource by uid.
func (m ResourceVirtualmachineclassbindings) GetByUid(search string) ResourceVirtualMachineClassBinding {
	for _, res := range m.Virtualmachineclassbindings {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceVirtualMachineClassBinding
	return emptyResponse
}

// Function to return VirtualMachineClassBinding resource by label.
func (m ResourceVirtualmachineclassbindings) GetByLabels(search map[string]string) []ResourceVirtualMachineClassBinding {
	var Response []ResourceVirtualMachineClassBinding
	for _, res := range m.Virtualmachineclassbindings {
		if len(res.Metadata.Labels) != 0 {
			if stringhelper.CompareLabels(search, res.Metadata.Labels) {
				Response = append(Response, res)
			}
		}
	}
	return Response
}

// Function to return KubernetesCluster resource by name.
func (m ResourceKubernetesclusters) GetByName(search string) ResourceKubernetesCluster {
	for _, resource := range m.Kubernetesclusters {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceKubernetesCluster
	return emptyResponse
}

// Function to return KubernetesCluster resource by namespace.
func (m ResourceKubernetesclusters) GetByNamespace(search string) ResourceKubernetesCluster {
	for _, res := range m.Kubernetesclusters {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceKubernetesCluster
	return emptyResponse
}

// Function to return KubernetesCluster resource by uid.
func (m ResourceKubernetesclusters) GetByUid(search string) ResourceKubernetesCluster {
	for _, res := range m.Kubernetesclusters {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceKubernetesCluster
	return emptyResponse
}

// Function to return KubernetesCluster resource by label.
func (m ResourceKubernetesclusters) GetByLabels(search map[string]string) []ResourceKubernetesCluster {
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
func (m ResourceClusterorders) GetByName(search string) ResourceClusterOrder {
	for _, resource := range m.Clusterorders {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceClusterOrder
	return emptyResponse
}

// Function to return ClusterOrder resource by namespace.
func (m ResourceClusterorders) GetByNamespace(search string) ResourceClusterOrder {
	for _, res := range m.Clusterorders {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceClusterOrder
	return emptyResponse
}

// Function to return ClusterOrder resource by uid.
func (m ResourceClusterorders) GetByUid(search string) ResourceClusterOrder {
	for _, res := range m.Clusterorders {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceClusterOrder
	return emptyResponse
}

// Function to return ClusterOrder resource by label.
func (m ResourceClusterorders) GetByLabels(search map[string]string) []ResourceClusterOrder {
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
func (m ResourceProjects) GetByName(search string) ResourceProject {
	for _, resource := range m.Projects {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceProject
	return emptyResponse
}

// Function to return Project resource by namespace.
func (m ResourceProjects) GetByNamespace(search string) ResourceProject {
	for _, res := range m.Projects {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceProject
	return emptyResponse
}

// Function to return Project resource by uid.
func (m ResourceProjects) GetByUid(search string) ResourceProject {
	for _, res := range m.Projects {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceProject
	return emptyResponse
}

// Function to return Project resource by label.
func (m ResourceProjects) GetByLabels(search map[string]string) []ResourceProject {
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
func (m ResourceConfigurations) GetByName(search string) ResourceConfiguration {
	for _, resource := range m.Configurations {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceConfiguration
	return emptyResponse
}

// Function to return Configuration resource by namespace.
func (m ResourceConfigurations) GetByNamespace(search string) ResourceConfiguration {
	for _, res := range m.Configurations {
		if res.Metadata.Namespace == search {
			return res
		}
	}
	var emptyResponse ResourceConfiguration
	return emptyResponse
}

// Function to return Configuration resource by uid.
func (m ResourceConfigurations) GetByUid(search string) ResourceConfiguration {
	for _, res := range m.Configurations {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceConfiguration
	return emptyResponse
}

// Function to return Configuration resource by label.
func (m ResourceConfigurations) GetByLabels(search map[string]string) []ResourceConfiguration {
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
func (m ResourceClustercompliancereports) GetByName(search string) ResourceClusterComplianceReport {
	for _, resource := range m.Clustercompliancereports {
		if resource.Metadata.Name == search {
			return resource
		}
	}
	var emptyResponse ResourceClusterComplianceReport
	return emptyResponse
}

// Function to return ClusterComplianceReport resource by uid.
func (m ResourceClustercompliancereports) GetByUid(search string) ResourceClusterComplianceReport {
	for _, res := range m.Clustercompliancereports {
		if res.Metadata.Uid == search {
			return res
		}
	}
	var emptyResponse ResourceClusterComplianceReport
	return emptyResponse
}

// Function to return ClusterComplianceReport resource by label.
func (m ResourceClustercompliancereports) GetByLabels(search map[string]string) []ResourceClusterComplianceReport {
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
