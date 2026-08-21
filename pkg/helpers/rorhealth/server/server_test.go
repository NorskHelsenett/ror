package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/netip"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
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

// failingChecker is a health checker that always reports a failing status.
type failingChecker struct{}

func (failingChecker) CheckHealth(_ context.Context) []rorhealth.Check {
	return []rorhealth.Check{{
		Status: rorhealth.StatusFail,
		Output: "dependency down",
	}}
}

// TestHealthMuxLivenessIgnoresDependencies verifies that the liveness endpoint
// reports the process as alive (HTTP 200) even when a registered dependency is
// failing, so that a dependency outage does not trigger a pod restart.
func TestHealthMuxLivenessIgnoresDependencies(t *testing.T) {
	registered := rorhealth.Register(context.Background(), "test-failing-dep", failingChecker{})
	defer registered.Deregister()

	mux := newHealthMux()

	req := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("liveness with failing dependency = %d, want %d", rec.Code, http.StatusOK)
	}
}

// TestHealthMuxReadinessReflectsDependencies verifies that the readiness
// endpoint fails (HTTP 500) when a registered dependency is failing, so the pod
// is taken out of service while remaining alive.
func TestHealthMuxReadinessReflectsDependencies(t *testing.T) {
	registered := rorhealth.Register(context.Background(), "test-failing-dep", failingChecker{})
	defer registered.Deregister()

	mux := newHealthMux()

	req := httptest.NewRequest(http.MethodGet, "/health/ready", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("readiness with failing dependency = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
}

// TestHealthMuxLivenessMethods verifies the liveness handler's HTTP method
// handling.
func TestHealthMuxLivenessMethods(t *testing.T) {
	mux := newHealthMux()

	tests := []struct {
		method   string
		expected int
	}{
		{http.MethodGet, http.StatusOK},
		{http.MethodHead, http.StatusOK},
		{http.MethodOptions, http.StatusNoContent},
		{http.MethodPost, http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/health/live", nil)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			if rec.Code != tt.expected {
				t.Errorf("liveness %s = %d, want %d", tt.method, rec.Code, tt.expected)
			}
		})
	}
}

// passingChecker is a health checker that always reports a passing status.
type passingChecker struct{}

func (passingChecker) CheckHealth(_ context.Context) []rorhealth.Check {
	return []rorhealth.Check{{Status: rorhealth.StatusPass}}
}

// newTestGate builds a gate with deterministic (zero) jitter for tests.
func newTestGate(grace time.Duration, critical ...string) *livenessGate {
	set := make(map[string]struct{}, len(critical))
	for _, c := range critical {
		set[c] = struct{}{}
	}
	return &livenessGate{
		grace:        grace,
		jitter:       0,
		critical:     set,
		failingSince: make(map[string]time.Time),
	}
}

// TestLivenessGateTripsAfterGrace verifies the gate trips once a critical
// dependency has been failing for longer than the grace.
func TestLivenessGateTripsAfterGrace(t *testing.T) {
	registered := rorhealth.Register(context.Background(), "gate-failing-dep", failingChecker{})
	defer registered.Deregister()

	g := newTestGate(20 * time.Millisecond)
	g.sample(context.Background())

	if _, _, tripped := g.tripped(); tripped {
		t.Fatal("gate tripped immediately, want not tripped before grace elapses")
	}

	time.Sleep(40 * time.Millisecond)

	name, _, tripped := g.tripped()
	if !tripped {
		t.Fatal("gate did not trip after grace elapsed")
	}
	if name != "gate-failing-dep" {
		t.Errorf("tripped check = %q, want %q", name, "gate-failing-dep")
	}
}

// TestLivenessGateIgnoresNonCritical verifies a failing check that is not in the
// critical allowlist never trips liveness.
func TestLivenessGateIgnoresNonCritical(t *testing.T) {
	registered := rorhealth.Register(context.Background(), "gate-noncritical-dep", failingChecker{})
	defer registered.Deregister()

	g := newTestGate(10*time.Millisecond, "mongodb")
	g.sample(context.Background())
	time.Sleep(30 * time.Millisecond)

	if _, _, tripped := g.tripped(); tripped {
		t.Fatal("gate tripped on a non-critical dependency")
	}
}

// TestLivenessGatePassingDependencyDoesNotTrip verifies a healthy dependency
// never trips liveness.
func TestLivenessGatePassingDependencyDoesNotTrip(t *testing.T) {
	registered := rorhealth.Register(context.Background(), "gate-passing-dep", passingChecker{})
	defer registered.Deregister()

	g := newTestGate(10 * time.Millisecond)
	g.sample(context.Background())
	time.Sleep(30 * time.Millisecond)

	if _, _, tripped := g.tripped(); tripped {
		t.Fatal("gate tripped on a passing dependency")
	}
}

// TestLivenessGateRecovers verifies the failing timer is cleared once the
// dependency recovers, so a recovered dependency no longer trips liveness.
func TestLivenessGateRecovers(t *testing.T) {
	registered := rorhealth.Register(context.Background(), "gate-flapping-dep", failingChecker{})

	g := newTestGate(10 * time.Millisecond)
	g.sample(context.Background())
	time.Sleep(30 * time.Millisecond)
	if _, _, tripped := g.tripped(); !tripped {
		t.Fatal("gate did not trip while dependency failing")
	}

	// Dependency recovers: re-register a passing checker under the same name.
	registered.Deregister()
	recovered := rorhealth.Register(context.Background(), "gate-flapping-dep", passingChecker{})
	defer recovered.Deregister()

	g.sample(context.Background())
	if _, _, tripped := g.tripped(); tripped {
		t.Fatal("gate stayed tripped after dependency recovered")
	}
}

// TestHealthMuxLivenessEscalates verifies the liveness endpoint returns 500 once
// the gate has tripped for a chronically failing critical dependency.
func TestHealthMuxLivenessEscalates(t *testing.T) {
	registered := rorhealth.Register(context.Background(), "mux-failing-dep", failingChecker{})
	defer registered.Deregister()

	gate = newTestGate(20 * time.Millisecond)
	defer func() { gate = nil }()

	gate.sample(context.Background())
	time.Sleep(40 * time.Millisecond)

	mux := newHealthMux()
	req := httptest.NewRequest(http.MethodGet, "/health/live", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("liveness after escalation = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
}

// TestParseGraceDuration verifies grace parsing accepts durations and bare
// seconds and rejects invalid input.
func TestParseGraceDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{"", 0},
		{"5m", 5 * time.Minute},
		{"30s", 30 * time.Second},
		{"90", 90 * time.Second},
		{"garbage", 0},
	}
	for _, tt := range tests {
		if got := parseGraceDuration(tt.input); got != tt.expected {
			t.Errorf("parseGraceDuration(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}
