FROM alpine:3.6

ARG BUILD_CORES

ARG PHP_VER=7.1.11
ARG LIBICONV_VERSION=1.15

LABEL description="PHP image based on Alpine" \
      maintainer="Sandeepone <sandeepone@gmail.com>" \
      php_version="PHP v$PHP_VER built from source" \
      repository.hub="gleez/php" \
      repository.url="https://github.com/gleez/docker-images"

ARG PHP_MIRROR=http://ch1.php.net

ARG PHP_CONF=" \
    --prefix=/usr \
    --libdir=/usr/lib/php \
    --datadir=/usr/share/php \
    --sysconfdir=/php/etc \
    --localstatedir=/php/var \
    --with-pear=/usr/share/php \
    --with-config-file-scan-dir=/usr/local/etc/php/conf.d \
    --with-config-file-path=/usr/local/etc/php \
    --with-pic \
    --with-pcre-regex=/usr \
    --disable-short-tags \
    --without-readline \
    --enable-bcmath=shared \
    --enable-fpm \
    --enable-ftp \
    --disable-cgi \
    --enable-mysqlnd \
    --enable-mbstring \
    --with-curl \
    --with-libedit \
    --with-openssl \
    --with-iconv=/usr/local \
    --with-gd \
    --with-jpeg-dir \
    --with-png-dir \
    --with-webp-dir \
    --with-xpm-dir=no \
    --with-freetype-dir \
    --enable-gd-native-ttf \
    --disable-gd-jis-conv \
    --enable-libxml \
    --enable-xml \
    --enable-soap \
    --enable-shmop \
    --with-libzip \
    --with-imap-ssl \
    --with-zlib"

ARG PHP_EXT_LIST=" \
    mysqli \
    gettext \
    bcmath \
    calendar \
    pdo_pgsql \
    pdo_mysql \
    pcntl \
    bz2 \
    enchant \
    intl \
    opcache \
    mcrypt \
    ldap \
    pgsql \
    pspell \
    exif \
    gmp \
    tidy \
    snmp \
    sockets \
    sysvmsg \
    sysvsem \
    sysvshm \
    wddx \
    xmlrpc \
    xsl \
    imap"

ARG CUSTOM_BUILD_PKGS=" \
    freetype-dev \
    openldap-dev \
    gmp-dev \
    libmcrypt-dev \
    icu-dev \
    postgresql-dev \
    libpng-dev \
    libwebp-dev \
    gd-dev \
    libjpeg-turbo-dev \
    libxpm-dev \
    libedit-dev \
    libxml2-dev \
    libressl-dev \
    libbz2 \
    sqlite-dev \
    gettext-dev \
    tidyhtml-dev \
    libzip-dev \
    libxslt-dev \
    libsodium-dev \
    net-snmp-dev \
    tzdata \
    geoip-dev \
    imagemagick-dev \
    libmemcached-dev \
    aspell-dev \
    enchant-dev \
    libevent-dev \
    nghttp2-dev \
    hiredis-dev \
    imap-dev"

ARG CUSTOM_PKGS=" \
    freetype \
    openldap \
    gmp \
    libmcrypt \
    bzip2-dev \
    icu \
    libpq \
    gettext \
    tidyhtml-libs \
    libzip \
    libxslt \
    libsodium \
    net-snmp-libs \
    geoip \
    imagemagick \
    libmemcached-libs \
    aspell-libs \
    aspell-en \
    enchant \
    libevent \
    nghttp2-libs \
    hiredis \
    c-client"

COPY rootfs /

# Apply stack smash protection to functions using local buffers and alloca()
# Make PHP's main executable position-independent (improves ASLR security mechanism, and has no performance impact on x86_64)
# Enable optimization (-O2)
# Enable linker optimization (this sorts the hash buckets to improve cache locality, and is non-default)
# Adds GNU HASH segments to generated executables (this is used if present, and is much faster than sysv hash; in this configuration, sysv hash is also generated)
# https://github.com/docker-library/php/issues/272
ENV PHP_CFLAGS="-fstack-protector-strong -fpic -fpie -O2"
ENV PHP_CPPFLAGS="$PHP_CFLAGS"
ENV PHP_LDFLAGS="-Wl,-O1 -Wl,--hash-style=both -pie"


