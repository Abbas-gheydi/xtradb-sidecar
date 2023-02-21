FROM perconalab/percona-xtrabackup:8.0.30
USER root
RUN export HTTPS_PROXY=http://10.21.13.142:3306 && microdnf install openssh-clients
COPY xtradb-sidecar /bin/
CMD ["xtradb-sidecar"]

