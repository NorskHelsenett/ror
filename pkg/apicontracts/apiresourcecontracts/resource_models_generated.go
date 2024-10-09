// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package apiresourcecontracts

// Resourcetypes allowed in the generic resource models.
type Resourcetypes interface {
	ResourceNamespace | ResourceNode | ResourcePersistentVolumeClaim | ResourceDeployment | ResourceStorageClass | ResourcePolicyReport | ResourceApplication | ResourceAppProject | ResourceCertificate | ResourceService | ResourcePod | ResourceReplicaSet | ResourceStatefulSet | ResourceDaemonSet | ResourceIngress | ResourceIngressClass | ResourceVulnerabilityReport | ResourceExposedSecretReport | ResourceConfigAuditReport | ResourceRbacAssessmentReport | ResourceTanzuKubernetesCluster | ResourceTanzuKubernetesRelease | ResourceVirtualMachineClass | ResourceVirtualMachineClassBinding | ResourceKubernetesCluster | ResourceClusterOrder | ResourceProject | ResourceConfiguration | ResourceClusterComplianceReport | ResourceClusterVulnerabilityReport | ResourceRoute | ResourceSlackMessage | ResourceVulnerabilityEvent | ResourceVulnerabilityWhitelist
}

// type for returning Namespace resources to internal functions
type ResourceNamespaces struct {
	Owner      ResourceOwnerReference `json:"owner"`
	Namespaces []ResourceNamespace    `json:"namespaces"`
}

// type for returning Node resources to internal functions
type ResourceNodes struct {
	Owner ResourceOwnerReference `json:"owner"`
	Nodes []ResourceNode         `json:"nodes"`
}

// type for returning PersistentVolumeClaim resources to internal functions
type ResourcePersistentvolumeclaims struct {
	Owner                  ResourceOwnerReference          `json:"owner"`
	Persistentvolumeclaims []ResourcePersistentVolumeClaim `json:"persistentvolumeclaims"`
}

// type for returning Deployment resources to internal functions
type ResourceDeployments struct {
	Owner       ResourceOwnerReference `json:"owner"`
	Deployments []ResourceDeployment   `json:"deployments"`
}

// type for returning StorageClass resources to internal functions
type ResourceStorageclasses struct {
	Owner          ResourceOwnerReference `json:"owner"`
	Storageclasses []ResourceStorageClass `json:"storageclasses"`
}

// type for returning PolicyReport resources to internal functions
type ResourcePolicyreports struct {
	Owner         ResourceOwnerReference `json:"owner"`
	Policyreports []ResourcePolicyReport `json:"policyreports"`
}

// type for returning Application resources to internal functions
type ResourceApplications struct {
	Owner        ResourceOwnerReference `json:"owner"`
	Applications []ResourceApplication  `json:"applications"`
}

// type for returning AppProject resources to internal functions
type ResourceAppprojects struct {
	Owner       ResourceOwnerReference `json:"owner"`
	Appprojects []ResourceAppProject   `json:"appprojects"`
}

// type for returning Certificate resources to internal functions
type ResourceCertificates struct {
	Owner        ResourceOwnerReference `json:"owner"`
	Certificates []ResourceCertificate  `json:"certificates"`
}

// type for returning Service resources to internal functions
type ResourceServices struct {
	Owner    ResourceOwnerReference `json:"owner"`
	Services []ResourceService      `json:"services"`
}

// type for returning Pod resources to internal functions
type ResourcePods struct {
	Owner ResourceOwnerReference `json:"owner"`
	Pods  []ResourcePod          `json:"pods"`
}

// type for returning ReplicaSet resources to internal functions
type ResourceReplicasets struct {
	Owner       ResourceOwnerReference `json:"owner"`
	Replicasets []ResourceReplicaSet   `json:"replicasets"`
}

// type for returning StatefulSet resources to internal functions
type ResourceStatefulsets struct {
	Owner        ResourceOwnerReference `json:"owner"`
	Statefulsets []ResourceStatefulSet  `json:"statefulsets"`
}

// type for returning DaemonSet resources to internal functions
type ResourceDaemonsets struct {
	Owner      ResourceOwnerReference `json:"owner"`
	Daemonsets []ResourceDaemonSet    `json:"daemonsets"`
}

// type for returning Ingress resources to internal functions
type ResourceIngresses struct {
	Owner     ResourceOwnerReference `json:"owner"`
	Ingresses []ResourceIngress      `json:"ingresses"`
}

// type for returning IngressClass resources to internal functions
type ResourceIngressclasses struct {
	Owner          ResourceOwnerReference `json:"owner"`
	Ingressclasses []ResourceIngressClass `json:"ingressclasses"`
}

