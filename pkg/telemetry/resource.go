package telemetry

import (
	"context"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"
)

// newResource creates an OpenTelemetry resource with service information and attributes.
//
// This resource contains identifying information about the service such as its name,
// version, and environment, which will be attached to all telemetry data (traces, metrics, logs).
//
// Parameters:
//   - ctx: The context for resource creation operations
//   - opts: Configuration options that may contain service information
//
// Returns:
//   - *resource.Resource: A configured OpenTelemetry resource with service attributes
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
