package rorconfig

import (
	"testing"
)

func TestConfigconstsMap_GetConfigConstByName(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		expected ConfigConst
	}{
		{
			name:     "existing key - ROLE",
			key:      "ROLE",
			expected: ROLE,
		},
		{
			name:     "existing key with different value - ROR_URL",
			key:      "ROR_URL",
			expected: ConfigConst("API_ENDPOINT"),
		},
		{
			name:     "non-existing key",
			key:      "NON_EXISTING_KEY",
			expected: ConfigConst("NON_EXISTING_KEY"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConfigConsts.GetConfigConstByName(tt.key)
			if result != tt.expected {
				t.Errorf("GetConfigConstByName(%q) = %q, want %q", tt.key, result, tt.expected)
			}
		})
	}
}
func TestConfigconstsMap_GetDescription(t *testing.T) {
	tests := []struct {
		name     string
		key      ConfigConst
		expected string
	}{
		{
			name:     "existing key with empty description - ROLE",
			key:      ROLE,
			expected: "",
		},
		{
			name:     "existing key with empty description - HTTP_HOST",
			key:      HTTP_HOST,
			expected: "",
		},
		{
			name:     "deprecated key with description - HEALTH_ENDPOINT",
			key:      HEALTH_ENDPOINT,
			expected: "use HTTP_HEALTH_HOST / HTTP_HEALTH_PORT instead",
		},
		{
			name:     "non-existing key",
			key:      ConfigConst("NON_EXISTING_KEY"),
			expected: "Local env variable not in central list",
		},
		{
			name:     "empty key",
			key:      ConfigConst(""),
			expected: "Local env variable not in central list",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConfigConsts.GetDescription(tt.key)
			if result != tt.expected {
				t.Errorf("GetDescription(%q) = %q, want %q", tt.key, result, tt.expected)
			}
		})
	}
}
