package mocktransport_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport"
)

func TestMockTransportImplementsInterface(t *testing.T) {
	// Test that the mock transport can be instantiated
	transport := mocktransport.NewRorMockTransport()

	if transport == nil {
		t.Fatal("Expected transport to be non-nil")
	}

	// Test basic functionality
	err := transport.Ping()
	if err != nil {
		t.Errorf("Expected ping to succeed, got error: %v", err)
	}

	// Test status
	status := transport.Status()
	if !status.IsEstablished() {
		t.Error("Expected status to be established")
	}

	// Test API secret
	secret := transport.GetApiSecret()
	if secret == "" {
		t.Error("Expected API secret to be set")
	}

	// Test role
	role := transport.GetRole()
	if role == "" {
		t.Error("Expected role to be set")
	}
}

func TestMockTransportClients(t *testing.T) {
	transport := mocktransport.NewRorMockTransport()

	// Test that all client interfaces are available
	if transport.Info() == nil {
		t.Error("Expected Info client to be available")
	}

	if transport.Clusters() == nil {
		t.Error("Expected Clusters client to be available")
	}

	if transport.Datacenters() == nil {
		t.Error("Expected Datacenters client to be available")
	}

	if transport.Workspaces() == nil {
		t.Error("Expected Workspaces client to be available")
	}

	if transport.Projects() == nil {
		t.Error("Expected Projects client to be available")
	}

	if transport.Resources() == nil {
		t.Error("Expected Resources client to be available")
	}

	if transport.ResourcesV2() == nil {
		t.Error("Expected ResourcesV2 client to be available")
	}

	if transport.Metrics() == nil {
		t.Error("Expected Metrics client to be available")
	}

	if transport.Stream() == nil {
		t.Error("Expected Stream client to be available")
	}

	if transport.Streamv2() == nil {
		t.Error("Expected Streamv2 client to be available")
	}

	if transport.AclV1() == nil {
		t.Error("Expected AclV1 client to be available")
	}

	if transport.Self() == nil {
		t.Error("Expected Self client to be available")
	}
}

func TestMockTransportBasicOperations(t *testing.T) {
	transport := mocktransport.NewRorMockTransport()

	// Test Info client
	version, err := transport.Info().GetVersion()
	if err != nil {
		t.Errorf("Expected version call to succeed, got error: %v", err)
	}
	if version == "" {
		t.Error("Expected version to be non-empty")
	}

	// Test Clusters client
	clusterSelf, err := transport.Clusters().GetSelf()
	if err != nil {
		t.Errorf("Expected cluster self call to succeed, got error: %v", err)
	}
	if clusterSelf.ClusterId == "" {
		t.Error("Expected cluster ID to be non-empty")
	}

	// Test basic CRUD operations
	clusters, err := transport.Clusters().GetAll()
	if err != nil {
		t.Errorf("Expected get all clusters to succeed, got error: %v", err)
	}
	if clusters == nil || len(*clusters) == 0 {
		t.Error("Expected clusters to be returned")
	}
}
