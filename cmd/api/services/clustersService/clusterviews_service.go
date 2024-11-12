package clustersservice

import (
	"context"
	"fmt"
	resourcesservice "github.com/NorskHelsenett/ror/cmd/api/services/resourcesService"
	"github.com/NorskHelsenett/ror/internal/models/viewsmodels"
	viewsrepo "github.com/NorskHelsenett/ror/internal/mongodbrepo/repositories/viewsRepo"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

func GetViewPolicyreport(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (viewsmodels.PolicyreportView, error) {
	var policyReportsView viewsmodels.PolicyreportView

	policyreports, _ := resourcesservice.GetPolicyreports(ctx, ownerref)
	err := policyReportsView.ImportData(policyreports)
	if err != nil {
		return policyReportsView, err
	}
	return policyReportsView, nil
}

func GetViewVulnerabilityReports(ctx context.Context, ownerref apiresourcecontracts.ResourceOwnerReference) (viewsmodels.VulnerabilityReportsView, error) {
	var vulnerabilityReportsView viewsmodels.VulnerabilityReportsView

	vulnerabilityreports, _ := resourcesservice.GetVulnerabilityreports(ctx, ownerref)
	err := vulnerabilityReportsView.ImportData(vulnerabilityreports)
	if err != nil {
		return vulnerabilityReportsView, err
	}
	return vulnerabilityReportsView, nil
}

func GetViewVulnerabilityReportsById(ctx context.Context, cveId string) (viewsmodels.VulnerabilityById, error) {
	var view viewsmodels.VulnerabilityById

	data, err := viewsrepo.GetViewVulnerabilityReportsById(ctx, cveId)

	view.ImportData(data)
	if err != nil {
		return view, err
	}
	return view, nil
}

func GetGlobalViewVulnerabilityReportsById(ctx context.Context, cveId string) ([]viewsmodels.GlobalVulnerabilityReportsViewById, error) {
	data, err := viewsrepo.GetGlobalViewVulnerabilityReportsById(ctx, cveId)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetViewVulnerabilityReportsGlobal(ctx context.Context) ([]viewsmodels.VulnerabilityReportsView, error) {
	data, err := viewsrepo.GetViewVulnerabilityReportsGlobal(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetViewPolicyReportSummary(ctx context.Context, querytype viewsmodels.PolicyreportGlobalQueryType, clusterID string) ([]viewsmodels.PolicyreportGlobal, error) {
	var ReportsView []viewsmodels.PolicyreportGlobal
	query := viewsmodels.PolicyreportGlobalQuery{
		Type:     querytype,
		Internal: false,
	}
	ReportsView, err := viewsrepo.GetPolicyReportSummary(ctx, query, clusterID)
	if err != nil {
		return ReportsView, err
	}
	return ReportsView, nil
}

func GetClusterComplianceReports(ctx context.Context, clusterId string) ([]viewsmodels.ComplianceReport, error) {
	reports, err := viewsrepo.GetComplianceReports(ctx, clusterId)
	if err != nil {
		return nil, fmt.Errorf("error when fetching compliance reports from repository: %w", err)
	}
	return reports, nil
}

func GetClusterComplianceReportsGlobal(ctx context.Context) ([]viewsmodels.ComplianceReport, error) {
	reports, err := viewsrepo.GetComplianceReportsGlobal(ctx)
	if err != nil {
		return nil, fmt.Errorf("error when fetching compliance reports from repository: %w", err)
	}
	return reports, nil
}
