# Mock Transport Implementation

This document describes the complete mock transport implementation for the ROR client.

## Overview

The mock transport (`RorMockTransport`) is a complete implementation of the `RorTransport` interface that provides mock responses for all API endpoints. This is useful for testing, development, and integration scenarios where you don't want to connect to a real ROR API server.

## Architecture

The mock transport follows the same structure as the REST transport, with separate packages for each client:

```
mocktransport/
├── transport.go                     # Main transport implementation
├── transport_test.go                # Tests
├── mocktransportinfo/               # Info client mock
├── mocktransportclusters/           # Clusters client mock
├── mocktransportdatacenter/         # Datacenter client mock
├── mocktransportworkspaces/         # Workspaces client mock
├── mocktransportprojects/           # Projects client mock
├── mocktransportresources/          # V1 Resources client mock
├── mocktransportresourcesv2/        # V2 Resources client mock
├── mocktransportmetrics/            # Metrics client mock
├── mocktransportstream/             # V1 Stream client mock
├── mocktransportstreamv2/           # V2 Stream client mock
├── mocktransportacl/                # ACL client mock
├── mocktransportself/               # Self client mock
└── mocktransportstatus/             # Transport status mock
```

## Implementation

### Main Transport (transport.go)

The `RorMockTransport` struct implements all methods required by the `RorTransport` interface:

-   `Status()` - Returns mock transport status
-   `Stream()` - Returns V1 stream client mock
-   `Info()` - Returns info client mock
-   `Datacenters()` - Returns datacenter client mock
-   `Clusters()` - Returns clusters client mock
-   `Self()` - Returns self client mock
-   `Workspaces()` - Returns workspaces client mock
-   `Projects()` - Returns projects client mock
-   `Resources()` - Returns V1 resources client mock
-   `ResourcesV2()` - Returns V2 resources client mock
-   `Metrics()` - Returns metrics client mock
-   `Streamv2()` - Returns V2 stream client mock
-   `AclV1()` - Returns ACL client mock
-   `Ping()` - Always returns nil (success)
-   `GetApiSecret()` - Returns configurable mock API secret
-   `GetRole()` - Returns configurable mock role

### Client Implementations

Each client mock provides realistic mock responses:

#### Info Client

-   `GetVersion()` - Returns "1.1.1"

#### Clusters Client

-   `GetSelf()` - Returns mock cluster self info
-   `GetById(id)` - Returns mock cluster with given ID
-   `GetAll()` - Returns list of mock clusters
-   `Create()` - Returns mock created cluster ID
-   `Register()` - Returns mock API key
-   Full CRUD operations with appropriate validation

#### Datacenters Client

-   `Get()` - Returns list of mock datacenters
-   `GetById(id)` - Returns mock datacenter by ID
-   `GetByName(name)` - Returns mock datacenter by name
-   `Post()`, `Put()` - Create/update operations

#### Workspaces Client

-   `GetByName()`, `GetById()` - Retrieve operations
-   `GetAll()` - List all workspaces
-   `GetKubeconfig()` - Returns mock kubeconfig

#### Projects Client

-   `GetById()`, `Get()`, `GetAll()` - Standard retrieval operations

#### Resources Client (V1)

-   All resource CRUD operations for various resource types
-   Hash list operations
-   Vulnerability reports
-   Cluster orders
-   Applications, PVCs, Routes, etc.

#### Resources Client (V2)

-   `Get()`, `GetByUid()` - Resource queries
-   `Update()`, `UpdateOne()` - Resource updates
-   `Delete()` - Resource deletion
-   `Exists()` - Resource existence checks
-   `GetOwnHashes()` - Hash list operations

#### Metrics Client

-   `PostReport()` - Accepts metrics reports
-   `CreatePVC()` - PVC metrics

#### Stream Clients (V1 & V2)

-   `StartEventstream()` - Returns channel with mock events
-   `StartEventstreamWithCallback()` - Callback-based event streaming
-   `BroadcastEvent()` (V2 only) - Event broadcasting

#### ACL Client

-   Full CRUD operations for ACL items
-   Access checking (always returns true in mock)
-   Filtering support

#### Self Client

-   `Get()` - Returns mock self information
-   `CreateOrUpdateApiKey()` - Returns mock API key

## Usage

```go
package main

import (
    "fmt"
    "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport"
)

func main() {
    // Create mock transport
    transport := mocktransport.NewRorMockTransport()

    // Use like any other transport
    version, err := transport.Info().GetVersion()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Version: %s\n", version)

    // Get clusters
    clusters, err := transport.Clusters().GetAll()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Found %d clusters\n", len(*clusters))

    // Test connection
    err = transport.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Connection successful")
}
```

## Configuration

The mock transport supports some configuration:

```go
transport := mocktransport.NewRorMockTransport()

// Set custom API secret for testing
transport.SetApiSecret("custom-test-secret")

// Set custom role for testing
transport.SetRole("test-role")
```

## Testing

The implementation includes comprehensive tests that verify:

-   Interface compliance
-   All client availability
-   Basic operation functionality
-   Error handling

Run tests with:

```bash
go test ./pkg/clients/rorclient/transports/mocktransport/ -v
```

## Features

-   ✅ Complete `RorTransport` interface implementation
-   ✅ All client methods implemented with realistic mock data
-   ✅ Proper error handling and validation
-   ✅ Configurable API secrets and roles
-   ✅ Event streaming support
-   ✅ Comprehensive test coverage
-   ✅ No external dependencies
-   ✅ Thread-safe operations

## Benefits

1. **Testing** - Perfect for unit tests and integration tests
2. **Development** - Develop against consistent mock data
3. **CI/CD** - Run tests without external dependencies
4. **Prototyping** - Rapid prototyping of ROR client usage
5. **Documentation** - Examples of expected API responses
