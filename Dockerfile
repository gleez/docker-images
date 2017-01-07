# ------------------------------------------------------------------------------
# Based on a work at https://github.com/docker/docker.
# ------------------------------------------------------------------------------
# Pull base image.
FROM alpine:3.4
MAINTAINER Sandeep Sangamreddi <sandeepone@gmail.com>

RUN apk --update add build-base g++ make curl wget openssl-dev git libxml2-dev sshfs bash tmux nodejs supervisor \
 && rm -f /var/cache/apk/*\
 && git clone https://github.com/c9/core.git /cloud9 \
 && curl -s -L https://raw.githubusercontent.com/c9/install/master/link.sh | bash \
 && /cloud9/scripts/install-sdk.sh \
 && sed -i -e 's_127.0.0.1_0.0.0.0_g' /cloud9/configs/standalone.js \
 && mkdir /workspace \
 && mkdir -p /var/log/supervisor

# Add supervisord conf
ADD supervisord.conf /etc/

# ------------------------------------------------------------------------------
# Add volumes
VOLUME /workspace

# ------------------------------------------------------------------------------
# Expose ports.
EXPOSE 3000

# ------------------------------------------------------------------------------
# Start supervisor, define default command.
ENTRYPOINT ["supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf"]