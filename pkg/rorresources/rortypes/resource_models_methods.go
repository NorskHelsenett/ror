// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rortypes

import (
	"fmt"

	"github.com/mitchellh/hashstructure/v2"
)

// (r *ResourceNamespace) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceNamespace) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceNamespace) Get returns a pointer to the resource of type ResourceNamespace
func (r *ResourceNamespace) Get() *ResourceNamespace {
	return r
}

// (r *ResourceNode) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceNode) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceNode) Get returns a pointer to the resource of type ResourceNode
func (r *ResourceNode) Get() *ResourceNode {
	return r
}

// (r *ResourcePersistentVolumeClaim) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourcePersistentVolumeClaim) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourcePersistentVolumeClaim) Get returns a pointer to the resource of type ResourcePersistentVolumeClaim
func (r *ResourcePersistentVolumeClaim) Get() *ResourcePersistentVolumeClaim {
	return r
}

// (r *ResourceDeployment) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceDeployment) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceDeployment) Get returns a pointer to the resource of type ResourceDeployment
func (r *ResourceDeployment) Get() *ResourceDeployment {
	return r
}

// (r *ResourceStorageClass) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceStorageClass) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceStorageClass) Get returns a pointer to the resource of type ResourceStorageClass
func (r *ResourceStorageClass) Get() *ResourceStorageClass {
	return r
}

// (r *ResourcePolicyReport) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourcePolicyReport) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourcePolicyReport) Get returns a pointer to the resource of type ResourcePolicyReport
func (r *ResourcePolicyReport) Get() *ResourcePolicyReport {
	return r
}

// (r *ResourceApplication) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceApplication) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceApplication) Get returns a pointer to the resource of type ResourceApplication
func (r *ResourceApplication) Get() *ResourceApplication {
	return r
}

// (r *ResourceAppProject) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceAppProject) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceAppProject) Get returns a pointer to the resource of type ResourceAppProject
func (r *ResourceAppProject) Get() *ResourceAppProject {
	return r
}

// (r *ResourceCertificate) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceCertificate) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceCertificate) Get returns a pointer to the resource of type ResourceCertificate
func (r *ResourceCertificate) Get() *ResourceCertificate {
	return r
}

// (r *ResourceService) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceService) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceService) Get returns a pointer to the resource of type ResourceService
func (r *ResourceService) Get() *ResourceService {
	return r
}

// (r *ResourcePod) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourcePod) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourcePod) Get returns a pointer to the resource of type ResourcePod
func (r *ResourcePod) Get() *ResourcePod {
	return r
}

// (r *ResourceReplicaSet) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceReplicaSet) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceReplicaSet) Get returns a pointer to the resource of type ResourceReplicaSet
func (r *ResourceReplicaSet) Get() *ResourceReplicaSet {
	return r
}

// (r *ResourceStatefulSet) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceStatefulSet) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceStatefulSet) Get returns a pointer to the resource of type ResourceStatefulSet
func (r *ResourceStatefulSet) Get() *ResourceStatefulSet {
	return r
}

// (r *ResourceDaemonSet) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceDaemonSet) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceDaemonSet) Get returns a pointer to the resource of type ResourceDaemonSet
func (r *ResourceDaemonSet) Get() *ResourceDaemonSet {
	return r
}

// (r *ResourceIngress) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceIngress) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceIngress) Get returns a pointer to the resource of type ResourceIngress
func (r *ResourceIngress) Get() *ResourceIngress {
	return r
}

// (r *ResourceIngressClass) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceIngressClass) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceIngressClass) Get returns a pointer to the resource of type ResourceIngressClass
func (r *ResourceIngressClass) Get() *ResourceIngressClass {
	return r
}

// (r *ResourceVulnerabilityReport) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceVulnerabilityReport) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceVulnerabilityReport) Get returns a pointer to the resource of type ResourceVulnerabilityReport
func (r *ResourceVulnerabilityReport) Get() *ResourceVulnerabilityReport {
	return r
}

// (r *ResourceExposedSecretReport) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceExposedSecretReport) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceExposedSecretReport) Get returns a pointer to the resource of type ResourceExposedSecretReport
func (r *ResourceExposedSecretReport) Get() *ResourceExposedSecretReport {
	return r
}

// (r *ResourceConfigAuditReport) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceConfigAuditReport) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceConfigAuditReport) Get returns a pointer to the resource of type ResourceConfigAuditReport
func (r *ResourceConfigAuditReport) Get() *ResourceConfigAuditReport {
	return r
}

