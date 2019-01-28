#!/bin/bash
set -ex
# First start gateway, e.g. `WIDGETS_V2_ERROR_RATE=2 make dc-up-gateway`

# get TOKEN
eval $(docker logs --since 20m gomesh_auth-v2alpha_1 2>&1 | grep -o 'TOKEN=.*' | tail -1)

# make 10 req/s for 5m
hey -H "Authorization: Bearer $TOKEN" -c 1  -q 10 -z 5m http://localhost/v2/orgs/myorg/users/myuser
sleep 180

# make 30 req/s for 5m
hey -H "Authorization: Bearer $TOKEN" -c 3  -q 10 -z 5m http://localhost/v2/orgs/myorg/users/myuser
sleep 180

# make 200 req/s for 5m
hey -H "Authorization: Bearer $TOKEN" -c 20 -q 10 -z 5m http://localhost/v2/orgs/myorg/users/myuser
