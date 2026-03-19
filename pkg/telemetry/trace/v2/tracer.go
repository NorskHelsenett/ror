package trace

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"

	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/credentials"
)

var (
	globalTracer *tracer
	initOnce     sync.Once
	initErr      error
)

type tracer struct {
	provider *sdktrace.TracerProvider
	tracer   oteltrace.Tracer
}

// Init initializes the global tracer singleton. It is safe to call multiple
// times; only the first call takes effect. Subsequent calls return the error
// (if any) from the first initialization.
func Init(ctx context.Context, serviceName string, opts ...Option) error {
	initOnce.Do(func() {
		initErr = initTracer(ctx, serviceName, opts...)
	})
	return initErr
}

func initTracer(ctx context.Context, serviceName string, opts ...Option) error {
	cfg := defaultConfig()
	for _, o := range opts {
		o(&cfg)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return fmt.Errorf("trace: failed to create resource: %w", err)
	}

	exporterOpts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(cfg.endpoint),
	}
	if cfg.isInsecure() {
		exporterOpts = append(exporterOpts, otlptracegrpc.WithInsecure())
	} else {
		exporterOpts = append(exporterOpts, otlptracegrpc.WithTLSCredentials(
			credentials.NewTLS(&tls.Config{MinVersion: tls.VersionTLS12}),
		))
	}

	exporter, err := otlptracegrpc.New(ctx, exporterOpts...)
	if err != nil {
		return fmt.Errorf("trace: failed to create exporter: %w", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(cfg.sampler),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)

	otel.SetTracerProvider(provider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	globalTracer = &tracer{
		provider: provider,
		tracer:   provider.Tracer(serviceName),
	}
	return nil
}

// Shutdown gracefully shuts down the tracer provider, flushing any pending spans.
func Shutdown(ctx context.Context) error {
	if globalTracer == nil {
		return nil
	}
	return globalTracer.provider.Shutdown(ctx)
}

// InitDefault initializes the global tracer using the ROLE config as the
// service name and OPENTELEMETRY_COLLECTOR_ENDPOINT as the endpoint.
// Additional options can be passed to override defaults.
func InitWithDefault(ctx context.Context, opts ...Option) error {
	serviceName := rorconfig.GetString(rorconfig.ROLE)
	endpoint := rorconfig.GetString(rorconfig.OPENTELEMETRY_COLLECTOR_ENDPOINT)
	if endpoint != "" {
		opts = append([]Option{WithEndpoint(endpoint)}, opts...)
	}
	return Init(ctx, serviceName, opts...)
}
