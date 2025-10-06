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

	RABBITMQ_HOST             ConfigConst = "RABBITMQ_HOST"
	RABBITMQ_PORT             ConfigConst = "RABBITMQ_PORT"
	RABBITMQ_BROADCAST_NAME   ConfigConst = "RABBITMQ_BROADCAST_NAME"
	RABBITMQ_CONNECTIONSTRING ConfigConst = "RABBITMQ_CONNECTIONSTRING"

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

type ConfigConstData struct {
	value       string
	deprecated  bool
	description string
}

// ...existing code...
var ConfigConstsMap = map[ConfigConst]ConfigConstData{
	ConfigConst("ROLE"):                               {value: "ROLE", deprecated: false, description: ""},
	ConfigConst("HTTP_HOST"):                          {value: "HTTP_HOST", deprecated: false, description: ""},
	ConfigConst("HTTP_PORT"):                          {value: "HTTP_PORT", deprecated: false, description: ""},
	ConfigConst("HTTP_HEALTH_HOST"):                   {value: "HTTP_HEALTH_HOST", deprecated: false, description: ""},
	ConfigConst("HTTP_HEALTH_PORT"):                   {value: "HTTP_HEALTH_PORT", deprecated: false, description: ""},
	ConfigConst("HTTP_TIMEOUT"):                       {value: "HTTP_TIMEOUT", deprecated: false, description: ""},
	ConfigConst("HEALTH_ENDPOINT"):                    {value: "HEALTH_ENDPOINT", deprecated: true, description: "use HTTP_HEALTH_HOST / HTTP_HEALTH_PORT instead"},
	ConfigConst("API_KEY_SALT"):                       {value: "API_KEY_SALT", deprecated: false, description: ""},
	ConfigConst("API_KEY"):                            {value: "API_KEY", deprecated: false, description: ""},
	ConfigConst("API_KEY_SECRET"):                     {value: "API_KEY_SECRET", deprecated: false, description: ""},
	ConfigConst("API_ENDPOINT"):                       {value: "ROR_URL", deprecated: false, description: ""},
	ConfigConst("DEVELOPMENT"):                        {value: "DEVELOPMENT", deprecated: false, description: ""},
	ConfigConst("PORT"):                               {value: "PORT", deprecated: false, description: ""},
	ConfigConst("POD_NAMESPACE"):                      {value: "POD_NAMESPACE", deprecated: false, description: ""},
	ConfigConst("CLUSTER_ID"):                         {value: "CLUSTER_ID", deprecated: false, description: ""},
	ConfigConst("ERROR_COUNT"):                        {value: "ERROR_COUNT", deprecated: false, description: ""},
	ConfigConst("HELSEGITLAB_BASE_URL"):               {value: "HELSEGITLAB_BASE_URL", deprecated: false, description: ""},
	ConfigConst("ENVIRONMENT"):                        {value: "ENVIRONMENT", deprecated: false, description: ""},
	ConfigConst("LDAP_CONFIGS"):                       {value: "LDAP_CONFIGS", deprecated: false, description: ""},
	ConfigConst("GIT_REPO_URL"):                       {value: "GIT_REPO_URL", deprecated: false, description: ""},
	ConfigConst("GIT_BRANCHConfigConst"):              {value: "GIT_BRANCH", deprecated: false, description: ""},
	ConfigConst("GIT_TOKEN"):                          {value: "GIT_TOKEN", deprecated: false, description: ""},
	ConfigConst("GIT_PATH"):                           {value: "GIT_PATH", deprecated: false, description: ""},
	ConfigConst("ROR_OPERATOR_NAMESPACE"):             {value: "ROR_OPERATOR_NAMESPACE", deprecated: false, description: ""},
	ConfigConst("LDAP_CERTIFICATE_FOLDER"):            {value: "LDAP_CERTIFICATE_FOLDER", deprecated: false, description: ""},
	ConfigConst("OIDC_PROVIDER"):                      {value: "OIDC_PROVIDER", deprecated: false, description: ""},
	ConfigConst("OIDC_CLIENT_ID"):                     {value: "OIDC_CLIENT_ID", deprecated: false, description: ""},
	ConfigConst("OIDC_DEVICE_CLIENT_IDConfigConst"):   {value: "OIDC_DEVICE_CLIENT_ID", deprecated: false, description: ""},
	ConfigConst("OIDC_SKIP_ISSUER_VERIFY"):            {value: "SKIP_OIDC_ISSUER_VERIFY", deprecated: false, description: ""},
	ConfigConst("VAULT_URL"):                          {value: "VAULT_URL", deprecated: false, description: ""},
	ConfigConst("MONGODB_HOST"):                       {value: "MONGODB_HOST", deprecated: false, description: ""},
	ConfigConst("MONGODB_PORT"):                       {value: "MONGODB_PORT", deprecated: false, description: ""},
	ConfigConst("MONGODB_DATABASE"):                   {value: "MONGODB_DATABASE", deprecated: false, description: ""},
	ConfigConst("MONGODB_URL"):                        {value: "MONGODB_URL", deprecated: false, description: ""},
	ConfigConst("MONGO_DATABASEConfigConst"):          {value: "MONGO_DATABASE", deprecated: false, description: ""},
	ConfigConst("RABBITMQ_HOST"):                      {value: "RABBITMQ_HOST", deprecated: false, description: ""},
	ConfigConst("RABBITMQ_PORT"):                      {value: "RABBITMQ_PORT", deprecated: false, description: ""},
	ConfigConst("RABBITMQ_BROADCAST_NAMEConfigConst"): {value: "RABBITMQ_BROADCAST_NAME", deprecated: false, description: ""},
	ConfigConst("RABBITMQ_CONNECTIONSTRING"):          {value: "RABBITMQ_CONNECTIONSTRING", deprecated: false, description: ""},
	ConfigConst("REDIS_HOST"):                         {value: "REDIS_HOST", deprecated: false, description: ""},
	ConfigConst("REDIS_PORT"):                         {value: "REDIS_PORT", deprecated: false, description: ""},
	ConfigConst("TRACER_ID"):                          {value: "TRACER_ID", deprecated: false, description: ""},
	ConfigConst("ENABLE_TRACING"):                     {value: "ENABLE_TRACING", deprecated: false, description: ""},
	ConfigConst("OPENTELEMETRY_COLLECTOR_ENDPOINT"):   {value: "OPENTELEMETRY_COLLECTOR_ENDPOINT", deprecated: false, description: ""},
	ConfigConst("PROFILER_ENABLED"):                   {value: "PROFILER_ENABLED", deprecated: false, description: ""},
	ConfigConst("STARTUP_SLEEP_TIME"):                 {value: "STARTUP_SLEEP_TIME", deprecated: false, description: ""},
	ConfigConst("GIN_USE_CORS"):                       {value: "USE_CORS", deprecated: false, description: ""},
	ConfigConst("GIN_ALLOW_ORIGINS"):                  {value: "ALLOW_ORIGINS", deprecated: false, description: ""},
	ConfigConst("VERSION"):                            {value: "VERSION", deprecated: false, description: ""},
	ConfigConst("COMMIT"):                             {value: "COMMIT", deprecated: false, description: ""},
	ConfigConst("DEX_HOST"):                           {value: "DEX_HOST", deprecated: false, description: ""},
	ConfigConst("DEX_PORT"):                           {value: "DEX_PORT", deprecated: false, description: ""},
	ConfigConst("DEX_GRPC_PORT"):                      {value: "DEX_GRPC_PORT", deprecated: false, description: ""},
	ConfigConst("DEX_CERT_FILEPATH"):                  {value: "DEX_CERT_FILEPATH", deprecated: false, description: ""},
	ConfigConst("DEX_VAULT_PATH"):                     {value: "DEX_VAULT_PATH", deprecated: false, description: ""},
	ConfigConst("DEX_TLS"):                            {value: "DEX_TLS", deprecated: false, description: ""},
	ConfigConst("SLACK_BOT_TOKEN"):                    {value: "SLACK_BOT_TOKEN", deprecated: false, description: ""},
	ConfigConst("SLACK_APP_TOKEN"):                    {value: "SLACK_APP_TOKEN", deprecated: false, description: ""},
	ConfigConst("CONTAINER_REG_PREFIX"):               {value: "CONTAINER_REG_PREFIX", deprecated: false, description: ""},
	ConfigConst("CONTAINER_REG_IMAGE_PATH"):           {value: "CONTAINER_REG_IMAGE_PATH", deprecated: false, description: ""},
	ConfigConst("CONTAINER_REG_HELM_PATH"):            {value: "CONTAINER_REG_HELM_PATH", deprecated: false, description: ""},
	ConfigConst("OPERATOR_BACKOFF_LIMIT"):             {value: "BACKOFF_LIMIT", deprecated: false, description: ""},
	ConfigConst("OPERATOR_DEADLINE_SECONDS"):          {value: "DEADLINE_SECONDS", deprecated: false, description: ""},
	ConfigConst("OPERATOR_JOB_SERVICE_ACCOUNT"):       {value: "JOB_SERVICE_ACCOUNT", deprecated: false, description: ""},
	ConfigConst("OPERATOR_APPLOG_SECRET_NAME"):        {value: "APPLOG_SECRET_NAME", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_KUBECONFIG"):             {value: "KUBECONFIG", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_DELETE_KUBECONFIG"):      {value: "DELETE_KUBECONFIG", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_KUBE_VSPHERE_PATH"):      {value: "KUBE_VSPHERE_PATH", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_KUBECTL_PATH"):           {value: "KUBECTL_PATH", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_DATACENTER_URL"):         {value: "DATACENTER_URL", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_USERNAME"):               {value: "TANZU_USERNAME", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_PASSWORD"):               {value: "TANZU_PWD", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_TOKEN_EXPIRY"):           {value: "TOKEN_EXPIRY", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_LOGIN_EVERY_MINUTE"):     {value: "LOGIN_EVERY_MINUTE", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_HEALTH_PORT"):            {value: "HTTP_PORT", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_TANZU_ACCESS"):           {value: "TANZU_ACCESS", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_DATACENTER"):             {value: "DATACENTER", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_PROVIDER"):               {value: "PROVIDER", deprecated: false, description: ""},
	ConfigConst("TANZU_AGENT_REGION"):                 {value: "REGION", deprecated: false, description: ""},
	ConfigConst("TANZU_AUTH_HEALTH_PORT"):             {value: "HTTP_PORT", deprecated: false, description: ""},
	ConfigConst("TANZU_AUTH_KUBE_VSPHERE_PATH"):       {value: "KUBE_VSPHERE_PATH", deprecated: false, description: ""},
	ConfigConst("TANZU_AUTH_KUBECTL_PATH"):            {value: "KUBECTL_PATH", deprecated: false, description: ""},
	ConfigConst("TANZU_AUTH_BASE_URL"):                {value: "TANZU_AUTH_BASE_URL", deprecated: false, description: ""},
	ConfigConst("TANZU_AUTH_CONFIG_FOLDER_PATH"):      {value: "TANZU_AUTH_CONFIG_FOLDER_PATH", deprecated: false, description: ""},
	ConfigConst("MS_HTTP_PORT"):                       {value: "HTTP_PORT", deprecated: false, description: ""},
	ConfigConst("MS_HTTP_BIND_PORT"):                  {value: "HTTP_BIND_PORT", deprecated: false, description: ""},
	ConfigConst("LOCAL_KUBERNETES_ROR_BASE_URL"):      {value: "LOCAL_KUBERNETES_ROR_BASE_URL", deprecated: false, description: ""},
	ConfigConst("ENABLE_PPROF"):                       {value: "ENABLE_PPROF", deprecated: false, description: ""},
}
