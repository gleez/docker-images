Ubuntu Linux HHVM Dockerfile
=============

This repository contains Dockerfile of HHVM running on Ubuntu Linux for Docker's automated build published to the public Docker Hub Registry.

The Base Image
---------------------

This is built on top of `ubuntu:16.04`, so make sure you're building against
the latest version of that:

```
$ docker pull ubuntu:xenial
```

# Installation

## Install Docker.

Download automated build from public Docker Hub Registry: docker pull gleez/hhvm

(alternatively, you can build an image from Dockerfile: docker build -t="gleez/hhvm" github.com/gleez/docker-images/hhvm)

## Usage

    docker run -it -d gleez/hhvm
    
You can add a workspace as a volume directory with the argument *-v /your-path/workspace/:/workspace/* like this :

	docker run -it --rm --name my-running-script -v "$PWD":/workspace gleez/hhvm hhvm your-script.php
    
## Build and run with custom config directory

Get the latest version from github

    git clone https://github.com/gleez/docker-images
    cd docker-images/hhvm

Build it

    sudo docker build --force-rm=true --tag="$USER/hhvm:latest" .
    
And run

    sudo docker run -d -v /your-path/workspace/:/workspace/ $USER/hhvm:latest

    
Enjoy !!

## License

[MIT](http://opensource.org/licenses/MIT)
