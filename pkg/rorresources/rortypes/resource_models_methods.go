// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package rortypes

import (
	"fmt"

	"github.com/mitchellh/hashstructure/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// (r ResourceNamespace) GetName returns the name of the resource
func (r ResourceNamespace) GetName() string {
	return r.Metadata.Name
}

// (r ResourceNamespace) GetUID returns the UID of the resource
func (r ResourceNamespace) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceNamespace) GetAPIVersion returns the APIVersion of the resource
func (r ResourceNamespace) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceNamespace) GetKind returns the kind of the resource
func (r ResourceNamespace) GetKind() string {
	return string(r.Kind)
}

// (r ResourceNamespace) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceNamespace) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceNamespace) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceNamespace) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceNamespace) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceNamespace) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
}

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

// (r ResourceNode) GetName returns the name of the resource
func (r ResourceNode) GetName() string {
	return r.Metadata.Name
}

// (r ResourceNode) GetUID returns the UID of the resource
func (r ResourceNode) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceNode) GetAPIVersion returns the APIVersion of the resource
func (r ResourceNode) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceNode) GetKind returns the kind of the resource
func (r ResourceNode) GetKind() string {
	return string(r.Kind)
}

// (r ResourceNode) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceNode) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceNode) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceNode) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceNode) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceNode) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourcePersistentVolumeClaim) GetName returns the name of the resource
func (r ResourcePersistentVolumeClaim) GetName() string {
	return r.Metadata.Name
}

// (r ResourcePersistentVolumeClaim) GetUID returns the UID of the resource
func (r ResourcePersistentVolumeClaim) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourcePersistentVolumeClaim) GetAPIVersion returns the APIVersion of the resource
func (r ResourcePersistentVolumeClaim) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourcePersistentVolumeClaim) GetKind returns the kind of the resource
func (r ResourcePersistentVolumeClaim) GetKind() string {
	return string(r.Kind)
}

// (r ResourcePersistentVolumeClaim) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourcePersistentVolumeClaim) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourcePersistentVolumeClaim) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourcePersistentVolumeClaim) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourcePersistentVolumeClaim) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourcePersistentVolumeClaim) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceDeployment) GetName returns the name of the resource
func (r ResourceDeployment) GetName() string {
	return r.Metadata.Name
}

// (r ResourceDeployment) GetUID returns the UID of the resource
func (r ResourceDeployment) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceDeployment) GetAPIVersion returns the APIVersion of the resource
func (r ResourceDeployment) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceDeployment) GetKind returns the kind of the resource
func (r ResourceDeployment) GetKind() string {
	return string(r.Kind)
}

// (r ResourceDeployment) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceDeployment) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceDeployment) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceDeployment) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceDeployment) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceDeployment) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceStorageClass) GetName returns the name of the resource
func (r ResourceStorageClass) GetName() string {
	return r.Metadata.Name
}

// (r ResourceStorageClass) GetUID returns the UID of the resource
func (r ResourceStorageClass) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceStorageClass) GetAPIVersion returns the APIVersion of the resource
func (r ResourceStorageClass) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceStorageClass) GetKind returns the kind of the resource
func (r ResourceStorageClass) GetKind() string {
	return string(r.Kind)
}

// (r ResourceStorageClass) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceStorageClass) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceStorageClass) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceStorageClass) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceStorageClass) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceStorageClass) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourcePolicyReport) GetName returns the name of the resource
func (r ResourcePolicyReport) GetName() string {
	return r.Metadata.Name
}

// (r ResourcePolicyReport) GetUID returns the UID of the resource
func (r ResourcePolicyReport) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourcePolicyReport) GetAPIVersion returns the APIVersion of the resource
func (r ResourcePolicyReport) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourcePolicyReport) GetKind returns the kind of the resource
func (r ResourcePolicyReport) GetKind() string {
	return string(r.Kind)
}

// (r ResourcePolicyReport) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourcePolicyReport) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourcePolicyReport) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourcePolicyReport) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourcePolicyReport) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourcePolicyReport) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceApplication) GetName returns the name of the resource
func (r ResourceApplication) GetName() string {
	return r.Metadata.Name
}

