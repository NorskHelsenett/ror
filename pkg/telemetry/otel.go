package telemetry

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

func SetupOTel(ctx context.Context, opts ...Option) (shutdown func(context.Context) error, err error) {
	// we enable all signals by default
	cfg := config{
		WithLogger: true,
		WithMeter:  true,
		WithTracer: true,
	}
	cfg = applyEnvOptions(cfg)
	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}

	var shutdownFuncs []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	res := newResource(ctx, opts...)

	if cfg.WithTracer {
		tracerProvider, err := newTracerProvider(ctx, res)
		if err != nil {
			handleErr(err)
			return
		}
		shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
		otel.SetTracerProvider(tracerProvider)
	}

	if cfg.WithMeter {
		meterProvider, err := newMeterProvider(res)
		if err != nil {
			handleErr(err)
			return
		}
		shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
		otel.SetMeterProvider(meterProvider)
	}

	if cfg.WithLogger {
		loggerProvider, err := newLoggerProvider(res)
		if err != nil {
			handleErr(err)
			return
		}
		shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
		global.SetLoggerProvider(loggerProvider)
	}

	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

// add resource information to logs, metrics and traces
func newResource(ctx context.Context, opts ...Option) *resource.Resource {
	attributes := []attribute.KeyValue{}
	attributes = append(attributes,
		semconv.ServiceInstanceID(
			uuid.NewString(),
		),
	)

	cfg := config{}
	cfg = applyEnvOptions(cfg)
	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}

	if cfg.ServiceName != "" {
		attributes = append(attributes,
			semconv.ServiceName(cfg.ServiceName),
		)
	}

	if cfg.DeploymentEnvironment != "" {
		attributes = append(attributes,
			semconv.DeploymentEnvironmentName(cfg.DeploymentEnvironment),
		)
	}
	// We do not care if any detectors error out, so we ignore the returned error
	res, _ := resource.New(ctx,
		resource.WithAttributes(attributes...),

		resource.WithFromEnv(),

		resource.WithContainer(),
		resource.WithContainerID(),

		resource.WithHost(),
		resource.WithHostID(),

		resource.WithOS(),
		resource.WithOSDescription(),
		resource.WithOSType(),

		resource.WithProcess(),
		resource.WithProcessCommandArgs(),
		resource.WithProcessExecutableName(),
		resource.WithProcessExecutablePath(),
		resource.WithProcessOwner(),
		resource.WithProcessPID(),
		resource.WithProcessRuntimeDescription(),
		resource.WithProcessRuntimeName(),
		resource.WithProcessRuntimeVersion(),

		resource.WithTelemetrySDK(),
	)

	return res

}

// creates a tracer provider with otlpgrpc exporter
// TODO: Add support for multiple exporters
func newTracerProvider(ctx context.Context, res *resource.Resource) (*trace.TracerProvider, error) {
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	tracerProvider := trace.NewTracerProvider(
		trace.WithResource(res),
		trace.WithBatcher(traceExporter),
	)
	return tracerProvider, nil
}

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

// creates a logger provider with stdout exporter
// TODO: Add support for multiple exporters
func newLoggerProvider(res *resource.Resource) (*log.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()
	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(
		log.WithResource(res),
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}
