# Fault Tolerance
## With Envoy

# Benchmark

Reliable server benchmark

```shell
$ go install github.com/bojand/ghz/cmd/ghz
$ ghz -call omgrpc.users.v2.Users/Get -d '{"name": "foo"}' -insecure -proto ./protos/users/v2/users.proto 0.0.0.0:80

Summary:
  Count:	200
  Total:	184.49 ms
  Slowest:	51.26 ms
  Fastest:	22.23 ms
  Average:	37.43 ms
  Requests/sec:	1084.05

Response time histogram:
  22.229 [1]	|∎
  25.131 [26]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  28.034 [15]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  30.937 [4]	|∎∎∎∎
  33.840 [0]	|
  36.743 [33]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  39.646 [41]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  42.549 [20]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  45.452 [18]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  48.355 [24]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  51.258 [18]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎

Latency distribution:
  10% in 23.39 ms
  25% in 34.39 ms
  50% in 38.15 ms
  75% in 43.54 ms
  90% in 47.31 ms
  95% in 49.56 ms
  99% in 51.15 ms
Status code distribution:
  [OK]	200 responses
```

```
$ ghz -call omgrpc.users.v2.Users/Get -d '{"name": "foo"}' -insecure -O json -proto ./protos/users/v2/users.proto 0.0.0.0:80 | jq '.statusCodeDistribution.OK / .count * 100.0'
100
```

100% success rate.

# Fault Injection

Add 10% error rate:

```go
func (s *Server) List(ctx context.Context, r *widgets.ListRequest) (*widgets.ListResponse, error) {
	if rand.Float64() < 0.10 {
		return nil, status.Errorf(codes.Unavailable, "random failure")
	}
	...
}
```

See ~90% success rate:

```
$ ghz -call omgrpc.users.v2.Users/Get -d '{"name": "foo"}' -insecure -n 1000 -O json -proto ./protos/users/v2/users.proto 0.0.0.0:80 | jq '.statusCodeDistribution.OK / .count * 100.0'
88.8
```

Add retry middleware:

```go
func serve(config config) error {
	conn, err := grpc.Dial(
		config.WidgetsAddr,
		grpc.WithAuthority("widgets-v1"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(3),
		)),
	)
	...
}
```

See improved success rate:

```
$ ghz -call omgrpc.users.v2.Users/Get -d '{"name": "foo"}' -insecure -n 1000 -O json -proto ./protos/users/v2/users.proto 0.0.0.0:80 | jq '.statusCodeDistribution.OK / .count * 100.0'
99.8
```

# Envoy Retries

```yaml
static_resources:
  - name: egress
    filter_chains:
    - filters:
      - name: envoy.http_connection_manager
        config:
          stat_prefix: egress_http
          route_config:
            name: local_route
            virtual_hosts:
            - name: widgets-v1
              domains: ["widgets-v1"]
              include_request_attempt_count: true
              routes:
              - match:
                  prefix: "/"
                route:
                  cluster: widgets-v1
                  retry_policy:
                    retry_on: unavailable
                    num_retries: 5
          http_filters:
          - name: envoy.router
```

```shell
$ ghz -call omgrpc.users.v2.Users/Get -d '{"name": "foo"}' -insecure -n 1000 -O json -proto ./protos/users/v2/users.proto 0.0.0.0:80 | jq '.statusCodeDistribution.OK / .count * 100.0'
96.6
```

Better than 90% but less than grpc middleware. What gives?

# Circuit Breaking

```yaml
static_resources:
  clusters:
  - name: widgets-v1
    circuit_breakers:
      thresholds:
        - max_retries: 10
```

```shell
$ ghz -call omgrpc.users.v2.Users/Get -d '{"name": "foo"}' -insecure -n 1000 -O json -proto ./protos/users/v2/users.proto 0.0.0.0:80 | jq '.statusCodeDistribution.OK / .count * 100.0'
99.9
```

http://localhost:9090/graph?g0.range_input=1h&g0.expr=envoy_cluster_upstream_rq_retry_overflow&g0.tab=0
