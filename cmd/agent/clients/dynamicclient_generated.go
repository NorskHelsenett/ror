// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package clients

import "k8s.io/apimachinery/pkg/runtime/schema"

func InitSchema() []schema.GroupVersionResource {
	var schemas []schema.GroupVersionResource
	schemas = append(schemas, schema.GroupVersionResource{Group: "", Version: "v1", Resource: "namespaces"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "", Version: "v1", Resource: "nodes"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "", Version: "v1", Resource: "persistentvolumeclaims"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "deployments"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "storageclasses"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "wgpolicyk8s.io", Version: "v1alpha2", Resource: "policyreports"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "applications"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "argoproj.io", Version: "v1alpha1", Resource: "appprojects"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "cert-manager.io", Version: "v1", Resource: "certificates"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "", Version: "v1", Resource: "services"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "replicasets"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "statefulsets"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "apps", Version: "v1", Resource: "daemonsets"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "ingresses"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "networking.k8s.io", Version: "v1", Resource: "ingressclasses"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "aquasecurity.github.io", Version: "v1alpha1", Resource: "vulnerabilityreports"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "aquasecurity.github.io", Version: "v1alpha1", Resource: "exposedsecretreports"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "aquasecurity.github.io", Version: "v1alpha1", Resource: "configauditreports"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "aquasecurity.github.io", Version: "v1alpha1", Resource: "rbacassessmentreports"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "aquasecurity.github.io", Version: "v1alpha1", Resource: "clustercompliancereports"})

	return schemas
}
