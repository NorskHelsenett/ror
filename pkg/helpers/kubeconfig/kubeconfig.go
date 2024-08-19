package kubeconfig

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

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

// IsExpired check if the selected contexts token is expired
func (k *KubeConfig) IsExpired(context string) (bool, error) {
	if k == nil {
		return true, nil
	}
	if errs := k.HandleErrors(); errs != nil {
		return true, errs
	}

	if k.Config.Contexts[context] == nil || k.Config.Contexts[context].AuthInfo == "" {
		return true, nil
	}

	authInfo, exists := k.Config.AuthInfos[k.Config.Contexts[context].AuthInfo]
	if !exists {
		return true, nil
	}

	if authInfo.Token == "" {
		// we can't check if the token is expired if it's not set
		return true, nil
	}
	token, _, err := new(jwt.Parser).ParseUnverified(authInfo.Token, jwt.MapClaims{})

	if err != nil {
		return true, err
	}

	claims := token.Claims.(jwt.MapClaims)

	expiry := int64(claims["exp"].(float64))
	now := time.Now().Unix()

	if now > expiry {
		return true, nil
	}

	return false, nil
}

// LoadFromBytes loads the kubeconfig file
func (k *KubeConfig) MergeYaml(yaml []byte) *KubeConfig {

	mergeConfig, err := clientcmd.Load(yaml)
	if err != nil {
		k.Errors = append(k.Errors, err)
	}

	if k.Config != nil {
		for key, v := range mergeConfig.Clusters {
			k.Config.Clusters[key] = v
		}

		for key, v := range mergeConfig.AuthInfos {
			k.Config.AuthInfos[key] = v
		}

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
func (k *KubeConfig) GetContext() (string, error) {
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
