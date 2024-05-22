package configconsts

const (
	ROLE                 = "ROLE"
	HTTP_PORT            = "HTTP_PORT"
	HTTP_TIMEOUT         = "HTTP_TIMEOUT"
	HEALTH_ENDPOINT      = "HEALTH_ENDPOINT"
	API_KEY_SALT         = "API_KEY_SALT"
	API_KEY              = "API_KEY"
	API_KEY_SECRET       = "API_KEY_SECRET"
	API_ENDPOINT         = "ROR_URL"
	DEVELOPMENT          = "DEVELOPMENT"
	LOCALHOST            = "LOCALHOST"
	PORT                 = "PORT"
	POD_NAMESPACE        = "POD_NAMESPACE"
	CLUSTER_ID           = "CLUSTER_ID"
	ERROR_COUNT          = "ERROR_COUNT"
	HELSEGITLAB_BASE_URL = "HELSEGITLAB_BASE_URL"
	ENVIRONMENT          = "ENVIRONMENT"
	LDAP_CONFIGS         = "LDAP_CONFIGS"

	ROR_OPERATOR_NAMESPACE = "ROR_OPERATOR_NAMESPACE"

	LDAP_CERTIFICATE_FOLDER = "LDAP_CERTIFICATE_FOLDER"

	OIDC_PROVIDER           = "OIDC_PROVIDER"
	OIDC_CLIENT_ID          = "OIDC_CLIENT_ID"
	OIDC_DEVICE_CLIENT_ID   = "OIDC_DEVICE_CLIENT_ID"
	OIDC_SKIP_ISSUER_VERIFY = "SKIP_OIDC_ISSUER_VERIFY"

	VAULT_URL = "VAULT_URL"

	MONGODB_HOST     = "MONGODB_HOST"
	MONGODB_PORT     = "MONGODB_PORT"
	MONGODB_DATABASE = "MONGODB_DATABASE"
	MONGODB_URL      = "MONGODB_URL"
	MONGO_DATABASE   = "MONGO_DATABASE"

	RABBITMQ_HOST             = "RABBITMQ_HOST"
	RABBITMQ_PORT             = "RABBITMQ_PORT"
	RABBITMQ_BROADCAST_NAME   = "RABBITMQ_BROADCAST_NAME"
	RABBITMQ_CONNECTIONSTRING = "RABBITMQ_CONNECTIONSTRING"

	REDIS_HOST = "REDIS_HOST"
	REDIS_PORT = "REDIS_PORT"

	TRACER_ID                        = "TRACER_ID"
	ENABLE_TRACING                   = "ENABLE_TRACING"
	OPENTELEMETRY_COLLECTOR_ENDPOINT = "OPENTELEMETRY_COLLECTOR_ENDPOINT"

	PROFILER_ENABLED   = "PROFILER_ENABLED"
	STARTUP_SLEEP_TIME = "STARTUP_SLEEP_TIME"

	GIN_USE_CORS      = "USE_CORS"
	GIN_ALLOW_ORIGINS = "ALLOW_ORIGINS"

	VERSION           = "VERSION"
	COMMIT            = "COMMIT"
	DEX_HOST          = "DEX_HOST"
	DEX_PORT          = "DEX_PORT"
	DEX_GRPC_PORT     = "DEX_GRPC_PORT"
	DEX_CERT_FILEPATH = "DEX_CERT_FILEPATH"
	DEX_VAULT_PATH    = "DEX_VAULT_PATH"
	DEX_TLS           = "DEX_TLS"

	SLACK_BOT_TOKEN = "SLACK_BOT_TOKEN"
	SLACK_APP_TOKEN = "SLACK_APP_TOKEN"

	CONTAINER_REG_PREFIX     = "CONTAINER_REG_PREFIX"
	CONTAINER_REG_IMAGE_PATH = "CONTAINER_REG_IMAGE_PATH"
	CONTAINER_REG_HELM_PATH  = "CONTAINER_REG_HELM_PATH"

	// Operator
	OPERATOR_BACKOFF_LIMIT       = "BACKOFF_LIMIT"
	OPERATOR_DEADLINE_SECONDS    = "DEADLINE_SECONDS"
	OPERATOR_JOB_SERVICE_ACCOUNT = "JOB_SERVICE_ACCOUNT"
	OPERATOR_APPLOG_SECRET_NAME  = "APPLOG_SECRET_NAME"

	// Tanzu agent
	TANZU_AGENT_KUBECONFIG         = "KUBECONFIG"
	TANZU_AGENT_DELETE_KUBECONFIG  = "DELETE_KUBECONFIG"
	TANZU_AGENT_KUBE_VSPHERE_PATH  = "KUBE_VSPHERE_PATH"
	TANZU_AGENT_KUBECTL_PATH       = "KUBECTL_PATH"
	TANZU_AGENT_DATACENTER_URL     = "DATACENTER_URL"
	TANZU_AGENT_USERNAME           = "TANZU_USERNAME"
	TANZU_AGENT_PASSWORD           = "TANZU_PWD"
	TANZU_AGENT_TOKEN_EXPIRY       = "TOKEN_EXPIRY"
	TANZU_AGENT_LOGIN_EVERY_MINUTE = "LOGIN_EVERY_MINUTE"
	TANZU_AGENT_HEALTH_PORT        = "HTTP_PORT"
	TANZU_AGENT_TANZU_ACCESS       = "TANZU_ACCESS"
	TANZU_AGENT_DATACENTER         = "DATACENTER"

	// Tanzu Auth
	TANZU_AUTH_HEALTH_PORT        = "HTTP_PORT"
	TANZU_AUTH_KUBE_VSPHERE_PATH  = "KUBE_VSPHERE_PATH"
	TANZU_AUTH_KUBECTL_PATH       = "KUBECTL_PATH"
	TANZU_AUTH_BASE_URL           = "TANZU_AUTH_BASE_URL"
	TANZU_AUTH_CONFIG_FOLDER_PATH = "TANZU_AUTH_CONFIG_FOLDER_PATH"

	// ms
	MS_HTTP_PORT      = "HTTP_PORT"
	MS_HTTP_BIND_PORT = "HTTP_BIND_PORT"

	LOCAL_KUBERNETES_ROR_BASE_URL = "LOCAL_KUBERNETES_ROR_BASE_URL"
)
