Percona XtraDB Cluster docker image
===================================

The docker image is available right now at `percona/percona-xtradb-cluster:5.7`.
The image supports work in Docker Network, including overlay networks,
so that you can install Percona XtraDB Cluster nodes on different boxes.
There is an initial support for the etcd discovery service.

Basic usage
-----------

For an example, see the `start_node.sh` script.

The `CLUSTER_NAME` environment variable should be set, and the easiest to do it is:
`export CLUSTER_NAME=cluster1`

The script will try to create an overlay network `${CLUSTER_NAME}_net`.
If you want to have a bridge network or network with a specific parameter,
create it in advance.
For example:
`docker network create -d bridge ${CLUSTER_NAME}_net`

The Docker image accepts the following parameters:
* One of `MYSQL_ROOT_PASSWORD_FILE`, `MYSQL_ROOT_PASSWORD`, `MYSQL_ALLOW_EMPTY_PASSWORD` or `MYSQL_RANDOM_ROOT_PASSWORD` must be defined
* The image will create the user `xtrabackup@localhost` for the XtraBackup SST method. If you want to use a password for the `xtrabackup` user, set `XTRABACKUP_PASSWORD`. 
* If you want to use the discovery service (right now only `etcd` is supported), set the address to `DISCOVERY_SERVICE`. The image will automatically find a running cluser by `CLUSTER_NAME` and join to the existing cluster (or start a new one).
* If you want to start without the discovery service, use the `CLUSTER_JOIN` variable. Empty variables will start a new cluster, To join an existing cluster, set `CLUSTER_JOIN` to the list of IP addresses running cluster nodes.

+## `MYSQL_ROOT_PASSWORD_FILE`
+
+This variable specifies a file that will be read for the root user account. This can be a mounted file when you run your container. This can also be used in the scope of the Docker Secrets (Swarm mode) functionality.
+



Running a Docker overlay network
------------------------------

The following link is a great introduction with easy steps on how to run a Docker overlay network: http://chunqi.li/2015/11/09/docker-multi-host-networking/


Running with ProxySQL
---------------------

The ProxySQL image https://hub.docker.com/r/perconalab/proxysql/
provides an integration with Percona XtraDB Cluster and discovery service.

You can start proxysql image by
```
docker run -d -p 3306:3306 -p 6032:6032 --net=$NETWORK_NAME --name=${CLUSTER_NAME}_proxysql \
        -e CLUSTER_NAME=$CLUSTER_NAME \
        -e ETCD_HOST=$ETCD_HOST \
        -e MYSQL_ROOT_PASSWORD=Theistareyk \
        -e MYSQL_PROXY_USER=proxyuser \
        -e MYSQL_PROXY_PASSWORD=s3cret \
        perconalab/proxysql
```

where `MYSQL_ROOT_PASSWORD` is the root password for the MySQL nodes. The password is needed to register the proxy user. The user `MYSQL_PROXY_USER` with password `MYSQL_PROXY_PASSWORD` will be registered on all Percona XtraDB Cluster nodes.


Running `docker exec -it ${CLUSTER_NAME}_proxysql add_cluster_nodes.sh` will register all nodes in the ProxySQL.



TODO 
    
    dont create xtrabackup user if not in cluster mode
    allow setting up mysql config variables at runtime using env variabes


    HEALTCHECK during initial mysql bootstrap shuts down the container before is syncs with the other nodes
    issue - https://github.com/docker/docker/issues/26664
    dig +short tasks.serviceName - no way to get which service am I part of
    docker returns tasks with dig even if they are still starting(unhealthy) - dig +short tasks.serviceName
    when scaling up start tasks Consecutively - https://github.com/docker/docker/issues/30194
    containers loose connectivity between each other after some time


RUN 
docker service create \
--publish 3306:3306 \
--name mysql \
--replicas 1 \
--network mysql_net \
-e MYSQL_ROOT_PASSWORD=1 \
-e CLUSTER_NAME=cl  \
-e SERVICE_ID={{.Service.ID}} \
-e SERVICE_NAME={{.Service.Name}} \
-e TASK_ID={{.Task.ID}} \
-e TASK_NAME={{.Task.Name}} \
-e TASK_SLOT={{.Task.Slot}} \
--mount type=bind,src=/home/apps/volumes/mysql/{{.Task.Slot}},dst=/var/lib/mysql \
--mount type=bind,src=/home/apps/log/mysql/{{.Task.Slot}},dst=/var/log/mysql \
gleez/mysql:5.7 bash -c "(/entrypoint.sh mysqld &) && while true; do  dig +short tasks.mysql | xargs -n1 ping -w1 -c1 ; sleep 2; done"