// (r ResourceApplication) GetUID returns the UID of the resource
func (r ResourceApplication) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceApplication) GetAPIVersion returns the APIVersion of the resource
func (r ResourceApplication) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceApplication) GetKind returns the kind of the resource
func (r ResourceApplication) GetKind() string {
	return string(r.Kind)
}

// (r ResourceApplication) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceApplication) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceApplication) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceApplication) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceApplication) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceApplication) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceAppProject) GetName returns the name of the resource
func (r ResourceAppProject) GetName() string {
	return r.Metadata.Name
}

// (r ResourceAppProject) GetUID returns the UID of the resource
func (r ResourceAppProject) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceAppProject) GetAPIVersion returns the APIVersion of the resource
func (r ResourceAppProject) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceAppProject) GetKind returns the kind of the resource
func (r ResourceAppProject) GetKind() string {
	return string(r.Kind)
}

// (r ResourceAppProject) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceAppProject) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceAppProject) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceAppProject) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceAppProject) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceAppProject) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceCertificate) GetName returns the name of the resource
func (r ResourceCertificate) GetName() string {
	return r.Metadata.Name
}

// (r ResourceCertificate) GetUID returns the UID of the resource
func (r ResourceCertificate) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceCertificate) GetAPIVersion returns the APIVersion of the resource
func (r ResourceCertificate) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceCertificate) GetKind returns the kind of the resource
func (r ResourceCertificate) GetKind() string {
	return string(r.Kind)
}

// (r ResourceCertificate) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceCertificate) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceCertificate) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceCertificate) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceCertificate) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceCertificate) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceService) GetName returns the name of the resource
func (r ResourceService) GetName() string {
	return r.Metadata.Name
}

// (r ResourceService) GetUID returns the UID of the resource
func (r ResourceService) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceService) GetAPIVersion returns the APIVersion of the resource
func (r ResourceService) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceService) GetKind returns the kind of the resource
func (r ResourceService) GetKind() string {
	return string(r.Kind)
}

// (r ResourceService) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceService) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceService) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceService) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceService) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceService) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourcePod) GetName returns the name of the resource
func (r ResourcePod) GetName() string {
	return r.Metadata.Name
}

// (r ResourcePod) GetUID returns the UID of the resource
func (r ResourcePod) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourcePod) GetAPIVersion returns the APIVersion of the resource
func (r ResourcePod) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourcePod) GetKind returns the kind of the resource
func (r ResourcePod) GetKind() string {
	return string(r.Kind)
}

// (r ResourcePod) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourcePod) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourcePod) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourcePod) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourcePod) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourcePod) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceReplicaSet) GetName returns the name of the resource
func (r ResourceReplicaSet) GetName() string {
	return r.Metadata.Name
}

// (r ResourceReplicaSet) GetUID returns the UID of the resource
func (r ResourceReplicaSet) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceReplicaSet) GetAPIVersion returns the APIVersion of the resource
func (r ResourceReplicaSet) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceReplicaSet) GetKind returns the kind of the resource
func (r ResourceReplicaSet) GetKind() string {
	return string(r.Kind)
}

// (r ResourceReplicaSet) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceReplicaSet) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceReplicaSet) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceReplicaSet) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceReplicaSet) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceReplicaSet) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceStatefulSet) GetName returns the name of the resource
func (r ResourceStatefulSet) GetName() string {
	return r.Metadata.Name
}

// (r ResourceStatefulSet) GetUID returns the UID of the resource
func (r ResourceStatefulSet) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceStatefulSet) GetAPIVersion returns the APIVersion of the resource
func (r ResourceStatefulSet) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceStatefulSet) GetKind returns the kind of the resource
func (r ResourceStatefulSet) GetKind() string {
	return string(r.Kind)
}

// (r ResourceStatefulSet) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceStatefulSet) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceStatefulSet) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceStatefulSet) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceStatefulSet) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceStatefulSet) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceDaemonSet) GetName returns the name of the resource
func (r ResourceDaemonSet) GetName() string {
	return r.Metadata.Name
}

