# Getting started with ROR development

## Prerequisites

-   Linux distro or WSL2 for windows
-   Docker runtime
    -   wsl2 tips: https://learn.microsoft.com/en-us/windows/wsl/systemd
-   Golang SDK (if you want to change and debug ROR) https://go.dev

Optional:

-   Docker Desktop (https://www.docker.com/products/docker-desktop/)
-   Talosctl (https://www.talos.dev/v1.8/introduction/quickstart/)
-   Kind (https://kind.sigs.k8s.io)
-   K3d (https://k3d.io/v5.7.4/#releases)
-   Python for running documentation with mkdocs

## Clone

1.  Create a folder on you computer where you want to put the code
2.  git clone (ROR GIT URL, https:// or git://)

## Hardware demands:

Minimum 16 gb RAM, but this will potentially painfull... Recommended is 32 gb RAM or more

## Run with docker

```bash
./r.sh api web
```

This runs containers; **dex**, **openldap**, **vault**, **rabbitmq**, **mongodb**, **mongo-express**, **redis**, **ms-auth**, **ms-kind**, **ms-talos**. Does not use that much memory

If you want to run **_all_**, this includes all the container above, and jaeger and opentelemetry collector

```bash
docker compose up
```

## Login to ROR-web

Open your favorite browser, and go to http://localhost:11000
Log in with:
Accounts:

-   "super admin"
    -   Login with `superadmin@ror.dev` and `S3cret!`
-   Read only admin
    -   Login with `readadmin@ror.dev` and `S3cret!`
-   developer 1
    -   Login with `dev1@ror.dev` and `S3cret!`
-   developer 2
    -   Login with `dev2@ror.dev` and `S3cret!`

To see swagger for ROR Api, go to http://localhost:10000/swagger/index.html

### Environment Variables

-   &lt;repo root&gt;/`.env` has the default settings for docker compose
-   Env variables used during development are set in `hacks/docker-compose/`
-   Env varaibles used in cluster are set with charts in `charts/`

## Needed infrastructure

| Service       | What                  | Url                                                                | ReadMe link                                                                           | Comment                                     |
| ------------- | --------------------- | ------------------------------------------------------------------ | ------------------------------------------------------------------------------------- | ------------------------------------------- |
| DEX           | Authentication        | www: http://localhost:5556, <br /> grpc api: http://localhost:5557 | [dex doc](https://dexidp.io/docs/) [docker hub](https://hub.docker.com/r/bitnami/dex) | Reachable from inside and outside of docker |
| Openldap      | Mocking users         | http://localhost:389                                               |                                                                                       |                                             |
| MongoDb       | Document database     | localhost:27017                                                    |                                                                                       |                                             |
| Mongo-Express | Gui for document base | http://localhost:8081                                              |                                                                                       |                                             |
| RabbitMq      | Message bus           | GUI: http://localhost:15672, <br />, amqp port: localhost:5672     |                                                                                       |                                             |
| Vault         | Secrets handling      | GUI: http://localhost:8200                                         |                                                                                       |                                             |
| Redis         | Cache                 | GUI: http://localhost:6379                                         |                                                                                       |                                             |
| Redis-insight | Cache insight         | GUI: http://localhost:8001                                         |                                                                                       |                                             |

## NHN-ROR services

| Service   | What            | Url                    | Port | ReadMe link                                    | Comment                   |
| --------- | --------------- | ---------------------- | ---- | ---------------------------------------------- | ------------------------- |
| ROR-Api   | WebApi          | http://localhost:10000 | 8080 | [ReadMe.md](./src/backend/ror-api/ReadMe.md)   |                           |
| ROR-Admin | Adminportal GUI | http://localhost:11000 | 8090 | [ReadMe.md](./src/clients/ror-admin/README.md) |                           |
| ROR-Agent | K8s agent       | http://localhost:8100  | 8100 | [ReadMe.md](./src/clients/ror-agent/README.md) | Not run by docker-compose |

## Documentation

We pull documentation from code using **_some go package_**. Thus all functions should be annotated with a comment describing its use and any caveats. We keep system documentation in `cmd/docs/`, some files are copied in from .md files located in other parts of the repo using the `cmd/docs/collectdocs.sh` script. If you see any documentation that is out of date or wrong, please update it.
