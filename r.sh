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
while getopts "e:" arg; do
  case $arg in
    e) envfile=$OPTARG;;
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

docker compose --env-file $envfile up openldap dex init-dex-db vault mongodb rabbitmq mongo-express redis ms-auth ms-kind ms-talos $services