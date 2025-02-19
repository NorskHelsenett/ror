package telemetry

import "os"

type config struct {
	ServiceName           string
	DeploymentEnvironment string
	WithLogger            bool
	WithMeter             bool
	WithTracer            bool
}

type Option interface {
	apply(config config) config
}

type serviceNameOption string

func (o serviceNameOption) apply(cfg config) config {
	cfg.ServiceName = string(o)
	return cfg
}

func WithServiceName(name string) Option {
	return serviceNameOption(name)
}

type deploymentEnvironmentOption string

func (o deploymentEnvironmentOption) apply(cfg config) config {
	cfg.DeploymentEnvironment = string(o)
	return cfg
}

func WithDeploymentEnvironment(environment string) Option {
	return deploymentEnvironmentOption(environment)
}

type withLoggerOption bool

func (o withLoggerOption) apply(cfg config) config {
	cfg.WithLogger = bool(o)
	return cfg
}

func WithLogger() Option {
	return withLoggerOption(true)
}

type withMeterOption bool

func (o withMeterOption) apply(cfg config) config {
	cfg.WithMeter = bool(o)
	return cfg
}

func WithMeter() Option {
	return withMeterOption(true)
}

type withTracerOption bool

func (o withTracerOption) apply(cfg config) config {
	cfg.WithTracer = bool(o)
	return cfg
}

func WithTracer() Option {
	return withTracerOption(true)
}

func applyEnvOptions(cfg config) config {
	opts := getOptionsFromEnv()
	for _, opt := range opts {
		cfg = opt.apply(cfg)
	}
	return cfg
}

func getOptionsFromEnv() []Option {
	opts := []Option{}
	serviceName, ok := os.LookupEnv("SERVICE_NAME")
	if ok {
		opts = append(opts, serviceNameOption(serviceName))
	}
	deploymentEnvironment, ok := os.LookupEnv("DEPLOYMENT_ENVIRONMENT")
	if ok {
		opts = append(opts, deploymentEnvironmentOption(deploymentEnvironment))
	}
	return opts
}
