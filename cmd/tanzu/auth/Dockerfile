ARG DOCKER_MIRROR=docker.io/
FROM ${DOCKER_MIRROR}alpine:latest
LABEL org.opencontainers.image.source https://github.com/norskhelsenett/ror
RUN apk --no-cache add ca-certificates gcompat
WORKDIR /

COPY --chown=1001:2000 cmd/tanzu/auth/app /app
RUN chmod 755 /app/kubectl /app/kubectl-vsphere /app/entrypoint.sh
COPY --chown=1001:2000 dist/tanzuAuth /app/tanzuAuth
COPY cmd/tanzu/auth/version.json /version.json

ENTRYPOINT ["/app/tanzuAuth"]
