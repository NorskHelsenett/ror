package rordefs

import (
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type ApiVersions string

const (
	ApiVersionV1 = "v1"
	ApiVersionV2 = "v2"
)

type ApiResourceType string

const (
	ApiResourceTypeUnknown    ApiResourceType = ""
	ApiResourceTypeAgent      ApiResourceType = "Agent"
	ApiResourceTypeVmAgent    ApiResourceType = "VmAgent"
	ApiResourceTypeTanzuAgent ApiResourceType = "TanzuAgent"
	ApiResourceTypeInternal   ApiResourceType = "Internal"
)

// ApiResources
// The type describing a list of resources implemented in ror
type ApiResources []ApiResource

// ApiResource
// The type describing a resource implemented in ror
type ApiResource struct {
	metav1.TypeMeta `json:",inline"`
	Plural          string
	Namespaced      bool
	Types           []ApiResourceType
	Versions        []ApiVersions
}

// GetApiVersion
// Generates the apiVersion from the resource object to match with kubernetes api resources
func (m ApiResource) GetApiVersion() string {
	return m.APIVersion
}

func (m ApiResource) GetVersion() string {
	return m.GroupVersionKind().Version
}

func (m ApiResource) GetGroup() string {
	return m.GroupVersionKind().Group
}

func (m ApiResource) GetKind() string {
	return m.Kind
}

func (m ApiResource) GetResource() string {
	return m.Plural
}

func (m ApiResource) GetGroupVersionKind() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   m.GroupVersionKind().Group,
		Version: m.GroupVersionKind().Version,
		Kind:    m.GetKind(),
	}
}

func (m ApiResource) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    m.GroupVersionKind().Group,
		Version:  m.GroupVersionKind().Version,
		Resource: m.GetResource(),
	}
}

func (m ApiResource) PluralCapitalized() string {
	caser := cases.Title(language.Und)

	return caser.String(m.Plural)
}

func (m ApiResources) GetSchemasByType(resourceType ApiResourceType) []schema.GroupVersionResource {
	var resources []schema.GroupVersionResource
	for _, resource := range m.GetResourcesByType(resourceType) {
		resources = append(resources, resource.GetGroupVersionResource())
	}

	return resources
}

func (m ApiResources) GetResourcesByType(resourceType ApiResourceType) ApiResources {
	var resources ApiResources
	for _, resource := range m {
		if slices.Contains(resource.Types, resourceType) {
			resources = append(resources, resource)
		}
	}
	return resources
}

func (m ApiResources) GetResourcesByVersion(version ApiVersions) ApiResources {
	var resources ApiResources
	for _, resource := range m {
		if slices.Contains(resource.Versions, version) {
			resources = append(resources, resource)
		}
	}
	return resources
}

// Deprecated: migrate to the ApiResources.GetSchemasByType
func GetSchemasByType(resourceType ApiResourceType) []schema.GroupVersionResource {
	var resources []schema.GroupVersionResource
	for _, resource := range GetResourcesByType(resourceType) {
		resources = append(resources, resource.GetGroupVersionResource())
	}

	return resources
}

// Deprecated: migrate to the ApiResources.GetResourcesByType
func GetResourcesByType(resourceType ApiResourceType) []ApiResource {
	var resources []ApiResource
	for _, resource := range Resourcedefs {
		if slices.Contains(resource.Types, resourceType) {
			resources = append(resources, resource)
		}
	}

	return resources
}
