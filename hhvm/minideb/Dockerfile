FROM bitnami/minideb:jessie

RUN apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0x5a16e7281be7a449 \
  && apt-get update -y && apt-get install -y --no-install-recommends software-properties-common \
  && echo deb http://dl.hhvm.com/debian jessie main | tee /etc/apt/sources.list.d/hhvm.list \
  && apt-get update -y \
  && apt-get install -y --no-install-recommends hhvm \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

