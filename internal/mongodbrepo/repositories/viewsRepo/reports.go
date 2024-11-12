package viewsrepo

import (
	"context"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/models/viewsmodels"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	ResourceCollectionName = "resources"
)

func GetPolicyReportSummary(ctx context.Context, query viewsmodels.PolicyreportGlobalQuery, clusterID string) ([]viewsmodels.PolicyreportGlobal, error) {

	kind := "PolicyReport"
	apiversion := "wgpolicyk8s.io/v1alpha2"

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"kind": kind, "apiversion": apiversion, "internal": query.Internal}})
	switch query.Type {
	case viewsmodels.PolicyreportGlobalQueryTypePolicy:
		aggregationPipeline = append(aggregationPipeline, bson.M{"$group": bson.M{"_id": bson.M{"cluster": "$owner.subject", "policy": "$resource.metadata.name"}, "pass": bson.M{"$sum": "$resource.summary.pass"}, "fail": bson.M{"$sum": "$resource.summary.fail"}}})
		aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "cluster": "$_id.cluster", "policy": "$_id.policy", "fail": "$fail", "pass": "$pass"}})
		aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"cluster": bson.M{"$regex": clusterID}}})
	case viewsmodels.PolicyreportGlobalQueryTypeCluster:
		aggregationPipeline = append(aggregationPipeline, bson.M{"$group": bson.M{"_id": bson.M{"cluster": "$owner.subject"}, "pass": bson.M{"$sum": "$resource.summary.pass"}, "fail": bson.M{"$sum": "$resource.summary.fail"}}})
		aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "cluster": "$_id.cluster", "fail": "$fail", "pass": "$pass"}})
	default:
		return []viewsmodels.PolicyreportGlobal{}, fmt.Errorf("unknown type %s", query.Type)
	}
	results := make([]viewsmodels.PolicyreportGlobal, 0)
	err := mongodb.Aggregate(ctx, ResourceCollectionName, aggregationPipeline, &results)

	return results, err
}

func GetViewVulnerabilityReportsById(ctx context.Context, cveId string) ([]viewsmodels.VulnerabilityByIdDB, error) {
	kind := "VulnerabilityReport"
	apiversion := "aquasecurity.github.io/v1alpha1"

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"kind": kind, "apiversion": apiversion}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$unwind": "$resource.report.vulnerabilities"})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"resource.report.vulnerabilities.vulnerabilityid": cveId}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "owner": 1, "internal": 1, "ownerref": "$resource.metadata.ownerreferences", "scanner": "$resource.report.scanner", "artifact": "$resource.report.artifact", "vulnerabilities": "$resource.report.vulnerabilities"}})

	results := make([]viewsmodels.VulnerabilityByIdDB, 0)
	err := mongodb.Aggregate(ctx, ResourceCollectionName, aggregationPipeline, &results)

	return results, err
}

func GetGlobalViewVulnerabilityReportsById(ctx context.Context, cveId string) ([]viewsmodels.GlobalVulnerabilityReportsViewById, error) {
	kind := "VulnerabilityReport"
	apiversion := "aquasecurity.github.io/v1alpha1"

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"kind": kind, "apiversion": apiversion, "resource.report.vulnerabilities.vulnerabilityid": cveId, "internal": false}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "clusterid": "$owner.subject"}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$group": bson.M{"_id": "$clusterid"}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$lookup": bson.M{"from": "clusters", "localField": "_id", "foreignField": "clusterid", "pipeline": []bson.M{{"$project": bson.M{"_id": 0, "environment": 1, "projectid": "$metadata.projectid", "criticality": "$metadata.criticality", "sensitivity": "$metadata.sensitivity"}}}, "as": "clusterLookup"}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$replaceRoot": bson.M{"newRoot": bson.M{"$mergeObjects": bson.A{bson.M{"$arrayElemAt": bson.A{"$clusterLookup", 0}}, "$$ROOT"}}}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$lookup": bson.M{"from": "projects", "localField": "projectid", "foreignField": "_id", "pipeline": []bson.M{{"$project": bson.M{"_id": 1, "name": 1}}}, "as": "projectLookup"}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$addFields": bson.M{"project": bson.M{"$arrayElemAt": bson.A{"$projectLookup", 0}}}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$set": bson.M{"clusterid": "$_id"}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$unset": bson.A{"clusterLookup", "projectLookup", "_id", "projectid"}})

	var results []viewsmodels.GlobalVulnerabilityReportsViewById
	err := mongodb.Aggregate(ctx, ResourceCollectionName, aggregationPipeline, &results)

	return results, err
}

