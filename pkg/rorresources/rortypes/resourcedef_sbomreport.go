package rortypes

// Trivy SBOM Report structure based on aquasecurity.github.io/v1alpha1
type ResourceSbomReport struct {
	Report ResourceSbomReportsReport `json:"report"`
}

type ResourceSbomReportsReport struct {
	Scanner         AquaReportScanner           `json:"scanner"`
	Artifact        ResourceSbomReportsArtifact `json:"artifact"`
	Registry        ResourceSbomReportsRegistry `json:"registry,omitempty"`
	Summary         ResourceSbomReportsSummary  `json:"summary"`
	Components      ResourceSbomReportsBom      `json:"components"`
	UpdateTimestamp string                      `json:"updateTimestamp"`
}

type ResourceSbomReportsArtifact struct {
	Repository string `json:"repository,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Digest     string `json:"digest,omitempty"`
	MimeType   string `json:"mimeType,omitempty"`
}

type ResourceSbomReportsRegistry struct {
	Server string `json:"server,omitempty"`
}

type ResourceSbomReportsSummary struct {
	ComponentsCount   int `json:"componentsCount"`
	DependenciesCount int `json:"dependenciesCount"`
}

type ResourceSbomReportsBom struct {
	BomFormat    string                            `json:"bomFormat"`
	SpecVersion  string                            `json:"specVersion"`
	SerialNumber string                            `json:"serialNumber,omitempty"`
	Version      int                               `json:"version,omitempty"`
	Metadata     ResourceSbomReportsBomMetadata    `json:"metadata,omitempty"`
	Components   []ResourceSbomReportsComponent    `json:"components,omitempty"`
	Dependencies []ResourceSbomReportsComponentDep `json:"dependencies,omitempty"`
}

type ResourceSbomReportsBomMetadata struct {
	Timestamp string                              `json:"timestamp,omitempty"`
	Tools     ResourceSbomReportsBomMetadataTools `json:"tools,omitempty"`
	Component ResourceSbomReportsComponent        `json:"component,omitempty"`
}

type ResourceSbomReportsBomMetadataTools struct {
	Components []ResourceSbomReportsComponent `json:"components,omitempty"`
}

type ResourceSbomReportsComponent struct {
	BomRef     string                                 `json:"bom-ref,omitempty"`
	Type       string                                 `json:"type,omitempty"`
	Name       string                                 `json:"name,omitempty"`
	Group      string                                 `json:"group,omitempty"`
	Version    string                                 `json:"version,omitempty"`
	Purl       string                                 `json:"purl,omitempty"`
	Hashes     []ResourceSbomReportsComponentHash     `json:"hashes,omitempty"`
	Licenses   []ResourceSbomReportsComponentLicense  `json:"licenses,omitempty"`
	Properties []ResourceSbomReportsComponentProperty `json:"properties,omitempty"`
	Supplier   ResourceSbomReportsComponentSupplier   `json:"supplier,omitempty"`
}

type ResourceSbomReportsComponentHash struct {
	Alg     string `json:"alg,omitempty"`
	Content string `json:"content,omitempty"`
}

type ResourceSbomReportsComponentLicense struct {
	Expression string                                     `json:"expression,omitempty"`
	License    ResourceSbomReportsComponentLicenseDetails `json:"license,omitempty"`
}

type ResourceSbomReportsComponentLicenseDetails struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type ResourceSbomReportsComponentProperty struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type ResourceSbomReportsComponentSupplier struct {
	Name    string                                        `json:"name,omitempty"`
	URL     []string                                      `json:"url,omitempty"`
	Contact []ResourceSbomReportsComponentSupplierContact `json:"contact,omitempty"`
}

type ResourceSbomReportsComponentSupplierContact struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

type ResourceSbomReportsComponentDep struct {
	Ref       string   `json:"ref,omitempty"`
	DependsOn []string `json:"dependsOn,omitempty"`
}
