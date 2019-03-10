FROM alpine:3.9

# File Author / Maintainer
LABEL name="Alpine-JQ" \
      maintainer="Nick Lalumiere <kurrier@thelysium.com>" \
      version=$JQ_VERSION \
      git-commit=$GIT_COMMIT

# Update Repos and Install curl/jq
RUN echo http://dl-cdn.alpinelinux.org/alpine/v3.9/main > /etc/apk/repositories && \
    echo http://dl-cdn.alpinelinux.org/alpine/v3.9/community >> /etc/apk/repositories && \
    apk add --no-cache curl bash && \
    apk add --no-cache jq && \
    addgroup -g 1000 jq && \
    adduser -D -u 1000 -s /bin/sh -G jq jq

ENV JQ_VERSION="3.0"

# User
USER jq

# Start the service
#ENTRYPOINT ["/bin/sh"]