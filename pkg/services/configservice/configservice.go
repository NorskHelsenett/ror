package configservice

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/NorskHelsenett/ror/pkg/clients/vaultclient"
	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"
	"github.com/NorskHelsenett/ror/pkg/rlog"
)

var (
	Loaders   = ConfigLoaders{}
	initiated bool
)

func init() {
	InitConfigService()
}

type ConfigLoaders map[string]ConfigLoaderInterface

type ConfigLoaderInterface interface {
	Load(key string) (string, error)
}

func InitConfigService() {
	if initiated {
		Loaders.AddLoader("env", NewEnvConfigLoader())
		initiated = true
	}
}

func (cl ConfigLoaders) AddLoader(source string, loader ConfigLoaderInterface) {
	InitConfigService()
	if loader != nil {
		cl[source] = loader
	}
}

func (cl ConfigLoaders) GetLoader(source string) ConfigLoaderInterface {
	InitConfigService()
	loader, ok := cl[source]
	if !ok {
		return nil
	}
	return loader
}

type VaultConfigLoader struct {
	client *vaultclient.VaultClient
}

func (vcl VaultConfigLoader) Load(key string) (string, error) {
	return vcl.client.GetSecretValueFromPath(key)
}

// vaultclient, err := vaultclient.New(context.Background(), vaultclient.NewStaticVaultCredsHelper(""), "http://localhost:9200")
//
// configservice.Loaders.AddLoader("vault", configservice.NewVaultConfigLoader(apiconnections.Vaultclient))
func NewVaultConfigLoader(client *vaultclient.VaultClient) ConfigLoaderInterface {
	if client == nil {
		rlog.Warn("vault client is nil, vault config loader will not work")
		return nil
	}
	return VaultConfigLoader{client: client}
}

type EnvConfigLoader struct{}

func (ecl EnvConfigLoader) Load(key string) (string, error) {
	return os.Getenv(key), nil
}

func NewEnvConfigLoader() EnvConfigLoader {
	return EnvConfigLoader{}
}

func AddLoader(source string, loader ConfigLoaderInterface) {
	Loaders.AddLoader(source, loader)
}

func GetLoader(source string) ConfigLoaderInterface {
	return Loaders.GetLoader(source)
}

// ConfigService Function loads the configuration for a resource based on the resource config and the context. It filters the config based on the identity in the context and loads the config from the specified source (e.g. vault) if needed.
// Template functions registered via AddLoader (e.g. "vault", "env") are available in the template and produce the final value directly.
func Template(templatevalue string, ctx context.Context) (string, error) {
	InitConfigService()
	identity := rorcontext.MustGetIdentityFromRorContext(ctx)

	data := map[string]string{
		"id":   identity.GetId(),
		"test": identity.GetId(), // For testing purposes, can be removed later
	}

	funcMap := make(template.FuncMap, len(Loaders))
	for source, loader := range Loaders {
		loader := loader
		funcMap[source] = func(key string) (string, error) {
			resolvedKey, err := renderLoaderKeyTemplate(key, data)
			if err != nil {
				return "", err
			}
			return loader.Load(resolvedKey)
		}
	}

	tmpl, err := template.New("value").Delims("{{", "}}").Funcs(funcMap).Parse(templatevalue)
	if err != nil {
		return "", fmt.Errorf("failed to parse template %q: %w", templatevalue, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %q: %w", templatevalue, err)
	}

	return buf.String(), nil
}

func renderLoaderKeyTemplate(key string, data map[string]string) (string, error) {
	if !strings.Contains(key, "{{") {
		return key, nil
	}

	tmpl, err := template.New("loader-key").Delims("{{", "}}").Parse(key)
	if err != nil {
		return "", fmt.Errorf("failed to parse loader key %q: %w", key, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute loader key %q: %w", key, err)
	}

	return buf.String(), nil
}
