package rorconfig

import (
	"slices"
)

type ConfigConst string

const (
	ROLE               ConfigConst = "ROLE"
	HTTP_HOST          ConfigConst = "HTTP_HOST"
	HTTP_PORT          ConfigConst = "HTTP_PORT"
	HTTP_HEALTH_HOST   ConfigConst = "HTTP_HEALTH_HOST"
	HTTP_HEALTH_PORT   ConfigConst = "HTTP_HEALTH_PORT"
	HTTP_TIMEOUT       ConfigConst = "HTTP_TIMEOUT"
	HTTP_USE_CORS      ConfigConst = "HTTP_USE_CORS"
	HTTP_ALLOW_ORIGINS ConfigConst = "HTTP_ALLOW_ORIGINS"
	ROR_API_KEY_SALT   ConfigConst = "API_KEY_SALT"
	ROR_API_KEY        ConfigConst = "ROR_API_KEY"
	ROR_API_KEY_SECRET ConfigConst = "ROR_API_KEY_SECRET"
	ROR_API_ENDPOINT   ConfigConst = "ROR_API_ENDPOINT"
	DEVELOPMENT        ConfigConst = "DEVELOPMENT"
	ENVIRONMENT        ConfigConst = "ENVIRONMENT"

	ERROR_COUNT ConfigConst = "ERROR_COUNT"

	POD_NAMESPACE ConfigConst = "POD_NAMESPACE"
	CLUSTER_ID    ConfigConst = "CLUSTER_ID"

	GIT_REPO_URL ConfigConst = "GIT_REPO_URL"
	GIT_BRANCH   ConfigConst = "GIT_BRANCH"
	GIT_TOKEN    ConfigConst = "GIT_TOKEN"
	GIT_PATH     ConfigConst = "GIT_PATH"

	LDAP_CONFIGS            ConfigConst = "LDAP_CONFIGS"
	LDAP_CERTIFICATE_FOLDER ConfigConst = "LDAP_CERTIFICATE_FOLDER"

	OIDC_PROVIDER           ConfigConst = "OIDC_PROVIDER"
	OIDC_CLIENT_ID          ConfigConst = "OIDC_CLIENT_ID"
	OIDC_DEVICE_CLIENT_ID   ConfigConst = "OIDC_DEVICE_CLIENT_ID"
	OIDC_SKIP_ISSUER_VERIFY ConfigConst = "OIDC_SKIP_ISSUER_VERIFY"

	VAULT_URL ConfigConst = "VAULT_URL"

	MONGODB_HOST     ConfigConst = "MONGODB_HOST"
	MONGODB_PORT     ConfigConst = "MONGODB_PORT"
	MONGODB_DATABASE ConfigConst = "MONGODB_DATABASE"
	MONGO_DATABASE   ConfigConst = "MONGO_DATABASE"
	MONGODB_URL      ConfigConst = "MONGODB_URL"

	RABBITMQ_HOST             ConfigConst = "RABBITMQ_HOST"
	RABBITMQ_PORT             ConfigConst = "RABBITMQ_PORT"
	RABBITMQ_BROADCAST_NAME   ConfigConst = "RABBITMQ_BROADCAST_NAME"
	RABBITMQ_CONNECTIONSTRING ConfigConst = "RABBITMQ_CONNECTIONSTRING"

	KV_HOST ConfigConst = "KV_HOST"
	KV_PORT ConfigConst = "KV_PORT"

	TRACER_ID                        ConfigConst = "TRACER_ID"
	ENABLE_TRACING                   ConfigConst = "ENABLE_TRACING"
	OPENTELEMETRY_COLLECTOR_ENDPOINT ConfigConst = "OPENTELEMETRY_COLLECTOR_ENDPOINT"

	PROFILER_ENABLED   ConfigConst = "PROFILER_ENABLED"
	STARTUP_SLEEP_TIME ConfigConst = "STARTUP_SLEEP_TIME"

	VERSION ConfigConst = "VERSION"
	COMMIT  ConfigConst = "COMMIT"

	// ms
	MS_ENDPOINT  ConfigConst = "MS_ENDPOINT"
	MS_HTTP_PORT ConfigConst = "MS_HTTP_PORT"

	ENABLE_PPROF ConfigConst = "ENABLE_PPROF"
)

