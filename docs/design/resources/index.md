# Resources

## Quick words

### v1

ROR resource v1 is deprecated. It was not very modular and there was a lot of manual steps to create a new resource. Each type was its own full definition with very little reuse of code between them.

### v2

ROR resource v2 is to wrap all resources in a common resource type and generate code to work with each kind.

## What is a resource

A ROR-Resource (resource), is a structured piece of data with some common metadata attached. Each resource must be explicitly defined in the ROR resource package. This means that if you want to create a new resource you must submit a Pull Request to the ROR repo. At this time resources are not built to be registered at runtime.

One of RORs core [philosophies](https://norskhelsenett.github.io/ror/philosophy/) is vendor agnosticism. This means that if you are trying to add a resource for for example Azure virtual machines, you should generalize the resource so that it can be used for VMware virtual machines and so on. This keeps the number of resources down, and also makes it easier to swap vendor in the future.

### Resource queries

TODO

### Resource sets

TODO

## Creating a new resource

To create a new resource see our development guide [here](https://norskhelsenett.github.io/ror/development/generator/resources/)
