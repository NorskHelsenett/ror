package kubernetesclient

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	metricsv1beta1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// K8sClientsets represents a collection of Kubernetes clientsets.
type K8sClientsets struct {
	k8sConfig       *rest.Config
	k8sClientset    *kubernetes.Clientset
	discoveryClient *discovery.DiscoveryClient
	dynamicClient   dynamic.Interface
	metricsClient   *metrics.Clientset
	metricsV1Beta1  metricsv1beta1.MetricsV1beta1Interface
}

type K8SClientInterface interface {
	GetConfig() *rest.Config
	GetDiscoveryClient() (*discovery.DiscoveryClient, error)
	GetDynamicClient() (dynamic.Interface, error)
	GetMetricsClient() (*metrics.Clientset, error)
	GetSecret(namespace string, name string) (*v1.Secret, error)
	GetNodes() ([]v1.Node, error)
}

// NewK8sClientConfig initializes a new instance of the K8sClientsets struct with the Kubernetes configuration.
//
// Returns:
// - (*K8sClientsets): The initialized K8sClientsets object.
func NewK8sClientConfig() *K8sClientsets {
	k8sconfig, err := config.GetConfig()
	if err != nil {
		rlog.Error("Error trying to get k8s config", err)
		panic(err)
	}
	return &K8sClientsets{k8sConfig: k8sconfig}
}

// GetConfig retrieves the Kubernetes configuration.
//
// Returns:
// - (*rest.Config): The Kubernetes configuration.
func (c *K8sClientsets) GetConfig() *rest.Config {
	return c.k8sConfig
}

// GetKubernetesClientset returns a Kubernetes clientset for interacting with the Kubernetes API.
//
// Parameters:
// - kubeconfigPath (string): The path to the Kubernetes configuration file.
//
// Returns:
// - (*kubernetes.Clientset): The Kubernetes clientset.
// - (error): An error if the clientset creation fails.
func (c *K8sClientsets) GetKubernetesClientset() (*kubernetes.Clientset, error) {
	if c.k8sClientset == nil {
		var err error
		c.k8sClientset, err = kubernetes.NewForConfig(c.k8sConfig)
		if err != nil {
			rlog.Error("Error trying to create clientset for k8s", err)
			return c.k8sClientset, err
		}
	}
	return c.k8sClientset, nil
}

// GetDiscoveryClient returns the Kubernetes discovery client for retrieving API server information.
// If the discovery client has not been initialized, it creates a new client using the provided
// Kubernetes configuration.
//
// Returns:
// - (discovery.DiscoveryInterface): The discovery client.
// - (error): An error if the discovery client creation fails.
func (c *K8sClientsets) GetDiscoveryClient() (*discovery.DiscoveryClient, error) {
	if c.discoveryClient == nil {
		var err error
		c.discoveryClient, err = discovery.NewDiscoveryClientForConfig(c.k8sConfig)
		if err != nil {
			rlog.Error("Error trying to create discovery client for k8s", err)
			return c.discoveryClient, err
		}
	}
	return c.discoveryClient, nil
}

// GetDynamicClient returns the dynamic client for interacting with Kubernetes resources.
// If the dynamic client has not been initialized, it creates a new client using the
// provided Kubernetes configuration.
//
// Returns:
// - (dynamic.Interface): The dynamic client.
// - (error): An error if the dynamic client creation fails.
func (c *K8sClientsets) GetDynamicClient() (dynamic.Interface, error) {
	if c.dynamicClient == nil {
		var err error
		c.dynamicClient, err = dynamic.NewForConfig(c.k8sConfig)
		if err != nil {
			rlog.Error("Error trying to create dynamic client for k8s", err)
			return c.dynamicClient, err
		}
	}
	return c.dynamicClient, nil
}

