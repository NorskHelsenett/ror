package mongodb

import (
	"context"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
)

// TestPingNilClient verifies that pinging a connection whose client has not been
// established returns false instead of panicking with a nil-pointer
// dereference. This guard is what makes it safe to register the health check
// before the connection is established.
func TestPingNilClient(t *testing.T) {
	c := MongodbCon{} // Client is nil
	if c.ping(context.Background()) {
		t.Error("ping on a nil client should return false")
	}
}

// TestMongoStartupCheckerNotConnected verifies that the startup checker reports
// a failing status with a clear message while the connection is still being
// established, so the health endpoint surfaces mongodb as the blocking
// dependency.
func TestMongoStartupCheckerNotConnected(t *testing.T) {
	checker := &mongoStartupChecker{}

	checks := checker.CheckHealth(context.Background())

	if len(checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(checks))
	}
	if checks[0].Status != rorhealth.StatusFail {
		t.Errorf("expected StatusFail while connecting, got %v", checks[0].Status)
	}
	if checks[0].Output != "Connecting to mongodb" {
		t.Errorf("unexpected output: %q", checks[0].Output)
	}
}

// TestMongoStartupCheckerConnectedDelegates verifies that once the checker is
// marked connected it delegates to the live connection check instead of
// reporting the "Connecting" placeholder. With no real mongodb available the
// live check fails, but crucially it must not return the startup placeholder.
func TestMongoStartupCheckerConnectedDelegates(t *testing.T) {
	checker := &mongoStartupChecker{}
	checker.setConnected()

	checks := checker.CheckHealth(context.Background())

	if len(checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(checks))
	}
	if checks[0].Output == "Connecting to mongodb" {
		t.Error("connected checker should delegate to the live check, not report the startup placeholder")
	}
}
