# Getting started with ROR development

## Prerequisites

-   Linux distro or WSL2 for windows
-   Docker runtime
    - Docker CE: https://docs.docker.com/engine/install/ 
    - WSL2 tips: https://learn.microsoft.com/en-us/windows/wsl/systemd
-   Golang SDK (For debugging and changing) https://go.dev
- ROR API: https://github.com/NorskHelsenett/ror-api
- ROR Web: https://github.com/NorskHelsenett/ror-webapp

### Optional:
-   Docker Desktop (https://www.docker.com/products/docker-desktop/)
-   Talosctl (https://www.talos.dev/v1.8/introduction/quickstart/)
-   Kind (https://kind.sigs.k8s.io)
-   K3d (https://k3d.io/stable/)
-   Python for running documentation with mkdocs
-   RO Agent: https://github.com/NorskHelsenett/ror-agent

## Clone

1.  Create a folder on you computer where you want to put the code
2.  Clone the repository
```bash
git clone git@github.com:NorskHelsenett/ror.git
```

```bash
git clone https://github.com/NorskHelsenett/ror.git
```

## Hardware requirements:

| Recommendations | CPU | Memory |
| --------------  | --  | -----  |
| Minimum         | 2   | 16GB   |
| Recommended     | 4   | 32GB   |

## Install docker

### Linux
Installation steps for Linux:
https://docs.docker.com/engine/install
Recommended post-installation steps:
https://docs.docker.com/engine/install/linux-postinstall/

#### Fedora
<details>
  <summary>Fedora</summary>

### Installations:
    
```bash
sudo dnf -y install dnf-plugins-core
sudo dnf-3 config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo
```

:warning: if you receive errors like this, you might have an old Docker installation already installed:
```bash
- package docker-ce-3:27.3.1-1.fc40.x86_64 from docker-ce-stable conflicts with docker provided by moby-engine-24.0.5-4.fc40.x86_64 from fedora
- package moby-engine-24.0.5-4.fc40.x86_64 from fedora conflicts with docker-ce provided by docker-ce-3:27.3.1-1.fc40.x86_64 from docker-ce-stable
```

#### Install the Docker Engine

```bash
sudo dnf install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

#### Start the Docker engine

```bash
sudo systemctl enable --now docker
```

#### (Optional) Install Docker auto-complete 

https://docs.docker.com/engine/cli/completion/

#### (Optional) Test the docker installation 

```bash
sudo docker run hello-world
```

#### Manage Dockker as a non-root

Doc reference: https://docs.docker.com/engine/install/linux-postinstall/

#### Create the docker group.

```bash
sudo groupadd docker
```

#### Add your user to the docker group.

```bash
sudo usermod -aG docker $USER
```

Log out and log back in so that your group membership is re-evaluated.
:warning: If you're running Linux in a virtual machine, it may be necessary to restart the virtual machine for changes to take effect.

#### Verify

```bash
docker run hello-world
```

</details>

### Windows

https://learn.microsoft.com/en-us/windows/wsl/systemd

TODO

## Starting ROR

### Run with docker

*Note Specific environment variables need to be set up for ROR to run, see* [Environment Variables](#Environment-Variables)

To start the ROR infrastructure you can run:

```bash
./r.sh
```

Which will start the [Core infrastructure](#Core-infrastructure)

To include any [optional services](#Optional-services) you can add them as arguments as shown:

```bash
./r.sh jaeger opentelemetry-collector
```
  
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

-   &lt;repo root&gt;/`.env` has the default settings for docker compose
-   Env variables used during development are set in `hacks/docker-compose/`
-   Env variables used in cluster are set with charts in `charts/`

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
| MongoDB       | Document database     | http://localhost:27017                                      |                                                                                       |                                             |
| Mongo-Express | Gui for document base | http://localhost:8081                                       |                                                                                       |                                             |
| RabbitMQ      | Message bus           | GUI: http://localhost:15672, amqp port: localhost:5672      |                                                                                       |                                             |
| Vault         | Secrets handling      | GUI: http://localhost:8200                                  |                                                                                       |                                             |
| Redis         | Cache                 | GUI: http://localhost:6379                                  |                                                                                       |                                             |
| Redis-insight | Cache insight         | GUI: http://localhost:8001                                  |                                                                                       |                                             |

## Default users

| Service       | Username   | Password  |
| -------       | ---------- | --------- |
| MongoDB       | `someone`  | `S3cret!` |
| Mongo-Express | `someone`  | `S3cret!` |
| RabbitMQ      | `admin`    | `S3cret!` |

## Optional services

 Service                  | What                  | Url                                                                | ReadMe link                                                                           | Comment                                     |
| ----------------------- | --------------------- | ------------------------------------------------------------------ | ------------------------------------------------------------------------------------- | ------------------------------------------- |
| jaeger                  |                       |                                                                    |                                                                                       |                                             |
| opentelemetry-collector |                       |                                                                    |                                                                                       |                                             |

## ROR services

| Service    | What      | Url                    | Port | ReadMe link                                                | Comment                   |
| ---------  | --------- | ---------------------- | ---- | ---------------------------------------------------------- | ------------------------- |
| ROR-Api    | Api       | http://localhost:10000 | 8080 | [ror-api](https://github.com/NorskHelsenett/ror-api)       |                           |
| ROR-WebApp | Web       | http://localhost:11000 | 8090 | [ror-webapp](https://github.com/NorskHelsenett/ror-webapp) |                           |
| ROR-Agent  | K8s Agent | http://localhost:8100  | 8100 | [ror-agent](https://github.com/NorskHelsenett/ror-agent)   | Not run by docker-compose |

## Documentation

We pull documentation from code using **_some go package_**. Thus all functions should be annotated with a comment describing its use and any caveats. We keep system documentation in `cmd/docs/`, some files are copied in from .md files located in other parts of the repo using the `cmd/docs/collectdocs.sh` script. If you see any documentation that is out of date or wrong, please update it.
