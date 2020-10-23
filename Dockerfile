FROM alpinelinux/docker-cli

ADD puller /usr/local/bin/puller
RUN chmod +x /usr/local/bin/puller

ENTRYPOINT ["/usr/local/bin/puller"]