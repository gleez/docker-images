Ubuntu Linux Dockerfile
=============

This repository contains Dockerfile of deb builder running on Ubuntu Linux for building deb packages in Ubuntu container.

The Base Image
---------------------

This is built on top of `ubuntu:16.04`, so make sure you're building against
the latest version of that:

```
$ docker pull gleez/debbuild
```

## Usage

    docker run -it gleez/debbuild /bin/bash
    docker run -v /home/gleez/Projects/Docker/nginx/src:/build -it gleez/debbuild /bin/bash
    
You can add a workspace as a volume directory with the argument *-v /your-path/build/:/build/* like this :

	docker run -it -v "$PWD":/build gleez/debbuild /bin/bash
    
## Build and run with custom config directory

Get the latest version from github

    git clone https://github.com/gleez/docker-images
    cd docker-images/debbuild

Build it

    sudo  docker build --force-rm=true --tag="$USER/debbuild" .


    
Enjoy !!

## License

[MIT](http://opensource.org/licenses/MIT)