// (r ResourceDaemonSet) GetUID returns the UID of the resource
func (r ResourceDaemonSet) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceDaemonSet) GetAPIVersion returns the APIVersion of the resource
func (r ResourceDaemonSet) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceDaemonSet) GetKind returns the kind of the resource
func (r ResourceDaemonSet) GetKind() string {
	return string(r.Kind)
}

// (r ResourceDaemonSet) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceDaemonSet) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceDaemonSet) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceDaemonSet) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceDaemonSet) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceDaemonSet) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceIngress) GetName returns the name of the resource
func (r ResourceIngress) GetName() string {
	return r.Metadata.Name
}

// (r ResourceIngress) GetUID returns the UID of the resource
func (r ResourceIngress) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceIngress) GetAPIVersion returns the APIVersion of the resource
func (r ResourceIngress) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceIngress) GetKind returns the kind of the resource
func (r ResourceIngress) GetKind() string {
	return string(r.Kind)
}

// (r ResourceIngress) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceIngress) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceIngress) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceIngress) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceIngress) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceIngress) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceIngressClass) GetName returns the name of the resource
func (r ResourceIngressClass) GetName() string {
	return r.Metadata.Name
}

// (r ResourceIngressClass) GetUID returns the UID of the resource
func (r ResourceIngressClass) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceIngressClass) GetAPIVersion returns the APIVersion of the resource
func (r ResourceIngressClass) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceIngressClass) GetKind returns the kind of the resource
func (r ResourceIngressClass) GetKind() string {
	return string(r.Kind)
}

// (r ResourceIngressClass) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceIngressClass) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceIngressClass) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceIngressClass) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceIngressClass) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceIngressClass) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceVulnerabilityReport) GetName returns the name of the resource
func (r ResourceVulnerabilityReport) GetName() string {
	return r.Metadata.Name
}

// (r ResourceVulnerabilityReport) GetUID returns the UID of the resource
func (r ResourceVulnerabilityReport) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceVulnerabilityReport) GetAPIVersion returns the APIVersion of the resource
func (r ResourceVulnerabilityReport) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceVulnerabilityReport) GetKind returns the kind of the resource
func (r ResourceVulnerabilityReport) GetKind() string {
	return string(r.Kind)
}

// (r ResourceVulnerabilityReport) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceVulnerabilityReport) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceVulnerabilityReport) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceVulnerabilityReport) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceVulnerabilityReport) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceVulnerabilityReport) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceExposedSecretReport) GetName returns the name of the resource
func (r ResourceExposedSecretReport) GetName() string {
	return r.Metadata.Name
}

// (r ResourceExposedSecretReport) GetUID returns the UID of the resource
func (r ResourceExposedSecretReport) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceExposedSecretReport) GetAPIVersion returns the APIVersion of the resource
func (r ResourceExposedSecretReport) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceExposedSecretReport) GetKind returns the kind of the resource
func (r ResourceExposedSecretReport) GetKind() string {
	return string(r.Kind)
}

// (r ResourceExposedSecretReport) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceExposedSecretReport) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceExposedSecretReport) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceExposedSecretReport) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceExposedSecretReport) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceExposedSecretReport) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceConfigAuditReport) GetName returns the name of the resource
func (r ResourceConfigAuditReport) GetName() string {
	return r.Metadata.Name
}

// (r ResourceConfigAuditReport) GetUID returns the UID of the resource
func (r ResourceConfigAuditReport) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceConfigAuditReport) GetAPIVersion returns the APIVersion of the resource
func (r ResourceConfigAuditReport) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceConfigAuditReport) GetKind returns the kind of the resource
func (r ResourceConfigAuditReport) GetKind() string {
	return string(r.Kind)
}

// (r ResourceConfigAuditReport) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceConfigAuditReport) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceConfigAuditReport) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceConfigAuditReport) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceConfigAuditReport) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceConfigAuditReport) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceRbacAssessmentReport) GetName returns the name of the resource
func (r ResourceRbacAssessmentReport) GetName() string {
	return r.Metadata.Name
}