// type for returning VulnerabilityReport resources to internal functions
type ResourceVulnerabilityreports struct {
	Owner                ResourceOwnerReference        `json:"owner"`
	Vulnerabilityreports []ResourceVulnerabilityReport `json:"vulnerabilityreports"`
}

// type for returning ExposedSecretReport resources to internal functions
type ResourceExposedsecretreports struct {
	Owner                ResourceOwnerReference        `json:"owner"`
	Exposedsecretreports []ResourceExposedSecretReport `json:"exposedsecretreports"`
}

// type for returning ConfigAuditReport resources to internal functions
type ResourceConfigauditreports struct {
	Owner              ResourceOwnerReference      `json:"owner"`
	Configauditreports []ResourceConfigAuditReport `json:"configauditreports"`
}

// type for returning RbacAssessmentReport resources to internal functions
type ResourceRbacassessmentreports struct {
	Owner                 ResourceOwnerReference         `json:"owner"`
	Rbacassessmentreports []ResourceRbacAssessmentReport `json:"rbacassessmentreports"`
}

// type for returning TanzuKubernetesCluster resources to internal functions
type ResourceTanzukubernetesclusters struct {
	Owner                   ResourceOwnerReference           `json:"owner"`
	Tanzukubernetesclusters []ResourceTanzuKubernetesCluster `json:"tanzukubernetesclusters"`
}

// type for returning TanzuKubernetesRelease resources to internal functions
type ResourceTanzukubernetesreleases struct {
	Owner                   ResourceOwnerReference           `json:"owner"`
	Tanzukubernetesreleases []ResourceTanzuKubernetesRelease `json:"tanzukubernetesreleases"`
}

// type for returning VirtualMachineClass resources to internal functions
type ResourceVirtualmachineclasses struct {
	Owner                 ResourceOwnerReference        `json:"owner"`
	Virtualmachineclasses []ResourceVirtualMachineClass `json:"virtualmachineclasses"`
}

// type for returning VirtualMachineClassBinding resources to internal functions
type ResourceVirtualmachineclassbindings struct {
	Owner                       ResourceOwnerReference               `json:"owner"`
	Virtualmachineclassbindings []ResourceVirtualMachineClassBinding `json:"virtualmachineclassbindings"`
}

// type for returning KubernetesCluster resources to internal functions
type ResourceKubernetesclusters struct {
	Owner              ResourceOwnerReference      `json:"owner"`
	Kubernetesclusters []ResourceKubernetesCluster `json:"kubernetesclusters"`
}

// type for returning ClusterOrder resources to internal functions
type ResourceClusterorders struct {
	Owner         ResourceOwnerReference `json:"owner"`
	Clusterorders []ResourceClusterOrder `json:"clusterorders"`
}

// type for returning Project resources to internal functions
type ResourceProjects struct {
	Owner    ResourceOwnerReference `json:"owner"`
	Projects []ResourceProject      `json:"projects"`
}

// type for returning Configuration resources to internal functions
type ResourceConfigurations struct {
	Owner          ResourceOwnerReference  `json:"owner"`
	Configurations []ResourceConfiguration `json:"configurations"`
}

// type for returning ClusterComplianceReport resources to internal functions
type ResourceClustercompliancereports struct {
	Owner                    ResourceOwnerReference            `json:"owner"`
	Clustercompliancereports []ResourceClusterComplianceReport `json:"clustercompliancereports"`
}

// type for returning ClusterVulnerabilityReport resources to internal functions
type ResourceClustervulnerabilityreports struct {
	Owner                       ResourceOwnerReference               `json:"owner"`
	Clustervulnerabilityreports []ResourceClusterVulnerabilityReport `json:"clustervulnerabilityreports"`
}

// type for returning Route resources to internal functions
type ResourceRoutes struct {
	Owner  ResourceOwnerReference `json:"owner"`
	Routes []ResourceRoute        `json:"routes"`
}

// type for returning SlackMessage resources to internal functions
type ResourceSlackmessages struct {
	Owner         ResourceOwnerReference `json:"owner"`
	Slackmessages []ResourceSlackMessage `json:"slackmessages"`
}

// type for returning VulnerabilityEvent resources to internal functions
type ResourceVulnerabilityevents struct {
	Owner               ResourceOwnerReference       `json:"owner"`
	Vulnerabilityevents []ResourceVulnerabilityEvent `json:"vulnerabilityevents"`
}

// type for returning VulnerabilityWhitelist resources to internal functions
type ResourceVulnerabilitywhitelists struct {
	Owner                   ResourceOwnerReference           `json:"owner"`
	Vulnerabilitywhitelists []ResourceVulnerabilityWhitelist `json:"vulnerabilitywhitelists"`
}
