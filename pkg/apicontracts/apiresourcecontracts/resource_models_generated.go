// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package apiresourcecontracts

// Resourcetypes allowed in the generic resource models.// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type Resourcetypes interface {
	ResourceNamespace | ResourceNode | ResourcePersistentVolumeClaim | ResourceDeployment | ResourceStorageClass | ResourcePolicyReport | ResourceApplication | ResourceAppProject | ResourceCertificate | ResourceService | ResourcePod | ResourceReplicaSet | ResourceStatefulSet | ResourceDaemonSet | ResourceIngress | ResourceIngressClass | ResourceVulnerabilityReport | ResourceExposedSecretReport | ResourceConfigAuditReport | ResourceRbacAssessmentReport | ResourceTanzuKubernetesCluster | ResourceTanzuKubernetesRelease | ResourceVirtualMachineClass | ResourceKubernetesCluster | ResourceClusterOrder | ResourceProject | ResourceConfiguration | ResourceClusterComplianceReport | ResourceClusterVulnerabilityReport | ResourceRoute | ResourceSlackMessage | ResourceVulnerabilityEvent | ResourceVirtualMachine | ResourceEndpoints | ResourceNetworkPolicy | ResourceBackupJob
}

// type for returning Namespace resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListNamespaces struct {
	Owner      ResourceOwnerReference `json:"owner"`
	Namespaces []ResourceNamespace    `json:"namespaces"`
}

// type for returning Node resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListNodes struct {
	Owner ResourceOwnerReference `json:"owner"`
	Nodes []ResourceNode         `json:"nodes"`
}

// type for returning PersistentVolumeClaim resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListPersistentvolumeclaims struct {
	Owner                  ResourceOwnerReference          `json:"owner"`
	Persistentvolumeclaims []ResourcePersistentVolumeClaim `json:"persistentvolumeclaims"`
}

// type for returning Deployment resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListDeployments struct {
	Owner       ResourceOwnerReference `json:"owner"`
	Deployments []ResourceDeployment   `json:"deployments"`
}

// type for returning StorageClass resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListStorageclasses struct {
	Owner          ResourceOwnerReference `json:"owner"`
	Storageclasses []ResourceStorageClass `json:"storageclasses"`
}

// type for returning PolicyReport resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListPolicyreports struct {
	Owner         ResourceOwnerReference `json:"owner"`
	Policyreports []ResourcePolicyReport `json:"policyreports"`
}

// type for returning Application resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListApplications struct {
	Owner        ResourceOwnerReference `json:"owner"`
	Applications []ResourceApplication  `json:"applications"`
}

// type for returning AppProject resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListAppprojects struct {
	Owner       ResourceOwnerReference `json:"owner"`
	Appprojects []ResourceAppProject   `json:"appprojects"`
}

// type for returning Certificate resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListCertificates struct {
	Owner        ResourceOwnerReference `json:"owner"`
	Certificates []ResourceCertificate  `json:"certificates"`
}

// type for returning Service resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListServices struct {
	Owner    ResourceOwnerReference `json:"owner"`
	Services []ResourceService      `json:"services"`
}

// type for returning Pod resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListPods struct {
	Owner ResourceOwnerReference `json:"owner"`
	Pods  []ResourcePod          `json:"pods"`
}

// type for returning ReplicaSet resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListReplicasets struct {
	Owner       ResourceOwnerReference `json:"owner"`
	Replicasets []ResourceReplicaSet   `json:"replicasets"`
}

// type for returning StatefulSet resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListStatefulsets struct {
	Owner        ResourceOwnerReference `json:"owner"`
	Statefulsets []ResourceStatefulSet  `json:"statefulsets"`
}

// type for returning DaemonSet resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListDaemonsets struct {
	Owner      ResourceOwnerReference `json:"owner"`
	Daemonsets []ResourceDaemonSet    `json:"daemonsets"`
}

// type for returning Ingress resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListIngresses struct {
	Owner     ResourceOwnerReference `json:"owner"`
	Ingresses []ResourceIngress      `json:"ingresses"`
}

// type for returning IngressClass resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListIngressclasses struct {
	Owner          ResourceOwnerReference `json:"owner"`
	Ingressclasses []ResourceIngressClass `json:"ingressclasses"`
}

// type for returning VulnerabilityReport resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListVulnerabilityreports struct {
	Owner                ResourceOwnerReference        `json:"owner"`
	Vulnerabilityreports []ResourceVulnerabilityReport `json:"vulnerabilityreports"`
}

