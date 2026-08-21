package mongodb

import (
	"context"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

// TestGetMongoDbUsesReceiverClient verifies that a standalone MongodbCon (its
// own client, no Credentials, singleton uninitialized) uses the receiver's
// client instead of the package singleton. This guards against the regression
// where GetMongoDb hit the uninitialized singleton and dereferenced its nil
// Credentials, panicking for callers that construct their own connection (e.g.
// the resourcesv2service integration tests).
func TestGetMongoDbUsesReceiverClient(t *testing.T) {
	// mongo.Connect is lazy in the v2 driver, so this does not dial anything.
	cli, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	t.Cleanup(func() { _ = cli.Disconnect(context.Background()) })

	rc := MongodbCon{Client: cli, Database: "standalonedb"}
	db := rc.GetMongoDb()
	if db == nil {
		t.Fatal("GetMongoDb returned nil")
	}
	if db.Name() != "standalonedb" {
		t.Errorf("db name = %q, want standalonedb", db.Name())
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
