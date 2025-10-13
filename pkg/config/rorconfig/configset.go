package rorconfig

import (
	"fmt"
	"os"

	"github.com/NorskHelsenett/ror/pkg/rlog"
)

type rorConfigSet struct {
	autoload bool
	configs  configsMap
}

func (rc *rorConfigSet) LoadEnv(key ConfigConst) {
	if ConfigConsts.IsDeprecated(key) {
		rlog.Warn(fmt.Sprintf("Config %s is deprecated %s", ConfigConsts[key].value, ConfigConsts[key].description))
	}
	rc.configs[key] = ConfigData(os.Getenv(ConfigConsts.GetEnvVariable(key)))
}

func (rc *rorConfigSet) AutoLoadAllEnv() {

	loadDotEnv()
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
	if _, exists := rc.configs[key]; !exists || rc.configs[key] == "" {
		rc.configs[key] = ConfigData(anyToString(defaultValue))
	}
}

func (rc *rorConfigSet) SetWithProvider(key ConfigConst, provider SecretProvider) {
	proveidervalue := provider.GetSecret()
	rc.configs[key] = ConfigData(proveidervalue)
}

func (rc *rorConfigSet) Set(key ConfigConst, value any) {

	rc.configs[key] = ConfigData(anyToString(value))
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
