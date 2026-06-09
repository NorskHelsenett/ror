package rortypes

// K8s namepace struct
type ResourceClusterComplianceReport struct {
}

// (r ResourceClusterComplianceReport) Get returns a pointer to the resource of type ResourceClusterComplianceReport
func (r *ResourceClusterComplianceReport) Get() *ResourceClusterComplianceReport {
	return r
}

// ClusterComplianceReportinterface represents the interface for resources of the type clustercompliancereport
type ClusterComplianceReportinterface interface {
	Get() *ResourceClusterComplianceReport
}
