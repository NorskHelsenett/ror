// Package rorconfig provides a centralized configuration management system for ROR applications.
// It supports loading configuration values from environment variables, default values,
// and custom secret providers.
package rorconfig

import "time"

// configKey is a constraint that allows any type that can be represented as a string.
type configKey interface {
	~string
}

// SecretProvider is an interface for custom secret retrieval mechanisms.
// Implementations can provide secrets from various sources like Vault, Kubernetes secrets, etc.
type SecretProvider interface {
	// GetSecret retrieves a secret value from the provider
	GetSecret() string
}

// configsMap is an internal type that maps configuration constants to their data
type configsMap map[string]ConfigData

// config is the global configuration instance used by all package functions
var config = rorConfigSet{
	configs: make(configsMap),
}

// InitConfig initializes the configuration system by automatically loading
// all recognized environment variables. This should be called once at application startup.
func InitConfig() {
	config.AutoLoadAllEnv()
}

// IsSet checks if a configuration string(key) has been set (either from environment variable,
// default value, or explicitly set value).
// Returns true if the string(key) exists and has a value, false otherwise.
func IsSet[K configKey](key K) bool {
	return config.IsSet(string(key))
}

// Set explicitly sets a configuration value for the given string(key).
// This will override any environment variable or default value.
func Set[K configKey](key K, value any) {
	config.Set(string(key), value)
}

// SetDefault sets a default value for a configuration string(key).
// The default value will only be used if no environment variable or explicit value is set.
func SetDefault[K configKey](key K, defaultValue any) {
	config.SetDefault(string(key), defaultValue)
}

// GetConfigs returns a copy of all currently loaded configuration values.
// This is primarily used for debugging and testing purposes.
func GetConfigs() configsMap {
	test := config.configs.GetAll()
	return test
}

// SetConfigFromStruct registers configuration values from a struct using the rorconfig tag on each field.
// Tagged fields are written into the internal configuration map using their tag value as key.
// Supported field kinds are strings, booleans, integers, unsigned integers, and floats.
func SetConfigFromStruct(source any) error {
	return config.ImportStruct(source)
}

// AutomaticEnv loads configuration values from environment variables for all registered string(key)s.
// This is called internally by InitConfig and typically doesn't need to be called directly.
func AutomaticEnv() {
	config.AutoLoadEnv()
}

// SetWithProvider sets a configuration value using a custom SecretProvider.
// The provider's GetSecret() method will be called to retrieve the actual value.
// This is useful for integrating with external secret management systems.
func SetWithProvider[K configKey](key K, provider SecretProvider) {
	config.SetWithProvider(string(key), provider)
}

// GetString retrieves a configuration value as a string.
// Returns the string value of the configuration string(key), or empty string if not set.
func GetString[K configKey](key K) string {
	return config.GetString(string(key))
}

// GetBool retrieves a configuration value as a boolean.
// Returns true if the value is "true" (case-insensitive), false otherwise.
func GetBool[K configKey](key K) bool {
	return config.GetBool(string(key))
}

// GetInt retrieves a configuration value as an integer.
// Returns the integer value, or 0 if the value cannot be parsed as an integer.
func GetInt[K configKey](key K) int {
	return config.GetInt(string(key))
}

// GetInt64 retrieves a configuration value as a 64-bit integer.
// Returns the int64 value, or 0 if the value cannot be parsed as an int64.
func GetInt64[K configKey](key K) int64 {
	return config.GetInt64(string(key))
}

// GetFloat64 retrieves a configuration value as a 64-bit floating point number.
// Returns the float64 value, or 0.0 if the value cannot be parsed as a float64.
func GetFloat64[K configKey](key K) float64 {
	return config.GetFloat64(string(key))
}

// GetFloat32 retrieves a configuration value as a 32-bit floating point number.
// Returns the float32 value, or 0.0 if the value cannot be parsed as a float32.
func GetFloat32[K configKey](key K) float32 {
	return config.GetFloat32(string(key))
}

// GetUint retrieves a configuration value as an unsigned integer.
// Returns the uint value, or 0 if the value cannot be parsed as a uint.
func GetUint[K configKey](key K) uint {
	return config.GetUint(string(key))
}

// GetUint64 retrieves a configuration value as a 64-bit unsigned integer.
// Returns the uint64 value, or 0 if the value cannot be parsed as a uint64.
func GetUint64[K configKey](key K) uint64 {
	return config.GetUint64(string(key))
}

// GetUint32 retrieves a configuration value as a 32-bit unsigned integer.
// Returns the uint32 value, or 0 if the value cannot be parsed as a uint32.
func GetUint32[K configKey](key K) uint32 {
	return config.GetUint32(string(key))
}

// GetTime retrieves a configuration value as a time.Time.
// Returns the time.Time value, or the zero time if the value cannot be parsed as a time.
func GetTime[K configKey](key K) time.Time {
	return config.GetTime(string(key))
}

// GetTimeDuration retrieves a configuration value as a time.Duration.
// Returns the time.Duration value, or 0 if the value cannot be parsed as a duration.
func GetTimeDuration[K configKey](key K) time.Duration {
	return config.GetTimeDuration(string(key))
}
