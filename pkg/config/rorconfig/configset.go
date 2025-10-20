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

func (cm *configsMap) IsSet(key ConfigConst) bool {
	_, exists := (*cm)[key]
	return exists
}

func (cm *configsMap) IsEmpty(key ConfigConst) bool {
	value, exists := (*cm)[key]
	return !exists || value == ""
}

func (cm *configsMap) Set(key ConfigConst, data any) {

	(*cm)[key] = ConfigData(anyToString(data))
}

func (cm *configsMap) Unset(key ConfigConst) {
	delete(*cm, key)
}

func (cm *configsMap) Get(key ConfigConst) ConfigData {
	return (*cm)[key]
}

func (cm *configsMap) GetAll() configsMap {
	copyMap := make(configsMap)
	maps.Copy(copyMap, *cm)
	return copyMap
}

func (rc *rorConfigSet) LoadEnv(key ConfigConst) {
	envVar := ConfigConsts.GetEnvVariable(key)
	constData := ConfigConsts[key]

	if constData.deprecated {
		rlog.Warn(fmt.Sprintf("Config %s is deprecated %s", constData.value, constData.description))
	}

	data := os.Getenv(envVar)
	rc.configs.Set(ConfigConst(key), data)
}

func (rc *rorConfigSet) AutoLoadAllEnv() {
	for key, value := range readDotEnv() {
		os.Setenv(key, value)
	}
	for key, value := range ConfigConsts {
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
	if rc.configs.IsEmpty(key) {
		rc.configs.Set(key, defaultValue)
	}
}

func (rc *rorConfigSet) SetWithProvider(key ConfigConst, provider SecretProvider) {
	proveidervalue := provider.GetSecret()
	rc.configs.Set(key, proveidervalue)
}

func (rc *rorConfigSet) Set(key ConfigConst, value any) {
	rc.configs.Set(key, value)
}

func (rc *rorConfigSet) IsSet(key ConfigConst) bool {
	return rc.configs.IsSet(key)
}

func (rc *rorConfigSet) AutoLoadEnv() {
	rc.autoload = true
	for key := range rc.configs {
		rc.LoadEnv(key)
	}
}

func (rc *rorConfigSet) getValue(key ConfigConst) ConfigData {
	value := rc.configs.Get(key)
	if rc.configs.IsEmpty(key) && rc.autoload {
		rc.LoadEnv(key)
		value = rc.configs.Get(key)
	}
	return value
}

func (rc *rorConfigSet) GetString(key ConfigConst) string {
	return rc.getValue(key).String()
}
func (rc *rorConfigSet) GetBool(key ConfigConst) bool {
	return rc.getValue(key).Bool()
}
func (rc *rorConfigSet) GetInt(key ConfigConst) int {
	return rc.getValue(key).Int()
}
func (rc *rorConfigSet) GetInt64(key ConfigConst) int64 {
	return rc.getValue(key).Int64()
}
func (rc *rorConfigSet) GetFloat64(key ConfigConst) float64 {
	return rc.getValue(key).Float64()
}
func (rc *rorConfigSet) GetFloat32(key ConfigConst) float32 {
	return rc.getValue(key).Float32()
}
func (rc *rorConfigSet) GetUint(key ConfigConst) uint {
	return rc.getValue(key).Uint()
}
func (rc *rorConfigSet) GetUint64(key ConfigConst) uint64 {
	return rc.getValue(key).Uint64()
}
func (rc *rorConfigSet) GetUint32(key ConfigConst) uint32 {
	return rc.getValue(key).Uint32()
}
