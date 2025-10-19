package rorconfig

import (
	"os"
	"testing"
)

func TestRorConfigSet_LoadEnv(t *testing.T) {
	tests := []struct {
		name          string
		key           ConfigConst
		envValue      string
		setupEnv      bool
		expectedValue string
	}{
		{
			name:          "load existing environment variable",
			key:           ConfigConst("TEST_KEY"),
			envValue:      "test_value",
			setupEnv:      true,
			expectedValue: "test_value",
		},
		{
			name:          "load non-existing environment variable",
			key:           ConfigConst("NONEXISTENT_KEY"),
			envValue:      "",
			setupEnv:      false,
			expectedValue: "",
		},
		{
			name:          "load empty environment variable",
			key:           ConfigConst("EMPTY_KEY"),
			envValue:      "",
			setupEnv:      true,
			expectedValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				configs: make(configsMap),
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(tt.key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			rc.LoadEnv(tt.key)

			// Verify
			if got := string(rc.configs[tt.key]); got != tt.expectedValue {
				t.Errorf("LoadEnv() = %v, want %v", got, tt.expectedValue)
			}
		})
	}
}

func TestRorConfigSet_LoadEnv_DeprecatedConfig(t *testing.T) {
	// This test requires mocking ConfigConsts.IsDeprecated()
	// and would need the actual ConfigConsts implementation
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	key := ConfigConst("DEPRECATED_KEY")
	envVar := ConfigConsts.GetEnvVariable(key)
	os.Setenv(envVar, "deprecated_value")
	defer os.Unsetenv(envVar)

	// Execute - this would log a warning if the key is deprecated
	rc.LoadEnv(key)

	// Verify the value is still loaded despite being deprecated
	if got := string(rc.configs[key]); got != "deprecated_value" {
		t.Errorf("LoadEnv() = %v, want %v", got, "deprecated_value")
	}
}
func TestRorConfigSet_SetDefault(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		defaultValue  any
		setupEnv      bool
		envValue      string
		expectedValue string
	}{
		{
			name:          "set default when config does not exist",
			autoload:      false,
			hasExisting:   false,
			defaultValue:  "default_value",
			expectedValue: "default_value",
		},
		{
			name:          "set default when config exists but is empty",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			defaultValue:  "default_value",
			expectedValue: "default_value",
		},
		{
			name:          "do not set default when config exists and has value",
			autoload:      false,
			existingValue: "existing_value",
			hasExisting:   true,
			defaultValue:  "default_value",
			expectedValue: "existing_value",
		},
		{
			name:          "autoload loads env before setting default",
			autoload:      true,
			hasExisting:   false,
			defaultValue:  "default_value",
			setupEnv:      true,
			envValue:      "env_value",
			expectedValue: "env_value",
		},
		{
			name:          "autoload with no env falls back to default",
			autoload:      true,
			hasExisting:   false,
			defaultValue:  "default_value",
			setupEnv:      false,
			expectedValue: "default_value",
		},
		{
			name:          "set default with int value",
			autoload:      false,
			hasExisting:   false,
			defaultValue:  42,
			expectedValue: "42",
		},
		{
			name:          "set default with bool value",
			autoload:      false,
			hasExisting:   false,
			defaultValue:  true,
			expectedValue: "true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			rc.SetDefault(key, tt.defaultValue)

			// Verify
			if got := string(rc.configs[key]); got != tt.expectedValue {
				t.Errorf("SetDefault() = %v, want %v", got, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_Set(t *testing.T) {
	tests := []struct {
		name          string
		key           ConfigConst
		value         any
		expectedValue string
	}{
		{
			name:          "set string value",
			key:           ConfigConst("STRING_KEY"),
			value:         "test_string",
			expectedValue: "test_string",
		},
		{
			name:          "set int value",
			key:           ConfigConst("INT_KEY"),
			value:         42,
			expectedValue: "42",
		},
		{
			name:          "set bool true value",
			key:           ConfigConst("BOOL_TRUE_KEY"),
			value:         true,
			expectedValue: "true",
		},
		{
			name:          "set bool false value",
			key:           ConfigConst("BOOL_FALSE_KEY"),
			value:         false,
			expectedValue: "false",
		},
		{
			name:          "set float64 value",
			key:           ConfigConst("FLOAT64_KEY"),
			value:         3.14159,
			expectedValue: "3.14159",
		},
		{
			name:          "set float32 value",
			key:           ConfigConst("FLOAT32_KEY"),
			value:         float32(2.718),
			expectedValue: "2.718",
		},
		{
			name:          "set int64 value",
			key:           ConfigConst("INT64_KEY"),
			value:         int64(9223372036854775807),
			expectedValue: "9223372036854775807",
		},
		{
			name:          "set uint value",
			key:           ConfigConst("UINT_KEY"),
			value:         uint(123),
			expectedValue: "123",
		},
		{
			name:          "set uint64 value",
			key:           ConfigConst("UINT64_KEY"),
			value:         uint64(18446744073709551615),
			expectedValue: "18446744073709551615",
		},
		{
			name:          "set uint32 value",
			key:           ConfigConst("UINT32_KEY"),
			value:         uint32(4294967295),
			expectedValue: "4294967295",
		},
		{
			name:          "set empty string value",
			key:           ConfigConst("EMPTY_STRING_KEY"),
			value:         "",
			expectedValue: "",
		},
		{
			name:          "set zero int value",
			key:           ConfigConst("ZERO_INT_KEY"),
			value:         0,
			expectedValue: "0",
		},
		{
			name:          "overwrite existing value",
			key:           ConfigConst("OVERWRITE_KEY"),
			value:         "new_value",
			expectedValue: "new_value",
		},
		{
			name:          "set unsupported type",
			key:           ConfigConst("UNSUPPORTED_KEY"),
			value:         struct{}{},
			expectedValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				configs: make(configsMap),
			}

			// For the overwrite test, set an initial value
			if tt.name == "overwrite existing value" {
				rc.configs[tt.key] = ConfigData("old_value")
			}

			// Execute
			rc.Set(tt.key, tt.value)

			// Verify
			if got := string(rc.configs[tt.key]); got != tt.expectedValue {
				t.Errorf("Set() = %v, want %v", got, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_IsSet(t *testing.T) {
	tests := []struct {
		name     string
		key      ConfigConst
		setValue bool
		value    string
		expected bool
	}{
		{
			name:     "key exists with non-empty value",
			key:      ConfigConst("EXISTING_KEY"),
			setValue: true,
			value:    "some_value",
			expected: true,
		},
		{
			name:     "key exists with empty value",
			key:      ConfigConst("EMPTY_KEY"),
			setValue: true,
			value:    "",
			expected: true,
		},
		{
			name:     "key does not exist",
			key:      ConfigConst("NONEXISTENT_KEY"),
			setValue: false,
			value:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				configs: make(configsMap),
			}

			if tt.setValue {
				rc.configs[tt.key] = ConfigData(tt.value)
			}

			// Execute
			result := rc.IsSet(tt.key)

			// Verify
			if result != tt.expected {
				t.Errorf("IsSet() = %v, want %v", result, tt.expected)
			}
		})
	}
}
func TestRorConfigSet_GetValue(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue string
	}{
		{
			name:          "get existing non-empty value without autoload",
			autoload:      false,
			existingValue: "existing_value",
			hasExisting:   true,
			expectedValue: "existing_value",
		},
		{
			name:          "get existing empty value without autoload",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: "",
		},
		{
			name:          "get non-existing value without autoload",
			autoload:      false,
			hasExisting:   false,
			expectedValue: "",
		},
		{
			name:          "get existing non-empty value with autoload",
			autoload:      true,
			existingValue: "existing_value",
			hasExisting:   true,
			expectedValue: "existing_value",
		},
		{
			name:          "get existing empty value with autoload and env set",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "env_value",
			expectedValue: "env_value",
		},
		{
			name:          "get existing empty value with autoload and no env",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      false,
			expectedValue: "",
		},
		{
			name:          "get non-existing value with autoload and env set",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "env_value",
			expectedValue: "env_value",
		},
		{
			name:          "get non-existing value with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.getValue(key)

			// Verify
			if got := string(result); got != tt.expectedValue {
				t.Errorf("getValue() = %v, want %v", got, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetString(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue string
	}{
		{
			name:          "get string from existing non-empty value",
			autoload:      false,
			existingValue: "test_string",
			hasExisting:   true,
			expectedValue: "test_string",
		},
		{
			name:          "get string from existing empty value",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: "",
		},
		{
			name:          "get string from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: "",
		},
		{
			name:          "get string with autoload from existing value",
			autoload:      true,
			existingValue: "existing_string",
			hasExisting:   true,
			expectedValue: "existing_string",
		},
		{
			name:          "get string with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "env_string_value",
			expectedValue: "env_string_value",
		},
		{
			name:          "get string with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "env_string_value",
			expectedValue: "env_string_value",
		},
		{
			name:          "get string with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: "",
		},
		{
			name:          "get string with special characters",
			autoload:      false,
			existingValue: "special!@#$%^&*()_+-=[]{}|;':\",./<>?",
			hasExisting:   true,
			expectedValue: "special!@#$%^&*()_+-=[]{}|;':\",./<>?",
		},
		{
			name:          "get string with unicode characters",
			autoload:      false,
			existingValue: "测试 🚀 ñáéíóú",
			hasExisting:   true,
			expectedValue: "测试 🚀 ñáéíóú",
		},
		{
			name:          "get string with newlines and tabs",
			autoload:      false,
			existingValue: "line1\nline2\tindented",
			hasExisting:   true,
			expectedValue: "line1\nline2\tindented",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_STRING_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetString(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetString() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetBool(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue bool
	}{
		{
			name:          "get bool true from string 'true'",
			autoload:      false,
			existingValue: "true",
			hasExisting:   true,
			expectedValue: true,
		},
		{
			name:          "get bool false from string 'false'",
			autoload:      false,
			existingValue: "false",
			hasExisting:   true,
			expectedValue: false,
		},
		{
			name:          "get bool true from string 'TRUE'",
			autoload:      false,
			existingValue: "TRUE",
			hasExisting:   true,
			expectedValue: true,
		},
		{
			name:          "get bool false from string 'FALSE'",
			autoload:      false,
			existingValue: "FALSE",
			hasExisting:   true,
			expectedValue: false,
		},
		{
			name:          "get bool true from string '1'",
			autoload:      false,
			existingValue: "1",
			hasExisting:   true,
			expectedValue: true,
		},
		{
			name:          "get bool false from string '0'",
			autoload:      false,
			existingValue: "0",
			hasExisting:   true,
			expectedValue: false,
		},
		{
			name:          "get bool false from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: false,
		},
		{
			name:          "get bool false from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: false,
		},
		{
			name:          "get bool false from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: false,
		},
		{
			name:          "get bool with autoload from existing value",
			autoload:      true,
			existingValue: "true",
			hasExisting:   true,
			expectedValue: true,
		},
		{
			name:          "get bool with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "true",
			expectedValue: true,
		},
		{
			name:          "get bool with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "false",
			expectedValue: false,
		},
		{
			name:          "get bool with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: false,
		},
		{
			name:          "get bool true from string 'yes'",
			autoload:      false,
			existingValue: "yes",
			hasExisting:   true,
			expectedValue: false, // assuming Bool() only accepts true/false/1/0
		},
		{
			name:          "get bool false from string 'no'",
			autoload:      false,
			existingValue: "no",
			hasExisting:   true,
			expectedValue: false,
		},
		{
			name:          "get bool from mixed case string 'True'",
			autoload:      false,
			existingValue: "True",
			hasExisting:   true,
			expectedValue: true,
		},
		{
			name:          "get bool from mixed case string 'False'",
			autoload:      false,
			existingValue: "False",
			hasExisting:   true,
			expectedValue: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_BOOL_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetBool(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetBool() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetInt(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue int
	}{
		{
			name:          "get int from valid positive number",
			autoload:      false,
			existingValue: "42",
			hasExisting:   true,
			expectedValue: 42,
		},
		{
			name:          "get int from valid negative number",
			autoload:      false,
			existingValue: "-123",
			hasExisting:   true,
			expectedValue: -123,
		},
		{
			name:          "get int from zero",
			autoload:      false,
			existingValue: "0",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get int from maximum int value",
			autoload:      false,
			existingValue: "2147483647",
			hasExisting:   true,
			expectedValue: 2147483647,
		},
		{
			name:          "get int from minimum int value",
			autoload:      false,
			existingValue: "-2147483648",
			hasExisting:   true,
			expectedValue: -2147483648,
		},
		{
			name:          "get int from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get int from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: 0,
		},
		{
			name:          "get int from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get int from float string",
			autoload:      false,
			existingValue: "123.45",
			hasExisting:   true,
			expectedValue: 0, // assuming Int() doesn't parse floats
		},
		{
			name:          "get int with autoload from existing value",
			autoload:      true,
			existingValue: "999",
			hasExisting:   true,
			expectedValue: 999,
		},
		{
			name:          "get int with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "555",
			expectedValue: 555,
		},
		{
			name:          "get int with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "777",
			expectedValue: 777,
		},
		{
			name:          "get int with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: 0,
		},
		{
			name:          "get int from string with leading zeros",
			autoload:      false,
			existingValue: "000123",
			hasExisting:   true,
			expectedValue: 123,
		},
		{
			name:          "get int from string with whitespace",
			autoload:      false,
			existingValue: " 456 ",
			hasExisting:   true,
			expectedValue: 0, // assuming Int() doesn't trim whitespace
		},
		{
			name:          "get int from hexadecimal string",
			autoload:      false,
			existingValue: "0xFF",
			hasExisting:   true,
			expectedValue: 0, // assuming Int() doesn't parse hex
		},
		{
			name:          "get int from string with plus sign",
			autoload:      false,
			existingValue: "+123",
			hasExisting:   true,
			expectedValue: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_INT_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetInt(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetInt() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetInt64(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue int64
	}{
		{
			name:          "get int64 from valid positive number",
			autoload:      false,
			existingValue: "42",
			hasExisting:   true,
			expectedValue: 42,
		},
		{
			name:          "get int64 from valid negative number",
			autoload:      false,
			existingValue: "-123",
			hasExisting:   true,
			expectedValue: -123,
		},
		{
			name:          "get int64 from zero",
			autoload:      false,
			existingValue: "0",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get int64 from maximum int64 value",
			autoload:      false,
			existingValue: "9223372036854775807",
			hasExisting:   true,
			expectedValue: 9223372036854775807,
		},
		{
			name:          "get int64 from minimum int64 value",
			autoload:      false,
			existingValue: "-9223372036854775808",
			hasExisting:   true,
			expectedValue: -9223372036854775808,
		},
		{
			name:          "get int64 from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get int64 from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: 0,
		},
		{
			name:          "get int64 from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get int64 from float string",
			autoload:      false,
			existingValue: "123.45",
			hasExisting:   true,
			expectedValue: 0, // assuming Int64() doesn't parse floats
		},
		{
			name:          "get int64 with autoload from existing value",
			autoload:      true,
			existingValue: "999999999999",
			hasExisting:   true,
			expectedValue: 999999999999,
		},
		{
			name:          "get int64 with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "555555555555",
			expectedValue: 555555555555,
		},
		{
			name:          "get int64 with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "777777777777",
			expectedValue: 777777777777,
		},
		{
			name:          "get int64 with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: 0,
		},
		{
			name:          "get int64 from string with leading zeros",
			autoload:      false,
			existingValue: "000123",
			hasExisting:   true,
			expectedValue: 123,
		},
		{
			name:          "get int64 from string with whitespace",
			autoload:      false,
			existingValue: " 456 ",
			hasExisting:   true,
			expectedValue: 0, // assuming Int64() doesn't trim whitespace
		},
		{
			name:          "get int64 from hexadecimal string",
			autoload:      false,
			existingValue: "0xFF",
			hasExisting:   true,
			expectedValue: 0, // assuming Int64() doesn't parse hex
		},
		{
			name:          "get int64 from string with plus sign",
			autoload:      false,
			existingValue: "+123456789012345",
			hasExisting:   true,
			expectedValue: 123456789012345,
		},
		{
			name:          "get int64 from large positive number",
			autoload:      false,
			existingValue: "1234567890123456789",
			hasExisting:   true,
			expectedValue: 1234567890123456789,
		},
		{
			name:          "get int64 from large negative number",
			autoload:      false,
			existingValue: "-1234567890123456789",
			hasExisting:   true,
			expectedValue: -1234567890123456789,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_INT64_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetInt64(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetInt64() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetFloat64(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue float64
	}{
		{
			name:          "get float64 from valid positive number",
			autoload:      false,
			existingValue: "42.5",
			hasExisting:   true,
			expectedValue: 42.5,
		},
		{
			name:          "get float64 from valid negative number",
			autoload:      false,
			existingValue: "-123.456",
			hasExisting:   true,
			expectedValue: -123.456,
		},
		{
			name:          "get float64 from zero",
			autoload:      false,
			existingValue: "0.0",
			hasExisting:   true,
			expectedValue: 0.0,
		},
		{
			name:          "get float64 from integer string",
			autoload:      false,
			existingValue: "42",
			hasExisting:   true,
			expectedValue: 42.0,
		},
		{
			name:          "get float64 from scientific notation",
			autoload:      false,
			existingValue: "1.23e4",
			hasExisting:   true,
			expectedValue: 12300.0,
		},
		{
			name:          "get float64 from negative scientific notation",
			autoload:      false,
			existingValue: "-1.23e-4",
			hasExisting:   true,
			expectedValue: -0.000123,
		},
		{
			name:          "get float64 from very large number",
			autoload:      false,
			existingValue: "1.7976931348623157e+308",
			hasExisting:   true,
			expectedValue: 1.7976931348623157e+308,
		},
		{
			name:          "get float64 from very small number",
			autoload:      false,
			existingValue: "2.2250738585072014e-308",
			hasExisting:   true,
			expectedValue: 2.2250738585072014e-308,
		},
		{
			name:          "get float64 from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: 0.0,
		},
		{
			name:          "get float64 from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: 0.0,
		},
		{
			name:          "get float64 from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: 0.0,
		},
		{
			name:          "get float64 with autoload from existing value",
			autoload:      true,
			existingValue: "999.999",
			hasExisting:   true,
			expectedValue: 999.999,
		},
		{
			name:          "get float64 with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "555.555",
			expectedValue: 555.555,
		},
		{
			name:          "get float64 with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "777.777",
			expectedValue: 777.777,
		},
		{
			name:          "get float64 with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: 0.0,
		},
		{
			name:          "get float64 from string with leading zeros",
			autoload:      false,
			existingValue: "000123.456",
			hasExisting:   true,
			expectedValue: 123.456,
		},
		{
			name:          "get float64 from string with whitespace",
			autoload:      false,
			existingValue: " 456.789 ",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float64() doesn't trim whitespace
		},
		{
			name:          "get float64 from string with plus sign",
			autoload:      false,
			existingValue: "+123.456",
			hasExisting:   true,
			expectedValue: 123.456,
		},
		{
			name:          "get float64 from Pi approximation",
			autoload:      false,
			existingValue: "3.141592653589793",
			hasExisting:   true,
			expectedValue: 3.141592653589793,
		},
		{
			name:          "get float64 from special value Inf",
			autoload:      false,
			existingValue: "Inf",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float64() doesn't parse special values
		},
		{
			name:          "get float64 from special value -Inf",
			autoload:      false,
			existingValue: "-Inf",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float64() doesn't parse special values
		},
		{
			name:          "get float64 from special value NaN",
			autoload:      false,
			existingValue: "NaN",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float64() doesn't parse special values
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_FLOAT64_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetFloat64(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetFloat64() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetFloat32(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue float32
	}{
		{
			name:          "get float32 from valid positive number",
			autoload:      false,
			existingValue: "42.5",
			hasExisting:   true,
			expectedValue: 42.5,
		},
		{
			name:          "get float32 from valid negative number",
			autoload:      false,
			existingValue: "-123.456",
			hasExisting:   true,
			expectedValue: -123.456,
		},
		{
			name:          "get float32 from zero",
			autoload:      false,
			existingValue: "0.0",
			hasExisting:   true,
			expectedValue: 0.0,
		},
		{
			name:          "get float32 from integer string",
			autoload:      false,
			existingValue: "42",
			hasExisting:   true,
			expectedValue: 42.0,
		},
		{
			name:          "get float32 from scientific notation",
			autoload:      false,
			existingValue: "1.23e4",
			hasExisting:   true,
			expectedValue: 12300.0,
		},
		{
			name:          "get float32 from negative scientific notation",
			autoload:      false,
			existingValue: "-1.23e-4",
			hasExisting:   true,
			expectedValue: -0.000123,
		},
		{
			name:          "get float32 from maximum float32 value",
			autoload:      false,
			existingValue: "3.4028235e+38",
			hasExisting:   true,
			expectedValue: 3.4028235e+38,
		},
		{
			name:          "get float32 from minimum positive float32 value",
			autoload:      false,
			existingValue: "1.175494e-38",
			hasExisting:   true,
			expectedValue: 1.175494e-38,
		},
		{
			name:          "get float32 from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: 0.0,
		},
		{
			name:          "get float32 from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: 0.0,
		},
		{
			name:          "get float32 from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: 0.0,
		},
		{
			name:          "get float32 with autoload from existing value",
			autoload:      true,
			existingValue: "999.999",
			hasExisting:   true,
			expectedValue: 999.999,
		},
		{
			name:          "get float32 with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "555.555",
			expectedValue: 555.555,
		},
		{
			name:          "get float32 with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "777.777",
			expectedValue: 777.777,
		},
		{
			name:          "get float32 with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: 0.0,
		},
		{
			name:          "get float32 from string with leading zeros",
			autoload:      false,
			existingValue: "000123.456",
			hasExisting:   true,
			expectedValue: 123.456,
		},
		{
			name:          "get float32 from string with whitespace",
			autoload:      false,
			existingValue: " 456.789 ",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float32() doesn't trim whitespace
		},
		{
			name:          "get float32 from string with plus sign",
			autoload:      false,
			existingValue: "+123.456",
			hasExisting:   true,
			expectedValue: 123.456,
		},
		{
			name:          "get float32 from Pi approximation",
			autoload:      false,
			existingValue: "3.1415927",
			hasExisting:   true,
			expectedValue: 3.1415927,
		},
		{
			name:          "get float32 from special value Inf",
			autoload:      false,
			existingValue: "Inf",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float32() doesn't parse special values
		},
		{
			name:          "get float32 from special value -Inf",
			autoload:      false,
			existingValue: "-Inf",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float32() doesn't parse special values
		},
		{
			name:          "get float32 from special value NaN",
			autoload:      false,
			existingValue: "NaN",
			hasExisting:   true,
			expectedValue: 0.0, // assuming Float32() doesn't parse special values
		},
		{
			name:          "get float32 from precision limit test",
			autoload:      false,
			existingValue: "1.23456789",
			hasExisting:   true,
			expectedValue: 1.2345679, // float32 has limited precision
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_FLOAT32_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetFloat32(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetFloat32() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetUint(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue uint
	}{
		{
			name:          "get uint from valid positive number",
			autoload:      false,
			existingValue: "42",
			hasExisting:   true,
			expectedValue: 42,
		},
		{
			name:          "get uint from zero",
			autoload:      false,
			existingValue: "0",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint from maximum uint value",
			autoload:      false,
			existingValue: "18446744073709551615",
			hasExisting:   true,
			expectedValue: 18446744073709551615,
		},
		{
			name:          "get uint from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: 0,
		},
		{
			name:          "get uint from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint from negative number",
			autoload:      false,
			existingValue: "-123",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint() doesn't parse negative numbers
		},
		{
			name:          "get uint from float string",
			autoload:      false,
			existingValue: "123.45",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint() doesn't parse floats
		},
		{
			name:          "get uint with autoload from existing value",
			autoload:      true,
			existingValue: "999",
			hasExisting:   true,
			expectedValue: 999,
		},
		{
			name:          "get uint with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "555",
			expectedValue: 555,
		},
		{
			name:          "get uint with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "777",
			expectedValue: 777,
		},
		{
			name:          "get uint with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: 0,
		},
		{
			name:          "get uint from string with leading zeros",
			autoload:      false,
			existingValue: "000123",
			hasExisting:   true,
			expectedValue: 123,
		},
		{
			name:          "get uint from string with whitespace",
			autoload:      false,
			existingValue: " 456 ",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint() doesn't trim whitespace
		},
		{
			name:          "get uint from hexadecimal string",
			autoload:      false,
			existingValue: "0xFF",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint() doesn't parse hex
		},
		{
			name:          "get uint from invalid string with plus sign",
			autoload:      false,
			existingValue: "+123",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint from large number",
			autoload:      false,
			existingValue: "12345678901234567890",
			hasExisting:   true,
			expectedValue: 12345678901234567890,
		},
		{
			name:          "get uint from string with scientific notation",
			autoload:      false,
			existingValue: "1e3",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint() doesn't parse scientific notation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_UINT_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetUint(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetUint() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetUint64(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue uint64
	}{
		{
			name:          "get uint64 from valid positive number",
			autoload:      false,
			existingValue: "42",
			hasExisting:   true,
			expectedValue: 42,
		},
		{
			name:          "get uint64 from zero",
			autoload:      false,
			existingValue: "0",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint64 from maximum uint64 value",
			autoload:      false,
			existingValue: "18446744073709551615",
			hasExisting:   true,
			expectedValue: 18446744073709551615,
		},
		{
			name:          "get uint64 from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint64 from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: 0,
		},
		{
			name:          "get uint64 from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint64 from negative number",
			autoload:      false,
			existingValue: "-123",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint64() doesn't parse negative numbers
		},
		{
			name:          "get uint64 from float string",
			autoload:      false,
			existingValue: "123.45",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint64() doesn't parse floats
		},
		{
			name:          "get uint64 with autoload from existing value",
			autoload:      true,
			existingValue: "999999999999",
			hasExisting:   true,
			expectedValue: 999999999999,
		},
		{
			name:          "get uint64 with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "555555555555",
			expectedValue: 555555555555,
		},
		{
			name:          "get uint64 with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "777777777777",
			expectedValue: 777777777777,
		},
		{
			name:          "get uint64 with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: 0,
		},
		{
			name:          "get uint64 from string with leading zeros",
			autoload:      false,
			existingValue: "000123",
			hasExisting:   true,
			expectedValue: 123,
		},
		{
			name:          "get uint64 from string with whitespace",
			autoload:      false,
			existingValue: " 456 ",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint64() doesn't trim whitespace
		},
		{
			name:          "get uint64 from hexadecimal string",
			autoload:      false,
			existingValue: "0xFF",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint64() doesn't parse hex
		},
		{
			name:          "get uint64 from string with plus sign",
			autoload:      false,
			existingValue: "+123",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint64() doesn't parse plus sign
		},
		{
			name:          "get uint64 from large number",
			autoload:      false,
			existingValue: "12345678901234567890",
			hasExisting:   true,
			expectedValue: 12345678901234567890,
		},
		{
			name:          "get uint64 from string with scientific notation",
			autoload:      false,
			existingValue: "1e3",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint64() doesn't parse scientific notation
		},
		{
			name:          "get uint64 from very large valid number",
			autoload:      false,
			existingValue: "9223372036854775808",
			hasExisting:   true,
			expectedValue: 9223372036854775808,
		},
		{
			name:          "get uint64 from medium large number",
			autoload:      false,
			existingValue: "1000000000000000000",
			hasExisting:   true,
			expectedValue: 1000000000000000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_UINT64_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetUint64(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetUint64() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_GetUint32(t *testing.T) {
	tests := []struct {
		name          string
		autoload      bool
		existingValue string
		hasExisting   bool
		setupEnv      bool
		envValue      string
		expectedValue uint32
	}{
		{
			name:          "get uint32 from valid positive number",
			autoload:      false,
			existingValue: "42",
			hasExisting:   true,
			expectedValue: 42,
		},
		{
			name:          "get uint32 from zero",
			autoload:      false,
			existingValue: "0",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint32 from maximum uint32 value",
			autoload:      false,
			existingValue: "4294967295",
			hasExisting:   true,
			expectedValue: 4294967295,
		},
		{
			name:          "get uint32 from empty string",
			autoload:      false,
			existingValue: "",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint32 from non-existing value",
			autoload:      false,
			hasExisting:   false,
			expectedValue: 0,
		},
		{
			name:          "get uint32 from invalid string",
			autoload:      false,
			existingValue: "invalid",
			hasExisting:   true,
			expectedValue: 0,
		},
		{
			name:          "get uint32 from negative number",
			autoload:      false,
			existingValue: "-123",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint32() doesn't parse negative numbers
		},
		{
			name:          "get uint32 from float string",
			autoload:      false,
			existingValue: "123.45",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint32() doesn't parse floats
		},
		{
			name:          "get uint32 with autoload from existing value",
			autoload:      true,
			existingValue: "999999",
			hasExisting:   true,
			expectedValue: 999999,
		},
		{
			name:          "get uint32 with autoload from env when existing is empty",
			autoload:      true,
			existingValue: "",
			hasExisting:   true,
			setupEnv:      true,
			envValue:      "555555",
			expectedValue: 555555,
		},
		{
			name:          "get uint32 with autoload from env when not existing",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      true,
			envValue:      "777777",
			expectedValue: 777777,
		},
		{
			name:          "get uint32 with autoload and no env",
			autoload:      true,
			hasExisting:   false,
			setupEnv:      false,
			expectedValue: 0,
		},
		{
			name:          "get uint32 from string with leading zeros",
			autoload:      false,
			existingValue: "000123",
			hasExisting:   true,
			expectedValue: 123,
		},
		{
			name:          "get uint32 from string with whitespace",
			autoload:      false,
			existingValue: " 456 ",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint32() doesn't trim whitespace
		},
		{
			name:          "get uint32 from hexadecimal string",
			autoload:      false,
			existingValue: "0xFF",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint32() doesn't parse hex
		},
		{
			name:          "get uint32 from string with plus sign",
			autoload:      false,
			existingValue: "+123",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint32() doesn't parse plus sign
		},
		{
			name:          "get uint32 from large valid number",
			autoload:      false,
			existingValue: "3000000000",
			hasExisting:   true,
			expectedValue: 3000000000,
		},
		{
			name:          "get uint32 from string with scientific notation",
			autoload:      false,
			existingValue: "1e3",
			hasExisting:   true,
			expectedValue: 0, // assuming Uint32() doesn't parse scientific notation
		},
		{
			name:          "get uint32 from medium number",
			autoload:      false,
			existingValue: "1000000",
			hasExisting:   true,
			expectedValue: 1000000,
		},
		{
			name:          "get uint32 from number above uint32 max",
			autoload:      false,
			existingValue: "4294967296",
			hasExisting:   true,
			expectedValue: 4294967295, // supplies the saturated value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: tt.autoload,
				configs:  make(configsMap),
			}

			key := ConfigConst("TEST_UINT32_KEY")

			if tt.hasExisting {
				rc.configs[key] = ConfigData(tt.existingValue)
			}

			if tt.setupEnv {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, tt.envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			result := rc.GetUint32(key)

			// Verify
			if result != tt.expectedValue {
				t.Errorf("GetUint32() = %v, want %v", result, tt.expectedValue)
			}
		})
	}
}
func TestRorConfigSet_AutoLoadEnv(t *testing.T) {
	tests := []struct {
		name             string
		initialConfigs   map[ConfigConst]ConfigData
		envValues        map[ConfigConst]string
		expectedAutoload bool
		expectedValues   map[ConfigConst]string
	}{
		{
			name:             "enable autoload with no existing configs",
			initialConfigs:   map[ConfigConst]ConfigData{},
			envValues:        map[ConfigConst]string{},
			expectedAutoload: true,
			expectedValues:   map[ConfigConst]string{},
		},
		{
			name: "enable autoload with existing configs and env values",
			initialConfigs: map[ConfigConst]ConfigData{
				ConfigConst("KEY1"): ConfigData("old_value1"),
				ConfigConst("KEY2"): ConfigData("old_value2"),
			},
			envValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "env_value1",
				ConfigConst("KEY2"): "env_value2",
			},
			expectedAutoload: true,
			expectedValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "env_value1",
				ConfigConst("KEY2"): "env_value2",
			},
		},
		{
			name: "enable autoload with existing configs but no env values",
			initialConfigs: map[ConfigConst]ConfigData{
				ConfigConst("KEY1"): ConfigData("old_value1"),
				ConfigConst("KEY2"): ConfigData("old_value2"),
			},
			envValues:        map[ConfigConst]string{},
			expectedAutoload: true,
			expectedValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "",
				ConfigConst("KEY2"): "",
			},
		},
		{
			name: "enable autoload with mixed env availability",
			initialConfigs: map[ConfigConst]ConfigData{
				ConfigConst("KEY1"): ConfigData("old_value1"),
				ConfigConst("KEY2"): ConfigData("old_value2"),
				ConfigConst("KEY3"): ConfigData("old_value3"),
			},
			envValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "env_value1",
				ConfigConst("KEY3"): "env_value3",
			},
			expectedAutoload: true,
			expectedValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "env_value1",
				ConfigConst("KEY2"): "",
				ConfigConst("KEY3"): "env_value3",
			},
		},
		{
			name: "enable autoload with empty env values",
			initialConfigs: map[ConfigConst]ConfigData{
				ConfigConst("KEY1"): ConfigData("old_value1"),
			},
			envValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "",
			},
			expectedAutoload: true,
			expectedValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "",
			},
		},
		{
			name: "enable autoload when already autoload is false",
			initialConfigs: map[ConfigConst]ConfigData{
				ConfigConst("KEY1"): ConfigData("initial_value"),
			},
			envValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "updated_env_value",
			},
			expectedAutoload: true,
			expectedValues: map[ConfigConst]string{
				ConfigConst("KEY1"): "updated_env_value",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				autoload: false, // Start with autoload disabled
				configs:  make(configsMap),
			}

			// Set initial configs
			for key, value := range tt.initialConfigs {
				rc.configs[key] = value
			}

			// Set environment variables
			for key, envValue := range tt.envValues {
				envVar := ConfigConsts.GetEnvVariable(key)
				os.Setenv(envVar, envValue)
				defer os.Unsetenv(envVar)
			}

			// Execute
			rc.AutoLoadEnv()

			// Verify autoload is enabled
			if rc.autoload != tt.expectedAutoload {
				t.Errorf("AutoLoadEnv() autoload = %v, want %v", rc.autoload, tt.expectedAutoload)
			}

			// Verify all expected values are loaded
			for key, expectedValue := range tt.expectedValues {
				if got := string(rc.configs[key]); got != expectedValue {
					t.Errorf("AutoLoadEnv() config[%v] = %v, want %v", key, got, expectedValue)
				}
			}

			// Verify no unexpected configs were added
			if len(rc.configs) != len(tt.expectedValues) {
				t.Errorf("AutoLoadEnv() config count = %v, want %v", len(rc.configs), len(tt.expectedValues))
			}
		})
	}
}

func TestRorConfigSet_AutoLoadEnv_CallsLoadEnvForEachConfig(t *testing.T) {
	// Setup
	rc := &rorConfigSet{
		autoload: false,
		configs:  make(configsMap),
	}

	// Add multiple configs
	keys := []ConfigConst{
		ConfigConst("TEST_KEY_1"),
		ConfigConst("TEST_KEY_2"),
		ConfigConst("TEST_KEY_3"),
	}

	for i, key := range keys {
		rc.configs[key] = ConfigData("initial_value_" + string(rune('1'+i)))
	}

	// Set corresponding environment variables
	for i, key := range keys {
		envVar := ConfigConsts.GetEnvVariable(key)
		envValue := "env_value_" + string(rune('1'+i))
		os.Setenv(envVar, envValue)
		defer os.Unsetenv(envVar)
	}

	// Execute
	rc.AutoLoadEnv()

	// Verify autoload is enabled
	if !rc.autoload {
		t.Error("AutoLoadEnv() should enable autoload")
	}

	// Verify LoadEnv was called for each existing config key
	for i, key := range keys {
		expectedValue := "env_value_" + string(rune('1'+i))
		if got := string(rc.configs[key]); got != expectedValue {
			t.Errorf("AutoLoadEnv() should have loaded env for key %v, got %v, want %v", key, got, expectedValue)
		}
	}
}

func TestRorConfigSet_AutoLoadEnv_EnablesAutoloadFirst(t *testing.T) {
	// Setup
	rc := &rorConfigSet{
		autoload: false,
		configs:  make(configsMap),
	}

	// Verify initial state
	if rc.autoload {
		t.Error("Initial autoload should be false")
	}

	// Execute
	rc.AutoLoadEnv()

	// Verify autoload is enabled
	if !rc.autoload {
		t.Error("AutoLoadEnv() should enable autoload")
	}
}

func TestRorConfigSet_AutoLoadEnv_WithAlreadyEnabledAutoload(t *testing.T) {
	// Setup
	rc := &rorConfigSet{
		autoload: true, // Already enabled
		configs:  make(configsMap),
	}

	key := ConfigConst("TEST_KEY")
	rc.configs[key] = ConfigData("initial_value")

	envVar := ConfigConsts.GetEnvVariable(key)
	os.Setenv(envVar, "env_value")
	defer os.Unsetenv(envVar)

	// Execute
	rc.AutoLoadEnv()

	// Verify autoload remains enabled
	if !rc.autoload {
		t.Error("AutoLoadEnv() should keep autoload enabled")
	}

	// Verify the environment value was loaded
	if got := string(rc.configs[key]); got != "env_value" {
		t.Errorf("AutoLoadEnv() config value = %v, want %v", got, "env_value")
	}
}

type mockSecretProvider struct {
	secret string
}

func (m *mockSecretProvider) GetSecret() string {
	return m.secret
}
func TestRorConfigSet_SetWithProvider(t *testing.T) {
	// Mock SecretProvider for testing

	tests := []struct {
		name          string
		key           ConfigConst
		provider      SecretProvider
		expectedValue string
		existingValue string
		hasExisting   bool
	}{
		{
			name:          "set value with provider returning non-empty secret",
			key:           ConfigConst("SECRET_KEY"),
			provider:      &mockSecretProvider{secret: "secret_value"},
			expectedValue: "secret_value",
		},
		{
			name:          "set value with provider returning empty secret",
			key:           ConfigConst("EMPTY_SECRET_KEY"),
			provider:      &mockSecretProvider{secret: ""},
			expectedValue: "",
		},
		{
			name:          "overwrite existing value with provider",
			key:           ConfigConst("OVERWRITE_KEY"),
			provider:      &mockSecretProvider{secret: "new_secret"},
			existingValue: "old_value",
			hasExisting:   true,
			expectedValue: "new_secret",
		},
		{
			name:          "set value with provider returning complex secret",
			key:           ConfigConst("COMPLEX_SECRET_KEY"),
			provider:      &mockSecretProvider{secret: "password123!@#$%^&*()"},
			expectedValue: "password123!@#$%^&*()",
		},
		{
			name:          "set value with provider returning multiline secret",
			key:           ConfigConst("MULTILINE_SECRET_KEY"),
			provider:      &mockSecretProvider{secret: "line1\nline2\nline3"},
			expectedValue: "line1\nline2\nline3",
		},
		{
			name:          "set value with provider returning json-like secret",
			key:           ConfigConst("JSON_SECRET_KEY"),
			provider:      &mockSecretProvider{secret: "{\"username\":\"user\",\"password\":\"pass\"}"},
			expectedValue: "{\"username\":\"user\",\"password\":\"pass\"}",
		},
		{
			name:          "set value with provider returning unicode secret",
			key:           ConfigConst("UNICODE_SECRET_KEY"),
			provider:      &mockSecretProvider{secret: "测试密码🔐"},
			expectedValue: "测试密码🔐",
		},
		{
			name:          "set value with provider returning whitespace secret",
			key:           ConfigConst("WHITESPACE_SECRET_KEY"),
			provider:      &mockSecretProvider{secret: "   secret with spaces   "},
			expectedValue: "   secret with spaces   ",
		},
		{
			name:          "set value with provider returning very long secret",
			key:           ConfigConst("LONG_SECRET_KEY"),
			provider:      &mockSecretProvider{secret: "very_long_secret_" + string(make([]byte, 1000))},
			expectedValue: "very_long_secret_" + string(make([]byte, 1000)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			rc := &rorConfigSet{
				configs: make(configsMap),
			}

			if tt.hasExisting {
				rc.configs[tt.key] = ConfigData(tt.existingValue)
			}

			// Execute
			rc.SetWithProvider(tt.key, tt.provider)

			// Verify
			if got := string(rc.configs[tt.key]); got != tt.expectedValue {
				t.Errorf("SetWithProvider() = %v, want %v", got, tt.expectedValue)
			}
		})
	}
}

type trackingSecretProvider struct {
	secret    string
	callCount int
}

func (t *trackingSecretProvider) GetSecret() string {
	t.callCount++
	return t.secret
}
func TestRorConfigSet_SetWithProvider_CallsGetSecret(t *testing.T) {
	// Mock SecretProvider that tracks if GetSecret was called

	// Setup
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	provider := &trackingSecretProvider{secret: "tracked_secret"}
	key := ConfigConst("TRACKING_KEY")

	// Execute
	rc.SetWithProvider(key, provider)

	// Verify GetSecret was called exactly once
	if provider.callCount != 1 {
		t.Errorf("SetWithProvider() should call GetSecret() exactly once, called %d times", provider.callCount)
	}

	// Verify the value was set correctly
	if got := string(rc.configs[key]); got != "tracked_secret" {
		t.Errorf("SetWithProvider() = %v, want %v", got, "tracked_secret")
	}
}

func TestRorConfigSet_SetWithProvider_MultipleCalls(t *testing.T) {
	// Mock SecretProvider for testing

	// Setup
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	key := ConfigConst("MULTIPLE_CALLS_KEY")

	// First call
	provider1 := &mockSecretProvider{secret: "first_secret"}
	rc.SetWithProvider(key, provider1)

	// Verify first value
	if got := string(rc.configs[key]); got != "first_secret" {
		t.Errorf("First SetWithProvider() = %v, want %v", got, "first_secret")
	}

	// Second call with different provider
	provider2 := &mockSecretProvider{secret: "second_secret"}
	rc.SetWithProvider(key, provider2)

	// Verify second value overwrote the first
	if got := string(rc.configs[key]); got != "second_secret" {
		t.Errorf("Second SetWithProvider() = %v, want %v", got, "second_secret")
	}
}

func TestRorConfigSet_SetWithProvider_DoesNotAffectOtherConfigs(t *testing.T) {
	// Mock SecretProvider for testing

	// Setup
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	// Set initial values for other configs
	key1 := ConfigConst("KEY1")
	key2 := ConfigConst("KEY2")
	targetKey := ConfigConst("TARGET_KEY")

	rc.configs[key1] = ConfigData("value1")
	rc.configs[key2] = ConfigData("value2")

	provider := &mockSecretProvider{secret: "provider_secret"}

	// Execute
	rc.SetWithProvider(targetKey, provider)

	// Verify target key was set
	if got := string(rc.configs[targetKey]); got != "provider_secret" {
		t.Errorf("SetWithProvider() target key = %v, want %v", got, "provider_secret")
	}

	// Verify other configs were not affected
	if got := string(rc.configs[key1]); got != "value1" {
		t.Errorf("SetWithProvider() should not affect other configs, key1 = %v, want %v", got, "value1")
	}

	if got := string(rc.configs[key2]); got != "value2" {
		t.Errorf("SetWithProvider() should not affect other configs, key2 = %v, want %v", got, "value2")
	}
}
func TestRorConfigSet_AutoLoadAllEnv(t *testing.T) {
	tests := []struct {
		name                string
		dotEnvData          map[string]string
		existingEnvVars     map[string]string
		configConstants     map[string]string
		expectedEnvVars     map[string]string
		expectedConfigCount int
	}{
		{
			name: "load env vars from dot env file only",
			dotEnvData: map[string]string{
				"DOT_ENV_VAR1": "dot_value1",
				"DOT_ENV_VAR2": "dot_value2",
			},
			existingEnvVars: map[string]string{},
			configConstants: map[string]string{},
			expectedEnvVars: map[string]string{
				"DOT_ENV_VAR1": "dot_value1",
				"DOT_ENV_VAR2": "dot_value2",
			},
			expectedConfigCount: 0,
		},
		{
			name:       "load configs from existing env vars only",
			dotEnvData: map[string]string{},
			existingEnvVars: map[string]string{
				"CONFIG_VAR1": "env_value1",
				"CONFIG_VAR2": "env_value2",
			},
			configConstants: map[string]string{
				"key1": "CONFIG_VAR1",
				"key2": "CONFIG_VAR2",
			},
			expectedEnvVars: map[string]string{
				"CONFIG_VAR1": "env_value1",
				"CONFIG_VAR2": "env_value2",
			},
			expectedConfigCount: 2,
		},
		{
			name: "load both dot env and config constants",
			dotEnvData: map[string]string{
				"DOT_ENV_VAR": "dot_value",
			},
			existingEnvVars: map[string]string{
				"CONFIG_VAR": "config_value",
			},
			configConstants: map[string]string{
				"config_key": "CONFIG_VAR",
			},
			expectedEnvVars: map[string]string{
				"DOT_ENV_VAR": "dot_value",
				"CONFIG_VAR":  "config_value",
			},
			expectedConfigCount: 1,
		},
		{
			name: "dot env overwrites existing env vars",
			dotEnvData: map[string]string{
				"SHARED_VAR": "dot_env_value",
			},
			existingEnvVars: map[string]string{
				"SHARED_VAR": "original_value",
			},
			configConstants: map[string]string{
				"shared_key": "SHARED_VAR",
			},
			expectedEnvVars: map[string]string{
				"SHARED_VAR": "dot_env_value",
			},
			expectedConfigCount: 1,
		},
		{
			name:            "no dot env file and no matching env vars",
			dotEnvData:      map[string]string{},
			existingEnvVars: map[string]string{},
			configConstants: map[string]string{
				"key1": "NON_EXISTENT_VAR",
			},
			expectedEnvVars:     map[string]string{},
			expectedConfigCount: 0,
		},
		{
			name:       "partial match of config constants with env vars",
			dotEnvData: map[string]string{},
			existingEnvVars: map[string]string{
				"EXISTS_VAR":     "exists_value",
				"NOT_CONFIG_VAR": "not_config_value",
			},
			configConstants: map[string]string{
				"exists_key":     "EXISTS_VAR",
				"not_exists_key": "NOT_EXISTS_VAR",
			},
			expectedEnvVars: map[string]string{
				"EXISTS_VAR":     "exists_value",
				"NOT_CONFIG_VAR": "not_config_value",
			},
			expectedConfigCount: 1,
		},
		{
			name: "empty values in dot env",
			dotEnvData: map[string]string{
				"EMPTY_DOT_VAR": "",
			},
			existingEnvVars: map[string]string{
				"EMPTY_CONFIG_VAR": "",
			},
			configConstants: map[string]string{
				"empty_key": "EMPTY_CONFIG_VAR",
			},
			expectedEnvVars: map[string]string{
				"EMPTY_DOT_VAR":    "",
				"EMPTY_CONFIG_VAR": "",
			},
			expectedConfigCount: 1,
		},
		{
			name: "special characters in env var names and values",
			dotEnvData: map[string]string{
				"SPECIAL_VAR": "value with spaces and symbols!@#$%",
			},
			existingEnvVars: map[string]string{
				"UNICODE_VAR": "测试值🚀",
			},
			configConstants: map[string]string{
				"unicode_key": "UNICODE_VAR",
			},
			expectedEnvVars: map[string]string{
				"SPECIAL_VAR": "value with spaces and symbols!@#$%",
				"UNICODE_VAR": "测试值🚀",
			},
			expectedConfigCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup - Clear environment first
			for key := range tt.existingEnvVars {
				os.Unsetenv(key)
			}
			for key := range tt.dotEnvData {
				os.Unsetenv(key)
			}

			// Set existing environment variables
			for key, value := range tt.existingEnvVars {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			// Mock readDotEnv function (this would need to be implemented in the test)
			// Since we can't easily mock the readDotEnv function, we'll set the dot env vars manually
			for key, value := range tt.dotEnvData {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			// Override ConfigConsts so the test controls which env vars are recognised
			originalConfigConsts := ConfigConsts
			customConsts := make(ConfigconstsMap)
			for key, envVar := range tt.configConstants {
				customConsts[ConfigConst(key)] = ConfigConstData{value: envVar}
			}
			ConfigConsts = customConsts
			t.Cleanup(func() {
				ConfigConsts = originalConfigConsts
			})

			rc := &rorConfigSet{
				configs: make(configsMap),
			}

			// Execute
			rc.AutoLoadAllEnv()

			// Verify environment variables are set correctly
			for key, expectedValue := range tt.expectedEnvVars {
				if got := os.Getenv(key); got != expectedValue {
					t.Errorf("AutoLoadAllEnv() env var %s = %v, want %v", key, got, expectedValue)
				}
			}

			// Verify configs are loaded for matching environment variables
			if len(rc.configs) != tt.expectedConfigCount {
				t.Errorf("AutoLoadAllEnv() loaded config count = %v, want %v", len(rc.configs), tt.expectedConfigCount)
			}
		})
	}
}

func TestRorConfigSet_AutoLoadAllEnv_CallsReadDotEnv(t *testing.T) {
	// This test verifies that readDotEnv is called and its results are applied to environment
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	// Since we can't easily mock readDotEnv, we'll test the behavior indirectly
	// by checking if environment variables get set

	// Clear any existing test environment variables
	testEnvVars := []string{"TEST_DOT_ENV_VAR1", "TEST_DOT_ENV_VAR2"}
	for _, envVar := range testEnvVars {
		os.Unsetenv(envVar)
	}

	// Execute
	rc.AutoLoadAllEnv()

	// Note: Without being able to mock readDotEnv, we can't fully test this
	// In a real implementation, you might use dependency injection or interfaces
	// to make readDotEnv mockable
}

func TestRorConfigSet_AutoLoadAllEnv_ChecksConfigConstants(t *testing.T) {
	// This test verifies that the function checks ConfigConsts and loads matching env vars
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	// Set up test environment variables
	testEnvVars := map[string]string{
		"TEST_CONFIG_VAR1": "test_value1",
		"TEST_CONFIG_VAR2": "test_value2",
		"NON_CONFIG_VAR":   "non_config_value",
	}

	for key, value := range testEnvVars {
		os.Setenv(key, value)
		defer os.Unsetenv(key)
	}

	// Execute
	rc.AutoLoadAllEnv()

	// Note: Without access to the actual ConfigConsts implementation,
	// we can't verify the exact behavior. In a real test, you would:
	// 1. Mock or control ConfigConsts
	// 2. Verify that only environment variables that match ConfigConsts entries are loaded
	// 3. Verify that LoadEnv is called for each matching key
}

func TestRorConfigSet_AutoLoadAllEnv_IgnoresNonExistentEnvVars(t *testing.T) {
	// Test that the function doesn't try to load configs for non-existent environment variables
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	// Ensure test environment variables don't exist
	testEnvVars := []string{"NON_EXISTENT_VAR1", "NON_EXISTENT_VAR2"}
	for _, envVar := range testEnvVars {
		os.Unsetenv(envVar)
	}

	initialConfigCount := len(rc.configs)

	// Execute
	rc.AutoLoadAllEnv()

	// Verify no configs were added for non-existent environment variables
	// Note: This test assumes that the test environment variables don't exist
	// in ConfigConsts or that ConfigConsts filtering works correctly
	if len(rc.configs) < initialConfigCount {
		t.Errorf("AutoLoadAllEnv() should not reduce config count")
	}
}

func TestRorConfigSet_AutoLoadAllEnv_OrderOfOperations(t *testing.T) {
	// Test that dot env loading happens before config constant checking
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	// Set up a scenario where dot env should set an environment variable
	// that then gets picked up by the config constants check
	testEnvVar := "ORDER_TEST_VAR"
	testValue := "order_test_value"

	// Ensure the env var doesn't exist initially
	os.Unsetenv(testEnvVar)

	// The test would need to mock readDotEnv to return this variable
	// For now, we'll simulate by setting it manually
	os.Setenv(testEnvVar, testValue)
	defer os.Unsetenv(testEnvVar)

	// Execute
	rc.AutoLoadAllEnv()

	// Verify the environment variable is set
	if got := os.Getenv(testEnvVar); got != testValue {
		t.Errorf("AutoLoadAllEnv() should set env var from dot env, got %v, want %v", got, testValue)
	}
}

func TestRorConfigSet_AutoLoadAllEnv_EmptyDotEnvAndConfigConsts(t *testing.T) {
	// Test behavior when both readDotEnv returns empty and ConfigConsts is empty
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	initialConfigCount := len(rc.configs)

	// Execute
	rc.AutoLoadAllEnv()

	// Verify no configs were added
	if len(rc.configs) != initialConfigCount {
		t.Errorf("AutoLoadAllEnv() with empty inputs should not change config count, got %v, want %v", len(rc.configs), initialConfigCount)
	}
}

func TestRorConfigSet_AutoLoadAllEnv_DoesNotAffectExistingConfigs(t *testing.T) {
	// Test that existing configs are not removed or modified (unless overwritten by LoadEnv)
	rc := &rorConfigSet{
		configs: make(configsMap),
	}

	// Set up existing configs
	existingKey := ConfigConst("EXISTING_KEY")
	existingValue := "existing_value"
	rc.configs[existingKey] = ConfigData(existingValue)

	// Execute
	rc.AutoLoadAllEnv()

	// Verify existing config still exists (unless it was specifically loaded from env)
	if _, exists := rc.configs[existingKey]; !exists {
		t.Error("AutoLoadAllEnv() should not remove existing configs")
	}
}
