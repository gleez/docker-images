FROM debian:jessie

RUN apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0x5a16e7281be7a449 \
  && echo deb http://dl.hhvm.com/debian jessie main | tee /etc/apt/sources.list.d/hhvm.list \

## mysql 5.7
# && export GNUPGHOME="$(mktemp -d)" \
# && gpg --keyserver ha.pool.sks-keyservers.net --recv-keys A4A9406876FCBD3C456770C88C718D3B5072E1F5 \
# && gpg --export A4A9406876FCBD3C456770C88C718D3B5072E1F5 > /etc/apt/trusted.gpg.d/mysql.gpg \
# && echo deb http://repo.mysql.com/apt/debian/ jessie mysql-5.7 > /etc/apt/sources.list.d/mysql.list \

  && apt-get update -y \
  && apt-get install -y --no-install-recommends hhvm=3.21.0~jessie \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

