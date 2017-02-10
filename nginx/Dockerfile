FROM ubuntu:xenial
MAINTAINER Sandeep Sangamreddi <sandeepone@gmail.com>

ENV NGINX_VERSION 1.10.3-4~xenial

COPY nginx_${NGINX_VERSION}_amd64.deb /tmp/

RUN apt-get update \
##	&& apt-get upgrade -y \
	&& apt-get install --no-install-recommends --no-install-suggests -y \
				libc6 libexpat1 libgcc1 libgd3 libgeoip1 libhiredis0.13 \
				libluajit-5.1-2 lua-cjson libpcre3 libperl5.22 libstdc++6 libxml2 \
				libxslt1.1 zlib1g lsb-base adduser ca-certificates gettext-base \
	&& dpkg -i /tmp/nginx_${NGINX_VERSION}_amd64.deb \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/* /tmp/* \

# forward request and error logs to docker log collector
	&& ln -sf /dev/stdout /var/log/nginx/access.log \
	&& ln -sf /dev/stderr /var/log/nginx/error.log

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]