// (r ResourceRbacAssessmentReport) GetUID returns the UID of the resource
func (r ResourceRbacAssessmentReport) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceRbacAssessmentReport) GetAPIVersion returns the APIVersion of the resource
func (r ResourceRbacAssessmentReport) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceRbacAssessmentReport) GetKind returns the kind of the resource
func (r ResourceRbacAssessmentReport) GetKind() string {
	return string(r.Kind)
}

// (r ResourceRbacAssessmentReport) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceRbacAssessmentReport) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceRbacAssessmentReport) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceRbacAssessmentReport) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceRbacAssessmentReport) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceRbacAssessmentReport) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceTanzuKubernetesCluster) GetName returns the name of the resource
func (r ResourceTanzuKubernetesCluster) GetName() string {
	return r.Metadata.Name
}

// (r ResourceTanzuKubernetesCluster) GetUID returns the UID of the resource
func (r ResourceTanzuKubernetesCluster) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceTanzuKubernetesCluster) GetAPIVersion returns the APIVersion of the resource
func (r ResourceTanzuKubernetesCluster) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceTanzuKubernetesCluster) GetKind returns the kind of the resource
func (r ResourceTanzuKubernetesCluster) GetKind() string {
	return string(r.Kind)
}

// (r ResourceTanzuKubernetesCluster) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceTanzuKubernetesCluster) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceTanzuKubernetesCluster) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceTanzuKubernetesCluster) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceTanzuKubernetesCluster) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceTanzuKubernetesCluster) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceTanzuKubernetesRelease) GetName returns the name of the resource
func (r ResourceTanzuKubernetesRelease) GetName() string {
	return r.Metadata.Name
}

// (r ResourceTanzuKubernetesRelease) GetUID returns the UID of the resource
func (r ResourceTanzuKubernetesRelease) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceTanzuKubernetesRelease) GetAPIVersion returns the APIVersion of the resource
func (r ResourceTanzuKubernetesRelease) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceTanzuKubernetesRelease) GetKind returns the kind of the resource
func (r ResourceTanzuKubernetesRelease) GetKind() string {
	return string(r.Kind)
}

// (r ResourceTanzuKubernetesRelease) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceTanzuKubernetesRelease) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceTanzuKubernetesRelease) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceTanzuKubernetesRelease) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceTanzuKubernetesRelease) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceTanzuKubernetesRelease) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceVirtualMachineClass) GetName returns the name of the resource
func (r ResourceVirtualMachineClass) GetName() string {
	return r.Metadata.Name
}

// (r ResourceVirtualMachineClass) GetUID returns the UID of the resource
func (r ResourceVirtualMachineClass) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceVirtualMachineClass) GetAPIVersion returns the APIVersion of the resource
func (r ResourceVirtualMachineClass) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceVirtualMachineClass) GetKind returns the kind of the resource
func (r ResourceVirtualMachineClass) GetKind() string {
	return string(r.Kind)
}

// (r ResourceVirtualMachineClass) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceVirtualMachineClass) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceVirtualMachineClass) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceVirtualMachineClass) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceVirtualMachineClass) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceVirtualMachineClass) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceVirtualMachineClassBinding) GetName returns the name of the resource
func (r ResourceVirtualMachineClassBinding) GetName() string {
	return r.Metadata.Name
}

// (r ResourceVirtualMachineClassBinding) GetUID returns the UID of the resource
func (r ResourceVirtualMachineClassBinding) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceVirtualMachineClassBinding) GetAPIVersion returns the APIVersion of the resource
func (r ResourceVirtualMachineClassBinding) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceVirtualMachineClassBinding) GetKind returns the kind of the resource
func (r ResourceVirtualMachineClassBinding) GetKind() string {
	return string(r.Kind)
}

