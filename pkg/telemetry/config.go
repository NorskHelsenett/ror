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
