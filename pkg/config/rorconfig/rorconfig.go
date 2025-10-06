package rorconfig

import (
	"fmt"
	"os"
	"strings"

	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
	"github.com/NorskHelsenett/ror/pkg/rlog"
	"github.com/joho/godotenv"
)

type SecretProvider interface {
	GetSecret() string
}

var Config = rorConfigSet{
	configs: make(configsMap),
}

type rorConfigSet struct {
	clients  rorClients
	autoload bool
	configs  configsMap
}

type rorClients struct {
	rorclient *rorclient.RorClient
	k8sclient *kubernetesclient.K8sClientsets
}

type configsMap map[ConfigConst]ConfigData

func InitConfig() {
	Config.AutoLoadAllEnv()
}

func (rc *rorConfigSet) LoadEnv(key ConfigConst) {
	if ConfigConstsMap[key].deprecated {
		rlog.Warn(fmt.Sprintf("Config %s is deprecated %s", ConfigConstsMap[key].value, ConfigConstsMap[key].description))
	}
	rc.configs[key] = ConfigData(os.Getenv(ConfigConstsMap[ConfigConst(key)].value))
}

func loadDotEnv() {
	enfilevar := strings.Split(os.Getenv("ENV_FILE"), ",")
	if len(enfilevar) == 1 && enfilevar[0] == "" {
		enfilevar = []string{".env"}
	}
	godotenv.Load(enfilevar...)
}

func (rc *rorConfigSet) AutoLoadAllEnv() {

	loadDotEnv()
	for key, value := range ConfigConstsMap {
		_, exists := os.LookupEnv(value.value)
		if exists {
			rc.LoadEnv(ConfigConst(key))
		}
	}
}

func (rc *rorConfigSet) SetDefault(key ConfigConst, defaultValue any) {
	if rc.autoload {
		rc.LoadEnv(key)
	}
	if _, exists := rc.configs[key]; !exists || rc.configs[key] == "" {
		rc.configs[key] = ConfigData(AnyToString(defaultValue))
	}
}

func (rc *rorConfigSet) SetWithProvider(key ConfigConst, provider SecretProvider) {
	proveidervalue := provider.GetSecret()
	rc.configs[key] = ConfigData(proveidervalue)
}

func (rc *rorConfigSet) Set(key ConfigConst, value any) {

	rc.configs[key] = ConfigData(AnyToString(value))
}

func AnyToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case int64:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%f", v)
	case float32:
		return fmt.Sprintf("%f", v)
	case uint:
		return fmt.Sprintf("%d", v)
	case uint64:
		return fmt.Sprintf("%d", v)
	case uint32:
		return fmt.Sprintf("%d", v)
	case bool:
		return fmt.Sprintf("%t", v)
	default:
		rlog.Error("Unsupported type for config value", fmt.Errorf("type %T not supported", v))
		return ""
	}
}

func (rc *rorConfigSet) IsSet(key ConfigConst) bool {
	_, exists := rc.configs[key]
	return exists
}

func (rc *rorConfigSet) AutoLoadEnv() {
	rc.autoload = true
	for key := range rc.configs {
		rc.LoadEnv(key)
	}
}

func IsSet(key ConfigConst) bool {
	return Config.IsSet(key)
}

func Set(key ConfigConst, value any) {
	Config.Set(key, value)
}

func SetDefault(key ConfigConst, defaultValue any) {
	Config.SetDefault(key, defaultValue)
}

func GetConfigs() configsMap {
	return Config.configs
}

func AutomaticEnv() {
	Config.AutoLoadEnv()
}

///GetString, GetBool, GetInt, GetInt64, GetFloat64, GetFloat32, GetUint, GetUint64, GetUint32 are helper functions to get config values in different types.

func GetString(key ConfigConst) string {
	return Config.GetString(key)
}
func GetBool(key ConfigConst) bool {
	return Config.GetBool(key)
}
func GetInt(key ConfigConst) int {
	return Config.GetInt(key)
}
func GetInt64(key ConfigConst) int64 {
	return Config.GetInt64(key)
}
func GetFloat64(key ConfigConst) float64 {
	return Config.GetFloat64(key)
}
func GetFloat32(key ConfigConst) float32 {
	return Config.GetFloat32(key)
}
func GetUint(key ConfigConst) uint {
	return Config.GetUint(key)
}
func GetUint64(key ConfigConst) uint64 {
	return Config.GetUint64(key)
}
func GetUint32(key ConfigConst) uint32 {
	return Config.GetUint32(key)
}

// GetString, GetBool, GetInt, GetInt64, GetFloat64, GetFloat32, GetUint, GetUint64, GetUint32 are helper functions to get config values in different types.

func (rc *rorConfigSet) GetString(key ConfigConst) string {
	return rc.configs[key].String()
}
func (rc *rorConfigSet) GetBool(key ConfigConst) bool {
	return rc.configs[key].Bool()
}
func (rc *rorConfigSet) GetInt(key ConfigConst) int {
	return rc.configs[key].Int()
}
func (rc *rorConfigSet) GetInt64(key ConfigConst) int64 {
	return rc.configs[key].Int64()
}
func (rc *rorConfigSet) GetFloat64(key ConfigConst) float64 {
	return rc.configs[key].Float64()
}
func (rc *rorConfigSet) GetFloat32(key ConfigConst) float32 {
	return rc.configs[key].Float32()
}
func (rc *rorConfigSet) GetUint(key ConfigConst) uint {
	return rc.configs[key].Uint()
}
func (rc *rorConfigSet) GetUint64(key ConfigConst) uint64 {
	return rc.configs[key].Uint64()
}
func (rc *rorConfigSet) GetUint32(key ConfigConst) uint32 {
	return rc.configs[key].Uint32()
}

// GetK8sClient returns the Kubernetes client from the RorConfig.
func (rc *rorConfigSet) GetK8sClient() *kubernetesclient.K8sClientsets {
	return rc.clients.k8sclient
}

// GetRorClient returns the ROR client from the RorConfig.
func (rc *rorConfigSet) GetRorClient() *rorclient.RorClient {
	return rc.clients.rorclient
}
