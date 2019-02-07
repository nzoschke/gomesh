```yaml
---
descriptors:
  - key: remote_address
    rate_limit:
      requests_per_unit: 25
      unit: second
domain: gateway
```

```shell
# start the gateway with the widgets service erroring 2% of the time
$ WIDGETS_V2_ERROR_RATE=2 make dc-up-gateway

# 10 req/s for 5m
$ hey -H "Authorization: Bearer $TOKEN" -c 1  -q 10 -z 5m http://localhost/v2/orgs/myorg/users/myuser
Summary:
  Total:        300.0279 secs
  Slowest:      1.0350 secs
  Fastest:      0.0137 secs
  Average:      0.0280 secs
  Requests/sec: 9.9391
  
  Total data:   771820 bytes
  Size/request: 258 bytes

Response time histogram:
  0.014 [1]     |
  0.116 [2979]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.218 [0]     |
  0.320 [0]     |
  0.422 [0]     |
  0.524 [0]     |
  0.627 [0]     |
  0.729 [0]     |
  0.831 [0]     |
  0.933 [0]     |
  1.035 [2]     |

Status code distribution:
  [200] 2980 responses
  [403] 1    responses
  [503] 1    responses


# 30 req/s for 5m
$ hey -H "Authorization: Bearer $TOKEN" -c 3  -q 10 -z 5m http://localhost/v2/orgs/myorg/users/myuser
Summary:
  Total:        300.0979 secs
  Slowest:      1.0308 secs
  Fastest:      0.0042 secs
  Average:      0.0426 secs
  Requests/sec: 29.9002
  
  Total data:   1937061 bytes
  Size/request: 215 bytes

Response time histogram:
  0.004 [1]     |
  0.107 [8963]  |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.209 [6]     |
  0.312 [0]     |
  0.415 [0]     |
  0.517 [0]     |
  0.620 [0]     |
  0.723 [0]     |
  0.825 [0]     |
  0.928 [0]     |
  1.031 [3]     |

Status code distribution:
  [200] 7479 responses
  [403] 3    responses
  [429] 1487 responses
  [503] 4    responses


# 200 req/s for 5m
$ hey -H "Authorization: Bearer $TOKEN" -c 20 -q 10 -z 5m http://localhost/v2/orgs/myorg/users/myuser
Summary:
  Total:        300.1524 secs
  Slowest:      0.6477 secs
  Fastest:      0.0027 secs
  Average:      0.0691 secs
  Requests/sec: 169.8137
  
  Total data:   1648017 bytes
  Size/request: 32 bytes

Response time histogram:
  0.003 [1]     |
  0.067 [41077] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.132 [3217]  |■■■
  0.196 [1300]  |■
  0.261 [1578]  |■■
  0.325 [1656]  |■■
  0.390 [1445]  |■
  0.454 [534]   |■
  0.519 [150]   |
  0.583 [9]     |
  0.648 [3]     |

Status code distribution:
  [200] 6363  responses
  [403] 1158  responses
  [429] 43447 responses
  [503] 2     responses
```

# Results - Good

We see rate limiting is working as expected. 25 req/sec * 60sec/m * 5m = 7500 responses.

# Results - Bad - 403s

But what about the 403s? We see a cascading failure with higher concurrency.

Jaeger `http.status_code=403` -> 1000 Traces

TODO: SCREENSHOT

$ docker logs gomesh_auth-v2alpha-sidecar_1 2>&1 | grep 951d4082-e171-97c0-8c02-0ba623caea3c
listener=ingress start_time=2019-01-26T23:08:48.108Z req_method=POST req_path=/envoy.service.auth.v2alpha.Authorization/Check protocol=HTTP/2 response_code=0 response_flags=DC bytes_received=540 bytes_sent=0 duration=180 resp_x_envoy_upstream_service_time=- req_x_envoy_original_path=- req_x_forwarded_for=172.29.0.17 req_user_agent="-" req_x_request_id=951d4082-e171-97c0-8c02-0ba623caea3c req_uber_trace_id=5656d560be251a99:e5624eb658b79c6e:2a5b9f7d2ece9308:1 req_authority=auth-v2alpha-sidecar upstream_host=172.29.0.5:8002 upstream_cluster=local2 resp_grpc_status=- resp_grpc_message="-" trailer_grpc_status=- trailer_grpc_message="-"

