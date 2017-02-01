FROM mysql:5.7
MAINTAINER Sandeep Sangamreddi <sandeepone@gmail.com>
 
RUN \
    echo "deb http://repo.percona.com/apt jessie main" > /etc/apt/sources.list.d/percona-release.list \
    && apt-key adv --keyserver keys.gnupg.net --recv-keys 1C4CBDCDCD2EFD2A \
    && apt-get update \
	&& apt-get install --no-install-recommends --no-install-suggests -y --force-yes \
	percona-xtrabackup-24 percona-toolkit mytop \
	&& mkdir -p /backups \
	&& apt-get -y autoremove && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Allow mountable backup path
VOLUME ["/backups"]

# Copy the script to simplify backup command
COPY xtrabackup.sh /xtrabackup.sh