type EnvironmentVariableConfig struct {
	key         string
	deprecated  bool
	description string
}

type EnvironmentVariables []EnvironmentVariableConfig

// ...existing code...
var ConfigConsts = EnvironmentVariables{
	{key: "ROLE", deprecated: false, description: ""},
	{key: "HTTP_HOST", deprecated: false, description: ""},
	{key: "HTTP_PORT", deprecated: false, description: ""},
	{key: "HTTP_HEALTH_HOST", deprecated: false, description: ""},
	{key: "HTTP_HEALTH_PORT", deprecated: false, description: ""},
	{key: "HTTP_TIMEOUT", deprecated: false, description: ""},
	{key: "API_KEY_SALT", deprecated: false, description: ""},
	{key: "API_KEY", deprecated: false, description: ""},
	{key: "API_KEY_SECRET", deprecated: false, description: ""},
	{key: "ROR_URL", deprecated: false, description: ""},
	{key: "DEVELOPMENT", deprecated: false, description: ""},
	{key: "PORT", deprecated: false, description: ""},
	{key: "POD_NAMESPACE", deprecated: false, description: ""},
	{key: "CLUSTER_ID", deprecated: false, description: ""},
	{key: "ERROR_COUNT", deprecated: false, description: ""},
	{key: "HELSEGITLAB_BASE_URL", deprecated: false, description: ""},
	{key: "ENVIRONMENT", deprecated: true, description: "Dont use ENVIRONMENT as its ambigous. Use DEVELOPMENT instead."},
	{key: "LDAP_CONFIGS", deprecated: false, description: ""},
	{key: "GIT_REPO_URL", deprecated: false, description: ""},
	{key: "GIT_BRANCH", deprecated: false, description: ""},
	{key: "GIT_TOKEN", deprecated: false, description: ""},
	{key: "GIT_PATH", deprecated: false, description: ""},
	{key: "ROR_OPERATOR_NAMESPACE", deprecated: false, description: ""},
	{key: "LDAP_CERTIFICATE_FOLDER", deprecated: false, description: ""},
	{key: "OIDC_PROVIDER", deprecated: false, description: ""},
	{key: "OIDC_CLIENT_ID", deprecated: false, description: ""},
	{key: "OIDC_DEVICE_CLIENT_ID", deprecated: false, description: ""},
	{key: "SKIP_OIDC_ISSUER_VERIFY", deprecated: false, description: ""},
	{key: "VAULT_URL", deprecated: false, description: ""},
	{key: "MONGODB_HOST", deprecated: false, description: ""},
	{key: "MONGODB_PORT", deprecated: false, description: ""},
	{key: "MONGODB_DATABASE", deprecated: false, description: ""},
	{key: "MONGODB_URL", deprecated: false, description: ""},
	{key: "MONGO_DATABASE", deprecated: false, description: ""},
	{key: "RABBITMQ_HOST", deprecated: false, description: ""},
	{key: "RABBITMQ_PORT", deprecated: false, description: ""},
	{key: "RABBITMQ_BROADCAST_NAME", deprecated: false, description: ""},
	{key: "RABBITMQ_CONNECTIONSTRING", deprecated: false, description: ""},
	{key: "KV_HOST", deprecated: false, description: ""},
	{key: "KV_PORT", deprecated: false, description: ""},
	{key: "TRACER_ID", deprecated: false, description: ""},
	{key: "ENABLE_TRACING", deprecated: false, description: ""},
	{key: "OPENTELEMETRY_COLLECTOR_ENDPOINT", deprecated: false, description: ""},
	{key: "PROFILER_ENABLED", deprecated: false, description: ""},
	{key: "STARTUP_SLEEP_TIME", deprecated: false, description: ""},
	{key: "USE_CORS", deprecated: false, description: ""},
	{key: "ALLOW_ORIGINS", deprecated: false, description: ""},
	{key: "VERSION", deprecated: false, description: ""},
	{key: "COMMIT", deprecated: false, description: ""},
	{key: "DEX_HOST", deprecated: false, description: ""},
	{key: "DEX_PORT", deprecated: false, description: ""},
	{key: "DEX_GRPC_PORT", deprecated: false, description: ""},
	{key: "DEX_CERT_FILEPATH", deprecated: false, description: ""},
	{key: "DEX_VAULT_PATH", deprecated: false, description: ""},
	{key: "DEX_TLS", deprecated: false, description: ""},
	{key: "SLACK_BOT_TOKEN", deprecated: false, description: ""},
	{key: "SLACK_APP_TOKEN", deprecated: false, description: ""},
	{key: "CONTAINER_REG_PREFIX", deprecated: false, description: ""},
	{key: "CONTAINER_REG_IMAGE_PATH", deprecated: false, description: ""},
	{key: "CONTAINER_REG_HELM_PATH", deprecated: false, description: ""},
	{key: "BACKOFF_LIMIT", deprecated: false, description: ""},
	{key: "DEADLINE_SECONDS", deprecated: false, description: ""},
	{key: "JOB_SERVICE_ACCOUNT", deprecated: false, description: ""},
	{key: "APPLOG_SECRET_NAME", deprecated: false, description: ""},
	{key: "KUBECONFIG", deprecated: false, description: ""},
	{key: "DELETE_KUBECONFIG", deprecated: false, description: ""},
	{key: "KUBE_VSPHERE_PATH", deprecated: false, description: ""},
	{key: "KUBECTL_PATH", deprecated: false, description: ""},
	{key: "DATACENTER_URL", deprecated: false, description: ""},
	{key: "TANZU_USERNAME", deprecated: false, description: ""},
	{key: "TANZU_PWD", deprecated: false, description: ""},
	{key: "TOKEN_EXPIRY", deprecated: false, description: ""},
	{key: "LOGIN_EVERY_MINUTE", deprecated: false, description: ""},
	{key: "HTTP_PORT", deprecated: false, description: ""},
	{key: "TANZU_ACCESS", deprecated: false, description: ""},
	{key: "DATACENTER", deprecated: false, description: ""},
	{key: "PROVIDER", deprecated: false, description: ""},
	{key: "REGION", deprecated: false, description: ""},
	{key: "HTTP_PORT", deprecated: false, description: ""},
	{key: "KUBE_VSPHERE_PATH", deprecated: false, description: ""},
	{key: "KUBECTL_PATH", deprecated: false, description: ""},
	{key: "TANZU_AUTH_BASE_URL", deprecated: false, description: ""},
	{key: "TANZU_AUTH_CONFIG_FOLDER_PATH", deprecated: false, description: ""},
	{key: "HTTP_PORT", deprecated: false, description: ""},
	{key: "HTTP_BIND_PORT", deprecated: false, description: ""},
	{key: "LOCAL_KUBERNETES_ROR_BASE_URL", deprecated: false, description: ""},
	{key: "ENABLE_PPROF", deprecated: false, description: ""},
}

func (ev *EnvironmentVariables) IsSet(val string) bool {
	return slices.ContainsFunc(*ev, func(data EnvironmentVariableConfig) bool {
		return data.key == val
	})
}

// GetEnvVariableByKey returns the environment variable configration for the provided env var.
func (ev *EnvironmentVariables) GetEnvVariableConfigByKey(key string) (EnvironmentVariableConfig, bool) {
	for _, data := range *ev {
		if data.key == key {
			return data, true
		}
	}
	// Return a default config if not found
	return EnvironmentVariableConfig{}, false
}

func (ev *EnvironmentVariables) Add(key string) {
	if !ev.IsSet(key) {
		data := EnvironmentVariableConfig{
			key:         key,
			deprecated:  false,
			description: "Local env variable not in central list",
		}
		*ev = append(*ev, data)
	}
}

func (ev *EnvironmentVariables) IsDeprecated(key string) bool {
	var v EnvironmentVariableConfig
	v, _ = ev.GetEnvVariableConfigByKey(key)
	return v.deprecated
}

func (ev *EnvironmentVariables) GetDescription(key string) string {
	var v EnvironmentVariableConfig
	v, _ = ev.GetEnvVariableConfigByKey(key)
	return v.description
}
