// THIS FILE IS GENERATED, DO NOT EDIT
// ref: build/generator/main.go

package schemas

import "k8s.io/apimachinery/pkg/runtime/schema"

func InitSchema() []schema.GroupVersionResource {
	var schemas []schema.GroupVersionResource
	schemas = append(schemas, schema.GroupVersionResource{Group: "run.tanzu.vmware.com", Version: "v1alpha2", Resource: "tanzukubernetesclusters"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "run.tanzu.vmware.com", Version: "v1alpha2", Resource: "tanzukubernetesreleases"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "vmoperator.vmware.com", Version: "v1alpha1", Resource: "virtualmachineclasses"})
	schemas = append(schemas, schema.GroupVersionResource{Group: "vmoperator.vmware.com", Version: "v1alpha1", Resource: "virtualmachineclassbindings"})

	return schemas
}
