Alpine Linux Cloud9 v3 Dockerfile
=============

This repository contains Dockerfile of Cloud9 IDE running on Alpine Linux for Docker's automated build published to the public Docker Hub Registry.

# Base Docker Image
[alpine:3.4](https://github.com/gliderlabs/docker-alpine)

# Installation

## Install Docker.

Download automated build from public Docker Hub Registry: docker pull gleez/alpine-cloud9

(alternatively, you can build an image from Dockerfile: docker build -t="gleez/cloud9" github.com/gleez/docker-images/cloud9)

## Usage

    docker run -it -d -p 80:3000 gleez/cloud9
    
You can add a workspace as a volume directory with the argument *-v /your-path/workspace/:/workspace/* like this :

    docker run -it -d -p 80:3000 -v /your-path/workspace/:/workspace/ gleez/cloud9
    
## Build and run with custom config directory

Get the latest version from github

    git clone https://github.com/gleez/docker-images
    cd docker-images/cloud9

Build it

    sudo docker build --force-rm=true --tag="$USER/cloud9:latest" .
    
And run

    sudo docker run -d -p 80:3000 -v /your-path/workspace/:/workspace/ $USER/cloud9:latest

or

	sudo docker run -d --net=host -v /your-path/workspace/:/workspace/ $USER/cloud9:latest
    
Enjoy !!

## License

[MIT](http://opensource.org/licenses/MIT)
