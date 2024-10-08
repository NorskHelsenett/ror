// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package apiresourcecontracts

import "github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"

// Resourcetypes allowed in the generic resource models.
type Resourcetypes interface {
	ResourceNamespace | ResourceNode | ResourcePersistentVolumeClaim | ResourceDeployment | ResourceStorageClass | ResourcePolicyReport | ResourceApplication | ResourceAppProject | ResourceCertificate | ResourceService | ResourcePod | ResourceReplicaSet | ResourceStatefulSet | ResourceDaemonSet | ResourceIngress | ResourceIngressClass | ResourceVulnerabilityReport | ResourceExposedSecretReport | ResourceConfigAuditReport | ResourceRbacAssessmentReport | ResourceTanzuKubernetesCluster | ResourceTanzuKubernetesRelease | ResourceVirtualMachineClass | ResourceVirtualMachineClassBinding | ResourceKubernetesCluster | ResourceClusterOrder | ResourceProject | ResourceConfiguration | ResourceClusterComplianceReport | ResourceClusterVulnerabilityReport | ResourceRoute | ResourceSlackMessage | ResourceNotification
}

// type for returning Namespace resources to internal functions
type ResourceNamespaces struct {
	Owner      rortypes.RorResourceOwnerReference `json:"owner"`
	Namespaces []ResourceNamespace                `json:"namespaces"`
}

// type for returning Node resources to internal functions
type ResourceNodes struct {
	Owner rortypes.RorResourceOwnerReference `json:"owner"`
	Nodes []ResourceNode                     `json:"nodes"`
}

// type for returning PersistentVolumeClaim resources to internal functions
type ResourcePersistentvolumeclaims struct {
	Owner                  rortypes.RorResourceOwnerReference `json:"owner"`
	Persistentvolumeclaims []ResourcePersistentVolumeClaim    `json:"persistentvolumeclaims"`
}

// type for returning Deployment resources to internal functions
type ResourceDeployments struct {
	Owner       rortypes.RorResourceOwnerReference `json:"owner"`
	Deployments []ResourceDeployment               `json:"deployments"`
}

// type for returning StorageClass resources to internal functions
type ResourceStorageclasses struct {
	Owner          rortypes.RorResourceOwnerReference `json:"owner"`
	Storageclasses []ResourceStorageClass             `json:"storageclasses"`
}

// type for returning PolicyReport resources to internal functions
type ResourcePolicyreports struct {
	Owner         rortypes.RorResourceOwnerReference `json:"owner"`
	Policyreports []ResourcePolicyReport             `json:"policyreports"`
}

// type for returning Application resources to internal functions
type ResourceApplications struct {
	Owner        rortypes.RorResourceOwnerReference `json:"owner"`
	Applications []ResourceApplication              `json:"applications"`
}

// type for returning AppProject resources to internal functions
type ResourceAppprojects struct {
	Owner       rortypes.RorResourceOwnerReference `json:"owner"`
	Appprojects []ResourceAppProject               `json:"appprojects"`
}

// type for returning Certificate resources to internal functions
type ResourceCertificates struct {
	Owner        rortypes.RorResourceOwnerReference `json:"owner"`
	Certificates []ResourceCertificate              `json:"certificates"`
}

// type for returning Service resources to internal functions
type ResourceServices struct {
	Owner    rortypes.RorResourceOwnerReference `json:"owner"`
	Services []ResourceService                  `json:"services"`
}

// type for returning Pod resources to internal functions
type ResourcePods struct {
	Owner rortypes.RorResourceOwnerReference `json:"owner"`
	Pods  []ResourcePod                      `json:"pods"`
}

// type for returning ReplicaSet resources to internal functions
type ResourceReplicasets struct {
	Owner       rortypes.RorResourceOwnerReference `json:"owner"`
	Replicasets []ResourceReplicaSet               `json:"replicasets"`
}

// type for returning StatefulSet resources to internal functions
type ResourceStatefulsets struct {
	Owner        rortypes.RorResourceOwnerReference `json:"owner"`
	Statefulsets []ResourceStatefulSet              `json:"statefulsets"`
}

// type for returning DaemonSet resources to internal functions
type ResourceDaemonsets struct {
	Owner      rortypes.RorResourceOwnerReference `json:"owner"`
	Daemonsets []ResourceDaemonSet                `json:"daemonsets"`
}

// type for returning Ingress resources to internal functions
type ResourceIngresses struct {
	Owner     rortypes.RorResourceOwnerReference `json:"owner"`
	Ingresses []ResourceIngress                  `json:"ingresses"`
}

// type for returning IngressClass resources to internal functions
type ResourceIngressclasses struct {
	Owner          rortypes.RorResourceOwnerReference `json:"owner"`
	Ingressclasses []ResourceIngressClass             `json:"ingressclasses"`
}

