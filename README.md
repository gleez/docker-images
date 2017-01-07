Alpine Linux Cloud9 v3 Dockerfile
=============

This repository contains Dockerfile of Cloud9 IDE running on Alpine Linux for Docker's automated build published to the public Docker Hub Registry.

# Base Docker Image
[nodejs/docker-node](https://github.com/nodejs/docker-node)

# Installation

## Install Docker.

Download automated build from public Docker Hub Registry: docker pull gleez/alpine-cloud9

(alternatively, you can build an image from Dockerfile: docker build -t="gleez/alpine-cloud9" github.com/gleez/cloud9-docker)

## Usage

    docker run -it -d -p 80:3000 gleez/alpine-cloud9
    
You can add a workspace as a volume directory with the argument *-v /your-path/workspace/:/workspace/* like this :

    docker run -it -d -p 80:3000 -v /your-path/workspace/:/workspace/ gleez/alpine-cloud9
    
## Build and run with custom config directory

Get the latest version from github

    git clone https://github.com/gleez/alpine-cloud9
    cd cloud9-docker/

Build it

    sudo docker build --force-rm=true --tag="$USER/alpine-cloud9:latest" .
    
And run

    sudo docker run -d -p 80:3000 -v /your-path/workspace/:/workspace/ $USER/alpine-cloud9:latest

or

	sudo docker run -d --net=host -v /your-path/workspace/:/workspace/ $USER/alpine-cloud9:latest
    
Enjoy !!    
