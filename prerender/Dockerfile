FROM node:10.16-alpine

ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/

WORKDIR /app

# Installs latest Chromium package.
RUN addgroup -S prerender && adduser -S -g prerender prerender \
    && echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories \
    && echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories \
    && apk add --no-cache zstd nspr snappy ttf-opensans \
    ca-certificates \
    chromium@edge \
    harfbuzz@edge \
    nss@edge \
    freetype@edge \
    ttf-freefont@edge \
    && yarn add prerender@5.6.0 \
    && rm -rf /var/cache/*

COPY server.js /app/server.js

# Add user so we don't need --no-sandbox.
# Run everything after as non-privileged user.
USER prerender

EXPOSE 3000

CMD [ "node", "server.js" ]
