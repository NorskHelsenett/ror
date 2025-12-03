# Getting started with ROR development

## Prerequisites

-   Docker
-   Go

## Starting ROR

### ROR-API

Clone the ror-api repositiroy and run

```bash
git clone https://github.com/NorskHelsenett/ror-api
cp .env.example .env
docker compose up -d
go run cmd/api/main.go
```

### ROR WEB

Clone the ror-webapp repository

```bash
git clone https://github.com/NorskHelsenett/ror-webapp
npm i
nmp start
```

Point your browser to http://localhost:11000

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
