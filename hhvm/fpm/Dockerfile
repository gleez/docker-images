FROM gleez/hhvm:3.20.2

RUN mkdir -p /var/www

ADD server.ini /etc/hhvm/server.ini

EXPOSE 9000

CMD ["/usr/bin/hhvm", "-m", "server", "-c", "/etc/hhvm/server.ini"]
