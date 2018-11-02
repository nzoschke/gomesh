FROM envoyproxy/envoy:latest

COPY bin/envoy.sh      /usr/local/bin/
COPY bin/linux_amd64/  /usr/local/sbin/
COPY config/envoy/     /etc/envoy/
COPY gen/pb/           /etc/pb/

ENTRYPOINT ["/usr/bin/dumb-init", "--", "envoy.sh"]