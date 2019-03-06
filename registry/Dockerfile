FROM golang:1.11-alpine AS build

ENV DISTRIBUTION_DIR /go/src/github.com/docker/distribution
ENV DOCKER_BUILDTAGS include_oss include_gcs

ARG GOOS=linux
ARG GOARCH=amd64
ARG GOARM=6

RUN set -ex \
    && apk add --no-cache make git file

WORKDIR $DISTRIBUTION_DIR
COPY . $DISTRIBUTION_DIR
RUN CGO_ENABLED=0 make PREFIX=/go clean binaries && file ./bin/registry | grep "statically linked"

# https://github.com/docker/distribution-library-image/tree/master/amd64
# https://github.com/leoh0/distribution/commit/abcacf2d0e35278c6a42cf762893ce62f1db217e

# Build a minimal distribution container

FROM alpine:3.8

RUN set -ex \
    && apk add --no-cache ca-certificates apache2-utils

#COPY ./registry /bin/registry
COPY --from=build /go/src/github.com/docker/distribution/bin/registry /bin/registry
COPY ./config-example.yml /etc/docker/registry/config.yml

VOLUME ["/var/lib/registry"]
EXPOSE 5000

COPY docker-entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["/etc/docker/registry/config.yml"]