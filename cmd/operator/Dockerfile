ARG GCR_MIRROR=gcr.io/
FROM ${GCR_MIRROR}distroless/static:nonroot
LABEL org.opencontainers.image.source https://github.com/norskhelsenett/ror
WORKDIR /
COPY dist/operator /bin/operator
COPY cmd/operator/version.json /version.json

ENTRYPOINT ["/bin/operator"]
