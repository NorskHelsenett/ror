package telemetry

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

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
