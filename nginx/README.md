# Supported tags and respective `Dockerfile` links

-	[`1.10.2` (*nginx/Dockerfile*)](https://github.com/gleez/docker-images/blob/master/nginx/Dockerfile)

### Nginx  with OpenSSL 1.1.0c

# Modules

-	Added ngx_brotli.
-	Added ngx_cache_purge.
-	Added ngx_dav_ext.
-	Added ngx_headers_more_filter.
-	Added ngx_lua.
-	Added ngx_redis2.
-	Added ngx_srcache_filter.
-	Added ngx_substitutions_filter.
-	Added ngx_uploadprogress.
-	Added ngx_upstream_fair.
-	Added ngx_nchan.
-	Added Pagespeed 1.12.34.2-beta.

Note all are dynamic modules. check `/etc/nginx/modules`

# What is Nginx?

Nginx (pronounced "engine-x") is an open source reverse proxy server for HTTP, HTTPS, SMTP, POP3, and IMAP protocols, as well as a load balancer, HTTP cache, and a web server (origin server). The nginx project started with a strong focus on high concurrency, high performance and low memory usage. It is licensed under the 2-clause BSD-like license and it runs on Linux, BSD variants, Mac OS X, Solaris, AIX, HP-UX, as well as on other *nix flavors. It also has a proof of concept port for Microsoft Windows.

> [wikipedia.org/wiki/Nginx](https://en.wikipedia.org/wiki/Nginx)

![logo](https://raw.githubusercontent.com/docker-library/docs/01c12653951b2fe592c1f93a13b4e289ada0e3a1/nginx/logo.png)

# How to use this image

## hosting some simple static content

```console
$ docker run --name some-nginx -v /some/content:/usr/share/nginx/html:ro -d gleez/nginx:1.10.2
```

Alternatively, a simple `Dockerfile` can be used to generate a new image that includes the necessary content (which is a much cleaner solution than the bind mount above):

```dockerfile
FROM nginx
COPY static-html-directory /usr/share/nginx/html
```

Place this file in the same directory as your directory of content ("static-html-directory"), run `docker build -t some-content-nginx .`, then start your container:

```console
$ docker run --name some-nginx -d some-content-nginx
```

## exposing the port

```console
$ docker run --name some-nginx -d -p 8080:80 some-content-nginx
```

Then you can hit `http://localhost:8080` or `http://host-ip:8080` in your browser.

## complex configuration

```console
$ docker run --name some-nginx -v /some/nginx.conf:/etc/nginx/nginx.conf:ro -d gleez/nginx:1.10.2
```

For information on the syntax of the Nginx configuration files, see [the official documentation](http://nginx.org/en/docs/) (specifically the [Beginner's Guide](http://nginx.org/en/docs/beginners_guide.html#conf_structure)).

Be sure to include `daemon off;` in your custom configuration to ensure that Nginx stays in the foreground so that Docker can track the process properly (otherwise your container will stop immediately after starting)!

If you wish to adapt the default configuration, use something like the following to copy it from a running Nginx container:

```console
$ docker cp some-nginx:/etc/nginx/nginx.conf /some/nginx.conf
```

As above, this can also be accomplished more cleanly using a simple `Dockerfile`:

```dockerfile
FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf
```

Then, build with `docker build -t some-custom-nginx .` and run:

```console
$ docker run --name some-nginx -d some-custom-nginx
```

### using environment variables in nginx configuration

Out-of-the-box, Nginx doesn't support using environment variables inside most configuration blocks. But `envsubst` may be used as a workaround if you need to generate your nginx configuration dynamically before nginx starts.

Here is an example using docker-compose.yml:

```web:
  image: gleez/nginx:1.10.2
  volumes:
   - ./mysite.template:/etc/nginx/conf.d/mysite.template
  ports:
   - "8080:80"
  environment:
   - NGINX_HOST=foobar.com
   - NGINX_PORT=80
  command: /bin/bash -c "envsubst < /etc/nginx/conf.d/mysite.template > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'"
```

The `mysite.template` file may then contain variable references like this :

`listen       ${NGINX_PORT};
`

# Image Variants

The `nginx` images come in many flavors, each designed for a specific use case.

## `gleez/nginx:1.10.2`

This is the defacto image. If you are unsure about what your needs are, you probably want to use this one. It is designed to be used both as a throw away container (mount your source code and start the container to start your app), as well as the base to build other images off of.

# Supported Docker versions

This image is officially supported on Docker version 1.13.0.

Support for older versions (down to 1.6) is provided on a best-effort basis.

Please see [the Docker installation documentation](https://docs.docker.com/installation/) for details on how to upgrade your Docker daemon.
