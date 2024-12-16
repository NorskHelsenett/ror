# Getting started with ROR Client

## Prerequisites
    
- A ROR resourcedef - To create this see docs/ror-generator-client/docs/generator/getting-started.md

### Optional

- TODO

## Getting started

### Configuration variables

First to implement it we need to add the rorclient to the project

```
go get "github.com/NorskHelsenett/ror/pkg/clients/rorclient"
```

And then add configuration variables for the following parameters:


``BaseUrl``

The Url of the ROR instance you're communicating with.

---------

``AuthProvider``

The auth proivder from "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpauthprovider"
which is either:
httpauthprovider.NewNoAuthProvider()
or an actual AuthProvider

TODO

---------

``Version``

The current version of the client, for example 0.0.1:<commit>

---------

``Role``

The name of the client, for example vpshereAgent

---------

### Implementation

#### Example

##### Implementing the ROR client

<details>
  <summary>Example</summary>

```go
package rorclient

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
)

type config struct {
    RorUrl string
    RorRole string
    RorCommit string
    RorVersion string
}

// Populates the Config with values from environmental variables or static values 
func NewConfig() *Config {
...
...
...
}

var config Config = NewConfig()

type UpdateError struct {
	uuid    string
	status  int
	message string
}

// So we can add our own methods to the RorClient struct
type RorClient struct {
	rorclient.RorClient
}

// Constructs 
func NewRorClient(config config.Config) *RorClient {
	transport := resttransport.NewRorHttpTransport(&httpclient.HttpTransportClientConfig{
		BaseURL:      config.RorUrl,
		AuthProvider: httpauthprovider.NewNoAuthprovider(),
		Version:      rorversion.NewRorVersion(config.RorVersion, config.RorCommit),
		Role:         config.RorRole,
	})

	rorClient := RorClient{
		*rorclient.NewRorClient(transport),
	}

	return &rorClient
}
```

</details>

##### Converting the resource to the ROR resource

<details>
  <summary>Example</summary>

```go
package rorclient

// Adds or updates the VMs in ROR
func (r *RorClient) UpdateVms(ctx context.Context, vmResources []*rortypes.ResourceVirtualMachine) error {
	set := rorresources.NewResourceSet()

	names := []string{}
	for _, vm := range vmResources {
		names = append(names, vm.Name)
		res := rorresources.NewRorResource("VirtualMachine", "general.ror.internal/v1alpha1")

		res.RorMeta.Ownerref = rortypes.RorResourceOwnerReference{
			Scope:   aclmodels.Acl2ScopeRor,
			Subject: aclmodels.Acl2RorSubjectGlobal,
		}

		v5Uuid, err := virtualmachine.UuidV5FromCompositeId(vm.Id)
		if err != nil {
			return fmt.Errorf("could not get uuid from vm resource: %w", err)
		}

		res.Metadata.UID = types.UID(v5Uuid.String())
		res.RorMeta.Action = rortypes.K8sActionAdd
		res.Metadata.Name = vm.Name
		res.SetVirtualMachine(vm)
		set.Add(res)
	}

	res, err := r.ResourceV2().Update(ctx, *set)
}
```

</details>

##### Adding or updating using the ROR client

<details>
  <summary>Example</summary>

```go
package rorclient

// Adds or updates the VMs in ROR
func (r *RorClient) UpdateVms(ctx context.Context, vmResources []*rortypes.ResourceVirtualMachine) error {
		res, err := r.ResourceV2().Update(ctx, *set)
	if err != nil {
		return fmt.Errorf("could not update vm, ROR returned: %w", err)
	}

	var errors UpdateErrors
	for uuid, response := range *&res.Results {
		if response.Status > 299 {
			error := UpdateError{
				uuid:    uuid,
				status:  response.Status,
				message: response.Message,
			}

			errors.Errors = append(errors.Errors, error)
		}
	}

	if len(errors.Errors) != 0 {
		return errors
	}

	return nil
}

// Deletes VMs
func (r *RorClient) DeleteVms(ctx context.Context, uuids []string) error {
	for _, uid := range uuids {
		del, err := r.ResourceV2().Delete(ctx, uid)
		slog.Info("deleted vm", "info", del)
		if err != nil {
			return fmt.Errorf("could not delete vm, ROR returned %w", err)
		}
	}
	return nil
}
```

</details>

##### Deleting using the ROR client

<details>
  <summary>Example</summary>

```go
package rorclient

func (r *RorClient) DeleteVms(ctx context.Context, uuids []string) error {
	for _, uid := range uuids {
		del, err := r.ResourceV2().Delete(ctx, uid)
		slog.Info("deleted vm", "info", del)
		if err != nil {
			return fmt.Errorf("could not delete vm, ROR returned %w", err)
		}
	}
	return nil
}
```

</details>


### Instructions

