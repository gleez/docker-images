# docker-grimd

## Description

[Grimd](https://github.com/looterz/grimd) is a fast dns proxy, built to black-hole internet advertisements/malware servers.

## Usage

The recommended way to run this container is as follows:

from alpine image:
```bash
$ docker run -d --name grimd -p 53:53/tcp -p 53:53/udp gleez/grimd
```

## Parameters

* `-d` - as daemon **required**
* `--name grimd` - container name
* `-p 53:53/tcp` - dns port **required**
* `-p 53:53/udp` - dns port **required**
* `-p 8080:8080/tcp` - api port
* --cap-add=NET_ADMIN
* `-v <path to config>:/config` - where grimd read [grimd.toml](https://raw.githubusercontent.com/gleez/docker-images/master/grimd/grimd.toml) file

## Security

You should allow connections only from clients you trust. More information [here](https://github.com/looterz/grimd/wiki/Securing-on-linux).

### Build

Get the latest version from github

    git clone https://github.com/gleez/docker-images
    cd docker-images/grimd

Build it

```
sudo docker run --rm -it -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest bash
```

```
$ ./debian_cpcerts.bash
$ CGO_ENABLED=0 go build -ldflags '-s' -v -o grimd
```

#### Finally

```
sudo docker build --force-rm=true --tag="$USER/grimd" .
```

