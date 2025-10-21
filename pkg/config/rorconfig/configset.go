package rorconfig

import (
	"fmt"
	"maps"
	"os"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type rorConfigSet struct {
	autoload bool
	configs  configsMap
}

func (cm *configsMap) IsSet(key string) bool {
	_, exists := (*cm)[key]
	return exists
}

func (cm *configsMap) IsEmpty(key string) bool {
	value, exists := (*cm)[key]
	return !exists || value == ""
}

func (cm *configsMap) Set(key string, data any) {

	(*cm)[key] = ConfigData(anyToString(data))
}

func (cm *configsMap) Unset(key string) {
	delete(*cm, key)
}

func (cm *configsMap) Get(key string) ConfigData {
	return (*cm)[key]
}

func (cm *configsMap) GetAll() configsMap {
	copyMap := make(configsMap)
	maps.Copy(copyMap, *cm)
	return copyMap
}

func (rc *rorConfigSet) LoadEnv(key string) {
	constData, ok := ConfigConsts.GetEnvVariableConfigByKey(key)

	if !ok {
		ConfigConsts.Add(string(key))
	} else if ConfigConsts.IsDeprecated(key) {
		rlog.Warn(fmt.Sprintf("Config %s is deprecated %s", constData.key, constData.description))
	}

	data := os.Getenv(key)
	rc.configs.Set(key, data)
}

func (rc *rorConfigSet) AutoLoadAllEnv(local ...string) {
	if len(local) > 0 {
		for _, data := range local {
			ConfigConsts.Add(data)
		}
	}
	for key, value := range readDotEnv() {
		os.Setenv(key, value)
	}
	for _, value := range ConfigConsts {
		_, exists := os.LookupEnv(value.key)
		if exists {
			rc.LoadEnv(value.key)
		}
	}
}

func (rc *rorConfigSet) SetDefault(key string, defaultValue any) {
	if rc.autoload {
		rc.LoadEnv(key)
	}
	if rc.configs.IsEmpty(key) {
		rc.configs.Set(key, defaultValue)
	}
}

func (rc *rorConfigSet) SetWithProvider(key string, provider SecretProvider) {
	proveidervalue := provider.GetSecret()
	rc.configs.Set(key, proveidervalue)
}

func (rc *rorConfigSet) Set(key string, value any) {
	rc.configs.Set(key, value)
}

func (rc *rorConfigSet) IsSet(key string) bool {
	return rc.configs.IsSet(key)
}

func (rc *rorConfigSet) AutoLoadEnv() {
	rc.autoload = true
	for key := range rc.configs {
		rc.LoadEnv(key)
	}
}

func (rc *rorConfigSet) getValue(key string) ConfigData {
	value := rc.configs.Get(key)
	if rc.configs.IsEmpty(key) && rc.autoload {
		rc.LoadEnv(key)
		value = rc.configs.Get(key)
	}
	return value
}

func (rc *rorConfigSet) GetString(key string) string {
	return rc.getValue(key).String()
}
func (rc *rorConfigSet) GetBool(key string) bool {
	return rc.getValue(key).Bool()
}
func (rc *rorConfigSet) GetInt(key string) int {
	return rc.getValue(key).Int()
}
func (rc *rorConfigSet) GetInt64(key string) int64 {
	return rc.getValue(key).Int64()
}
func (rc *rorConfigSet) GetFloat64(key string) float64 {
	return rc.getValue(key).Float64()
}
func (rc *rorConfigSet) GetFloat32(key string) float32 {
	return rc.getValue(key).Float32()
}
func (rc *rorConfigSet) GetUint(key string) uint {
	return rc.getValue(key).Uint()
}
func (rc *rorConfigSet) GetUint64(key string) uint64 {
	return rc.getValue(key).Uint64()
}
func (rc *rorConfigSet) GetUint32(key string) uint32 {
	return rc.getValue(key).Uint32()
}