// GetMetricsClient returns the metrics client for interacting with Kubernetes metrics.
// If the metrics client has not been initialized, it creates a new client using the
// provided Kubernetes configuration.
//
// Returns:
// - (*metrics.Clientset): The metrics clientset.
// - (error): An error if the metrics client creation fails.
func (c *K8sClientsets) GetMetricsClient() (*metrics.Clientset, error) {
	if c.metricsClient == nil {
		var err error
		c.metricsClient, err = metrics.NewForConfig(c.k8sConfig)
		if err != nil {
			rlog.Error("failed to get metrics client", err)
			return c.metricsClient, err
		}
	}
	return c.metricsClient, nil
}

// GetMetricsV1Beta1Client returns the metrics v1beta1 client for interacting with Kubernetes metrics.
// If the metrics v1beta1 client has not been initialized, it creates a new client using the
// provided Kubernetes configuration.
//
// Returns:
// - (metricsv1beta1.MetricsV1beta1Interface): The metrics v1beta1 client.
// - (error): An error if the metrics v1beta1 client creation fails.
func (c *K8sClientsets) GetMetricsV1Beta1Client() (metricsv1beta1.MetricsV1beta1Interface, error) {
	if c.metricsV1Beta1 == nil {
		var err error
		c.metricsV1Beta1, err = metricsv1beta1.NewForConfig(c.k8sConfig)
		if err != nil {
			rlog.Error("failed to get metrics v1beta1 client", err)
			return c.metricsV1Beta1, err
		}
	}
	return c.metricsV1Beta1, nil
}

// GetSecret retrieves a Kubernetes secret by its name and namespace using the provided clientsets.
//
// Parameters:
// - c (*K8sClientsets): The K8sClientsets object that contains the necessary clientsets for interacting with secrets.
// - name (string): The name of the secret to retrieve.
// - namespace (string): The namespace of the secret.
//
// Returns:
// - (*corev1.Secret): The retrieved secret.
// - (error): An error if the secret retrieval fails.
func (c *K8sClientsets) GetSecret(namespace string, name string) (*v1.Secret, error) {
	client, err := c.GetKubernetesClientset()
	if err != nil {
		return nil, err
	}

	secret, err := client.CoreV1().Secrets(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return secret, nil
}

// CreateSecret creates a Kubernetes secret in the specified namespace using the provided clientsets.
//
// Parameters:
// - c (*K8sClientsets): The K8sClientsets object that contains the necessary clientsets for interacting with secrets.
// - namespace (string): The namespace in which to create the secret.
// - secret (*corev1.Secret): The secret to create.
//
// Returns:
// - (*corev1.Secret): The created secret.
// - (error): An error if the secret creation fails.
func (c *K8sClientsets) CreateSecret(namespace string, secret *v1.Secret) (*v1.Secret, error) {
	client, err := c.GetKubernetesClientset()
	if err != nil {
		return nil, err
	}

	ret, err := client.CoreV1().Secrets(namespace).Create(context.Background(), secret, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetNodes retrieves the list of Kubernetes nodes using the provided clientsets.
//
// Parameters:
// - c (*K8sClientsets): The K8sClientsets object that contains the necessary clientsets for interacting with nodes.
//
// Returns:
// - ([]corev1.Node, error): The list of retrieved nodes and an error if the retrieval fails.
func (c *K8sClientsets) GetNodes() ([]v1.Node, error) {
	client, err := c.GetKubernetesClientset()
	if err != nil {
		return nil, err
	}

	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return nodes.Items, nil
}

// GetNamespace retrieves a Kubernetes namespace by its name using the provided clientsets.
func (c *K8sClientsets) GetNamespace(namespace string) (*v1.Namespace, error) {
	client, err := c.GetKubernetesClientset()
	if err != nil {
		return nil, err
	}

	ns, err := client.CoreV1().Namespaces().Get(context.Background(), namespace, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return ns, nil
}

// MustInitializeKubernetesClient initializes a Kubernetes client and panics if the initialization fails.
//
// Returns:
// - (*K8sClientsets): The initialized K8sClientsets object.
func MustInitializeKubernetesClient() *K8sClientsets {
	kubernetesCli := NewK8sClientConfig()
	if kubernetesCli == nil {
		panic("failed to initialize kubernetes client")
	}
	return kubernetesCli
}
