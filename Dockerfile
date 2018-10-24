FROM envoyproxy/envoy:latest

COPY configs/*         /etc/envoy/
COPY bin/envoy.sh      /usr/local/bin/
COPY bin/linux_amd64/* /usr/local/bin/

ENTRYPOINT ["/usr/bin/dumb-init", "--", "/usr/local/bin/envoy.sh"]