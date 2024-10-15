package metricsrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"

	"github.com/NorskHelsenett/ror/pkg/clients/mongodb"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MetricsCollectionName = "metrics"
)

func WriteMetrics(metricsreport *apicontracts.MetricsReport, clusterId string, ctx context.Context) {
	db := mongodb.GetMongoDb()
	queries := make([]interface{}, 0)

	for _, metricUpdate := range metricsreport.Nodes {
		query := bson.M{
			"metadata": bson.M{
				"type": "node", "clusterId": clusterId,
				"name": metricUpdate.Name,
			},
			"timestamp":        metricUpdate.TimeStamp,
			"cpuusage":         metricUpdate.CpuUsage,
			"cpuallocated":     metricUpdate.CpuAllocated,
			"cpupercentage":    metricUpdate.CpuPercentage,
			"memoryusage":      metricUpdate.MemoryUsage,
			"memoryallocated":  metricUpdate.MemoryAllocated,
			"memorypercentage": metricUpdate.MemoryPercentage,
		}
		queries = append(queries, query)
	}
	_, err := db.Collection(MetricsCollectionName).InsertMany(ctx, queries)
	if err != nil {
		msg := "could not insert metrics"
		rlog.Error(msg, err)
	}
}

// GetPodMetricsByFilter returns timeseries data filterd with the apicontracts.MetricsFilter.
//
//   - Type must be set to either node or pod.
//   - Default time is common era (year 0 - now()).
//   - Default time resolution is hour. all results
func GetMetricsByFilter(filter apicontracts.MetricsFilter, ctx context.Context, mongodb *mongo.Database) ([]apicontracts.MetricsResult, error) {
	var metricsresult []apicontracts.MetricsResult
	var startOfTime, startTime, stopTime time.Time

	if filter.Type != "node" && filter.Type != "pod" {
		return metricsresult, fmt.Errorf("type has to be node or pod, got: %s", filter.Type)
	}

	if filter.Time.Start != startOfTime {
		startTime = filter.Time.Start
	}

	if filter.Time.Stop != startOfTime {
		stopTime = filter.Time.Stop
	} else {
		stopTime = time.Now()
	}

	aggregationFilter := bson.M{"metadata.type": filter.Type, "timestamp": bson.M{"$gte": startTime, "$lte": stopTime}}

	if filter.Metadata.Name != "" {
		aggregationFilter["metadata.name"] = filter.Metadata.Name
	}
	if filter.Metadata.Namespace != "" {
		aggregationFilter["metadata.namespace"] = filter.Metadata.Namespace
	}
	if filter.Metadata.ClusterId != "" {
		aggregationFilter["metadata.clusterId"] = filter.Metadata.ClusterId
	}

	aggregationProjection := bson.M{"date": bson.M{"$dateToParts": bson.M{"date": "$timestamp"}}, "cpuusage": 1, "memoryusage": 1, "memorypercentage": 1, "cpupercentage": 1, "memoryallocated": 1, "cpuallocated": 1}
	aggregationProjection["metadata.name"] = 1
	aggregationProjection["metadata.namespace"] = 1
	aggregationProjection["metadata.clusterId"] = 1

	timeResolution := filter.Time.Resolution
	if timeResolution == 0 {
		timeResolution = 4
	}

	aggregationGroupingTime := bson.M{"year": "$date.year"}
	if timeResolution > 1 {
		aggregationGroupingTime["month"] = "$date.month"
	}
	if timeResolution > 2 {
		aggregationGroupingTime["day"] = "$date.day"
	}
	if timeResolution > 3 {
		aggregationGroupingTime["hour"] = "$date.hour"
	}
	if timeResolution > 4 {
		aggregationGroupingTime["minute"] = "$date.minute"
	}

	aggregationGrouping := bson.M{"date": aggregationGroupingTime}

	if filter.GroupBy.Name {
		aggregationGrouping["name"] = "$metadata.name"
		aggregationGrouping["namespace"] = "$metadata.namespace"
		aggregationGrouping["clusterId"] = "$metadata.clusterId"
	}

	if filter.GroupBy.Namespace {
		aggregationGrouping["namespace"] = "$metadata.namespace"
		aggregationGrouping["clusterId"] = "$metadata.clusterId"
	}

	if filter.GroupBy.Cluster {
		aggregationGrouping["clusterId"] = "$metadata.clusterId"
	}

	var aggregationPipeline []bson.M
	aggregationPipeline = append(aggregationPipeline, bson.M{"$match": aggregationFilter})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$project": aggregationProjection})
	aggregationPipeline = append(aggregationPipeline, bson.M{"$group": bson.M{
		"_id":                 aggregationGrouping,
		"avgCpu":              bson.M{"$avg": "$cpuusage"},
		"avgMemory":           bson.M{"$avg": "$memoryusage"},
		"avgPercentageMemory": bson.M{"$avg": "$memorypercentage"},
		"avgPercentageCpu":    bson.M{"$avg": "$cpupercentage"},
		"avgAllocatedMemory":  bson.M{"$avg": "$memoryallocated"},
		"avgAllocatedCpu":     bson.M{"$avg": "$cpuallocated"},
	}})
	results, err := mongodb.Collection("metrics").Aggregate(ctx, aggregationPipeline)
	if err != nil {
		return nil, fmt.Errorf("could not fetch metrics: %v", err)
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		_ = results.Close(ctx)
	}(results, ctx)

	if results.RemainingBatchLength() == 0 {
		return metricsresult, nil
	}

	for results.Next(ctx) {
		var singleMetric apicontracts.MetricsResult
		if err = results.Decode(&singleMetric); err != nil {
			return nil, fmt.Errorf("could not fetch metric: %v", err)
		}
		metricsresult = append(metricsresult, singleMetric)
	}

	return metricsresult, nil
}
