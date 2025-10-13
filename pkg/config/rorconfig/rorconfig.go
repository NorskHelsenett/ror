// Package rorconfig provides a centralized configuration management system for ROR applications.
// It supports loading configuration values from environment variables, default values,
// and custom secret providers.
package rorconfig

// SecretProvider is an interface for custom secret retrieval mechanisms.
// Implementations can provide secrets from various sources like Vault, Kubernetes secrets, etc.
type SecretProvider interface {
	// GetSecret retrieves a secret value from the provider
	GetSecret() string
}

// configsMap is an internal type that maps configuration constants to their data
type configsMap map[ConfigConst]ConfigData

// config is the global configuration instance used by all package functions
var config = rorConfigSet{
	configs: make(configsMap),
}

// InitConfig initializes the configuration system by automatically loading
// all recognized environment variables. This should be called once at application startup.
func InitConfig() {
	config.AutoLoadAllEnv()
}

// IsSet checks if a configuration key has been set (either from environment variable,
// default value, or explicitly set value).
// Returns true if the key exists and has a value, false otherwise.
func IsSet(key ConfigConst) bool {
	return config.IsSet(key)
}

// Set explicitly sets a configuration value for the given key.
// This will override any environment variable or default value.
func Set(key ConfigConst, value any) {
	config.Set(key, value)
}

// SetDefault sets a default value for a configuration key.
// The default value will only be used if no environment variable or explicit value is set.
func SetDefault(key ConfigConst, defaultValue any) {
	config.SetDefault(key, defaultValue)
}

// GetConfigs returns a copy of all currently loaded configuration values.
// This is primarily used for debugging and testing purposes.
func GetConfigs() configsMap {
	return config.configs
}

// AutomaticEnv loads configuration values from environment variables for all registered keys.
// This is called internally by InitConfig and typically doesn't need to be called directly.
func AutomaticEnv() {
	config.AutoLoadEnv()
}

// SetWithProvider sets a configuration value using a custom SecretProvider.
// The provider's GetSecret() method will be called to retrieve the actual value.
// This is useful for integrating with external secret management systems.
func SetWithProvider(key ConfigConst, provider SecretProvider) {
	config.SetWithProvider(key, provider)
}

// GetString retrieves a configuration value as a string.
// Returns the string value of the configuration key, or empty string if not set.
func GetString(key ConfigConst) string {
	return config.GetString(key)
}

// GetBool retrieves a configuration value as a boolean.
// Returns true if the value is "true" (case-insensitive), false otherwise.
func GetBool(key ConfigConst) bool {
	return config.GetBool(key)
}

// GetInt retrieves a configuration value as an integer.
// Returns the integer value, or 0 if the value cannot be parsed as an integer.
func GetInt(key ConfigConst) int {
	return config.GetInt(key)
}

// GetInt64 retrieves a configuration value as a 64-bit integer.
// Returns the int64 value, or 0 if the value cannot be parsed as an int64.
func GetInt64(key ConfigConst) int64 {
	return config.GetInt64(key)
}

// GetFloat64 retrieves a configuration value as a 64-bit floating point number.
// Returns the float64 value, or 0.0 if the value cannot be parsed as a float64.
func GetFloat64(key ConfigConst) float64 {
	return config.GetFloat64(key)
}

// GetFloat32 retrieves a configuration value as a 32-bit floating point number.
// Returns the float32 value, or 0.0 if the value cannot be parsed as a float32.
func GetFloat32(key ConfigConst) float32 {
	return config.GetFloat32(key)
}

// GetUint retrieves a configuration value as an unsigned integer.
// Returns the uint value, or 0 if the value cannot be parsed as a uint.
func GetUint(key ConfigConst) uint {
	return config.GetUint(key)
}

// GetUint64 retrieves a configuration value as a 64-bit unsigned integer.
// Returns the uint64 value, or 0 if the value cannot be parsed as a uint64.
func GetUint64(key ConfigConst) uint64 {
	return config.GetUint64(key)
}

// GetUint32 retrieves a configuration value as a 32-bit unsigned integer.
// Returns the uint32 value, or 0 if the value cannot be parsed as a uint32.
func GetUint32(key ConfigConst) uint32 {
	return config.GetUint32(key)
}