// (r *ResourceRbacAssessmentReport) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceRbacAssessmentReport) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceRbacAssessmentReport) Get returns a pointer to the resource of type ResourceRbacAssessmentReport
func (r *ResourceRbacAssessmentReport) Get() *ResourceRbacAssessmentReport {
	return r
}

// (r *ResourceTanzuKubernetesCluster) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceTanzuKubernetesCluster) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceTanzuKubernetesCluster) Get returns a pointer to the resource of type ResourceTanzuKubernetesCluster
func (r *ResourceTanzuKubernetesCluster) Get() *ResourceTanzuKubernetesCluster {
	return r
}

// (r *ResourceTanzuKubernetesRelease) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceTanzuKubernetesRelease) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceTanzuKubernetesRelease) Get returns a pointer to the resource of type ResourceTanzuKubernetesRelease
func (r *ResourceTanzuKubernetesRelease) Get() *ResourceTanzuKubernetesRelease {
	return r
}

// (r *ResourceVirtualMachineClass) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceVirtualMachineClass) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceVirtualMachineClass) Get returns a pointer to the resource of type ResourceVirtualMachineClass
func (r *ResourceVirtualMachineClass) Get() *ResourceVirtualMachineClass {
	return r
}

// (r *ResourceVirtualMachineClassBinding) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceVirtualMachineClassBinding) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceVirtualMachineClassBinding) Get returns a pointer to the resource of type ResourceVirtualMachineClassBinding
func (r *ResourceVirtualMachineClassBinding) Get() *ResourceVirtualMachineClassBinding {
	return r
}

// (r *ResourceKubernetesCluster) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceKubernetesCluster) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceKubernetesCluster) Get returns a pointer to the resource of type ResourceKubernetesCluster
func (r *ResourceKubernetesCluster) Get() *ResourceKubernetesCluster {
	return r
}

// (r *ResourceClusterOrder) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceClusterOrder) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceClusterOrder) Get returns a pointer to the resource of type ResourceClusterOrder
func (r *ResourceClusterOrder) Get() *ResourceClusterOrder {
	return r
}

// (r *ResourceProject) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceProject) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceProject) Get returns a pointer to the resource of type ResourceProject
func (r *ResourceProject) Get() *ResourceProject {
	return r
}

// (r *ResourceConfiguration) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceConfiguration) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceConfiguration) Get returns a pointer to the resource of type ResourceConfiguration
func (r *ResourceConfiguration) Get() *ResourceConfiguration {
	return r
}

// (r *ResourceClusterComplianceReport) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceClusterComplianceReport) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceClusterComplianceReport) Get returns a pointer to the resource of type ResourceClusterComplianceReport
func (r *ResourceClusterComplianceReport) Get() *ResourceClusterComplianceReport {
	return r
}

// (r *ResourceClusterVulnerabilityReport) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceClusterVulnerabilityReport) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceClusterVulnerabilityReport) Get returns a pointer to the resource of type ResourceClusterVulnerabilityReport
func (r *ResourceClusterVulnerabilityReport) Get() *ResourceClusterVulnerabilityReport {
	return r
}

// (r *ResourceRoute) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceRoute) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceRoute) Get returns a pointer to the resource of type ResourceRoute
func (r *ResourceRoute) Get() *ResourceRoute {
	return r
}

// (r *ResourceSlackMessage) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceSlackMessage) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceSlackMessage) Get returns a pointer to the resource of type ResourceSlackMessage
func (r *ResourceSlackMessage) Get() *ResourceSlackMessage {
	return r
}

// (r *ResourceVulnerabilityEvent) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceVulnerabilityEvent) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceVulnerabilityEvent) Get returns a pointer to the resource of type ResourceVulnerabilityEvent
func (r *ResourceVulnerabilityEvent) Get() *ResourceVulnerabilityEvent {
	return r
}

// (r *ResourceVirtualMachine) GetRorHash calculates the hash of the resource
//
// it uses the hashstructure library to calculate the hash of the resource
// fields can be ignored by adding the tag `hash:"ignore"` to the field
func (r *ResourceVirtualMachine) GetRorHash() string {
	hash, err := hashstructure.Hash(r, hashstructure.FormatV2, nil)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", hash)
}

// (r ResourceVirtualMachine) Get returns a pointer to the resource of type ResourceVirtualMachine
func (r *ResourceVirtualMachine) Get() *ResourceVirtualMachine {
	return r
}
