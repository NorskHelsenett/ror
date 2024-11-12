ARG DOCKER_MIRROR=docker.io/
FROM ${DOCKER_MIRROR}alpine:latest
LABEL org.opencontainers.image.source https://github.com/norskhelsenett/ror

RUN apk --no-cache add ca-certificates gcompat

COPY --chown=1001:2000 cmd/tanzu/agent/app /app
COPY --chown=1001:2000 dist/tanzuAgent /app/tanzuAgent

ENTRYPOINT ["/app/entrypoint.sh"]
