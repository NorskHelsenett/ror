package rortypes

// ResourceTanzuKubernetesCluster
// K8s node struct
// Tanzu kubernetes release struct
type ResourceTanzuKubernetesRelease struct {
	CommonResource `json:",inline"`
	Spec           ResourceTanzuKubernetesReleaseSpec   `json:"spec"`
	Status         ResourceTanzuKubernetesReleaseStatus `json:"status"`
}

type ResourceTanzuKubernetesReleaseMetadata struct {
	Annotations                map[string]string                                      `json:"annotations"`
	ClusterName                string                                                 `json:"clusterName"`
	CreationTimestamp          string                                                 `json:"creationTimestamp"`
	DeletionGracePeriodSeconds int                                                    `json:"deletionGracePeriodSeconds"`
	DeletionTimestamp          string                                                 `json:"deletionTimestamp"`
	Finalizers                 []string                                               `json:"finalizers"`
	GenerateName               string                                                 `json:"generateName"`
	Generation                 int                                                    `json:"generation"`
	Labels                     map[string]string                                      `json:"labels"`
	ManagedFields              []ResourceTanzuKubernetesReleaseMetadataManagedField   `json:"managedFields"`
	Name                       string                                                 `json:"name"`
	Namespace                  string                                                 `json:"namespace"`
	OwnerReferences            []ResourceTanzuKubernetesReleaseMetadataOwnerReference `json:"ownerReferences"`
	//ResourceVersion            string                                                  `json:"resourceVersion"`
	SelfLink string `json:"selfLink"`
	Uid      string `json:"uid"`
}

type ResourceTanzuKubernetesReleaseMetadataManagedField struct {
	ApiVersion string `json:"apiVersion"`
	FieldsType string `json:"fieldsType"`
	//FieldsV1    map[string]string `json:"fieldsV1"`
	Manager     string `json:"manager"`
	Operation   string `json:"operation"`
	Subresource string `json:"subresource"`
	Time        string `json:"time"`
}

type ResourceTanzuKubernetesReleaseMetadataOwnerReference struct {
	ApiVersion         string `json:"apiVersion"`
	BlockOwnerDeletion bool   `json:"blockOwnerDeletion"`
	Controller         bool   `json:"controller"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Uid                string `json:"uid"`
}

type ResourceTanzuKubernetesReleaseSpec struct {
	Images            []ResourceTanzuKubernetesReleaseSpecImage      `json:"images"`
	KubernetesVersion string                                         `json:"kubernetesVersion"`
	NodeImageRef      ResourceTanzuKubernetesReleaseSpecNodeImageRef `json:"nodeImageRef"`
	Repository        string                                         `json:"repository"`
	Version           string                                         `json:"version"`
}

type ResourceTanzuKubernetesReleaseSpecImage struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type ResourceTanzuKubernetesReleaseSpecNodeImageRef struct {
	ApiVersion string `json:"apiVersion"`
	FieldPath  string `json:"fieldPath"`
	Kind       string `json:"kind"`
	Namespace  string `json:"namespace"`
	//ResourceVersion string `json:"resourceVersion"`
	Uid string `json:"uid"`
}

type ResourceTanzuKubernetesReleaseStatus struct {
	Conditions []ResourceTanzuKubernetesReleaseStatusCondition `json:"conditions"`
}

type ResourceTanzuKubernetesReleaseStatusCondition struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Severity           string `json:"severity"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}
