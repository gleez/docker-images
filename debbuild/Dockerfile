FROM ubuntu:xenial
MAINTAINER Sandeep Sangamreddi <sandeepone@gmail.com>

# Volumes
VOLUME /build  
VOLUME /release

# Install build dependencies
RUN apt-get update && apt-get -y install \  
    build-essential software-properties-common \
    devscripts \
    fakeroot \
    debhelper \
    automake make \
    autotools-dev \
    pkg-config \
    git htop nano\
    ca-certificates \
	libpcre3 libpcre3-dev \
	cmake gawk libmysqlclient-dev \
	libgeoip1 libgeoip-dev \
	libhiredis-dev quilt libperl-dev \
	google-perftools libgoogle-perftools-dev \
  	libxml2-dev libmcrypt-dev libicu-dev openssl binutils-dev \
  	libcap-dev zlib1g-dev libtbb-dev libonig-dev libpcre3-dev \
  	autoconf libtool libcurl4-openssl-dev wget memcached \
  	libreadline-dev libncurses5-dev libmemcached-dev libbz2-dev \
  	libc-client2007e-dev libgoogle-perftools-dev \
  	libcloog-ppl-dev libelf-dev libdwarf-dev libunwind8-dev subversion \
  	libarchive13 cmake-data libacl1 libattr1 libgd2-xpm-dev \
  	pkg-config libmagickwand-dev libiberty-dev libevent-dev libxslt-dev libgoogle-glog-dev \
  	automake libldap2-dev libkrb5-dev libyaml-dev gperf ocaml-native-compilers libpq-dev \
	dh-systemd libluajit-5.1-2 libluajit-5.1-common libluajit-5.1-dev libmhash-dev libmhash2 libperl-dev \
    --no-install-recommends

# overwrite this with 'CMD []' in a dependent Dockerfile
CMD ["/bin/bash"]