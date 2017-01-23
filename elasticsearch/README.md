# elasticsearch
Elasticsearch docker image of smaller than official version.

## Base Image
- anapsix/alpine-java:jdk8

## Environment
- Alpine Linux 3.4
- Oracle Java 8 JDK
- Elasticsearch 5.1.2

#### Version: 5.1.2, Build: c8c4c16/2017-01-11T20:18:39.146Z, JVM: 1.8.0_121

## Running Containers

You can run the default `elasticsearch` command simply:

```console
$ docker run -d gleez/elasticsearch
```

You can also pass in additional flags to `elasticsearch`:

```console
$ docker run -d gleez/elasticsearch -Des.node.name="TestNode"
```

This image comes with a default set of configuration files for `elasticsearch`, but if you want to provide your own set of configuration files, you can do so via a volume mounted at `/usr/share/elasticsearch/config`:

```console
$ docker run -d -v "$PWD/config":/usr/share/elasticsearch/config gleez/elasticsearch
```

This image is configured with a volume at `/usr/share/elasticsearch/data` to hold the persisted index data. Use that path if you would like to keep the data in a mounted volume:

```console
$ docker run -d -v "$PWD/esdata":/usr/share/elasticsearch/data gleez/elasticsearch
```

This image includes `EXPOSE 9200 9300` ([default `http.port`](http://www.elastic.co/guide/en/elasticsearch/reference/1.5/modules-http.html)), so standard container linking will make it automatically available to the linked containers.



## License

[MIT](http://opensource.org/licenses/MIT)
