# goStatic
A really small static web server for Docker


### The goal
My goal is to create to smallest docker container for my web static files. The advantage of Go is that you can generate a fully static binary, so that you don't need anything else.

### Features
 * A fully static web server in 5MB
 * Web server build for Docker
 * HTTPS by default
 * TLS certificates for a domain automatically from Let’s Encrypt
 * Light container
 * More security than official images (see below)
 * Log enabled

### Why?
Because the official Golang image is wayyyy to big (around 1/2Gb as you can see below) and could be unsecure.


For me, the whole point of containers is to have a light container...
Many links should provide you with additionnal info to see my point of view:

 * [Over 30% of Official Images in Docker Hub Contain High Priority Security Vulnerabilities](http://www.banyanops.com/blog/analyzing-docker-hub/)
 * [Create The Smallest Possible Docker Container](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/)
 * [Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07)
 * [Small Docker Images For Go Apps](http://www.centurylinklabs.com/small-docker-images-for-go-apps/)
 * [https://github.com/PierreZ/goStatic] (https://github.com/PierreZ/goStatic) -- forked and modified


### How to use
```
// HTTPS server
docker run -d -p 443:443 -v path/to/website:/srv/http --name goStatic gleez/gostatic

// HTTP server
docker run -d -p 80:80 -v path/to/website:/srv/http --name goStatic gleez/gostatic --forceHTTP
```

#### For Auto TLS
https://echo.labstack.com/cookbook/auto-tls

Needs this folder mounted "/var/www/.cache"


### Wow, such container! What are you using?

I'm using [echo](http://echo.labstack.com/) as a  micro web framework because he has great performance, and generate the static binary

I'm also using the https://github.com/broady/cacerts for the certs to avoid this error:

```
x509: failed to load system roots and no roots provided
```

### debian_cpcerts.bash

```
set -e -x

apt-get update
apt-get install -y ca-certificates
cp /etc/ssl/certs/ca-certificates.crt .
```

### Build

Get the latest version from github

    git clone https://github.com/gleez/docker-images
    cd docker-images/goStatic

Build it

```
sudo docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.7.5 bash
```

```
$ ./debian_cpcerts.bash
$ CGO_ENABLED=0 go build -ldflags '-s' -v -o goStatic
```

#### Finally

```
sudo docker build --force-rm=true --tag="$USER/gostatic" .
```

Enjoy !!

## License

[MIT](http://opensource.org/licenses/MIT)