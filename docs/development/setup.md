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

If developing on the frontend this one becomes mandatory:

- ROR Web: https://github.com/NorskHelsenett/ror-webapp

New version:

- ROR Web: https://github.com/NorskHelsenett/ror-web

Useful for spinning up the development environment:

- Docker Desktop (https://www.docker.com/products/docker-desktop/)

For kubernetes specific things:

- Talosctl (https://www.talos.dev/v1.8/introduction/quickstart/)
- Kind (https://kind.sigs.k8s.io)
- K3d (https://k3d.io/stable/)

For running documentation with mkdocs

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

### Run with docker

For environment variables an `env.example` file exists in `ror/hacks/env/env.example` that will be copied.
If any modifications is wanted - like adding your own dockerhub mirror - you can copy that to the root folder
and name it `.env` and edit it.

To start the ROR infrastructure you can run:

```bash
./r.sh
```

Which will start the [Core infrastructure](#Core-infrastructure)

To include any [optional services](#Optional-services) you can add them as arguments as shown:

```bash
./r.sh jaeger opentelemetry-collector
```

When the containers start you'll note that the following services will keep crashing,
This is intended as they're dependent on the API service which has yet to be started:

- dex
- ms-auth
- ms-talos
- ms-kind

### ROR API

#### Visual Studio Code

1. Open the repository in VSCode
2. Go to Debugging
3. On "Run and debug" select "Debug ROR-Api" or "Debug ROR-Api tests"

#### Terminal

1. If you haven't already, start r.sh
2. Start ./debug.sh

### ROR WEB

Start the core services as mentioned [Here](#Starting-ROR)

Start the API as mentioned [Here](#ROR-API)

Start the webapp as mentioned [Here](https://github.com/NorskHelsenett/ror-webapp)

#### VSCode

TODO

#### Terminal

TODO

### Environment Variables

- &lt;repo root&gt;/`.env` has the default settings for docker compose
- Env variables used during development are set in `hacks/docker-compose/`
- Env variables used in cluster are set with charts in `charts/`

## Login to ROR-web

Open your favorite browser, and go to http://localhost:11000
Log in with any of these accounts:

| Title           | Username             | Password  |
| --------------- | -------------------- | --------- |
| super admin     | `superadmin@ror.dev` | `S3cret!` |
| read only admin | `readadmin@ror.dev`  | `S3cret!` |
| developer 1     | `dev1@ror.dev`       | `S3cret!` |
| developer 2     | `dev2@ror.dev`       | `S3cret!` |

## Swagger

To see swagger for ROR Api, go to http://localhost:10000/swagger/index.html

## Core infrastructure

| Service       | What                  | Url                                                         | ReadMe link                                                                           | Comment                                     |
| ------------- | --------------------- | ----------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------- |
| DEX           | Authentication        | www: http://localhost:5556, grpc api: http://localhost:5557 | [dex doc](https://dexidp.io/docs/) [docker hub](https://hub.docker.com/r/bitnami/dex) | Reachable from inside and outside of docker |
| Openldap      | Mocking users         | http://localhost:389                                        |                                                                                       |                                             |
| MongoDB       | Document database     | mongodb://localhost:27017                                   |                                                                                       |                                             |
| Mongo-Express | Gui for document base | http://localhost:8081                                       |                                                                                       |                                             |
| RabbitMQ      | Message bus           | GUI: http://localhost:15672, amqp port: localhost:5672      |                                                                                       |                                             |
| Vault         | Secrets handling      | GUI: http://localhost:8200                                  |                                                                                       |                                             |
| Valkey        | Cache                 | GUI: http://localhost:6379                                  |                                                                                       |                                             |

## Default users

| Service       | Username  | Password  |
| ------------- | --------- | --------- |
| MongoDB       | `someone` | `S3cret!` |
| Mongo-Express | `test`    | `S3cr3t`  |
| RabbitMQ      | `admin`   | `S3cret!` |

## Optional services

| Service                 | What | Url | ReadMe link | Comment |
| ----------------------- | ---- | --- | ----------- | ------- |
| jaeger                  |      |     |             |         |
| opentelemetry-collector |      |     |             |         |

## ROR services

| Service    | What      | Url                    | Port | ReadMe link                                                | Comment                   |
| ---------- | --------- | ---------------------- | ---- | ---------------------------------------------------------- | ------------------------- |
| ROR-Api    | Api       | http://localhost:10000 | 8080 | [ror-api](https://github.com/NorskHelsenett/ror-api)       |                           |
| ROR-WebApp | Web       | http://localhost:11000 | 8090 | [ror-webapp](https://github.com/NorskHelsenett/ror-webapp) |                           |
| ROR-Agent  | K8s Agent | http://localhost:8100  | 8100 | [ror-agent](https://github.com/NorskHelsenett/ror-agent)   | Not run by docker-compose |

## Known issues

See [Known-issues](./installation/known-issues.md)

## Documentation

We pull documentation from code using **_some go package_**. Thus all functions should be annotated with a comment describing its use and any caveats. We keep system documentation in `cmd/docs/`, some files are copied in from .md files located in other parts of the repo using the `cmd/docs/collectdocs.sh` script. If you see any documentation that is out of date or wrong, please update it.
