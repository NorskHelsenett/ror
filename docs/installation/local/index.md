# local installation

To install ROR locally on your machine, certain prerequisites must be installed:

## Total Prerequisites

### Hardware requirements:

While anything less is possible we highly discourage it.

| Recommendations | CPU | Memory |
| --------------- | --- | ------ |
| Minimum         | 2   | 16GB   |
| Recommended     | 4   | 32GB   |

### Software

- A Linux distro or WSL2 for Windows
- Docker runtime
    - Docker CE: https://docs.docker.com/engine/install/
    - WSL2 Tips: https://learn.microsoft.com/en-us/windows/wsl/systemd
- Golang SDK (For debugging and changing) https://go.dev
- ROR API: https://github.com/NorskHelsenett/ror-api
- ROR Web: https://github.com/NorskHelsenett/ror-webapp

Currently a Windows installation is not supported.

### Optional:

- Docker Desktop (https://www.docker.com/products/docker-desktop/)
- Talosctl (https://www.talos.dev/v1.8/introduction/quickstart/)
- Kind (https://kind.sigs.k8s.io)
- K3d (https://k3d.io/stable/)
- Python for running documentation with mkdocs
- RO Agent: https://github.com/NorskHelsenett/ror-agent

## Start

To start with you need to first setup your repositories ready for development.
For how to do that please see (here)[./repos.md]

## Install docker

To install docker see (here)[./docker.md]

## Starting ROR

### Run with docker

_Note Specific environment variables need to be set up for ROR to run, see_ [Environment Variables](#Environment-Variables)

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

For ROR to work you require minimum the API, which can be found here:
https://github.com/NorskHelsenett/ror-api

#### Visual Studio Code

1. Open the repository in VSCode
2. Go to Debugging
3. On "Run and debug" select "Debug ROR-Api" or "Debug ROR-Api tests"

#### Terminal

TODO

### ROR WEB

Clone the ror-webapp repository

```bash
git clone https://github.com/NorskHelsenett/ror-webapp
```

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

See [Known-issues](known-issues.md)

## Documentation

We pull documentation from code using **_some go package_**. Thus all functions should be annotated with a comment describing its use and any caveats. We keep system documentation in `cmd/docs/`, some files are copied in from .md files located in other parts of the repo using the `cmd/docs/collectdocs.sh` script. If you see any documentation that is out of date or wrong, please update it.
