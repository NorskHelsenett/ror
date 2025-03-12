package kubeconfig

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"maps"

	"github.com/golang-jwt/jwt"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// KubeConfig is the kubeconfig file
type KubeConfig struct {
	Config *clientcmdapi.Config
	Path   string
	Errors []error
}

// LoadKubeConfig loads the kubeconfig file
func MustLoadFromDefaultFile() *KubeConfig {
	return MustLoadFromFile(getDefaultFilename())
}

// LoadFromFile loads the kubeconfig file
func MustLoadFromFile(path string) *KubeConfig {
	config, err := LoadFromFile(path)
	if err != nil {
		panic(err)
	}
	return config
}

// NewKubeConfig loads the kubeconfig file
func LoadFromDefaultFile() (*KubeConfig, error) {
	return LoadFromFile(getDefaultFilename())
}

// LoadFromFile loads the kubeconfig file
func LoadFromFile(path string) (*KubeConfig, error) {
	config, err := clientcmd.LoadFromFile(path)
	if err != nil {
		return nil, err
	}
	return &KubeConfig{
		Config: config,
		Path:   path,
	}, nil
}

// LoadOrNewKubeConfig loads the kubeconfig file or creates a new one
func LoadOrNewKubeConfig() *KubeConfig {
	config, err := LoadFromDefaultFile()
	if err != nil {
		return NewKubeConfig()
	}
	return config
}

// NewKubeConfig loads the kubeconfig file
func NewKubeConfig() *KubeConfig {
	config := clientcmdapi.NewConfig()
	path := getDefaultFilename()
	return &KubeConfig{
		Config: config,
		Path:   path,
	}
}

// NewKubeConfigFromYaml loads the kubeconfig file from yaml
func NewKubeConfigFromYaml(yaml []byte) *KubeConfig {
	config, err := clientcmd.Load(yaml)
	if err != nil {
		return &KubeConfig{
			Errors: []error{err},
		}
	}
	return &KubeConfig{
		Config: config,
	}
}

// IsExpired check if the selected contexts token is expired
func (k *KubeConfig) IsExpired(context string) (bool, error) {

	expires, err := k.ExpiresByContext(context)
	if err != nil {
		return true, err
	}
	return time.Now().After(expires), nil
}

// Expires returns the expiry time of the selected contexts token
func (k *KubeConfig) ExpiresByContext(context string) (time.Time, error) {

	isExpired := time.Unix(0, 0)
	if k == nil {
		fmt.Println("kubeconfig is nil")
		return isExpired, nil
	}
	if errs := k.HandleErrors(); errs != nil {
		fmt.Println("error checking kubeconfig")
		return isExpired, errs
	}

	if k.Config.Contexts[context] == nil || k.Config.Contexts[context].AuthInfo == "" {
		return isExpired, nil
	}
	authInfo, exists := k.Config.AuthInfos[k.Config.Contexts[context].AuthInfo]
	if !exists {
		fmt.Println("authinfo not found")
		return isExpired, nil
	}
	fmt.Println("authinfo not found", context, getExpiry(authInfo).String())
	return getExpiry(authInfo), nil
}

func (k *KubeConfig) Expires(AuthInforName string) (time.Time, error) {

	isExpired := time.Unix(0, 0)
	if k == nil {
		return isExpired, nil
	}
	if errs := k.HandleErrors(); errs != nil {
		return isExpired, errs
	}

	if _, ok := k.Config.AuthInfos[AuthInforName]; !ok {
		return isExpired, nil
	}
	authInfo, exists := k.Config.AuthInfos[AuthInforName]
	if !exists {
		return isExpired, nil
	}
	return getExpiry(authInfo), nil
}

func getExpiry(authInfo *clientcmdapi.AuthInfo) time.Time {
	if authInfo.Token == "" {
		// we can't check if the token is expired if it's not set
		return time.Unix(0, 0)
	}
	token, _, err := new(jwt.Parser).ParseUnverified(authInfo.Token, jwt.MapClaims{})

	if err != nil {
		return time.Unix(0, 0)
	}

	claims := token.Claims.(jwt.MapClaims)

	expiry := int64(claims["exp"].(float64))
	return time.Unix(expiry, 0)
}

// LoadFromBytes loads the kubeconfig file
func (k *KubeConfig) MergeYaml(yaml []byte) *KubeConfig {

	mergeConfig := NewKubeConfigFromYaml(yaml).Config

	if k.Config != nil {
		maps.Copy(k.Config.Clusters, mergeConfig.Clusters)
		maps.Copy(k.Config.AuthInfos, mergeConfig.AuthInfos)
		for key, v := range mergeConfig.Contexts {
			old, exists := k.Config.Contexts[key]
			k.Config.Contexts[key] = v
			if exists && old.Namespace != "" {
				k.Config.Contexts[key].Namespace = old.Namespace
			}
		}
	}

	if k.Config == nil {
		k.Config = mergeConfig
	}

	return k
}

