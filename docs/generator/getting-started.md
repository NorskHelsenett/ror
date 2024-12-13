# Getting started with creating a ROR resource definition

## Prerequisites

- Good knowledge about the resource you're planning to implement
- Golang SDK: https://go.dev

### Optional

- A clear picture of what you want to implement and how go generalize it

## Getting started

## General information

When creating a new resource for ROR the entire procedure needs to be followed.
If you're just updating the existing resources - As in not adding or removing structs - You don't need to follow any of these instructions from the ROR side
If you're adding more sub structs but not adding a core struct you need to rerun the generator again.

### Cloning the repository

1. Create a folder where you want to store the code.
2. Clone the repository to where you have Go installed.

```bash
git clone git@github.com:NorskHelsenett/ror.git
```

```bash
git clone https://github.com/NorskHelsenett/ror.git
```

For more information on this part see: https://norskhelsenett.github.io/ror/getting-started/


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
    - The package name is ``rortypes``
    - The structs all begin with ``Resource`` resulting your type to be ``Resource<YourResourceName>``
    - Your main struct should contain a ``spec`` and a ``status`` field
    - The ``spec`` is the desired state, while ``status`` is observed state

An example of a resourcedef:

```
package rortypes

type ResourceVirtualMachine struct {
        Id     string                       `json:"id"`
        Name   string                       `json:"name"`
        Spec   ResourceVirtualMachineSpec   `json:"spec"`
        Status ResourceVirtualMachineStatus `json:"status"`
}

// Desired state
type ResourceVirtualMachineSpec struct {
        Cpu             ResourceVirtualMachineCpuSpec             `json:"cpu"`
        Tags            []ResourceVirtualMachineTagSpec           `json:"tags"`
        Disks           []ResourceVirtualMachineDiskSpec          `json:"disks"`
        Memory          ResourceVirtualMachineMemorySpec          `json:"memory"`
        Networks        []ResourceVirtualMachineNetworkSpec       `json:"networks"`
        OperatingSystem ResourceVirtualMachineOperatingSystemSpec `json:"operatingSystem"`
}

// Observed state
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

#### Create the resource_input_filter definition 

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

#### Running the generator

1. Go to the root folder of ROR
2. Run the following command:

```
go run cmd/generator/main.go
```

3. And you're done, commit to the branch and make a pull/merge request.

### Afterwards

- To import data to this (newly) created ROR resource, you need to implement the ROR client, which you can read more about here: ror/docs/rorclient/getting-started.md
