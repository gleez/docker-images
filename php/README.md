Alpine Linux PHP Dockerfile
=============

This repository contains Dockerfile of PHP running on Alpine Linux for Docker's automated build published to the public Docker Hub Registry.

### PHP Extensions optional
* ldap
* pgsql
* mongo
* redis
* geoip
* imagick
* memcache
* apcu
* xdebug
* event
* swoole
* hprose

The Base Image
---------------------

This is built on top of `alpine:3.6`, so make sure you're building against
the latest version of that:

```
$ docker pull alpine:3.6
```

# Installation

## Install Docker.

Download automated build from public Docker Hub Registry: docker pull gleez/php

(alternatively, you can build an image from Dockerfile: docker build -t="gleez/php" github.com/gleez/docker-images/php)

## Usage

    docker run -it -d gleez/php

You can add a workspace as a volume directory with the argument *-v /your-path/workspace/:/var/www/html/* like this :

	docker run -it --rm --name my-running-script -v "$PWD":/workspace gleez/php php your-script.php

## Build and run with custom config directory

	Get the latest version from github

	    git clone https://github.com/gleez/docker-images
	    cd docker-images/php

	Build it

	    sudo docker build --force-rm=true --tag="$USER/php:latest" .

	And run

	    sudo docker run -d -v /your-path/workspace/:/var/www/html/ $USER/php:latest


	Enjoy !!

## License

	[MIT](http://opensource.org/licenses/MIT)
