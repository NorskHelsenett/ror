package rorconfig

import (
	"fmt"
	"maps"
	"os"
	"reflect"
	"time"

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
	return !exists || value.Value == ""
}

func (cm *configsMap) Set(key string, data any, source ...ConfigSource) {
	cd := ConfigData{
		Value: anyToString(data),
	}
	if len(source) == 1 {
		cd.source = source[0]
	}
	(*cm)[key] = cd
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
	rc.configs.Set(key, data, ConfigSourceEnv)
}

func (rc *rorConfigSet) ImportStruct(source any) error {
	if source == nil {
		return fmt.Errorf("source must be a struct or pointer to struct")
	}

	return rc.addStructValue(reflect.ValueOf(source))
}

func (rc *rorConfigSet) addStructValue(value reflect.Value) error {
	structValue, err := resolveStructValue(value, false)
	if err != nil {
		return err
	}

	typ := structValue.Type()
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		structField := typ.Field(i)

		if !field.CanInterface() {
			continue
		}

		tag := structField.Tag.Get("rorconfig")
		if tag == "" {
			nested, err := resolveStructValue(field, true)
			if err != nil {
				return err
			}
			if nested.IsValid() {
				if err := rc.addStructValue(nested); err != nil {
					return err
				}
			}
			continue
		}

		if err := rc.assignTaggedValue(tag, field, structField); err != nil {
			return err
		}
	}

	return nil
}

func resolveStructValue(value reflect.Value, allowNil bool) (reflect.Value, error) {
	if !value.IsValid() {
		if allowNil {
			return reflect.Value{}, nil
		}
		return reflect.Value{}, fmt.Errorf("source must be a struct or pointer to struct")
	}

	for value.Kind() == reflect.Pointer {
		if value.IsNil() {
			if allowNil {
				return reflect.Value{}, nil
			}
			return reflect.Value{}, fmt.Errorf("source must be a struct or pointer to struct")
		}
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		if allowNil {
			return reflect.Value{}, nil
		}
		return reflect.Value{}, fmt.Errorf("source must be a struct or pointer to struct")
	}

	return value, nil
}

func (rc *rorConfigSet) assignTaggedValue(tag string, field reflect.Value, structField reflect.StructField) error {
	resolved := field
	for resolved.Kind() == reflect.Pointer {
		if resolved.IsNil() {
			return nil
		}
		resolved = resolved.Elem()
	}

	if !resolved.IsValid() {
		return nil
	}

	var data any
	if resolved.IsZero() {
		return nil
	}

	switch resolved.Kind() {
	case reflect.String:
		data = resolved.String()
	case reflect.Bool:
		data = resolved.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if resolved.Type() == reflect.TypeFor[time.Duration]() {
			data = resolved.Interface().(time.Duration)
			break
		}
		data = resolved.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		data = resolved.Uint()
	case reflect.Float32, reflect.Float64:
		data = resolved.Float()
	case reflect.Struct:
		if resolved.Type() == reflect.TypeFor[time.Time]() {
			data = resolved.Interface().(time.Time).Format(time.RFC3339Nano)
			break
		}
		return fmt.Errorf("field %s has unsupported kind %s", structField.Name, resolved.Kind())
	default:
		return fmt.Errorf("field %s has unsupported kind %s", structField.Name, resolved.Kind())
	}

	rc.Set(tag, data, ConfigSourceConfigFile)
	return nil
}

func (rc *rorConfigSet) AutoLoadAllEnv(local ...string) {
	if len(local) > 0 {
		for _, data := range local {
			ConfigConsts.Add(data)
		}
	}
	for key, value := range readDotEnv() {
		err := os.Setenv(key, value)
		if err != nil {
			rlog.Error("failed to set env variable", err)
		}
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
		rc.configs.Set(key, defaultValue, ConfigSourceDefault)
	}
}

func (rc *rorConfigSet) SetWithProvider(key string, provider SecretProvider) {
	proveidervalue := provider.GetSecret()
	rc.configs.Set(key, proveidervalue)
}

func (rc *rorConfigSet) Set(key string, value any, source ...ConfigSource) {
	rc.configs.Set(key, value, source...)
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
func (rc *rorConfigSet) GetTime(key string) time.Time {
	return rc.getValue(key).Time()
}
func (rc *rorConfigSet) GetTimeDuration(key string) time.Duration {
	return rc.getValue(key).TimeDuration()
}