RUN NB_CORES=${BUILD_CORES-$(getconf _NPROCESSORS_CONF)} \

### Packages installation
  && BUILD_DEPS=" \
    linux-headers \
    libtool \
    build-base \
    pcre-dev \
    zlib-dev \
    wget \
    gnupg \
    autoconf \
    gcc \
    g++ \
    libc-dev \
    make \
    pkgconf \
    re2c \
    curl-dev \
    dpkg-dev dpkg \
    ${CUSTOM_BUILD_PKGS}" \

  && apk -U add \
    ${BUILD_DEPS} \
    s6 \
    su-exec \
    curl \
    ca-certificates \
    libedit \
    libxml2 \
    libressl \
    libwebp \
    gd \
    pcre \
    zlib \
    tar \
    xz \
    ${CUSTOM_PKGS} \

### Source downloading
  && wget ${PHP_MIRROR}/get/php-${PHP_VER}.tar.gz/from/this/mirror -O /tmp/php-${PHP_VER}.tar.gz \
  && wget ${PHP_MIRROR}/get/php-${PHP_VER}.tar.gz.asc/from/this/mirror -O /tmp/php-${PHP_VER}.tar.gz.asc \
  && wget http://ftp.gnu.org/pub/gnu/libiconv/libiconv-${LIBICONV_VERSION}.tar.gz -O /tmp/libiconv-${LIBICONV_VERSION}.tar.gz \
  && mkdir -p /usr/local/etc/php/conf.d \
  && mkdir -p /usr/src \
  && tar xzvf /tmp/php-${PHP_VER}.tar.gz -C /usr/src \
  && tar xzf /tmp/libiconv-${LIBICONV_VERSION}.tar.gz -C /usr/src \

### GNU Libiconv installation
  && cd /usr/src/libiconv-${LIBICONV_VERSION} \
  && ./configure --prefix=/usr/local \
  && make && make install && libtool --finish /usr/local/lib \

### PHP installation
  && mv /usr/src/php-${PHP_VER} /usr/src/php \
  && cd /usr/src/php \
  && export CFLAGS="$PHP_CFLAGS" \
   CPPFLAGS="$PHP_CPPFLAGS" \
   LDFLAGS="$PHP_LDFLAGS" \
  && gnuArch="$(dpkg-architecture --query DEB_BUILD_GNU_TYPE)" \
  && ./configure ${PHP_CONF} \
  && make -j ${NB_CORES} \
  && make install \

### Strip, clean, install modules
  && { find /usr/local/bin /usr/local/sbin -type f -perm +0111 -exec strip --strip-all '{}' + || true; } \
  && make clean \
  && chmod u+x /usr/local/bin/* /etc/s6.d/*/* \
  && docker-php-ext-install ${PHP_EXT_LIST} \
  && cp /usr/src/php/php.ini-production /usr/local/etc/php/php.ini \
##  && pear config-set php_ini /usr/local/etc/php/php.ini \
##  && pecl config-set php_ini /usr/local/etc/php/php.ini \
  && pecl channel-update pecl.php.net \
  && pecl install apcu memcached redis mongodb imagick geoip-beta swoole event xdebug hprose \
  && docker-php-ext-enable apcu memcached redis mongodb imagick geoip swoole event hprose \
  && pecl clear-cache \

 ### Sockets ext needs to be loaded before event ext
  && mv /usr/local/etc/php/conf.d/docker-php-ext-event.ini /usr/local/etc/php/conf.d/docker-php-ext-zz-event.ini \

