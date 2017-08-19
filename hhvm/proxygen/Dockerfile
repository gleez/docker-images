FROM gleez/hhvm:3.21.0

RUN mkdir -p /var/www

ADD server.ini /etc/hhvm/server.ini

EXPOSE 80

CMD ["/usr/bin/hhvm", "-m", "server", "-c", "/etc/hhvm/server.ini"]