// (r ResourceVirtualMachineClassBinding) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceVirtualMachineClassBinding) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceVirtualMachineClassBinding) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceVirtualMachineClassBinding) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceVirtualMachineClassBinding) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceVirtualMachineClassBinding) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceKubernetesCluster) GetName returns the name of the resource
func (r ResourceKubernetesCluster) GetName() string {
	return r.Metadata.Name
}

// (r ResourceKubernetesCluster) GetUID returns the UID of the resource
func (r ResourceKubernetesCluster) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceKubernetesCluster) GetAPIVersion returns the APIVersion of the resource
func (r ResourceKubernetesCluster) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceKubernetesCluster) GetKind returns the kind of the resource
func (r ResourceKubernetesCluster) GetKind() string {
	return string(r.Kind)
}

// (r ResourceKubernetesCluster) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceKubernetesCluster) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceKubernetesCluster) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceKubernetesCluster) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceKubernetesCluster) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceKubernetesCluster) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceClusterOrder) GetName returns the name of the resource
func (r ResourceClusterOrder) GetName() string {
	return r.Metadata.Name
}

// (r ResourceClusterOrder) GetUID returns the UID of the resource
func (r ResourceClusterOrder) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceClusterOrder) GetAPIVersion returns the APIVersion of the resource
func (r ResourceClusterOrder) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceClusterOrder) GetKind returns the kind of the resource
func (r ResourceClusterOrder) GetKind() string {
	return string(r.Kind)
}

// (r ResourceClusterOrder) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceClusterOrder) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceClusterOrder) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceClusterOrder) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceClusterOrder) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceClusterOrder) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceProject) GetName returns the name of the resource
func (r ResourceProject) GetName() string {
	return r.Metadata.Name
}

// (r ResourceProject) GetUID returns the UID of the resource
func (r ResourceProject) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceProject) GetAPIVersion returns the APIVersion of the resource
func (r ResourceProject) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceProject) GetKind returns the kind of the resource
func (r ResourceProject) GetKind() string {
	return string(r.Kind)
}

// (r ResourceProject) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceProject) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceProject) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceProject) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceProject) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceProject) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceConfiguration) GetName returns the name of the resource
func (r ResourceConfiguration) GetName() string {
	return r.Metadata.Name
}

// (r ResourceConfiguration) GetUID returns the UID of the resource
func (r ResourceConfiguration) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceConfiguration) GetAPIVersion returns the APIVersion of the resource
func (r ResourceConfiguration) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceConfiguration) GetKind returns the kind of the resource
func (r ResourceConfiguration) GetKind() string {
	return string(r.Kind)
}

// (r ResourceConfiguration) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceConfiguration) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceConfiguration) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceConfiguration) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceConfiguration) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceConfiguration) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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

// (r ResourceClusterComplianceReport) GetName returns the name of the resource
func (r ResourceClusterComplianceReport) GetName() string {
	return r.Metadata.Name
}

// (r ResourceClusterComplianceReport) GetUID returns the UID of the resource
func (r ResourceClusterComplianceReport) GetUID() string {
	return string(r.Metadata.UID)
}

// (r ResourceClusterComplianceReport) GetAPIVersion returns the APIVersion of the resource
func (r ResourceClusterComplianceReport) GetAPIVersion() string {
	return string(r.APIVersion)
}

// (r ResourceClusterComplianceReport) GetKind returns the kind of the resource
func (r ResourceClusterComplianceReport) GetKind() string {
	return string(r.Kind)
}

// (r ResourceClusterComplianceReport) GetMetadata returns the metav1.ObjectMeta of the resource
func (r ResourceClusterComplianceReport) GetMetadata() metav1.ObjectMeta {
	return r.Metadata
}

// (r ResourceClusterComplianceReport) GetRorMeta returns the ResourceRorMeta of the resource
func (r ResourceClusterComplianceReport) GetRorMeta() ResourceRorMeta {
	return r.RorMeta
}

// (r *ResourceClusterComplianceReport) SetRorMeta sets the ResourceRorMeta of the resource
func (r *ResourceClusterComplianceReport) SetRorMeta(rormeta ResourceRorMeta) error {
	r.RorMeta = rormeta
	r.RorMeta.Hash = r.GetRorHash()
	return nil
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