func GetViewVulnerabilityReportsGlobal(ctx context.Context) ([]viewsmodels.VulnerabilityReportsView, error) {
	kind := "VulnerabilityReport"
	apiversion := "aquasecurity.github.io/v1alpha1"

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"kind": kind, "apiversion": apiversion, "internal": false}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$group": bson.M{"_id": "$owner.subject", "criticalCount": bson.M{"$sum": "$resource.report.summary.criticalcount"}, "highCount": bson.M{"$sum": "$resource.report.summary.highcount"}, "mediumCount": bson.M{"$sum": "$resource.report.summary.mediumcount"}, "lowCount": bson.M{"$sum": "$resource.report.summary.lowcount"}}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "clusterid": "$_id", "criticalCount": 1, "highCount": 1, "mediumCount": 1, "lowCount": 1}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$lookup": bson.M{"from": "clusters", "localField": "clusterid", "foreignField": "clusterid", "pipeline": []bson.M{{"$project": bson.M{"_id": 0, "environment": 1, "metadata.projectid": 1}}}, "as": "metadata"}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$addFields": bson.M{"environment": bson.M{"$arrayElemAt": bson.A{"$metadata.environment", 0}}, "projectid": bson.M{"$arrayElemAt": bson.A{"$metadata.metadata.projectid", 0}}}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$lookup": bson.M{"from": "projects", "localField": "projectid", "foreignField": "_id", "pipeline": []bson.M{{"$project": bson.M{"_id": 1, "name": 1}}}, "as": "projectArray"}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$addFields": bson.M{"project": bson.M{"$arrayElemAt": bson.A{"$projectArray", 0}}}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$unset": bson.A{"projectid", "projectArray", "metadata"}})

	results := make([]viewsmodels.VulnerabilityReportsView, 0)
	err := mongodb.Aggregate(ctx, ResourceCollectionName, aggregationPipeline, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetComplianceReports(ctx context.Context, clusterId string) ([]viewsmodels.ComplianceReport, error) {
	kind := "ClusterComplianceReport"
	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"kind": kind, "owner.subject": clusterId, "resource.metadata.name": bson.M{"$in": bson.A{"cis", "nsa"}}}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "clusterid": "$owner.subject", "metadata.name": "$resource.metadata.name", "metadata.title": "$resource.status.summaryreport.title", "summary.failcount": "$resource.status.summary.failcount", "summary.passcount": "$resource.status.summary.passcount", "reports": "$resource.status.summaryreport.controlcheck"}})

	results := []viewsmodels.ComplianceReport{}
	err := mongodb.Aggregate(ctx, ResourceCollectionName, aggregationPipeline, &results)
	if err != nil {
		return nil, fmt.Errorf("error performing mongodb aggregation: %w", err)
	}
	return results, nil
}

func GetComplianceReportsGlobal(ctx context.Context) ([]viewsmodels.ComplianceReport, error) {
	kind := "ClusterComplianceReport"
	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": bson.M{"kind": kind, "resource.metadata.name": bson.M{"$in": bson.A{"cis", "nsa"}}}})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": bson.M{"_id": 0, "clusterid": "$owner.subject", "metadata.name": "$resource.metadata.name", "summary.failcount": "$resource.status.summary.failcount", "summary.passcount": "$resource.status.summary.passcount"}})

	results := []viewsmodels.ComplianceReport{}
	err := mongodb.Aggregate(ctx, ResourceCollectionName, aggregationPipeline, &results)
	if err != nil {
		return nil, fmt.Errorf("error performing mongodb aggregation: %w", err)
	}
	return results, nil
}
