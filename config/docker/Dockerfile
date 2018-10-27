FROM envoyproxy/envoy:latest

COPY bin/envoy.sh      /usr/local/bin/
COPY bin/linux_amd64/* /usr/local/bin/
COPY config/envoy/*    /etc/envoy/
COPY gen/pb/           /etc/

ENTRYPOINT ["/usr/bin/dumb-init", "--", "/usr/local/bin/envoy.sh"]