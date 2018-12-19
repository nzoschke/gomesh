CMDS = $(wildcard cmd/*)
BINS = $(CMDS:cmd/%=bin/linux_amd64/%)
bins: $(BINS)
$(BINS): bin/linux_amd64/%: cmd/%/main.go $(shell find . -name '*.go')
	GOOS=linux GOARCH=amd64 go build -o $@ $<

clean:
	rm bin/linux_amd64/*

COMPOSE_CMD = docker-compose -p gomesh
COMPOSE_FILES = -f config/docker/compose-api.yaml \
	-f config/docker/compose-mesh.yaml \
	-f config/docker/compose-proxy.yaml

configs/sidecar.yaml: cmd/envoy-cfg/main.go
	go run $< /tmp/sidecar.yaml
	docker run \
		-v/tmp/sidecar.yaml:/tmp/sidecar.yaml \
		envoyproxy/envoy:latest \
		envoy -c /tmp/sidecar.yaml --mode validate
	# mv /tmp/sidecar.yaml configs/sidecar.yaml

compose-build:
	$(COMPOSE_CMD) $(COMPOSE_FILES) build

compose-down:
	$(COMPOSE_CMD) $(COMPOSE_FILES) down

compose-api:
	make -j bins
	$(COMPOSE_CMD) -f config/docker/compose-api.yaml   up --abort-on-container-exit

compose-mesh:
	make -j bins
	$(COMPOSE_CMD) -f config/docker/compose-mesh.yaml  up --abort-on-container-exit

compose-proxy:
	make -j bins
	$(COMPOSE_CMD) -f config/docker/compose-proxy.yaml up --abort-on-container-exit
