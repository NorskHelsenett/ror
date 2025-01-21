// RorConfig represents a ror configuration and clients for ror and Kubernetes.
package rorclientconfig

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
	"github.com/NorskHelsenett/ror/pkg/clients/kubernetes/clusterinterregator"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"

	"github.com/NorskHelsenett/ror/pkg/config/rorversion"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

const ERR_SECRET_NOT_FOUND = "___secret_not_found___"

var RorConfig *RorClientConfig

var Kubernetes *kubernetesclient.K8sClientsets

func MustInitializeKubernetesClient() {
	Kubernetes = kubernetesclient.NewK8sClientConfig()
	if Kubernetes == nil {
		panic("failed to initialize kubernetes client")
	}
}

type RorClientConfig struct {
	config    ClientConfig
	apiKey    string
	role      string
	clusterId string
	rorClient *rorclient.RorClient
}

type ClientConfig struct {
	Role                     string
	Namespace                string
	ApiEndpoint              string
	ApiKey                   string
	ApiKeySecret             string
	RorVersion               rorversion.RorVersion
	MustInitializeKubernetes bool
}

func InitRorClientConfig(conf ClientConfig) {

	RorConfig = NewRorConfig(conf)
	if conf.MustInitializeKubernetes {
		MustInitializeKubernetesClient()
	}
	err := RorConfig.getConfig()
	if err != nil {
		rlog.Fatal("failed to get config", err)
	}
	err = RorConfig.InitRorClient()
	if err != nil {
		rlog.Fatal("failed to init ror client", err)
	}
}

func NewRorConfig(conf ClientConfig) *RorClientConfig {
	return &RorClientConfig{
		config: conf,
		role:   conf.Role,
	}
}

func (a *RorClientConfig) getConfig() error {
	// must get auth config
	rlog.Debug("Authenticating ror config provider")
	if Kubernetes != nil {
		rlog.Debug("Using kubernetes auth provider")
		a.kubernetesAuth()
	} else {
		rlog.Debug("Using env auth provider")
		a.envAuth()
	}
	if a.GetApiKey() == "" {
		return fmt.Errorf("failed to get api key")
	}

	// check if namespace is set and accessible
	if a.config.Namespace == "" {
		return fmt.Errorf("failed to get namespace")
	}
	_, err := Kubernetes.GetNamespace(a.config.Namespace)
	if err != nil {
		return fmt.Errorf("failed to get namespace %s", err)
	}

	// check if api endpoint is set and accessible
	if a.config.ApiEndpoint == "" {
		return fmt.Errorf("failed to get api endpoint")
	}

	if a.GetApiKey() == ERR_SECRET_NOT_FOUND {
		rlog.Info("api key secret not found, interregating cluster and registering new key")

		interregationreport, err := clusterinterregator.InterregateCluster(Kubernetes)

		if err != nil {
			return fmt.Errorf("failed to interregate cluster %s", err)
		}

		a.initUnathorizedRorClient()
		key, err := a.rorClient.Clusters().Register(apicontracts.AgentApiKeyModel{
			Identifier:     interregationreport.ClusterName,
			DatacenterName: interregationreport.Datacenter,
			WorkspaceName:  interregationreport.Workspace,
			Provider:       interregationreport.Provider,
			Type:           "Cluster",
		})
		if err != nil {
			return fmt.Errorf("failed to register cluster %s", err)
		}
		err = a.kubernetesCreateApiKeySecret(key)
		if err != nil {
			return fmt.Errorf("failed to create api key secret %s", err)

		}
	}

	return nil

}

// kubernetesAuth sets the api key from the secret defined in the config value ApiKeySecret
// in the namespace defined in the config value Namespace
// if the secret does not exist, it will be created
func (a *RorClientConfig) kubernetesAuth() {
	secret, err := Kubernetes.GetSecret(a.config.Namespace, a.config.ApiKeySecret)
	if err != nil {
		if errors.IsNotFound(err) {
			rlog.Warn("api key secret not found")
			a.SetApiKey(ERR_SECRET_NOT_FOUND)
			return
		} else {
			rlog.Error("failed to get api key secret", err)
			return
		}
	}
	a.SetApiKey(string(secret.Data["APIKEY"]))
}

func (a *RorClientConfig) kubernetesCreateApiKeySecret(apiKey string) error {
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      a.config.ApiKeySecret,
			Namespace: a.config.Namespace,
		},
		Type: v1.SecretTypeOpaque,
		StringData: map[string]string{
			"APIKEY": apiKey,
		},
	}
	_, err := Kubernetes.CreateSecret(a.config.Namespace, secret)
	if err != nil {
		rlog.Error("failed to create api key secret", err)
		return err
	}
	a.SetApiKey(apiKey)
	return nil

}

// envAuth sets the api key from the config value ApiKey
func (a *RorClientConfig) envAuth() {
	a.SetApiKey(a.config.ApiKey)
}
func (a *RorClientConfig) SetApiKey(apiKey string) {
	a.apiKey = apiKey
}

func (a *RorClientConfig) GetApiKey() string {
	return a.apiKey
}

func (a *RorClientConfig) SetRole(role string) {
	a.role = role
}

func (a *RorClientConfig) GetRole() string {
	return a.role
}

func (a *RorClientConfig) SetClusterId(clusterId string) {
	a.clusterId = clusterId
}

func (a *RorClientConfig) GetClusterId() string {
	return a.clusterId
}

func (a *RorClientConfig) GetKubernetesClientSet() *kubernetesclient.K8sClientsets {
	return Kubernetes
}

func (a *RorClientConfig) InitRorClient() error {
	httptransportconfig := httpclient.HttpTransportClientConfig{
		BaseURL:      a.config.ApiEndpoint,
		AuthProvider: httpauthprovider.NewAuthProvider(httpauthprovider.AuthPoviderTypeAPIKey, a.apiKey),
		Role:         a.role,
		Version:      a.config.RorVersion,
	}
	rorclienttransport := resttransport.NewRorHttpTransport(&httptransportconfig)

	a.rorClient = rorclient.NewRorClient(rorclienttransport)
	ver, err := a.rorClient.Info().GetVersion()
	if err != nil {
		return err
	}
	rlog.Info("connected to ror-api", rlog.String("version", ver))

	selfdata, err := a.rorClient.Clusters().GetSelf()
	if err != nil {
		return err
	}

	rlog.Info("connected to ror-api", rlog.String("clusterid", selfdata.ClusterId))
	a.SetClusterId(selfdata.ClusterId)
	return nil
}

func (a *RorClientConfig) initUnathorizedRorClient() {
	httptransportconfig := httpclient.HttpTransportClientConfig{
		BaseURL:      a.config.ApiEndpoint,
		AuthProvider: httpauthprovider.NewNoAuthprovider(),
		Role:         a.role,
		Version:      a.config.RorVersion,
	}
	rorclienttransport := resttransport.NewRorHttpTransport(&httptransportconfig)
	a.rorClient = rorclient.NewRorClient(rorclienttransport)
}

func (a *RorClientConfig) GetRorClient() *rorclient.RorClient {
	return a.rorClient
}

func (a *RorClientConfig) CreateOwnerref() rorresourceowner.RorResourceOwnerReference {
	return rorresourceowner.RorResourceOwnerReference{
		Scope:   aclmodels.Acl2ScopeCluster,
		Subject: aclmodels.Acl2Subject(a.GetClusterId()),
	}
}
