FROM ubuntu:xenial
MAINTAINER Sandeep Sangamreddi <sandeepone@gmail.com>

RUN apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0x5a16e7281be7a449 \
  && apt-get update -y && apt-get install -y software-properties-common \
  && add-apt-repository "deb http://dl.hhvm.com/ubuntu $(lsb_release -sc) main" \
  && apt-get update -y \
  && apt-get install -y hhvm=3.17.1~$(lsb_release -sc) \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*