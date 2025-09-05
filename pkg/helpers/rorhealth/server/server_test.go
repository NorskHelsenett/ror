package server

import (
	"net/netip"
	"testing"
)

func TestParseServerString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "ip and port",
			input:    "192.168.1.1:8080",
			expected: "192.168.1.1:8080",
		},
		{
			name:     "only port with colon prefix",
			input:    ":8080",
			expected: "0.0.0.0:8080",
		},
		{
			name:     "only ip with colon suffix",
			input:    "192.168.1.1:",
			expected: "192.168.1.1:9999",
		},
		{
			name:     "only port number",
			input:    "8080",
			expected: "0.0.0.0:8080",
		},
		{
			name:     "only ip address",
			input:    "192.168.1.1",
			expected: "192.168.1.1:9999",
		},
		{
			name:     "invalid format with multiple colons",
			input:    "192.168.1.1:8080:extra",
			expected: "0.0.0.0:9999",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "0.0.0.0:9999",
		},
		{
			name:     "just a colon",
			input:    ":",
			expected: "0.0.0.0:9999",
		},
		{
			name:     "hostname",
			input:    "localhost",
			expected: "localhost:9999",
		},
		{
			name:     "hostname with port",
			input:    "localhost:3000",
			expected: "localhost:3000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseServerString(tt.input)
			if result != tt.expected {
				t.Errorf("parseServerString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
func TestGetDefaultAddrPort(t *testing.T) {
	result := getDefaultAddrPort()

	// Check that the result is valid
	if !result.IsValid() {
		t.Error("getDefaultAddrPort() returned invalid AddrPort")
	}

	// Check the IP address
	expectedIP := netip.MustParseAddr("0.0.0.0")
	if result.Addr() != expectedIP {
		t.Errorf("getDefaultAddrPort() IP = %v, want %v", result.Addr(), expectedIP)
	}

	// Check the port
	expectedPort := uint16(9999)
	if result.Port() != expectedPort {
		t.Errorf("getDefaultAddrPort() port = %v, want %v", result.Port(), expectedPort)
	}

	// Check the string representation
	expectedString := "0.0.0.0:9999"
	if result.String() != expectedString {
		t.Errorf("getDefaultAddrPort() string = %v, want %v", result.String(), expectedString)
	}
}
