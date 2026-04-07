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
	if data != "" {
		rc.configs.Set(key, data, ConfigSourceEnv)
	}

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

// ExportToStruct creates a new instance of T and populates its fields from
// the configuration store. Fields must be tagged with `rorconfig:"KEY"` to be
// populated. Untagged struct fields are recursed into. Supports the same set
// of types as ImportStruct: string, bool, int*, uint*, float*, time.Time,
// time.Duration, and pointers to any of these.
func ExportToStruct[T any](rc *rorConfigSet) (*T, error) {
	return exportToStruct[T](rc, nil)
}

// exportToStructFiltered is like ExportToStruct but skips config keys whose
// source is in excludeSources. Used by SaveToFile to omit env/flag overrides.
func exportToStructFiltered[T any](rc *rorConfigSet, excludeSources []ConfigSource) (*T, error) {
	return exportToStruct[T](rc, excludeSources)
}

func exportToStruct[T any](rc *rorConfigSet, excludeSources []ConfigSource) (*T, error) {
	cfg := new(T)
	if err := rc.exportToValue(reflect.ValueOf(cfg), excludeSources); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (rc *rorConfigSet) exportToValue(value reflect.Value, excludeSources []ConfigSource) error {
	// Dereference pointers to reach the struct.
	for value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return fmt.Errorf("rorconfig: cannot export to nil pointer")
		}
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return fmt.Errorf("rorconfig: target must be a struct or pointer to struct")
	}

	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		structField := typ.Field(i)

		if !field.CanSet() {
			continue
		}

		tag := structField.Tag.Get("rorconfig")
		if tag == "" {
			// Recurse into untagged struct / pointer-to-struct fields.
			if err := rc.exportRecurseUntagged(field, excludeSources); err != nil {
				return err
			}
			continue
		}

		if !rc.configs.IsSet(tag) {
			continue
		}

		cd := rc.getValue(tag)
		if isExcludedSource(cd.source, excludeSources) {
			continue
		}
		if err := setFieldFromConfigData(field, cd); err != nil {
			return fmt.Errorf("rorconfig: field %s: %w", structField.Name, err)
		}
	}

	return nil
}

func isExcludedSource(source ConfigSource, excluded []ConfigSource) bool {
	for _, e := range excluded {
		if source == e {
			return true
		}
	}
	return false
}

func (rc *rorConfigSet) exportRecurseUntagged(field reflect.Value, excludeSources []ConfigSource) error {
	switch field.Kind() {
	case reflect.Struct:
		return rc.exportToValue(field.Addr(), excludeSources)
	case reflect.Pointer:
		if field.Type().Elem().Kind() != reflect.Struct {
			return nil
		}
		// Only allocate a nil pointer-to-struct when the store actually
		// holds at least one non-excluded value for its tagged fields.
		// This avoids producing empty YAML sections on save.
		if field.IsNil() {
			if !rc.hasAnyTaggedKey(field.Type().Elem(), excludeSources) {
				return nil
			}
			field.Set(reflect.New(field.Type().Elem()))
		}
		return rc.exportToValue(field, excludeSources)
	default:
		return nil
	}
}

// hasAnyTaggedKey reports whether the store contains a value for at least one
// rorconfig-tagged field in the given struct type (recursing into untagged
// nested structs).
func (rc *rorConfigSet) hasAnyTaggedKey(t reflect.Type, excludeSources []ConfigSource) bool {
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		tag := sf.Tag.Get("rorconfig")
		if tag != "" {
			if rc.configs.IsSet(tag) && !isExcludedSource(rc.configs.Get(tag).source, excludeSources) {
				return true
			}
			continue
		}
		// Recurse into untagged struct / pointer-to-struct fields.
		ft := sf.Type
		for ft.Kind() == reflect.Pointer {
			ft = ft.Elem()
		}
		if ft.Kind() == reflect.Struct && rc.hasAnyTaggedKey(ft, excludeSources) {
			return true
		}
	}
	return false
}

// setFieldFromConfigData sets a reflect.Value from a ConfigData string,
// performing the reverse conversion of assignTaggedValue.
func setFieldFromConfigData(field reflect.Value, cd ConfigData) error {
	// Handle pointer fields: allocate and set the inner value.
	if field.Kind() == reflect.Pointer {
		inner := reflect.New(field.Type().Elem())
		if err := setFieldFromConfigData(inner.Elem(), cd); err != nil {
			return err
		}
		field.Set(inner)
		return nil
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(cd.String())
	case reflect.Bool:
		field.SetBool(cd.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if field.Type() == reflect.TypeFor[time.Duration]() {
			field.Set(reflect.ValueOf(cd.TimeDuration()))
			return nil
		}
		field.SetInt(cd.Int64())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		field.SetUint(cd.Uint64())
	case reflect.Float32:
		field.Set(reflect.ValueOf(cd.Float32()))
	case reflect.Float64:
		field.SetFloat(cd.Float64())
	case reflect.Struct:
		if field.Type() == reflect.TypeFor[time.Time]() {
			field.Set(reflect.ValueOf(cd.Time()))
			return nil
		}
		return fmt.Errorf("unsupported struct type %s", field.Type())
	default:
		return fmt.Errorf("unsupported kind %s", field.Kind())
	}
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
