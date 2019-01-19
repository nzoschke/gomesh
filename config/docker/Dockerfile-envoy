FROM envoyproxy/envoy:v1.9.0

ARG JAEGER_VERSION=v0.4.2

RUN wget -qO /usr/local/lib/libjaegertracing_plugin.so \
  https://github.com/jaegertracing/jaeger-client-cpp/releases/download/$JAEGER_VERSION/libjaegertracing_plugin.linux_amd64.so

COPY bin/envoy.sh /usr/local/bin/
