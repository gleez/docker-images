# Smallest base image
FROM alpine:latest

MAINTAINER Sandeep Sangamreddi <sandeepone@gmail.com>

ENV version=fc327b2

# Add dependencies
RUN apk --update upgrade && \
	apk add --update ca-certificates && \
    rm -rf /tmp/* /var/tmp/* /var/cache/apk/* /var/cache/distfiles/*

##COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Download grimd binary
COPY grimd /bin/grimd

# Chmod binary and add config folder
RUN chmod a+x /bin/grimd && \
  mkdir -p ~/grimd/config && \
  ln -s ~/grimd/config /config

# Copy custom config file to volume
COPY grimd.toml ~/grimd/config/grimd.toml

# Forward appropriate ports
EXPOSE 53/udp 53/tcp 8080/tcp

# Read grimd.toml from volume
VOLUME ["/config", "/log"]

# Run grimd
ENTRYPOINT ["/bin/grimd", "-config", "~/grimd/config/grimd.toml", "-update"]