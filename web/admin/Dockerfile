ARG DOCKER_MIRROR=docker.io/
FROM ${DOCKER_MIRROR}nginxinc/nginx-unprivileged:alpine
LABEL org.opencontainers.image.source https://github.com/norskhelsenett/ror
WORKDIR /app

COPY nginx/http_prod.conf /etc/nginx/conf.d/default.conf
COPY dist/ror-web/browser /app
