
### Alpine image for Tox (toxcore) DHT bootstrap daemon

If you are familiar with Docker and would rather run the daemon in a Docker container, run the following from this directory:

```sh
sudo docker build --force-rm=true --tag="$USER/tox-bootstrapd" .

sudo useradd --home-dir /var/lib/tox-bootstrapd --create-home --system --shell /sbin/nologin --comment "Account to run Tox's DHT bootstrap daemon" --user-group tox-bootstrapd
sudo chmod 700 /var/lib/tox-bootstrapd

sudo docker run -d --name tox-bootstrapd --restart always -v /var/lib/tox-bootstrapd/:/var/lib/tox-bootstrapd/ -p 443:443 -p 3389:3389 -p 33445:33445 -p 33445:33445/udp gleez/tox-bootstrapd
```

We create a new user and protect its home directory in order to mount it in the Docker image, so that the kyepair the daemon uses would be stored on the host system, which makes it less likely that you would loose the keypair while playing with or updating the Docker container.

You can check logs for your public key or any errors:
```sh
sudo docker logs tox-bootstrapd
```

Note that the Docker container runs a script which pulls a list of bootstrap nodes off https://nodes.tox.chat/ and adds them in the config file.

### Updating

You want to make sure that the daemon uses the newest toxcore, as there might have been some changes done to the DHT, so it's advised to update the daemon at least once every month.

To update the daemon, all you need is to erase current container with its image:

```sh
sudo docker stop tox-bootstrapd
sudo docker rm tox-bootstrapd
sudo docker rmi tox-bootstrapd
```

Then rebuild and run the image again:

```sh
sudo docker build --force-rm=true --tag="$USER/tox-bootstrapd" .
sudo docker run -d --name tox-bootstrapd --restart always -v /var/lib/tox-bootstrapd/:/var/lib/tox-bootstrapd/ -p 443:443 -p 3389:3389 -p 33445:33445 -p 33445:33445/udp gleez/tox-bootstrapd
```

### Troubleshooting

- Check if the container is running:
```sh
sudo docker ps -a
```

- Check the log for errors:
```sh
sudo docker logs tox-bootstrapd
```