// Merge merges the kubeconfig file
func (k *KubeConfig) Merge(config *KubeConfig) *KubeConfig {
	if k.Config == nil {
		k.Config = config.Config
		return k
	}
	maps.Copy(k.Config.Clusters, config.Config.Clusters)
	maps.Copy(k.Config.AuthInfos, config.Config.AuthInfos)
	for key, v := range config.Config.Contexts {
		old, exists := k.Config.Contexts[key]
		k.Config.Contexts[key] = v
		if exists && old.Namespace != "" {
			k.Config.Contexts[key].Namespace = old.Namespace
		}
	}
	return k
}

// AddCluster adds a cluster to the kubeconfig file
func (k *KubeConfig) AddCluster(name string, cluster *clientcmdapi.Cluster) *KubeConfig {
	if k.Config == nil {
		k.Config = clientcmdapi.NewConfig()
	}
	k.Config.Clusters[name] = cluster
	return k
}

// AddAuthInfo adds an auth info to the kubeconfig file
func (k *KubeConfig) AddAuthInfo(name string, authInfo *clientcmdapi.AuthInfo) *KubeConfig {
	if k.Config == nil {
		k.Config = clientcmdapi.NewConfig()
	}
	k.Config.AuthInfos[name] = authInfo
	return k
}

// AddContext adds a context to the kubeconfig file
func (k *KubeConfig) AddContext(name string, context *clientcmdapi.Context) *KubeConfig {
	if k.Config == nil {
		k.Config = clientcmdapi.NewConfig()
	}
	k.Config.Contexts[name] = context
	return k
}

// Write writes the kubeconfig file
func (k *KubeConfig) Write() error {
	if errs := k.HandleErrors(); errs != nil {
		return errs
	}

	if k.Config == nil {
		return errors.New("kubeconfig is nil")
	}
	path := filepath.Dir(k.Path)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		homeinfo, err := os.Stat(homedir)
		if err != nil {
			return err
		}
		err = os.MkdirAll(path, homeinfo.Mode())
		if err != nil {
			return err
		}
	}

	return clientcmd.WriteToFile(*k.Config, k.Path)
}

// Handle errors return a string of errors
func (k *KubeConfig) HandleErrors() error {
	if len(k.Errors) == 0 {
		return nil
	}

	var errs string
	for _, err := range k.Errors {
		errs += err.Error() + "\n"
	}
	return fmt.Errorf("errors: %s", errs)
}

// SetCurrentContext sets the current context
func (k *KubeConfig) SetContext(context string) *KubeConfig {
	k.Config.CurrentContext = context
	return k
}

// GetCurrentContext gets the current context
func (k *KubeConfig) GetCurrentContext() (string, error) {
	if errs := k.HandleErrors(); errs != nil {
		return "", errs
	}
	return k.Config.CurrentContext, nil
}

// SetNamespace sets the namespace for the current context
func (k *KubeConfig) SetNamespace(namespace string) *KubeConfig {
	if errs := k.HandleErrors(); errs != nil {
		return k
	}

	context := k.Config.Contexts[k.Config.CurrentContext]
	context.Namespace = namespace
	k.Config.Contexts[k.Config.CurrentContext] = context
	err := k.Write()
	if err != nil {
		k.Errors = append(k.Errors, err)
	}
	return k
}

// getDefaultFilename returns the default filename
func getDefaultFilename() string {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	return loadingRules.GetDefaultFilename()
}

// GetFirstCluster returns the first cluster
func (k *KubeConfig) GetFirstCluster() (string, *clientcmdapi.Cluster) {
	if k.Config == nil {
		return "", nil
	}
	for name, cluster := range k.Config.Clusters {
		return name, cluster
	}
	return "", nil
}

// GetFirstAuthInfo returns the first auth info
func (k *KubeConfig) GetFirstAuthInfo() (string, *clientcmdapi.AuthInfo) {
	if k.Config == nil {
		return "", nil
	}
	for name, authInfo := range k.Config.AuthInfos {
		return name, authInfo
	}
	return "", nil
}

// GetFirstContext returns the first context
func (k *KubeConfig) GetFirstContext() (string, *clientcmdapi.Context) {
	if k.Config == nil {
		return "", nil
	}
	for name, context := range k.Config.Contexts {
		return name, context
	}
	return "", nil
}

// GetCluster returns the cluster
func (k *KubeConfig) GetCluster(name string) *clientcmdapi.Cluster {
	if k.Config == nil {
		return nil
	}
	return k.Config.Clusters[name]
}

// GetAuthInfo returns the auth info
func (k *KubeConfig) GetAuthInfo(name string) *clientcmdapi.AuthInfo {
	if k.Config == nil {
		return nil
	}
	return k.Config.AuthInfos[name]
}

// GetContext returns the context
func (k *KubeConfig) GetContext(name string) *clientcmdapi.Context {
	if k.Config == nil {
		return nil
	}
	return k.Config.Contexts[name]
}
