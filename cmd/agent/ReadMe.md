# NHN-ROR-Agent

K8s agent for NHN-Ror to report from clusters inside Privat Sky.

## Prerequisites

1. Golang 1.20.x or newer [GoDev](https://go.dev/dl)
2. OCI image builder (docker, podman, etc)

## Connect to cluster

- Locally ror-agent uses the `%userprofile%/.kube/config` as default to connect to the cluster.

### Create test cluster with k3d (https://k3d.io)

```bash
k3d cluster create k8s --api-port 65001 -p "10081:80@loadbalancer" --agents 2
```

Spinning up a cluster in docker-desktop, with a loadbalancer and 2 agents. [More info](https://k3d.io/v5.4.1/usage/exposing_services/)

# Get started, and run it

- Open repo root in your favorite IDE (VS Code, VS, Rider, etc)
- Open a terminal
- go to `<repo root>/src/clients/ror-agent`
- `go get` (install dependencies)
- `go build -o agent` -> results in a executable file (win: agent.exe, unix: agent)
- run `agent`

# Debug in Visual Studio Code

- Open `<repo root>` as workspace/folder in VS Code
- Open terminal, go to `<repo root>/src/clients/ror-agent`, and run `go get`
- Go to the debug button on the sidebar in VS Code
- Start `Debug Ror-agent` configuration
- Set your breakpoints in the code

# Health endpoint

Go to [health endpoint: https://localhost:8090/health](https://localhost:8090/health) to check health

# Trigger add/update/delete ingress changes

- Open terminal
- Go to `<repo root>\testdata`
- Apply test ingress manifest
  - `kubectl apply -f avi-ingress-external.yaml`
