# Getting started with creating a ROR resource definition

## Prerequisites

- Good knowledge about the resource you're planning to implement
- Golang SDK: https://go.dev

### Optional

- A clear picture of what you want to implement and how go generalize it

## General information

When creating a new resource, editing, or deleting an existing resource for ROR the generator MUST be ran, this is required for amongst other things to update the TypeScript types.

## Getting started

### Cloning the repository

To start with you need to first setup your repositories ready for development.
For how to do that please see [here](../installation/repos.md)

### Creating the base resource files

#### Create the resource definition

1. Creating it at:

```
ror/pkg/rorresources/rortypes
```

with the format:

```
resourcedef_<resourceName>.go
```

2. When creating this file the following rule must be followed:
    - The package name is `rortypes`
    - The structs all begin with `Resource` resulting your type to be `Resource<YourResourceName>`
    - Your main struct should contain a `Spec` and a `Status` field
    - The `Spec` is the desired state, while `Status` is observed state, more about that further down.

An example of a resourcedef:

```
package rortypes

type ResourceVirtualMachine struct {
        Id     string                       `json:"id"`
        Name   string                       `json:"name"`
        Spec   ResourceVirtualMachineSpec   `json:"spec"`
        Status ResourceVirtualMachineStatus `json:"status"`
}

type ResourceVirtualMachineSpec struct {
        Cpu             ResourceVirtualMachineCpuSpec             `json:"cpu"`
        Tags            []ResourceVirtualMachineTagSpec           `json:"tags"`
        Disks           []ResourceVirtualMachineDiskSpec          `json:"disks"`
        Memory          ResourceVirtualMachineMemorySpec          `json:"memory"`
        Networks        []ResourceVirtualMachineNetworkSpec       `json:"networks"`
        OperatingSystem ResourceVirtualMachineOperatingSystemSpec `json:"operatingSystem"`
}

type ResourceVirtualMachineStatus struct {
        Cpu             ResourceVirtualMachineCpuStatus             `json:"cpu"`
        Disks           []ResourceVirtualMachineDiskStatus          `json:"disks"`
        Memory          ResourceVirtualMachineMemoryStatus          `json:"memory"`
        Networks        []ResourceVirtualMachineNetworkStatus       `json:"networks"`
        OperatingSystem ResourceVirtualMachineOperatingSystemStatus `json:"operatingSystem"`
}
.
.
.
```

##### Spec

Spec will be what we desire the configuration to be for when we wish to change something, like cpu, disk, or memory for the example above.
Any parameter within here it is expected we're allowed to change if change from ROR is implmeneted.

If the resource is read-only Spec is necessary to implement.

##### Status

Status will be what we observe about this resource, like ids, the current cpu, disk, or memory for the example above.

#### Create the resource_input_filter definition

The resource_input_filter runs on import on that type, it can be used to censor, anonymize, remove, or change data before import.

While it has to be defined on every type, generally we create it and return nil for no action.

1. Create it at:

```
ror/pkg/rorresources/rortypes
```

with the format:

```
resouce_input_filter<resourceName>.go
```

2. Which should contain something to this example with <YourResouceType> swapped out with your core struct:

```
package rortypes

// (r *<ResourceName>) ApplyInputFilter Applies the input filter to the resource
func (r *<ResourceName>) ApplyInputFilter(cr *CommonResource) error {
        return nil
}
```

#### ROR definitions

1. Go to:

```bash
ror/pkg/rorresources/rordefs
```

2. And edit defs.go
3. On the top const definition, add your new agent in the format:

```
ApiResoureType<Name> ApiReesourceType = "<AgentName>"
```

4. And at the near bottom at at to the Resroucedefs slice:

```
{
        TypeMeta: metav1.TypeMeta{
                Kind:       "<Name",
                APIVersion: "<Version of your choice>",
        },
        Plural:     "<Plural of name>",
        Namespaced: false,
        Types:      []ApiResourceType{<Type>},
},
```

#### Running the generator

1. Go to the root folder of ROR
2. Run the following command:

```
go run cmd/generator/main.go
```

3. And you're done, commit to the branch and make a pull/merge request.