### Add nginx user & group (docker run -t -i gleez/nginx:1.12.0 id nginx)
  && addgroup -S nginx \
	&& adduser -D -S -h /var/cache/nginx -s /sbin/nologin -G nginx nginx \
  && addgroup -g 3000 -S app \
	&& adduser  -u 3000 -D -S -h /var/www -G app app \

  && mkdir /usr/local/etc/php-fpm.d \
  && { \
      echo '[global]'; \
      echo 'include=/usr/local/etc/php-fpm.d/*.conf'; \
} | tee /usr/local/etc/php-fpm.conf \
  && { \
    echo '[global]'; \
    echo 'error_log = /proc/self/fd/2'; \
    echo 'daemonize = no'; \
} | tee /usr/local/etc/docker.conf \
  && { \
    echo '[www]'; \
    echo '; if we send this to /proc/self/fd/1, it never appears'; \
    echo 'listen = [::]:9000'; \
    echo 'user = nginx'; \
    echo 'group = nginx'; \
    echo 'chdir = /var/www/html'; \
    echo; \
    echo 'pm = dynamic'; \
    echo 'pm.max_children = 9'; \
    echo 'pm.start_servers = 3'; \
    echo 'pm.max_spare_servers = 4'; \
    echo 'pm.min_spare_servers = 2'; \
    echo 'pm.max_requests = 20'; \
    echo; \
    echo 'request_terminate_timeout = 0'; \
    echo 'request_slowlog_timeout = 1s'; \
    echo 'slowlog = /proc/self/fd/2'; \
    echo '; Ensure worker stdout and stderr are sent to the main error log.'; \
    echo 'catch_workers_output = yes'; \
    echo; \
    echo 'access.log = /proc/self/fd/2'; \
    echo 'clear_env = no'; \
    echo 'php_flag[display_errors] = on'; \
    echo 'php_admin_value[expose_php] = Off'; \
    echo 'php_admin_value[memory_limit] = 128M'; \
} | tee /usr/local/etc/php-fpm.d/zz-docker.conf \
  && sed -i \
    -e "s|upload_max_filesize\s*=.*|upload_max_filesize = 99M|" \
    -e "s|max_file_uploads\s*=.*|max_file_uploads = 20|" \
    -e "s|post_max_size\s*=.*|post_max_size = 99M|" \
    -e "s|;cgi.fix_pathinfo\s*=.*|cgi.fix_pathinfo = 1|" \
    -e "s|;date.timezone\s*=.*|date.timezone = UTC|" \
    -e "s|mail.add_x_header\s*=.*|mail.add_x_header = Off|" \
    -e "s|expose_php\s*=.*|expose_php = Off|" \
    -e "s|;opcache.enable\s*=.*|opcache.enable = 1|" \
    -e "s|;opcache.enable_cli\s*=.*|opcache.enable_cli = 1|" \
    -e "s|;opcache.memory_consumption\s*=.*|opcache.memory_consumption = 192|" \
    -e "s|;opcache.max_accelerated_files\s*=.*|opcache.max_accelerated_files = 5000|" \
    -e "s|;opcache.revalidate_freq\s*=.*|opcache.revalidate_freq = 60|" \
    -e "s|;opcache.validate_timestamps\s*=.*|opcache.validate_timestamps = 1|" \
    -e "s|;opcache.interned_strings_buffer\s*=.*|opcache.interned_strings_buffer = 16|" \
    -e "s|;opcache.fast_shutdown\s*=.*|opcache.fast_shutdown = 1|" \
    /usr/local/etc/php/php.ini \

## Disable Xdebug
##  && rm /usr/local/etc/php/conf.d/docker-php-ext-xdebug.ini \

  && sed -i -e "s|extension=hprose.so|;extension=hprose.so|" /usr/local/etc/php/conf.d/docker-php-ext-hprose.ini \
  && sed -i -e "s|extension=swoole.so|;extension=swoole.so|" /usr/local/etc/php/conf.d/docker-php-ext-swoole.ini \
  && sed -i -e "s|extension=event.so|;extension=event.so|" /usr/local/etc/php/conf.d/docker-php-ext-zz-event.ini \

## Install PHPUnit && Composer
  && curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer \
  && curl -sSL https://phar.phpunit.de/phpunit.phar -o phpunit.phar \
  && chmod +x phpunit.phar && mv phpunit.phar /usr/local/bin/phpunit \

# Cleanup
  && apk del ${BUILD_DEPS} \
  && rm -rf /tmp/* /var/cache/apk/* /usr/src/* \
     /usr/include/php7 \
     /usr/lib/php7/build \
     /usr/lib/php7/modules/*.a \
     /usr/share/man \
     /usr/share/doc \
     /usr/share/info \
     /usr/share/locale

EXPOSE 9000
WORKDIR /var/www/html
