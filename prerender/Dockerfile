FROM node:10.19-alpine:3.11

ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/lib/chromium/

WORKDIR /app

# Installs latest Chromium package.
RUN addgroup -S prerender && adduser -S -g prerender prerender \
    && apk add --no-cache zstd nspr snappy ttf-opensans \
    ca-certificates \
    chromium \
    harfbuzz \
    nss \
    freetype \
    ttf-freefont \
    && yarn add prerender@5.6.0 \
    && rm -rf /var/cache/*

COPY server.js /app/server.js

# Add user so we don't need --no-sandbox.
# Run everything after as non-privileged user.
USER prerender

EXPOSE 3000

CMD [ "node", "server.js" ]
