package rabbitmqclient

import (
	"context"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
)

// TestRabbitmqStartupCheckerNotConnected verifies that the startup checker
// reports a failing status with a clear message while the connection is still
// being established, so the health endpoint surfaces rabbitmq as the blocking
// dependency.
func TestRabbitmqStartupCheckerNotConnected(t *testing.T) {
	checker := &rabbitmqStartupChecker{conn: &rabbitmqcon{}}

	checks := checker.CheckHealth(context.Background())

	if len(checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(checks))
	}
	if checks[0].Status != rorhealth.StatusFail {
		t.Errorf("expected StatusFail while connecting, got %v", checks[0].Status)
	}
	if checks[0].Output != "Connecting to rabbitmq" {
		t.Errorf("unexpected output: %q", checks[0].Output)
	}
}

// TestRabbitmqStartupCheckerConnectedDelegates verifies that once the checker is
// marked connected it delegates to the live connection check instead of
// reporting the "Connecting" placeholder. With no real broker the underlying
// connection reports not-connected, but it must not return the startup
// placeholder.
func TestRabbitmqStartupCheckerConnectedDelegates(t *testing.T) {
	checker := &rabbitmqStartupChecker{conn: &rabbitmqcon{}}
	checker.setConnected()

	checks := checker.CheckHealth(context.Background())

	if len(checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(checks))
	}
	if checks[0].Output == "Connecting to rabbitmq" {
		t.Error("connected checker should delegate to the live check, not report the startup placeholder")
	}
}

// TestValidateConfig verifies the configuration validation used before a
// connection attempt produces clear, actionable errors.
func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		conn    *rabbitmqcon
		wantErr bool
	}{
		{
			name:    "missing host",
			conn:    &rabbitmqcon{Port: "5672", Credentials: staticCreds{}},
			wantErr: true,
		},
		{
			name:    "missing port",
			conn:    &rabbitmqcon{Host: "localhost", Credentials: staticCreds{}},
			wantErr: true,
		},
		{
			name:    "missing credentials",
			conn:    &rabbitmqcon{Host: "localhost", Port: "5672"},
			wantErr: true,
		},
		{
			name:    "valid",
			conn:    &rabbitmqcon{Host: "localhost", Port: "5672", Credentials: staticCreds{}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.conn.validateConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// staticCreds is a minimal credshelper.CredHelper for tests.
type staticCreds struct{}

func (staticCreds) GetUsername() string              { return "user" }
func (staticCreds) GetPassword() string              { return "pass" }
func (staticCreds) GetCredentials() (string, string) { return "user", "pass" }
