package rorconfig

type ConfigConst string

const (
	ROLE             ConfigConst = "ROLE"
	HTTP_HOST        ConfigConst = "HTTP_HOST"
	HTTP_PORT        ConfigConst = "HTTP_PORT"
	HTTP_HEALTH_HOST ConfigConst = "HTTP_HEALTH_HOST"
	HTTP_HEALTH_PORT ConfigConst = "HTTP_HEALTH_PORT"
	HTTP_TIMEOUT     ConfigConst = "HTTP_TIMEOUT"
	// Deprecated: use HTTP_HEALTH_HOST / HTTP_HEALTH_PORT instead
	HEALTH_ENDPOINT      ConfigConst = "HEALTH_ENDPOINT"
	API_KEY_SALT         ConfigConst = "API_KEY_SALT"
	API_KEY              ConfigConst = "API_KEY"
	API_KEY_SECRET       ConfigConst = "API_KEY_SECRET"
	API_ENDPOINT         ConfigConst = "ROR_URL"
	DEVELOPMENT          ConfigConst = "DEVELOPMENT"
	PORT                 ConfigConst = "PORT"
	POD_NAMESPACE        ConfigConst = "POD_NAMESPACE"
	CLUSTER_ID           ConfigConst = "CLUSTER_ID"
	ERROR_COUNT          ConfigConst = "ERROR_COUNT"
	HELSEGITLAB_BASE_URL ConfigConst = "HELSEGITLAB_BASE_URL"
	ENVIRONMENT          ConfigConst = "ENVIRONMENT"
	LDAP_CONFIGS         ConfigConst = "LDAP_CONFIGS"

	GIT_REPO_URL          ConfigConst = "GIT_REPO_URL"
	GIT_BRANCHConfigConst ConfigConst = "GIT_BRANCH"
	GIT_TOKEN             ConfigConst = "GIT_TOKEN"
	GIT_PATH              ConfigConst = "GIT_PATH"

	ROR_OPERATOR_NAMESPACE = "ROR_OPERATOR_NAMESPACE"

	LDAP_CERTIFICATE_FOLDER = "LDAP_CERTIFICATE_FOLDER"

	OIDC_PROVIDER                    ConfigConst = "OIDC_PROVIDER"
	OIDC_CLIENT_ID                   ConfigConst = "OIDC_CLIENT_ID"
	OIDC_DEVICE_CLIENT_IDConfigConst ConfigConst = "OIDC_DEVICE_CLIENT_ID"
	OIDC_SKIP_ISSUER_VERIFY          ConfigConst = "SKIP_OIDC_ISSUER_VERIFY"

	VAULT_URL = "VAULT_URL"

	MONGODB_HOST              ConfigConst = "MONGODB_HOST"
	MONGODB_PORT              ConfigConst = "MONGODB_PORT"
	MONGODB_DATABASE          ConfigConst = "MONGODB_DATABASE"
	MONGODB_URL               ConfigConst = "MONGODB_URL"
	MONGO_DATABASEConfigConst ConfigConst = "MONGO_DATABASE"

	RABBITMQ_HOST                      ConfigConst = "RABBITMQ_HOST"
	RABBITMQ_PORT                      ConfigConst = "RABBITMQ_PORT"
	RABBITMQ_BROADCAST_NAMEConfigConst ConfigConst = "RABBITMQ_BROADCAST_NAME"
	RABBITMQ_CONNECTIONSTRING          ConfigConst = "RABBITMQ_CONNECTIONSTRING"

	REDIS_HOST = "REDIS_HOST"
	REDIS_PORT = "REDIS_PORT"

	TRACER_ID                        ConfigConst = "TRACER_ID"
	ENABLE_TRACING                   ConfigConst = "ENABLE_TRACING"
	OPENTELEMETRY_COLLECTOR_ENDPOINT ConfigConst = "OPENTELEMETRY_COLLECTOR_ENDPOINT"

	PROFILER_ENABLEDConfigConst             = "PROFILER_ENABLED"
	STARTUP_SLEEP_TIME          ConfigConst = "STARTUP_SLEEP_TIME"

	GIN_USE_CORS      ConfigConst = "USE_CORS"
	GIN_ALLOW_ORIGINS ConfigConst = "ALLOW_ORIGINS"

	VERSION           ConfigConst = "VERSION"
	COMMIT            ConfigConst = "COMMIT"
	DEX_HOST          ConfigConst = "DEX_HOST"
	DEX_PORT          ConfigConst = "DEX_PORT"
	DEX_GRPC_PORT     ConfigConst = "DEX_GRPC_PORT"
	DEX_CERT_FILEPATH ConfigConst = "DEX_CERT_FILEPATH"
	DEX_VAULT_PATH    ConfigConst = "DEX_VAULT_PATH"
	DEX_TLS           ConfigConst = "DEX_TLS"

	SLACK_BOT_TOKEN = "SLACK_BOT_TOKEN" // #nosec G101 Jest the name of the variable holding the value
	SLACK_APP_TOKEN = "SLACK_APP_TOKEN" // #nosec G101 Jest the name of the variable holding the value

	CONTAINER_REG_PREFIX     ConfigConst = "CONTAINER_REG_PREFIX"
	CONTAINER_REG_IMAGE_PATH ConfigConst = "CONTAINER_REG_IMAGE_PATH"
	CONTAINER_REG_HELM_PATH  ConfigConst = "CONTAINER_REG_HELM_PATH"

	// Operator
	OPERATOR_BACKOFF_LIMIT       ConfigConst = "BACKOFF_LIMIT"
	OPERATOR_DEADLINE_SECONDS    ConfigConst = "DEADLINE_SECONDS"
	OPERATOR_JOB_SERVICE_ACCOUNT ConfigConst = "JOB_SERVICE_ACCOUNT"
	OPERATOR_APPLOG_SECRET_NAME  ConfigConst = "APPLOG_SECRET_NAME" // #nosec G101 Jest the name of the variable holding the value

	// Tanzu agent
	TANZU_AGENT_KUBECONFIG         ConfigConst = "KUBECONFIG"
	TANZU_AGENT_DELETE_KUBECONFIG  ConfigConst = "DELETE_KUBECONFIG"
	TANZU_AGENT_KUBE_VSPHERE_PATH  ConfigConst = "KUBE_VSPHERE_PATH"
	TANZU_AGENT_KUBECTL_PATH       ConfigConst = "KUBECTL_PATH"
	TANZU_AGENT_DATACENTER_URL     ConfigConst = "DATACENTER_URL"
	TANZU_AGENT_USERNAME           ConfigConst = "TANZU_USERNAME"
	TANZU_AGENT_PASSWORD           ConfigConst = "TANZU_PWD"    // #nosec G101 Jest the name of the variable holding the value
	TANZU_AGENT_TOKEN_EXPIRY       ConfigConst = "TOKEN_EXPIRY" // #nosec G101 Jest the name of the variable holding the value
	TANZU_AGENT_LOGIN_EVERY_MINUTE ConfigConst = "LOGIN_EVERY_MINUTE"
	TANZU_AGENT_HEALTH_PORT        ConfigConst = "HTTP_PORT"
	TANZU_AGENT_TANZU_ACCESS       ConfigConst = "TANZU_ACCESS"
	TANZU_AGENT_DATACENTER         ConfigConst = "DATACENTER"
	TANZU_AGENT_PROVIDER           ConfigConst = "PROVIDER"
	TANZU_AGENT_REGION             ConfigConst = "REGION"

	// Tanzu Auth
	TANZU_AUTH_HEALTH_PORT        ConfigConst = "HTTP_PORT"
	TANZU_AUTH_KUBE_VSPHERE_PATH  ConfigConst = "KUBE_VSPHERE_PATH"
	TANZU_AUTH_KUBECTL_PATH       ConfigConst = "KUBECTL_PATH"
	TANZU_AUTH_BASE_URL           ConfigConst = "TANZU_AUTH_BASE_URL"
	TANZU_AUTH_CONFIG_FOLDER_PATH ConfigConst = "TANZU_AUTH_CONFIG_FOLDER_PATH"

	// ms
	MS_HTTP_PORT      ConfigConst = "HTTP_PORT"
	MS_HTTP_BIND_PORT ConfigConst = "HTTP_BIND_PORT"

	LOCAL_KUBERNETES_ROR_BASE_URL ConfigConst = "LOCAL_KUBERNETES_ROR_BASE_URL"
	ENABLE_PPROF                  ConfigConst = "ENABLE_PPROF"
)
