# Installation of Docker

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
