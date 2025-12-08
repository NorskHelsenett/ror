# Setup

## Prerequisites

### Mandatory

- A Linux distro or WSL2 for Windows
- Docker runtime
    - Docker CE: https://docs.docker.com/engine/install/
    - WSL2 tips: https://learn.microsoft.com/en-us/windows/wsl/systemd
- Golang SDK (For coding, debugging, and testing) https://go.dev
- ROR API: https://github.com/NorskHelsenett/ror-api

### Optional

If developing on the frontend the frontend repository is required:

- ROR Web: https://github.com/NorskHelsenett/ror-webapp

New version:

- ROR Web: https://github.com/NorskHelsenett/ror-web

Useful for spinning up the development environment:

- Docker Desktop (https://www.docker.com/products/docker-desktop/)

For kubernetes specific things:

- Talosctl (https://www.talos.dev/v1.8/introduction/quickstart/)
- Kind (https://kind.sigs.k8s.io)
- K3d (https://k3d.io/stable/)

For running documentation with mkdocs:

- Python

For working on the [ror Agent](../components/distributables/ror-agent/index.md):

- RO Agent: https://github.com/NorskHelsenett/ror-agent

## Hardware requirements:

| Recommendations | CPU | Memory |
| --------------- | --- | ------ |
| Minimum         | 2   | 16GB   |
| Recommended     | 4   | 32GB   |

# Setting up you repositories

To develop on ROR it is highly recommended to setup a Go Workspace with all relevant projects.

See [here](./installation/repos.md) for instructions.

## Install docker

To run the development environment Docker is required.

See [here](./installation/docker.md) for instructions.

## Starting the environment

For instructions on how to start the environment see the ror-api [README.md](https://github.com/NorskHelsenett/ror-api/blob/main/README.md) for instructions.

### ROR WEB

For instructions on how to start the webapp see the ror-web [README.md](https://github.com/NorskHelsenett/ror-web) for instructions.

## Swagger

To see swagger for ROR Api, go to http://localhost:10000/swagger/index.html

## Core infrastructure

For default urls see [here](./installation/default-urls.md).

## Default users

For default users see [here](./installation/default-accounts.md).

## Known issues

See [Known-issues](./installation/known-issues.md)

## Documentation

We pull documentation from code using **_some go package_**. Thus all functions should be annotated with a comment describing its use and any caveats. We keep system documentation in `cmd/docs/`, some files are copied in from .md files located in other parts of the repo using the `cmd/docs/collectdocs.sh` script. If you see any documentation that is out of date or wrong, please update it.
