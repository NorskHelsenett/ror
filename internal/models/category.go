package models

type AuditCategory string

const (
	AuditCategoryUnknown         AuditCategory = "Unknown"
	AuditCategoryDatacenter      AuditCategory = "Datacenter"
	AuditCategoryPrice           AuditCategory = "Price"
	AuditCategoryProject         AuditCategory = "Project"
	AuditCategoryWorkspace       AuditCategory = "Workspace"
	AuditCategoryConfiguration   AuditCategory = "Configuration"
	AuditCategoryClusterMetadata AuditCategory = "ClusterMetadata"
	AuditCategoryApikey          AuditCategory = "ApiKey"
	AuditCategoryAcl             AuditCategory = "Acl"
	AuditCategorySwitchboard     AuditCategory = "Ruleset"
	AuditCategoryKubeconfig      AuditCategory = "Kubeconfig"
)
