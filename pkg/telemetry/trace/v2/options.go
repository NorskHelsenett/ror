package trace

import (
	"strings"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type traceConfig struct {
	endpoint string
	insecure *bool // nil = auto-detect from endpoint
	sampler  sdktrace.Sampler
}

func defaultConfig() traceConfig {
	return traceConfig{
		sampler: sdktrace.ParentBased(sdktrace.AlwaysSample()),
	}
}

func (c traceConfig) isInsecure() bool {
	if c.insecure != nil {
		return *c.insecure
	}
	return !strings.HasSuffix(c.endpoint, ":443")
}

// Option configures the tracer initialization.
type Option func(*traceConfig)

// WithEndpoint sets the OTLP gRPC collector endpoint (e.g. "localhost:4317").
func WithEndpoint(endpoint string) Option {
	return func(c *traceConfig) {
		c.endpoint = endpoint
	}
}

// WithInsecure disables TLS for the gRPC connection to the collector.
// By default, TLS is disabled unless the endpoint port is 443.
func WithInsecure() Option {
	return func(c *traceConfig) {
		insecure := true
		c.insecure = &insecure
	}
}

// WithSampler overrides the default sampler (ParentBased(AlwaysSample)).
func WithSampler(sampler sdktrace.Sampler) Option {
	return func(c *traceConfig) {
		c.sampler = sampler
	}
}
