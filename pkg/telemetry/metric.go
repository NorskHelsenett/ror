package telemetry

import (
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

// creates a tracer provider with prometheus exporter
// TODO: Add support for multiple exporters
func newMeterProvider(res *resource.Resource) (*metric.MeterProvider, error) {
	metricExporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metricExporter),
	)
	return meterProvider, nil
}
