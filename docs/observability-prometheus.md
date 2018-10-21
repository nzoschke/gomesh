# Observability
## With Envoy and Prometheus

https://www.envoyproxy.io/docs/envoy/latest/configuration/cluster_manager/cluster_stats
https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/stats

docker logs -f 573d627ab8b1 2>/dev/null | grep -o '{.*}' --line-buffered | jq

# SLA

Upstream response rates
  - 400
  - 500