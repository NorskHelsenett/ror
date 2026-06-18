package redisdb

import (
	"context"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
)

// TestPingNilClient verifies that pinging an unconnected connection returns false
// instead of panicking on a nil client. The health checker is registered before
// connect succeeds, so it must tolerate a nil client.
func TestPingNilClient(t *testing.T) {
	rc := rediscon{}

	if rc.Ping() {
		t.Error("expected Ping to return false for a nil client")
	}
	if rc.PingWithContext(context.Background()) {
		t.Error("expected PingWithContext to return false for a nil client")
	}
}

// TestRedisStartupCheckerNotConnected verifies that the startup checker reports a
// failing status with a clear message while the connection is still being
// established, so the health endpoint surfaces redis as the blocking dependency.
func TestRedisStartupCheckerNotConnected(t *testing.T) {
	checker := &redisStartupChecker{conn: &rediscon{}}

	checks := checker.CheckHealth(context.Background())

	if len(checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(checks))
	}
	if checks[0].Status != rorhealth.StatusFail {
		t.Errorf("expected StatusFail while connecting, got %v", checks[0].Status)
	}
	if checks[0].Output != "Connecting to redis" {
		t.Errorf("unexpected output: %q", checks[0].Output)
	}
}

// TestRedisStartupCheckerConnectedDelegates verifies that once the checker is
// marked connected it delegates to the live connection check instead of
// reporting the "Connecting" placeholder. With no real server the underlying
// connection reports not-connected, but it must not return the startup
// placeholder.
func TestRedisStartupCheckerConnectedDelegates(t *testing.T) {
	checker := &redisStartupChecker{conn: &rediscon{}}
	checker.setConnected()

	checks := checker.CheckHealth(context.Background())

	if len(checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(checks))
	}
	if checks[0].Output == "Connecting to redis" {
		t.Error("connected checker should delegate to the live check, not report the startup placeholder")
	}
}