// type for returning VulnerabilityReport resources to internal functions
type ResourceVulnerabilityreports struct {
	Owner                rortypes.RorResourceOwnerReference `json:"owner"`
	Vulnerabilityreports []ResourceVulnerabilityReport      `json:"vulnerabilityreports"`
}

// type for returning ExposedSecretReport resources to internal functions
type ResourceExposedsecretreports struct {
	Owner                rortypes.RorResourceOwnerReference `json:"owner"`
	Exposedsecretreports []ResourceExposedSecretReport      `json:"exposedsecretreports"`
}

// type for returning ConfigAuditReport resources to internal functions
type ResourceConfigauditreports struct {
	Owner              rortypes.RorResourceOwnerReference `json:"owner"`
	Configauditreports []ResourceConfigAuditReport        `json:"configauditreports"`
}

// type for returning RbacAssessmentReport resources to internal functions
type ResourceRbacassessmentreports struct {
	Owner                 rortypes.RorResourceOwnerReference `json:"owner"`
	Rbacassessmentreports []ResourceRbacAssessmentReport     `json:"rbacassessmentreports"`
}

// type for returning TanzuKubernetesCluster resources to internal functions
type ResourceTanzukubernetesclusters struct {
	Owner                   rortypes.RorResourceOwnerReference `json:"owner"`
	Tanzukubernetesclusters []ResourceTanzuKubernetesCluster   `json:"tanzukubernetesclusters"`
}

// type for returning TanzuKubernetesRelease resources to internal functions
type ResourceTanzukubernetesreleases struct {
	Owner                   rortypes.RorResourceOwnerReference `json:"owner"`
	Tanzukubernetesreleases []ResourceTanzuKubernetesRelease   `json:"tanzukubernetesreleases"`
}

// type for returning VirtualMachineClass resources to internal functions
type ResourceVirtualmachineclasses struct {
	Owner                 rortypes.RorResourceOwnerReference `json:"owner"`
	Virtualmachineclasses []ResourceVirtualMachineClass      `json:"virtualmachineclasses"`
}

// type for returning VirtualMachineClassBinding resources to internal functions
type ResourceVirtualmachineclassbindings struct {
	Owner                       rortypes.RorResourceOwnerReference   `json:"owner"`
	Virtualmachineclassbindings []ResourceVirtualMachineClassBinding `json:"virtualmachineclassbindings"`
}

// type for returning KubernetesCluster resources to internal functions
type ResourceKubernetesclusters struct {
	Owner              rortypes.RorResourceOwnerReference `json:"owner"`
	Kubernetesclusters []ResourceKubernetesCluster        `json:"kubernetesclusters"`
}

// type for returning ClusterOrder resources to internal functions
type ResourceClusterorders struct {
	Owner         rortypes.RorResourceOwnerReference `json:"owner"`
	Clusterorders []ResourceClusterOrder             `json:"clusterorders"`
}

// type for returning Project resources to internal functions
type ResourceProjects struct {
	Owner    rortypes.RorResourceOwnerReference `json:"owner"`
	Projects []ResourceProject                  `json:"projects"`
}

// type for returning Configuration resources to internal functions
type ResourceConfigurations struct {
	Owner          rortypes.RorResourceOwnerReference `json:"owner"`
	Configurations []ResourceConfiguration            `json:"configurations"`
}

// type for returning ClusterComplianceReport resources to internal functions
type ResourceClustercompliancereports struct {
	Owner                    rortypes.RorResourceOwnerReference `json:"owner"`
	Clustercompliancereports []ResourceClusterComplianceReport  `json:"clustercompliancereports"`
}

// type for returning ClusterVulnerabilityReport resources to internal functions
type ResourceClustervulnerabilityreports struct {
	Owner                       rortypes.RorResourceOwnerReference   `json:"owner"`
	Clustervulnerabilityreports []ResourceClusterVulnerabilityReport `json:"clustervulnerabilityreports"`
}

// type for returning Route resources to internal functions
type ResourceRoutes struct {
	Owner  rortypes.RorResourceOwnerReference `json:"owner"`
	Routes []ResourceRoute                    `json:"routes"`
}

// type for returning SlackMessage resources to internal functions
type ResourceSlackmessages struct {
	Owner         rortypes.RorResourceOwnerReference `json:"owner"`
	Slackmessages []ResourceSlackMessage             `json:"slackmessages"`
}

// type for returning Notification resources to internal functions
type ResourceNotifications struct {
	Owner         rortypes.RorResourceOwnerReference `json:"owner"`
	Notifications []ResourceNotification             `json:"notifications"`
}