listener=egress  start_time=2019-01-26T23:08:48.236Z req_method=POST req_path=/oauth2/introspect protocol=HTTP/1.1 response_code=200 response_flags=- bytes_received=100 bytes_sent=151 duration=10 resp_x_envoy_upstream_service_time=9 req_x_envoy_original_path=- req_x_forwarded_for=- req_user_agent="Swagger-Codegen/1.0.0/go" req_x_request_id=951d4082-e171-97c0-8c02-0ba623caea3c req_uber_trace_id=5656d560be251a99:b198657f61e87d37:e5624eb658b79c6e:1 req_authority=hydra-sidecar upstream_host=172.29.0.16:10000 upstream_cluster=hydra-sidecar resp_grpc_status=- resp_grpc_message="-" trailer_grpc_status=- trailer_grpc_message="-"

proxy response / SLA

rate(envoy_cluster_upstream_rq{envoy_response_code!="200", envoy_response_code!='429'}[5m])


# Config Improvements


```diff
@@ -78,6 +78,7 @@ static_resources:
                  - config:
                       grpc_service:
                         envoy_grpc:
                           cluster_name: auth-v2alpha-sidecar
+                        timeout: 0.5s
                    name: envoy.ext_authz

@@ -157,6 +157,7 @@ static_resources:
                           route:
                             cluster_header: :authority
                             retry_policy:
+                              num_retries: 3
                               retry_on: cancelled,deadline-exceeded,internal,resource-exhausted,unavailable
```


```shell
$ hey -H "Authorization: Bearer $TOKEN" -c 20 -q 10 -z 5m http://localhost/v2/orgs/myorg/users/myuser

Summary:
  Total:        300.0479 secs
  Slowest:      0.6900 secs
  Fastest:      0.0023 secs
  Average:      0.0753 secs
  Requests/sec: 163.3072
  
  Total data:   1948975 bytes
  Size/request: 39 bytes

Response time histogram:
  0.002 [1]     |
  0.071 [39821] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.140 [2246]  |■■
  0.209 [1326]  |■
  0.277 [1717]  |■■
  0.346 [1567]  |■■
  0.415 [1397]  |■
  0.484 [708]   |■
  0.552 [166]   |
  0.621 [45]    |
  0.690 [6]     |

Status code distribution:
  [200] 7525  responses
  [429] 41475 responses


---
Where are the 403s coming from?

Jaeger `http.status_code=403` -> 26 Traces

async auth-v2alpha-sidecar egress
Service:proxy Duration:199.69ms Start Time:27.55ms
Tags
    upstream_cluster="auth-v2alpha-sidecar"
    component="proxy"
    grpc.status_code="14"
    error="true"

ingress
Service:auth-v2alpha Duration:154.6ms Start Time:52.78ms
Tags
    http.status_code="0"
    error="true"
    guid:x-request-id=39a72121-0614-91bf-89bf-2b4a80f98311

Unavailable Code = 14


```prom
envoy_cluster_upstream_rq
envoy_cluster_upstream_rq{envoy_response_code!="200"}
envoy_cluster_upstream_rq{envoy_response_code != "200", envoy_response_code != "429"}
```

$ docker-compose -f config/docker/compose-gateway.yaml --project-directory . logs | grep 39a72121-0614-91bf-89bf-2b4a80f98311
auth-v2alpha-sidecar_1  | listener=ingress start_time=2019-01-26T17:01:50.035Z req_method=POST req_path=/envoy.service.auth.v2alpha.Authorization/Check protocol=HTTP/2 response_code=0 response_flags=DC bytes_received=541 bytes_sent=0 duration=197 resp_x_envoy_upstream_service_time=- req_x_envoy_original_path=- req_x_forwarded_for=172.21.0.17 req_user_agent="-" req_x_request_id=39a72121-0614-91bf-89bf-2b4a80f98311 req_uber_trace_id=7ff89654474795a5:136188289d78fabe:fdd9cf9b9f11069d:1 req_authority=auth-v2alpha-sidecar upstream_host=172.21.0.4:8002 upstream_cluster=local2 resp_grpc_status=- resp_grpc_message="-" trailer_grpc_status=- trailer_grpc_message="-"

DC = Downstream connection termination

```
envoy_cluster_upstream_rq{envoy_cluster_name="auth-v2alpha-sidecar",envoy_response_code="504",instance="172.19.0.17:9901",job="proxy"}
```

Why doesn't this work?

```prom
envoy_cluster_grpc{envoy_grpc_code='14'}
```

```prom
envoy_cluster_upstream_rq_xx{envoy_response_code_class="4"}
envoy_cluster_upstream_rq_xx{envoy_cluster_name="users-v2-sidecar",envoy_response_code_class="4",instance="172.19.0.17:9901",job="proxy"}

rate(envoy_cluster_upstream_rq_xx{envoy_response_code_class="4"}[5m])
```