// type for returning ExposedSecretReport resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListExposedsecretreports struct {
	Owner                ResourceOwnerReference        `json:"owner"`
	Exposedsecretreports []ResourceExposedSecretReport `json:"exposedsecretreports"`
}

// type for returning ConfigAuditReport resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListConfigauditreports struct {
	Owner              ResourceOwnerReference      `json:"owner"`
	Configauditreports []ResourceConfigAuditReport `json:"configauditreports"`
}

// type for returning RbacAssessmentReport resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListRbacassessmentreports struct {
	Owner                 ResourceOwnerReference         `json:"owner"`
	Rbacassessmentreports []ResourceRbacAssessmentReport `json:"rbacassessmentreports"`
}

// type for returning TanzuKubernetesCluster resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListTanzukubernetesclusters struct {
	Owner                   ResourceOwnerReference           `json:"owner"`
	Tanzukubernetesclusters []ResourceTanzuKubernetesCluster `json:"tanzukubernetesclusters"`
}

// type for returning TanzuKubernetesRelease resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListTanzukubernetesreleases struct {
	Owner                   ResourceOwnerReference           `json:"owner"`
	Tanzukubernetesreleases []ResourceTanzuKubernetesRelease `json:"tanzukubernetesreleases"`
}

// type for returning VirtualMachineClass resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListVirtualmachineclasses struct {
	Owner                 ResourceOwnerReference        `json:"owner"`
	Virtualmachineclasses []ResourceVirtualMachineClass `json:"virtualmachineclasses"`
}

// type for returning KubernetesCluster resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListKubernetesclusters struct {
	Owner              ResourceOwnerReference      `json:"owner"`
	Kubernetesclusters []ResourceKubernetesCluster `json:"kubernetesclusters"`
}

// type for returning ClusterOrder resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListClusterorders struct {
	Owner         ResourceOwnerReference `json:"owner"`
	Clusterorders []ResourceClusterOrder `json:"clusterorders"`
}

// type for returning Project resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListProjects struct {
	Owner    ResourceOwnerReference `json:"owner"`
	Projects []ResourceProject      `json:"projects"`
}

// type for returning Configuration resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListConfigurations struct {
	Owner          ResourceOwnerReference  `json:"owner"`
	Configurations []ResourceConfiguration `json:"configurations"`
}

// type for returning ClusterComplianceReport resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListClustercompliancereports struct {
	Owner                    ResourceOwnerReference            `json:"owner"`
	Clustercompliancereports []ResourceClusterComplianceReport `json:"clustercompliancereports"`
}

// type for returning ClusterVulnerabilityReport resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListClustervulnerabilityreports struct {
	Owner                       ResourceOwnerReference               `json:"owner"`
	Clustervulnerabilityreports []ResourceClusterVulnerabilityReport `json:"clustervulnerabilityreports"`
}

// type for returning Route resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListRoutes struct {
	Owner  ResourceOwnerReference `json:"owner"`
	Routes []ResourceRoute        `json:"routes"`
}

// type for returning SlackMessage resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListSlackmessages struct {
	Owner         ResourceOwnerReference `json:"owner"`
	Slackmessages []ResourceSlackMessage `json:"slackmessages"`
}

// type for returning VulnerabilityEvent resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListVulnerabilityevents struct {
	Owner               ResourceOwnerReference       `json:"owner"`
	Vulnerabilityevents []ResourceVulnerabilityEvent `json:"vulnerabilityevents"`
}

// type for returning VirtualMachine resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListVirtualmachines struct {
	Owner           ResourceOwnerReference   `json:"owner"`
	Virtualmachines []ResourceVirtualMachine `json:"VirtualMachines"`
}

// type for returning Endpoints resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListEndpoints struct {
	Owner     ResourceOwnerReference `json:"owner"`
	Endpoints []ResourceEndpoints    `json:"endpoints"`
}

// type for returning NetworkPolicy resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListNetworkpolicies struct {
	Owner           ResourceOwnerReference  `json:"owner"`
	Networkpolicies []ResourceNetworkPolicy `json:"networkpolicies"`
}

// type for returning BackupJob resources to internal functions// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceListBackupjobs struct {
	Owner      ResourceOwnerReference `json:"owner"`
	Backupjobs []ResourceBackupJob    `json:"backupjobs"`
}
