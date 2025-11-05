#!/bin/sh
if [ ! -d ./.vscode ]; then
  mkdir -p ./.vscode;
fi

if [ ! -f ./.vscode/launch.json ]; then
  cp -v ./hacks/vscode/launch.json ./.vscode/launch.json;
fi

if [ ! -f ./.vscode/settings.json ]; then
  cp -v ./hacks/vscode/settings.json ./.vscode/settings.json;
fi

envfile=".env"

if [ ! -f "$envfile" ]; then
  cp -v ./hacks/env/env.template "$envfile";
fi

while getopts "e:" arg; do
  case $arg in
    e) envfile=$OPTARG;;
    *) echo "error: Invalid option"
       exit 1 ;;
  esac
done

skip_next=false
services=""
for arg in "$@"; do
    if [ "$skip_next" = true ]; then
        skip_next=false
        continue
    fi

    if [ "$arg" = "-e" ]; then
        skip_next=true
        continue
    fi

    services="$services $arg"
done

if ! test -f "$envfile"; then
  echo "File $envfile does not exist, exiting ..."
  exit 1
fi

if docker compose > /dev/null 2>&1; then
  compose_cmd="docker compose"
elif command -v docker-compose; then
  compose_cmd="docker-compose"
else
  echo "Error: Docker Compose is not installed. Please install it and try again."
  exit 1
fi
version=$($compose_cmd version --short | cut -d "." -f 1)
if [ "$version" -lt 2 ]; then
    echo "Docker Compose version is to low, needs to be v2.0.0 or higher."
    exit 1
fi

$compose_cmd --env-file $envfile up openldap dex init-dex-db vault mongodb rabbitmq mongo-express valkey ms-auth ms-kind $services
